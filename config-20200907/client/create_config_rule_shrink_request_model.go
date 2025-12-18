// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateConfigRuleShrinkRequest
	GetClientToken() *string
	SetConfigRuleName(v string) *CreateConfigRuleShrinkRequest
	GetConfigRuleName() *string
	SetConfigRuleTriggerTypes(v string) *CreateConfigRuleShrinkRequest
	GetConfigRuleTriggerTypes() *string
	SetDescription(v string) *CreateConfigRuleShrinkRequest
	GetDescription() *string
	SetExcludeRegionIdsScope(v string) *CreateConfigRuleShrinkRequest
	GetExcludeRegionIdsScope() *string
	SetExcludeResourceGroupIdsScope(v string) *CreateConfigRuleShrinkRequest
	GetExcludeResourceGroupIdsScope() *string
	SetExcludeResourceIdsScope(v string) *CreateConfigRuleShrinkRequest
	GetExcludeResourceIdsScope() *string
	SetExcludeTagsScope(v []*CreateConfigRuleShrinkRequestExcludeTagsScope) *CreateConfigRuleShrinkRequest
	GetExcludeTagsScope() []*CreateConfigRuleShrinkRequestExcludeTagsScope
	SetExtendContent(v string) *CreateConfigRuleShrinkRequest
	GetExtendContent() *string
	SetInputParametersShrink(v string) *CreateConfigRuleShrinkRequest
	GetInputParametersShrink() *string
	SetMaximumExecutionFrequency(v string) *CreateConfigRuleShrinkRequest
	GetMaximumExecutionFrequency() *string
	SetRegionIdsScope(v string) *CreateConfigRuleShrinkRequest
	GetRegionIdsScope() *string
	SetResourceGroupIdsScope(v string) *CreateConfigRuleShrinkRequest
	GetResourceGroupIdsScope() *string
	SetResourceIdsScope(v string) *CreateConfigRuleShrinkRequest
	GetResourceIdsScope() *string
	SetResourceNameScope(v string) *CreateConfigRuleShrinkRequest
	GetResourceNameScope() *string
	SetResourceTypesScopeShrink(v string) *CreateConfigRuleShrinkRequest
	GetResourceTypesScopeShrink() *string
	SetRiskLevel(v int32) *CreateConfigRuleShrinkRequest
	GetRiskLevel() *int32
	SetSourceIdentifier(v string) *CreateConfigRuleShrinkRequest
	GetSourceIdentifier() *string
	SetSourceOwner(v string) *CreateConfigRuleShrinkRequest
	GetSourceOwner() *string
	SetTagShrink(v string) *CreateConfigRuleShrinkRequest
	GetTagShrink() *string
	SetTagKeyLogicScope(v string) *CreateConfigRuleShrinkRequest
	GetTagKeyLogicScope() *string
	SetTagKeyScope(v string) *CreateConfigRuleShrinkRequest
	GetTagKeyScope() *string
	SetTagValueScope(v string) *CreateConfigRuleShrinkRequest
	GetTagValueScope() *string
	SetTagsScope(v []*CreateConfigRuleShrinkRequestTagsScope) *CreateConfigRuleShrinkRequest
	GetTagsScope() []*CreateConfigRuleShrinkRequestTagsScope
}

type CreateConfigRuleShrinkRequest struct {
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
	ExcludeTagsScope []*CreateConfigRuleShrinkRequestExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	InputParametersShrink *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
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
	ResourceTypesScopeShrink *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
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
	TagsScope []*CreateConfigRuleShrinkRequestTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s CreateConfigRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateConfigRuleShrinkRequest) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *CreateConfigRuleShrinkRequest) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *CreateConfigRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConfigRuleShrinkRequest) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *CreateConfigRuleShrinkRequest) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *CreateConfigRuleShrinkRequest) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *CreateConfigRuleShrinkRequest) GetExcludeTagsScope() []*CreateConfigRuleShrinkRequestExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *CreateConfigRuleShrinkRequest) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *CreateConfigRuleShrinkRequest) GetInputParametersShrink() *string {
	return s.InputParametersShrink
}

