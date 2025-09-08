// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBaseSystemRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeBaseSystemRulesResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeBaseSystemRulesResponseBodyRules) *DescribeBaseSystemRulesResponseBody
	GetRules() []*DescribeBaseSystemRulesResponseBodyRules
	SetTotalCount(v int64) *DescribeBaseSystemRulesResponseBody
	GetTotalCount() *int64
}

type DescribeBaseSystemRulesResponseBody struct {
	// example:
	//
	// 80736FA5-FA87-55F6-AA69-C5477C6FE6D0
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeBaseSystemRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBaseSystemRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseSystemRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBaseSystemRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBaseSystemRulesResponseBody) GetRules() []*DescribeBaseSystemRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeBaseSystemRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeBaseSystemRulesResponseBody) SetRequestId(v string) *DescribeBaseSystemRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBody) SetRules(v []*DescribeBaseSystemRulesResponseBodyRules) *DescribeBaseSystemRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeBaseSystemRulesResponseBody) SetTotalCount(v int64) *DescribeBaseSystemRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBaseSystemRulesResponseBodyRules struct {
	// example:
	//
	// CVE-2021-34538
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// example:
	//
	// rule description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// sqli
	DetectType *string `json:"DetectType,omitempty" xml:"DetectType,omitempty"`
	// example:
	//
	// super_strict
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// block
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// example:
	//
	// 113089
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// systemRuleTest
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	RuleStatus *int32 `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// example:
	//
	// 1665460629000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeBaseSystemRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseSystemRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetCveId() *string {
	return s.CveId
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetDetectType() *string {
	return s.DetectType
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRuleAction() *string {
	return s.RuleAction
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetRuleStatus() *int32 {
	return s.RuleStatus
}

func (s *DescribeBaseSystemRulesResponseBodyRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetCveId(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.CveId = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetDescription(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetDetectType(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.DetectType = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRiskLevel(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.RiskLevel = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRuleAction(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.RuleAction = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRuleId(v int64) *DescribeBaseSystemRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRuleName(v string) *DescribeBaseSystemRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetRuleStatus(v int32) *DescribeBaseSystemRulesResponseBodyRules {
	s.RuleStatus = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) SetUpdateTime(v int64) *DescribeBaseSystemRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

func (s *DescribeBaseSystemRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
