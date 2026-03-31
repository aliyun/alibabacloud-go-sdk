// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRules(v *ListAggregateConfigRulesResponseBodyConfigRules) *ListAggregateConfigRulesResponseBody
	GetConfigRules() *ListAggregateConfigRulesResponseBodyConfigRules
	SetRequestId(v string) *ListAggregateConfigRulesResponseBody
	GetRequestId() *string
}

type ListAggregateConfigRulesResponseBody struct {
	// The queried rules.
	ConfigRules *ListAggregateConfigRulesResponseBodyConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22EF8287-2C9A-4F1F-80A6-CEFA7612689D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBody) GetConfigRules() *ListAggregateConfigRulesResponseBodyConfigRules {
	return s.ConfigRules
}

func (s *ListAggregateConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateConfigRulesResponseBody) SetConfigRules(v *ListAggregateConfigRulesResponseBodyConfigRules) *ListAggregateConfigRulesResponseBody {
	s.ConfigRules = v
	return s
}

func (s *ListAggregateConfigRulesResponseBody) SetRequestId(v string) *ListAggregateConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBody) Validate() error {
	if s.ConfigRules != nil {
		if err := s.ConfigRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateConfigRulesResponseBodyConfigRules struct {
	// The details of the rule.
	ConfigRuleList []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList `json:"ConfigRuleList,omitempty" xml:"ConfigRuleList,omitempty" type:"Repeated"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRules) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRules) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) GetConfigRuleList() []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	return s.ConfigRuleList
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) SetConfigRuleList(v []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) *ListAggregateConfigRulesResponseBodyConfigRules {
	s.ConfigRuleList = v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) SetPageNumber(v int32) *ListAggregateConfigRulesResponseBodyConfigRules {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) SetPageSize(v int32) *ListAggregateConfigRulesResponseBodyConfigRules {
	s.PageSize = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) SetTotalCount(v int64) *ListAggregateConfigRulesResponseBodyConfigRules {
	s.TotalCount = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) Validate() error {
	if s.ConfigRuleList != nil {
		for _, item := range s.ConfigRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList struct {
	// The ID of the management account to which the rules belong.
	//
	// example:
	//
	// 100931896542****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The type of the remediation template. Only OOS is returned, which indicates CloudOps Orchestration Service.
	//
	// example:
	//
	// OOS
	AutomationType *string `json:"AutomationType,omitempty" xml:"AutomationType,omitempty"`
	// The compliance evaluation result.
	Compliance *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance `json:"Compliance,omitempty" xml:"Compliance,omitempty" type:"Struct"`
	// The ARN of the rule.
	//
	// example:
	//
	// acs:config::100931896542****:rule/cr-fdc8626622af00f9****
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cr-fdc8626622af00f9****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-rule-name
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- ACTIVE: The rule is being used to monitor resource configurations.
	//
	// 	- DELETING: The rule is being deleted.
	//
	// 	- EVALUATING: The rule is triggered and is being used to monitor resource configurations.
	//
	// 	- INACTIVE: The rule is disabled.
	//
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The information about the creation of the rule.
	CreateBy   *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	CreateDate *string                                                                `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// The description of the test rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The types of resources evaluated by the rule. Multiple resource types are separated with commas (,).
	//
	// example:
	//
	// ACS::EIP::EipAddress
	ResourceTypesScope *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
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
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The identifier of the rule.
	//
	// 	- If the rule is a managed rule, the value of this parameter is the name of the managed rule.
	//
	// 	- If the rule is a custom rule, the value of this parameter is the Alibaba Cloud Resource Name (ARN) of a function.
	//
	// example:
	//
	// eip-bandwidth-limit
	SourceIdentifier *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- CUSTOM_FC: a custom rule.
	//
	// 	- ALIYUN: a managed rule.
	//
	// example:
	//
	// ALIYUN
	SourceOwner *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	// The tags of the rule.
	Tags []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetAutomationType() *string {
	return s.AutomationType
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetCompliance() *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	return s.Compliance
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetConfigRuleArn() *string {
	return s.ConfigRuleArn
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetCreateBy() *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	return s.CreateBy
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetDescription() *string {
	return s.Description
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetSourceIdentifier() *string {
	return s.SourceIdentifier
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetSourceOwner() *string {
	return s.SourceOwner
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GetTags() []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	return s.Tags
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetAccountId(v int64) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.AccountId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetAutomationType(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.AutomationType = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetCompliance(v *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Compliance = v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleArn(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleId(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleName(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleState(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleState = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetCreateBy(v *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.CreateBy = v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetCreateDate(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.CreateDate = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetDescription(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Description = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetResourceTypesScope(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ResourceTypesScope = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetRiskLevel(v int32) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetSourceIdentifier(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.SourceIdentifier = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetSourceOwner(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.SourceOwner = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetTags(v []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Tags = v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) Validate() error {
	if s.Compliance != nil {
		if err := s.Compliance.Validate(); err != nil {
			return err
		}
	}
	if s.CreateBy != nil {
		if err := s.CreateBy.Validate(); err != nil {
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
	return nil
}

type ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance struct {
	// The compliance evaluation result. Valid values:
	//
	// 	- COMPLIANT: The resources are evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resources are evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resources.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The number of evaluation resources that correspond to the summary result of the rule compliance evaluation.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) GetCount() *int32 {
	return s.Count
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) SetComplianceType(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) SetCount(v int32) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	s.Count = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) Validate() error {
	return dara.Validate(s)
}

type ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy struct {
	// The account group ID.
	//
	// example:
	//
	// ca-f632626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The name of the account group.
	//
	// example:
	//
	// Test_Group
	AggregatorName *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	// The compliance package ID.
	//
	// example:
	//
	// cp-fdc8626622af00f9****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance package.
	//
	// example:
	//
	// test-pack-name
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	// The ID of the management account that created the rule.
	//
	// example:
	//
	// 100931896542****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The name of the management account that create the rule.
	//
	// example:
	//
	// Alice
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The type of the creator of the rule. The value is fixed to AGGREGATOR, which indicates an account group.
	//
	// example:
	//
	// AGGREGATOR
	CreatorType *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetAggregatorName() *string {
	return s.AggregatorName
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetCreatorType() *string {
	return s.CreatorType
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetAggregatorId(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetAggregatorName(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.AggregatorName = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCompliancePackId(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCompliancePackName(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CompliancePackName = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCreatorId(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CreatorId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCreatorName(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CreatorName = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCreatorType(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CreatorType = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) Validate() error {
	return dara.Validate(s)
}

type ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags struct {
	// The tag key of the rule.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the rule.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) GetKey() *string {
	return s.Key
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) GetValue() *string {
	return s.Value
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) SetKey(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	s.Key = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) SetValue(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	s.Value = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) Validate() error {
	return dara.Validate(s)
}
