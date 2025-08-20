// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateScratchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTemplateScratchShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateTemplateScratchShrinkRequest
	GetDescription() *string
	SetExecutionMode(v string) *UpdateTemplateScratchShrinkRequest
	GetExecutionMode() *string
	SetLogicalIdStrategy(v string) *UpdateTemplateScratchShrinkRequest
	GetLogicalIdStrategy() *string
	SetPreferenceParametersShrink(v string) *UpdateTemplateScratchShrinkRequest
	GetPreferenceParametersShrink() *string
	SetRegionId(v string) *UpdateTemplateScratchShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateTemplateScratchShrinkRequest
	GetResourceGroupId() *string
	SetSourceResourceGroupShrink(v string) *UpdateTemplateScratchShrinkRequest
	GetSourceResourceGroupShrink() *string
	SetSourceResourcesShrink(v string) *UpdateTemplateScratchShrinkRequest
	GetSourceResourcesShrink() *string
	SetSourceTagShrink(v string) *UpdateTemplateScratchShrinkRequest
	GetSourceTagShrink() *string
	SetTemplateScratchId(v string) *UpdateTemplateScratchShrinkRequest
	GetTemplateScratchId() *string
}

type UpdateTemplateScratchShrinkRequest struct {
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
	PreferenceParametersShrink *string `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty"`
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
	SourceResourceGroupShrink *string `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty"`
	// The source resources.
	//
	// >  You must specify only one of the following parameters: SourceResources, SourceTag, and SourceResourceGroup.
	SourceResourcesShrink *string `json:"SourceResources,omitempty" xml:"SourceResources,omitempty"`
	// The source tag.
	//
	// >  You must specify only one of the following parameters: SourceResources, SourceTag, and SourceResourceGroup.
	SourceTagShrink *string `json:"SourceTag,omitempty" xml:"SourceTag,omitempty"`
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

func (s UpdateTemplateScratchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateScratchShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTemplateScratchShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTemplateScratchShrinkRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *UpdateTemplateScratchShrinkRequest) GetLogicalIdStrategy() *string {
	return s.LogicalIdStrategy
}

func (s *UpdateTemplateScratchShrinkRequest) GetPreferenceParametersShrink() *string {
	return s.PreferenceParametersShrink
}

func (s *UpdateTemplateScratchShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTemplateScratchShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTemplateScratchShrinkRequest) GetSourceResourceGroupShrink() *string {
	return s.SourceResourceGroupShrink
}

func (s *UpdateTemplateScratchShrinkRequest) GetSourceResourcesShrink() *string {
	return s.SourceResourcesShrink
}

func (s *UpdateTemplateScratchShrinkRequest) GetSourceTagShrink() *string {
	return s.SourceTagShrink
}

func (s *UpdateTemplateScratchShrinkRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *UpdateTemplateScratchShrinkRequest) SetClientToken(v string) *UpdateTemplateScratchShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetDescription(v string) *UpdateTemplateScratchShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetExecutionMode(v string) *UpdateTemplateScratchShrinkRequest {
	s.ExecutionMode = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetLogicalIdStrategy(v string) *UpdateTemplateScratchShrinkRequest {
	s.LogicalIdStrategy = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetPreferenceParametersShrink(v string) *UpdateTemplateScratchShrinkRequest {
	s.PreferenceParametersShrink = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetRegionId(v string) *UpdateTemplateScratchShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetResourceGroupId(v string) *UpdateTemplateScratchShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetSourceResourceGroupShrink(v string) *UpdateTemplateScratchShrinkRequest {
	s.SourceResourceGroupShrink = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetSourceResourcesShrink(v string) *UpdateTemplateScratchShrinkRequest {
	s.SourceResourcesShrink = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetSourceTagShrink(v string) *UpdateTemplateScratchShrinkRequest {
	s.SourceTagShrink = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetTemplateScratchId(v string) *UpdateTemplateScratchShrinkRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
