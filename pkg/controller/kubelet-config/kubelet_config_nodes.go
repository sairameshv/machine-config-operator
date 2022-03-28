package kubeletconfig

import (
	"context"
	"fmt"
	"k8s.io/kubelet/config/v1beta1"
	"reflect"
	"time"

	"github.com/clarketm/json"
	"github.com/golang/glog"
	osev1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/retry"

	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	ctrlcommon "github.com/openshift/machine-config-operator/pkg/controller/common"
	"github.com/openshift/machine-config-operator/pkg/version"
)

func (ctrl *Controller) nodeWorker() {
	for ctrl.processNextNodeWorkItem() {
	}
}

func (ctrl *Controller) processNextNodeWorkItem() bool {
	key, quit := ctrl.nodeQueue.Get()
	if quit {
		return false
	}
	defer ctrl.nodeQueue.Done(key)

	err := ctrl.syncNodeHandler(key.(string))
	ctrl.handleNodeErr(err, key)
	return true
}

func (ctrl *Controller) handleNodeErr(err error, key interface{}) {
	if err == nil {
		ctrl.nodeQueue.Forget(key)
		return
	}

	if ctrl.nodeQueue.NumRequeues(key) < maxRetries {
		glog.V(2).Infof("Error syncing node configuration %v: %v", key, err)
		ctrl.nodeQueue.AddRateLimited(key)
		return
	}

	utilruntime.HandleError(err)
	glog.V(2).Infof("Dropping node config %q out of the queue: %v", key, err)
	ctrl.nodeQueue.Forget(key)
	ctrl.nodeQueue.AddAfter(key, 1*time.Minute)
}

func (ctrl *Controller) syncNodeHandler(key string) error {
	startTime := time.Now()
	glog.V(4).Infof("Started syncing node handler %q (%v)", key, startTime)
	defer func() {
		glog.V(4).Infof("Finished syncing node handler %q (%v)", key, time.Since(startTime))
	}()

	// Fetch the Node
	node, err := ctrl.nodeLister.Get(ctrlcommon.ClusterNodeInstanceName)
	if errors.IsNotFound(err) {
		glog.V(2).Infof("Node configuration %v is missing, using default", key)
		node = createNewDefaultNodeconfig()
	} else if err != nil {
		glog.V(2).Infof("%v", err)
		err := fmt.Errorf("could not fetch Node: %v", err)
		return err
	} else if err == nil {
		// checking if the Node spec is empty and accordingly returning from here.
		if reflect.DeepEqual(node.Spec, osev1.NodeSpec{}) {
			glog.V(2).Info("empty Node resource found")
			return nil
		}
	}

	// Find all MachineConfigPools
	mcpPools, err := ctrl.mcpLister.List(labels.Everything())
	if err != nil {
		return err
	}

	for _, pool := range mcpPools {
		role := pool.Name
		// the configuration change will be applied only on the worker nodes.
		if role == "master" {
			continue
		}

		// Get KubeletConfig
		managedKey, err := getManagedNodeKey(pool, ctrl.client)
		if err != nil {
			return err
		}
		kc, err := ctrl.client.MachineconfigurationV1().KubeletConfigs().Get(context.TODO(), managedKey, metav1.GetOptions{})
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		isNotFound := errors.IsNotFound(err)
		if isNotFound {
			ignConfig := ctrlcommon.NewIgnConfig()
			kc, err = ctrlcommon.KubeletConfigFromIgnConfig(role, managedKey, ignConfig)
			if err != nil {
				return err
			}
		}

		// updating the node status based on the latest kubelet config CR status
		err = ctrl.updateNodestatus(kc, node)
		if err != nil {
			return err
		}
		// restricting the sync if the node status condition is still Progressing/Degraded
		nodeCondition := fetchNodeconditionstatus(node)
		if nodeCondition == osev1.ConditionFalse {
			return fmt.Errorf("unable to modify the kubelet configuration on the node, node condition status not ready")
		}

		var originalKubeConfig = new(v1beta1.KubeletConfiguration)

		// updating the kubelet configuration with the Node specific configuration.
		err = updateOriginalKubeConfigwithNodeConfig(node, originalKubeConfig)
		if err != nil {
			return err
		}

		byteData, err := json.Marshal(originalKubeConfig)
		if err != nil {
			return err
		}
		if byteData == nil {
			continue
		}
		kc.Spec.KubeletConfig.Raw = byteData
		kc.ObjectMeta.Annotations = map[string]string{
			ctrlcommon.GeneratedByControllerVersionAnnotationKey: version.Hash,
		}
		// Create or Update, on conflict retry
		if err := retry.RetryOnConflict(updateBackoff, func() error {
			var err error
			if isNotFound {
				_, err = ctrl.client.MachineconfigurationV1().KubeletConfigs().Create(context.TODO(), kc, metav1.CreateOptions{})
			} else {
				_, err = ctrl.client.MachineconfigurationV1().KubeletConfigs().Update(context.TODO(), kc, metav1.UpdateOptions{})
			}
			return err
		}); err != nil {
			return fmt.Errorf("Could not Create/Update KubeletConfig: %v", err)
		}
		glog.Infof("Applied Node configuration %v on MachineConfigPool %v", key, pool.Name)

		// updating the node status based on the latest kubelet config CR status
		err = ctrl.updateNodestatus(kc, node)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctrl *Controller) enqueueNode(node *osev1.Node) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(node)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Couldn't get key for object %#v: %v", node, err))
		return
	}
	ctrl.nodeQueue.Add(key)
}

