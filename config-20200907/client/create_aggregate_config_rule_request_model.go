// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetAccountIdsScope() *string
	SetAggregatorId(v string) *CreateAggregateConfigRuleRequest
	GetAggregatorId() *string
	SetClientToken(v string) *CreateAggregateConfigRuleRequest
	GetClientToken() *string
	SetConfigRuleName(v string) *CreateAggregateConfigRuleRequest
	GetConfigRuleName() *string
	SetConfigRuleTriggerTypes(v string) *CreateAggregateConfigRuleRequest
	GetConfigRuleTriggerTypes() *string
	SetDescription(v string) *CreateAggregateConfigRuleRequest
	GetDescription() *string
	SetExcludeAccountIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetExcludeAccountIdsScope() *string
	SetExcludeFolderIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetExcludeFolderIdsScope() *string
	SetExcludeRegionIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*CreateAggregateConfigRuleRequestExcludeTagsScope) *CreateAggregateConfigRuleRequest
	GetExcludeTagsScope() []*CreateAggregateConfigRuleRequestExcludeTagsScope
	SetExtendContent(v string) *CreateAggregateConfigRuleRequest
	GetExtendContent() *string
	SetFolderIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetFolderIdsScope() *string
	SetInputParameters(v map[string]interface{}) *CreateAggregateConfigRuleRequest
	GetInputParameters() map[string]interface{}
	SetMaximumExecutionFrequency(v string) *CreateAggregateConfigRuleRequest
	GetMaximumExecutionFrequency() *string
	SetRegionIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateAggregateConfigRuleRequest
	GetResourceIdsScope() *string
	SetResourceNameScope(v string) *CreateAggregateConfigRuleRequest
	GetResourceNameScope() *string
	SetResourceTypesScope(v []*string) *CreateAggregateConfigRuleRequest
	GetResourceTypesScope() []*string
	SetRiskLevel(v int32) *CreateAggregateConfigRuleRequest
	GetRiskLevel() *int32
	SetSourceIdentifier(v string) *CreateAggregateConfigRuleRequest
	GetSourceIdentifier() *string
	SetSourceOwner(v string) *CreateAggregateConfigRuleRequest
	GetSourceOwner() *string
	SetTag(v []*CreateAggregateConfigRuleRequestTag) *CreateAggregateConfigRuleRequest
	GetTag() []*CreateAggregateConfigRuleRequestTag
	SetTagKeyLogicScope(v string) *CreateAggregateConfigRuleRequest
	GetTagKeyLogicScope() *string
	SetTagKeyScope(v string) *CreateAggregateConfigRuleRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *CreateAggregateConfigRuleRequest
	GetTagValueScope() *string
	SetTagsScope(v []*CreateAggregateConfigRuleRequestTagsScope) *CreateAggregateConfigRuleRequest
	GetTagsScope() []*CreateAggregateConfigRuleRequestTagsScope
}

