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

type ServicePerimeterEgressPolicyEgressFromInitParameters struct {

	// A list of identities that are allowed access through this EgressPolicy.
	// Should be in the format of email address. The email address should
	// represent individual user or service account only.
	Identities []*string `json:"identities,omitempty" tf:"identities,omitempty"`

	// Specifies the type of identities that are allowed access to outside the
	// perimeter. If left unspecified, then members of identities field will
	// be allowed access.
	// Possible values are: ANY_IDENTITY, ANY_USER_ACCOUNT, ANY_SERVICE_ACCOUNT.
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type,omitempty"`

	// Whether to enforce traffic restrictions based on sources field. If the sources field is non-empty, then this field must be set to SOURCE_RESTRICTION_ENABLED.
	// Possible values are: SOURCE_RESTRICTION_UNSPECIFIED, SOURCE_RESTRICTION_ENABLED, SOURCE_RESTRICTION_DISABLED.
	SourceRestriction *string `json:"sourceRestriction,omitempty" tf:"source_restriction,omitempty"`

	// Sources that this EgressPolicy authorizes access from.
	// Structure is documented below.
	Sources []ServicePerimeterEgressPolicyEgressFromSourcesInitParameters `json:"sources,omitempty" tf:"sources,omitempty"`
}

type ServicePerimeterEgressPolicyEgressFromObservation struct {

	// A list of identities that are allowed access through this EgressPolicy.
	// Should be in the format of email address. The email address should
	// represent individual user or service account only.
	Identities []*string `json:"identities,omitempty" tf:"identities,omitempty"`

	// Specifies the type of identities that are allowed access to outside the
	// perimeter. If left unspecified, then members of identities field will
	// be allowed access.
	// Possible values are: ANY_IDENTITY, ANY_USER_ACCOUNT, ANY_SERVICE_ACCOUNT.
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type,omitempty"`

	// Whether to enforce traffic restrictions based on sources field. If the sources field is non-empty, then this field must be set to SOURCE_RESTRICTION_ENABLED.
	// Possible values are: SOURCE_RESTRICTION_UNSPECIFIED, SOURCE_RESTRICTION_ENABLED, SOURCE_RESTRICTION_DISABLED.
	SourceRestriction *string `json:"sourceRestriction,omitempty" tf:"source_restriction,omitempty"`

	// Sources that this EgressPolicy authorizes access from.
	// Structure is documented below.
	Sources []ServicePerimeterEgressPolicyEgressFromSourcesObservation `json:"sources,omitempty" tf:"sources,omitempty"`
}

type ServicePerimeterEgressPolicyEgressFromParameters struct {

	// A list of identities that are allowed access through this EgressPolicy.
	// Should be in the format of email address. The email address should
	// represent individual user or service account only.
	// +kubebuilder:validation:Optional
	Identities []*string `json:"identities,omitempty" tf:"identities,omitempty"`

	// Specifies the type of identities that are allowed access to outside the
	// perimeter. If left unspecified, then members of identities field will
	// be allowed access.
	// Possible values are: ANY_IDENTITY, ANY_USER_ACCOUNT, ANY_SERVICE_ACCOUNT.
	// +kubebuilder:validation:Optional
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type,omitempty"`

	// Whether to enforce traffic restrictions based on sources field. If the sources field is non-empty, then this field must be set to SOURCE_RESTRICTION_ENABLED.
	// Possible values are: SOURCE_RESTRICTION_UNSPECIFIED, SOURCE_RESTRICTION_ENABLED, SOURCE_RESTRICTION_DISABLED.
	// +kubebuilder:validation:Optional
	SourceRestriction *string `json:"sourceRestriction,omitempty" tf:"source_restriction,omitempty"`

	// Sources that this EgressPolicy authorizes access from.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Sources []ServicePerimeterEgressPolicyEgressFromSourcesParameters `json:"sources,omitempty" tf:"sources,omitempty"`
}

type ServicePerimeterEgressPolicyEgressFromSourcesInitParameters struct {

	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`
}

type ServicePerimeterEgressPolicyEgressFromSourcesObservation struct {

	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`
}

