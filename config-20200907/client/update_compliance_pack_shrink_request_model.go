// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompliancePackShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCompliancePackShrinkRequest
	GetClientToken() *string
	SetCompliancePackId(v string) *UpdateCompliancePackShrinkRequest
	GetCompliancePackId() *string
	SetCompliancePackName(v string) *UpdateCompliancePackShrinkRequest
	GetCompliancePackName() *string
	SetConfigRulesShrink(v string) *UpdateCompliancePackShrinkRequest
	GetConfigRulesShrink() *string
	SetDescription(v string) *UpdateCompliancePackShrinkRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *UpdateCompliancePackShrinkRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateCompliancePackShrinkRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateCompliancePackShrinkRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*UpdateCompliancePackShrinkRequestExcludeTagsScope) *UpdateCompliancePackShrinkRequest
	GetExcludeTagsScope() []*UpdateCompliancePackShrinkRequestExcludeTagsScope
	SetRegionIdsScope(v string) *UpdateCompliancePackShrinkRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateCompliancePackShrinkRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateCompliancePackShrinkRequest
	GetResourceIdsScope() *string
	SetRiskLevel(v int32) *UpdateCompliancePackShrinkRequest
	GetRiskLevel() *int32
	SetTagShrink(v string) *UpdateCompliancePackShrinkRequest
	GetTagShrink() *string
	SetTagKeyScope(v string) *UpdateCompliancePackShrinkRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *UpdateCompliancePackShrinkRequest
	GetTagValueScope() *string
	SetTagsScope(v []*UpdateCompliancePackShrinkRequestTagsScope) *UpdateCompliancePackShrinkRequest
	GetTagsScope() []*UpdateCompliancePackShrinkRequestTagsScope
}

type UpdateCompliancePackShrinkRequest struct {
	// A client token to ensure the idempotence of the request. Generate a unique token for each request. The `ClientToken` value can contain only ASCII characters and must be no more than 64 characters long.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the compliance pack.
	//
	// For more information, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-a8a8626622af0082****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance pack.
	//
	// For more information, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// example:
	//
	// 等保三级预检合规包
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The rules in the compliance pack.
	//
	// If you leave this parameter empty when you modify the compliance pack, the original rules are retained. If you specify new rules, they replace the original rules.
	ConfigRulesShrink *string `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
	// The description of the compliance pack.
	//
	// example:
	//
	// 基于等保2.0三级标准，提供持续检测合规性的建议模板，帮助您提前自检并修复问题，以便快速通过正式检测。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The compliance pack does not evaluate resources in the specified regions. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The compliance pack does not evaluate resources in the specified resource groups. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The compliance pack does not evaluate the specified resources. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// 23642660635687****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The excluded tag scope.
	ExcludeTagsScope []*UpdateCompliancePackShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The compliance pack evaluates only resources in the specified regions. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The compliance pack evaluates only resources in the specified resource groups. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzdibsjjc****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The compliance pack evaluates only the specified resources. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The risk level of the compliance pack. Valid values:
	//
	// - 1: High risk.
	//
	// - 2: Medium risk.
	//
	// - 3: Low risk.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Deprecated
	//
	// The tags of the resource. This parameter is deprecated. Ignore this parameter because it is no longer valid.
	//
	// You can add up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The compliance pack evaluates only resources that have the specified tag key.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// The compliance pack evaluates only resources that have the specified tag key and value.
	//
	// > You must use TagValueScope with TagKeyScope.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tag scope.
	TagsScope []*UpdateCompliancePackShrinkRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s UpdateCompliancePackShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCompliancePackShrinkRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *UpdateCompliancePackShrinkRequest) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *UpdateCompliancePackShrinkRequest) GetConfigRulesShrink() *string {
	return s.ConfigRulesShrink
}

func (s *UpdateCompliancePackShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCompliancePackShrinkRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateCompliancePackShrinkRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateCompliancePackShrinkRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateCompliancePackShrinkRequest) GetExcludeTagsScope() []*UpdateCompliancePackShrinkRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *UpdateCompliancePackShrinkRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateCompliancePackShrinkRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateCompliancePackShrinkRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateCompliancePackShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateCompliancePackShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *UpdateCompliancePackShrinkRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *UpdateCompliancePackShrinkRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *UpdateCompliancePackShrinkRequest) GetTagsScope() []*UpdateCompliancePackShrinkRequestTagsScope {
	return s.TagsScope
}

func (s *UpdateCompliancePackShrinkRequest) SetClientToken(v string) *UpdateCompliancePackShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetCompliancePackId(v string) *UpdateCompliancePackShrinkRequest {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetCompliancePackName(v string) *UpdateCompliancePackShrinkRequest {
	s.CompliancePackName = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetConfigRulesShrink(v string) *UpdateCompliancePackShrinkRequest {
	s.ConfigRulesShrink = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetDescription(v string) *UpdateCompliancePackShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetExcludeRegionIdsScope(v string) *UpdateCompliancePackShrinkRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateCompliancePackShrinkRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetExcludeResourceIdsScope(v string) *UpdateCompliancePackShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetExcludeTagsScope(v []*UpdateCompliancePackShrinkRequestExcludeTagsScope) *UpdateCompliancePackShrinkRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetRegionIdsScope(v string) *UpdateCompliancePackShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetResourceGroupIdsScope(v string) *UpdateCompliancePackShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetResourceIdsScope(v string) *UpdateCompliancePackShrinkRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetRiskLevel(v int32) *UpdateCompliancePackShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetTagShrink(v string) *UpdateCompliancePackShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetTagKeyScope(v string) *UpdateCompliancePackShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetTagValueScope(v string) *UpdateCompliancePackShrinkRequest {
	s.TagValueScope = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetTagsScope(v []*UpdateCompliancePackShrinkRequestTagsScope) *UpdateCompliancePackShrinkRequest {
	s.TagsScope = v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) Validate() error {
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

type UpdateCompliancePackShrinkRequestExcludeTagsScope struct {
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

func (s UpdateCompliancePackShrinkRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackShrinkRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackShrinkRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateCompliancePackShrinkRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateCompliancePackShrinkRequestExcludeTagsScope) SetTagKey(v string) *UpdateCompliancePackShrinkRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequestExcludeTagsScope) SetTagValue(v string) *UpdateCompliancePackShrinkRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type UpdateCompliancePackShrinkRequestTagsScope struct {
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

func (s UpdateCompliancePackShrinkRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackShrinkRequestTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackShrinkRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateCompliancePackShrinkRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateCompliancePackShrinkRequestTagsScope) SetTagKey(v string) *UpdateCompliancePackShrinkRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequestTagsScope) SetTagValue(v string) *UpdateCompliancePackShrinkRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
