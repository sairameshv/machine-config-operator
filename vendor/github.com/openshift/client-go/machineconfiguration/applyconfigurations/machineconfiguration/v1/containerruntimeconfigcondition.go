// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	machineconfigurationv1 "github.com/openshift/api/machineconfiguration/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerRuntimeConfigConditionApplyConfiguration represents a declarative configuration of the ContainerRuntimeConfigCondition type for use
// with apply.
type ContainerRuntimeConfigConditionApplyConfiguration struct {
	Type               *machineconfigurationv1.ContainerRuntimeConfigStatusConditionType `json:"type,omitempty"`
	Status             *corev1.ConditionStatus                                           `json:"status,omitempty"`
	LastTransitionTime *metav1.Time                                                      `json:"lastTransitionTime,omitempty"`
	Reason             *string                                                           `json:"reason,omitempty"`
	Message            *string                                                           `json:"message,omitempty"`
}

// ContainerRuntimeConfigConditionApplyConfiguration constructs a declarative configuration of the ContainerRuntimeConfigCondition type for use with
// apply.
func ContainerRuntimeConfigCondition() *ContainerRuntimeConfigConditionApplyConfiguration {
	return &ContainerRuntimeConfigConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ContainerRuntimeConfigConditionApplyConfiguration) WithType(value machineconfigurationv1.ContainerRuntimeConfigStatusConditionType) *ContainerRuntimeConfigConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ContainerRuntimeConfigConditionApplyConfiguration) WithStatus(value corev1.ConditionStatus) *ContainerRuntimeConfigConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *ContainerRuntimeConfigConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *ContainerRuntimeConfigConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ContainerRuntimeConfigConditionApplyConfiguration) WithReason(value string) *ContainerRuntimeConfigConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ContainerRuntimeConfigConditionApplyConfiguration) WithMessage(value string) *ContainerRuntimeConfigConditionApplyConfiguration {
	b.Message = &value
	return b
}
