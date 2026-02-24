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
	SetConditions(v string) *UpdateAggregateConfigRuleShrinkRequest
	GetConditions() *string
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
	// The rule applies only to resources of the specified member accounts. Separate multiple member account IDs with a comma (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// 115748125982****
	AccountIdsScope *string `json:"AccountIdsScope,omitempty" xml:"AccountIdsScope,omitempty"`
	// The ID of the account group.
	//
	// For more information, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// A client token used to ensure the idempotence of the request. Make sure that the client token is unique for each request. The `ClientToken` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {"ComplianceConditions":"{\"operator\":\"and\",\"children\":[{\"operator\":\"StringEquals\",\"featurePath\":\"$.Status\",\"desired\":\"1\",\"featureSource\":\"CONFIGURATION\"}]}"}
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The rule ID.
	//
	// For more information, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-4e3d626622af0080****
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
	// > You can modify this parameter only for custom rules.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// 最多可以定义6组标签。如果资源同时具有指定的所有标签，则视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The member accounts to exclude. The rule does not apply to resources of these member accounts. Separate multiple member account IDs with a comma (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// 120886317861****
	ExcludeAccountIdsScope *string `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	// The folders to exclude. The rule does not apply to resources of member accounts in these folders. Separate multiple folder IDs with a comma (,).
	//
	// > - This parameter applies only to rules in a global account group.
	//
	// >
	//
	// > - This parameter applies only to rule templates.
	//
	// example:
	//
	// fd-pWmkqZ****
	ExcludeFolderIdsScope *string `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	// The regions to exclude. The rule does not apply to resources in these regions. Separate multiple region IDs with a comma (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The resource groups to exclude. The rule does not apply to resources in these resource groups. Separate multiple resource group IDs with a comma (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The resources to exclude. The rule does not apply to these resources. Separate multiple resource IDs with a comma (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The excluded tag scope.
	ExcludeTagsScope []*UpdateAggregateConfigRuleShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The rule applies only to resources of member accounts in the specified folders.
	//
	// > - This parameter applies only to rules in a global account group.
	//
	// >
	//
	// > - This parameter applies only to rule templates.
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
	// - TwentyFour_Hours: 24 hours.
	//
	// > This parameter is required if you set `ConfigRuleTriggerTypes` to `ScheduledNotification`.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The rule applies only to resources in the specified regions. Separate multiple region IDs with a comma (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The rule applies only to resources in the specified resource groups. Separate multiple resource group IDs with a comma (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The rule applies only to the specified resources. Separate multiple resource IDs with a comma (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The rule applies only to resources with the specified name.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// i-xxx
	ResourceNameScope *string `json:"ResourceNameScope,omitempty" xml:"ResourceNameScope,omitempty"`
	// The resource types that the rule evaluates. Separate multiple resource types with a comma (,).
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
	// The tags of the resource. This input parameter is deprecated and is ignored if specified.
	//
	// A maximum of 20 tags can be attached.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The logical relationship for multiple tags in the `TagsScope` parameter. For example, if you set the `TagsScope` parameter to `"TagsScope.1.TagKey":"a","TagsScope.1.TagValue":"a","TagsScope.2.TagKey":"b","TagsScope.2.TagValue":"b"` and set this parameter to `AND`, the rule applies only to resources that have both the `a:a` and `b:b` tags. If you do not set this parameter, the default value is `OR`.
	//
	// This parameter also works with the deprecated `TagKeyScope` parameter. For example, if you set the `TagKeyScope` parameter to `ECS,OSS` and set this parameter to `AND`, the rule applies only to resources that have both the `ECS` and `OSS` tags.
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
	// > This parameter applies only to rule templates. You must specify both `TagKeyScope` and `TagValueScope`.
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
	// > This parameter applies only to rule templates. You must specify both `TagKeyScope` and `TagValueScope`.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tag scope.
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

func (s *UpdateAggregateConfigRuleShrinkRequest) GetConditions() *string {
	return s.Conditions
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

func (s *UpdateAggregateConfigRuleShrinkRequest) SetConditions(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.Conditions = &v
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
