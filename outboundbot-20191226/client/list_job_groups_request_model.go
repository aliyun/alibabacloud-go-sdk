// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncQuery(v bool) *ListJobGroupsRequest
	GetAsyncQuery() *bool
	SetEndTime(v int64) *ListJobGroupsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListJobGroupsRequest
	GetInstanceId() *string
	SetJobGroupStatusFilter(v string) *ListJobGroupsRequest
	GetJobGroupStatusFilter() *string
	SetOnlyMinConcurrencyEnabled(v bool) *ListJobGroupsRequest
	GetOnlyMinConcurrencyEnabled() *bool
	SetPageNumber(v int32) *ListJobGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobGroupsRequest
	GetPageSize() *int32
	SetSearchText(v string) *ListJobGroupsRequest
	GetSearchText() *string
	SetStartTime(v int64) *ListJobGroupsRequest
	GetStartTime() *int64
}

type ListJobGroupsRequest struct {
	// Whether to asynchronously retrieve the query result.
	//
	// example:
	//
	// true
	AsyncQuery *bool `json:"AsyncQuery,omitempty" xml:"AsyncQuery,omitempty"`
	// Filter by the end of task group creation time.
	//
	// example:
	//
	// 1579965079000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Business ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Filter condition: Filter by job status. You can select multiple statuses simultaneously, separated by commas, with a logical OR relationship.
	//
	// Valid values:
	//
	// - **Draft**: Draft.
	//
	// - **Scheduling**: Scheduling.
	//
	// - **Executing**: Executing.
	//
	// - **Completed**: Completed.
	//
	// - **Paused**: Suspended.
	//
	// - **Failed**: Failed.
	//
	// - **Cancelled**: Cancelled.
	//
	// - **Initializing**: Initializing.
	//
	// example:
	//
	// Draft
	JobGroupStatusFilter *string `json:"JobGroupStatusFilter,omitempty" xml:"JobGroupStatusFilter,omitempty"`
	// View jobs with effective concurrency guarantee values.
	//
	// example:
	//
	// false
	OnlyMinConcurrencyEnabled *bool `json:"OnlyMinConcurrencyEnabled,omitempty" xml:"OnlyMinConcurrencyEnabled,omitempty"`
	// Page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of records per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Search keyword
	//
	// example:
	//
	// 面试
	SearchText *string `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	// Filter by task group creation time (start).
	//
	// example:
	//
	// 1578965079000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListJobGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListJobGroupsRequest) GetAsyncQuery() *bool {
	return s.AsyncQuery
}

func (s *ListJobGroupsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListJobGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListJobGroupsRequest) GetJobGroupStatusFilter() *string {
	return s.JobGroupStatusFilter
}

func (s *ListJobGroupsRequest) GetOnlyMinConcurrencyEnabled() *bool {
	return s.OnlyMinConcurrencyEnabled
}

func (s *ListJobGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobGroupsRequest) GetSearchText() *string {
	return s.SearchText
}

func (s *ListJobGroupsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListJobGroupsRequest) SetAsyncQuery(v bool) *ListJobGroupsRequest {
	s.AsyncQuery = &v
	return s
}

func (s *ListJobGroupsRequest) SetEndTime(v int64) *ListJobGroupsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobGroupsRequest) SetInstanceId(v string) *ListJobGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListJobGroupsRequest) SetJobGroupStatusFilter(v string) *ListJobGroupsRequest {
	s.JobGroupStatusFilter = &v
	return s
}

func (s *ListJobGroupsRequest) SetOnlyMinConcurrencyEnabled(v bool) *ListJobGroupsRequest {
	s.OnlyMinConcurrencyEnabled = &v
	return s
}

func (s *ListJobGroupsRequest) SetPageNumber(v int32) *ListJobGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobGroupsRequest) SetPageSize(v int32) *ListJobGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobGroupsRequest) SetSearchText(v string) *ListJobGroupsRequest {
	s.SearchText = &v
	return s
}

func (s *ListJobGroupsRequest) SetStartTime(v int64) *ListJobGroupsRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobGroupsRequest) Validate() error {
	return dara.Validate(s)
}
