// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *UpdateAggregateCompliancePackRequest
	GetAggregatorId() *string
	SetClientToken(v string) *UpdateAggregateCompliancePackRequest
	GetClientToken() *string
	SetCompliancePackId(v string) *UpdateAggregateCompliancePackRequest
	GetCompliancePackId() *string
	SetCompliancePackName(v string) *UpdateAggregateCompliancePackRequest
	GetCompliancePackName() *string
	SetConfigRules(v []*UpdateAggregateCompliancePackRequestConfigRules) *UpdateAggregateCompliancePackRequest
	GetConfigRules() []*UpdateAggregateCompliancePackRequestConfigRules
	SetDescription(v string) *UpdateAggregateCompliancePackRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *UpdateAggregateCompliancePackRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateAggregateCompliancePackRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateAggregateCompliancePackRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*UpdateAggregateCompliancePackRequestExcludeTagsScope) *UpdateAggregateCompliancePackRequest
	GetExcludeTagsScope() []*UpdateAggregateCompliancePackRequestExcludeTagsScope
	SetRegionIdsScope(v string) *UpdateAggregateCompliancePackRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateAggregateCompliancePackRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateAggregateCompliancePackRequest
	GetResourceIdsScope() *string
	SetRiskLevel(v int32) *UpdateAggregateCompliancePackRequest
	GetRiskLevel() *int32
	SetTag(v []*UpdateAggregateCompliancePackRequestTag) *UpdateAggregateCompliancePackRequest
	GetTag() []*UpdateAggregateCompliancePackRequestTag
	SetTagKeyScope(v string) *UpdateAggregateCompliancePackRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *UpdateAggregateCompliancePackRequest
	GetTagValueScope() *string
	SetTagsScope(v []*UpdateAggregateCompliancePackRequestTagsScope) *UpdateAggregateCompliancePackRequest
	GetTagsScope() []*UpdateAggregateCompliancePackRequestTagsScope
}

type UpdateAggregateCompliancePackRequest struct {
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
	ConfigRules []*UpdateAggregateCompliancePackRequestConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
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
	ExcludeTagsScope []*UpdateAggregateCompliancePackRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	Tag []*UpdateAggregateCompliancePackRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	TagsScope []*UpdateAggregateCompliancePackRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s UpdateAggregateCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregateCompliancePackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAggregateCompliancePackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *UpdateAggregateCompliancePackRequest) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *UpdateAggregateCompliancePackRequest) GetConfigRules() []*UpdateAggregateCompliancePackRequestConfigRules {
	return s.ConfigRules
}

func (s *UpdateAggregateCompliancePackRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggregateCompliancePackRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateAggregateCompliancePackRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateAggregateCompliancePackRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateAggregateCompliancePackRequest) GetExcludeTagsScope() []*UpdateAggregateCompliancePackRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *UpdateAggregateCompliancePackRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateAggregateCompliancePackRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateAggregateCompliancePackRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateAggregateCompliancePackRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateAggregateCompliancePackRequest) GetTag() []*UpdateAggregateCompliancePackRequestTag {
	return s.Tag
}

func (s *UpdateAggregateCompliancePackRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *UpdateAggregateCompliancePackRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *UpdateAggregateCompliancePackRequest) GetTagsScope() []*UpdateAggregateCompliancePackRequestTagsScope {
	return s.TagsScope
}

