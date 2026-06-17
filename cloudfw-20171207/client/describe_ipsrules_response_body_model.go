// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPSRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDropCount(v int64) *DescribeIPSRulesResponseBody
	GetDropCount() *int64
	SetHighRiskCount(v int64) *DescribeIPSRulesResponseBody
	GetHighRiskCount() *int64
	SetOpenCount(v int64) *DescribeIPSRulesResponseBody
	GetOpenCount() *int64
	SetPageNo(v int64) *DescribeIPSRulesResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeIPSRulesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeIPSRulesResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeIPSRulesResponseBodyRules) *DescribeIPSRulesResponseBody
	GetRules() []*DescribeIPSRulesResponseBodyRules
	SetTotalCount(v int64) *DescribeIPSRulesResponseBody
	GetTotalCount() *int64
	SetUserDefineCount(v int64) *DescribeIPSRulesResponseBody
	GetUserDefineCount() *int64
}

type DescribeIPSRulesResponseBody struct {
	// The number of rules that have the `drop` action.
	//
	// example:
	//
	// 976
	DropCount *int64 `json:"DropCount,omitempty" xml:"DropCount,omitempty"`
	// The number of high-risk rules.
	//
	// example:
	//
	// 518
	HighRiskCount *int64 `json:"HighRiskCount,omitempty" xml:"HighRiskCount,omitempty"`
	// The total number of enabled rules.
	//
	// example:
	//
	// 1752
	OpenCount *int64 `json:"OpenCount,omitempty" xml:"OpenCount,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6DF55258-1448-5386-B263-4771D081****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of rules.
	Rules []*DescribeIPSRulesResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of user-defined rules.
	//
	// example:
	//
	// 32
	UserDefineCount *int64 `json:"UserDefineCount,omitempty" xml:"UserDefineCount,omitempty"`
}

func (s DescribeIPSRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPSRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIPSRulesResponseBody) GetDropCount() *int64 {
	return s.DropCount
}

func (s *DescribeIPSRulesResponseBody) GetHighRiskCount() *int64 {
	return s.HighRiskCount
}

func (s *DescribeIPSRulesResponseBody) GetOpenCount() *int64 {
	return s.OpenCount
}

func (s *DescribeIPSRulesResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeIPSRulesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeIPSRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIPSRulesResponseBody) GetRules() []*DescribeIPSRulesResponseBodyRules {
	return s.Rules
}

func (s *DescribeIPSRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeIPSRulesResponseBody) GetUserDefineCount() *int64 {
	return s.UserDefineCount
}

func (s *DescribeIPSRulesResponseBody) SetDropCount(v int64) *DescribeIPSRulesResponseBody {
	s.DropCount = &v
	return s
}

func (s *DescribeIPSRulesResponseBody) SetHighRiskCount(v int64) *DescribeIPSRulesResponseBody {
	s.HighRiskCount = &v
	return s
}

func (s *DescribeIPSRulesResponseBody) SetOpenCount(v int64) *DescribeIPSRulesResponseBody {
	s.OpenCount = &v
	return s
}

func (s *DescribeIPSRulesResponseBody) SetPageNo(v int64) *DescribeIPSRulesResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeIPSRulesResponseBody) SetPageSize(v int64) *DescribeIPSRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIPSRulesResponseBody) SetRequestId(v string) *DescribeIPSRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIPSRulesResponseBody) SetRules(v []*DescribeIPSRulesResponseBodyRules) *DescribeIPSRulesResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeIPSRulesResponseBody) SetTotalCount(v int64) *DescribeIPSRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIPSRulesResponseBody) SetUserDefineCount(v int64) *DescribeIPSRulesResponseBody {
	s.UserDefineCount = &v
	return s
}

func (s *DescribeIPSRulesResponseBody) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIPSRulesResponseBodyRules struct {
	// The target application.
	//
	// example:
	//
	// SMB
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// The attack type.
	//
	// example:
	//
	// Exception connection
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The current action.
	//
	// example:
	//
	// alert
	CurrentMode *string `json:"CurrentMode,omitempty" xml:"CurrentMode,omitempty"`
	// The CVE ID.
	//
	// example:
	//
	// cve-2024-38816
	Cve *string `json:"Cve,omitempty" xml:"Cve,omitempty"`
	// The default action.
	//
	// example:
	//
	// alert
	DefaultMode *string `json:"DefaultMode,omitempty" xml:"DefaultMode,omitempty"`
	// A description of the rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The engine mode.
	//
	// This parameter takes effect only when `DefaultMode` is set to `drop`.
	//
	// example:
	//
	// 1
	RuleClass *int32 `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// f532f59d-2026-436b-8209-e04d8ebc2****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule level.
	//
	// example:
	//
	// 1
	RuleLevel *int64 `json:"RuleLevel,omitempty" xml:"RuleLevel,omitempty"`
	// The rule name.
	//
	// example:
	//
	// Nmap扫描探测
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule UUID.
	//
	// example:
	//
	// d401b0cb-dc64-4bbe-bba0-3e7c744****
	RuleUuid *string `json:"RuleUuid,omitempty" xml:"RuleUuid,omitempty"`
	// The UNIX timestamp for when the rule was last updated.
	//
	// example:
	//
	// 1775101028
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Specifies whether the rule is user-defined.
	//
	// example:
	//
	// 1
	UserDefined *string `json:"UserDefined,omitempty" xml:"UserDefined,omitempty"`
	// The user-defined status of the rule.
	//
	// example:
	//
	// 1
	UserStatus *int32 `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeIPSRulesResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPSRulesResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeIPSRulesResponseBodyRules) GetAttackApp() *string {
	return s.AttackApp
}

func (s *DescribeIPSRulesResponseBodyRules) GetAttackType() *string {
	return s.AttackType
}

func (s *DescribeIPSRulesResponseBodyRules) GetCurrentMode() *string {
	return s.CurrentMode
}

func (s *DescribeIPSRulesResponseBodyRules) GetCve() *string {
	return s.Cve
}

func (s *DescribeIPSRulesResponseBodyRules) GetDefaultMode() *string {
	return s.DefaultMode
}

func (s *DescribeIPSRulesResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeIPSRulesResponseBodyRules) GetRuleClass() *int32 {
	return s.RuleClass
}

func (s *DescribeIPSRulesResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeIPSRulesResponseBodyRules) GetRuleLevel() *int64 {
	return s.RuleLevel
}

func (s *DescribeIPSRulesResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeIPSRulesResponseBodyRules) GetRuleUuid() *string {
	return s.RuleUuid
}

func (s *DescribeIPSRulesResponseBodyRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeIPSRulesResponseBodyRules) GetUserDefined() *string {
	return s.UserDefined
}

func (s *DescribeIPSRulesResponseBodyRules) GetUserStatus() *int32 {
	return s.UserStatus
}

func (s *DescribeIPSRulesResponseBodyRules) SetAttackApp(v string) *DescribeIPSRulesResponseBodyRules {
	s.AttackApp = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetAttackType(v string) *DescribeIPSRulesResponseBodyRules {
	s.AttackType = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetCurrentMode(v string) *DescribeIPSRulesResponseBodyRules {
	s.CurrentMode = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetCve(v string) *DescribeIPSRulesResponseBodyRules {
	s.Cve = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetDefaultMode(v string) *DescribeIPSRulesResponseBodyRules {
	s.DefaultMode = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetDescription(v string) *DescribeIPSRulesResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetRuleClass(v int32) *DescribeIPSRulesResponseBodyRules {
	s.RuleClass = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetRuleId(v string) *DescribeIPSRulesResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetRuleLevel(v int64) *DescribeIPSRulesResponseBodyRules {
	s.RuleLevel = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetRuleName(v string) *DescribeIPSRulesResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetRuleUuid(v string) *DescribeIPSRulesResponseBodyRules {
	s.RuleUuid = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetUpdateTime(v int64) *DescribeIPSRulesResponseBodyRules {
	s.UpdateTime = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetUserDefined(v string) *DescribeIPSRulesResponseBodyRules {
	s.UserDefined = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) SetUserStatus(v int32) *DescribeIPSRulesResponseBodyRules {
	s.UserStatus = &v
	return s
}

func (s *DescribeIPSRulesResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
