// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRule(v *GetConfigRuleResponseBodyConfigRule) *GetConfigRuleResponseBody
	GetConfigRule() *GetConfigRuleResponseBodyConfigRule
	SetRequestId(v string) *GetConfigRuleResponseBody
	GetRequestId() *string
}

type GetConfigRuleResponseBody struct {
	// The details of the rule.
	ConfigRule *GetConfigRuleResponseBodyConfigRule `json:"ConfigRule,omitempty" xml:"ConfigRule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 811234F4-C3AB-4D15-B90B-F55016D1B5AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBody) GetConfigRule() *GetConfigRuleResponseBodyConfigRule {
	return s.ConfigRule
}

func (s *GetConfigRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigRuleResponseBody) SetConfigRule(v *GetConfigRuleResponseBodyConfigRule) *GetConfigRuleResponseBody {
	s.ConfigRule = v
	return s
}

func (s *GetConfigRuleResponseBody) SetRequestId(v string) *GetConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigRuleResponseBody) Validate() error {
	if s.ConfigRule != nil {
		if err := s.ConfigRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConfigRuleResponseBodyConfigRule struct {
	// The ID of the Alibaba Cloud account to which the rule belongs.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The compliance statistics of the rule.
	Compliance *GetConfigRuleResponseBodyConfigRuleCompliance `json:"Compliance,omitempty" xml:"Compliance,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the rule.
	//
	// example:
	//
	// acs:config::100931896542****:rule/cr-7f7d626622af0041****
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The execution status of the rule.
	ConfigRuleEvaluationStatus *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus `json:"ConfigRuleEvaluationStatus,omitempty" xml:"ConfigRuleEvaluationStatus,omitempty" type:"Struct"`
	// The rule ID.
	//
	// example:
	//
	// cr-7f7d626622af0041****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// RAM用户开启MFA
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// - ACTIVE: The rule is enabled.
	//
	// - DELETING: The rule is being deleted.
	//
	// - EVALUATING: The rule is being used to evaluate resource configurations.
	//
	// - INACTIVE: The rule is disabled.
	//
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// - ScheduledNotification: The rule is triggered periodically.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	ConfigRuleTriggerTypes *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	// The information about the creator of the rule.
	CreateBy *GetConfigRuleResponseBodyConfigRuleCreateBy `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	// The timestamp when the rule was created. Unit: milliseconds.
	//
	// example:
	//
	// 1604684022000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// RAM用户开启MFA，视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the regions where the rule does not apply. The rule does not evaluate resources in these regions. Separate multiple region IDs with a comma (,).
	//
	// example:
	//
	// cn-hangzhou
	ExcludeRegionIdsScope *string `json:"ExcludeRegionIdsScope,omitempty" xml:"ExcludeRegionIdsScope,omitempty"`
	// The IDs of the resource groups where the rule does not apply. The rule does not evaluate resources in these resource groups. Separate multiple resource group IDs with a comma (,).
	//
	// example:
	//
	// rg-aekzdibsjjc****
	ExcludeResourceGroupIdsScope *string `json:"ExcludeResourceGroupIdsScope,omitempty" xml:"ExcludeResourceGroupIdsScope,omitempty"`
	// The IDs of the resources that are not evaluated by the rule. Separate multiple resource IDs with a comma (,).
	//
	// example:
	//
	// 23642660635687****
	ExcludeResourceIdsScope *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	// The tags of the resources that are not evaluated by the rule.
	ExcludeTagsScope []*GetConfigRuleResponseBodyConfigRuleExcludeTagsScope `json:"ExcludeTagsScope,omitempty" xml:"ExcludeTagsScope,omitempty" type:"Repeated"`
	// The extended content. This parameter is used only to specify the trigger time for a rule that is triggered on a 24-hour cycle.
	//
	// example:
	//
	// {"fixedHour":"12"}
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The input parameters of the rule.
	InputParameters map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The details of the managed rule.
	ManagedRule *GetConfigRuleResponseBodyConfigRuleManagedRule `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" type:"Struct"`
	// The execution frequency of the rule. Valid values:
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
	// > This parameter is returned only when the rule is triggered periodically.
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
	// The IDs of the regions where the rule applies. The rule evaluates only resources in these regions.
	//
	// example:
	//
	// global
	RegionIdsScope *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	// The IDs of the resource groups where the rule applies. The rule evaluates only resources in these resource groups.
	//
	// example:
	//
	// rg-aekzdibsjjc****
	ResourceGroupIdsScope *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	// The IDs of the resources that are evaluated by the rule. Separate multiple resource IDs with a comma (,).
	//
	// example:
	//
	// eip-8vbf3x310fn56ijfd****
	ResourceIdsScope *string `json:"ResourceIdsScope,omitempty" xml:"ResourceIdsScope,omitempty"`
	// The rule evaluates only resources that have the specified names.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// i-xxx
	ResourceNameScope *string `json:"ResourceNameScope,omitempty" xml:"ResourceNameScope,omitempty"`
	// The types of the resources that are evaluated by the rule.
	//
	// example:
	//
	// ACS::RAM::User
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	// The risk level of the rule. Valid values:
	//
	// - 1: high
	//
	// - 2: medium
	//
	// - 3: low
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The effective scope of the rule.
	Scope *GetConfigRuleResponseBodyConfigRuleScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// The source of the rule.
	Source *GetConfigRuleResponseBodyConfigRuleSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	// This parameter is not returned for rules that are created using the `TagsScope` parameter.
	//
	// This parameter is returned for rules that are created using the deprecated TagKeyScope parameter. We do not recommend that you use the `TagKeyScope` parameter. For example, if `TagKeyScope` is set to `ECS,OSS` and this parameter is set to `AND`, the rule applies only to resources that have both the `ECS` and `OSS` tags.
	//
	// Valid values:
	//
	// - AND
	//
	// - OR
	//
	// example:
	//
	// OR
	TagKeyLogicScope *string `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter instead.
	//
	// The rule applies only to resources with the specified tag.
	//
	// > The `TagKeyScope` and `TagValueScope` parameters are returned at the same time.
	//
	// example:
	//
	// RAM
	TagKeyScope *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use the `TagsScope` parameter instead.
	//
	// The rule applies only to resources with the specified tag.
	//
	// > The `TagKeyScope` and `TagValueScope` parameters are returned at the same time.
	//
	// example:
	//
	// MFA
	TagValueScope *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
	// The tags of the resource.
	Tags []*GetConfigRuleResponseBodyConfigRuleTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tag-based scope.
	TagsScope []*GetConfigRuleResponseBodyConfigRuleTagsScope `json:"TagsScope,omitempty" xml:"TagsScope,omitempty" type:"Repeated"`
}

func (s GetConfigRuleResponseBodyConfigRule) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRule) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRule) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetConfigRuleResponseBodyConfigRule) GetCompliance() *GetConfigRuleResponseBodyConfigRuleCompliance {
	return s.Compliance
}

func (s *GetConfigRuleResponseBodyConfigRule) GetConfigRuleArn() *string {
	return s.ConfigRuleArn
}

func (s *GetConfigRuleResponseBodyConfigRule) GetConfigRuleEvaluationStatus() *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	return s.ConfigRuleEvaluationStatus
}

func (s *GetConfigRuleResponseBodyConfigRule) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *GetConfigRuleResponseBodyConfigRule) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *GetConfigRuleResponseBodyConfigRule) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *GetConfigRuleResponseBodyConfigRule) GetConfigRuleTriggerTypes() *string {
	return s.ConfigRuleTriggerTypes
}

func (s *GetConfigRuleResponseBodyConfigRule) GetCreateBy() *GetConfigRuleResponseBodyConfigRuleCreateBy {
	return s.CreateBy
}

func (s *GetConfigRuleResponseBodyConfigRule) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetConfigRuleResponseBodyConfigRule) GetDescription() *string {
	return s.Description
}

func (s *GetConfigRuleResponseBodyConfigRule) GetExcludeRegionIdsScope() *string {
	return s.ExcludeRegionIdsScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetExcludeResourceGroupIdsScope() *string {
	return s.ExcludeResourceGroupIdsScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetExcludeResourceIdsScope() *string {
	return s.ExcludeResourceIdsScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetExcludeTagsScope() []*GetConfigRuleResponseBodyConfigRuleExcludeTagsScope {
	return s.ExcludeTagsScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *GetConfigRuleResponseBodyConfigRule) GetInputParameters() map[string]interface{} {
	return s.InputParameters
}

func (s *GetConfigRuleResponseBodyConfigRule) GetManagedRule() *GetConfigRuleResponseBodyConfigRuleManagedRule {
	return s.ManagedRule
}

func (s *GetConfigRuleResponseBodyConfigRule) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *GetConfigRuleResponseBodyConfigRule) GetModifiedTimestamp() *int64 {
	return s.ModifiedTimestamp
}

func (s *GetConfigRuleResponseBodyConfigRule) GetRegionIdsScope() *string {
	return s.RegionIdsScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetResourceGroupIdsScope() *string {
	return s.ResourceGroupIdsScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetResourceIdsScope() *string {
	return s.ResourceIdsScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetResourceNameScope() *string {
	return s.ResourceNameScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *GetConfigRuleResponseBodyConfigRule) GetScope() *GetConfigRuleResponseBodyConfigRuleScope {
	return s.Scope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetSource() *GetConfigRuleResponseBodyConfigRuleSource {
	return s.Source
}

func (s *GetConfigRuleResponseBodyConfigRule) GetTagKeyLogicScope() *string {
	return s.TagKeyLogicScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetTagKeyScope() *string {
	return s.TagKeyScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetTagValueScope() *string {
	return s.TagValueScope
}

func (s *GetConfigRuleResponseBodyConfigRule) GetTags() []*GetConfigRuleResponseBodyConfigRuleTags {
	return s.Tags
}

func (s *GetConfigRuleResponseBodyConfigRule) GetTagsScope() []*GetConfigRuleResponseBodyConfigRuleTagsScope {
	return s.TagsScope
}

func (s *GetConfigRuleResponseBodyConfigRule) SetAccountId(v int64) *GetConfigRuleResponseBodyConfigRule {
	s.AccountId = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetCompliance(v *GetConfigRuleResponseBodyConfigRuleCompliance) *GetConfigRuleResponseBodyConfigRule {
	s.Compliance = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleArn(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleArn = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleEvaluationStatus(v *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleEvaluationStatus = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleId(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleId = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleName(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleName = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleState(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleState = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleTriggerTypes(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetCreateBy(v *GetConfigRuleResponseBodyConfigRuleCreateBy) *GetConfigRuleResponseBodyConfigRule {
	s.CreateBy = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetCreateTimestamp(v int64) *GetConfigRuleResponseBodyConfigRule {
	s.CreateTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetDescription(v string) *GetConfigRuleResponseBodyConfigRule {
	s.Description = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetExcludeRegionIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ExcludeRegionIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetExcludeResourceGroupIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ExcludeResourceGroupIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetExcludeResourceIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetExcludeTagsScope(v []*GetConfigRuleResponseBodyConfigRuleExcludeTagsScope) *GetConfigRuleResponseBodyConfigRule {
	s.ExcludeTagsScope = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetExtendContent(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ExtendContent = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetInputParameters(v map[string]interface{}) *GetConfigRuleResponseBodyConfigRule {
	s.InputParameters = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetManagedRule(v *GetConfigRuleResponseBodyConfigRuleManagedRule) *GetConfigRuleResponseBodyConfigRule {
	s.ManagedRule = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetMaximumExecutionFrequency(v string) *GetConfigRuleResponseBodyConfigRule {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetModifiedTimestamp(v int64) *GetConfigRuleResponseBodyConfigRule {
	s.ModifiedTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetRegionIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.RegionIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetResourceGroupIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetResourceIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ResourceIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetResourceNameScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ResourceNameScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetResourceTypesScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ResourceTypesScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetRiskLevel(v int32) *GetConfigRuleResponseBodyConfigRule {
	s.RiskLevel = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetScope(v *GetConfigRuleResponseBodyConfigRuleScope) *GetConfigRuleResponseBodyConfigRule {
	s.Scope = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetSource(v *GetConfigRuleResponseBodyConfigRuleSource) *GetConfigRuleResponseBodyConfigRule {
	s.Source = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetTagKeyLogicScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.TagKeyLogicScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetTagKeyScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.TagKeyScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetTagValueScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.TagValueScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetTags(v []*GetConfigRuleResponseBodyConfigRuleTags) *GetConfigRuleResponseBodyConfigRule {
	s.Tags = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetTagsScope(v []*GetConfigRuleResponseBodyConfigRuleTagsScope) *GetConfigRuleResponseBodyConfigRule {
	s.TagsScope = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) Validate() error {
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
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
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

type GetConfigRuleResponseBodyConfigRuleCompliance struct {
	// The compliance evaluation result. Valid values:
	//
	// - COMPLIANT: The resource is compliant.
	//
	// - NON_COMPLIANT: The resource is non-compliant.
	//
	// - NOT_APPLICABLE: The rule does not apply to the resource.
	//
	// - INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// NON_COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The number of resources that are evaluated based on the compliance result.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleCompliance) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleCompliance) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleCompliance) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *GetConfigRuleResponseBodyConfigRuleCompliance) GetCount() *int32 {
	return s.Count
}

func (s *GetConfigRuleResponseBodyConfigRuleCompliance) SetComplianceType(v string) *GetConfigRuleResponseBodyConfigRuleCompliance {
	s.ComplianceType = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCompliance) SetCount(v int32) *GetConfigRuleResponseBodyConfigRuleCompliance {
	s.Count = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCompliance) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus struct {
	// The timestamp when the rule was first activated. Unit: milliseconds.
	//
	// example:
	//
	// 1624932221993
	FirstActivatedTimestamp *int64 `json:"FirstActivatedTimestamp,omitempty" xml:"FirstActivatedTimestamp,omitempty"`
	// Indicates whether the rule has been evaluated. Valid values:
	//
	// - true: The rule has been evaluated.
	//
	// - false: The rule has not been evaluated.
	//
	// example:
	//
	// true
	FirstEvaluationStarted *bool `json:"FirstEvaluationStarted,omitempty" xml:"FirstEvaluationStarted,omitempty"`
	// The error code returned for the last failed execution of the rule.
	//
	// example:
	//
	// TimeOut
	LastErrorCode *string `json:"LastErrorCode,omitempty" xml:"LastErrorCode,omitempty"`
	// The error message returned for the last failed execution of the rule.
	//
	// example:
	//
	// Time out
	LastErrorMessage *string `json:"LastErrorMessage,omitempty" xml:"LastErrorMessage,omitempty"`
	// The timestamp when the last failed evaluation of the rule ended. Unit: milliseconds.
	//
	// example:
	//
	// 1614687022000
	LastFailedEvaluationTimestamp *int64 `json:"LastFailedEvaluationTimestamp,omitempty" xml:"LastFailedEvaluationTimestamp,omitempty"`
	// The timestamp when the last failed invocation of the rule started. Unit: milliseconds.
	//
	// example:
	//
	// 1614687022000
	LastFailedInvocationTimestamp *int64 `json:"LastFailedInvocationTimestamp,omitempty" xml:"LastFailedInvocationTimestamp,omitempty"`
	// The timestamp when the last successful evaluation of the rule ended. Unit: milliseconds.
	//
	// example:
	//
	// 1624932227486
	LastSuccessfulEvaluationTimestamp *int64 `json:"LastSuccessfulEvaluationTimestamp,omitempty" xml:"LastSuccessfulEvaluationTimestamp,omitempty"`
	// The timestamp when the last successful invocation of the rule started. Unit: milliseconds.
	//
	// example:
	//
	// 1624932227476
	LastSuccessfulInvocationTimestamp *int64 `json:"LastSuccessfulInvocationTimestamp,omitempty" xml:"LastSuccessfulInvocationTimestamp,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetFirstActivatedTimestamp() *int64 {
	return s.FirstActivatedTimestamp
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetFirstEvaluationStarted() *bool {
	return s.FirstEvaluationStarted
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastErrorCode() *string {
	return s.LastErrorCode
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastErrorMessage() *string {
	return s.LastErrorMessage
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastFailedEvaluationTimestamp() *int64 {
	return s.LastFailedEvaluationTimestamp
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastFailedInvocationTimestamp() *int64 {
	return s.LastFailedInvocationTimestamp
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastSuccessfulEvaluationTimestamp() *int64 {
	return s.LastSuccessfulEvaluationTimestamp
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GetLastSuccessfulInvocationTimestamp() *int64 {
	return s.LastSuccessfulInvocationTimestamp
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstActivatedTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstActivatedTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstEvaluationStarted(v bool) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstEvaluationStarted = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorCode(v string) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorCode = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorMessage(v string) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorMessage = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedEvaluationTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedEvaluationTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedInvocationTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedInvocationTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulEvaluationTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulEvaluationTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulInvocationTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulInvocationTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleResponseBodyConfigRuleCreateBy struct {
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
	// OSS合规基线
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the Alibaba Cloud account that was used to create the rule.
	//
	// example:
	//
	// 100931896542****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// Alice
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleCreateBy) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleCreateBy) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) SetCompliancePackId(v string) *GetConfigRuleResponseBodyConfigRuleCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) SetCompliancePackName(v string) *GetConfigRuleResponseBodyConfigRuleCreateBy {
	s.CompliancePackName = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorId(v string) *GetConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorId = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorName(v string) *GetConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorName = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleResponseBodyConfigRuleExcludeTagsScope struct {
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

func (s GetConfigRuleResponseBodyConfigRuleExcludeTagsScope) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleExcludeTagsScope) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleExcludeTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *GetConfigRuleResponseBodyConfigRuleExcludeTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *GetConfigRuleResponseBodyConfigRuleExcludeTagsScope) SetTagKey(v string) *GetConfigRuleResponseBodyConfigRuleExcludeTagsScope {
	s.TagKey = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleExcludeTagsScope) SetTagValue(v string) *GetConfigRuleResponseBodyConfigRuleExcludeTagsScope {
	s.TagValue = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleExcludeTagsScope) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleResponseBodyConfigRuleManagedRule struct {
	// The details of the required input parameters of the managed rule.
	//
	// example:
	//
	// {}
	CompulsoryInputParameterDetails map[string]interface{} `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	// The description of the managed rule.
	//
	// example:
	//
	// ECS磁盘未因欠费或安全等原因而被锁定，视为“合规”。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the managed rule.
	//
	// example:
	//
	// ram-user-mfa-check
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The list of rule labels.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the managed rule.
	//
	// example:
	//
	// RAM用户开启MFA
	ManagedRuleName *string `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty"`
	// The details of the optional input parameters of the managed rule.
	//
	// example:
	//
	// {}
	OptionalInputParameterDetails map[string]interface{} `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
	// The source details of the managed rule.
	SourceDetails []*GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetConfigRuleResponseBodyConfigRuleManagedRule) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleManagedRule) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) GetCompulsoryInputParameterDetails() map[string]interface{} {
	return s.CompulsoryInputParameterDetails
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) GetDescription() *string {
	return s.Description
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) GetLabels() []*string {
	return s.Labels
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) GetManagedRuleName() *string {
	return s.ManagedRuleName
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) GetOptionalInputParameterDetails() map[string]interface{} {
	return s.OptionalInputParameterDetails
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) GetSourceDetails() []*GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	return s.SourceDetails
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetCompulsoryInputParameterDetails(v map[string]interface{}) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetDescription(v string) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.Description = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetIdentifier(v string) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.Identifier = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetLabels(v []*string) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.Labels = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetManagedRuleName(v string) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.ManagedRuleName = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetOptionalInputParameterDetails(v map[string]interface{}) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetSourceDetails(v []*GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.SourceDetails = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) Validate() error {
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

type GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails struct {
	// The event source.
	//
	// > Only Cloud Config events are supported. The value is aliyun.config.
	//
	// example:
	//
	// aliyun.config
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The execution frequency of the rule. Valid values:
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
	// > This parameter is returned only when the rule is triggered periodically.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// - ScheduledNotification: The rule is triggered periodically.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GetEventSource() *string {
	return s.EventSource
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GetMessageType() *string {
	return s.MessageType
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetEventSource(v string) *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.EventSource = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMaximumExecutionFrequency(v string) *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMessageType(v string) *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MessageType = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleResponseBodyConfigRuleScope struct {
	// The list of resource types that are evaluated by the rule. You can also view this information in the ResourceTypesScope field.
	ComplianceResourceTypes []*string `json:"ComplianceResourceTypes,omitempty" xml:"ComplianceResourceTypes,omitempty" type:"Repeated"`
}

func (s GetConfigRuleResponseBodyConfigRuleScope) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleScope) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleScope) GetComplianceResourceTypes() []*string {
	return s.ComplianceResourceTypes
}

func (s *GetConfigRuleResponseBodyConfigRuleScope) SetComplianceResourceTypes(v []*string) *GetConfigRuleResponseBodyConfigRuleScope {
	s.ComplianceResourceTypes = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleScope) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleResponseBodyConfigRuleSource struct {
	// The identifier of the rule.
	//
	// - If the rule is a managed rule, the value of this parameter is the identifier of the managed rule.
	//
	// - If the rule is a custom rule, the value of this parameter is the ARN of the function.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:100931896542****:services/ConfigService.LATEST/functions/specific-config
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The owner of the rule. Valid values:
	//
	// - CUSTOM_FC: a custom rule.
	//
	// - ALIYUN: a managed rule.
	//
	// example:
	//
	// ALIYUN
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The source details.
	SourceDetails []*GetConfigRuleResponseBodyConfigRuleSourceSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetConfigRuleResponseBodyConfigRuleSource) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleSource) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) GetOwner() *string {
	return s.Owner
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) GetSourceDetails() []*GetConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	return s.SourceDetails
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) SetIdentifier(v string) *GetConfigRuleResponseBodyConfigRuleSource {
	s.Identifier = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) SetOwner(v string) *GetConfigRuleResponseBodyConfigRuleSource {
	s.Owner = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) SetSourceDetails(v []*GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) *GetConfigRuleResponseBodyConfigRuleSource {
	s.SourceDetails = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) Validate() error {
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

type GetConfigRuleResponseBodyConfigRuleSourceSourceDetails struct {
	// The event source.
	//
	// > Only Cloud Config events are supported. The value is aliyun.config.
	//
	// example:
	//
	// aliyun.config
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The execution frequency of the rule. Valid values:
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
	// > This parameter is returned only when the rule is triggered periodically.
	//
	// example:
	//
	// One_Hour
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// - ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	//
	// - ScheduledNotification: The rule is triggered periodically.
	//
	// example:
	//
	// ConfigurationItemChangeNotification
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) GetEventSource() *string {
	return s.EventSource
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) GetMaximumExecutionFrequency() *string {
	return s.MaximumExecutionFrequency
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) GetMessageType() *string {
	return s.MessageType
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetEventSource(v string) *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.EventSource = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMaximumExecutionFrequency(v string) *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMessageType(v string) *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MessageType = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleResponseBodyConfigRuleTags struct {
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

func (s GetConfigRuleResponseBodyConfigRuleTags) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleTags) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetConfigRuleResponseBodyConfigRuleTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetConfigRuleResponseBodyConfigRuleTags) SetTagKey(v string) *GetConfigRuleResponseBodyConfigRuleTags {
	s.TagKey = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleTags) SetTagValue(v string) *GetConfigRuleResponseBodyConfigRuleTags {
	s.TagValue = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleTags) Validate() error {
	return dara.Validate(s)
}

type GetConfigRuleResponseBodyConfigRuleTagsScope struct {
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

func (s GetConfigRuleResponseBodyConfigRuleTagsScope) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleTagsScope) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleTagsScope) GetTagKey() *string {
	return s.TagKey
}

func (s *GetConfigRuleResponseBodyConfigRuleTagsScope) GetTagValue() *string {
	return s.TagValue
}

func (s *GetConfigRuleResponseBodyConfigRuleTagsScope) SetTagKey(v string) *GetConfigRuleResponseBodyConfigRuleTagsScope {
	s.TagKey = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleTagsScope) SetTagValue(v string) *GetConfigRuleResponseBodyConfigRuleTagsScope {
	s.TagValue = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleTagsScope) Validate() error {
	return dara.Validate(s)
}
