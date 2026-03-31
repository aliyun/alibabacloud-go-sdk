// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBaseRuleChangeLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeBaseRuleChangeLogResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeBaseRuleChangeLogResponseBodyRules) *DescribeBaseRuleChangeLogResponseBody
	GetRules() []*DescribeBaseRuleChangeLogResponseBodyRules
	SetTotalCount(v int64) *DescribeBaseRuleChangeLogResponseBody
	GetTotalCount() *int64
}

type DescribeBaseRuleChangeLogResponseBody struct {
	// example:
	//
	// 6FBF08CB-8691-5B65-BBF8-***
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeBaseRuleChangeLogResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// example:
	//
	// 63
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBaseRuleChangeLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseRuleChangeLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBaseRuleChangeLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBaseRuleChangeLogResponseBody) GetRules() []*DescribeBaseRuleChangeLogResponseBodyRules {
	return s.Rules
}

func (s *DescribeBaseRuleChangeLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeBaseRuleChangeLogResponseBody) SetRequestId(v string) *DescribeBaseRuleChangeLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBaseRuleChangeLogResponseBody) SetRules(v []*DescribeBaseRuleChangeLogResponseBodyRules) *DescribeBaseRuleChangeLogResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeBaseRuleChangeLogResponseBody) SetTotalCount(v int64) *DescribeBaseRuleChangeLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBaseRuleChangeLogResponseBody) Validate() error {
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

type DescribeBaseRuleChangeLogResponseBodyRules struct {
	// example:
	//
	// CVE-2021-34538
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// 42755
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// Header XSS Scanner Behavior
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1665460629000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeBaseRuleChangeLogResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeBaseRuleChangeLogResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) GetCveId() *string {
	return s.CveId
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) GetOperation() *string {
	return s.Operation
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) SetCveId(v string) *DescribeBaseRuleChangeLogResponseBodyRules {
	s.CveId = &v
	return s
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) SetOperation(v string) *DescribeBaseRuleChangeLogResponseBodyRules {
	s.Operation = &v
	return s
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) SetRuleId(v int64) *DescribeBaseRuleChangeLogResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) SetRuleName(v string) *DescribeBaseRuleChangeLogResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) SetUpdateTime(v int64) *DescribeBaseRuleChangeLogResponseBodyRules {
	s.UpdateTime = &v
	return s
}

func (s *DescribeBaseRuleChangeLogResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
