// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRules(v *ListConfigRulesResponseBodyConfigRules) *ListConfigRulesResponseBody
	GetConfigRules() *ListConfigRulesResponseBodyConfigRules
	SetRequestId(v string) *ListConfigRulesResponseBody
	GetRequestId() *string
}

type ListConfigRulesResponseBody struct {
	// The information about the rules.
	ConfigRules *ListConfigRulesResponseBodyConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AC3A7E12-72E6-5CC9-A5C1-D8D8919829A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBody) GetConfigRules() *ListConfigRulesResponseBodyConfigRules {
	return s.ConfigRules
}

func (s *ListConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigRulesResponseBody) SetConfigRules(v *ListConfigRulesResponseBodyConfigRules) *ListConfigRulesResponseBody {
	s.ConfigRules = v
	return s
}

func (s *ListConfigRulesResponseBody) SetRequestId(v string) *ListConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigRulesResponseBody) Validate() error {
	if s.ConfigRules != nil {
		if err := s.ConfigRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConfigRulesResponseBodyConfigRules struct {
	// The details of the rule.
	ConfigRuleList []*ListConfigRulesResponseBodyConfigRulesConfigRuleList `json:"ConfigRuleList,omitempty" xml:"ConfigRuleList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
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

func (s ListConfigRulesResponseBodyConfigRules) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRules) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRules) GetConfigRuleList() []*ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	return s.ConfigRuleList
}

func (s *ListConfigRulesResponseBodyConfigRules) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConfigRulesResponseBodyConfigRules) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConfigRulesResponseBodyConfigRules) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConfigRulesResponseBodyConfigRules) SetConfigRuleList(v []*ListConfigRulesResponseBodyConfigRulesConfigRuleList) *ListConfigRulesResponseBodyConfigRules {
	s.ConfigRuleList = v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRules) SetPageNumber(v int32) *ListConfigRulesResponseBodyConfigRules {
	s.PageNumber = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRules) SetPageSize(v int32) *ListConfigRulesResponseBodyConfigRules {
	s.PageSize = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRules) SetTotalCount(v int64) *ListConfigRulesResponseBodyConfigRules {
	s.TotalCount = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRules) Validate() error {
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

type ListConfigRulesResponseBodyConfigRulesConfigRuleList struct {
	// The ID of the account to which the rule belongs.
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
	// The compliance aggregation result of the rule.
	Compliance *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance `json:"Compliance,omitempty" xml:"Compliance,omitempty" type:"Struct"`
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
	// 	- ACTIVE: The rule is enabled.
	//
	// 	- DELETING: The rule is being deleted.
	//
	// 	- EVALUATING: The rule is being used to evaluate resource configurations.
	//
	// 	- INACTIVE: The rule is disabled.
	//
	// example:
	//
	// ACTIVE
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The information about the creation of the rule.
	CreateBy   *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	CreateDate *string                                                       `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
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
	// 	- 1: high.
	//
	// 	- 2: medium.
	//
	// 	- 3: low.
	//
	// example:
	//
	// 1
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The identifier of the rule.
	//
	// 	- If the rule is a managed rule, the value of this parameter is the identifier of the managed rule.
	//
	// 	- If the rule is a custom rule, the value of this parameter is the Alibaba Cloud Resource Name (ARN) of the rule.
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
	Tags []*ListConfigRulesResponseBodyConfigRulesConfigRuleListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleList) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleList) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetAutomationType() *string {
	return s.AutomationType
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetCompliance() *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	return s.Compliance
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetConfigRuleArn() *string {
	return s.ConfigRuleArn
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetConfigRuleName() *string {
	return s.ConfigRuleName
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetConfigRuleState() *string {
	return s.ConfigRuleState
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetCreateBy() *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	return s.CreateBy
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetDescription() *string {
	return s.Description
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetResourceTypesScope() *string {
	return s.ResourceTypesScope
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetSourceIdentifier() *string {
	return s.SourceIdentifier
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetSourceOwner() *string {
	return s.SourceOwner
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) GetTags() []*ListConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	return s.Tags
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetAccountId(v int64) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.AccountId = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetAutomationType(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.AutomationType = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetCompliance(v *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Compliance = v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleArn(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleId(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleName(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleState(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleState = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetCreateBy(v *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.CreateBy = v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetCreateDate(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.CreateDate = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetDescription(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Description = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetResourceTypesScope(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ResourceTypesScope = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetRiskLevel(v int32) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.RiskLevel = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetSourceIdentifier(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.SourceIdentifier = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetSourceOwner(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.SourceOwner = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetTags(v []*ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Tags = v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) Validate() error {
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

type ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance struct {
	// The compliance evaluation result of the rule. Valid values:
	//
	// 	- COMPLIANT: The resources are evaluated as compliant.
	//
	// 	- NON_COMPLIANT: The resources are evaluated as non-compliant.
	//
	// 	- NOT_APPLICABLE: The rule does not apply to the resources.
	//
	// 	- INSUFFICIENT_DATA: No resource data is available.
	//
	// example:
	//
	// COMPLIANT
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The number of resources that are evaluated based on the rule.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) GetComplianceType() *string {
	return s.ComplianceType
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) GetCount() *int32 {
	return s.Count
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) SetComplianceType(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) SetCount(v int32) *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	s.Count = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) Validate() error {
	return dara.Validate(s)
}

type ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy struct {
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
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GetCompliancePackName() *string {
	return s.CompliancePackName
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCompliancePackId(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCompliancePackName(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CompliancePackName = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) Validate() error {
	return dara.Validate(s)
}

type ListConfigRulesResponseBodyConfigRulesConfigRuleListTags struct {
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

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) GetKey() *string {
	return s.Key
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) GetValue() *string {
	return s.Value
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) SetKey(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	s.Key = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) SetValue(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	s.Value = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) Validate() error {
	return dara.Validate(s)
}
