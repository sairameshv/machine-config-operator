// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	machineconfigurationv1alpha1 "github.com/openshift/api/machineconfiguration/v1alpha1"
	internal "github.com/openshift/client-go/machineconfiguration/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// MachineConfigNodeApplyConfiguration represents a declarative configuration of the MachineConfigNode type for use
// with apply.
type MachineConfigNodeApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *MachineConfigNodeSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *MachineConfigNodeStatusApplyConfiguration `json:"status,omitempty"`
}

// MachineConfigNode constructs a declarative configuration of the MachineConfigNode type for use with
// apply.
func MachineConfigNode(name string) *MachineConfigNodeApplyConfiguration {
	b := &MachineConfigNodeApplyConfiguration{}
	b.WithName(name)
	b.WithKind("MachineConfigNode")
	b.WithAPIVersion("machineconfiguration.openshift.io/v1alpha1")
	return b
}

// ExtractMachineConfigNode extracts the applied configuration owned by fieldManager from
// machineConfigNode. If no managedFields are found in machineConfigNode for fieldManager, a
// MachineConfigNodeApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// machineConfigNode must be a unmodified MachineConfigNode API object that was retrieved from the Kubernetes API.
// ExtractMachineConfigNode provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractMachineConfigNode(machineConfigNode *machineconfigurationv1alpha1.MachineConfigNode, fieldManager string) (*MachineConfigNodeApplyConfiguration, error) {
	return extractMachineConfigNode(machineConfigNode, fieldManager, "")
}

// ExtractMachineConfigNodeStatus is the same as ExtractMachineConfigNode except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractMachineConfigNodeStatus(machineConfigNode *machineconfigurationv1alpha1.MachineConfigNode, fieldManager string) (*MachineConfigNodeApplyConfiguration, error) {
	return extractMachineConfigNode(machineConfigNode, fieldManager, "status")
}

func extractMachineConfigNode(machineConfigNode *machineconfigurationv1alpha1.MachineConfigNode, fieldManager string, subresource string) (*MachineConfigNodeApplyConfiguration, error) {
	b := &MachineConfigNodeApplyConfiguration{}
	err := managedfields.ExtractInto(machineConfigNode, internal.Parser().Type("com.github.openshift.api.machineconfiguration.v1alpha1.MachineConfigNode"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(machineConfigNode.Name)

	b.WithKind("MachineConfigNode")
	b.WithAPIVersion("machineconfiguration.openshift.io/v1alpha1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithKind(value string) *MachineConfigNodeApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithAPIVersion(value string) *MachineConfigNodeApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithName(value string) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithGenerateName(value string) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithNamespace(value string) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithUID(value types.UID) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithResourceVersion(value string) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithGeneration(value int64) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithCreationTimestamp(value metav1.Time) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *MachineConfigNodeApplyConfiguration) WithLabels(entries map[string]string) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *MachineConfigNodeApplyConfiguration) WithAnnotations(entries map[string]string) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *MachineConfigNodeApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *MachineConfigNodeApplyConfiguration) WithFinalizers(values ...string) *MachineConfigNodeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *MachineConfigNodeApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithSpec(value *MachineConfigNodeSpecApplyConfiguration) *MachineConfigNodeApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *MachineConfigNodeApplyConfiguration) WithStatus(value *MachineConfigNodeStatusApplyConfiguration) *MachineConfigNodeApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *MachineConfigNodeApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}