type ServicePerimeterEgressPolicyEgressFromSourcesParameters struct {

	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	// +kubebuilder:validation:Optional
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToInitParameters struct {

	// A list of external resources that are allowed to be accessed. A request
	// matches if it contains an external resource in this list (Example:
	// s3://bucket/path). Currently '*' is not allowed.
	ExternalResources []*string `json:"externalResources,omitempty" tf:"external_resources,omitempty"`

	// A list of ApiOperations that this egress rule applies to. A request matches
	// if it contains an operation/service in this list.
	// Structure is documented below.
	Operations []ServicePerimeterEgressPolicyEgressToOperationsInitParameters `json:"operations,omitempty" tf:"operations,omitempty"`

	// A list of resources, currently only projects in the form
	// projects/<projectnumber>, that match this to stanza. A request matches
	// if it contains a resource in this list. If * is specified for resources,
	// then this EgressTo rule will authorize access to all resources outside
	// the perimeter.
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToObservation struct {

	// A list of external resources that are allowed to be accessed. A request
	// matches if it contains an external resource in this list (Example:
	// s3://bucket/path). Currently '*' is not allowed.
	ExternalResources []*string `json:"externalResources,omitempty" tf:"external_resources,omitempty"`

	// A list of ApiOperations that this egress rule applies to. A request matches
	// if it contains an operation/service in this list.
	// Structure is documented below.
	Operations []ServicePerimeterEgressPolicyEgressToOperationsObservation `json:"operations,omitempty" tf:"operations,omitempty"`

	// A list of resources, currently only projects in the form
	// projects/<projectnumber>, that match this to stanza. A request matches
	// if it contains a resource in this list. If * is specified for resources,
	// then this EgressTo rule will authorize access to all resources outside
	// the perimeter.
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToOperationsInitParameters struct {

	// API methods or permissions to allow. Method or permission must belong
	// to the service specified by serviceName field. A single MethodSelector
	// entry with * specified for the method field will allow all methods
	// AND permissions for the service specified in serviceName.
	// Structure is documented below.
	MethodSelectors []ServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsInitParameters `json:"methodSelectors,omitempty" tf:"method_selectors,omitempty"`

	// The name of the API whose methods or permissions the IngressPolicy or
	// EgressPolicy want to allow. A single ApiOperation with serviceName
	// field set to * will allow all methods AND permissions for all services.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsInitParameters struct {

	// Value for method should be a valid method name for the corresponding
	// serviceName in ApiOperation. If * used as value for method,
	// then ALL methods and permissions are allowed.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Value for permission should be a valid Cloud IAM permission for the
	// corresponding serviceName in ApiOperation.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsObservation struct {

	// Value for method should be a valid method name for the corresponding
	// serviceName in ApiOperation. If * used as value for method,
	// then ALL methods and permissions are allowed.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Value for permission should be a valid Cloud IAM permission for the
	// corresponding serviceName in ApiOperation.
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsParameters struct {

	// Value for method should be a valid method name for the corresponding
	// serviceName in ApiOperation. If * used as value for method,
	// then ALL methods and permissions are allowed.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Value for permission should be a valid Cloud IAM permission for the
	// corresponding serviceName in ApiOperation.
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToOperationsObservation struct {

	// API methods or permissions to allow. Method or permission must belong
	// to the service specified by serviceName field. A single MethodSelector
	// entry with * specified for the method field will allow all methods
	// AND permissions for the service specified in serviceName.
	// Structure is documented below.
	MethodSelectors []ServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsObservation `json:"methodSelectors,omitempty" tf:"method_selectors,omitempty"`

	// The name of the API whose methods or permissions the IngressPolicy or
	// EgressPolicy want to allow. A single ApiOperation with serviceName
	// field set to * will allow all methods AND permissions for all services.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToOperationsParameters struct {

	// API methods or permissions to allow. Method or permission must belong
	// to the service specified by serviceName field. A single MethodSelector
	// entry with * specified for the method field will allow all methods
	// AND permissions for the service specified in serviceName.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MethodSelectors []ServicePerimeterEgressPolicyEgressToOperationsMethodSelectorsParameters `json:"methodSelectors,omitempty" tf:"method_selectors,omitempty"`

	// The name of the API whose methods or permissions the IngressPolicy or
	// EgressPolicy want to allow. A single ApiOperation with serviceName
	// field set to * will allow all methods AND permissions for all services.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ServicePerimeterEgressPolicyEgressToParameters struct {

	// A list of external resources that are allowed to be accessed. A request
	// matches if it contains an external resource in this list (Example:
	// s3://bucket/path). Currently '*' is not allowed.
	// +kubebuilder:validation:Optional
	ExternalResources []*string `json:"externalResources,omitempty" tf:"external_resources,omitempty"`

	// A list of ApiOperations that this egress rule applies to. A request matches
	// if it contains an operation/service in this list.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Operations []ServicePerimeterEgressPolicyEgressToOperationsParameters `json:"operations,omitempty" tf:"operations,omitempty"`

	// A list of resources, currently only projects in the form
	// projects/<projectnumber>, that match this to stanza. A request matches
	// if it contains a resource in this list. If * is specified for resources,
	// then this EgressTo rule will authorize access to all resources outside
	// the perimeter.
	// +kubebuilder:validation:Optional
	Resources []*string `json:"resources,omitempty" tf:"resources,omitempty"`
}

type ServicePerimeterEgressPolicyInitParameters struct {

	// Defines conditions on the source of a request causing this EgressPolicy to apply.
	// Structure is documented below.
	EgressFrom *ServicePerimeterEgressPolicyEgressFromInitParameters `json:"egressFrom,omitempty" tf:"egress_from,omitempty"`

	// Defines the conditions on the ApiOperation and destination resources that
	// cause this EgressPolicy to apply.
	// Structure is documented below.
	EgressTo *ServicePerimeterEgressPolicyEgressToInitParameters `json:"egressTo,omitempty" tf:"egress_to,omitempty"`

	// The name of the Service Perimeter to add this resource to.
	Perimeter *string `json:"perimeter,omitempty" tf:"perimeter,omitempty"`
}

type ServicePerimeterEgressPolicyObservation struct {

	// Defines conditions on the source of a request causing this EgressPolicy to apply.
	// Structure is documented below.
	EgressFrom *ServicePerimeterEgressPolicyEgressFromObservation `json:"egressFrom,omitempty" tf:"egress_from,omitempty"`

	// Defines the conditions on the ApiOperation and destination resources that
	// cause this EgressPolicy to apply.
	// Structure is documented below.
	EgressTo *ServicePerimeterEgressPolicyEgressToObservation `json:"egressTo,omitempty" tf:"egress_to,omitempty"`

	// an identifier for the resource with format {{perimeter}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Service Perimeter to add this resource to.
	Perimeter *string `json:"perimeter,omitempty" tf:"perimeter,omitempty"`
}

type ServicePerimeterEgressPolicyParameters struct {

	// Defines conditions on the source of a request causing this EgressPolicy to apply.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EgressFrom *ServicePerimeterEgressPolicyEgressFromParameters `json:"egressFrom,omitempty" tf:"egress_from,omitempty"`

	// Defines the conditions on the ApiOperation and destination resources that
	// cause this EgressPolicy to apply.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EgressTo *ServicePerimeterEgressPolicyEgressToParameters `json:"egressTo,omitempty" tf:"egress_to,omitempty"`

	// The name of the Service Perimeter to add this resource to.
	// +kubebuilder:validation:Optional
	Perimeter *string `json:"perimeter,omitempty" tf:"perimeter,omitempty"`
}

// ServicePerimeterEgressPolicySpec defines the desired state of ServicePerimeterEgressPolicy
type ServicePerimeterEgressPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicePerimeterEgressPolicyParameters `json:"forProvider"`
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
	InitProvider ServicePerimeterEgressPolicyInitParameters `json:"initProvider,omitempty"`
}

// ServicePerimeterEgressPolicyStatus defines the observed state of ServicePerimeterEgressPolicy.
type ServicePerimeterEgressPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicePerimeterEgressPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServicePerimeterEgressPolicy is the Schema for the ServicePerimeterEgressPolicys API. Manage a single EgressPolicy in the 'status' (enforced) configuration for a service perimeter.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServicePerimeterEgressPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.egressFrom) || (has(self.initProvider) && has(self.initProvider.egressFrom))",message="spec.forProvider.egressFrom is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.egressTo) || (has(self.initProvider) && has(self.initProvider.egressTo))",message="spec.forProvider.egressTo is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.perimeter) || (has(self.initProvider) && has(self.initProvider.perimeter))",message="spec.forProvider.perimeter is a required parameter"
	Spec   ServicePerimeterEgressPolicySpec   `json:"spec"`
	Status ServicePerimeterEgressPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicePerimeterEgressPolicyList contains a list of ServicePerimeterEgressPolicys
type ServicePerimeterEgressPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServicePerimeterEgressPolicy `json:"items"`
}

// Repository type metadata.
var (
	ServicePerimeterEgressPolicy_Kind             = "ServicePerimeterEgressPolicy"
	ServicePerimeterEgressPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServicePerimeterEgressPolicy_Kind}.String()
	ServicePerimeterEgressPolicy_KindAPIVersion   = ServicePerimeterEgressPolicy_Kind + "." + CRDGroupVersion.String()
	ServicePerimeterEgressPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ServicePerimeterEgressPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ServicePerimeterEgressPolicy{}, &ServicePerimeterEgressPolicyList{})
}
