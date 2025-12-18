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
	// The client token that you want to use to ensure the idempotency of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.``
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// required-tags
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// 	- ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// 	- ScheduledNotification: The rule is periodically triggered.
	//
	// >  If a rule supports the preceding trigger types, separate the types with a comma (,).
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
	// example-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ExcludeRegionIdsScope
	//
	// example:
	//
	// cn-shanghai
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// ExcludeResourceGroupIdsScope
	//
	// example:
	//
	// rg-bnczc6r7rml****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The ID of the resource to be excluded from the compliance evaluations performed by the rule. Separate multiple resource IDs with commas (,).
	//
	// >  This parameter applies only to managed rules.
	//
	// example:
	//
	// lb-t4nbowvtbkss7t326****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// ExcludeTagsScope
	ExcludeTagsScope []*CreateConfigRuleRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// Optional field, only used in conjunction with the 24-hour cycle execution to set the trigger time.
	//
	// example:
	//
	// {"fixedHour":"13"}
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The input parameter of the rule.
	//
	// example:
	//
	// {"tag1Key":"ECS","tag1Value":"test"}
	InputParameters map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The intervals at which the rule is triggered. Valid values:
	//
	// 	- One_Hour: 1 hour.
	//
	// 	- Three_Hours: 3 hours.
	//
	// 	- Six_Hours: 6 hours.
	//
	// 	- Twelve_Hours: 12 hours.
	//
	// 	- TwentyFour_Hours (default): 24 hours.
	//
	// >  This parameter is required if the ConfigRuleTriggerTypes parameter is set to ScheduledNotification.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The ID of the region to which the rule applies. Separate multiple region IDs with commas (,).
	//
	// >  This parameter applies only to managed rules.
	//
	// example:
	//
	// cn-hangzhou
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The ID of the resource group to which the rule applies. Separate multiple resource group IDs with commas (,).
	//
	// >  This parameter applies only to managed rules.
	//
	// example:
	//
	// rg-aekzc7r7rhx****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// ResourceIdsScope
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
	// This parameter is required.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceTypesScope []*string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
	// The risk level of the resources that do not comply with the rule. Valid values:
	//
	// 	- 1: high.
	//
	// 	- 2: medium.
	//
	// 	- 3: low.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The ID of the rule.
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
	// The type of the rule Valid values:
	//
	// 	- ALIYUN: managed rule.
	//
	// 	- CUSTOM_FC: custom rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	SourceOwner *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	// rule attached tags
	Tag []*CreateConfigRuleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The tag key used to filter resources. The rule applies only to the resources with the specified tag key.
	//
	// >  This parameter applies only to managed rules. You must specify both `TagKeyScope` and `TagValueScope`.
	//
	// example:
	//
	// ECS
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// The tag value used to filter resources. The rule applies only to the resources that use the specified tag value.
	//
	// >  This parameter applies only to managed rules. You must specify both `TagKeyScope` and `TagValueScope`.
	//
	// example:
	//
	// test
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// TagsScope
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
	// TagKey
	//
	// example:
	//
	// key-2
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// TagValue
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
	// The tag keys.
	//
	// The tag keys cannot be an empty string. The tag keys can be up to 64 characters in length. The tag keys cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
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
	// TagKey
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// TagValue
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
