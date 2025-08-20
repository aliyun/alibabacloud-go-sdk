// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFeatureDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDriftDetection(v *GetFeatureDetailsResponseBodyDriftDetection) *GetFeatureDetailsResponseBody
	GetDriftDetection() *GetFeatureDetailsResponseBodyDriftDetection
	SetRequestId(v string) *GetFeatureDetailsResponseBody
	GetRequestId() *string
	SetResourceCleaner(v *GetFeatureDetailsResponseBodyResourceCleaner) *GetFeatureDetailsResponseBody
	GetResourceCleaner() *GetFeatureDetailsResponseBodyResourceCleaner
	SetResourceImport(v *GetFeatureDetailsResponseBodyResourceImport) *GetFeatureDetailsResponseBody
	GetResourceImport() *GetFeatureDetailsResponseBodyResourceImport
	SetTemplateParameterConstraints(v *GetFeatureDetailsResponseBodyTemplateParameterConstraints) *GetFeatureDetailsResponseBody
	GetTemplateParameterConstraints() *GetFeatureDetailsResponseBodyTemplateParameterConstraints
	SetTemplateScratch(v *GetFeatureDetailsResponseBodyTemplateScratch) *GetFeatureDetailsResponseBody
	GetTemplateScratch() *GetFeatureDetailsResponseBodyTemplateScratch
	SetTerraform(v *GetFeatureDetailsResponseBodyTerraform) *GetFeatureDetailsResponseBody
	GetTerraform() *GetFeatureDetailsResponseBodyTerraform
}

type GetFeatureDetailsResponseBody struct {
	// Details of the drift detection feature.
	DriftDetection *GetFeatureDetailsResponseBodyDriftDetection `json:"DriftDetection,omitempty" xml:"DriftDetection,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EBF833DA-D0E2-52BE-92E2-59CA56BE834E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the resource cleaner feature.
	ResourceCleaner *GetFeatureDetailsResponseBodyResourceCleaner `json:"ResourceCleaner,omitempty" xml:"ResourceCleaner,omitempty" type:"Struct"`
	// Details of the resource import feature.
	ResourceImport *GetFeatureDetailsResponseBodyResourceImport `json:"ResourceImport,omitempty" xml:"ResourceImport,omitempty" type:"Struct"`
	// Details of the template parameter constraint feature.
	TemplateParameterConstraints *GetFeatureDetailsResponseBodyTemplateParameterConstraints `json:"TemplateParameterConstraints,omitempty" xml:"TemplateParameterConstraints,omitempty" type:"Struct"`
	// Details of the scenario feature.
	TemplateScratch *GetFeatureDetailsResponseBodyTemplateScratch `json:"TemplateScratch,omitempty" xml:"TemplateScratch,omitempty" type:"Struct"`
	// Details of the Terraform hosting feature.
	Terraform *GetFeatureDetailsResponseBodyTerraform `json:"Terraform,omitempty" xml:"Terraform,omitempty" type:"Struct"`
}

func (s GetFeatureDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBody) GetDriftDetection() *GetFeatureDetailsResponseBodyDriftDetection {
	return s.DriftDetection
}

func (s *GetFeatureDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFeatureDetailsResponseBody) GetResourceCleaner() *GetFeatureDetailsResponseBodyResourceCleaner {
	return s.ResourceCleaner
}

func (s *GetFeatureDetailsResponseBody) GetResourceImport() *GetFeatureDetailsResponseBodyResourceImport {
	return s.ResourceImport
}

func (s *GetFeatureDetailsResponseBody) GetTemplateParameterConstraints() *GetFeatureDetailsResponseBodyTemplateParameterConstraints {
	return s.TemplateParameterConstraints
}

func (s *GetFeatureDetailsResponseBody) GetTemplateScratch() *GetFeatureDetailsResponseBodyTemplateScratch {
	return s.TemplateScratch
}

func (s *GetFeatureDetailsResponseBody) GetTerraform() *GetFeatureDetailsResponseBodyTerraform {
	return s.Terraform
}

func (s *GetFeatureDetailsResponseBody) SetDriftDetection(v *GetFeatureDetailsResponseBodyDriftDetection) *GetFeatureDetailsResponseBody {
	s.DriftDetection = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetRequestId(v string) *GetFeatureDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetResourceCleaner(v *GetFeatureDetailsResponseBodyResourceCleaner) *GetFeatureDetailsResponseBody {
	s.ResourceCleaner = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetResourceImport(v *GetFeatureDetailsResponseBodyResourceImport) *GetFeatureDetailsResponseBody {
	s.ResourceImport = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetTemplateParameterConstraints(v *GetFeatureDetailsResponseBodyTemplateParameterConstraints) *GetFeatureDetailsResponseBody {
	s.TemplateParameterConstraints = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetTemplateScratch(v *GetFeatureDetailsResponseBodyTemplateScratch) *GetFeatureDetailsResponseBody {
	s.TemplateScratch = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetTerraform(v *GetFeatureDetailsResponseBodyTerraform) *GetFeatureDetailsResponseBody {
	s.Terraform = v
	return s
}

func (s *GetFeatureDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyDriftDetection struct {
	// The resource types that are supported by the drift detection feature.
	SupportedResourceTypes []*string `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyDriftDetection) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyDriftDetection) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyDriftDetection) GetSupportedResourceTypes() []*string {
	return s.SupportedResourceTypes
}