func (s *UpdateAggregateCompliancePackRequest) SetAggregatorId(v string) *UpdateAggregateCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetClientToken(v string) *UpdateAggregateCompliancePackRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetCompliancePackId(v string) *UpdateAggregateCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetCompliancePackName(v string) *UpdateAggregateCompliancePackRequest {
	s.CompliancePackName = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetConfigRules(v []*UpdateAggregateCompliancePackRequestConfigRules) *UpdateAggregateCompliancePackRequest {
	s.ConfigRules = v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetDescription(v string) *UpdateAggregateCompliancePackRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetExcludeRegionIdsScope(v string) *UpdateAggregateCompliancePackRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateAggregateCompliancePackRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetExcludeResourceIdsScope(v string) *UpdateAggregateCompliancePackRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetExcludeTagsScope(v []*UpdateAggregateCompliancePackRequestExcludeTagsScope) *UpdateAggregateCompliancePackRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetRegionIdsScope(v string) *UpdateAggregateCompliancePackRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetResourceGroupIdsScope(v string) *UpdateAggregateCompliancePackRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetResourceIdsScope(v string) *UpdateAggregateCompliancePackRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetRiskLevel(v int32) *UpdateAggregateCompliancePackRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetTag(v []*UpdateAggregateCompliancePackRequestTag) *UpdateAggregateCompliancePackRequest {
	s.Tag = v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetTagKeyScope(v string) *UpdateAggregateCompliancePackRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetTagValueScope(v string) *UpdateAggregateCompliancePackRequest {
	s.TagValueScope = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetTagsScope(v []*UpdateAggregateCompliancePackRequestTagsScope) *UpdateAggregateCompliancePackRequest {
	s.TagsScope = v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) Validate() error {
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

type UpdateAggregateCompliancePackRequestConfigRules struct {
	// The rule ID. CloudConfig adds the existing rule to the compliance pack.
	//
	// You must specify either `ManagedRuleIdentifier` or `ConfigRuleId`. If you specify both parameters, `ConfigRuleId` takes precedence. For more information, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// cr-e918626622af000f****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// 检测闲置弹性公网IP
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The parameters of the rule.
	ConfigRuleParameters []*UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// 弹性公网已绑定到ECS或者NAT实例，非闲置状态，视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the rule template. CloudConfig automatically creates a rule based on the rule template identifier and adds the rule to the compliance pack.
	//
	// You must specify either `ManagedRuleIdentifier` or `ConfigRuleId`. If you specify both parameters, `ConfigRuleId` takes precedence. For more information, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// eip-bandwidth-limit
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The risk level of the rule. Valid values:
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
}

func (s UpdateAggregateCompliancePackRequestConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequestConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) GetConfigRuleParameters() []*UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	return s.ConfigRuleParameters
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetConfigRuleId(v string) *UpdateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetConfigRuleName(v string) *UpdateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetConfigRuleParameters(v []*UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) *UpdateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetDescription(v string) *UpdateAggregateCompliancePackRequestConfigRules {
	s.Description = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetManagedRuleIdentifier(v string) *UpdateAggregateCompliancePackRequestConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetRiskLevel(v int32) *UpdateAggregateCompliancePackRequestConfigRules {
	s.RiskLevel = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) Validate() error {
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

type UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters struct {
	// The name of the rule parameter.
	//
	// You must specify `ParameterName` and `ParameterValue` together, or leave both empty. If a rule template has a parameter without a default value, you must specify the parameter. For more information, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// bandwidth
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the rule parameter.
	//
	// You must specify `ParameterName` and `ParameterValue` together, or leave both empty. If a rule template has a parameter without a default value, you must specify the parameter. For more information, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// 20
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterName(v string) *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterValue(v string) *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateAggregateCompliancePackRequestExcludeTagsScope struct {
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

func (s UpdateAggregateCompliancePackRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAggregateCompliancePackRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAggregateCompliancePackRequestExcludeTagsScope) SetTagKey(v string) *UpdateAggregateCompliancePackRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestExcludeTagsScope) SetTagValue(v string) *UpdateAggregateCompliancePackRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type UpdateAggregateCompliancePackRequestTag struct {
	// The tag key of the resource.
	//
	// You can add up to 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// You can add up to 20 tag values.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateAggregateCompliancePackRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequestTag) GetKey() *string {
	return s.Key
}

func (s *UpdateAggregateCompliancePackRequestTag) GetValue() *string {
	return s.Value
}

func (s *UpdateAggregateCompliancePackRequestTag) SetKey(v string) *UpdateAggregateCompliancePackRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestTag) SetValue(v string) *UpdateAggregateCompliancePackRequestTag {
	s.Value = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestTag) Validate() error {
	return dara.Validate(s)
}

type UpdateAggregateCompliancePackRequestTagsScope struct {
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

func (s UpdateAggregateCompliancePackRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequestTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAggregateCompliancePackRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAggregateCompliancePackRequestTagsScope) SetTagKey(v string) *UpdateAggregateCompliancePackRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestTagsScope) SetTagValue(v string) *UpdateAggregateCompliancePackRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
