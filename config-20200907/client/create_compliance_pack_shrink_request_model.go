// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompliancePackShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCompliancePackShrinkRequest
	GetClientToken() *string
	SetCompliancePackName(v string) *CreateCompliancePackShrinkRequest
	GetCompliancePackName() *string
	SetCompliancePackTemplateId(v string) *CreateCompliancePackShrinkRequest
	GetCompliancePackTemplateId() *string
	SetConfigRulesShrink(v string) *CreateCompliancePackShrinkRequest
	GetConfigRulesShrink() *string
	SetDefaultEnable(v bool) *CreateCompliancePackShrinkRequest
	GetDefaultEnable() *bool
	SetDescription(v string) *CreateCompliancePackShrinkRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *CreateCompliancePackShrinkRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateCompliancePackShrinkRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateCompliancePackShrinkRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*CreateCompliancePackShrinkRequestExcludeTagsScope) *CreateCompliancePackShrinkRequest
	GetExcludeTagsScope() []*CreateCompliancePackShrinkRequestExcludeTagsScope
	SetRegionIdsScope(v string) *CreateCompliancePackShrinkRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateCompliancePackShrinkRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateCompliancePackShrinkRequest
	GetResourceIdsScope() *string
	SetRiskLevel(v int32) *CreateCompliancePackShrinkRequest
	GetRiskLevel() *int32
	SetTagShrink(v string) *CreateCompliancePackShrinkRequest
	GetTagShrink() *string
	SetTagKeyScope(v string) *CreateCompliancePackShrinkRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *CreateCompliancePackShrinkRequest
	GetTagValueScope() *string
	SetTagsScope(v []*CreateCompliancePackShrinkRequestTagsScope) *CreateCompliancePackShrinkRequest
	GetTagsScope() []*CreateCompliancePackShrinkRequestTagsScope
	SetTemplateContent(v string) *CreateCompliancePackShrinkRequest
	GetTemplateContent() *string
}

type CreateCompliancePackShrinkRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CompliancePackName       *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// if can be null:
	// false
	ConfigRulesShrink            *string                                              `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
	DefaultEnable                *bool                                                `json:"DefaultEnable,omitempty" xml:"DefaultEnable,omitempty"`
	Description                  *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeRegionIdsScope        *string                                              `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	ExcludeResourceGroupIdsScope *string                                              `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	ExcludeResourceIdsScope      *string                                              `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	ExcludeTagsScope             []*CreateCompliancePackShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	RegionIdsScope               *string                                              `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope        *string                                              `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceIdsScope             *string                                              `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	RiskLevel                    *int32                                               `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	TagShrink                    *string                                              `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TagKeyScope                  *string                                              `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope                *string                                              `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	TagsScope                    []*CreateCompliancePackShrinkRequestTagsScope        `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
	TemplateContent              *string                                              `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
}

func (s CreateCompliancePackShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCompliancePackShrinkRequest) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *CreateCompliancePackShrinkRequest) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *CreateCompliancePackShrinkRequest) GetConfigRulesShrink() *string {
	return s.ConfigRulesShrink
}

func (s *CreateCompliancePackShrinkRequest) GetDefaultEnable() *bool {
	return s.DefaultEnable
}

func (s *CreateCompliancePackShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCompliancePackShrinkRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateCompliancePackShrinkRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateCompliancePackShrinkRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateCompliancePackShrinkRequest) GetExcludeTagsScope() []*CreateCompliancePackShrinkRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *CreateCompliancePackShrinkRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateCompliancePackShrinkRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateCompliancePackShrinkRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateCompliancePackShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateCompliancePackShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateCompliancePackShrinkRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *CreateCompliancePackShrinkRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *CreateCompliancePackShrinkRequest) GetTagsScope() []*CreateCompliancePackShrinkRequestTagsScope {
	return s.TagsScope
}

func (s *CreateCompliancePackShrinkRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *CreateCompliancePackShrinkRequest) SetClientToken(v string) *CreateCompliancePackShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetCompliancePackName(v string) *CreateCompliancePackShrinkRequest {
	s.CompliancePackName = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetCompliancePackTemplateId(v string) *CreateCompliancePackShrinkRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetConfigRulesShrink(v string) *CreateCompliancePackShrinkRequest {
	s.ConfigRulesShrink = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetDefaultEnable(v bool) *CreateCompliancePackShrinkRequest {
	s.DefaultEnable = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetDescription(v string) *CreateCompliancePackShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetExcludeRegionIdsScope(v string) *CreateCompliancePackShrinkRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetExcludeResourceGroupIdsScope(v string) *CreateCompliancePackShrinkRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetExcludeResourceIdsScope(v string) *CreateCompliancePackShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetExcludeTagsScope(v []*CreateCompliancePackShrinkRequestExcludeTagsScope) *CreateCompliancePackShrinkRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetRegionIdsScope(v string) *CreateCompliancePackShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetResourceGroupIdsScope(v string) *CreateCompliancePackShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetResourceIdsScope(v string) *CreateCompliancePackShrinkRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetRiskLevel(v int32) *CreateCompliancePackShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetTagShrink(v string) *CreateCompliancePackShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetTagKeyScope(v string) *CreateCompliancePackShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetTagValueScope(v string) *CreateCompliancePackShrinkRequest {
	s.TagValueScope = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetTagsScope(v []*CreateCompliancePackShrinkRequestTagsScope) *CreateCompliancePackShrinkRequest {
	s.TagsScope = v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetTemplateContent(v string) *CreateCompliancePackShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) Validate() error {
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

type CreateCompliancePackShrinkRequestExcludeTagsScope struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateCompliancePackShrinkRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackShrinkRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackShrinkRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateCompliancePackShrinkRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateCompliancePackShrinkRequestExcludeTagsScope) SetTagKey(v string) *CreateCompliancePackShrinkRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateCompliancePackShrinkRequestExcludeTagsScope) SetTagValue(v string) *CreateCompliancePackShrinkRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateCompliancePackShrinkRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type CreateCompliancePackShrinkRequestTagsScope struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateCompliancePackShrinkRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackShrinkRequestTagsScope) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackShrinkRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateCompliancePackShrinkRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateCompliancePackShrinkRequestTagsScope) SetTagKey(v string) *CreateCompliancePackShrinkRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateCompliancePackShrinkRequestTagsScope) SetTagValue(v string) *CreateCompliancePackShrinkRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateCompliancePackShrinkRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
