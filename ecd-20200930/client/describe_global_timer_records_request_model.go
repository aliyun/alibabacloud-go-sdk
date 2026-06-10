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
	SetDisplayResultName(v string) *DescribeGlobalTimerRecordsRequest
	GetDisplayResultName() *string
	SetGroupId(v string) *DescribeGlobalTimerRecordsRequest
	GetGroupId() *string
	SetMaxResults(v string) *DescribeGlobalTimerRecordsRequest
	GetMaxResults() *string
	SetNextToken(v string) *DescribeGlobalTimerRecordsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeGlobalTimerRecordsRequest
	GetRegionId() *string
	SetResourceTypes(v []*string) *DescribeGlobalTimerRecordsRequest
	GetResourceTypes() []*string
	SetResultCategory(v string) *DescribeGlobalTimerRecordsRequest
	GetResultCategory() *string
	SetRetryable(v bool) *DescribeGlobalTimerRecordsRequest
	GetRetryable() *bool
	SetSearchRegionId(v string) *DescribeGlobalTimerRecordsRequest
	GetSearchRegionId() *string
	SetTimerResult(v string) *DescribeGlobalTimerRecordsRequest
	GetTimerResult() *string
	SetTimerTypes(v []*string) *DescribeGlobalTimerRecordsRequest
	GetTimerTypes() []*string
	SetWuyingServerIds(v []*string) *DescribeGlobalTimerRecordsRequest
	GetWuyingServerIds() []*string
}

type DescribeGlobalTimerRecordsRequest struct {
	// The batch ID for a scheduled task execution.
	//
	// example:
	//
	// ccg-****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// A list of cloud desktop IDs.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// example:
	//
	// FAILED
	DisplayResultName *string `json:"DisplayResultName,omitempty" xml:"DisplayResultName,omitempty"`
	// The scheduled task group ID.
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
	// The token used to start the next query.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to list the regions available in Elastic Desktop Service.
	//
	// example:
	//
	// cn-shanghai
	RegionId      *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// Filters the results by execution status. Valid values:
	//
	// example:
	//
	// SUCCEED
	ResultCategory *string `json:"ResultCategory,omitempty" xml:"ResultCategory,omitempty"`
	// example:
	//
	// true
	Retryable *bool `json:"Retryable,omitempty" xml:"Retryable,omitempty"`
	// The ID of the region to filter by. Only records for cloud desktops in this region are returned.
	//
	// example:
	//
	// cn-shanghai
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	// The execution result of the scheduled task. Valid values:
	//
	// example:
	//
	// RUNNING
	TimerResult *string `json:"TimerResult,omitempty" xml:"TimerResult,omitempty"`
	// The types of scheduled tasks.
	TimerTypes      []*string `json:"TimerTypes,omitempty" xml:"TimerTypes,omitempty" type:"Repeated"`
	WuyingServerIds []*string `json:"WuyingServerIds,omitempty" xml:"WuyingServerIds,omitempty" type:"Repeated"`
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

func (s *DescribeGlobalTimerRecordsRequest) GetDisplayResultName() *string {
	return s.DisplayResultName
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

func (s *DescribeGlobalTimerRecordsRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *DescribeGlobalTimerRecordsRequest) GetResultCategory() *string {
	return s.ResultCategory
}

func (s *DescribeGlobalTimerRecordsRequest) GetRetryable() *bool {
	return s.Retryable
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

func (s *DescribeGlobalTimerRecordsRequest) GetWuyingServerIds() []*string {
	return s.WuyingServerIds
}

func (s *DescribeGlobalTimerRecordsRequest) SetBatchId(v string) *DescribeGlobalTimerRecordsRequest {
	s.BatchId = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetDesktopIds(v []*string) *DescribeGlobalTimerRecordsRequest {
	s.DesktopIds = v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetDisplayResultName(v string) *DescribeGlobalTimerRecordsRequest {
	s.DisplayResultName = &v
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

func (s *DescribeGlobalTimerRecordsRequest) SetResourceTypes(v []*string) *DescribeGlobalTimerRecordsRequest {
	s.ResourceTypes = v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetResultCategory(v string) *DescribeGlobalTimerRecordsRequest {
	s.ResultCategory = &v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) SetRetryable(v bool) *DescribeGlobalTimerRecordsRequest {
	s.Retryable = &v
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

func (s *DescribeGlobalTimerRecordsRequest) SetWuyingServerIds(v []*string) *DescribeGlobalTimerRecordsRequest {
	s.WuyingServerIds = v
	return s
}

func (s *DescribeGlobalTimerRecordsRequest) Validate() error {
	return dara.Validate(s)
}
