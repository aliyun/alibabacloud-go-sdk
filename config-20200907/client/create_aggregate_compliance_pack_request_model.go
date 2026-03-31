// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateCompliancePackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *CreateAggregateCompliancePackRequest
	GetAggregatorId() *string
	SetClientToken(v string) *CreateAggregateCompliancePackRequest
	GetClientToken() *string
	SetCompliancePackName(v string) *CreateAggregateCompliancePackRequest
	GetCompliancePackName() *string
	SetCompliancePackTemplateId(v string) *CreateAggregateCompliancePackRequest
	GetCompliancePackTemplateId() *string
	SetConfigRules(v []*CreateAggregateCompliancePackRequestConfigRules) *CreateAggregateCompliancePackRequest
	GetConfigRules() []*CreateAggregateCompliancePackRequestConfigRules
	SetDefaultEnable(v bool) *CreateAggregateCompliancePackRequest
	GetDefaultEnable() *bool
	SetDescription(v string) *CreateAggregateCompliancePackRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *CreateAggregateCompliancePackRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateAggregateCompliancePackRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateAggregateCompliancePackRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*CreateAggregateCompliancePackRequestExcludeTagsScope) *CreateAggregateCompliancePackRequest
	GetExcludeTagsScope() []*CreateAggregateCompliancePackRequestExcludeTagsScope
	SetRegionIdsScope(v string) *CreateAggregateCompliancePackRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateAggregateCompliancePackRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateAggregateCompliancePackRequest
	GetResourceIdsScope() *string
	SetRiskLevel(v int32) *CreateAggregateCompliancePackRequest
	GetRiskLevel() *int32
	SetTag(v []*CreateAggregateCompliancePackRequestTag) *CreateAggregateCompliancePackRequest
	GetTag() []*CreateAggregateCompliancePackRequestTag
	SetTagKeyScope(v string) *CreateAggregateCompliancePackRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *CreateAggregateCompliancePackRequest
	GetTagValueScope() *string
	SetTagsScope(v []*CreateAggregateCompliancePackRequestTagsScope) *CreateAggregateCompliancePackRequest
	GetTagsScope() []*CreateAggregateCompliancePackRequestTagsScope
	SetTemplateContent(v string) *CreateAggregateCompliancePackRequest
	GetTemplateContent() *string
}

type CreateAggregateCompliancePackRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The client token that you want to use to ensure the idempotency of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.``
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the compliance package.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-compliance-pack-name
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the compliance package template from which you want to create a compliance package.
	//
	// For more information about how to obtain the ID of a compliance package template, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// ct-5f26ff4e06a300c4****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The rules in the compliance package.
	//
	// >  You must configure this parameter or the `TemplateContent` parameter.
	ConfigRules []*CreateAggregateCompliancePackRequestConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	// Specifies whether to enable the rule together with the compliance package. Valid values:
	//
	// 	- true: The system enables the rule together with the compliance package.
	//
	// 	- false: The system does not enable the rule together with the compliance package.
	//
	// example:
	//
	// false
	DefaultEnable *bool `json:"DefaultEnable,omitempty" xml:"DefaultEnable,omitempty"`
	// The description of the compliance package.
	//
	// example:
	//
	// Test compliance pack descripaiton.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the regions excluded from the compliance evaluations performed by the compliance package. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The IDs of the resource groups excluded from the compliance evaluations performed by the rule. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The ID of the resource that you do not want to evaluate by using the compliance package. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// eip-8vbf3x310fn56ijfd****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The tags that are excluded.
	ExcludeTagsScope []*CreateAggregateCompliancePackRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The IDs of the resources to which the rule applies. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The risk level of the resources that are not compliant with the rules in the compliance package. Valid values:
	//
	// 	- 1: high
	//
	// 	- 2 (default): medium
	//
	// 	- 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	Tag []*CreateAggregateCompliancePackRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The tag key of the resource that you want to evaluate by using the compliance package.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// The tag value of the resource that you want to evaluate by using the compliance package.
	//
	// >  You must configure the TagValueScope parameter together with the TagKeyScope parameter.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tag scope.
	TagsScope []*CreateAggregateCompliancePackRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
	// The information about the template that is used to create the compliance package. You can call the GetAggregateCompliancePack operation to view the details of an existing compliance package. You can also write a compliance package template. For more information, see [Write a compliance package template in a configuration file](https://help.aliyun.com/document_detail/2659733.html).
	//
	// >  You must configure this parameter or the `TemplateContent` parameter.
	//
	// example:
	//
	// { "configRuleTemplates": [ { "configRuleName": "condition-rule-example", "scope": { "complianceResourceTypes": [ "ACS::ECS::Instance" ] }, "description": "", "source": { "owner": "CUSTOM_CONFIGURATION", "identifier": "acs-config-configuration", "sourceDetails": [ { "messageType": "ScheduledNotification", "maximumExecutionFrequency": "Twelve_Hours" }, { "messageType": "ConfigurationItemChangeNotification" } ], "conditions": "{\\\\"ComplianceConditions\\\\":\\\\"{\\\\\\\\\\"operator\\\\\\\\\\":\\\\\\\\\\"and\\\\\\\\\\",\\\\\\\\\\"children\\\\\\\\\\":[{\\\\\\\\\\"operator\\\\\\\\\\":\\\\\\\\\\"GreaterOrEquals\\\\\\\\\\",\\\\\\\\\\"featurePath\\\\\\\\\\":\\\\\\\\\\"$.Cpu\\\\\\\\\\",\\\\\\\\\\"featureSource\\\\\\\\\\":\\\\\\\\\\"CONFIGURATION\\\\\\\\\\",\\\\\\\\\\"desired\\\\\\\\\\":\\\\\\\\\\"2\\\\\\\\\\"}]}\\\\"}" }, "inputParameters": {} }, { "configRuleName": "oss-bucket-referer-limit", "scope": { "complianceResourceTypes": [ "ACS::OSS::Bucket" ] }, "description": "If the hotlink protection feature is enabled for the Object Storage Service (OSS) bucket and the Referer is added to a specific whitelist, the evaluation result is compliant.", "source": { "owner": "ALIYUN", "identifier": "oss-bucket-referer-limit", "sourceDetails": [ { "messageType": "ConfigurationItemChangeNotification" } ] }, "inputParameters": { "allowEmptyReferer": "true", "allowReferers": "http://www.aliyun.com" } } ] }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
}

func (s CreateAggregateCompliancePackRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *CreateAggregateCompliancePackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAggregateCompliancePackRequest) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *CreateAggregateCompliancePackRequest) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *CreateAggregateCompliancePackRequest) GetConfigRules() []*CreateAggregateCompliancePackRequestConfigRules {
	return s.ConfigRules
}

func (s *CreateAggregateCompliancePackRequest) GetDefaultEnable() *bool {
	return s.DefaultEnable
}

func (s *CreateAggregateCompliancePackRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAggregateCompliancePackRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateAggregateCompliancePackRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateAggregateCompliancePackRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateAggregateCompliancePackRequest) GetExcludeTagsScope() []*CreateAggregateCompliancePackRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *CreateAggregateCompliancePackRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateAggregateCompliancePackRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateAggregateCompliancePackRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateAggregateCompliancePackRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateAggregateCompliancePackRequest) GetTag() []*CreateAggregateCompliancePackRequestTag {
	return s.Tag
}

func (s *CreateAggregateCompliancePackRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *CreateAggregateCompliancePackRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *CreateAggregateCompliancePackRequest) GetTagsScope() []*CreateAggregateCompliancePackRequestTagsScope {
	return s.TagsScope
}

func (s *CreateAggregateCompliancePackRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *CreateAggregateCompliancePackRequest) SetAggregatorId(v string) *CreateAggregateCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetClientToken(v string) *CreateAggregateCompliancePackRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetCompliancePackName(v string) *CreateAggregateCompliancePackRequest {
	s.CompliancePackName = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetCompliancePackTemplateId(v string) *CreateAggregateCompliancePackRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetConfigRules(v []*CreateAggregateCompliancePackRequestConfigRules) *CreateAggregateCompliancePackRequest {
	s.ConfigRules = v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetDefaultEnable(v bool) *CreateAggregateCompliancePackRequest {
	s.DefaultEnable = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetDescription(v string) *CreateAggregateCompliancePackRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetExcludeRegionIdsScope(v string) *CreateAggregateCompliancePackRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetExcludeResourceGroupIdsScope(v string) *CreateAggregateCompliancePackRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetExcludeResourceIdsScope(v string) *CreateAggregateCompliancePackRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetExcludeTagsScope(v []*CreateAggregateCompliancePackRequestExcludeTagsScope) *CreateAggregateCompliancePackRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetRegionIdsScope(v string) *CreateAggregateCompliancePackRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetResourceGroupIdsScope(v string) *CreateAggregateCompliancePackRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetResourceIdsScope(v string) *CreateAggregateCompliancePackRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetRiskLevel(v int32) *CreateAggregateCompliancePackRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetTag(v []*CreateAggregateCompliancePackRequestTag) *CreateAggregateCompliancePackRequest {
	s.Tag = v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetTagKeyScope(v string) *CreateAggregateCompliancePackRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetTagValueScope(v string) *CreateAggregateCompliancePackRequest {
	s.TagValueScope = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetTagsScope(v []*CreateAggregateCompliancePackRequestTagsScope) *CreateAggregateCompliancePackRequest {
	s.TagsScope = v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetTemplateContent(v string) *CreateAggregateCompliancePackRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) Validate() error {
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

type CreateAggregateCompliancePackRequestConfigRules struct {
	// The rule ID. If you configure this parameter, Cloud Config adds the rule that has the specified ID to the compliance package.
	//
	// You need to only configure the `ManagedRuleIdentifier` or `ConfigRuleId` parameter. If you configure both parameters, the value of the `ConfigRuleId` parameter takes precedence. For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// cr-e918626622af000f****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// eip-bandwidth-limit
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The input parameters of the rule.
	ConfigRuleParameters []*CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	// The rule description.
	//
	// example:
	//
	// Test rule description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the managed rule. Cloud Config automatically creates a rule based on the identifier of the managed rule and adds the rule to the current compliance package.
	//
	// You need to only configure the `ManagedRuleIdentifier` or `ConfigRuleId` parameter. If you configure both parameters, the value of the `ConfigRuleId` parameter takes precedence. For more information about how to obtain the identifier of a managed rule, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// eip-bandwidth-limit
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
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

func (s CreateAggregateCompliancePackRequestConfigRules) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackRequestConfigRules) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequestConfigRules) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *CreateAggregateCompliancePackRequestConfigRules) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *CreateAggregateCompliancePackRequestConfigRules) GetConfigRuleParameters() []*CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	return s.ConfigRuleParameters
}

func (s *CreateAggregateCompliancePackRequestConfigRules) GetDescription() *string {
	return s.Description
}

func (s *CreateAggregateCompliancePackRequestConfigRules) GetManagedRuleIdentifier() *string {
	return s.ManagedRuleIdentifier
}

func (s *CreateAggregateCompliancePackRequestConfigRules) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetConfigRuleId(v string) *CreateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetConfigRuleName(v string) *CreateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetConfigRuleParameters(v []*CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) *CreateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetDescription(v string) *CreateAggregateCompliancePackRequestConfigRules {
	s.Description = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetManagedRuleIdentifier(v string) *CreateAggregateCompliancePackRequestConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetRiskLevel(v int32) *CreateAggregateCompliancePackRequestConfigRules {
	s.RiskLevel = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) Validate() error {
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

type CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters struct {
	// The name of the input parameter.
	//
	// You must configure the `ParameterName` and `ParameterValue` parameters or neither of them. If the managed rule has an input parameter but no default value exists, you must configure this parameter. For more information about how to obtain the name of an input parameter for a managed rule, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// bandwidth
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the input parameter.
	//
	// You must configure the `ParameterName` and `ParameterValue` parameters or neither of them. If the managed rule has an input parameter but no default value exists, you must configure this parameter. For more information about how to obtain the value of an input parameter for a managed rule, see [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html).
	//
	// example:
	//
	// 10
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterName(v string) *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterValue(v string) *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) Validate() error {
	return dara.Validate(s)
}

type CreateAggregateCompliancePackRequestExcludeTagsScope struct {
	// The tag key.
	//
	// example:
	//
	// 4
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// user
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateAggregateCompliancePackRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAggregateCompliancePackRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAggregateCompliancePackRequestExcludeTagsScope) SetTagKey(v string) *CreateAggregateCompliancePackRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestExcludeTagsScope) SetTagValue(v string) *CreateAggregateCompliancePackRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type CreateAggregateCompliancePackRequestTag struct {
	// The key of the tag that is added to the resource.
	//
	// You can add up to 20 tag keys to a resource.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is added to the resource.
	//
	// You can add up to 20 tag values to a resource.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAggregateCompliancePackRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAggregateCompliancePackRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAggregateCompliancePackRequestTag) SetKey(v string) *CreateAggregateCompliancePackRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestTag) SetValue(v string) *CreateAggregateCompliancePackRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateAggregateCompliancePackRequestTagsScope struct {
	// The tag key.
	//
	// example:
	//
	// tagKey1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tagValue1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateAggregateCompliancePackRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackRequestTagsScope) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAggregateCompliancePackRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAggregateCompliancePackRequestTagsScope) SetTagKey(v string) *CreateAggregateCompliancePackRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestTagsScope) SetTagValue(v string) *CreateAggregateCompliancePackRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
