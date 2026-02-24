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
	SetConditions(v string) *CreateAggregateConfigRuleShrinkRequest
	GetConditions() *string
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
	// The rule is effective only for resources of the specified member accounts. Separate multiple member account IDs with commas (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// 115748125982****
	AccountIdsScope *string `json:"AccountIdsScope,omitempty" xml:"AccountIdsScope,omitempty"`
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// A client token to ensure that the request is idempotent. Generate a unique value from your client for each request. The `ClientToken` parameter must contain only ASCII characters and be no more than 64 characters long.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// {"ComplianceConditions":"{\"operator\":\"and\",\"children\":[{\"operator\":\"StringEquals\",\"featurePath\":\"$.Status\",\"desired\":\"1\",\"featureSource\":\"CONFIGURATION\"}]}"}
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
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
	// 最多可以定义6组标签。如果资源同时具有指定的所有标签，则视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule is not effective for resources of the specified member accounts. The resources of the specified member accounts are not evaluated. Separate multiple member account IDs with commas (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// 120886317861****
	ExcludeAccountIdsScope *string `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	// The rule is not effective for resources of the member accounts in the specified folders. The resources of the member accounts in the specified folders are not evaluated. Separate multiple folder IDs with commas (,).
	//
	// > - This parameter applies only to rules of a global account group.
	//
	// >
	//
	// > - This parameter applies only to rule templates.
	//
	// example:
	//
	// fd-pWmkqZ****
	ExcludeFolderIdsScope *string `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	// The rule is not effective for resources in the specified regions. The resources in the specified regions are not evaluated. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The rule is not effective for resources in the specified resource groups. The resources in the specified resource groups are not evaluated. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The rule is not effective for the specified resources. The specified resources are not evaluated. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The scope of the tags to be excluded.
	ExcludeTagsScope []*CreateAggregateConfigRuleShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The extended content. This parameter specifies the trigger time for a rule that runs on a 24-hour cycle.
	//
	// example:
	//
	// {"fixedHour":"12"}
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The rule is effective only for resources of the member accounts in the specified folders. Separate multiple folder IDs with commas (,).
	//
	// > - This parameter applies only to rules of a global account group.
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
	// - TwentyFour_Hours (default): 24 hours.
	//
	// > This parameter is required if you set `ConfigRuleTriggerTypes` to `ScheduledNotification`.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The rule is effective only for resources in the specified regions. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The rule is effective only for resources in the specified resource groups. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The rule is effective only for the specified resources. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The rule is effective only for resources that have the specified names.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// i-xxx
	ResourceNameScope *string `json:"ResourceNameScope,omitempty" xml:"ResourceNameScope,omitempty"`
	// The resource types to be evaluated by the rule. Separate multiple resource types with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypesScopeShrink *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - 1: high
	//
	// - 2: medium
	//
	// - 3: low
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The identifier of the rule.
	//
	// - If you set `SourceOwner` to `ALIYUN`, enter the identifier of the rule template, such as `required-tags`.
	//
	//   > For more information about how to query the identifier of a rule template, see [List of rule templates](https://help.aliyun.com/document_detail/127404.html).
	//
	// - If you set `SourceOwner` to `CUSTOM_FC`, enter the Alibaba Cloud Resource Name (ARN) of the function in Function Compute.
	//
	//   The ARN is in the format of `acs:fc:{region}:{accountId}:services/{serviceName}.LATEST/functions/{functionName}`. For example, `acs:fc:cn-hangzhou:120886317861****:services/service-test.LATEST/functions/config-test`.
	//
	//   > For more information about how to obtain the ARN of a function, see [ListFunctions](https://help.aliyun.com/document_detail/415752.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// required-tags
	SourceIdentifier *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	// The type of the rule. Valid values:
	//
	// - ALIYUN: rule template
	//
	// - CUSTOM_FC: custom rule
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	SourceOwner *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	// The tags to add to the rule. You can add up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The logical relationship for multiple tags in the `TagsScope` parameter. For example, if you set the `TagsScope` parameter to `"TagsScope.1.TagKey":"a","TagsScope.1.TagValue":"a","TagsScope.2.TagKey":"b","TagsScope.2.TagValue":"b"` and set this parameter to `AND`, the rule applies only to resources that have both the `a:a` and `b:b` tags. The default value is `OR`.
	//
	// This parameter can also be used for the deprecated `TagKeyScope` parameter, but this is not recommended. For example, if you set `TagKeyScope` to `ECS,OSS` and set this parameter to `AND`, the rule applies only to resources that have both the `ECS` and `OSS` tags.
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
	// This parameter is deprecated. Use the `TagsScope` parameter.
	//
	// The rule is effective only for resources that have the specified tag keys. Separate multiple tag keys with commas (,).
	//
	// > This parameter applies only to rule templates. The `TagKeyScope` and `TagValueScope` parameters must be used together.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter.
	//
	// The rule is effective only for resources that have the specified tag values.
	//
	// > This parameter applies only to rule templates. The `TagKeyScope` and `TagValueScope` parameters must be used together.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The scope of the tags.
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

func (s *CreateAggregateConfigRuleShrinkRequest) GetConditions() *string {
	return s.Conditions
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

func (s *CreateAggregateConfigRuleShrinkRequest) SetConditions(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.Conditions = &v
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
	// The tag key of the resource to be excluded.
	//
	// example:
	//
	// key-2
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource to be excluded.
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
	// The tag key of the resource.
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource.
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
