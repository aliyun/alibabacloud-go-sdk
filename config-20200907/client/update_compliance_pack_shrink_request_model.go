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
	// The client token that you want to use to ensure the idempotency of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.``
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-a8a8626622af0082****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance package.
	//
	// For more information about how to obtain the name of a compliance package, see [ListCompliancePacks](https://help.aliyun.com/document_detail/263332.html).
	//
	// example:
	//
	// The name of the compliance package.
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The rules in the compliance package.
	//
	// If you leave this parameter empty, the rules in the compliance package remain unchanged. If you configure this parameter, Cloud Config replaces the existing rules in the compliance package with the specified rules.
	ConfigRulesShrink *string `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
	// The description of the compliance package.
	//
	// example:
	//
	// The description of the compliance package.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the regions to which the rule not applies. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// ExcludeResourceGroupIdsScope. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The ID of the resource that you do not want to evaluate by using the compliance package. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// 23642660635687****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// ExcludeTagsScope
	ExcludeTagsScope []*UpdateCompliancePackShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The ID of the region whose resources you want to evaluate by using the compliance package. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The ID of the resource group whose resources you want to evaluate by using the compliance package. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzdibsjjc****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The IDs of the resources included from the compliance evaluations performed by the rule. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The risk level of the resources that are not compliant with the rules in the compliance package. Valid values:
	//
	// 	- 1: high risk level
	//
	// 	- 2: medium risk level
	//
	// 	- 3: low risk level
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The tag key of the resource that you want to evaluate by using the compliance package.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// The tag value of the resource that you want to evaluate by using the compliance package.
	//
	// >  You must configure the TagValueScope parameter together with the TagValueScope parameter.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// TagsScope
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
