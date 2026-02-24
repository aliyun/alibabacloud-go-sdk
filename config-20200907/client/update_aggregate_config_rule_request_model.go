// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetAccountIdsScope() *string
	SetAggregatorId(v string) *UpdateAggregateConfigRuleRequest
	GetAggregatorId() *string
	SetClientToken(v string) *UpdateAggregateConfigRuleRequest
	GetClientToken() *string
	SetConditions(v string) *UpdateAggregateConfigRuleRequest
	GetConditions() *string
	SetConfigRuleId(v string) *UpdateAggregateConfigRuleRequest
	GetConfigRuleId() *string
	SetConfigRuleName(v string) *UpdateAggregateConfigRuleRequest
	GetConfigRuleName() *string
	SetConfigRuleTriggerTypes(v string) *UpdateAggregateConfigRuleRequest
	GetConfigRuleTriggerTypes() *string
	SetDescription(v string) *UpdateAggregateConfigRuleRequest
	GetDescription() *string
	SetExcludeAccountIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetExcludeAccountIdsScope() *string
	SetExcludeFolderIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetExcludeFolderIdsScope() *string
	SetExcludeRegionIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*UpdateAggregateConfigRuleRequestExcludeTagsScope) *UpdateAggregateConfigRuleRequest
	GetExcludeTagsScope() []*UpdateAggregateConfigRuleRequestExcludeTagsScope
	SetFolderIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetFolderIdsScope() *string
	SetInputParameters(v map[string]interface{}) *UpdateAggregateConfigRuleRequest
	GetInputParameters() map[string]interface{}
	SetMaximumExecutionFrequency(v string) *UpdateAggregateConfigRuleRequest
	GetMaximumExecutionFrequency() *string
	SetRegionIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *UpdateAggregateConfigRuleRequest
	GetResourceIdsScope() *string
	SetResourceNameScope(v string) *UpdateAggregateConfigRuleRequest
	GetResourceNameScope() *string
	SetResourceTypesScope(v []*string) *UpdateAggregateConfigRuleRequest
	GetResourceTypesScope() []*string
	SetRiskLevel(v int32) *UpdateAggregateConfigRuleRequest
	GetRiskLevel() *int32
	SetTag(v []*UpdateAggregateConfigRuleRequestTag) *UpdateAggregateConfigRuleRequest
	GetTag() []*UpdateAggregateConfigRuleRequestTag
	SetTagKeyLogicScope(v string) *UpdateAggregateConfigRuleRequest
	GetTagKeyLogicScope() *string
	SetTagKeyScope(v string) *UpdateAggregateConfigRuleRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *UpdateAggregateConfigRuleRequest
	GetTagValueScope() *string
	SetTagsScope(v []*UpdateAggregateConfigRuleRequestTagsScope) *UpdateAggregateConfigRuleRequest
	GetTagsScope() []*UpdateAggregateConfigRuleRequestTagsScope
}

type UpdateAggregateConfigRuleRequest struct {
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
	ExcludeTagsScope []*UpdateAggregateConfigRuleRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	InputParameters map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
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
	// The tags of the resource. This input parameter is deprecated and is ignored if specified.
	//
	// A maximum of 20 tags can be attached.
	Tag []*UpdateAggregateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	TagsScope []*UpdateAggregateConfigRuleRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s UpdateAggregateConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleRequest) GetAccountIdsScope() *string {
	return s.AccountIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregateConfigRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAggregateConfigRuleRequest) GetConditions() *string {
	return s.Conditions
}

func (s *UpdateAggregateConfigRuleRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *UpdateAggregateConfigRuleRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *UpdateAggregateConfigRuleRequest) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *UpdateAggregateConfigRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggregateConfigRuleRequest) GetExcludeAccountIdsScope() *string {
	return s.ExcludeAccountIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetExcludeFolderIdsScope() *string {
	return s.ExcludeFolderIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetExcludeTagsScope() []*UpdateAggregateConfigRuleRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetFolderIdsScope() *string {
	return s.FolderIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetInputParameters() map[string]interface{} {
	return s.InputParameters
}

func (s *UpdateAggregateConfigRuleRequest) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *UpdateAggregateConfigRuleRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *UpdateAggregateConfigRuleRequest) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *UpdateAggregateConfigRuleRequest) GetResourceTypesScope() []*string {
	return s.ResourceTypesScope
}

func (s *UpdateAggregateConfigRuleRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *UpdateAggregateConfigRuleRequest) GetTag() []*UpdateAggregateConfigRuleRequestTag {
	return s.Tag
}

func (s *UpdateAggregateConfigRuleRequest) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *UpdateAggregateConfigRuleRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *UpdateAggregateConfigRuleRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *UpdateAggregateConfigRuleRequest) GetTagsScope() []*UpdateAggregateConfigRuleRequestTagsScope {
	return s.TagsScope
}

func (s *UpdateAggregateConfigRuleRequest) SetAccountIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.AccountIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetAggregatorId(v string) *UpdateAggregateConfigRuleRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetClientToken(v string) *UpdateAggregateConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetConditions(v string) *UpdateAggregateConfigRuleRequest {
	s.Conditions = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetConfigRuleId(v string) *UpdateAggregateConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetConfigRuleName(v string) *UpdateAggregateConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetConfigRuleTriggerTypes(v string) *UpdateAggregateConfigRuleRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetDescription(v string) *UpdateAggregateConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeAccountIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeFolderIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeRegionIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeResourceIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeTagsScope(v []*UpdateAggregateConfigRuleRequestExcludeTagsScope) *UpdateAggregateConfigRuleRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetFolderIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.FolderIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetInputParameters(v map[string]interface{}) *UpdateAggregateConfigRuleRequest {
	s.InputParameters = v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetMaximumExecutionFrequency(v string) *UpdateAggregateConfigRuleRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetRegionIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetResourceIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetResourceNameScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ResourceNameScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetResourceTypesScope(v []*string) *UpdateAggregateConfigRuleRequest {
	s.ResourceTypesScope = v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetRiskLevel(v int32) *UpdateAggregateConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetTag(v []*UpdateAggregateConfigRuleRequestTag) *UpdateAggregateConfigRuleRequest {
	s.Tag = v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetTagKeyLogicScope(v string) *UpdateAggregateConfigRuleRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetTagKeyScope(v string) *UpdateAggregateConfigRuleRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetTagValueScope(v string) *UpdateAggregateConfigRuleRequest {
	s.TagValueScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetTagsScope(v []*UpdateAggregateConfigRuleRequestTagsScope) *UpdateAggregateConfigRuleRequest {
	s.TagsScope = v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) Validate() error {
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

type UpdateAggregateConfigRuleRequestExcludeTagsScope struct {
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

func (s UpdateAggregateConfigRuleRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAggregateConfigRuleRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAggregateConfigRuleRequestExcludeTagsScope) SetTagKey(v string) *UpdateAggregateConfigRuleRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequestExcludeTagsScope) SetTagValue(v string) *UpdateAggregateConfigRuleRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type UpdateAggregateConfigRuleRequestTag struct {
	// The tag key of the resource.
	//
	// You can add up to 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// You can add up to 20 tag values.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateAggregateConfigRuleRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleRequestTag) GetKey() *string {
	return s.Key
}

func (s *UpdateAggregateConfigRuleRequestTag) GetValue() *string {
	return s.Value
}

func (s *UpdateAggregateConfigRuleRequestTag) SetKey(v string) *UpdateAggregateConfigRuleRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequestTag) SetValue(v string) *UpdateAggregateConfigRuleRequestTag {
	s.Value = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequestTag) Validate() error {
	return dara.Validate(s)
}

type UpdateAggregateConfigRuleRequestTagsScope struct {
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

func (s UpdateAggregateConfigRuleRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigRuleRequestTagsScope) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateAggregateConfigRuleRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateAggregateConfigRuleRequestTagsScope) SetTagKey(v string) *UpdateAggregateConfigRuleRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequestTagsScope) SetTagValue(v string) *UpdateAggregateConfigRuleRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
