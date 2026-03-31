// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRule(v *GetAggregateConfigRuleResponseBodyConfigRule) *GetAggregateConfigRuleResponseBody
	GetConfigRule() *GetAggregateConfigRuleResponseBodyConfigRule
	SetRequestId(v string) *GetAggregateConfigRuleResponseBody
	GetRequestId() *string
}

type GetAggregateConfigRuleResponseBody struct {
	// The information about the rules.
	ConfigRule *GetAggregateConfigRuleResponseBodyConfigRule `json:"ConfigRule,omitempty" xml:"ConfigRule,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 811234F4-C3AB-4D15-B90B-F55016D1B5AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBody) GetConfigRule() *GetAggregateConfigRuleResponseBodyConfigRule {
	return s.ConfigRule
}

func (s *GetAggregateConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateConfigRuleResponseBody) SetConfigRule(v *GetAggregateConfigRuleResponseBodyConfigRule) *GetAggregateConfigRuleResponseBody {
	s.ConfigRule = v
	return s
}

func (s *GetAggregateConfigRuleResponseBody) SetRequestId(v string) *GetAggregateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBody) Validate() error {
	if s.ConfigRule != nil {
		if err := s.ConfigRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateConfigRuleResponseBodyConfigRule struct {
	// The ID of the Alibaba Cloud account to which the rule belongs.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The IDs of the members to which the rule applies. Separate multiple member IDs with commas (,).
	//
	// example:
	//
	// 120886317861****
	AccountIdsScope *string `json:"AccountIdsScope,omitempty" xml:"AccountIdsScope,omitempty"`
	// The details of compliance evaluation results.
	Compliance *GetAggregateConfigRuleResponseBodyConfigRuleCompliance `json:"Compliance,omitempty" xml:"Compliance,omitempty" type:"Struct"`
	// The ARN of the managed rule.
	//
	// example:
	//
	// acs:config::100931896542****:rule/cr-7f7d626622af0041****
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The information about compliance evaluations performed by the rule.
	ConfigRuleEvaluationStatus *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus `json:"ConfigRuleEvaluationStatus,omitempty" xml:"ConfigRuleEvaluationStatus,omitempty" type:"Struct"`
	// The ID of the rule.
	//
	// example:
	//
	// cr-7f7d626622af0041****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the monitoring rule.
	//
	// example:
	//
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- ACTIVE: The rule is being used to monitor resource configurations.
	//
	// 	- DELETING: The rule is being deleted.
	//
	// 	- EVALUATING: The rule is triggered and is being used to monitor resource configurations.
	//
	// 	- INACTIVE: The rule is disabled and is no longer used to monitor resource configurations.
	//
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// 	- ConfigurationItemChangeNotification: The managed rule is triggered by configuration changes.
	//
	// 	- ScheduledNotification: The managed rule is periodically triggered.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// The information about the creation of the rule.
	CreateBy *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	// The timestamp when the rule was created. Unit: milliseconds.
	//
	// example:
	//
	// 1604684022000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the managed rule.
	//
	// example:
	//
	// The description of the managed rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the member account to which the rule does not apply, which means that the resources within the member account are not evaluated based on the rule.
	//
	// >  This parameter applies only to a managed rule.
	//
	// example:
	//
	// 120886317861****
	ExcludeAccountIdsScope *string `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	// The ID of the resource directory to which the rule does not apply, which means that the resources within member accounts in the resource directory are not evaluated based on the rule.
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
	// The IDs of the regions excluded from the compliance evaluations performed by the rule. Separate multiple region IDs with commas (,).
	//
	// example:
	//
	// cn-hangzhou
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The IDs of the resource groups excluded from the compliance evaluations performed by the rule. Separate multiple resource group IDs with commas (,).
	//
	// example:
	//
	// rg-aekzdibsjjc****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The ID of the resource excluded from the compliance evaluations performed by the rule.
	//
	// example:
	//
	// 23642660635687****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The scope of the tag that is excluded.
	ExcludeTagsScope []*GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
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
	// 	- This parameter applies only to rules of a global account group.
	//
	// 	- This parameter applies only to managed rules.
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
	// The details of the managed rule.
	ManagedRule *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" type:"Struct"`
	// The intervals at which the managed rule is triggered. Valid values:
	//
	// 	- One_Hour: 1 hour.
	//
	// 	- Three_Hours: 3 hours.
	//
	// 	- Six_Hours: 6 hours.
	//
	// 	- Twelve_Hours: 12 hours
	//
	// 	- TwentyFour_Hours: 24 hours
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The timestamp when the rule was last updated. Unit: milliseconds.
	//
	// example:
	//
	// 1614687022000
	ModifiedTimestamp *int64 `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	// The ID of the region to which the rule applies.
	//
	// example:
	//
	// global
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The ID of the resource group to which the rule applies.
	//
	// example:
	//
	// rg-aekzdibsjjc****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The IDs of the resources to which the rule applies. Separate multiple resource IDs with commas (,).
	//
	// example:
	//
	// eip-8vbf3x310fn56ijfd****
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
	// The type of the resource evaluated by the rule.
	//
	// example:
	//
	// ACS::RAM::User
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The risk level of the resources that are not compliant with the rule. Valid values:
	//
	// 	- 1: high risk level
	//
	// 	- 2: medium risk level
	//
	// 	- 3: low risk level
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The information about how the rule was created.
	Source *GetAggregateConfigRuleResponseBodyConfigRuleSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// When retrieving details of rules created using the parameter `TagsScope`, this field will not be returned.
	//
	// To retrieve rules created using the deprecated field `TagKeyScope` (not recommended): for example, when the parameter `TagKeyScope` has a value of ECS,OSS, if this parameter is set to `AND`, it means that the rule only applies to resources bound with both labels ECS and OSS.
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
	// example:
	//
	// RAM
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. We recommend that you use the `TagsScope` parameter.
	//
	// The tag value used to filter resources. The rule applies only to the resources with the specified tag value.
	//
	// example:
	//
	// MFA
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The list of tags.
	Tags []*GetAggregateConfigRuleResponseBodyConfigRuleTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tag scope.
	TagsScope []*GetAggregateConfigRuleResponseBodyConfigRuleTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRule) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRule) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetAccountIdsScope() *string {
	return s.AccountIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetCompliance() *GetAggregateConfigRuleResponseBodyConfigRuleCompliance {
	return s.Compliance
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetConfigRuleArn() *string {
	return s.ConfigRuleArn
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetConfigRuleEvaluationStatus() *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	return s.ConfigRuleEvaluationStatus
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetCreateBy() *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	return s.CreateBy
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetDescription() *string {
	return s.Description
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetExcludeAccountIdsScope() *string {
	return s.ExcludeAccountIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetExcludeFolderIdsScope() *string {
	return s.ExcludeFolderIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetExcludeTagsScope() []*GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetFolderIdsScope() *string {
	return s.FolderIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetInputParameters() map[string]interface{} {
	return s.InputParameters
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetManagedRule() *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	return s.ManagedRule
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetModifiedTimestamp() *int64 {
	return s.ModifiedTimestamp
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetSource() *GetAggregateConfigRuleResponseBodyConfigRuleSource {
	return s.Source
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetTags() []*GetAggregateConfigRuleResponseBodyConfigRuleTags {
	return s.Tags
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) GetTagsScope() []*GetAggregateConfigRuleResponseBodyConfigRuleTagsScope {
	return s.TagsScope
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetAccountId(v int64) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.AccountId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetAccountIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.AccountIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetCompliance(v *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.Compliance = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleArn(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleArn = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleEvaluationStatus(v *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleEvaluationStatus = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleId(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleName(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleState(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleState = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleTriggerTypes(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetCreateBy(v *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.CreateBy = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetCreateTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.CreateTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetDescription(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.Description = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeAccountIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeFolderIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeRegionIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeResourceGroupIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeResourceIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeTagsScope(v []*GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeTagsScope = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExtendContent(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExtendContent = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetFolderIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.FolderIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetInputParameters(v map[string]interface{}) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.InputParameters = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetManagedRule(v *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ManagedRule = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetMaximumExecutionFrequency(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetModifiedTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ModifiedTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetRegionIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.RegionIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetResourceGroupIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetResourceIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ResourceIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetResourceNameScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ResourceNameScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetResourceTypesScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ResourceTypesScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetRiskLevel(v int32) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.RiskLevel = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetSource(v *GetAggregateConfigRuleResponseBodyConfigRuleSource) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.Source = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetTagKeyLogicScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.TagKeyLogicScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetTagKeyScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.TagKeyScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetTagValueScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.TagValueScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetTags(v []*GetAggregateConfigRuleResponseBodyConfigRuleTags) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.Tags = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetTagsScope(v []*GetAggregateConfigRuleResponseBodyConfigRuleTagsScope) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.TagsScope = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) Validate() error {
	if s.Compliance != nil {
		if err := s.Compliance.Validate(); err != nil {
			return err
		}
	}
	if s.ConfigRuleEvaluationStatus != nil {
		if err := s.ConfigRuleEvaluationStatus.Validate(); err != nil {
			return err
		}
	}
	if s.CreateBy != nil {
		if err := s.CreateBy.Validate(); err != nil {
			return err
		}
	}
	if s.ExcludeTagsScope != nil {
		for _, item := range s.ExcludeTagsScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ManagedRule != nil {
		if err := s.ManagedRule.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
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

type GetAggregateConfigRuleResponseBodyConfigRuleCompliance struct {
	// The statistics on the compliance evaluation results by compliance type. Valid values:
	//
	// 	- COMPLIANT: The resources are evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resources are evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to your resources.
	//
	// 	- INSUFFICIENT_DATA: No resource data is available.
	//
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The number of evaluated resources.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleCompliance) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleCompliance) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) GetCount() *int32 {
	return s.Count
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) SetComplianceType(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCompliance {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) SetCount(v int32) *GetAggregateConfigRuleResponseBodyConfigRuleCompliance {
	s.Count = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) Validate() error {
	return dara.Validate(s)
}

type GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus struct {
	// The timestamp when the rule was first triggered.
	//
	// example:
	//
	// 1624932221993
	FirstActivatedTimestamp *int64 `json:"FirstActivatedTimestamp,omitempty" xml:"FirstActivatedTimestamp,omitempty"`
	// Indicates whether resources were evaluated based on the rule. Valid values:
	//
	// 	- true: Resources were evaluated based on the rule.
	//
	// 	- false: Resources were not evaluated based on the rule.
	//
	// example:
	//
	// true
	FirstEvaluationStarted *bool `json:"FirstEvaluationStarted,omitempty" xml:"FirstEvaluationStarted,omitempty"`
	// The error code returned for the last failed compliance evaluation.
	//
	// example:
	//
	// TimeOut
	LastErrorCode *string `json:"LastErrorCode,omitempty" xml:"LastErrorCode,omitempty"`
	// The error message returned for the last failed compliance evaluation.
	//
	// example:
	//
	// time out
	LastErrorMessage *string `json:"LastErrorMessage,omitempty" xml:"LastErrorMessage,omitempty"`
	// The timestamp when the last failed compliance evaluation of the rule ended. Unit: milliseconds.
	//
	// example:
	//
	// 1614687022000
	LastFailedEvaluationTimestamp *int64 `json:"LastFailedEvaluationTimestamp,omitempty" xml:"LastFailedEvaluationTimestamp,omitempty"`
	// The timestamp when the last failed compliance evaluation of the rule started. Unit: milliseconds.
	//
	// example:
	//
	// 1614687022000
	LastFailedInvocationTimestamp *int64 `json:"LastFailedInvocationTimestamp,omitempty" xml:"LastFailedInvocationTimestamp,omitempty"`
	// The timestamp when the last successful compliance evaluation of the rule ended. Unit: milliseconds.
	//
	// example:
	//
	// 1624932227486
	LastSuccessfulEvaluationTimestamp *int64 `json:"LastSuccessfulEvaluationTimestamp,omitempty" xml:"LastSuccessfulEvaluationTimestamp,omitempty"`
	// The timestamp when the last successful compliance evaluation of the rule started. Unit: milliseconds.
	//
	// example:
	//
	// 1624932227476
	LastSuccessfulInvocationTimestamp *int64 `json:"LastSuccessfulInvocationTimestamp,omitempty" xml:"LastSuccessfulInvocationTimestamp,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetFirstActivatedTimestamp() *int64 {
	return s.FirstActivatedTimestamp
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetFirstEvaluationStarted() *bool {
	return s.FirstEvaluationStarted
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastErrorCode() *string {
	return s.LastErrorCode
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastErrorMessage() *string {
	return s.LastErrorMessage
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastFailedEvaluationTimestamp() *int64 {
	return s.LastFailedEvaluationTimestamp
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastFailedInvocationTimestamp() *int64 {
	return s.LastFailedInvocationTimestamp
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastSuccessfulEvaluationTimestamp() *int64 {
	return s.LastSuccessfulEvaluationTimestamp
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastSuccessfulInvocationTimestamp() *int64 {
	return s.LastSuccessfulInvocationTimestamp
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstActivatedTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstActivatedTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstEvaluationStarted(v bool) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstEvaluationStarted = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorCode(v string) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorCode = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorMessage(v string) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorMessage = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedEvaluationTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedEvaluationTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedInvocationTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedInvocationTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulEvaluationTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulEvaluationTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulInvocationTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulInvocationTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) Validate() error {
	return dara.Validate(s)
}

type GetAggregateConfigRuleResponseBodyConfigRuleCreateBy struct {
	// The ID of the account group.
	//
	// example:
	//
	// ca-04b3fd170e340007****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The name of the account group.
	//
	// example:
	//
	// Test_Group
	AggregatorName *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	// The ID of the compliance package.
	//
	// example:
	//
	// cp-541e626622af008****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance package.
	//
	// example:
	//
	// The name of the compliance package.
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the account that was used to create the rule.
	//
	// example:
	//
	// 100931896542****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the account that was used to create the rule.
	//
	// example:
	//
	// Alice
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The type of the entity to which the rule belongs. The value is fixed to `AGGREGATOR`, which indicates an account group.
	//
	// example:
	//
	// AGGREGATOR
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GetAggregatorName() *string {
	return s.AggregatorName
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GetCreatorType() *string {
	return s.CreatorType
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetAggregatorId(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetAggregatorName(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.AggregatorName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCompliancePackId(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCompliancePackName(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CompliancePackName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorId(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorName(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorType(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorType = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) Validate() error {
	return dara.Validate(s)
}

type GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope struct {
	// The key of the tag.
	//
	// example:
	//
	// key-2
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value-2
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope) SetTagKey(v string) *GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope) SetTagValue(v string) *GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type GetAggregateConfigRuleResponseBodyConfigRuleManagedRule struct {
	// The required input parameters of the managed rule.
	//
	// example:
	//
	// {}
	CompulsoryInputParameterDetails map[string]interface{} `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	// The description of the managed rule.
	//
	// example:
	//
	// The description of the managed rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the managed rule.
	//
	// example:
	//
	// ram-user-mfa-check
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The tags of the managed rule.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the managed rule.
	//
	// example:
	//
	// ram-user-mfa-check
	ManagedRuleName *string `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty"`
	// The optional input parameters of the managed rule.
	//
	// example:
	//
	// {}
	OptionalInputParameterDetails map[string]interface{} `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
	// The details of the source of the managed rule.
	SourceDetails []*GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GetCompulsoryInputParameterDetails() map[string]interface{} {
	return s.CompulsoryInputParameterDetails
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GetDescription() *string {
	return s.Description
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GetLabels() []*string {
	return s.Labels
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GetManagedRuleName() *string {
	return s.ManagedRuleName
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GetOptionalInputParameterDetails() map[string]interface{} {
	return s.OptionalInputParameterDetails
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GetSourceDetails() []*GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	return s.SourceDetails
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetCompulsoryInputParameterDetails(v map[string]interface{}) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetDescription(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.Description = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetIdentifier(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.Identifier = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetLabels(v []*string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.Labels = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetManagedRuleName(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.ManagedRuleName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetOptionalInputParameterDetails(v map[string]interface{}) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetSourceDetails(v []*GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.SourceDetails = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) Validate() error {
	if s.SourceDetails != nil {
		for _, item := range s.SourceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails struct {
	// The event source of the managed rule.
	//
	// >  Only events related to Cloud Config are supported. The value is fixed to aliyun.config.
	//
	// example:
	//
	// aliyun.config
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The intervals at which the managed rule is triggered. Valid values:
	//
	// 	- One_Hour: 1 hour.
	//
	// 	- Three_Hours: 3 hours.
	//
	// 	- Six_Hours: 6 hours.
	//
	// 	- Twelve_Hours: 12 hours
	//
	// 	- TwentyFour_Hours: 24 hours
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// 	- ConfigurationItemChangeNotification: The managed rule is triggered by configuration changes.
	//
	// 	- ScheduledNotification: The managed rule is periodically triggered.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GetEventSource() *string {
	return s.EventSource
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GetMessageType() *string {
	return s.MessageType
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetEventSource(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.EventSource = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMaximumExecutionFrequency(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMessageType(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MessageType = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) Validate() error {
	return dara.Validate(s)
}

type GetAggregateConfigRuleResponseBodyConfigRuleSource struct {
	// The identifier of the rule.
	//
	// 	- If the rule was created based on a managed rule, the value of this parameter is the name of the managed rule.
	//
	// 	- If the rule is a custom rule, the value of this parameter is the Alibaba Cloud Resource Name (ARN) of the relevant function in Function Compute.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:100931896542****:services/ConfigService.LATEST/functions/specific-config
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The way in which the rule was created. Valid values:
	//
	// 	- CUSTOM_FC: The rule is a custom rule.
	//
	// 	- ALIYUN: The rule was created based on a managed rule of Alibaba Cloud.
	//
	// example:
	//
	// ALIYUN
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The details of the source of the rule.
	SourceDetails []*GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleSource) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleSource) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) GetOwner() *string {
	return s.Owner
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) GetSourceDetails() []*GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	return s.SourceDetails
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) SetIdentifier(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSource {
	s.Identifier = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) SetOwner(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSource {
	s.Owner = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) SetSourceDetails(v []*GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) *GetAggregateConfigRuleResponseBodyConfigRuleSource {
	s.SourceDetails = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) Validate() error {
	if s.SourceDetails != nil {
		for _, item := range s.SourceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails struct {
	// The event source of the managed rule.
	//
	// >  Only events related to Cloud Config are supported. The value is fixed to aliyun.config.
	//
	// example:
	//
	// aliyun.config
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The intervals at which the managed rule is triggered. Valid values:
	//
	// 	- One_Hour: 1 hour.
	//
	// 	- Three_Hours: 3 hours.
	//
	// 	- Six_Hours: 6 hours.
	//
	// 	- Twelve_Hours: 12 hours
	//
	// 	- TwentyFour_Hours: 24 hours
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// 	- ConfigurationItemChangeNotification: The managed rule is triggered by configuration changes.
	//
	// 	- ScheduledNotification: The managed rule is periodically triggered.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) GetEventSource() *string {
	return s.EventSource
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) GetMessageType() *string {
	return s.MessageType
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetEventSource(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.EventSource = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMaximumExecutionFrequency(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMessageType(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MessageType = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) Validate() error {
	return dara.Validate(s)
}

type GetAggregateConfigRuleResponseBodyConfigRuleTags struct {
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

func (s GetAggregateConfigRuleResponseBodyConfigRuleTags) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleTags) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTags) SetTagKey(v string) *GetAggregateConfigRuleResponseBodyConfigRuleTags {
	s.TagKey = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTags) SetTagValue(v string) *GetAggregateConfigRuleResponseBodyConfigRuleTags {
	s.TagValue = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTags) Validate() error {
	return dara.Validate(s)
}

type GetAggregateConfigRuleResponseBodyConfigRuleTagsScope struct {
	// The key of the tag.
	//
	// example:
	//
	// key-1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value-1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleTagsScope) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleTagsScope) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTagsScope) SetTagKey(v string) *GetAggregateConfigRuleResponseBodyConfigRuleTagsScope {
	s.TagKey = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTagsScope) SetTagValue(v string) *GetAggregateConfigRuleResponseBodyConfigRuleTagsScope {
	s.TagValue = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleTagsScope) Validate() error {
	return dara.Validate(s)
}
