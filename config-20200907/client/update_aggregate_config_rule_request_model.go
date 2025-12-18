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
	ExcludeTagsScope []*UpdateAggregateConfigRuleRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	Tag []*UpdateAggregateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// The tag key cannot be an empty string. The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs`:. The tag key cannot contain `http://` or `https://`.
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
