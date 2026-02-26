// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateConfigRuleRequest
	GetClientToken() *string
	SetConditions(v string) *CreateConfigRuleRequest
	GetConditions() *string
	SetConfigRuleName(v string) *CreateConfigRuleRequest
	GetConfigRuleName() *string
	SetConfigRuleTriggerTypes(v string) *CreateConfigRuleRequest
	GetConfigRuleTriggerTypes() *string
	SetDescription(v string) *CreateConfigRuleRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *CreateConfigRuleRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateConfigRuleRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateConfigRuleRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*CreateConfigRuleRequestExcludeTagsScope) *CreateConfigRuleRequest
	GetExcludeTagsScope() []*CreateConfigRuleRequestExcludeTagsScope
	SetExtendContent(v string) *CreateConfigRuleRequest
	GetExtendContent() *string
	SetInputParameters(v map[string]interface{}) *CreateConfigRuleRequest
	GetInputParameters() map[string]interface{}
	SetMaximumExecutionFrequency(v string) *CreateConfigRuleRequest
	GetMaximumExecutionFrequency() *string
	SetRegionIdsScope(v string) *CreateConfigRuleRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateConfigRuleRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateConfigRuleRequest
	GetResourceIdsScope() *string
	SetResourceNameScope(v string) *CreateConfigRuleRequest
	GetResourceNameScope() *string
	SetResourceTypesScope(v []*string) *CreateConfigRuleRequest
	GetResourceTypesScope() []*string
	SetRiskLevel(v int32) *CreateConfigRuleRequest
	GetRiskLevel() *int32
	SetSourceIdentifier(v string) *CreateConfigRuleRequest
	GetSourceIdentifier() *string
	SetSourceOwner(v string) *CreateConfigRuleRequest
	GetSourceOwner() *string
	SetTag(v []*CreateConfigRuleRequestTag) *CreateConfigRuleRequest
	GetTag() []*CreateConfigRuleRequestTag
	SetTagKeyLogicScope(v string) *CreateConfigRuleRequest
	GetTagKeyLogicScope() *string
	SetTagKeyScope(v string) *CreateConfigRuleRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *CreateConfigRuleRequest
	GetTagValueScope() *string
	SetTagsScope(v []*CreateConfigRuleRequestTagsScope) *CreateConfigRuleRequest
	GetTagsScope() []*CreateConfigRuleRequestTagsScope
}

