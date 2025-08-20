// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateScratchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTemplateScratchResponseBody
	GetRequestId() *string
	SetTemplateScratch(v *GetTemplateScratchResponseBodyTemplateScratch) *GetTemplateScratchResponseBody
	GetTemplateScratch() *GetTemplateScratchResponseBodyTemplateScratch
}

type GetTemplateScratchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A8E0EF98-6FBD-5656-8298-FC8194F0F7B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource scenario.
	TemplateScratch *GetTemplateScratchResponseBodyTemplateScratch `json:"TemplateScratch,omitempty" xml:"TemplateScratch,omitempty" type:"Struct"`
}

func (s GetTemplateScratchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateScratchResponseBody) GetTemplateScratch() *GetTemplateScratchResponseBodyTemplateScratch {
	return s.TemplateScratch
}

func (s *GetTemplateScratchResponseBody) SetRequestId(v string) *GetTemplateScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateScratchResponseBody) SetTemplateScratch(v *GetTemplateScratchResponseBodyTemplateScratch) *GetTemplateScratchResponseBody {
	s.TemplateScratch = v
	return s
}

func (s *GetTemplateScratchResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTemplateScratchResponseBodyTemplateScratch struct {
	// The time at which the resource scenario was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-22T01:49:22
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the resource scenario.
	//
	// example:
	//
	// The description of the resource scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The status code of the resource scenario that fails to be created.
	//
	// > This parameter is returned only if you set Status to GENERATE_FAILED.
	//
	// example:
	//
	// InvalidZoneId
	FailedCode *string `json:"FailedCode,omitempty" xml:"FailedCode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// 	- LongTypePrefixAndIndexSuffix (default): long-type prefix + index-type suffix
	//
	// 	- LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	//
	// 	- ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	//
	// example:
	//
	// LongTypePrefixAndIndexSuffix
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The preference parameters of the resource scenario.
	PreferenceParameters []*GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty" type:"Repeated"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzmhzoaad5oq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	SourceResourceGroup *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty" type:"Struct"`
	// The source resources.
	SourceResources []*GetTemplateScratchResponseBodyTemplateScratchSourceResources `json:"SourceResources,omitempty" xml:"SourceResources,omitempty" type:"Repeated"`
	// The source tag.
	SourceTag *GetTemplateScratchResponseBodyTemplateScratchSourceTag `json:"SourceTag,omitempty" xml:"SourceTag,omitempty" type:"Struct"`
	// The preset information of the stack.
	StackProvision *GetTemplateScratchResponseBodyTemplateScratchStackProvision `json:"StackProvision,omitempty" xml:"StackProvision,omitempty" type:"Struct"`
	// The stacks that are associated with the resource scenario.
	Stacks []*GetTemplateScratchResponseBodyTemplateScratchStacks `json:"Stacks,omitempty" xml:"Stacks,omitempty" type:"Repeated"`
	// The state of the resource scenario. Valid values:
	//
	// 	- GENERATE_IN_PROGRESS: The resource scenario is being created.
	//
	// 	- GENERATE_COMPLETE: The resource scenario is created.
	//
	// 	- GENERATE_FAILED: The resource scenario fails to be created.
	//
	// example:
	//
	// GENERATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the resource scenario fails to be created.
	//
	// > This parameter is returned only if you set Status to GENERATE_FAILED.
	//
	// example:
	//
	// Resource ALIYUN::ECS::VPC vpc-m5eauuq80anx59v28***	- could not be found for template scratch.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The resource scenario data.
	TemplateScratchData map[string]interface{} `json:"TemplateScratchData,omitempty" xml:"TemplateScratchData,omitempty"`
	// The ID of the resource scenario.
	//
	// example:
	//
	// ts-7f7a704cf71c49a6****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The type of the resource scenario. Valid values:
	//
	// 	- ResourceImport: resource management
	//
	// 	- ArchitectureReplication: resource replication
	//
	// example:
	//
	// ArchitectureReplication
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
	// The time at which the resource scenario was updated.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-22T01:49:23
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratch) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratch) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetDescription() *string {
	return s.Description
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetFailedCode() *string {
	return s.FailedCode
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetLogicalIdStrategy() *string {
	return s.LogicalIdStrategy
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetPreferenceParameters() []*GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters {
	return s.PreferenceParameters
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetSourceResourceGroup() *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup {
	return s.SourceResourceGroup
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetSourceResources() []*GetTemplateScratchResponseBodyTemplateScratchSourceResources {
	return s.SourceResources
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetSourceTag() *GetTemplateScratchResponseBodyTemplateScratchSourceTag {
	return s.SourceTag
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetStackProvision() *GetTemplateScratchResponseBodyTemplateScratchStackProvision {
	return s.StackProvision
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetStacks() []*GetTemplateScratchResponseBodyTemplateScratchStacks {
	return s.Stacks
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetStatus() *string {
	return s.Status
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetTemplateScratchData() map[string]interface{} {
	return s.TemplateScratchData
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetTemplateScratchType() *string {
	return s.TemplateScratchType
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetCreateTime(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetDescription(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.Description = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetFailedCode(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.FailedCode = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetLogicalIdStrategy(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.LogicalIdStrategy = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetPreferenceParameters(v []*GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) *GetTemplateScratchResponseBodyTemplateScratch {
	s.PreferenceParameters = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetResourceGroupId(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetSourceResourceGroup(v *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) *GetTemplateScratchResponseBodyTemplateScratch {
	s.SourceResourceGroup = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetSourceResources(v []*GetTemplateScratchResponseBodyTemplateScratchSourceResources) *GetTemplateScratchResponseBodyTemplateScratch {
	s.SourceResources = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetSourceTag(v *GetTemplateScratchResponseBodyTemplateScratchSourceTag) *GetTemplateScratchResponseBodyTemplateScratch {
	s.SourceTag = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetStackProvision(v *GetTemplateScratchResponseBodyTemplateScratchStackProvision) *GetTemplateScratchResponseBodyTemplateScratch {
	s.StackProvision = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetStacks(v []*GetTemplateScratchResponseBodyTemplateScratchStacks) *GetTemplateScratchResponseBodyTemplateScratch {
	s.Stacks = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetStatus(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.Status = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetStatusReason(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.StatusReason = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetTemplateScratchData(v map[string]interface{}) *GetTemplateScratchResponseBodyTemplateScratch {
	s.TemplateScratchData = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetTemplateScratchId(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.TemplateScratchId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetTemplateScratchType(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.TemplateScratchType = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetUpdateTime(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.UpdateTime = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) Validate() error {
	return dara.Validate(s)
}

type GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters struct {
	// The parameter name.
	//
	// example:
	//
	// DeletionPolicy
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// Retain
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) SetParameterKey(v string) *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) SetParameterValue(v string) *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) Validate() error {
	return dara.Validate(s)
}

type GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup struct {
	// The ID of the source resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource type filters.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) GetResourceTypeFilter() []*string {
	return s.ResourceTypeFilter
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) SetResourceGroupId(v string) *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) SetResourceTypeFilter(v []*string) *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup {
	s.ResourceTypeFilter = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) Validate() error {
	return dara.Validate(s)
}

type GetTemplateScratchResponseBodyTemplateScratchSourceResources struct {
	// The related resource type filters.
	RelatedResourceTypeFilter []*string `json:"RelatedResourceTypeFilter,omitempty" xml:"RelatedResourceTypeFilter,omitempty" type:"Repeated"`
	// The resource ID.
	//
	// example:
	//
	// vpc-m5e7cv7e9mz69sszb****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceResources) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceResources) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) GetRelatedResourceTypeFilter() []*string {
	return s.RelatedResourceTypeFilter
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) SetRelatedResourceTypeFilter(v []*string) *GetTemplateScratchResponseBodyTemplateScratchSourceResources {
	s.RelatedResourceTypeFilter = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) SetResourceId(v string) *GetTemplateScratchResponseBodyTemplateScratchSourceResources {
	s.ResourceId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) SetResourceType(v string) *GetTemplateScratchResponseBodyTemplateScratchSourceResources {
	s.ResourceType = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) Validate() error {
	return dara.Validate(s)
}

type GetTemplateScratchResponseBodyTemplateScratchSourceTag struct {
	// The source tags.
	//
	// example:
	//
	// {"a": "b"}
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The resource type filters.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceTag) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceTag) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceTag) GetResourceTags() map[string]interface{} {
	return s.ResourceTags
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceTag) GetResourceTypeFilter() []*string {
	return s.ResourceTypeFilter
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceTag) SetResourceTags(v map[string]interface{}) *GetTemplateScratchResponseBodyTemplateScratchSourceTag {
	s.ResourceTags = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceTag) SetResourceTypeFilter(v []*string) *GetTemplateScratchResponseBodyTemplateScratchSourceTag {
	s.ResourceTypeFilter = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceTag) Validate() error {
	return dara.Validate(s)
}

type GetTemplateScratchResponseBodyTemplateScratchStackProvision struct {
	// Indicates whether the resource is replicated by calling the [CreateStack](https://help.aliyun.com/document_detail/132086.html) operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Creatable *bool `json:"Creatable,omitempty" xml:"Creatable,omitempty"`
	// Indicates whether the resource is managed by calling the [CreateChangeSet](https://help.aliyun.com/document_detail/131051.html) operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Importable *bool `json:"Importable,omitempty" xml:"Importable,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchStackProvision) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchStackProvision) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStackProvision) GetCreatable() *bool {
	return s.Creatable
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStackProvision) GetImportable() *bool {
	return s.Importable
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStackProvision) SetCreatable(v bool) *GetTemplateScratchResponseBodyTemplateScratchStackProvision {
	s.Creatable = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStackProvision) SetImportable(v bool) *GetTemplateScratchResponseBodyTemplateScratchStackProvision {
	s.Importable = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStackProvision) Validate() error {
	return dara.Validate(s)
}

type GetTemplateScratchResponseBodyTemplateScratchStacks struct {
	// The region ID of the stack.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID.
	//
	// example:
	//
	// 3708bf6a-3a67-44d4-9eb1-c56704b9****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The purpose of the stack. Valid values:
	//
	// 	- ResourceImport: resource management
	//
	// 	- ArchitectureReplication: resource replication
	//
	// example:
	//
	// ArchitectureReplication
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchStacks) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchStacks) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) GetStackId() *string {
	return s.StackId
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) GetUsageType() *string {
	return s.UsageType
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) SetRegionId(v string) *GetTemplateScratchResponseBodyTemplateScratchStacks {
	s.RegionId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) SetStackId(v string) *GetTemplateScratchResponseBodyTemplateScratchStacks {
	s.StackId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) SetUsageType(v string) *GetTemplateScratchResponseBodyTemplateScratchStacks {
	s.UsageType = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) Validate() error {
	return dara.Validate(s)
}
