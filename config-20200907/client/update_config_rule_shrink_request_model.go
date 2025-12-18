// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateConfigRuleShrinkRequest
	GetClientToken() *string
	SetConfigRuleId(v string) *UpdateConfigRuleShrinkRequest
	GetConfigRuleId() *string
	SetConfigRuleName(v string) *UpdateConfigRuleShrinkRequest
	GetConfigRuleName() *string
	SetConfigRuleTriggerTypes(v string) *UpdateConfigRuleShrinkRequest
	GetConfigRuleTriggerTypes() *string
	SetDescription(v string) *UpdateConfigRuleShrinkRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *UpdateConfigRuleShrinkRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateConfigRuleShrinkRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateConfigRuleShrinkRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*UpdateConfigRuleShrinkRequestExcludeTagsScope) *UpdateConfigRuleShrinkRequest
	GetExcludeTagsScope() []*UpdateConfigRuleShrinkRequestExcludeTagsScope
	SetExtendContent(v string) *UpdateConfigRuleShrinkRequest
	GetExtendContent() *string
	SetInputParametersShrink(v string) *UpdateConfigRuleShrinkRequest
	GetInputParametersShrink() *string
	SetMaximumExecutionFrequency(v string) *UpdateConfigRuleShrinkRequest
	GetMaximumExecutionFrequency() *string
	SetRegionIdsScope(v string) *UpdateConfigRuleShrinkRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateConfigRuleShrinkRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateConfigRuleShrinkRequest
	GetResourceIdsScope() *string
	SetResourceNameScope(v string) *UpdateConfigRuleShrinkRequest
	GetResourceNameScope() *string
	SetResourceTypesScopeShrink(v string) *UpdateConfigRuleShrinkRequest
	GetResourceTypesScopeShrink() *string
	SetRiskLevel(v int32) *UpdateConfigRuleShrinkRequest
	GetRiskLevel() *int32
	SetTagShrink(v string) *UpdateConfigRuleShrinkRequest
	GetTagShrink() *string
	SetTagKeyLogicScope(v string) *UpdateConfigRuleShrinkRequest
	GetTagKeyLogicScope() *string
	SetTagKeyScope(v string) *UpdateConfigRuleShrinkRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *UpdateConfigRuleShrinkRequest
	GetTagValueScope() *string
	SetTagsScope(v []*UpdateConfigRuleShrinkRequestTagsScope) *UpdateConfigRuleShrinkRequest
	GetTagsScope() []*UpdateConfigRuleShrinkRequestTagsScope
}

type UpdateConfigRuleShrinkRequest struct {
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
	ExcludeTagsScope []*UpdateConfigRuleShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	InputParametersShrink *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
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
	ResourceTypesScopeShrink *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
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
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
	TagsScope []*UpdateConfigRuleShrinkRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s UpdateConfigRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateConfigRuleShrinkRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *UpdateConfigRuleShrinkRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *UpdateConfigRuleShrinkRequest) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *UpdateConfigRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConfigRuleShrinkRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateConfigRuleShrinkRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateConfigRuleShrinkRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateConfigRuleShrinkRequest) GetExcludeTagsScope() []*UpdateConfigRuleShrinkRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *UpdateConfigRuleShrinkRequest) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *UpdateConfigRuleShrinkRequest) GetInputParametersShrink() *string {
	return s.InputParametersShrink
}

func (s *UpdateConfigRuleShrinkRequest) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *UpdateConfigRuleShrinkRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateConfigRuleShrinkRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateConfigRuleShrinkRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateConfigRuleShrinkRequest) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *UpdateConfigRuleShrinkRequest) GetResourceTypesScopeShrink() *string {
	return s.ResourceTypesScopeShrink
}

func (s *UpdateConfigRuleShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateConfigRuleShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *UpdateConfigRuleShrinkRequest) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *UpdateConfigRuleShrinkRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *UpdateConfigRuleShrinkRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *UpdateConfigRuleShrinkRequest) GetTagsScope() []*UpdateConfigRuleShrinkRequestTagsScope {
	return s.TagsScope
}

func (s *UpdateConfigRuleShrinkRequest) SetClientToken(v string) *UpdateConfigRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetConfigRuleId(v string) *UpdateConfigRuleShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetConfigRuleName(v string) *UpdateConfigRuleShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetConfigRuleTriggerTypes(v string) *UpdateConfigRuleShrinkRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetDescription(v string) *UpdateConfigRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetExcludeRegionIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetExcludeResourceIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetExcludeTagsScope(v []*UpdateConfigRuleShrinkRequestExcludeTagsScope) *UpdateConfigRuleShrinkRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetExtendContent(v string) *UpdateConfigRuleShrinkRequest {
	s.ExtendContent = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetInputParametersShrink(v string) *UpdateConfigRuleShrinkRequest {
	s.InputParametersShrink = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetMaximumExecutionFrequency(v string) *UpdateConfigRuleShrinkRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetRegionIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetResourceGroupIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetResourceIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetResourceNameScope(v string) *UpdateConfigRuleShrinkRequest {
	s.ResourceNameScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetResourceTypesScopeShrink(v string) *UpdateConfigRuleShrinkRequest {
	s.ResourceTypesScopeShrink = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetRiskLevel(v int32) *UpdateConfigRuleShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetTagShrink(v string) *UpdateConfigRuleShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetTagKeyLogicScope(v string) *UpdateConfigRuleShrinkRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetTagKeyScope(v string) *UpdateConfigRuleShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetTagValueScope(v string) *UpdateConfigRuleShrinkRequest {
	s.TagValueScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetTagsScope(v []*UpdateConfigRuleShrinkRequestTagsScope) *UpdateConfigRuleShrinkRequest {
	s.TagsScope = v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) Validate() error {
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

type UpdateConfigRuleShrinkRequestExcludeTagsScope struct {
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

func (s UpdateConfigRuleShrinkRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleShrinkRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleShrinkRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateConfigRuleShrinkRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateConfigRuleShrinkRequestExcludeTagsScope) SetTagKey(v string) *UpdateConfigRuleShrinkRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequestExcludeTagsScope) SetTagValue(v string) *UpdateConfigRuleShrinkRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type UpdateConfigRuleShrinkRequestTagsScope struct {
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

func (s UpdateConfigRuleShrinkRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleShrinkRequestTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleShrinkRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateConfigRuleShrinkRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateConfigRuleShrinkRequestTagsScope) SetTagKey(v string) *UpdateConfigRuleShrinkRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequestTagsScope) SetTagValue(v string) *UpdateConfigRuleShrinkRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
