// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateConfigRuleRequest
	GetClientToken() *string
	SetConfigRuleId(v string) *UpdateConfigRuleRequest
	GetConfigRuleId() *string
	SetConfigRuleName(v string) *UpdateConfigRuleRequest
	GetConfigRuleName() *string
	SetConfigRuleTriggerTypes(v string) *UpdateConfigRuleRequest
	GetConfigRuleTriggerTypes() *string
	SetDescription(v string) *UpdateConfigRuleRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *UpdateConfigRuleRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateConfigRuleRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateConfigRuleRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*UpdateConfigRuleRequestExcludeTagsScope) *UpdateConfigRuleRequest
	GetExcludeTagsScope() []*UpdateConfigRuleRequestExcludeTagsScope
	SetExtendContent(v string) *UpdateConfigRuleRequest
	GetExtendContent() *string
	SetInputParameters(v map[string]interface{}) *UpdateConfigRuleRequest
	GetInputParameters() map[string]interface{}
	SetMaximumExecutionFrequency(v string) *UpdateConfigRuleRequest
	GetMaximumExecutionFrequency() *string
	SetRegionIdsScope(v string) *UpdateConfigRuleRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateConfigRuleRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateConfigRuleRequest
	GetResourceIdsScope() *string
	SetResourceNameScope(v string) *UpdateConfigRuleRequest
	GetResourceNameScope() *string
	SetResourceTypesScope(v []*string) *UpdateConfigRuleRequest
	GetResourceTypesScope() []*string
	SetRiskLevel(v int32) *UpdateConfigRuleRequest
	GetRiskLevel() *int32
	SetTag(v []*UpdateConfigRuleRequestTag) *UpdateConfigRuleRequest
	GetTag() []*UpdateConfigRuleRequestTag
	SetTagKeyLogicScope(v string) *UpdateConfigRuleRequest
	GetTagKeyLogicScope() *string
	SetTagKeyScope(v string) *UpdateConfigRuleRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *UpdateConfigRuleRequest
	GetTagValueScope() *string
	SetTagsScope(v []*UpdateConfigRuleRequestTagsScope) *UpdateConfigRuleRequest
	GetTagsScope() []*UpdateConfigRuleRequestTagsScope
}

