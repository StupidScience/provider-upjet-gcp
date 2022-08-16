/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConditionObservation struct {
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type OrganizationIAMBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OrganizationIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// The numeric ID of the organization in which you want to manage the audit logging config.
	// +kubebuilder:validation:Required
	OrgID *string `json:"orgId" tf:"org_id,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// OrganizationIAMBindingSpec defines the desired state of OrganizationIAMBinding
type OrganizationIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrganizationIAMBindingParameters `json:"forProvider"`
}

// OrganizationIAMBindingStatus defines the observed state of OrganizationIAMBinding.
type OrganizationIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrganizationIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationIAMBinding is the Schema for the OrganizationIAMBindings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type OrganizationIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationIAMBindingSpec   `json:"spec"`
	Status            OrganizationIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationIAMBindingList contains a list of OrganizationIAMBindings
type OrganizationIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	OrganizationIAMBinding_Kind             = "OrganizationIAMBinding"
	OrganizationIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrganizationIAMBinding_Kind}.String()
	OrganizationIAMBinding_KindAPIVersion   = OrganizationIAMBinding_Kind + "." + CRDGroupVersion.String()
	OrganizationIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(OrganizationIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&OrganizationIAMBinding{}, &OrganizationIAMBindingList{})
}
