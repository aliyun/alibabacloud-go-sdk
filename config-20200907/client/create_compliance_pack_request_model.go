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
	// test-pack-name
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the compliance package template.
	//
	// You can call the [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html) operation to obtain the ID of the compliance package.
	//
	// example:
	//
	// ct-5f26ff4e06a300c4****
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	// The rules in the compliance package. You must specify either this parameter or TemplateContent.
	//
	// if can be null:
	// false
	ConfigRules []*CreateCompliancePackRequestConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
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
	// Test pack description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ExcludeRegionIdsScope
	//
	// example:
	//
	// cn-hangzhou
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
	// eip-8vbf3x310fn56ijfd****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// ExcludeTagsScope
	ExcludeTagsScope []*CreateCompliancePackRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	// ResourceIdsScope
	//
	// example:
	//
	// eip-8vbf3x310fn56ijfd****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The risk level of the resources that are not compliant with the rules in the compliance package. Default value: 2. Valid values:
	//
	// 	- 1: high.
	//
	// 	- 2: medium.
	//
	// 	- 3: low.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	Tag []*CreateCompliancePackRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// TagsScope
	TagsScope []*CreateCompliancePackRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
	// The information about the template that is used to generate the compliance package. You can call an API operation to view the details of an existing compliance package or write a compliance package template. For more information, see [Write a compliance package template in a configuration file](https://help.aliyun.com/document_detail/2659733.html). You must specify one of ConfigRules and TemplateContent.
	//
	// example:
	//
	// { "configRuleTemplates": [ { "configRuleName": "condition-rule-example", "scope": { "complianceResourceTypes": [ "ACS::ECS::Instance" ] }, "description": "", "source": { "owner": "CUSTOM_CONFIGURATION", "identifier": "acs-config-configuration", "sourceDetails": [ { "messageType": "ScheduledNotification", "maximumExecutionFrequency": "Twelve_Hours" }, { "messageType": "ConfigurationItemChangeNotification" } ], "conditions": "{\\\\"ComplianceConditions\\\\":\\\\"{\\\\\\\\\\"operator\\\\\\\\\\":\\\\\\\\\\"and\\\\\\\\\\",\\\\\\\\\\"children\\\\\\\\\\":[{\\\\\\\\\\"operator\\\\\\\\\\":\\\\\\\\\\"GreaterOrEquals\\\\\\\\\\",\\\\\\\\\\"featurePath\\\\\\\\\\":\\\\\\\\\\"$.Cpu\\\\\\\\\\",\\\\\\\\\\"featureSource\\\\\\\\\\":\\\\\\\\\\"CONFIGURATION\\\\\\\\\\",\\\\\\\\\\"desired\\\\\\\\\\":\\\\\\\\\\"2\\\\\\\\\\"}]}\\\\"}" }, "inputParameters": {} }, { "configRuleName": "oss-bucket-referer-limit", "scope": { "complianceResourceTypes": [ "ACS::OSS::Bucket" ] }, "description": "If the hotlink protection feature is enabled for the Object Storage Service (OSS) bucket and the Referer is added to a specific whitelist, the evaluation result is compliant.", "source": { "owner": "ALIYUN", "identifier": "oss-bucket-referer-limit", "sourceDetails": [ { "messageType": "ConfigurationItemChangeNotification" } ] }, "inputParameters": { "allowEmptyReferer": "true", "allowReferers": "http://www.aliyun.com" } } ] }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
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
	// The rule ID. If you specify this parameter, Cloud Config adds the rule that has the specified ID to the compliance package.
	//
	// You need to only specify `ManagedRuleIdentifier` or `ConfigRuleId`. If you specify both parameters, Cloud Config adds a rule based on the value of `ConfigRuleId`. You can call the [ListConfigRules](https://help.aliyun.com/document_detail/169607.html) operation to obtain the rule ID.
	//
	// example:
	//
	// cr-e918626622af000f****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// eip-bandwidth-limit
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The input parameters of the rule.
	ConfigRuleParameters []*CreateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	// The description of the rule.
	//
	// example:
	//
	// The description of the test rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the managed rule. Cloud Config automatically creates a managed rule based on the specified identifier and adds the rule to the compliance package.
	//
	// You need to only specify `ManagedRuleIdentifier` or `ConfigRuleId`. If you specify both parameters, Cloud Config adds a rule based on the value of `ConfigRuleId`. You can call the [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html) operation to obtain the identifier of the managed rule.
	//
	// example:
	//
	// eip-bandwidth-limit
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The risk level of the resources that do not comply with the rule. Valid values:
	//
	// 	- 1: high.
	//
	// 	- 2: medium.
	//
	// 	- 3: low.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
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
	// The name of the input parameter.
	//
	// You must specify both `ParameterName` and `ParameterValue` or neither of them. If the managed rule has an input parameter but no default value is specified, you must specify this parameter. You can call the [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html) operation to obtain the names of input parameters of the managed rule.
	//
	// example:
	//
	// bandwidth
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the input parameter.
	//
	// You must specify both `ParameterName` and `ParameterValue` or neither of them. If the managed rule has an input parameter but no default value is specified, you must specify this parameter. You can call the [ListCompliancePackTemplates](https://help.aliyun.com/document_detail/261176.html) operation to obtain the values of input parameters of the managed rule.
	//
	// example:
	//
	// 10
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
	// TagKey
	//
	// example:
	//
	// key-2
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// TagValue
	//
	// example:
	//
	// value-2
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
	// The tag keys.
	//
	// The tag keys cannot be an empty string. The tag keys can be up to 64 characters in length. The tag keys cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be an empty string or up to 128 characters in length. The tag values cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// value-1
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
	// Tagkey
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// TagValue
	//
	// example:
	//
	// value-1
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
