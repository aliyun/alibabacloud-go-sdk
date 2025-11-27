// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobExecutorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ListJobExecutorsRequest
	GetJobId() *string
	SetPageNumber(v int32) *ListJobExecutorsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJobExecutorsRequest
	GetPageSize() *int32
	SetTaskName(v string) *ListJobExecutorsRequest
	GetTaskName() *string
}

type ListJobExecutorsRequest struct {
	// The ID of the job.
	//
	// example:
	//
	// job-xxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The job name.
	//
	// example:
	//
	// task0
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListJobExecutorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutorsRequest) GoString() string {
	return s.String()
}

func (s *ListJobExecutorsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListJobExecutorsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobExecutorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobExecutorsRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListJobExecutorsRequest) SetJobId(v string) *ListJobExecutorsRequest {
	s.JobId = &v
	return s
}

func (s *ListJobExecutorsRequest) SetPageNumber(v int32) *ListJobExecutorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobExecutorsRequest) SetPageSize(v int32) *ListJobExecutorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutorsRequest) SetTaskName(v string) *ListJobExecutorsRequest {
	s.TaskName = &v
	return s
}

func (s *ListJobExecutorsRequest) Validate() error {
	return dara.Validate(s)
}
