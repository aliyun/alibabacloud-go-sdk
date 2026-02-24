// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCompliancePackRequest
	GetClientToken() *string
	SetCompliancePackId(v string) *UpdateCompliancePackRequest
	GetCompliancePackId() *string
	SetCompliancePackName(v string) *UpdateCompliancePackRequest
	GetCompliancePackName() *string
	SetConfigRules(v []*UpdateCompliancePackRequestConfigRules) *UpdateCompliancePackRequest
	GetConfigRules() []*UpdateCompliancePackRequestConfigRules
	SetDescription(v string) *UpdateCompliancePackRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *UpdateCompliancePackRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateCompliancePackRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateCompliancePackRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*UpdateCompliancePackRequestExcludeTagsScope) *UpdateCompliancePackRequest
	GetExcludeTagsScope() []*UpdateCompliancePackRequestExcludeTagsScope
	SetRegionIdsScope(v string) *UpdateCompliancePackRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateCompliancePackRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateCompliancePackRequest
	GetResourceIdsScope() *string
	SetRiskLevel(v int32) *UpdateCompliancePackRequest
	GetRiskLevel() *int32
	SetTag(v []*UpdateCompliancePackRequestTag) *UpdateCompliancePackRequest
	GetTag() []*UpdateCompliancePackRequestTag
	SetTagKeyScope(v string) *UpdateCompliancePackRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *UpdateCompliancePackRequest
	GetTagValueScope() *string
	SetTagsScope(v []*UpdateCompliancePackRequestTagsScope) *UpdateCompliancePackRequest
	GetTagsScope() []*UpdateCompliancePackRequestTagsScope
}

type UpdateCompliancePackRequest struct {
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
	ConfigRules []*UpdateCompliancePackRequestConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
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
	ExcludeTagsScope []*UpdateCompliancePackRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	Tag []*UpdateCompliancePackRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	TagsScope []*UpdateCompliancePackRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s UpdateCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCompliancePackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *UpdateCompliancePackRequest) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *UpdateCompliancePackRequest) GetConfigRules() []*UpdateCompliancePackRequestConfigRules {
	return s.ConfigRules
}

func (s *UpdateCompliancePackRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCompliancePackRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateCompliancePackRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateCompliancePackRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateCompliancePackRequest) GetExcludeTagsScope() []*UpdateCompliancePackRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *UpdateCompliancePackRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateCompliancePackRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateCompliancePackRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateCompliancePackRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateCompliancePackRequest) GetTag() []*UpdateCompliancePackRequestTag {
	return s.Tag
}

func (s *UpdateCompliancePackRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *UpdateCompliancePackRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *UpdateCompliancePackRequest) GetTagsScope() []*UpdateCompliancePackRequestTagsScope {
	return s.TagsScope
}

