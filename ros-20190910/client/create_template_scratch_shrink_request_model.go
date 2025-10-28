// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateScratchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTemplateScratchShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTemplateScratchShrinkRequest
	GetDescription() *string
	SetExecutionMode(v string) *CreateTemplateScratchShrinkRequest
	GetExecutionMode() *string
	SetLogicalIdStrategy(v string) *CreateTemplateScratchShrinkRequest
	GetLogicalIdStrategy() *string
	SetPreferenceParametersShrink(v string) *CreateTemplateScratchShrinkRequest
	GetPreferenceParametersShrink() *string
	SetRegionId(v string) *CreateTemplateScratchShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTemplateScratchShrinkRequest
	GetResourceGroupId() *string
	SetSourceResourceGroupShrink(v string) *CreateTemplateScratchShrinkRequest
	GetSourceResourceGroupShrink() *string
	SetSourceResourcesShrink(v string) *CreateTemplateScratchShrinkRequest
	GetSourceResourcesShrink() *string
	SetSourceTagShrink(v string) *CreateTemplateScratchShrinkRequest
	GetSourceTagShrink() *string
	SetTags(v []*CreateTemplateScratchShrinkRequestTags) *CreateTemplateScratchShrinkRequest
	GetTags() []*CreateTemplateScratchShrinkRequestTags
	SetTemplateScratchType(v string) *CreateTemplateScratchShrinkRequest
	GetTemplateScratchType() *string
}

type CreateTemplateScratchShrinkRequest struct {
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
	PreferenceParametersShrink *string `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty"`
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
	SourceResourceGroupShrink *string `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty"`
	// The source resources.
	//
	// When you set TemplateScratchType to ArchitectureDetection, you can specify SourceResources to detect the architecture data of all resources associated with the specified source resources. For example, if you set SourceResources to the ID of a Classic Load Balancer (CLB) instance, the architecture data of all resources, such as the Elastic Compute Service (ECS) instance, vSwitch, and VPC, associated with the CLB instance is detected.
	//
	// If you set TemplateScratchType to ArchitectureDetection, you can specify up to 20 source resources. In other cases, you can specify up to 200 source resources.
	SourceResourcesShrink *string `json:"SourceResources,omitempty" xml:"SourceResources,omitempty"`
	// The source tag.
	SourceTagShrink *string `json:"SourceTag,omitempty" xml:"SourceTag,omitempty"`
	// The tags of the resource scenario.
	Tags []*CreateTemplateScratchShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s CreateTemplateScratchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTemplateScratchShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTemplateScratchShrinkRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *CreateTemplateScratchShrinkRequest) GetLogicalIdStrategy() *string {
	return s.LogicalIdStrategy
}

func (s *CreateTemplateScratchShrinkRequest) GetPreferenceParametersShrink() *string {
	return s.PreferenceParametersShrink
}

func (s *CreateTemplateScratchShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTemplateScratchShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTemplateScratchShrinkRequest) GetSourceResourceGroupShrink() *string {
	return s.SourceResourceGroupShrink
}

func (s *CreateTemplateScratchShrinkRequest) GetSourceResourcesShrink() *string {
	return s.SourceResourcesShrink
}

func (s *CreateTemplateScratchShrinkRequest) GetSourceTagShrink() *string {
	return s.SourceTagShrink
}

func (s *CreateTemplateScratchShrinkRequest) GetTags() []*CreateTemplateScratchShrinkRequestTags {
	return s.Tags
}

func (s *CreateTemplateScratchShrinkRequest) GetTemplateScratchType() *string {
	return s.TemplateScratchType
}

func (s *CreateTemplateScratchShrinkRequest) SetClientToken(v string) *CreateTemplateScratchShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetDescription(v string) *CreateTemplateScratchShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetExecutionMode(v string) *CreateTemplateScratchShrinkRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetLogicalIdStrategy(v string) *CreateTemplateScratchShrinkRequest {
	s.LogicalIdStrategy = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetPreferenceParametersShrink(v string) *CreateTemplateScratchShrinkRequest {
	s.PreferenceParametersShrink = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetRegionId(v string) *CreateTemplateScratchShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetResourceGroupId(v string) *CreateTemplateScratchShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetSourceResourceGroupShrink(v string) *CreateTemplateScratchShrinkRequest {
	s.SourceResourceGroupShrink = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetSourceResourcesShrink(v string) *CreateTemplateScratchShrinkRequest {
	s.SourceResourcesShrink = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetSourceTagShrink(v string) *CreateTemplateScratchShrinkRequest {
	s.SourceTagShrink = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetTags(v []*CreateTemplateScratchShrinkRequestTags) *CreateTemplateScratchShrinkRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetTemplateScratchType(v string) *CreateTemplateScratchShrinkRequest {
	s.TemplateScratchType = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) Validate() error {
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

type CreateTemplateScratchShrinkRequestTags struct {
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

func (s CreateTemplateScratchShrinkRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateScratchShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchShrinkRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateTemplateScratchShrinkRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateTemplateScratchShrinkRequestTags) SetKey(v string) *CreateTemplateScratchShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequestTags) SetValue(v string) *CreateTemplateScratchShrinkRequestTags {
	s.Value = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequestTags) Validate() error {
	return dara.Validate(s)
}
