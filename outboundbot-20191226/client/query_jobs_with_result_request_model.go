// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryJobsWithResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndActualTimeFilter(v int64) *QueryJobsWithResultRequest
	GetEndActualTimeFilter() *int64
	SetHasAnsweredFilter(v bool) *QueryJobsWithResultRequest
	GetHasAnsweredFilter() *bool
	SetHasHangUpByRejectionFilter(v bool) *QueryJobsWithResultRequest
	GetHasHangUpByRejectionFilter() *bool
	SetHasReachedEndOfFlowFilter(v bool) *QueryJobsWithResultRequest
	GetHasReachedEndOfFlowFilter() *bool
	SetInstanceId(v string) *QueryJobsWithResultRequest
	GetInstanceId() *string
	SetJobFailureReasonsFilter(v string) *QueryJobsWithResultRequest
	GetJobFailureReasonsFilter() *string
	SetJobGroupId(v string) *QueryJobsWithResultRequest
	GetJobGroupId() *string
	SetJobStatusFilter(v string) *QueryJobsWithResultRequest
	GetJobStatusFilter() *string
	SetLabelsJson(v []*string) *QueryJobsWithResultRequest
	GetLabelsJson() []*string
	SetPageNumber(v int32) *QueryJobsWithResultRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryJobsWithResultRequest
	GetPageSize() *int32
	SetQueryText(v string) *QueryJobsWithResultRequest
	GetQueryText() *string
	SetStartActualTimeFilter(v int64) *QueryJobsWithResultRequest
	GetStartActualTimeFilter() *int64
	SetTaskStatusFilter(v string) *QueryJobsWithResultRequest
	GetTaskStatusFilter() *string
}

type QueryJobsWithResultRequest struct {
	EndActualTimeFilter *int64 `json:"EndActualTimeFilter,omitempty" xml:"EndActualTimeFilter,omitempty"`
	// example:
	//
	// true
	HasAnsweredFilter *bool `json:"HasAnsweredFilter,omitempty" xml:"HasAnsweredFilter,omitempty"`
	// example:
	//
	// false
	HasHangUpByRejectionFilter *bool `json:"HasHangUpByRejectionFilter,omitempty" xml:"HasHangUpByRejectionFilter,omitempty"`
	// example:
	//
	// true
	HasReachedEndOfFlowFilter *bool `json:"HasReachedEndOfFlowFilter,omitempty" xml:"HasReachedEndOfFlowFilter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9d53cd72-4050-4419-8c17-acc0bf158147
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["NoAnswer"]
	JobFailureReasonsFilter *string `json:"JobFailureReasonsFilter,omitempty" xml:"JobFailureReasonsFilter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ad16fc35-d824-4102-a606-2be51c1aa6dd
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// example:
	//
	// Succeeded
	JobStatusFilter *string   `json:"JobStatusFilter,omitempty" xml:"JobStatusFilter,omitempty"`
	LabelsJson      []*string `json:"LabelsJson,omitempty" xml:"LabelsJson,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1882020****
	QueryText             *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	StartActualTimeFilter *int64  `json:"StartActualTimeFilter,omitempty" xml:"StartActualTimeFilter,omitempty"`
	// example:
	//
	// Succeeded
	TaskStatusFilter *string `json:"TaskStatusFilter,omitempty" xml:"TaskStatusFilter,omitempty"`
}

func (s QueryJobsWithResultRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryJobsWithResultRequest) GoString() string {
	return s.String()
}

func (s *QueryJobsWithResultRequest) GetEndActualTimeFilter() *int64 {
	return s.EndActualTimeFilter
}

func (s *QueryJobsWithResultRequest) GetHasAnsweredFilter() *bool {
	return s.HasAnsweredFilter
}

func (s *QueryJobsWithResultRequest) GetHasHangUpByRejectionFilter() *bool {
	return s.HasHangUpByRejectionFilter
}

func (s *QueryJobsWithResultRequest) GetHasReachedEndOfFlowFilter() *bool {
	return s.HasReachedEndOfFlowFilter
}

func (s *QueryJobsWithResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryJobsWithResultRequest) GetJobFailureReasonsFilter() *string {
	return s.JobFailureReasonsFilter
}

func (s *QueryJobsWithResultRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *QueryJobsWithResultRequest) GetJobStatusFilter() *string {
	return s.JobStatusFilter
}

func (s *QueryJobsWithResultRequest) GetLabelsJson() []*string {
	return s.LabelsJson
}

func (s *QueryJobsWithResultRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryJobsWithResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryJobsWithResultRequest) GetQueryText() *string {
	return s.QueryText
}

func (s *QueryJobsWithResultRequest) GetStartActualTimeFilter() *int64 {
	return s.StartActualTimeFilter
}

func (s *QueryJobsWithResultRequest) GetTaskStatusFilter() *string {
	return s.TaskStatusFilter
}

func (s *QueryJobsWithResultRequest) SetEndActualTimeFilter(v int64) *QueryJobsWithResultRequest {
	s.EndActualTimeFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetHasAnsweredFilter(v bool) *QueryJobsWithResultRequest {
	s.HasAnsweredFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetHasHangUpByRejectionFilter(v bool) *QueryJobsWithResultRequest {
	s.HasHangUpByRejectionFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetHasReachedEndOfFlowFilter(v bool) *QueryJobsWithResultRequest {
	s.HasReachedEndOfFlowFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetInstanceId(v string) *QueryJobsWithResultRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetJobFailureReasonsFilter(v string) *QueryJobsWithResultRequest {
	s.JobFailureReasonsFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetJobGroupId(v string) *QueryJobsWithResultRequest {
	s.JobGroupId = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetJobStatusFilter(v string) *QueryJobsWithResultRequest {
	s.JobStatusFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetLabelsJson(v []*string) *QueryJobsWithResultRequest {
	s.LabelsJson = v
	return s
}

func (s *QueryJobsWithResultRequest) SetPageNumber(v int32) *QueryJobsWithResultRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetPageSize(v int32) *QueryJobsWithResultRequest {
	s.PageSize = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetQueryText(v string) *QueryJobsWithResultRequest {
	s.QueryText = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetStartActualTimeFilter(v int64) *QueryJobsWithResultRequest {
	s.StartActualTimeFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) SetTaskStatusFilter(v string) *QueryJobsWithResultRequest {
	s.TaskStatusFilter = &v
	return s
}

func (s *QueryJobsWithResultRequest) Validate() error {
	return dara.Validate(s)
}