func (s *GetFeatureDetailsResponseBodyDriftDetection) SetSupportedResourceTypes(v []*string) *GetFeatureDetailsResponseBodyDriftDetection {
	s.SupportedResourceTypes = v
	return s
}

func (s *GetFeatureDetailsResponseBodyDriftDetection) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyResourceCleaner struct {
	// The resource types that can be cleaned up.
	SupportedResourceTypes []*GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyResourceCleaner) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyResourceCleaner) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyResourceCleaner) GetSupportedResourceTypes() []*GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes {
	return s.SupportedResourceTypes
}

func (s *GetFeatureDetailsResponseBodyResourceCleaner) SetSupportedResourceTypes(v []*GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) *GetFeatureDetailsResponseBodyResourceCleaner {
	s.SupportedResourceTypes = v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceCleaner) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes struct {
	// The resource type that supports the resource cleaner feature.
	//
	// example:
	//
	// ECS:Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The names of the side effects that may be caused by the cleanup operation performed on the resources of the specified type.
	SideEffects []*string `json:"SideEffects,omitempty" xml:"SideEffects,omitempty" type:"Repeated"`
	// The names of the filters that are supported by the resource type.
	SupportedFilters []*string `json:"SupportedFilters,omitempty" xml:"SupportedFilters,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) GetSideEffects() []*string {
	return s.SideEffects
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) GetSupportedFilters() []*string {
	return s.SupportedFilters
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) SetResourceType(v string) *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) SetSideEffects(v []*string) *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes {
	s.SideEffects = v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) SetSupportedFilters(v []*string) *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes {
	s.SupportedFilters = v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyResourceImport struct {
	// The resource types that are supported by the resource import feature.
	SupportedResourceTypes []*GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyResourceImport) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyResourceImport) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyResourceImport) GetSupportedResourceTypes() []*GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes {
	return s.SupportedResourceTypes
}

func (s *GetFeatureDetailsResponseBodyResourceImport) SetSupportedResourceTypes(v []*GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) *GetFeatureDetailsResponseBodyResourceImport {
	s.SupportedResourceTypes = v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceImport) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes struct {
	// The resource identifiers.
	ResourceIdentifiers []*string `json:"ResourceIdentifiers,omitempty" xml:"ResourceIdentifiers,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::Disk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) GetResourceIdentifiers() []*string {
	return s.ResourceIdentifiers
}