type CreateAggregateConfigRuleRequest struct {
	// The IDs of the member accounts to which the rule applies, which means that the resources within the member accounts are evaluated based on the rule. Separate multiple member account IDs with commas (,).
	//
	// example:
	//
	// 115748125982****
	AccountIdsScope *string `json:"AccountIdsScope,omitempty" xml:"AccountIdsScope,omitempty"`
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The rule name.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-default-encryption-kms
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// 	- ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// 	- ScheduledNotification: The rule is periodically triggered.
	//
	// This parameter is required.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// description of rule
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the member account to which the rule does not apply, which means that the resources within the member account are not evaluated based on the rule. Separate multiple member account IDs with commas (,).
	//
	// > This parameter applies only to a managed rule.
	//
	// example:
	//
	// 120886317861****
	ExcludeAccountIdsScope *string `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	// The ID of the resource directory to which the rule does not apply, which means that the resources within member accounts in the resource directory are not evaluated based on the rule. Separate multiple resource directory IDs with commas (,).
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
	// ExcludeResourceGroupIdsScope. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The ID of the resource to be excluded from the compliance evaluations performed by the rule. Separate multiple resource IDs with commas (,).
	//
	// > This parameter applies only to a managed rule.
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The scope of the tag that is excluded.
	ExcludeTagsScope []*CreateAggregateConfigRuleRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The extended content, which is temporarily only used to configure the trigger time with a 24-hour cycle trigger.
	//
	// example:
	//
	// {"fixedHour":"12"}
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The ID of the resource directory to which the rule applies, which means that the resources within member accounts in the resource directory are evaluated based on the rule.
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
	// The input parameter of the rule.
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
	// The ID of the region to which the rule applies. Separate multiple region IDs with commas (,).
	//
	// > This parameter applies only to a managed rule.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The ID of the resource group to which the rule applies. Separate multiple resource group IDs with commas (,).
	//
	// > This parameter applies only to a managed rule.
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
	// The type of the resource evaluated by the rule. Separate multiple resource types with commas (,).
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The identifier of the rule.
	//
	// 	- If you set the SourceOwner parameter to ALIYUN, set this parameter to the name of the managed rule.
	//
	// 	- If you set the SourceOwner parameter to CUSTOM_FC, set this parameter to the Alibaba Cloud Resource Name (ARN) of the relevant function in Function Compute.
	//
	// For more information about how to query the name of a managed rule, see [Managed rules](https://help.aliyun.com/document_detail/127404.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// required-tags
	SourceIdentifier *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- ALIYUN: a managed rule.
	//
	// 	- CUSTOM_FC: a custom rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	SourceOwner *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	// The tags.
	Tag []*CreateAggregateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The tag key used to filter resources. The rule applies only to the resources with the specified tag key. Separate multiple parameter values with commas (,).
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
	TagsScope []*CreateAggregateConfigRuleRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s CreateAggregateConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleRequest) GetAccountIdsScope() *string {
	return s.AccountIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *CreateAggregateConfigRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAggregateConfigRuleRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *CreateAggregateConfigRuleRequest) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *CreateAggregateConfigRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAggregateConfigRuleRequest) GetExcludeAccountIdsScope() *string {
	return s.ExcludeAccountIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetExcludeFolderIdsScope() *string {
	return s.ExcludeFolderIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetExcludeTagsScope() []*CreateAggregateConfigRuleRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *CreateAggregateConfigRuleRequest) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *CreateAggregateConfigRuleRequest) GetFolderIdsScope() *string {
	return s.FolderIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetInputParameters() map[string]interface{} {
	return s.InputParameters
}

func (s *CreateAggregateConfigRuleRequest) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *CreateAggregateConfigRuleRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateAggregateConfigRuleRequest) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *CreateAggregateConfigRuleRequest) GetResourceTypesScope() []*string {
	return s.ResourceTypesScope
}

func (s *CreateAggregateConfigRuleRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateAggregateConfigRuleRequest) GetSourceIdentifier() *string {
	return s.SourceIdentifier
}

func (s *CreateAggregateConfigRuleRequest) GetSourceOwner() *string {
	return s.SourceOwner
}

func (s *CreateAggregateConfigRuleRequest) GetTag() []*CreateAggregateConfigRuleRequestTag {
	return s.Tag
}

func (s *CreateAggregateConfigRuleRequest) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *CreateAggregateConfigRuleRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *CreateAggregateConfigRuleRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *CreateAggregateConfigRuleRequest) GetTagsScope() []*CreateAggregateConfigRuleRequestTagsScope {
	return s.TagsScope
}

func (s *CreateAggregateConfigRuleRequest) SetAccountIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.AccountIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetAggregatorId(v string) *CreateAggregateConfigRuleRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetClientToken(v string) *CreateAggregateConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetConfigRuleName(v string) *CreateAggregateConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetConfigRuleTriggerTypes(v string) *CreateAggregateConfigRuleRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetDescription(v string) *CreateAggregateConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeAccountIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeFolderIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeRegionIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeResourceGroupIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeResourceIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeTagsScope(v []*CreateAggregateConfigRuleRequestExcludeTagsScope) *CreateAggregateConfigRuleRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExtendContent(v string) *CreateAggregateConfigRuleRequest {
	s.ExtendContent = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetFolderIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.FolderIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetInputParameters(v map[string]interface{}) *CreateAggregateConfigRuleRequest {
	s.InputParameters = v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetMaximumExecutionFrequency(v string) *CreateAggregateConfigRuleRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetRegionIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetResourceGroupIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetResourceIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetResourceNameScope(v string) *CreateAggregateConfigRuleRequest {
	s.ResourceNameScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetResourceTypesScope(v []*string) *CreateAggregateConfigRuleRequest {
	s.ResourceTypesScope = v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetRiskLevel(v int32) *CreateAggregateConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetSourceIdentifier(v string) *CreateAggregateConfigRuleRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetSourceOwner(v string) *CreateAggregateConfigRuleRequest {
	s.SourceOwner = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetTag(v []*CreateAggregateConfigRuleRequestTag) *CreateAggregateConfigRuleRequest {
	s.Tag = v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetTagKeyLogicScope(v string) *CreateAggregateConfigRuleRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetTagKeyScope(v string) *CreateAggregateConfigRuleRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetTagValueScope(v string) *CreateAggregateConfigRuleRequest {
	s.TagValueScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetTagsScope(v []*CreateAggregateConfigRuleRequestTagsScope) *CreateAggregateConfigRuleRequest {
	s.TagsScope = v
	return s
}

func (s *CreateAggregateConfigRuleRequest) Validate() error {
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

type CreateAggregateConfigRuleRequestExcludeTagsScope struct {
	// The tag key of the resource that you want to exclude.
	//
	// example:
	//
	// key-2
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource that you want to exclude.
	//
	// example:
	//
	// value-2
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateAggregateConfigRuleRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAggregateConfigRuleRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAggregateConfigRuleRequestExcludeTagsScope) SetTagKey(v string) *CreateAggregateConfigRuleRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateAggregateConfigRuleRequestExcludeTagsScope) SetTagValue(v string) *CreateAggregateConfigRuleRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateAggregateConfigRuleRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type CreateAggregateConfigRuleRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
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

func (s CreateAggregateConfigRuleRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAggregateConfigRuleRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAggregateConfigRuleRequestTag) SetKey(v string) *CreateAggregateConfigRuleRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAggregateConfigRuleRequestTag) SetValue(v string) *CreateAggregateConfigRuleRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAggregateConfigRuleRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateAggregateConfigRuleRequestTagsScope struct {
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

func (s CreateAggregateConfigRuleRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleRequestTagsScope) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAggregateConfigRuleRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAggregateConfigRuleRequestTagsScope) SetTagKey(v string) *CreateAggregateConfigRuleRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateAggregateConfigRuleRequestTagsScope) SetTagValue(v string) *CreateAggregateConfigRuleRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateAggregateConfigRuleRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
