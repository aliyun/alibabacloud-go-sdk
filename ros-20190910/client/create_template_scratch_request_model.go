// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateScratchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTemplateScratchRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTemplateScratchRequest
	GetDescription() *string
	SetExecutionMode(v string) *CreateTemplateScratchRequest
	GetExecutionMode() *string
	SetLogicalIdStrategy(v string) *CreateTemplateScratchRequest
	GetLogicalIdStrategy() *string
	SetPreferenceParameters(v []*CreateTemplateScratchRequestPreferenceParameters) *CreateTemplateScratchRequest
	GetPreferenceParameters() []*CreateTemplateScratchRequestPreferenceParameters
	SetRegionId(v string) *CreateTemplateScratchRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTemplateScratchRequest
	GetResourceGroupId() *string
	SetSourceResourceGroup(v *CreateTemplateScratchRequestSourceResourceGroup) *CreateTemplateScratchRequest
	GetSourceResourceGroup() *CreateTemplateScratchRequestSourceResourceGroup
	SetSourceResources(v []*CreateTemplateScratchRequestSourceResources) *CreateTemplateScratchRequest
	GetSourceResources() []*CreateTemplateScratchRequestSourceResources
	SetSourceTag(v *CreateTemplateScratchRequestSourceTag) *CreateTemplateScratchRequest
	GetSourceTag() *CreateTemplateScratchRequestSourceTag
	SetTags(v []*CreateTemplateScratchRequestTags) *CreateTemplateScratchRequest
	GetTags() []*CreateTemplateScratchRequestTags
	SetTemplateScratchType(v string) *CreateTemplateScratchRequest
	GetTemplateScratchType() *string
}

type CreateTemplateScratchRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the resource scenario.
	//
	// example:
	//
	// Replicate a VPC.
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
	PreferenceParameters []*CreateTemplateScratchRequestPreferenceParameters `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty" type:"Repeated"`
	// The region ID of the resource scenario.
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
	SourceResourceGroup *CreateTemplateScratchRequestSourceResourceGroup `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty" type:"Struct"`
	// The source resources.
	//
	// When you set TemplateScratchType to ArchitectureDetection, you can specify SourceResources to detect the architecture data of all resources associated with the specified source resources. For example, if you set SourceResources to the ID of a Classic Load Balancer (CLB) instance, the architecture data of all resources, such as the Elastic Compute Service (ECS) instance, vSwitch, and VPC, associated with the CLB instance is detected.
	//
	// If you set TemplateScratchType to ArchitectureDetection, you can specify up to 20 source resources. In other cases, you can specify up to 200 source resources.
	SourceResources []*CreateTemplateScratchRequestSourceResources `json:"SourceResources,omitempty" xml:"SourceResources,omitempty" type:"Repeated"`
	// The source tag.
	SourceTag *CreateTemplateScratchRequestSourceTag `json:"SourceTag,omitempty" xml:"SourceTag,omitempty" type:"Struct"`
	// The tags of the resource scenario.
	Tags []*CreateTemplateScratchRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the resource scenario. Valid values:
	//
	// 	- ArchitectureReplication: resource replication
	//
	// 	- ArchitectureDetection: resource detection
	//
	// 	- ResourceImport: resource management
	//
	// 	- ResourceMigration: resource migration
	//
	// >  The valid values of the ParameterKey and ParameterValue request parameters vary based on the value of TemplateScratchType. For more information, see the "**Additional information about request parameters**" section of this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// ArchitectureReplication
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
}

func (s CreateTemplateScratchRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTemplateScratchRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTemplateScratchRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *CreateTemplateScratchRequest) GetLogicalIdStrategy() *string {
	return s.LogicalIdStrategy
}

func (s *CreateTemplateScratchRequest) GetPreferenceParameters() []*CreateTemplateScratchRequestPreferenceParameters {
	return s.PreferenceParameters
}

