// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalTimerBatchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeGlobalTimerBatchesResponseBody
	GetCount() *int32
	SetNextToken(v string) *DescribeGlobalTimerBatchesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeGlobalTimerBatchesResponseBody
	GetRequestId() *string
	SetResults(v []*DescribeGlobalTimerBatchesResponseBodyResults) *DescribeGlobalTimerBatchesResponseBody
	GetResults() []*DescribeGlobalTimerBatchesResponseBodyResults
}

type DescribeGlobalTimerBatchesResponseBody struct {
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E54EB497-D7B7-5F04-B744-D8DFA7B******
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DescribeGlobalTimerBatchesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DescribeGlobalTimerBatchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerBatchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerBatchesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeGlobalTimerBatchesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGlobalTimerBatchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalTimerBatchesResponseBody) GetResults() []*DescribeGlobalTimerBatchesResponseBodyResults {
	return s.Results
}

func (s *DescribeGlobalTimerBatchesResponseBody) SetCount(v int32) *DescribeGlobalTimerBatchesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBody) SetNextToken(v string) *DescribeGlobalTimerBatchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBody) SetRequestId(v string) *DescribeGlobalTimerBatchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBody) SetResults(v []*DescribeGlobalTimerBatchesResponseBodyResults) *DescribeGlobalTimerBatchesResponseBody {
	s.Results = v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalTimerBatchesResponseBodyResults struct {
	// example:
	//
	// ccg-0cvfvf6u1enx1****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// 2023-08-03T08:27:29Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// example:
	//
	// 0
	RunningCount *int32 `json:"RunningCount,omitempty" xml:"RunningCount,omitempty"`
	// example:
	//
	// 0
	SkippedCount *int32 `json:"SkippedCount,omitempty" xml:"SkippedCount,omitempty"`
	// example:
	//
	// 0
	SucceedCount *int32 `json:"SucceedCount,omitempty" xml:"SucceedCount,omitempty"`
	// example:
	//
	// TimerBoot
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeGlobalTimerBatchesResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerBatchesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) GetBatchId() *string {
	return s.BatchId
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) GetRunningCount() *int32 {
	return s.RunningCount
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) GetSkippedCount() *int32 {
	return s.SkippedCount
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) GetSucceedCount() *int32 {
	return s.SucceedCount
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) GetTimerType() *string {
	return s.TimerType
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) SetBatchId(v string) *DescribeGlobalTimerBatchesResponseBodyResults {
	s.BatchId = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) SetCreateTime(v string) *DescribeGlobalTimerBatchesResponseBodyResults {
	s.CreateTime = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) SetFailedCount(v int32) *DescribeGlobalTimerBatchesResponseBodyResults {
	s.FailedCount = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) SetRunningCount(v int32) *DescribeGlobalTimerBatchesResponseBodyResults {
	s.RunningCount = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) SetSkippedCount(v int32) *DescribeGlobalTimerBatchesResponseBodyResults {
	s.SkippedCount = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) SetSucceedCount(v int32) *DescribeGlobalTimerBatchesResponseBodyResults {
	s.SucceedCount = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) SetTimerType(v string) *DescribeGlobalTimerBatchesResponseBodyResults {
	s.TimerType = &v
	return s
}

func (s *DescribeGlobalTimerBatchesResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
