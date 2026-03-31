// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompliancePack(v *GetCompliancePackResponseBodyCompliancePack) *GetCompliancePackResponseBody
	GetCompliancePack() *GetCompliancePackResponseBodyCompliancePack
	SetRequestId(v string) *GetCompliancePackResponseBody
	GetRequestId() *string
}

type GetCompliancePackResponseBody struct {
	// The details of the compliance package.
	CompliancePack *GetCompliancePackResponseBodyCompliancePack `json:"CompliancePack,omitempty" xml:"CompliancePack,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EC7AED1-172F-42AE-9C12-295BC2ADB751
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCompliancePackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBody) GetCompliancePack() *GetCompliancePackResponseBodyCompliancePack {
	return s.CompliancePack
}

func (s *GetCompliancePackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCompliancePackResponseBody) SetCompliancePack(v *GetCompliancePackResponseBodyCompliancePack) *GetCompliancePackResponseBody {
	s.CompliancePack = v
	return s
}

func (s *GetCompliancePackResponseBody) SetRequestId(v string) *GetCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCompliancePackResponseBody) Validate() error {
	if s.CompliancePack != nil {
		if err := s.CompliancePack.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCompliancePackResponseBodyCompliancePack struct {
	// The ID of the Alibaba Cloud account to which the compliance package belongs.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-a8a8626622af0082****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance package.
	//
	// example:
	//
	// example-pack-name
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the compliance package template.
	//
	// example:
	//
	// ct-5f26ff4e06a300c4****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The rules in the compliance package.
	ConfigRules []*GetCompliancePackResponseBodyCompliancePackConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	// The timestamp when the compliance package was created. Unit: milliseconds.
	//
	// example:
	//
	// 1624245766000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the compliance package.
	//
	// example:
	//
	// example-name
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The risk level of the resources that are not compliant with the rules in the compliance package. Valid values:
	//
	// 	- 1: high
	//
	// 	- 2: medium
	//
	// 	- 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The resource group for which the compliance package took effect.
	Scope *GetCompliancePackResponseBodyCompliancePackScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// The status of the compliance package. Valid values:
	//
	// 	- ACTIVE: The compliance package is normal.
	//
	// 	- CREATING: The compliance package is being created.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*GetCompliancePackResponseBodyCompliancePackTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The information about the current compliance package template. The rules in the template do not contain custom function rules. You can quickly create the same compliance package for other accounts or account groups based on the template information.
	//
	// example:
	//
	// {
	//
	//     "configRuleTemplates": [
	//
	//         {
	//
	//             "configRuleName": "rule-example",
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
	//             "configRuleName": "name",
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
	//             "description": "description-1",
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

func (s GetCompliancePackResponseBodyCompliancePack) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePack) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetConfigRules() []*GetCompliancePackResponseBodyCompliancePackConfigRules {
	return s.ConfigRules
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetDescription() *string {
	return s.Description
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetScope() *GetCompliancePackResponseBodyCompliancePackScope {
	return s.Scope
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetStatus() *string {
	return s.Status
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetTags() []*GetCompliancePackResponseBodyCompliancePackTags {
	return s.Tags
}

func (s *GetCompliancePackResponseBodyCompliancePack) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetAccountId(v int64) *GetCompliancePackResponseBodyCompliancePack {
	s.AccountId = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetCompliancePackId(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.CompliancePackId = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetCompliancePackName(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.CompliancePackName = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetCompliancePackTemplateId(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetConfigRules(v []*GetCompliancePackResponseBodyCompliancePackConfigRules) *GetCompliancePackResponseBodyCompliancePack {
	s.ConfigRules = v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetCreateTimestamp(v int64) *GetCompliancePackResponseBodyCompliancePack {
	s.CreateTimestamp = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetDescription(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.Description = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetRiskLevel(v int32) *GetCompliancePackResponseBodyCompliancePack {
	s.RiskLevel = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetScope(v *GetCompliancePackResponseBodyCompliancePackScope) *GetCompliancePackResponseBodyCompliancePack {
	s.Scope = v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetStatus(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.Status = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetTags(v []*GetCompliancePackResponseBodyCompliancePackTags) *GetCompliancePackResponseBodyCompliancePack {
	s.Tags = v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetTemplateContent(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.TemplateContent = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) Validate() error {
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

type GetCompliancePackResponseBodyCompliancePackConfigRules struct {
	// The rule ID.
	//
	// example:
	//
	// cr-a260626622af0005****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// eip-bandwidth-limit
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The input parameters of the rule.
	ConfigRuleParameters []*GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	// The rule description.
	//
	// example:
	//
	// example-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the managed rule.
	//
	// example:
	//
	// eip-bandwidth-limit
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The type of the resource evaluated based on the rule. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::EIP::EipAddress
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The risk level of the resources that do not comply with the rule. Valid values:
	//
	// 	- 1: high
	//
	// 	- 2: medium
	//
	// 	- 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetCompliancePackResponseBodyCompliancePackConfigRules) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePackConfigRules) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) GetConfigRuleParameters() []*GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	return s.ConfigRuleParameters
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) GetDescription() *string {
	return s.Description
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleId(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleName(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleParameters(v []*GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetDescription(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.Description = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetManagedRuleIdentifier(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetResourceTypesScope(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ResourceTypesScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetRiskLevel(v int32) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.RiskLevel = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) Validate() error {
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

type GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters struct {
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
	// Indicates whether the input parameters are required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GetRequired() *bool {
	return s.Required
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetParameterName(v string) *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetParameterValue(v string) *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetRequired(v bool) *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.Required = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) Validate() error {
	return dara.Validate(s)
}

type GetCompliancePackResponseBodyCompliancePackScope struct {
	// The IDs of regions that are excluded. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The IDs of the resource groups whose resources you do not want to evaluate by using the compliance package. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The ID of the resource that you do not want to evaluate by using the compliance package.
	//
	// example:
	//
	// eip-8vbf3x310fn56ijfd****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The scope of the tag that is excluded.
	//
	// This parameter is required.
	ExcludeTagsScope []*GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The ID of the region whose resources you want to evaluate by using the compliance package.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The ID of the resource group whose resources you want to evaluate by using the compliance package.
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
	// The tag key of the resource that you want to evaluate by using the compliance package.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// The tag value of the resource that you want to evaluate by using the compliance package.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tag scope.
	//
	// This parameter is required.
	TagsScope []*GetCompliancePackResponseBodyCompliancePackScopeTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s GetCompliancePackResponseBodyCompliancePackScope) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePackScope) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetExcludeTagsScope() []*GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) GetTagsScope() []*GetCompliancePackResponseBodyCompliancePackScopeTagsScope {
	return s.TagsScope
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetExcludeRegionIdsScope(v string) *GetCompliancePackResponseBodyCompliancePackScope {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetExcludeResourceGroupIdsScope(v string) *GetCompliancePackResponseBodyCompliancePackScope {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetExcludeResourceIdsScope(v string) *GetCompliancePackResponseBodyCompliancePackScope {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetExcludeTagsScope(v []*GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) *GetCompliancePackResponseBodyCompliancePackScope {
	s.ExcludeTagsScope = v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetRegionIdsScope(v string) *GetCompliancePackResponseBodyCompliancePackScope {
	s.RegionIdsScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetResourceGroupIdsScope(v string) *GetCompliancePackResponseBodyCompliancePackScope {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetResourceIdsScope(v string) *GetCompliancePackResponseBodyCompliancePackScope {
	s.ResourceIdsScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetTagKeyScope(v string) *GetCompliancePackResponseBodyCompliancePackScope {
	s.TagKeyScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetTagValueScope(v string) *GetCompliancePackResponseBodyCompliancePackScope {
	s.TagValueScope = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) SetTagsScope(v []*GetCompliancePackResponseBodyCompliancePackScopeTagsScope) *GetCompliancePackResponseBodyCompliancePackScope {
	s.TagsScope = v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScope) Validate() error {
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

type GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope struct {
	// The key of the tag.
	//
	// example:
	//
	// key-2
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value-2
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) SetTagKey(v string) *GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) SetTagValue(v string) *GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type GetCompliancePackResponseBodyCompliancePackScopeTagsScope struct {
	// The key of the tag.
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetCompliancePackResponseBodyCompliancePackScopeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePackScopeTagsScope) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeTagsScope) SetTagKey(v string) *GetCompliancePackResponseBodyCompliancePackScopeTagsScope {
	s.TagKey = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeTagsScope) SetTagValue(v string) *GetCompliancePackResponseBodyCompliancePackScopeTagsScope {
	s.TagValue = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackScopeTagsScope) Validate() error {
	return dara.Validate(s)
}

type GetCompliancePackResponseBodyCompliancePackTags struct {
	// The tag keys of the resource.
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag values of the resource.
	//
	// example:
	//
	// value-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetCompliancePackResponseBodyCompliancePackTags) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePackTags) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePackTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetCompliancePackResponseBodyCompliancePackTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetCompliancePackResponseBodyCompliancePackTags) SetTagKey(v string) *GetCompliancePackResponseBodyCompliancePackTags {
	s.TagKey = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackTags) SetTagValue(v string) *GetCompliancePackResponseBodyCompliancePackTags {
	s.TagValue = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackTags) Validate() error {
	return dara.Validate(s)
}