func (s *CreateTemplateScratchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTemplateScratchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTemplateScratchRequest) GetSourceResourceGroup() *CreateTemplateScratchRequestSourceResourceGroup {
	return s.SourceResourceGroup
}

func (s *CreateTemplateScratchRequest) GetSourceResources() []*CreateTemplateScratchRequestSourceResources {
	return s.SourceResources
}

func (s *CreateTemplateScratchRequest) GetSourceTag() *CreateTemplateScratchRequestSourceTag {
	return s.SourceTag
}

func (s *CreateTemplateScratchRequest) GetTags() []*CreateTemplateScratchRequestTags {
	return s.Tags
}

func (s *CreateTemplateScratchRequest) GetTemplateScratchType() *string {
	return s.TemplateScratchType
}

func (s *CreateTemplateScratchRequest) SetClientToken(v string) *CreateTemplateScratchRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetDescription(v string) *CreateTemplateScratchRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetExecutionMode(v string) *CreateTemplateScratchRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetLogicalIdStrategy(v string) *CreateTemplateScratchRequest {
	s.LogicalIdStrategy = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetPreferenceParameters(v []*CreateTemplateScratchRequestPreferenceParameters) *CreateTemplateScratchRequest {
	s.PreferenceParameters = v
	return s
}

func (s *CreateTemplateScratchRequest) SetRegionId(v string) *CreateTemplateScratchRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetResourceGroupId(v string) *CreateTemplateScratchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetSourceResourceGroup(v *CreateTemplateScratchRequestSourceResourceGroup) *CreateTemplateScratchRequest {
	s.SourceResourceGroup = v
	return s
}

func (s *CreateTemplateScratchRequest) SetSourceResources(v []*CreateTemplateScratchRequestSourceResources) *CreateTemplateScratchRequest {
	s.SourceResources = v
	return s
}

func (s *CreateTemplateScratchRequest) SetSourceTag(v *CreateTemplateScratchRequestSourceTag) *CreateTemplateScratchRequest {
	s.SourceTag = v
	return s
}

func (s *CreateTemplateScratchRequest) SetTags(v []*CreateTemplateScratchRequestTags) *CreateTemplateScratchRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateScratchRequest) SetTemplateScratchType(v string) *CreateTemplateScratchRequest {
	s.TemplateScratchType = &v
	return s
}

func (s *CreateTemplateScratchRequest) Validate() error {
	if s.PreferenceParameters != nil {
		for _, item := range s.PreferenceParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SourceResourceGroup != nil {
		if err := s.SourceResourceGroup.Validate(); err != nil {
			return err
		}
	}
	if s.SourceResources != nil {
		for _, item := range s.SourceResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SourceTag != nil {
		if err := s.SourceTag.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTemplateScratchRequestPreferenceParameters struct {
	// The parameter name.
	//
	// For more information about the valid values of ParameterKey, see the "**Additional information about request parameters**" section of this topic.
	//
	// >
	//
	// 	- PreferenceParameters is optional. If you specify PreferenceParameters, you must specify ParameterKey and ParameterValue.
	//
	// 	- You must set ParameterKey to DeletionPolicy when TemplateScratchType is set to ResourceImport.
	//
	// This parameter is required.
	//
	// example:
	//
	// DeletionPolicy
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value. The value is an assignment to the parameter key.
	//
	// For more information about the valid values of ParameterValue, see the "**Additional information about request parameters**" section of this topic.
	//
	// >  PreferenceParameters is optional. If you specify PreferenceParameters, you must specify ParameterKey and ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// Retain
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateTemplateScratchRequestPreferenceParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchRequestPreferenceParameters) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestPreferenceParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateTemplateScratchRequestPreferenceParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateTemplateScratchRequestPreferenceParameters) SetParameterKey(v string) *CreateTemplateScratchRequestPreferenceParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateTemplateScratchRequestPreferenceParameters) SetParameterValue(v string) *CreateTemplateScratchRequestPreferenceParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateTemplateScratchRequestPreferenceParameters) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateScratchRequestSourceResourceGroup struct {
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

func (s CreateTemplateScratchRequestSourceResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchRequestSourceResourceGroup) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestSourceResourceGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTemplateScratchRequestSourceResourceGroup) GetResourceTypeFilter() []*string {
	return s.ResourceTypeFilter
}

func (s *CreateTemplateScratchRequestSourceResourceGroup) SetResourceGroupId(v string) *CreateTemplateScratchRequestSourceResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateScratchRequestSourceResourceGroup) SetResourceTypeFilter(v []*string) *CreateTemplateScratchRequestSourceResourceGroup {
	s.ResourceTypeFilter = v
	return s
}

