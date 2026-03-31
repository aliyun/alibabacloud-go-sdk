// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRuleEvaluationStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationResults(v []*ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) *ListAggregateConfigRuleEvaluationStatisticsResponseBody
	GetEvaluationResults() []*ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults
	SetRequestId(v string) *ListAggregateConfigRuleEvaluationStatisticsResponseBody
	GetRequestId() *string
}

type ListAggregateConfigRuleEvaluationStatisticsResponseBody struct {
	// The statistics of compliance evaluation results.
	EvaluationResults []*ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9EFA436B-FC6F-513B-9DB8-C96E6CEBE5E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBody) GetEvaluationResults() []*ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	return s.EvaluationResults
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBody) SetEvaluationResults(v []*ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) *ListAggregateConfigRuleEvaluationStatisticsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBody) SetRequestId(v string) *ListAggregateConfigRuleEvaluationStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBody) Validate() error {
	if s.EvaluationResults != nil {
		for _, item := range s.EvaluationResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults struct {
	// The ID of the account group.
	//
	// example:
	//
	// ca-edd3626622af00b3****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The number of resources that are evaluated as non-compliant.
	//
	// example:
	//
	// 25
	NonCompliantResourceCnt *int32 `json:"NonCompliantResourceCnt,omitempty" xml:"NonCompliantResourceCnt,omitempty"`
	// The number of rules based on which resources are evaluated as non-compliant.
	//
	// example:
	//
	// 3
	NonCompliantRuleCnt *int32 `json:"NonCompliantRuleCnt,omitempty" xml:"NonCompliantRuleCnt,omitempty"`
	// The date on which the statistics are obtained.
	//
	// example:
	//
	// 2023-06-27
	StatisticDate *string `json:"StatisticDate,omitempty" xml:"StatisticDate,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 153
	TotalResourceCnt *int32 `json:"TotalResourceCnt,omitempty" xml:"TotalResourceCnt,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 10
	TotalRuleCnt *int32 `json:"TotalRuleCnt,omitempty" xml:"TotalRuleCnt,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetNonCompliantResourceCnt() *int32 {
	return s.NonCompliantResourceCnt
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetNonCompliantRuleCnt() *int32 {
	return s.NonCompliantRuleCnt
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetStatisticDate() *string {
	return s.StatisticDate
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetTotalResourceCnt() *int32 {
	return s.TotalResourceCnt
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetTotalRuleCnt() *int32 {
	return s.TotalRuleCnt
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetAggregatorId(v string) *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetNonCompliantResourceCnt(v int32) *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.NonCompliantResourceCnt = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetNonCompliantRuleCnt(v int32) *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.NonCompliantRuleCnt = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetStatisticDate(v string) *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.StatisticDate = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetTotalResourceCnt(v int32) *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.TotalResourceCnt = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetTotalRuleCnt(v int32) *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.TotalRuleCnt = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) Validate() error {
	return dara.Validate(s)
}
