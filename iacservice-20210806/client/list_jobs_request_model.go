// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobType(v string) *ListJobsRequest
	GetJobType() *string
	SetPageNumber(v int32) *ListJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListJobsRequest
	GetStatus() *string
	SetTaskType(v string) *ListJobsRequest
	GetTaskType() *string
}

type ListJobsRequest struct {
	// example:
	//
	// Default
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Errored
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s ListJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListJobsRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListJobsRequest) SetJobType(v string) *ListJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) SetTaskType(v string) *ListJobsRequest {
	s.TaskType = &v
	return s
}

func (s *ListJobsRequest) Validate() error {
	return dara.Validate(s)
}
