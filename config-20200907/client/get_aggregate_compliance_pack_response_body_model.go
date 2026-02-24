// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePack(v *GetAggregateCompliancePackResponseBodyCompliancePack) *GetAggregateCompliancePackResponseBody
	GetCompliancePack() *GetAggregateCompliancePackResponseBodyCompliancePack
	SetRequestId(v string) *GetAggregateCompliancePackResponseBody
	GetRequestId() *string
}

type GetAggregateCompliancePackResponseBody struct {
	// The details of the compliance pack.
	CompliancePack *GetAggregateCompliancePackResponseBodyCompliancePack `json:"CompliancePack,omitempty" xml:"CompliancePack,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBody) GetCompliancePack() *GetAggregateCompliancePackResponseBodyCompliancePack {
	return s.CompliancePack
}

func (s *GetAggregateCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateCompliancePackResponseBody) SetCompliancePack(v *GetAggregateCompliancePackResponseBodyCompliancePack) *GetAggregateCompliancePackResponseBody {
	s.CompliancePack = v
	return s
}

func (s *GetAggregateCompliancePackResponseBody) SetRequestId(v string) *GetAggregateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBody) Validate() error {
	if s.CompliancePack != nil {
		if err := s.CompliancePack.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateCompliancePackResponseBodyCompliancePack struct {
	// The ID of the management account to which the compliance pack belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the account group.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance pack.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance pack.
	//
	// example:
	//
	// 等保三级预检合规包
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the compliance pack template.
	//
	// example:
	//
	// ct-5f26ff4e06a300c4****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The list of rules in the compliance pack.
	ConfigRules []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	// The timestamp when the compliance pack was created. Unit: milliseconds.
	//
	// example:
	//
	// 1624243657000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the compliance pack.
	//
	// example:
	//
	// 基于等保2.0三级标准，提供持续检测合规性的建议模板，帮助您提前自检并修复问题，以便快速通过正式检测。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The risk level of the compliance pack. Valid values:
	//
	// - 1: high
	//
	// - 2: medium
	//
	// - 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The evaluation scope of the compliance pack.
	Scope *GetAggregateCompliancePackResponseBodyCompliancePackScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// The status of the compliance pack. Valid values:
	//
	// - ACTIVE: The compliance pack is active.
	//
	// - CREATING: The compliance pack is being created.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resource.
	Tags []*GetAggregateCompliancePackResponseBodyCompliancePackTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The information about the compliance pack template. The list of rules in the template does not include user-defined function rules. You can use the template to quickly create the same compliance pack for other accounts or account groups.
	//
	// example:
	//
	// {
	//
	//     "configRuleTemplates": [
	//
	//         {
	//
	//             "configRuleName": "自定义条件规则示例",
	//
	//             "scope": {
	//
	//                 "complianceResourceTypes": [
	//
	//                     "ACS::ECS::Instance"
	//
	//                 ]
	//
	//             },
	//
	//             "description": "",
	//
	//             "source": {
	//
	//                 "owner": "CUSTOM_CONFIGURATION",
	//
	//                 "identifier": "acs-config-configuration",
	//
	//                 "sourceDetails": [
	//
	//                     {
	//
	//                         "messageType": "ScheduledNotification",
	//
	//                         "maximumExecutionFrequency": "Twelve_Hours"
	//
	//                     },
	//
	//                     {
	//
	//                         "messageType": "ConfigurationItemChangeNotification"
	//
	//                     }
	//
	//                 ],
	//
	//                 "conditions": "{\\"ComplianceConditions\\":\\"{\\\\\\"operator\\\\\\":\\\\\\"and\\\\\\",\\\\\\"children\\\\\\":[{\\\\\\"operator\\\\\\":\\\\\\"GreaterOrEquals\\\\\\",\\\\\\"featurePath\\\\\\":\\\\\\"$.Cpu\\\\\\",\\\\\\"featureSource\\\\\\":\\\\\\"CONFIGURATION\\\\\\",\\\\\\"desired\\\\\\":\\\\\\"2\\\\\\"}]}\\"}"
	//
	//             },
	//
	//             "inputParameters": {}
	//
	//         },
	//
	//         {
	//
	//             "configRuleName": "OSS存储空间Referer在指定的防盗链白名单中",
	//
	//             "scope": {
	//
	//                 "complianceResourceTypes": [
	//
	//                     "ACS::OSS::Bucket"
	//
	//                 ]
	//
	//             },
	//
	//             "description": "OSS存储空间开启防盗链并且Referer在指定白名单中，视为“合规”。",
	//
	//             "source": {
	//
	//                 "owner": "ALIYUN",
	//
	//                 "identifier": "oss-bucket-referer-limit",
	//
	//                 "sourceDetails": [
	//
	//                     {
	//
	//                         "messageType": "ConfigurationItemChangeNotification"
	//
	//                     }
	//
	//                 ]
	//
	//             },
	//
	//             "inputParameters": {
	//
	//                 "allowEmptyReferer": "true",
	//
	//                 "allowReferers": "http://www.aliyun.com"
	//
	//             }
	//
	//         }
	//
	//     ]
	//
	// }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
}

func (s GetAggregateCompliancePackResponseBodyCompliancePack) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePack) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetConfigRules() []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	return s.ConfigRules
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetDescription() *string {
	return s.Description
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetScope() *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	return s.Scope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetStatus() *string {
	return s.Status
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetTags() []*GetAggregateCompliancePackResponseBodyCompliancePackTags {
	return s.Tags
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetAccountId(v int64) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.AccountId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetAggregatorId(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetCompliancePackId(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetCompliancePackName(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.CompliancePackName = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetCompliancePackTemplateId(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetConfigRules(v []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.ConfigRules = v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetCreateTimestamp(v int64) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.CreateTimestamp = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetDescription(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.Description = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetRiskLevel(v int32) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.RiskLevel = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetScope(v *GetAggregateCompliancePackResponseBodyCompliancePackScope) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.Scope = v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetStatus(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.Status = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetTags(v []*GetAggregateCompliancePackResponseBodyCompliancePackTags) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.Tags = v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetTemplateContent(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.TemplateContent = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) Validate() error {
	if s.ConfigRules != nil {
		for _, item := range s.ConfigRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
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

type GetAggregateCompliancePackResponseBodyCompliancePackConfigRules struct {
	// The ID of the rule.
	//
	// example:
	//
	// cr-a260626622af0005****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// 弹性IP实例带宽满足最低要求
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The input parameters of the rule.
	ConfigRuleParameters []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// 弹性IP实例可用带宽大于等于指定参数值，视为“合规”。默认值：10 MB。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the rule.
	//
	// example:
	//
	// eip-bandwidth-limit
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The types of the resources that are evaluated by the rule. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::EIP::EipAddress
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - 1: high
	//
	// - 2: medium
	//
	// - 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GetConfigRuleParameters() []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	return s.ConfigRuleParameters
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GetDescription() *string {
	return s.Description
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleId(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleName(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleParameters(v []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetDescription(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.Description = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetManagedRuleIdentifier(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetResourceTypesScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ResourceTypesScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetRiskLevel(v int32) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.RiskLevel = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) Validate() error {
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

type GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters struct {
	// The name of the input parameter.
	//
	// example:
	//
	// bandwidth
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the input parameter.
	//
	// example:
	//
	// 10
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	// Indicates whether the input parameter is required. Valid values:
	//
	// - true: The input parameter is required.
	//
	// - false: The input parameter is not required.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GetRequired() *bool {
	return s.Required
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetParameterName(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetParameterValue(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetRequired(v bool) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.Required = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) Validate() error {
	return dara.Validate(s)
}

type GetAggregateCompliancePackResponseBodyCompliancePackScope struct {
	// The IDs of the regions to exclude. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The IDs of the resource groups to exclude. The compliance pack does not apply to resources in these groups. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The IDs of the resources to exclude. The compliance pack does not apply to these resources. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// eip-8vbf3x310fn56ijfd****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The excluded tag scope.
	ExcludeTagsScope []*GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The IDs of the regions where the compliance pack applies. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The IDs of the resource groups where the compliance pack applies. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The IDs of the resources to which the rule applies. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// eip-8vbf3x310fn56ijfd****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The tag key of the resources to which the compliance pack applies.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// The tag value of the resources to which the compliance pack applies.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tag scope.
	TagsScope []*GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackScope) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackScope) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetExcludeTagsScope() []*GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) GetTagsScope() []*GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope {
	return s.TagsScope
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetExcludeRegionIdsScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetExcludeResourceGroupIdsScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetExcludeResourceIdsScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetExcludeTagsScope(v []*GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.ExcludeTagsScope = v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetRegionIdsScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.RegionIdsScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetResourceGroupIdsScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetResourceIdsScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.ResourceIdsScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetTagKeyScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.TagKeyScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetTagValueScope(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.TagValueScope = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) SetTagsScope(v []*GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope) *GetAggregateCompliancePackResponseBodyCompliancePackScope {
	s.TagsScope = v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScope) Validate() error {
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

type GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope struct {
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

func (s GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) SetTagKey(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) SetTagValue(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope struct {
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

func (s GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope) SetTagKey(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope {
	s.TagKey = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope) SetTagValue(v string) *GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope {
	s.TagValue = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackScopeTagsScope) Validate() error {
	return dara.Validate(s)
}

type GetAggregateCompliancePackResponseBodyCompliancePackTags struct {
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

func (s GetAggregateCompliancePackResponseBodyCompliancePackTags) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackTags) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackTags) SetTagKey(v string) *GetAggregateCompliancePackResponseBodyCompliancePackTags {
	s.TagKey = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackTags) SetTagValue(v string) *GetAggregateCompliancePackResponseBodyCompliancePackTags {
	s.TagValue = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackTags) Validate() error {
	return dara.Validate(s)
}
