// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSIntelligentAclRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeHttpDDoSIntelligentAclRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHttpDDoSIntelligentAclRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHttpDDoSIntelligentAclRulesResponseBody
	GetRequestId() *string
	SetRuleInfos(v []*DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) *DescribeHttpDDoSIntelligentAclRulesResponseBody
	GetRuleInfos() []*DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos
	SetTotalCount(v int64) *DescribeHttpDDoSIntelligentAclRulesResponseBody
	GetTotalCount() *int64
}

type DescribeHttpDDoSIntelligentAclRulesResponseBody struct {
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
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleInfos []*DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos `json:"RuleInfos,omitempty" xml:"RuleInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHttpDDoSIntelligentAclRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSIntelligentAclRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) GetRuleInfos() []*DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos {
	return s.RuleInfos
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) SetPageNumber(v int32) *DescribeHttpDDoSIntelligentAclRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) SetPageSize(v int32) *DescribeHttpDDoSIntelligentAclRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) SetRequestId(v string) *DescribeHttpDDoSIntelligentAclRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) SetRuleInfos(v []*DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) *DescribeHttpDDoSIntelligentAclRulesResponseBody {
	s.RuleInfos = v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) SetTotalCount(v int64) *DescribeHttpDDoSIntelligentAclRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBody) Validate() error {
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

type DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos struct {
	// example:
	//
	// deny
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {"$and":[{"key":"URI","opValue":"prefix-match","values":"/"}]}
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// 1000030
	LogRuleId *int64 `json:"LogRuleId,omitempty" xml:"LogRuleId,omitempty"`
	// example:
	//
	// 1800
	PunishTime *int64 `json:"PunishTime,omitempty" xml:"PunishTime,omitempty"`
	// example:
	//
	// test.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// 20569929
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// smart_cc_***
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) GetCondition() *string {
	return s.Condition
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) GetLogRuleId() *int64 {
	return s.LogRuleId
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) GetPunishTime() *int64 {
	return s.PunishTime
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) GetRecordName() *string {
	return s.RecordName
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) SetAction(v string) *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos {
	s.Action = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) SetCondition(v string) *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos {
	s.Condition = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) SetLogRuleId(v int64) *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos {
	s.LogRuleId = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) SetPunishTime(v int64) *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos {
	s.PunishTime = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) SetRecordName(v string) *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos {
	s.RecordName = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) SetRuleId(v int64) *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos {
	s.RuleId = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) SetRuleName(v string) *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos {
	s.RuleName = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentAclRulesResponseBodyRuleInfos) Validate() error {
	return dara.Validate(s)
}
