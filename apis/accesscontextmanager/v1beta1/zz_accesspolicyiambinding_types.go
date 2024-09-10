// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessPolicyIAMBindingConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type AccessPolicyIAMBindingConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type AccessPolicyIAMBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type AccessPolicyIAMBindingInitParameters struct {
	Condition *AccessPolicyIAMBindingConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type AccessPolicyIAMBindingObservation struct {
	Condition *AccessPolicyIAMBindingConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type AccessPolicyIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition *AccessPolicyIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// AccessPolicyIAMBindingSpec defines the desired state of AccessPolicyIAMBinding
type AccessPolicyIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPolicyIAMBindingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider AccessPolicyIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// AccessPolicyIAMBindingStatus defines the observed state of AccessPolicyIAMBinding.
type AccessPolicyIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPolicyIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AccessPolicyIAMBinding is the Schema for the AccessPolicyIAMBindings API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AccessPolicyIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.members) || (has(self.initProvider) && has(self.initProvider.members))",message="spec.forProvider.members is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   AccessPolicyIAMBindingSpec   `json:"spec"`
	Status AccessPolicyIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPolicyIAMBindingList contains a list of AccessPolicyIAMBindings
type AccessPolicyIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPolicyIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	AccessPolicyIAMBinding_Kind             = "AccessPolicyIAMBinding"
	AccessPolicyIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPolicyIAMBinding_Kind}.String()
	AccessPolicyIAMBinding_KindAPIVersion   = AccessPolicyIAMBinding_Kind + "." + CRDGroupVersion.String()
	AccessPolicyIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(AccessPolicyIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPolicyIAMBinding{}, &AccessPolicyIAMBindingList{})
}