func (s *CreateConfigRuleShrinkRequest) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *CreateConfigRuleShrinkRequest) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *CreateConfigRuleShrinkRequest) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *CreateConfigRuleShrinkRequest) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *CreateConfigRuleShrinkRequest) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *CreateConfigRuleShrinkRequest) GetResourceTypesScopeShrink() *string {
	return s.ResourceTypesScopeShrink
}

func (s *CreateConfigRuleShrinkRequest) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *CreateConfigRuleShrinkRequest) GetSourceIdentifier() *string {
	return s.SourceIdentifier
}

func (s *CreateConfigRuleShrinkRequest) GetSourceOwner() *string {
	return s.SourceOwner
}

func (s *CreateConfigRuleShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateConfigRuleShrinkRequest) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *CreateConfigRuleShrinkRequest) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *CreateConfigRuleShrinkRequest) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *CreateConfigRuleShrinkRequest) GetTagsScope() []*CreateConfigRuleShrinkRequestTagsScope {
	return s.TagsScope
}

func (s *CreateConfigRuleShrinkRequest) SetClientToken(v string) *CreateConfigRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetConfigRuleName(v string) *CreateConfigRuleShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetConfigRuleTriggerTypes(v string) *CreateConfigRuleShrinkRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetDescription(v string) *CreateConfigRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetExcludeRegionIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetExcludeResourceGroupIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetExcludeResourceIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetExcludeTagsScope(v []*CreateConfigRuleShrinkRequestExcludeTagsScope) *CreateConfigRuleShrinkRequest {
	s.ExcludeTagsScope = v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetExtendContent(v string) *CreateConfigRuleShrinkRequest {
	s.ExtendContent = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetInputParametersShrink(v string) *CreateConfigRuleShrinkRequest {
	s.InputParametersShrink = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetMaximumExecutionFrequency(v string) *CreateConfigRuleShrinkRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetRegionIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetResourceGroupIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetResourceIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.ResourceIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetResourceNameScope(v string) *CreateConfigRuleShrinkRequest {
	s.ResourceNameScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetResourceTypesScopeShrink(v string) *CreateConfigRuleShrinkRequest {
	s.ResourceTypesScopeShrink = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetRiskLevel(v int32) *CreateConfigRuleShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetSourceIdentifier(v string) *CreateConfigRuleShrinkRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetSourceOwner(v string) *CreateConfigRuleShrinkRequest {
	s.SourceOwner = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetTagShrink(v string) *CreateConfigRuleShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetTagKeyLogicScope(v string) *CreateConfigRuleShrinkRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetTagKeyScope(v string) *CreateConfigRuleShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetTagValueScope(v string) *CreateConfigRuleShrinkRequest {
	s.TagValueScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetTagsScope(v []*CreateConfigRuleShrinkRequestTagsScope) *CreateConfigRuleShrinkRequest {
	s.TagsScope = v
	return s
}

func (s *CreateConfigRuleShrinkRequest) Validate() error {
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

type CreateConfigRuleShrinkRequestExcludeTagsScope struct {
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

func (s CreateConfigRuleShrinkRequestExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleShrinkRequestExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleShrinkRequestExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateConfigRuleShrinkRequestExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateConfigRuleShrinkRequestExcludeTagsScope) SetTagKey(v string) *CreateConfigRuleShrinkRequestExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateConfigRuleShrinkRequestExcludeTagsScope) SetTagValue(v string) *CreateConfigRuleShrinkRequestExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateConfigRuleShrinkRequestExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type CreateConfigRuleShrinkRequestTagsScope struct {
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

func (s CreateConfigRuleShrinkRequestTagsScope) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRuleShrinkRequestTagsScope) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleShrinkRequestTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateConfigRuleShrinkRequestTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateConfigRuleShrinkRequestTagsScope) SetTagKey(v string) *CreateConfigRuleShrinkRequestTagsScope {
	s.TagKey = &v
	return s
}

func (s *CreateConfigRuleShrinkRequestTagsScope) SetTagValue(v string) *CreateConfigRuleShrinkRequestTagsScope {
	s.TagValue = &v
	return s
}

func (s *CreateConfigRuleShrinkRequestTagsScope) Validate() error {
	return dara.Validate(s)
}
