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
	SetConditions(v string) *UpdateConfigRuleRequest
	GetConditions() *string
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
	// A client token used to ensure the idempotence of the request. Generate a unique token on your client for each request. The `ClientToken` parameter supports only ASCII characters and must not exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The condition for a custom conditional rule, in JSON format.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"ComplianceConditions":"{\\"operator\\":\\"and\\",\\"children\\":[{\\"operator\\":\\"StringEquals\\",\\"featurePath\\":\\"$.Status\\",\\"desired\\":\\"1\\",\\"featureSource\\":\\"CONFIGURATION\\"}]}"}
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
	// - ConfigurationItemChangeNotification: The rule triggers on configuration changes.
	//
	// - ScheduledNotification: The rule triggers on a schedule.
	//
	// > You can modify this parameter only for custom rules.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// The rule description. The description can be up to 500 characters long.
	//
	// example:
	//
	// 最多可以定义6组标签。如果资源同时具有指定的所有标签，则视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The regions where the rule does not apply. To specify multiple region IDs, separate them with a comma (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The resource groups where the rule does not apply. To specify multiple resource group IDs, separate them with a comma (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The resources that the rule does not evaluate. To specify multiple resource IDs, separate them with a comma (,).
	//
	// > This parameter applies only to managed rules.
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The tags used to exclude resources.
	ExcludeTagsScope []*UpdateConfigRuleRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// Optional. Extended content used with a 24-hour trigger period to set the trigger time.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"fixedHour":"12"}
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The rule parameters.
	//
	// example:
	//
	// {"tag1Key":"ECS","tag1Value":"test"}
	InputParameters map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The frequency at which the rule runs. Valid values:
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
	// > Set this parameter when `ConfigRuleTriggerTypes` is set to `ScheduledNotification`.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The regions where the rule applies. To specify multiple region IDs, separate them with a comma (,).
	//
	// > This parameter applies only to managed rules.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The resource groups where the rule applies. To specify multiple resource group IDs, separate them with a comma (,).
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
	// The tags of the resource. This parameter is deprecated. Ignore it. Values passed for this parameter have no effect.
	//
	// You can add up to 20 tags.
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
	ResourceTypesScope []*string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
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
	// The tags of the resource. This parameter is deprecated. Ignore it. Values passed for this parameter have no effect.
	//
	// You can add up to 20 tags.
	Tag []*UpdateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The logical relationship between tags in the `TagsScope` parameter. For example, if you set `TagsScope` to `"TagsScope.1.TagKey":"a","TagsScope.1.TagValue":"a","TagsScope.2.TagKey":"b","TagsScope.2.TagValue":"b"` and set this parameter to `AND`, the rule applies only to resources that have both the `a:a` and `b:b` tags. If you omit this parameter, the default logic is `OR`.
	//
	// This parameter also works with the deprecated `TagKeyScope` parameter, but this is not recommended. For example, if you set `TagKeyScope` to `ECS,OSS` and set this parameter to `AND`, the rule applies only to resources that have both the `ECS` and `OSS` tags.
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
	// The rule applies only to resources that have the specified tag.
	//
	// > This parameter applies only to managed rules. You must specify both `TagKeyScope` and `TagValueScope`.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter instead.
	//
	// The rule applies only to resources that have the specified tag.
	//
	// > This parameter applies only to managed rules. You must specify both `TagKeyScope` and `TagValueScope`.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// Scope of the tag
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

func (s *UpdateConfigRuleRequest) GetConditions() *string {
	return s.Conditions
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

func (s *UpdateConfigRuleRequest) SetConditions(v string) *UpdateConfigRuleRequest {
	s.Conditions = &v
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
	// The tag key.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
