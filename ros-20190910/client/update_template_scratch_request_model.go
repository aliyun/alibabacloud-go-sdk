// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateScratchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTemplateScratchRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateTemplateScratchRequest
	GetDescription() *string
	SetExecutionMode(v string) *UpdateTemplateScratchRequest
	GetExecutionMode() *string
	SetLogicalIdStrategy(v string) *UpdateTemplateScratchRequest
	GetLogicalIdStrategy() *string
	SetPreferenceParameters(v []*UpdateTemplateScratchRequestPreferenceParameters) *UpdateTemplateScratchRequest
	GetPreferenceParameters() []*UpdateTemplateScratchRequestPreferenceParameters
	SetRegionId(v string) *UpdateTemplateScratchRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateTemplateScratchRequest
	GetResourceGroupId() *string
	SetSourceResourceGroup(v *UpdateTemplateScratchRequestSourceResourceGroup) *UpdateTemplateScratchRequest
	GetSourceResourceGroup() *UpdateTemplateScratchRequestSourceResourceGroup
	SetSourceResources(v []*UpdateTemplateScratchRequestSourceResources) *UpdateTemplateScratchRequest
	GetSourceResources() []*UpdateTemplateScratchRequestSourceResources
	SetSourceTag(v *UpdateTemplateScratchRequestSourceTag) *UpdateTemplateScratchRequest
	GetSourceTag() *UpdateTemplateScratchRequestSourceTag
	SetTemplateScratchId(v string) *UpdateTemplateScratchRequest
	GetTemplateScratchId() *string
}

type UpdateTemplateScratchRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the scenario.
	//
	// example:
	//
	// The description of the scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- Async (default)
	//
	// 	- Sync
	//
	// > If you have a wide scope of resources, Sync takes longer. If you set ExecutionMode to Sync, we recommend that you specify ClientToken to prevent the execution timeout.
	//
	// example:
	//
	// Sync
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// 	- LongTypePrefixAndIndexSuffix: long-type prefix + index-type suffix
	//
	// 	- LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	//
	// 	- ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	//
	// >  If you set TemplateScratchType to ArchitectureDetection, the default value of LogicalIdStrategy is LongTypePrefixAndHashSuffix. In other cases, the default value of LogicalIdStrategy is LongTypePrefixAndIndexSuffix.
	//
	// example:
	//
	// LongTypePrefixAndIndexSuffix
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The preference parameters of the resource scenario.
	PreferenceParameters []*UpdateTemplateScratchRequestPreferenceParameters `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty" type:"Repeated"`
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	//
	// >  You must specify only one of the following parameters: SourceResources, SourceTag, and SourceResourceGroup.
	SourceResourceGroup *UpdateTemplateScratchRequestSourceResourceGroup `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty" type:"Struct"`
	// The source resources.
	//
	// >  You must specify only one of the following parameters: SourceResources, SourceTag, and SourceResourceGroup.
	SourceResources []*UpdateTemplateScratchRequestSourceResources `json:"SourceResources,omitempty" xml:"SourceResources,omitempty" type:"Repeated"`
	// The source tag.
	//
	// >  You must specify only one of the following parameters: SourceResources, SourceTag, and SourceResourceGroup.
	SourceTag *UpdateTemplateScratchRequestSourceTag `json:"SourceTag,omitempty" xml:"SourceTag,omitempty" type:"Struct"`
	// The ID of the resource scenario.
	//
	// The valid values of the ParameterKey and ParameterValue request parameters vary based on the IDs of different types of resource scenarios. For more information, see the "Additional information about request parameters" section of this topic.
	//
	// >  You can call the [ListTemplateScratches](https://help.aliyun.com/document_detail/610832.html) operation to query the ID of a resource scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// ts-7f7a704cf71c49a6****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s UpdateTemplateScratchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateScratchRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTemplateScratchRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTemplateScratchRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *UpdateTemplateScratchRequest) GetLogicalIdStrategy() *string {
	return s.LogicalIdStrategy
}

func (s *UpdateTemplateScratchRequest) GetPreferenceParameters() []*UpdateTemplateScratchRequestPreferenceParameters {
	return s.PreferenceParameters
}

func (s *UpdateTemplateScratchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTemplateScratchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTemplateScratchRequest) GetSourceResourceGroup() *UpdateTemplateScratchRequestSourceResourceGroup {
	return s.SourceResourceGroup
}

func (s *UpdateTemplateScratchRequest) GetSourceResources() []*UpdateTemplateScratchRequestSourceResources {
	return s.SourceResources
}

func (s *UpdateTemplateScratchRequest) GetSourceTag() *UpdateTemplateScratchRequestSourceTag {
	return s.SourceTag
}