func (s *CreateTemplateScratchRequestSourceResourceGroup) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateScratchRequestSourceResources struct {
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// >
	//
	// 	- This parameter takes effect only when TemplateScratchType is set to ArchitectureDetection.
	//
	// 	- The region ID of a global resource is `global`. For example, the region ID of the ALIYUN::CDN::Domain global resource is `global`.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The related resource type filters.
	RelatedResourceTypeFilter []*string `json:"RelatedResourceTypeFilter,omitempty" xml:"RelatedResourceTypeFilter,omitempty" type:"Repeated"`
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

func (s CreateTemplateScratchRequestSourceResources) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchRequestSourceResources) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestSourceResources) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTemplateScratchRequestSourceResources) GetRelatedResourceTypeFilter() []*string {
	return s.RelatedResourceTypeFilter
}

func (s *CreateTemplateScratchRequestSourceResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateTemplateScratchRequestSourceResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateTemplateScratchRequestSourceResources) SetRegionId(v string) *CreateTemplateScratchRequestSourceResources {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateScratchRequestSourceResources) SetRelatedResourceTypeFilter(v []*string) *CreateTemplateScratchRequestSourceResources {
	s.RelatedResourceTypeFilter = v
	return s
}

func (s *CreateTemplateScratchRequestSourceResources) SetResourceId(v string) *CreateTemplateScratchRequestSourceResources {
	s.ResourceId = &v
	return s
}

func (s *CreateTemplateScratchRequestSourceResources) SetResourceType(v string) *CreateTemplateScratchRequestSourceResources {
	s.ResourceType = &v
	return s
}

func (s *CreateTemplateScratchRequestSourceResources) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateScratchRequestSourceTag struct {
	// The source tags that consist of key-value pairs. If you want to specify only the tag key, you must leave the tag value empty. Example: `{"TagKey": ""}`.
	//
	// You can add up to 10 source tags.
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

func (s CreateTemplateScratchRequestSourceTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchRequestSourceTag) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestSourceTag) GetResourceTags() map[string]interface{} {
	return s.ResourceTags
}

func (s *CreateTemplateScratchRequestSourceTag) GetResourceTypeFilter() []*string {
	return s.ResourceTypeFilter
}

func (s *CreateTemplateScratchRequestSourceTag) SetResourceTags(v map[string]interface{}) *CreateTemplateScratchRequestSourceTag {
	s.ResourceTags = v
	return s
}

func (s *CreateTemplateScratchRequestSourceTag) SetResourceTypeFilter(v []*string) *CreateTemplateScratchRequestSourceTag {
	s.ResourceTypeFilter = v
	return s
}

func (s *CreateTemplateScratchRequestSourceTag) Validate() error {
	return dara.Validate(s)
}

type CreateTemplateScratchRequestTags struct {
	// The tag key of the resource scenario.
	//
	// > Tags is optional. If you want to specify Tags, you must specify Key.
	//
	// This parameter is required.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource scenario.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTemplateScratchRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateTemplateScratchRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateTemplateScratchRequestTags) SetKey(v string) *CreateTemplateScratchRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTemplateScratchRequestTags) SetValue(v string) *CreateTemplateScratchRequestTags {
	s.Value = &v
	return s
}

func (s *CreateTemplateScratchRequestTags) Validate() error {
	return dara.Validate(s)
}
