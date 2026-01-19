// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSIntelligentRateLimitRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody
	GetRequestId() *string
	SetRuleInfos(v []*DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody
	GetRuleInfos() []*DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos
	SetTotalCount(v int64) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody
	GetTotalCount() *int64
}

type DescribeHttpDDoSIntelligentRateLimitRulesResponseBody struct {
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
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleInfos []*DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos `json:"RuleInfos,omitempty" xml:"RuleInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) GetRuleInfos() []*DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	return s.RuleInfos
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) SetPageNumber(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) SetPageSize(v int32) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) SetRequestId(v string) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) SetRuleInfos(v []*DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody {
	s.RuleInfos = v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) SetTotalCount(v int64) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBody) Validate() error {
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

type DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos struct {
	// example:
	//
	// js
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// {"$and":[{"key":"URI","opValue":"prefix-match","values":"/"}]}
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// 100030
	LogRuleId *int64 `json:"LogRuleId,omitempty" xml:"LogRuleId,omitempty"`
	// example:
	//
	// 86400
	PunishTime *int64 `json:"PunishTime,omitempty" xml:"PunishTime,omitempty"`
	// example:
	//
	// {"threshold":2000,"interval":5,"target":"ip","ttl":600}
	RateLimit *string `json:"RateLimit,omitempty" xml:"RateLimit,omitempty"`
	// example:
	//
	// test.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// example:
	//
	// 20110849
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// inner_cc_client_ip_ratelimit
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// {"field":"ip","mode":"count"}
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
}

func (s DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetCondition() *string {
	return s.Condition
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetLogRuleId() *int64 {
	return s.LogRuleId
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetPunishTime() *int64 {
	return s.PunishTime
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetRateLimit() *string {
	return s.RateLimit
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetRecordName() *string {
	return s.RecordName
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetAction(v string) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.Action = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetCondition(v string) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.Condition = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetLogRuleId(v int64) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.LogRuleId = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetPunishTime(v int64) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.PunishTime = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetRateLimit(v string) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.RateLimit = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetRecordName(v string) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.RecordName = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetRuleId(v int64) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.RuleId = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetRuleName(v string) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.RuleName = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) SetStatistics(v string) *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos {
	s.Statistics = &v
	return s
}

func (s *DescribeHttpDDoSIntelligentRateLimitRulesResponseBodyRuleInfos) Validate() error {
	return dara.Validate(s)
}
