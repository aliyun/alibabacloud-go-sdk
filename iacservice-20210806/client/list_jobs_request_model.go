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
	// The job type.
	//
	// example:
	//
	// Default
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The job status. Valid values:
	//
	// - Pending: The initial status after the job is created.
	//
	// - PlanQueued: After the job is created, if no workflow is available, the job is queued.
	//
	// - Planning: The resource job is in the Plan execution phase.
	//
	// - ConfigProactiveInProgress: Compliance pre-check is in progress. The compliance pre-check feature must be enabled for the account.
	//
	// - ConfigProactiveSuccess: Compliance pre-check succeeded. The compliance pre-check feature must be enabled for the account.
	//
	// - Planned: The resource job has completed the Plan execution.
	//
	// - PlannedAndFinished: After the Plan execution is complete, no diff is found. This is a final status.
	//
	// - Confirmed: The resource job is waiting for confirmation after the Plan execution is complete.
	//
	// - ApplyQueued: During job execution, if no workflow is available, the job is queued.
	//
	// - Applying: The resource job is in the Apply execution phase.
	//
	// - Applied: The resource job has completed the Apply execution. This is a final status.
	//
	// - Errored: The job execution encountered an error. This is a final status.
	//
	// - Canceled: The job execution was canceled. This is a final status.
	//
	// - Discarded: The plan of the resource job was discarded. This is a final status.
	//
	// - ConfigProactiveFailure: Compliance pre-check failed. The compliance pre-check feature must be enabled for the account.
	//
	// example:
	//
	// Errored
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task type. Valid values:
	//
	// - Task: regular task. This is the default value.
	//
	// - SceneTestingTask: scenario-based testing task.
	//
	// example:
	//
	// SceneTestingTask
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