type UpdateConfigRuleRequest struct {
	// The client token that you want to use to ensure the idempotency of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.``
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the rule.
	//
	// For more information about how to query the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-a260626622af0005****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// For more information about how to query the name of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// 	- ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// 	- ScheduledNotification: The rule is periodically triggered.
	//
	// >  This parameter applies only to custom rules.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// The description of the rule. You can enter up to 500 characters.
	//
	// example:
	//
	// The description of the rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the regions excluded from the compliance evaluations performed by the rule. Separate multiple region IDs with commas (,).
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
	// The IDs of the resources excluded from the compliance evaluations performed by the rule. Separate multiple resource IDs with commas (,).
	//
	// >  This parameter applies only to a managed rule.
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The scope of the tag that is excluded.
	ExcludeTagsScope []*UpdateConfigRuleRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// Optional. The extended content of the resource. This parameter can be used together with the MaximumExecutionFrequency parameter when the MaximumExecutionFrequency parameter is set to TwentyFour_Hours to specify the trigger time.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"fixedHour":"12"}
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The input parameters of the rule.
	//
	// example:
	//
	// {"tag1Key":"ECS","tag1Value":"test"}
	InputParameters map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The interval at which the rule is triggered. Valid values:
	//
	// 	- One_Hour
	//
	// 	- Three_Hours
	//
	// 	- Six_Hours
	//
	// 	- Twelve_Hours
	//
	// 	- TwentyFour_Hours (default)
	//
	// >  This parameter is required if the `ConfigRuleTriggerTypes` parameter is set to `ScheduledNotification`.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The IDs of the regions to which the rule applies. Separate multiple region IDs with commas (,).
	//
	// >  This parameter applies only to a managed rule.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The IDs of the resource groups to which the rule applies. Separate multiple resource group IDs with commas (,).
	//
	// >  This parameter applies only to a managed rule.
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
	// The names of the resource to which the rule applies.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// i-xxx
	ResourceNameScope *string `json:"ResourceNameScope,omitempty" xml:"ResourceNameScope,omitempty"`
	// The type of the resource to be evaluated by the rule. Separate multiple resource types with commas (,).
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypesScope []*string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
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
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Deprecated
	//
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	Tag []*UpdateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The logical relationship when parameter `TagsScope` takes multiple values, for example: When the parameter `TagsScope` is `"TagsScope.1.TagKey":"a", "TagsScope.1.TagValue":"a", "TagsScope.2.TagKey":"b", "TagsScope.2.TagValue":"b"`, if this parameter is set to` AND`, it means that the rule only applies to resources bound with both tags `a:a` and `b:b`. If not specified, the default logic is `OR`.
	//
	// It can also be used for the deprecated field `TagKeyScope` (not recommended), for example: When the parameter `TagKeyScope` has a value of `ECS`,`OSS`, if this parameter is set to `AND`, it means that the rule only applies to resources bound with both labels `ECS` and `OSS`.
	//
	// Values:
	//
	//  - AND: And.
	//
	//  - OR: Or.
	//
	// example:
	//
	// AND
	TagKeyLogicScope *string `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. We recommend that you use the `TagsScope` parameter.
	//
	// The tag key used to filter resources. The rule applies only to the resources with the specified tag key.
	//
	// >  This parameter applies only to a managed rule. You must configure the `TagKeyScope` and `TagValueScope` parameters at the same time.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. We recommend that you use the `TagsScope` parameter.
	//
	// The tag value used to filter resources. The rule applies only to the resources that use the specified tag value.
	//
	// >  This parameter applies only to a managed rule. You must configure the `TagKeyScope` and `TagValueScope` parameters at the same time.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tag scope.
	TagsScope []*UpdateConfigRuleRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s UpdateConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateConfigRuleRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *UpdateConfigRuleRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *UpdateConfigRuleRequest) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *UpdateConfigRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConfigRuleRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateConfigRuleRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateConfigRuleRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateConfigRuleRequest) GetExcludeTagsScope() []*UpdateConfigRuleRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *UpdateConfigRuleRequest) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *UpdateConfigRuleRequest) GetInputParameters() map[string]interface{} {
	return s.InputParameters
}

func (s *UpdateConfigRuleRequest) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *UpdateConfigRuleRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateConfigRuleRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateConfigRuleRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateConfigRuleRequest) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *UpdateConfigRuleRequest) GetResourceTypesScope() []*string {
	return s.ResourceTypesScope
}

func (s *UpdateConfigRuleRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateConfigRuleRequest) GetTag() []*UpdateConfigRuleRequestTag {
	return s.Tag
}

func (s *UpdateConfigRuleRequest) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *UpdateConfigRuleRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *UpdateConfigRuleRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *UpdateConfigRuleRequest) GetTagsScope() []*UpdateConfigRuleRequestTagsScope {
	return s.TagsScope
}

func (s *UpdateConfigRuleRequest) SetClientToken(v string) *UpdateConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetConfigRuleId(v string) *UpdateConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetConfigRuleName(v string) *UpdateConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetConfigRuleTriggerTypes(v string) *UpdateConfigRuleRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetDescription(v string) *UpdateConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetExcludeRegionIdsScope(v string) *UpdateConfigRuleRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateConfigRuleRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetExcludeResourceIdsScope(v string) *UpdateConfigRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetExcludeTagsScope(v []*UpdateConfigRuleRequestExcludeTagsScope) *UpdateConfigRuleRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *UpdateConfigRuleRequest) SetExtendContent(v string) *UpdateConfigRuleRequest {
	s.ExtendContent = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetInputParameters(v map[string]interface{}) *UpdateConfigRuleRequest {
	s.InputParameters = v
	return s
}

func (s *UpdateConfigRuleRequest) SetMaximumExecutionFrequency(v string) *UpdateConfigRuleRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetRegionIdsScope(v string) *UpdateConfigRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetResourceGroupIdsScope(v string) *UpdateConfigRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetResourceIdsScope(v string) *UpdateConfigRuleRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetResourceNameScope(v string) *UpdateConfigRuleRequest {
	s.ResourceNameScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetResourceTypesScope(v []*string) *UpdateConfigRuleRequest {
	s.ResourceTypesScope = v
	return s
}

func (s *UpdateConfigRuleRequest) SetRiskLevel(v int32) *UpdateConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetTag(v []*UpdateConfigRuleRequestTag) *UpdateConfigRuleRequest {
	s.Tag = v
	return s
}

func (s *UpdateConfigRuleRequest) SetTagKeyLogicScope(v string) *UpdateConfigRuleRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetTagKeyScope(v string) *UpdateConfigRuleRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetTagValueScope(v string) *UpdateConfigRuleRequest {
	s.TagValueScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetTagsScope(v []*UpdateConfigRuleRequestTagsScope) *UpdateConfigRuleRequest {
	s.TagsScope = v
	return s
}

func (s *UpdateConfigRuleRequest) Validate() error {
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

type UpdateConfigRuleRequestExcludeTagsScope struct {
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

func (s UpdateConfigRuleRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateConfigRuleRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateConfigRuleRequestExcludeTagsScope) SetTagKey(v string) *UpdateConfigRuleRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateConfigRuleRequestExcludeTagsScope) SetTagValue(v string) *UpdateConfigRuleRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateConfigRuleRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type UpdateConfigRuleRequestTag struct {
	// The key of tag N to add to the key pair. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: **1 to 20**. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` and `acs:`.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateConfigRuleRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleRequestTag) GetKey() *string {
	return s.Key
}

func (s *UpdateConfigRuleRequestTag) GetValue() *string {
	return s.Value
}

func (s *UpdateConfigRuleRequestTag) SetKey(v string) *UpdateConfigRuleRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateConfigRuleRequestTag) SetValue(v string) *UpdateConfigRuleRequestTag {
	s.Value = &v
	return s
}

func (s *UpdateConfigRuleRequestTag) Validate() error {
	return dara.Validate(s)
}

type UpdateConfigRuleRequestTagsScope struct {
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

func (s UpdateConfigRuleRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleRequestTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateConfigRuleRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateConfigRuleRequestTagsScope) SetTagKey(v string) *UpdateConfigRuleRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateConfigRuleRequestTagsScope) SetTagValue(v string) *UpdateConfigRuleRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateConfigRuleRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