func (ctrl *Controller) updateNode(old, cur interface{}) {
	oldNode := old.(*osev1.Node)
	newNode := cur.(*osev1.Node)
	if !reflect.DeepEqual(oldNode.Spec, newNode.Spec) {
		// skipping the update in case of the Worker-Latency-Profile type transition from "Default" to "LowUpdateSlowReaction" and vice-versa
		// (TODO) Ideally the user request has to be honoured, the transition need to be from Default -> Medium -> Low or vice-versa.
		// Restricting the request for now until this process is automated in future.
		if (oldNode.Spec.WorkerLatencyProfile == osev1.DefaultUpdateDefaultReaction && newNode.Spec.WorkerLatencyProfile == osev1.LowUpdateSlowReaction) || (oldNode.Spec.WorkerLatencyProfile == osev1.LowUpdateSlowReaction && newNode.Spec.WorkerLatencyProfile == osev1.DefaultUpdateDefaultReaction) {
			glog.Infof("Skipping the Update Node event, name: %s, transition not allowed from old WorkerLatencyProfile: %s to new WorkerLatencyProfile: %s", newNode.Name, oldNode.Spec.WorkerLatencyProfile, newNode.Spec.WorkerLatencyProfile)
			return
		}
		glog.V(4).Infof("Update Node event, name: %s", newNode.Name)
		ctrl.enqueueNode(newNode)
	}
}

func (ctrl *Controller) addNode(obj interface{}) {
	node := obj.(*osev1.Node)
	glog.V(4).Infof("Add Node event %s", node.Name)
	ctrl.enqueueNode(node)
}

func (ctrl *Controller) deleteNode(obj interface{}) {
	node, ok := obj.(*osev1.Node)
	if !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Couldn't get object from tombstone %#v", obj))
			return
		}
		node, ok = tombstone.Obj.(*osev1.Node)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Tombstone contained object that is not a KubeletConfig %#v", obj))
			return
		}
	}
	glog.V(4).Infof("Deleted Node %s and restored default config", node.Name)
}

// fetchNodeconditionstatus fetches the condition status written by the MCO from the provided node object's status
func fetchNodeconditionstatus(node *osev1.Node) osev1.ConditionStatus {
	if node != nil {
		if len(node.Status.WorkerLatencyProfileStatus.Conditions) > 0 {
			for _, condition := range node.Status.WorkerLatencyProfileStatus.Conditions {
				if condition.Owner == osev1.MachineConfigOperator {
					return condition.Status
				}
			}
		}
	}
	return osev1.ConditionUnknown
}

// updateNodestatus updates the status of the node object based on the machineconfigpool status
func (ctrl *Controller) updateNodestatus(kc *mcfgv1.KubeletConfig, node *osev1.Node) error {
	if kc == nil || node == nil {
		return fmt.Errorf("unable to update the node status, incomplete data, kubelet config: %v, node: %v", kc, node)
	}
	var nodeCondition osev1.WorkerLatencyStatusCondition

	nodeCondition.Owner = osev1.MachineConfigOperator
	nodeCondition.LastTransitionTime = metav1.Now()
	nodeCondition.Status = osev1.ConditionUnknown

	// updating the node CR status based on the latest KubeletConfig CR status
	if len(kc.Status.Conditions) > 0 {
		switch kc.Status.Conditions[len(kc.Status.Conditions)-1].Type {
		case mcfgv1.KubeletConfigFailure:
			nodeCondition.Status = osev1.ConditionFalse
			nodeCondition.Type = osev1.WorkerLatencyProfileDegraded
		case mcfgv1.KubeletConfigSuccess:
			nodeCondition.Status = osev1.ConditionTrue
			nodeCondition.Type = osev1.WorkerLatencyProfileComplete
			// (TODO) further cases can be added in future based on the Kubelet Config CR Status condition types
		default:
			nodeCondition.Status = osev1.ConditionUnknown
		}
		nodeCondition.Reason = kc.Status.Conditions[len(kc.Status.Conditions)-1].Reason
		nodeCondition.Message = kc.Status.Conditions[len(kc.Status.Conditions)-1].Message
	}
	var (
		index     int
		isPresent bool
	)
	for index = range node.Status.WorkerLatencyProfileStatus.Conditions {
		if node.Status.WorkerLatencyProfileStatus.Conditions[index].Owner == osev1.MachineConfigOperator {
			isPresent = true
			break
		}
	}
	if isPresent {
		node.Status.WorkerLatencyProfileStatus.Conditions[index] = *nodeCondition.DeepCopy()
	} else {
		node.Status.WorkerLatencyProfileStatus.Conditions = append(node.Status.WorkerLatencyProfileStatus.Conditions, nodeCondition)
	}
	_, err := ctrl.configClient.ConfigV1().Nodes().UpdateStatus(context.TODO(), node, metav1.UpdateOptions{})
	return err
}
