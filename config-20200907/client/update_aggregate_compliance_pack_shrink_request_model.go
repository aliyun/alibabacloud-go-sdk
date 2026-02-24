// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateCompliancePackShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetAggregatorId() *string
	SetClientToken(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetClientToken() *string
	SetCompliancePackId(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetCompliancePackId() *string
	SetCompliancePackName(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetCompliancePackName() *string
	SetConfigRulesShrink(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetConfigRulesShrink() *string
	SetDescription(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) *UpdateAggregateCompliancePackShrinkRequest
	GetExcludeTagsScope() []*UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope
	SetRegionIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetResourceIdsScope() *string
	SetRiskLevel(v int32) *UpdateAggregateCompliancePackShrinkRequest
	GetRiskLevel() *int32
	SetTagShrink(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetTagShrink() *string
	SetTagKeyScope(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *UpdateAggregateCompliancePackShrinkRequest
	GetTagValueScope() *string
	SetTagsScope(v []*UpdateAggregateCompliancePackShrinkRequestTagsScope) *UpdateAggregateCompliancePackShrinkRequest
	GetTagsScope() []*UpdateAggregateCompliancePackShrinkRequestTagsScope
}

type UpdateAggregateCompliancePackShrinkRequest struct {
	// The ID of the account group.
	//
	// For more information, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// A client token. It is used to ensure the idempotence of the request. Generate a value that is unique among different requests. The `ClientToken` parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the compliance pack.
	//
	// For more information, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance pack.
	//
	// For more information, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// example:
	//
	// 等保三级预检合规包
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The rules in the compliance pack.
	//
	// If you leave this parameter empty when you modify the compliance pack, the existing rules are not changed. If you specify new rules, the new rules replace the existing ones.
	ConfigRulesShrink *string `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
	// The description of the compliance pack.
	//
	// example:
	//
	// 基于等保2.0三级标准，提供持续检测合规性的建议模板，帮助您提前自检并修复问题，以便快速通过正式检测。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rules are not effective for resources in the specified regions. Resources in these regions are not evaluated. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The rules are not effective for resources in the specified resource groups. Resources in these resource groups are not evaluated. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The compliance pack is not effective for the specified resources. The specified resources are not evaluated. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// eip-8vbf3x310fn56ijfd****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The excluded tag scope.
	ExcludeTagsScope []*UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The compliance pack is effective only for resources in the specified regions. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The compliance pack is effective only for resources in the specified resource groups. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The rules are effective only for the specified resources. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The risk level of the compliance pack. Valid values:
	//
	// - 1: high risk.
	//
	// - 2: medium risk.
	//
	// - 3: low risk.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Deprecated
	//
	// The tags of the resource. This parameter is deprecated and no longer takes effect.
	//
	// You can add up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The compliance pack is effective only for resources that have the specified tag key.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// The compliance pack is effective only for resources that have the specified tag key and tag value.
	//
	// > You must specify TagValueScope together with TagKeyScope.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tag scope.
	TagsScope []*UpdateAggregateCompliancePackShrinkRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s UpdateAggregateCompliancePackShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetConfigRulesShrink() *string {
	return s.ConfigRulesShrink
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetExcludeTagsScope() []*UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) GetTagsScope() []*UpdateAggregateCompliancePackShrinkRequestTagsScope {
	return s.TagsScope
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetAggregatorId(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetClientToken(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetCompliancePackId(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetCompliancePackName(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.CompliancePackName = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetConfigRulesShrink(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ConfigRulesShrink = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetDescription(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetExcludeRegionIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetExcludeResourceIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetExcludeTagsScope(v []*UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) *UpdateAggregateCompliancePackShrinkRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetRegionIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetResourceGroupIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetResourceIdsScope(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetRiskLevel(v int32) *UpdateAggregateCompliancePackShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetTagShrink(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetTagKeyScope(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetTagValueScope(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.TagValueScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetTagsScope(v []*UpdateAggregateCompliancePackShrinkRequestTagsScope) *UpdateAggregateCompliancePackShrinkRequest {
	s.TagsScope = v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) Validate() error {
	if s.ExcludeTagsScope != nil {
		for _, item := range s.ExcludeTagsScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagsScope != nil {
		for _, item := range s.TagsScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope struct {
	// The tag key.
	//
	// example:
	//
	// key-2
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value-2
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) SetTagKey(v string) *UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) SetTagValue(v string) *UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type UpdateAggregateCompliancePackShrinkRequestTagsScope struct {
	// The tag key.
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UpdateAggregateCompliancePackShrinkRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackShrinkRequestTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackShrinkRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAggregateCompliancePackShrinkRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAggregateCompliancePackShrinkRequestTagsScope) SetTagKey(v string) *UpdateAggregateCompliancePackShrinkRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequestTagsScope) SetTagValue(v string) *UpdateAggregateCompliancePackShrinkRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
