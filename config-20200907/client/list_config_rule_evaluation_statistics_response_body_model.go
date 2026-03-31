// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRuleEvaluationStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationResults(v []*ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) *ListConfigRuleEvaluationStatisticsResponseBody
	GetEvaluationResults() []*ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults
	SetRequestId(v string) *ListConfigRuleEvaluationStatisticsResponseBody
	GetRequestId() *string
}

type ListConfigRuleEvaluationStatisticsResponseBody struct {
	// The statistics of compliance evaluation results.
	EvaluationResults []*ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4FF2B787-347E-5299-A196-2C0448DEA341
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigRuleEvaluationStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationStatisticsResponseBody) GetEvaluationResults() []*ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	return s.EvaluationResults
}

func (s *ListConfigRuleEvaluationStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigRuleEvaluationStatisticsResponseBody) SetEvaluationResults(v []*ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) *ListConfigRuleEvaluationStatisticsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponseBody) SetRequestId(v string) *ListConfigRuleEvaluationStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponseBody) Validate() error {
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

type ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults struct {
	// The number of resources that are evaluated as non-compliant.
	//
	// example:
	//
	// 30
	NonCompliantResourceCnt *int32 `json:"NonCompliantResourceCnt,omitempty" xml:"NonCompliantResourceCnt,omitempty"`
	// The number of rules based on which resources are evaluated as non-compliant.
	//
	// example:
	//
	// 5
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
	// 91
	TotalResourceCnt *int32 `json:"TotalResourceCnt,omitempty" xml:"TotalResourceCnt,omitempty"`
	// The total number of rules.
	//
	// example:
	//
	// 13
	TotalRuleCnt *int32 `json:"TotalRuleCnt,omitempty" xml:"TotalRuleCnt,omitempty"`
}

func (s ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetNonCompliantResourceCnt() *int32 {
	return s.NonCompliantResourceCnt
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetNonCompliantRuleCnt() *int32 {
	return s.NonCompliantRuleCnt
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetStatisticDate() *string {
	return s.StatisticDate
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetTotalResourceCnt() *int32 {
	return s.TotalResourceCnt
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) GetTotalRuleCnt() *int32 {
	return s.TotalRuleCnt
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetNonCompliantResourceCnt(v int32) *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.NonCompliantResourceCnt = &v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetNonCompliantRuleCnt(v int32) *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.NonCompliantRuleCnt = &v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetStatisticDate(v string) *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.StatisticDate = &v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetTotalResourceCnt(v int32) *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.TotalResourceCnt = &v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) SetTotalRuleCnt(v int32) *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults {
	s.TotalRuleCnt = &v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponseBodyEvaluationResults) Validate() error {
	return dara.Validate(s)
}