func (s *UpdateTemplateScratchRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *UpdateTemplateScratchRequest) SetClientToken(v string) *UpdateTemplateScratchRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetDescription(v string) *UpdateTemplateScratchRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetExecutionMode(v string) *UpdateTemplateScratchRequest {
	s.ExecutionMode = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetLogicalIdStrategy(v string) *UpdateTemplateScratchRequest {
	s.LogicalIdStrategy = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetPreferenceParameters(v []*UpdateTemplateScratchRequestPreferenceParameters) *UpdateTemplateScratchRequest {
	s.PreferenceParameters = v
	return s
}

func (s *UpdateTemplateScratchRequest) SetRegionId(v string) *UpdateTemplateScratchRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetResourceGroupId(v string) *UpdateTemplateScratchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetSourceResourceGroup(v *UpdateTemplateScratchRequestSourceResourceGroup) *UpdateTemplateScratchRequest {
	s.SourceResourceGroup = v
	return s
}

func (s *UpdateTemplateScratchRequest) SetSourceResources(v []*UpdateTemplateScratchRequestSourceResources) *UpdateTemplateScratchRequest {
	s.SourceResources = v
	return s
}

func (s *UpdateTemplateScratchRequest) SetSourceTag(v *UpdateTemplateScratchRequestSourceTag) *UpdateTemplateScratchRequest {
	s.SourceTag = v
	return s
}

func (s *UpdateTemplateScratchRequest) SetTemplateScratchId(v string) *UpdateTemplateScratchRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *UpdateTemplateScratchRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateScratchRequestPreferenceParameters struct {
	// The parameter name.
	//
	// For more information about the valid values of ParameterKey, see the "**Additional information about request parameters**" section of this topic.
	//
	// >- PreferenceParameters is optional. If you specify PreferenceParameters, you must specify both ParameterKey and ParameterValue.
	//
	// > - If you set TemplateScratchType to ResourceImport, you must set ParameterKey to DeletionPolicy.
	//
	// This parameter is required.
	//
	// example:
	//
	// DeletionPolicy
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value. The value of ParameterValue varies based on the value of ParameterKey.
	//
	// For more information about the valid values of ParameterKey, see the "**Additional information about request parameters**" section of this topic.
	//
	// >  PreferenceParameters is optional. If you specify PreferenceParameters, you must specify both ParameterKey and ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// Retain
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateTemplateScratchRequestPreferenceParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateScratchRequestPreferenceParameters) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequestPreferenceParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateTemplateScratchRequestPreferenceParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateTemplateScratchRequestPreferenceParameters) SetParameterKey(v string) *UpdateTemplateScratchRequestPreferenceParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateTemplateScratchRequestPreferenceParameters) SetParameterValue(v string) *UpdateTemplateScratchRequestPreferenceParameters {
	s.ParameterValue = &v
	return s
}

func (s *UpdateTemplateScratchRequestPreferenceParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateScratchRequestSourceResourceGroup struct {
	// The ID of the source resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource types for filtering resources.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s UpdateTemplateScratchRequestSourceResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateScratchRequestSourceResourceGroup) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequestSourceResourceGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTemplateScratchRequestSourceResourceGroup) GetResourceTypeFilter() []*string {
	return s.ResourceTypeFilter
}

func (s *UpdateTemplateScratchRequestSourceResourceGroup) SetResourceGroupId(v string) *UpdateTemplateScratchRequestSourceResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateScratchRequestSourceResourceGroup) SetResourceTypeFilter(v []*string) *UpdateTemplateScratchRequestSourceResourceGroup {
	s.ResourceTypeFilter = v
	return s
}

func (s *UpdateTemplateScratchRequestSourceResourceGroup) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateScratchRequestSourceResources struct {
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1m6fww66xbntjyc****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::ECS::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s UpdateTemplateScratchRequestSourceResources) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateScratchRequestSourceResources) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequestSourceResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateTemplateScratchRequestSourceResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateTemplateScratchRequestSourceResources) SetResourceId(v string) *UpdateTemplateScratchRequestSourceResources {
	s.ResourceId = &v
	return s
}

func (s *UpdateTemplateScratchRequestSourceResources) SetResourceType(v string) *UpdateTemplateScratchRequestSourceResources {
	s.ResourceType = &v
	return s
}

func (s *UpdateTemplateScratchRequestSourceResources) Validate() error {
	return dara.Validate(s)
}

type UpdateTemplateScratchRequestSourceTag struct {
	// The source tags. A tag contains a tag key and a tag value.
	//
	// If you want to specify only the tag key, you must leave the tag value empty. Example: {"TagKey": ""}.
	//
	// If you set TemplateScratchType to ArchitectureDetection, you can add up to 5 source tags. In other cases, you can add up to 10 source tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"a": "b"}
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The resource types for filtering resources.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s UpdateTemplateScratchRequestSourceTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateScratchRequestSourceTag) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequestSourceTag) GetResourceTags() map[string]interface{} {
	return s.ResourceTags
}

func (s *UpdateTemplateScratchRequestSourceTag) GetResourceTypeFilter() []*string {
	return s.ResourceTypeFilter
}

func (s *UpdateTemplateScratchRequestSourceTag) SetResourceTags(v map[string]interface{}) *UpdateTemplateScratchRequestSourceTag {
	s.ResourceTags = v
	return s
}

func (s *UpdateTemplateScratchRequestSourceTag) SetResourceTypeFilter(v []*string) *UpdateTemplateScratchRequestSourceTag {
	s.ResourceTypeFilter = v
	return s
}

func (s *UpdateTemplateScratchRequestSourceTag) Validate() error {
	return dara.Validate(s)
}
