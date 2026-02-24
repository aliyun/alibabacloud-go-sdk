// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateCompliancePackRequest
	GetClientToken() *string
	SetCompliancePackName(v string) *CreateCompliancePackRequest
	GetCompliancePackName() *string
	SetCompliancePackTemplateId(v string) *CreateCompliancePackRequest
	GetCompliancePackTemplateId() *string
	SetConfigRules(v []*CreateCompliancePackRequestConfigRules) *CreateCompliancePackRequest
	GetConfigRules() []*CreateCompliancePackRequestConfigRules
	SetDefaultEnable(v bool) *CreateCompliancePackRequest
	GetDefaultEnable() *bool
	SetDescription(v string) *CreateCompliancePackRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *CreateCompliancePackRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateCompliancePackRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateCompliancePackRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*CreateCompliancePackRequestExcludeTagsScope) *CreateCompliancePackRequest
	GetExcludeTagsScope() []*CreateCompliancePackRequestExcludeTagsScope
	SetRegionIdsScope(v string) *CreateCompliancePackRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateCompliancePackRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateCompliancePackRequest
	GetResourceIdsScope() *string
	SetRiskLevel(v int32) *CreateCompliancePackRequest
	GetRiskLevel() *int32
	SetTag(v []*CreateCompliancePackRequestTag) *CreateCompliancePackRequest
	GetTag() []*CreateCompliancePackRequestTag
	SetTagKeyScope(v string) *CreateCompliancePackRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *CreateCompliancePackRequest
	GetTagValueScope() *string
	SetTagsScope(v []*CreateCompliancePackRequestTagsScope) *CreateCompliancePackRequest
	GetTagsScope() []*CreateCompliancePackRequestTagsScope
	SetTemplateContent(v string) *CreateCompliancePackRequest
	GetTemplateContent() *string
}

type CreateCompliancePackRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CompliancePackName       *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// if can be null:
	// false
	ConfigRules                  []*CreateCompliancePackRequestConfigRules      `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	DefaultEnable                *bool                                          `json:"DefaultEnable,omitempty" xml:"DefaultEnable,omitempty"`
	Description                  *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeRegionIdsScope        *string                                        `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	ExcludeResourceGroupIdsScope *string                                        `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	ExcludeResourceIdsScope      *string                                        `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	ExcludeTagsScope             []*CreateCompliancePackRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	RegionIdsScope               *string                                        `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope        *string                                        `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceIdsScope             *string                                        `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	RiskLevel                    *int32                                         `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Tag                          []*CreateCompliancePackRequestTag              `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TagKeyScope                  *string                                        `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope                *string                                        `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	TagsScope                    []*CreateCompliancePackRequestTagsScope        `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
	TemplateContent              *string                                        `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
}

func (s CreateCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCompliancePackRequest) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *CreateCompliancePackRequest) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *CreateCompliancePackRequest) GetConfigRules() []*CreateCompliancePackRequestConfigRules {
	return s.ConfigRules
}

func (s *CreateCompliancePackRequest) GetDefaultEnable() *bool {
	return s.DefaultEnable
}

func (s *CreateCompliancePackRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCompliancePackRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateCompliancePackRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateCompliancePackRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateCompliancePackRequest) GetExcludeTagsScope() []*CreateCompliancePackRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *CreateCompliancePackRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateCompliancePackRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateCompliancePackRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateCompliancePackRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateCompliancePackRequest) GetTag() []*CreateCompliancePackRequestTag {
	return s.Tag
}

func (s *CreateCompliancePackRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *CreateCompliancePackRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *CreateCompliancePackRequest) GetTagsScope() []*CreateCompliancePackRequestTagsScope {
	return s.TagsScope
}

func (s *CreateCompliancePackRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *CreateCompliancePackRequest) SetClientToken(v string) *CreateCompliancePackRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCompliancePackRequest) SetCompliancePackName(v string) *CreateCompliancePackRequest {
	s.CompliancePackName = &v
	return s
}

func (s *CreateCompliancePackRequest) SetCompliancePackTemplateId(v string) *CreateCompliancePackRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *CreateCompliancePackRequest) SetConfigRules(v []*CreateCompliancePackRequestConfigRules) *CreateCompliancePackRequest {
	s.ConfigRules = v
	return s
}

func (s *CreateCompliancePackRequest) SetDefaultEnable(v bool) *CreateCompliancePackRequest {
	s.DefaultEnable = &v
	return s
}

func (s *CreateCompliancePackRequest) SetDescription(v string) *CreateCompliancePackRequest {
	s.Description = &v
	return s
}

func (s *CreateCompliancePackRequest) SetExcludeRegionIdsScope(v string) *CreateCompliancePackRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateCompliancePackRequest) SetExcludeResourceGroupIdsScope(v string) *CreateCompliancePackRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateCompliancePackRequest) SetExcludeResourceIdsScope(v string) *CreateCompliancePackRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateCompliancePackRequest) SetExcludeTagsScope(v []*CreateCompliancePackRequestExcludeTagsScope) *CreateCompliancePackRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *CreateCompliancePackRequest) SetRegionIdsScope(v string) *CreateCompliancePackRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateCompliancePackRequest) SetResourceGroupIdsScope(v string) *CreateCompliancePackRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateCompliancePackRequest) SetResourceIdsScope(v string) *CreateCompliancePackRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateCompliancePackRequest) SetRiskLevel(v int32) *CreateCompliancePackRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateCompliancePackRequest) SetTag(v []*CreateCompliancePackRequestTag) *CreateCompliancePackRequest {
	s.Tag = v
	return s
}

func (s *CreateCompliancePackRequest) SetTagKeyScope(v string) *CreateCompliancePackRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateCompliancePackRequest) SetTagValueScope(v string) *CreateCompliancePackRequest {
	s.TagValueScope = &v
	return s
}

func (s *CreateCompliancePackRequest) SetTagsScope(v []*CreateCompliancePackRequestTagsScope) *CreateCompliancePackRequest {
	s.TagsScope = v
	return s
}

func (s *CreateCompliancePackRequest) SetTemplateContent(v string) *CreateCompliancePackRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateCompliancePackRequest) Validate() error {
	if s.ConfigRules != nil {
		for _, item := range s.ConfigRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExcludeTagsScope != nil {
		for _, item := range s.ExcludeTagsScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
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

type CreateCompliancePackRequestConfigRules struct {
	ConfigRuleId          *string                                                       `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName        *string                                                       `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleParameters  []*CreateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	Description           *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagedRuleIdentifier *string                                                       `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	RiskLevel             *int32                                                        `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s CreateCompliancePackRequestConfigRules) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackRequestConfigRules) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequestConfigRules) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *CreateCompliancePackRequestConfigRules) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *CreateCompliancePackRequestConfigRules) GetConfigRuleParameters() []*CreateCompliancePackRequestConfigRulesConfigRuleParameters {
	return s.ConfigRuleParameters
}

func (s *CreateCompliancePackRequestConfigRules) GetDescription() *string {
	return s.Description
}

func (s *CreateCompliancePackRequestConfigRules) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *CreateCompliancePackRequestConfigRules) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateCompliancePackRequestConfigRules) SetConfigRuleId(v string) *CreateCompliancePackRequestConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetConfigRuleName(v string) *CreateCompliancePackRequestConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetConfigRuleParameters(v []*CreateCompliancePackRequestConfigRulesConfigRuleParameters) *CreateCompliancePackRequestConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetDescription(v string) *CreateCompliancePackRequestConfigRules {
	s.Description = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetManagedRuleIdentifier(v string) *CreateCompliancePackRequestConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetRiskLevel(v int32) *CreateCompliancePackRequestConfigRules {
	s.RiskLevel = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) Validate() error {
	if s.ConfigRuleParameters != nil {
		for _, item := range s.ConfigRuleParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCompliancePackRequestConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateCompliancePackRequestConfigRulesConfigRuleParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackRequestConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequestConfigRulesConfigRuleParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *CreateCompliancePackRequestConfigRulesConfigRuleParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterName(v string) *CreateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterValue(v string) *CreateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRulesConfigRuleParameters) Validate() error {
	return dara.Validate(s)
}

type CreateCompliancePackRequestExcludeTagsScope struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateCompliancePackRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateCompliancePackRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateCompliancePackRequestExcludeTagsScope) SetTagKey(v string) *CreateCompliancePackRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateCompliancePackRequestExcludeTagsScope) SetTagValue(v string) *CreateCompliancePackRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateCompliancePackRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type CreateCompliancePackRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCompliancePackRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCompliancePackRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCompliancePackRequestTag) SetKey(v string) *CreateCompliancePackRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCompliancePackRequestTag) SetValue(v string) *CreateCompliancePackRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCompliancePackRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateCompliancePackRequestTagsScope struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateCompliancePackRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackRequestTagsScope) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateCompliancePackRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateCompliancePackRequestTagsScope) SetTagKey(v string) *CreateCompliancePackRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateCompliancePackRequestTagsScope) SetTagValue(v string) *CreateCompliancePackRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateCompliancePackRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