func (s *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) SetResourceIdentifiers(v []*string) *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes {
	s.ResourceIdentifiers = v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) SetResourceType(v string) *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTemplateParameterConstraints struct {
	// The resource types that support the template parameter constraint feature.
	SupportedResourceTypes []*GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTemplateParameterConstraints) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTemplateParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraints) GetSupportedResourceTypes() []*GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes {
	return s.SupportedResourceTypes
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraints) SetSupportedResourceTypes(v []*GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) *GetFeatureDetailsResponseBodyTemplateParameterConstraints {
	s.SupportedResourceTypes = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraints) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes struct {
	// The names of properties that are supported by the resource type.
	Properties []*string `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::Disk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) GetProperties() []*string {
	return s.Properties
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) SetProperties(v []*string) *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes {
	s.Properties = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) SetResourceType(v string) *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTemplateScratch struct {
	// The resource types that are supported by the scenario feature.
	SupportedResourceTypes []*GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTemplateScratch) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTemplateScratch) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTemplateScratch) GetSupportedResourceTypes() []*GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	return s.SupportedResourceTypes
}

func (s *GetFeatureDetailsResponseBodyTemplateScratch) SetSupportedResourceTypes(v []*GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) *GetFeatureDetailsResponseBodyTemplateScratch {
	s.SupportedResourceTypes = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratch) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes struct {
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::Disk
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Indicates whether the resource scope can be specified by source resource group. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SourceResourceGroupSupported *bool `json:"SourceResourceGroupSupported,omitempty" xml:"SourceResourceGroupSupported,omitempty"`
	// Indicates whether the resource scope can be specified by source resource. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SourceResourcesSupported *bool `json:"SourceResourcesSupported,omitempty" xml:"SourceResourcesSupported,omitempty"`
	// Indicates whether the resource scope can be specified by source tag, resource group, or resource. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SourceSupported *bool `json:"SourceSupported,omitempty" xml:"SourceSupported,omitempty"`
	// Indicates whether the resource scope can be specified by source tag. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SourceTagSupported *bool `json:"SourceTagSupported,omitempty" xml:"SourceTagSupported,omitempty"`
	// The scenario types that are supported.
	SupportedTemplateScratchTypes []*string `json:"SupportedTemplateScratchTypes,omitempty" xml:"SupportedTemplateScratchTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) GetSourceResourceGroupSupported() *bool {
	return s.SourceResourceGroupSupported
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) GetSourceResourcesSupported() *bool {
	return s.SourceResourcesSupported
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) GetSourceSupported() *bool {
	return s.SourceSupported
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) GetSourceTagSupported() *bool {
	return s.SourceTagSupported
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) GetSupportedTemplateScratchTypes() []*string {
	return s.SupportedTemplateScratchTypes
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetResourceType(v string) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSourceResourceGroupSupported(v bool) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SourceResourceGroupSupported = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSourceResourcesSupported(v bool) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SourceResourcesSupported = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSourceSupported(v bool) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SourceSupported = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSourceTagSupported(v bool) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SourceTagSupported = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSupportedTemplateScratchTypes(v []*string) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SupportedTemplateScratchTypes = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTerraform struct {
	// The resource types that support the scenario feature.
	SupportedResourceTypes *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Struct"`
	// The Terraform versions.
	SupportedVersions []*GetFeatureDetailsResponseBodyTerraformSupportedVersions `json:"SupportedVersions,omitempty" xml:"SupportedVersions,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraform) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraform) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraform) GetSupportedResourceTypes() *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	return s.SupportedResourceTypes
}

func (s *GetFeatureDetailsResponseBodyTerraform) GetSupportedVersions() []*GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	return s.SupportedVersions
}

func (s *GetFeatureDetailsResponseBodyTerraform) SetSupportedResourceTypes(v *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) *GetFeatureDetailsResponseBodyTerraform {
	s.SupportedResourceTypes = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraform) SetSupportedVersions(v []*GetFeatureDetailsResponseBodyTerraformSupportedVersions) *GetFeatureDetailsResponseBodyTerraform {
	s.SupportedVersions = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraform) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes struct {
	// The resource types that support the custom tag feature.
	CustomTag []*string `json:"CustomTag,omitempty" xml:"CustomTag,omitempty" type:"Repeated"`
	// The resource types that support the price inquiry feature.
	EstimateCost []*string `json:"EstimateCost,omitempty" xml:"EstimateCost,omitempty" type:"Repeated"`
	// The resource types that support the resource group feature.
	ResourceGroup []*string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Repeated"`
	// The resource type that support the risk check feature.
	StackOperationRisk *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk `json:"StackOperationRisk,omitempty" xml:"StackOperationRisk,omitempty" type:"Struct"`
	// The resource types that support the system tag `acs:ros:stackId`.
	SystemTag []*string `json:"SystemTag,omitempty" xml:"SystemTag,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) GetCustomTag() []*string {
	return s.CustomTag
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) GetEstimateCost() []*string {
	return s.EstimateCost
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) GetResourceGroup() []*string {
	return s.ResourceGroup
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) GetStackOperationRisk() *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk {
	return s.StackOperationRisk
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) GetSystemTag() []*string {
	return s.SystemTag
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetCustomTag(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.CustomTag = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetEstimateCost(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.EstimateCost = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetResourceGroup(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.ResourceGroup = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetStackOperationRisk(v *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.StackOperationRisk = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetSystemTag(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.SystemTag = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk struct {
	// The resource types that support the risk check performed to detect risks caused by a stack deletion operation.
	DeleteStack []*string `json:"DeleteStack,omitempty" xml:"DeleteStack,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) GetDeleteStack() []*string {
	return s.DeleteStack
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) SetDeleteStack(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk {
	s.DeleteStack = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTerraformSupportedVersions struct {
	// The names and versions of the providers that correspond to the Terraform versions.
	ProviderVersions []*GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions `json:"ProviderVersions,omitempty" xml:"ProviderVersions,omitempty" type:"Repeated"`
	// The Terraform version.
	//
	// example:
	//
	// 1.0.11
	TerraformVersion *string `json:"TerraformVersion,omitempty" xml:"TerraformVersion,omitempty"`
	// The Terraform version that is supported by ROS. The parameter value is the same as the value of the Transform parameter in a Terraform template.
	//
	// example:
	//
	// Aliyun::Terraform-v1.0
	Transform *string `json:"Transform,omitempty" xml:"Transform,omitempty"`
	// The Terraform versions that can be updated in ROS.
	UpdateAllowedTransforms []*string `json:"UpdateAllowedTransforms,omitempty" xml:"UpdateAllowedTransforms,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedVersions) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedVersions) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) GetProviderVersions() []*GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions {
	return s.ProviderVersions
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) GetTransform() *string {
	return s.Transform
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) GetUpdateAllowedTransforms() []*string {
	return s.UpdateAllowedTransforms
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) SetProviderVersions(v []*GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) *GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	s.ProviderVersions = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) SetTerraformVersion(v string) *GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	s.TerraformVersion = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) SetTransform(v string) *GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	s.Transform = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) SetUpdateAllowedTransforms(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	s.UpdateAllowedTransforms = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) Validate() error {
	return dara.Validate(s)
}

type GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions struct {
	// The name of the provider.
	//
	// example:
	//
	// alicloud
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// The provider versions.
	SupportedVersions []*string `json:"SupportedVersions,omitempty" xml:"SupportedVersions,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) String() string {
	return dara.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) GetProviderName() *string {
	return s.ProviderName
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) GetSupportedVersions() []*string {
	return s.SupportedVersions
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) SetProviderName(v string) *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions {
	s.ProviderName = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) SetSupportedVersions(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions {
	s.SupportedVersions = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) Validate() error {
	return dara.Validate(s)
}
