// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateConfigRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetAccountIdsScope() *string
	SetAggregatorId(v string) *CreateAggregateConfigRuleShrinkRequest
	GetAggregatorId() *string
	SetClientToken(v string) *CreateAggregateConfigRuleShrinkRequest
	GetClientToken() *string
	SetConfigRuleName(v string) *CreateAggregateConfigRuleShrinkRequest
	GetConfigRuleName() *string
	SetConfigRuleTriggerTypes(v string) *CreateAggregateConfigRuleShrinkRequest
	GetConfigRuleTriggerTypes() *string
	SetDescription(v string) *CreateAggregateConfigRuleShrinkRequest
	GetDescription() *string
	SetExcludeAccountIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetExcludeAccountIdsScope() *string
	SetExcludeFolderIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetExcludeFolderIdsScope() *string
	SetExcludeRegionIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) *CreateAggregateConfigRuleShrinkRequest
	GetExcludeTagsScope() []*CreateAggregateConfigRuleShrinkRequestExcludeTagsScope
	SetExtendContent(v string) *CreateAggregateConfigRuleShrinkRequest
	GetExtendContent() *string
	SetFolderIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetFolderIdsScope() *string
	SetInputParametersShrink(v string) *CreateAggregateConfigRuleShrinkRequest
	GetInputParametersShrink() *string
	SetMaximumExecutionFrequency(v string) *CreateAggregateConfigRuleShrinkRequest
	GetMaximumExecutionFrequency() *string
	SetRegionIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetResourceIdsScope() *string
	SetResourceNameScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetResourceNameScope() *string
	SetResourceTypesScopeShrink(v string) *CreateAggregateConfigRuleShrinkRequest
	GetResourceTypesScopeShrink() *string
	SetRiskLevel(v int32) *CreateAggregateConfigRuleShrinkRequest
	GetRiskLevel() *int32
	SetSourceIdentifier(v string) *CreateAggregateConfigRuleShrinkRequest
	GetSourceIdentifier() *string
	SetSourceOwner(v string) *CreateAggregateConfigRuleShrinkRequest
	GetSourceOwner() *string
	SetTagShrink(v string) *CreateAggregateConfigRuleShrinkRequest
	GetTagShrink() *string
	SetTagKeyLogicScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetTagKeyLogicScope() *string
	SetTagKeyScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *CreateAggregateConfigRuleShrinkRequest
	GetTagValueScope() *string
	SetTagsScope(v []*CreateAggregateConfigRuleShrinkRequestTagsScope) *CreateAggregateConfigRuleShrinkRequest
	GetTagsScope() []*CreateAggregateConfigRuleShrinkRequestTagsScope
}

type CreateAggregateConfigRuleShrinkRequest struct {
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
	ExcludeTagsScope []*CreateAggregateConfigRuleShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	ResourceTypesScopeShrink *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
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
	TagsScope []*CreateAggregateConfigRuleShrinkRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s CreateAggregateConfigRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetAccountIdsScope() *string {
	return s.AccountIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetExcludeAccountIdsScope() *string {
	return s.ExcludeAccountIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetExcludeFolderIdsScope() *string {
	return s.ExcludeFolderIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetExcludeTagsScope() []*CreateAggregateConfigRuleShrinkRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetFolderIdsScope() *string {
	return s.FolderIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetInputParametersShrink() *string {
	return s.InputParametersShrink
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetResourceTypesScopeShrink() *string {
	return s.ResourceTypesScopeShrink
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetSourceIdentifier() *string {
	return s.SourceIdentifier
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetSourceOwner() *string {
	return s.SourceOwner
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) GetTagsScope() []*CreateAggregateConfigRuleShrinkRequestTagsScope {
	return s.TagsScope
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetAccountIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.AccountIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetAggregatorId(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetClientToken(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetConfigRuleName(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetConfigRuleTriggerTypes(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetDescription(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeAccountIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeFolderIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeRegionIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeResourceGroupIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeResourceIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeTagsScope(v []*CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExtendContent(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExtendContent = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetFolderIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.FolderIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetInputParametersShrink(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.InputParametersShrink = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetMaximumExecutionFrequency(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetRegionIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetResourceGroupIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetResourceIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetResourceNameScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ResourceNameScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetResourceTypesScopeShrink(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ResourceTypesScopeShrink = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetRiskLevel(v int32) *CreateAggregateConfigRuleShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetSourceIdentifier(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetSourceOwner(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.SourceOwner = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetTagShrink(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetTagKeyLogicScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetTagKeyScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetTagValueScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.TagValueScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetTagsScope(v []*CreateAggregateConfigRuleShrinkRequestTagsScope) *CreateAggregateConfigRuleShrinkRequest {
	s.TagsScope = v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) Validate() error {
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

type CreateAggregateConfigRuleShrinkRequestExcludeTagsScope struct {
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

func (s CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) SetTagKey(v string) *CreateAggregateConfigRuleShrinkRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) SetTagValue(v string) *CreateAggregateConfigRuleShrinkRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type CreateAggregateConfigRuleShrinkRequestTagsScope struct {
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

func (s CreateAggregateConfigRuleShrinkRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigRuleShrinkRequestTagsScope) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleShrinkRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateAggregateConfigRuleShrinkRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateAggregateConfigRuleShrinkRequestTagsScope) SetTagKey(v string) *CreateAggregateConfigRuleShrinkRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequestTagsScope) SetTagValue(v string) *CreateAggregateConfigRuleShrinkRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
