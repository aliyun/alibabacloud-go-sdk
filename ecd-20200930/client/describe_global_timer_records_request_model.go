// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalTimerRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *DescribeGlobalTimerRecordsRequest
	GetBatchId() *string
	SetDesktopIds(v []*string) *DescribeGlobalTimerRecordsRequest
	GetDesktopIds() []*string
	SetGroupId(v string) *DescribeGlobalTimerRecordsRequest
	GetGroupId() *string
	SetMaxResults(v string) *DescribeGlobalTimerRecordsRequest
	GetMaxResults() *string
	SetNextToken(v string) *DescribeGlobalTimerRecordsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeGlobalTimerRecordsRequest
	GetRegionId() *string
	SetResultCategory(v string) *DescribeGlobalTimerRecordsRequest
	GetResultCategory() *string
	SetSearchRegionId(v string) *DescribeGlobalTimerRecordsRequest
	GetSearchRegionId() *string
	SetTimerResult(v string) *DescribeGlobalTimerRecordsRequest
	GetTimerResult() *string
	SetTimerTypes(v []*string) *DescribeGlobalTimerRecordsRequest
	GetTimerTypes() []*string
}

type DescribeGlobalTimerRecordsRequest struct {
	// The ID of the batch in which the scheduled task is executed.
	//
	// example:
	//
	// ccg-****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The cloud computer IDs.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The ID of the scheduled task group.
	//
	// example:
	//
	// ccg-0cvfvf6u1enx1****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the execution result. You can specify this parameter to filter the execution results.
	//
	// Valid values:
	//
	// 	- FAILED: The execution is successful.
	//
	// 	- FAILED: The execution failed.
	//
	// 	- RUNNING: The execution is in progress.
	//
	// 	- SKIPPED: The execution is skipped.
	//
	// example:
	//
	// SUCCEED
	ResultCategory *string `json:"ResultCategory,omitempty" xml:"ResultCategory,omitempty"`
	// The ID of the searched region. You can specify this parameter to filter cloud computers in specific regions.
	//
	// example:
	//
	// cn-shanghai
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// The execution result of the scheduled task.
	//
	// Valid values:
	//
	// 	- CONNECTED_NOT_RUN: The cloud computer is connected, but the scheduled task is not executed.
	//
	// 	- PAUSED: The scheduled task is suspended.
	//
	// 	- COMPLETED: The scheduled task is executed.
	//
	// 	- FAILED: The scheduled task failed to be executed.
	//
	// 	- RUNNING: The scheduled task is being executed.
	//
	// 	- TERMINATED: The scheduled task is stopped.
	//
	// example:
	//
	// RUNNING
	TimerResult *string `json:"TimerResult,omitempty" xml:"TimerResult,omitempty"`
	// The scheduled tasks.
	TimerTypes []*string `json:"TimerTypes,omitempty" xml:"TimerTypes,omitempty" type:"Repeated"`
}

func (s DescribeGlobalTimerRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerRecordsRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *DescribeGlobalTimerRecordsRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *DescribeGlobalTimerRecordsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGlobalTimerRecordsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *DescribeGlobalTimerRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeGlobalTimerRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalTimerRecordsRequest) GetResultCategory() *string {
	return s.ResultCategory
}

func (s *DescribeGlobalTimerRecordsRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeGlobalTimerRecordsRequest) GetTimerResult() *string {
	return s.TimerResult
}

func (s *DescribeGlobalTimerRecordsRequest) GetTimerTypes() []*string {
	return s.TimerTypes
}

func (s *DescribeGlobalTimerRecordsRequest) SetBatchId(v string) *DescribeGlobalTimerRecordsRequest {
	s.BatchId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetDesktopIds(v []*string) *DescribeGlobalTimerRecordsRequest {
	s.DesktopIds = v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetGroupId(v string) *DescribeGlobalTimerRecordsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetMaxResults(v string) *DescribeGlobalTimerRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetNextToken(v string) *DescribeGlobalTimerRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetRegionId(v string) *DescribeGlobalTimerRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetResultCategory(v string) *DescribeGlobalTimerRecordsRequest {
	s.ResultCategory = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetSearchRegionId(v string) *DescribeGlobalTimerRecordsRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetTimerResult(v string) *DescribeGlobalTimerRecordsRequest {
	s.TimerResult = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetTimerTypes(v []*string) *DescribeGlobalTimerRecordsRequest {
	s.TimerTypes = v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) Validate() error {
	return dara.Validate(s)
}
