// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateConfigRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetAccountIdsScope() *string
	SetAggregatorId(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetAggregatorId() *string
	SetClientToken(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetClientToken() *string
	SetConfigRuleId(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetConfigRuleId() *string
	SetConfigRuleName(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetConfigRuleName() *string
	SetConfigRuleTriggerTypes(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetConfigRuleTriggerTypes() *string
	SetDescription(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetDescription() *string
	SetExcludeAccountIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetExcludeAccountIdsScope() *string
	SetExcludeFolderIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetExcludeFolderIdsScope() *string
	SetExcludeRegionIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) *UpdateAggregateConfigRuleShrinkRequest
	GetExcludeTagsScope() []*UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope
	SetFolderIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetFolderIdsScope() *string
	SetInputParametersShrink(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetInputParametersShrink() *string
	SetMaximumExecutionFrequency(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetMaximumExecutionFrequency() *string
	SetRegionIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetResourceIdsScope() *string
	SetResourceNameScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetResourceNameScope() *string
	SetResourceTypesScopeShrink(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetResourceTypesScopeShrink() *string
	SetRiskLevel(v int32) *UpdateAggregateConfigRuleShrinkRequest
	GetRiskLevel() *int32
	SetTagShrink(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetTagShrink() *string
	SetTagKeyLogicScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetTagKeyLogicScope() *string
	SetTagKeyScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetTagValueScope() *string
	SetTagsScope(v []*UpdateAggregateConfigRuleShrinkRequestTagsScope) *UpdateAggregateConfigRuleShrinkRequest
	GetTagsScope() []*UpdateAggregateConfigRuleShrinkRequestTagsScope
}

type UpdateAggregateConfigRuleShrinkRequest struct {
	// The IDs of the member accounts to which the rule applies, which means that the resources within the member accounts are evaluated based on the rule. Separate multiple member account IDs with commas (,).
	//
	// example:
	//
	// 115748125982****
	AccountIdsScope *string `json:"AccountIdsScope,omitempty" xml:"AccountIdsScope,omitempty"`
	// The ID of the account group.
	//
	// For more information about how to query the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the rule.
	//
	// For more information about how to query the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-4e3d626622af0080****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// For more information about how to query the name of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// example:
	//
	// test_rule
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// 	- ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// 	- ScheduledNotification: The rule is periodically triggered.
	//
	// >  This parameter applies only to a custom rule.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// test_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the member accounts to which the rule does not apply, which means that the resources within the member accounts are not evaluated based on the rule. Separate multiple member account IDs with commas (,).
	//
	// >  This parameter applies only to a managed rule.
	//
	// example:
	//
	// 120886317861****
	ExcludeAccountIdsScope *string `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	// The IDs of the resource directories to which the rule does not apply, which means that the resources within member accounts in the resource directories are not evaluated based on the rule. Separate multiple resource directory IDs with commas (,).
	//
	// >
	//
	// 	- This parameter applies only to a rule of a global account group.
	//
	// 	- This parameter applies only to a managed rule.
	//
	// example:
	//
	// fd-pWmkqZ****
	ExcludeFolderIdsScope *string `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	// The IDs of the regions to which the rule not applies. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The IDs of the resource groups to which the rule not applies. Separate multiple resource group IDs with commas (,).
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
	// Exclude the specific tag scope of resources .
	ExcludeTagsScope []*UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The IDs of the resource directories to which the rule applies, which means that the resources within member accounts in the resource directories are evaluated based on the rule.
	//
	// >
	//
	// 	- This parameter applies only to a rule of a global account group.
	//
	// 	- This parameter applies only to a managed rule.
	//
	// example:
	//
	// fd-ZtHsRH****
	FolderIdsScope *string `json:"FolderIdsScope,omitempty" xml:"FolderIdsScope,omitempty"`
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
	// 	- TwentyFour_Hours
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
	// The IDs of the resources included from the compliance evaluations performed by the rule. Separate multiple resource IDs with commas (,).
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
	// The valid tag scope of resources.
	TagsScope []*UpdateAggregateConfigRuleShrinkRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s UpdateAggregateConfigRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetAccountIdsScope() *string {
	return s.AccountIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetExcludeAccountIdsScope() *string {
	return s.ExcludeAccountIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetExcludeFolderIdsScope() *string {
	return s.ExcludeFolderIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetExcludeTagsScope() []*UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetFolderIdsScope() *string {
	return s.FolderIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetInputParametersShrink() *string {
	return s.InputParametersShrink
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetResourceTypesScopeShrink() *string {
	return s.ResourceTypesScopeShrink
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) GetTagsScope() []*UpdateAggregateConfigRuleShrinkRequestTagsScope {
	return s.TagsScope
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetAccountIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.AccountIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetAggregatorId(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetClientToken(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetConfigRuleId(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetConfigRuleName(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetConfigRuleTriggerTypes(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetDescription(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeAccountIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeFolderIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeRegionIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeResourceIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeTagsScope(v []*UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetFolderIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.FolderIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetInputParametersShrink(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.InputParametersShrink = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetMaximumExecutionFrequency(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetRegionIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetResourceIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetResourceNameScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ResourceNameScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetResourceTypesScopeShrink(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ResourceTypesScopeShrink = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetRiskLevel(v int32) *UpdateAggregateConfigRuleShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetTagShrink(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetTagKeyLogicScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetTagKeyScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetTagValueScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.TagValueScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetTagsScope(v []*UpdateAggregateConfigRuleShrinkRequestTagsScope) *UpdateAggregateConfigRuleShrinkRequest {
	s.TagsScope = v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) Validate() error {
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

type UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope struct {
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

func (s UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) SetTagKey(v string) *UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) SetTagValue(v string) *UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type UpdateAggregateConfigRuleShrinkRequestTagsScope struct {
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

func (s UpdateAggregateConfigRuleShrinkRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleShrinkRequestTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleShrinkRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAggregateConfigRuleShrinkRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAggregateConfigRuleShrinkRequestTagsScope) SetTagKey(v string) *UpdateAggregateConfigRuleShrinkRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequestTagsScope) SetTagValue(v string) *UpdateAggregateConfigRuleShrinkRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