type CreateConfigRuleRequest struct {
	// A client token used to ensure request idempotence. Generate a unique token on your client. The `ClientToken` parameter can contain only ASCII characters and cannot exceed 64 characters.
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
	// The trigger that invokes the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule runs when a resource configuration changes.
	//
	// - ScheduledNotification: The rule runs on a regular schedule.
	//
	// > If a rule has multiple triggers, separate them with commas (,).
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
	// The rule does not apply to resources in the specified regions. The compliance of resources in these regions is not evaluated. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The rule does not apply to resources in the specified resource groups. The compliance of resources in these resource groups is not evaluated. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The rule does not apply to the specified resources. The compliance of these resources is not evaluated. Separate multiple resource IDs with commas (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The scope of the tags to exclude.
	ExcludeTagsScope []*CreateConfigRuleRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// Extended content. This parameter specifies the trigger time for a 24-hour evaluation cycle.
	//
	// example:
	//
	// {"fixedHour":"13"}
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The input parameters for the rule.
	//
	// You can get the input parameters of a rule by calling the [GetManagedRule](https://help.aliyun.com/document_detail/606993.html) operation. View the `CompulsoryInputParameterDetails` and `OptionalInputParameterDetails` parameters to learn about the required and optional parameters.
	//
	// The format of the input parameters is `{"Parameter 1 Name":"Parameter 1 Value","Parameter 2 Name":"Parameter 2 Value"}`.
	//
	// example:
	//
	// {"key1":"value1","key2":"value2"}
	InputParameters map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The frequency at which the rule runs. Valid values:
	//
	// - One_Hour: every hour.
	//
	// - Three_Hours: every three hours.
	//
	// - Six_Hours: every six hours.
	//
	// - Twelve_Hours: every twelve hours.
	//
	// - TwentyFour_Hours (default): every twenty-four hours.
	//
	// > This parameter is required if you set ConfigRuleTriggerTypes to ScheduledNotification.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The rule applies only to resources in the specified regions. Separate multiple region IDs with commas (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The rule applies only to resources in the specified resource groups. Separate multiple resource group IDs with commas (,).
	//
	// > This parameter applies only to rule templates.
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The rule applies to the specified resources. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// lb-5cmbowstbkss9ta03****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The rule applies only to resources that have the specified names.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// i-xxx
	ResourceNameScope *string `json:"ResourceNameScope,omitempty" xml:"ResourceNameScope,omitempty"`
	// The resource types to evaluate. Separate multiple resource types with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypesScope []*string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
	// The risk level of the rule. Valid values:
	//
	// - 1: high.
	//
	// - 2: medium.
	//
	// - 3: low.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The identifier of the rule.
	//
	// - If you set `SourceOwner` to `ALIYUN`, specify the identifier of the rule template. Example: `required-tags`.
	//
	//   > To query the identifier of a rule template, see [List of rule templates](https://help.aliyun.com/document_detail/127404.html).
	//
	// - If you set `SourceOwner` to `CUSTOM_CONFIGURATION`, set this parameter to `acs-config-configuration`.
	//
	// - If you set `SourceOwner` to `CUSTOM_FC`, specify the Alibaba Cloud Resource Name (ARN) of the function.
	//
	//   The ARN must be in the format `acs:fc:{Region}:{AccountID}:services/{ServiceName}.LATEST/functions/{FunctionName}`. Example: `acs:fc:cn-hangzhou:120886317861****:services/service-test.LATEST/functions/config-test`.
	//
	//   > To obtain the ARN of a function, see [ListFunctions](https://help.aliyun.com/document_detail/415752.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// required-tags
	SourceIdentifier *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	// The type of rule to create. Valid values:
	//
	// - ALIYUN: rule template.
	//
	// - CUSTOM_FC: custom Function Compute rule.
	//
	// - CUSTOM_CONFIGURATION: custom condition rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	SourceOwner *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	// The tags of the rule to create.
	Tag []*CreateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The logical operator used when you specify multiple tags for the `TagsScope` parameter. For example, if you set `TagsScope` to `"TagsScope.1.TagKey":"a","TagsScope.1.TagValue":"a","TagsScope.2.TagKey":"b","TagsScope.2.TagValue":"b"` and set this parameter to `AND`, the rule applies only to resources that have both the `a:a` and `b:b` tags. If you do not specify this parameter, the default value `OR` is used.
	//
	// This parameter also works with the deprecated `TagKeyScope` parameter. For example, if you set `TagKeyScope` to `ECS,OSS` and set this parameter to `AND`, the rule applies only to resources that have both the `ECS` and `OSS` tags.
	//
	// Valid values:
	//
	// - AND: Use AND logic.
	//
	// - OR: Use OR logic.
	//
	// example:
	//
	// AND
	TagKeyLogicScope *string `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter instead.
	//
	// The rule applies only to resources that have the specified tag key.
	//
	// > This parameter applies only to managed rules. You must set both `TagKeyScope` and `TagValueScope`.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter instead.
	//
	// The rule applies only to resources that have the specified tag value.
	//
	// > This parameter applies only to rule templates. You must set both `TagKeyScope` and `TagValueScope`.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The scope of the tags.
	TagsScope []*CreateConfigRuleRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s CreateConfigRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateConfigRuleRequest) GetConditions() *string {
	return s.Conditions
}

func (s *CreateConfigRuleRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *CreateConfigRuleRequest) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *CreateConfigRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConfigRuleRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateConfigRuleRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateConfigRuleRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateConfigRuleRequest) GetExcludeTagsScope() []*CreateConfigRuleRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *CreateConfigRuleRequest) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *CreateConfigRuleRequest) GetInputParameters() map[string]interface{} {
	return s.InputParameters
}

func (s *CreateConfigRuleRequest) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *CreateConfigRuleRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateConfigRuleRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateConfigRuleRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateConfigRuleRequest) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *CreateConfigRuleRequest) GetResourceTypesScope() []*string {
	return s.ResourceTypesScope
}

func (s *CreateConfigRuleRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateConfigRuleRequest) GetSourceIdentifier() *string {
	return s.SourceIdentifier
}

func (s *CreateConfigRuleRequest) GetSourceOwner() *string {
	return s.SourceOwner
}

func (s *CreateConfigRuleRequest) GetTag() []*CreateConfigRuleRequestTag {
	return s.Tag
}

func (s *CreateConfigRuleRequest) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *CreateConfigRuleRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *CreateConfigRuleRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *CreateConfigRuleRequest) GetTagsScope() []*CreateConfigRuleRequestTagsScope {
	return s.TagsScope
}

func (s *CreateConfigRuleRequest) SetClientToken(v string) *CreateConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigRuleRequest) SetConditions(v string) *CreateConfigRuleRequest {
	s.Conditions = &v
	return s
}

func (s *CreateConfigRuleRequest) SetConfigRuleName(v string) *CreateConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateConfigRuleRequest) SetConfigRuleTriggerTypes(v string) *CreateConfigRuleRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *CreateConfigRuleRequest) SetDescription(v string) *CreateConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigRuleRequest) SetExcludeRegionIdsScope(v string) *CreateConfigRuleRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetExcludeResourceGroupIdsScope(v string) *CreateConfigRuleRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetExcludeResourceIdsScope(v string) *CreateConfigRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetExcludeTagsScope(v []*CreateConfigRuleRequestExcludeTagsScope) *CreateConfigRuleRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *CreateConfigRuleRequest) SetExtendContent(v string) *CreateConfigRuleRequest {
	s.ExtendContent = &v
	return s
}

func (s *CreateConfigRuleRequest) SetInputParameters(v map[string]interface{}) *CreateConfigRuleRequest {
	s.InputParameters = v
	return s
}

func (s *CreateConfigRuleRequest) SetMaximumExecutionFrequency(v string) *CreateConfigRuleRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *CreateConfigRuleRequest) SetRegionIdsScope(v string) *CreateConfigRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetResourceGroupIdsScope(v string) *CreateConfigRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetResourceIdsScope(v string) *CreateConfigRuleRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetResourceNameScope(v string) *CreateConfigRuleRequest {
	s.ResourceNameScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetResourceTypesScope(v []*string) *CreateConfigRuleRequest {
	s.ResourceTypesScope = v
	return s
}

func (s *CreateConfigRuleRequest) SetRiskLevel(v int32) *CreateConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateConfigRuleRequest) SetSourceIdentifier(v string) *CreateConfigRuleRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *CreateConfigRuleRequest) SetSourceOwner(v string) *CreateConfigRuleRequest {
	s.SourceOwner = &v
	return s
}

func (s *CreateConfigRuleRequest) SetTag(v []*CreateConfigRuleRequestTag) *CreateConfigRuleRequest {
	s.Tag = v
	return s
}

func (s *CreateConfigRuleRequest) SetTagKeyLogicScope(v string) *CreateConfigRuleRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetTagKeyScope(v string) *CreateConfigRuleRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetTagValueScope(v string) *CreateConfigRuleRequest {
	s.TagValueScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetTagsScope(v []*CreateConfigRuleRequestTagsScope) *CreateConfigRuleRequest {
	s.TagsScope = v
	return s
}

func (s *CreateConfigRuleRequest) Validate() error {
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

type CreateConfigRuleRequestExcludeTagsScope struct {
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

func (s CreateConfigRuleRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateConfigRuleRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateConfigRuleRequestExcludeTagsScope) SetTagKey(v string) *CreateConfigRuleRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateConfigRuleRequestExcludeTagsScope) SetTagValue(v string) *CreateConfigRuleRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateConfigRuleRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type CreateConfigRuleRequestTag struct {
	// The tag key of the resource.
	//
	// You can attach up to 20 tag keys.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// You can attach up to 20 tag values.
	//
	// example:
	//
	// value-1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConfigRuleRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleRequestTag) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateConfigRuleRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateConfigRuleRequestTag) SetKey(v string) *CreateConfigRuleRequestTag {
	s.Key = &v
	return s
}

func (s *CreateConfigRuleRequestTag) SetValue(v string) *CreateConfigRuleRequestTag {
	s.Value = &v
	return s
}

func (s *CreateConfigRuleRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateConfigRuleRequestTagsScope struct {
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

func (s CreateConfigRuleRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleRequestTagsScope) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateConfigRuleRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateConfigRuleRequestTagsScope) SetTagKey(v string) *CreateConfigRuleRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateConfigRuleRequestTagsScope) SetTagValue(v string) *CreateConfigRuleRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateConfigRuleRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
