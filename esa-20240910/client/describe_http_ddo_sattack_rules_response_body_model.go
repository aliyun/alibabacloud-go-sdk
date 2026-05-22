// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeHttpDDoSAttackRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHttpDDoSAttackRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHttpDDoSAttackRulesResponseBody
	GetRequestId() *string
	SetRuleInfos(v []*DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) *DescribeHttpDDoSAttackRulesResponseBody
	GetRuleInfos() []*DescribeHttpDDoSAttackRulesResponseBodyRuleInfos
	SetTotalCount(v int32) *DescribeHttpDDoSAttackRulesResponseBody
	GetTotalCount() *int32
}

type DescribeHttpDDoSAttackRulesResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D4030CD2-0D9D-5E92-B358-421AE58307C6
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleInfos []*DescribeHttpDDoSAttackRulesResponseBodyRuleInfos `json:"RuleInfos,omitempty" xml:"RuleInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHttpDDoSAttackRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) GetRuleInfos() []*DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	return s.RuleInfos
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) SetPageNumber(v int32) *DescribeHttpDDoSAttackRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) SetPageSize(v int32) *DescribeHttpDDoSAttackRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) SetRequestId(v string) *DescribeHttpDDoSAttackRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) SetRuleInfos(v []*DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) *DescribeHttpDDoSAttackRulesResponseBody {
	s.RuleInfos = v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) SetTotalCount(v int32) *DescribeHttpDDoSAttackRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBody) Validate() error {
	if s.RuleInfos != nil {
		for _, item := range s.RuleInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHttpDDoSAttackRulesResponseBodyRuleInfos struct {
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// deny
	DefaultAction *string `json:"DefaultAction,omitempty" xml:"DefaultAction,omitempty"`
	// example:
	//
	// 100010
	LogRuleId *int32 `json:"LogRuleId,omitempty" xml:"LogRuleId,omitempty"`
	// example:
	//
	// The HTTP request\\"s Accept header contains invalid features#1
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// example:
	//
	// 20203578
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// global_01_s
	RuleIdInfo *string `json:"RuleIdInfo,omitempty" xml:"RuleIdInfo,omitempty"`
	// example:
	//
	// Illegal request。
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// on
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GetDefaultAction() *string {
	return s.DefaultAction
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GetLogRuleId() *int32 {
	return s.LogRuleId
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GetRuleId() *int32 {
	return s.RuleId
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GetRuleIdInfo() *string {
	return s.RuleIdInfo
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) SetAction(v string) *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	s.Action = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) SetDefaultAction(v string) *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	s.DefaultAction = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) SetLogRuleId(v int32) *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	s.LogRuleId = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) SetRuleDesc(v string) *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	s.RuleDesc = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) SetRuleId(v int32) *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	s.RuleId = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) SetRuleIdInfo(v string) *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	s.RuleIdInfo = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) SetRuleName(v string) *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	s.RuleName = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) SetStatus(v string) *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos {
	s.Status = &v
	return s
}

func (s *DescribeHttpDDoSAttackRulesResponseBodyRuleInfos) Validate() error {
	return dara.Validate(s)
}