func (s *UpdateCompliancePackRequest) SetClientToken(v string) *UpdateCompliancePackRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetCompliancePackId(v string) *UpdateCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetCompliancePackName(v string) *UpdateCompliancePackRequest {
	s.CompliancePackName = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetConfigRules(v []*UpdateCompliancePackRequestConfigRules) *UpdateCompliancePackRequest {
	s.ConfigRules = v
	return s
}

func (s *UpdateCompliancePackRequest) SetDescription(v string) *UpdateCompliancePackRequest {
	s.Description = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetExcludeRegionIdsScope(v string) *UpdateCompliancePackRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateCompliancePackRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetExcludeResourceIdsScope(v string) *UpdateCompliancePackRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetExcludeTagsScope(v []*UpdateCompliancePackRequestExcludeTagsScope) *UpdateCompliancePackRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *UpdateCompliancePackRequest) SetRegionIdsScope(v string) *UpdateCompliancePackRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetResourceGroupIdsScope(v string) *UpdateCompliancePackRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetResourceIdsScope(v string) *UpdateCompliancePackRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetRiskLevel(v int32) *UpdateCompliancePackRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetTag(v []*UpdateCompliancePackRequestTag) *UpdateCompliancePackRequest {
	s.Tag = v
	return s
}

func (s *UpdateCompliancePackRequest) SetTagKeyScope(v string) *UpdateCompliancePackRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetTagValueScope(v string) *UpdateCompliancePackRequest {
	s.TagValueScope = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetTagsScope(v []*UpdateCompliancePackRequestTagsScope) *UpdateCompliancePackRequest {
	s.TagsScope = v
	return s
}

func (s *UpdateCompliancePackRequest) Validate() error {
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

type UpdateCompliancePackRequestConfigRules struct {
	// The rule ID. CloudConfig adds an existing rule to the compliance pack.
	//
	// You must specify either `ManagedRuleIdentifier` or `ConfigRuleId`. If you specify both parameters, `ConfigRuleId` takes precedence. For more information, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
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
	ConfigRuleParameters []*UpdateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// 弹性公网已绑定到ECS或者NAT实例，非闲置状态，视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the rule template. CloudConfig automatically creates a rule based on this identifier and adds the rule to the compliance pack.
	//
	// You must specify either `ManagedRuleIdentifier` or `ConfigRuleId`. If you specify both parameters, `ConfigRuleId` takes precedence. For more information, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// eip-bandwidth-limit
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The risk level of the rule. Valid values:
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
}

func (s UpdateCompliancePackRequestConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackRequestConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequestConfigRules) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *UpdateCompliancePackRequestConfigRules) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *UpdateCompliancePackRequestConfigRules) GetConfigRuleParameters() []*UpdateCompliancePackRequestConfigRulesConfigRuleParameters {
	return s.ConfigRuleParameters
}

func (s *UpdateCompliancePackRequestConfigRules) GetDescription() *string {
	return s.Description
}

func (s *UpdateCompliancePackRequestConfigRules) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *UpdateCompliancePackRequestConfigRules) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateCompliancePackRequestConfigRules) SetConfigRuleId(v string) *UpdateCompliancePackRequestConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetConfigRuleName(v string) *UpdateCompliancePackRequestConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetConfigRuleParameters(v []*UpdateCompliancePackRequestConfigRulesConfigRuleParameters) *UpdateCompliancePackRequestConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetDescription(v string) *UpdateCompliancePackRequestConfigRules {
	s.Description = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetManagedRuleIdentifier(v string) *UpdateCompliancePackRequestConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetRiskLevel(v int32) *UpdateCompliancePackRequestConfigRules {
	s.RiskLevel = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) Validate() error {
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

type UpdateCompliancePackRequestConfigRulesConfigRuleParameters struct {
	// The name of the rule parameter.
	//
	// You must specify `ParameterName` and `ParameterValue` together, or leave them both empty. If a rule template has a parameter without a default value, you must specify this parameter. For more information, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// bandwidth
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the rule parameter.
	//
	// You must specify `ParameterName` and `ParameterValue` together, or leave them both empty. If a rule template has a parameter without a default value, you must specify this parameter. For more information, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// 20
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateCompliancePackRequestConfigRulesConfigRuleParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackRequestConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequestConfigRulesConfigRuleParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *UpdateCompliancePackRequestConfigRulesConfigRuleParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterName(v string) *UpdateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterValue(v string) *UpdateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRulesConfigRuleParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateCompliancePackRequestExcludeTagsScope struct {
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

func (s UpdateCompliancePackRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateCompliancePackRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateCompliancePackRequestExcludeTagsScope) SetTagKey(v string) *UpdateCompliancePackRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateCompliancePackRequestExcludeTagsScope) SetTagValue(v string) *UpdateCompliancePackRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateCompliancePackRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type UpdateCompliancePackRequestTag struct {
	// The tag key of the resource.
	//
	// If you specify this parameter, it cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https:// `.
	//
	// You can specify the tag keys of up to 20 tags at a time.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. If you specify this parameter, it can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateCompliancePackRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequestTag) GetKey() *string {
	return s.Key
}

func (s *UpdateCompliancePackRequestTag) GetValue() *string {
	return s.Value
}

func (s *UpdateCompliancePackRequestTag) SetKey(v string) *UpdateCompliancePackRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateCompliancePackRequestTag) SetValue(v string) *UpdateCompliancePackRequestTag {
	s.Value = &v
	return s
}

func (s *UpdateCompliancePackRequestTag) Validate() error {
	return dara.Validate(s)
}

type UpdateCompliancePackRequestTagsScope struct {
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

func (s UpdateCompliancePackRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackRequestTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateCompliancePackRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateCompliancePackRequestTagsScope) SetTagKey(v string) *UpdateCompliancePackRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateCompliancePackRequestTagsScope) SetTagValue(v string) *UpdateCompliancePackRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateCompliancePackRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
