// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalTimerRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeGlobalTimerRecordsResponseBody
	GetCount() *int32
	SetNextToken(v string) *DescribeGlobalTimerRecordsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeGlobalTimerRecordsResponseBody
	GetRequestId() *string
	SetResults(v []*DescribeGlobalTimerRecordsResponseBodyResults) *DescribeGlobalTimerRecordsResponseBody
	GetResults() []*DescribeGlobalTimerRecordsResponseBodyResults
}

type DescribeGlobalTimerRecordsResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6mnFXZiT7NdvGNgkInJ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 63740E03-1B4B-5A18-AC27-2745A4F2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response parameters.
	Results []*DescribeGlobalTimerRecordsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DescribeGlobalTimerRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerRecordsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeGlobalTimerRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGlobalTimerRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalTimerRecordsResponseBody) GetResults() []*DescribeGlobalTimerRecordsResponseBodyResults {
	return s.Results
}

func (s *DescribeGlobalTimerRecordsResponseBody) SetCount(v int32) *DescribeGlobalTimerRecordsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBody) SetNextToken(v string) *DescribeGlobalTimerRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBody) SetRequestId(v string) *DescribeGlobalTimerRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBody) SetResults(v []*DescribeGlobalTimerRecordsResponseBodyResults) *DescribeGlobalTimerRecordsResponseBody {
	s.Results = v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBody) Validate() error {
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

type DescribeGlobalTimerRecordsResponseBodyResults struct {
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The ID of the batch in which the scheduled task is executed.
	//
	// example:
	//
	// ccg-0cvfvf6u1enx1****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The time when the execution record was created.
	//
	// example:
	//
	// 2023-08-03T08:27:29Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-0c951fy9arnk9****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud computer name.
	//
	// example:
	//
	// DesktopName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The time when the scheduled task ended.
	//
	// example:
	//
	// 2025-01-21T02:00:45Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scheduled task group.
	//
	// example:
	//
	// ccg-xxxx
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	// The execution result of the scheduled task.
	//
	// example:
	//
	// RUNNING
	TimerResult *string `json:"TimerResult,omitempty" xml:"TimerResult,omitempty"`
	// The type of the scheduled task.
	//
	// example:
	//
	// TimerBoot
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeGlobalTimerRecordsResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerRecordsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetActionType() *string {
	return s.ActionType
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetBatchId() *string {
	return s.BatchId
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetContext() *string {
	return s.Context
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetTimerGroupId() *string {
	return s.TimerGroupId
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetTimerResult() *string {
	return s.TimerResult
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) GetTimerType() *string {
	return s.TimerType
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetActionType(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.ActionType = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetBatchId(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.BatchId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetContext(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.Context = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetCreateTime(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.CreateTime = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetDesktopId(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.DesktopId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetDesktopName(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetFinishTime(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.FinishTime = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetRegionId(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetTimerGroupId(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.TimerGroupId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetTimerResult(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.TimerResult = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) SetTimerType(v string) *DescribeGlobalTimerRecordsResponseBodyResults {
	s.TimerType = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
