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
	SetConditions(v string) *CreateAggregateConfigRuleRequest
	GetConditions() *string
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
	// The rule applies only to resources of the specified member accounts. Separate multiple member account IDs with commas (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// 115748125982****
	AccountIdsScope *string `json:"AccountIdsScope,omitempty" xml:"AccountIdsScope,omitempty"`
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [the referenced document](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// A client token that ensures the request is idempotent. Generate a unique value from your client for each request. The `ClientToken` parameter must contain only ASCII characters and be no more than 64 characters long.
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The conditions for a custom condition rule, in JSON format.
	//
	// example:
	//
	// {"ComplianceConditions":"{\\"operator\\":\\"and\\",\\"children\\":[{\\"operator\\":\\"StringEquals\\",\\"featurePath\\":\\"$.Status\\",\\"desired\\":\\"1\\",\\"featureSource\\":\\"CONFIGURATION\\"}]}"}
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 存在所有指定标签
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The trigger type for the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule triggers when a resource configuration changes.
	//
	// - ScheduledNotification: The rule triggers on a schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// A description of the rule.
	//
	// example:
	//
	// 最多可以定义6组标签。如果资源同时具有指定的所有标签，则视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule does not apply to resources of the specified member accounts. Resources in these accounts are not evaluated. Separate multiple member account IDs with commas (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// 120886317861****
	ExcludeAccountIdsScope *string `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	// The rule does not apply to resources of member accounts in the specified folders. Resources in these folders are not evaluated. Separate multiple folder IDs with commas (,).
	//
	// > - This parameter applies only to global account group rules.
	//
	// >
	//
	// > - This parameter applies only to rule templates.
	//
	// example:
	//
	// fd-pWmkqZ****
	ExcludeFolderIdsScope *string `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	// The rule does not apply to resources in the specified regions. Resources in these regions are not evaluated. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The rule does not apply to resources in the specified resource groups. Resources in these groups are not evaluated. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The rule does not apply to the specified resources. These resources are not evaluated. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The scope of tags to exclude.
	ExcludeTagsScope []*CreateAggregateConfigRuleRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// Extended content. This parameter currently supports only setting the trigger time for rules that run on a 24-hour cycle.
	//
	// example:
	//
	// {"fixedHour":"12"}
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The rule applies only to resources of member accounts in the specified folders. Separate multiple folder IDs with commas (,).
	//
	// > - This parameter applies only to global account group rules.
	//
	// >
	//
	// > - This parameter applies only to rule templates.
	//
	// example:
	//
	// fd-ZtHsRH****
	FolderIdsScope *string `json:"FolderIdsScope,omitempty" xml:"FolderIdsScope,omitempty"`
	// The input parameters for the rule.
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
	// > Set this parameter if you set `ConfigRuleTriggerTypes` to `ScheduledNotification`.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The rule applies only to resources in the specified regions. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The rule applies only to resources in the specified resource groups. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The rule applies only to the specified resources. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The rule applies only to resources with the specified names.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// i-xxx
	ResourceNameScope *string `json:"ResourceNameScope,omitempty" xml:"ResourceNameScope,omitempty"`
	// The resource types to evaluate. Separate multiple types with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypesScope []*string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
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
	// - If `SourceOwner` is `ALIYUN`, enter the identifier of the rule template, such as `required-tags`.
	//
	//   > For more information about how to query rule template identifiers, see [the referenced document](https://help.aliyun.com/document_detail/127404.html).
	//
	// - If `SourceOwner` is `CUSTOM_CONFIGURATION`, enter `acs-config-configuration`.
	//
	// - If `SourceOwner` is `CUSTOM_FC`, enter the Alibaba Cloud Resource Name (ARN) of the Function Compute function.
	//
	//   The ARN format is `acs:fc:{region}:{accountId}:services/{serviceName}.LATEST/functions/{functionName}`. For example, `acs:fc:cn-hangzhou:120886317861****:services/service-test.LATEST/functions/config-test`.
	//
	//   > For more information about how to obtain a function ARN, see [the referenced document](https://help.aliyun.com/document_detail/415752.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// required-tags
	SourceIdentifier *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	// The type of rule. Valid values:
	//
	// - ALIYUN: rule template
	//
	// - CUSTOM_FC: custom Function Compute rule
	//
	// - CUSTOM_CONFIGURATION: custom condition rule
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	SourceOwner *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	// The tag of the rule.
	Tag []*CreateAggregateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The logical relationship between multiple tags in the `TagsScope` parameter. For example, if you set `TagsScope` to `"TagsScope.1.TagKey":"a","TagsScope.1.TagValue":"a","TagsScope.2.TagKey":"b","TagsScope.2.TagValue":"b"` and set this parameter to `AND`, the rule applies only to resources that have both the `a:a` and `b:b` tags. The default value is `OR`.
	//
	// You can also use this parameter with the deprecated `TagKeyScope` parameter, but this is not recommended. For example, if you set `TagKeyScope` to `ECS,OSS` and set this parameter to `AND`, the rule applies only to resources that have both the `ECS` and `OSS` tags.
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
	// The rule applies only to resources that have the specified tag keys. Separate multiple tag keys with commas (,).
	//
	// > This parameter applies only to rule templates. Set both `TagKeyScope` and `TagValueScope` together.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter instead.
	//
	// The rule applies only to resources that have the specified tag values.
	//
	// > This parameter applies only to rule templates. Set both `TagKeyScope` and `TagValueScope` together.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The scope of tags to include.
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

func (s *CreateAggregateConfigRuleRequest) GetConditions() *string {
	return s.Conditions
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

func (s *CreateAggregateConfigRuleRequest) SetConditions(v string) *CreateAggregateConfigRuleRequest {
	s.Conditions = &v
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
	// The tag key of the resource to exclude.
	//
	// example:
	//
	// key-2
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource to exclude.
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
	// The key of the tag.
	//
	// You can add up to 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// You can add up to 20 tag values.
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
