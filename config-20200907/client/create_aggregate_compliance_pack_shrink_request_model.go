// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateCompliancePackShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *CreateAggregateCompliancePackShrinkRequest
	GetAggregatorId() *string
	SetClientToken(v string) *CreateAggregateCompliancePackShrinkRequest
	GetClientToken() *string
	SetCompliancePackName(v string) *CreateAggregateCompliancePackShrinkRequest
	GetCompliancePackName() *string
	SetCompliancePackTemplateId(v string) *CreateAggregateCompliancePackShrinkRequest
	GetCompliancePackTemplateId() *string
	SetConfigRulesShrink(v string) *CreateAggregateCompliancePackShrinkRequest
	GetConfigRulesShrink() *string
	SetDefaultEnable(v bool) *CreateAggregateCompliancePackShrinkRequest
	GetDefaultEnable() *bool
	SetDescription(v string) *CreateAggregateCompliancePackShrinkRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) *CreateAggregateCompliancePackShrinkRequest
	GetExcludeTagsScope() []*CreateAggregateCompliancePackShrinkRequestExcludeTagsScope
	SetRegionIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest
	GetResourceIdsScope() *string
	SetRiskLevel(v int32) *CreateAggregateCompliancePackShrinkRequest
	GetRiskLevel() *int32
	SetTagShrink(v string) *CreateAggregateCompliancePackShrinkRequest
	GetTagShrink() *string
	SetTagKeyScope(v string) *CreateAggregateCompliancePackShrinkRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *CreateAggregateCompliancePackShrinkRequest
	GetTagValueScope() *string
	SetTagsScope(v []*CreateAggregateCompliancePackShrinkRequestTagsScope) *CreateAggregateCompliancePackShrinkRequest
	GetTagsScope() []*CreateAggregateCompliancePackShrinkRequestTagsScope
	SetTemplateContent(v string) *CreateAggregateCompliancePackShrinkRequest
	GetTemplateContent() *string
}

type CreateAggregateCompliancePackShrinkRequest struct {
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
	ConfigRulesShrink *string `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
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
	ExcludeTagsScope []*CreateAggregateCompliancePackShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
	TagsScope []*CreateAggregateCompliancePackShrinkRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
	// The information about the template that is used to create the compliance package. You can call the GetAggregateCompliancePack operation to view the details of an existing compliance package. You can also write a compliance package template. For more information, see [Write a compliance package template in a configuration file](https://help.aliyun.com/document_detail/2659733.html).
	//
	// >  You must configure this parameter or the `TemplateContent` parameter.
	//
	// example:
	//
	// { "configRuleTemplates": [ { "configRuleName": "condition-rule-example", "scope": { "complianceResourceTypes": [ "ACS::ECS::Instance" ] }, "description": "", "source": { "owner": "CUSTOM_CONFIGURATION", "identifier": "acs-config-configuration", "sourceDetails": [ { "messageType": "ScheduledNotification", "maximumExecutionFrequency": "Twelve_Hours" }, { "messageType": "ConfigurationItemChangeNotification" } ], "conditions": "{\\\\"ComplianceConditions\\\\":\\\\"{\\\\\\\\\\"operator\\\\\\\\\\":\\\\\\\\\\"and\\\\\\\\\\",\\\\\\\\\\"children\\\\\\\\\\":[{\\\\\\\\\\"operator\\\\\\\\\\":\\\\\\\\\\"GreaterOrEquals\\\\\\\\\\",\\\\\\\\\\"featurePath\\\\\\\\\\":\\\\\\\\\\"$.Cpu\\\\\\\\\\",\\\\\\\\\\"featureSource\\\\\\\\\\":\\\\\\\\\\"CONFIGURATION\\\\\\\\\\",\\\\\\\\\\"desired\\\\\\\\\\":\\\\\\\\\\"2\\\\\\\\\\"}]}\\\\"}" }, "inputParameters": {} }, { "configRuleName": "oss-bucket-referer-limit", "scope": { "complianceResourceTypes": [ "ACS::OSS::Bucket" ] }, "description": "If the hotlink protection feature is enabled for the Object Storage Service (OSS) bucket and the Referer is added to a specific whitelist, the evaluation result is compliant.", "source": { "owner": "ALIYUN", "identifier": "oss-bucket-referer-limit", "sourceDetails": [ { "messageType": "ConfigurationItemChangeNotification" } ] }, "inputParameters": { "allowEmptyReferer": "true", "allowReferers": "http://www.aliyun.com" } } ] }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
}

func (s CreateAggregateCompliancePackShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetCompliancePackTemplateId() *string {
	return s.CompliancePackTemplateId
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetConfigRulesShrink() *string {
	return s.ConfigRulesShrink
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetDefaultEnable() *bool {
	return s.DefaultEnable
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetExcludeTagsScope() []*CreateAggregateCompliancePackShrinkRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetTagsScope() []*CreateAggregateCompliancePackShrinkRequestTagsScope {
	return s.TagsScope
}

func (s *CreateAggregateCompliancePackShrinkRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetAggregatorId(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetClientToken(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetCompliancePackName(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.CompliancePackName = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetCompliancePackTemplateId(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetConfigRulesShrink(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ConfigRulesShrink = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetDefaultEnable(v bool) *CreateAggregateCompliancePackShrinkRequest {
	s.DefaultEnable = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetDescription(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetExcludeRegionIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetExcludeResourceGroupIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetExcludeResourceIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetExcludeTagsScope(v []*CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) *CreateAggregateCompliancePackShrinkRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetRegionIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetResourceGroupIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetResourceIdsScope(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetRiskLevel(v int32) *CreateAggregateCompliancePackShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetTagShrink(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetTagKeyScope(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetTagValueScope(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.TagValueScope = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetTagsScope(v []*CreateAggregateCompliancePackShrinkRequestTagsScope) *CreateAggregateCompliancePackShrinkRequest {
	s.TagsScope = v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetTemplateContent(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) Validate() error {
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

type CreateAggregateCompliancePackShrinkRequestExcludeTagsScope struct {
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

func (s CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) SetTagKey(v string) *CreateAggregateCompliancePackShrinkRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) SetTagValue(v string) *CreateAggregateCompliancePackShrinkRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type CreateAggregateCompliancePackShrinkRequestTagsScope struct {
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

func (s CreateAggregateCompliancePackShrinkRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateCompliancePackShrinkRequestTagsScope) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackShrinkRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAggregateCompliancePackShrinkRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAggregateCompliancePackShrinkRequestTagsScope) SetTagKey(v string) *CreateAggregateCompliancePackShrinkRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequestTagsScope) SetTagValue(v string) *CreateAggregateCompliancePackShrinkRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
