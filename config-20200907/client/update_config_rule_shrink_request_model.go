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
	SetConditions(v string) *UpdateConfigRuleShrinkRequest
	GetConditions() *string
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
	// A client token used to ensure the idempotence of the request. You can use a client to generate the token, but you must make sure that the token is unique for different requests. The `ClientToken` parameter can contain only ASCII characters and cannot be more than 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// {"ComplianceConditions":"{\"operator\":\"and\",\"children\":[{\"operator\":\"StringEquals\",\"featurePath\":\"$.Status\",\"desired\":\"1\",\"featureSource\":\"CONFIGURATION\"}]}"}
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The rule ID.
	//
	// For more information, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-a260626622af0005****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// For more information, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// 存在所有指定标签
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// - ScheduledNotification: The rule is triggered on a regular basis.
	//
	// > This parameter can be modified only for custom rules.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// The description of the rule. The description can be up to 500 characters in length.
	//
	// example:
	//
	// 最多可以定义6组标签。如果资源同时具有指定的所有标签，则视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The regions where the rule is not effective. The system does not evaluate resources in these regions. To specify multiple region IDs, separate them with a comma (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The resource groups where the rule is not effective. The system does not evaluate resources in these resource groups. To specify multiple resource group IDs, separate them with a comma (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The resources that are not evaluated by the rule. The system does not evaluate these resources. To specify multiple resource IDs, separate them with a comma (,).
	//
	// > This parameter applies only to managed rules.
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The tags that are used to exclude resources.
	ExcludeTagsScope []*UpdateConfigRuleShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The extended content. This parameter is optional. You can use this parameter with a 24-hour trigger period to set the trigger time.
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
	// The frequency at which the rule is run. Valid values:
	//
	// - One_Hour: 1 hour.
	//
	// - Three_Hours: 3 hours.
	//
	// - Six_Hours: 6 hours.
	//
	// - Twelve_Hours: 12 hours.
	//
	// - TwentyFour_Hours (default): 24 hours.
	//
	// > This parameter is required when `ConfigRuleTriggerTypes` is set to `ScheduledNotification`.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The regions where the rule is effective. The rule evaluates only resources in these regions. To specify multiple region IDs, separate them with a comma (,).
	//
	// > This parameter applies only to managed rules.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The resource groups where the rule is effective. The rule evaluates only resources in these resource groups. To specify multiple resource group IDs, separate them with a comma (,).
	//
	// > This parameter applies only to managed rules.
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The resources that the rule evaluates. To specify multiple resource IDs, separate them with a comma (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The names of the resources that the rule evaluates.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// i-xxx
	ResourceNameScope *string `json:"ResourceNameScope,omitempty" xml:"ResourceNameScope,omitempty"`
	// The resource types that the rule evaluates. To specify multiple resource types, separate them with a comma (,).
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypesScopeShrink *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
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
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// Deprecated
	//
	// The tags of the resource. This parameter is deprecated. The value that you specify for this parameter does not take effect.
	//
	// You can add up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The logical relationship for the tags that you specify for the `TagsScope` parameter. For example, if you set the `TagsScope` parameter to `"TagsScope.1.TagKey":"a","TagsScope.1.TagValue":"a","TagsScope.2.TagKey":"b","TagsScope.2.TagValue":"b"` and set this parameter to `AND`, the rule applies only to resources that have both the `a:a` and `b:b` tags. If you do not specify this parameter, the `OR` logic is used by default.
	//
	// This parameter also applies to the deprecated `TagKeyScope` parameter, but this is not recommended. For example, if you set the `TagKeyScope` parameter to `ECS,OSS` and set this parameter to `AND`, the rule applies only to resources that have both the `ECS` and `OSS` tags.
	//
	// Valid values:
	//
	// - AND
	//
	// - OR
	//
	// example:
	//
	// AND
	TagKeyLogicScope *string `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter instead.
	//
	// The rule is effective only for resources that have the specified tag.
	//
	// > This parameter applies only to managed rules. The `TagKeyScope` and `TagValueScope` parameters must be specified at the same time.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter instead.
	//
	// The rule is effective only for resources that have the specified tag.
	//
	// > This parameter applies only to managed rules. The `TagKeyScope` and `TagValueScope` parameters must be specified at the same time.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tags that are used to filter resources.
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

func (s *UpdateConfigRuleShrinkRequest) GetConditions() *string {
	return s.Conditions
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

func (s *UpdateConfigRuleShrinkRequest) SetConditions(v string) *UpdateConfigRuleShrinkRequest {
	s.Conditions = &v
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
