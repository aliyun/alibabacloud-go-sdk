// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalysisTask interface {
	dara.Model
	String() string
	GoString() string
	SetCompletedAt(v int64) *AnalysisTask
	GetCompletedAt() *int64
	SetCreatedAt(v int64) *AnalysisTask
	GetCreatedAt() *int64
	SetErrorCode(v string) *AnalysisTask
	GetErrorCode() *string
	SetErrorMessage(v string) *AnalysisTask
	GetErrorMessage() *string
	SetReport(v map[string]interface{}) *AnalysisTask
	GetReport() map[string]interface{}
	SetSelectedCount(v int32) *AnalysisTask
	GetSelectedCount() *int32
	SetStatus(v string) *AnalysisTask
	GetStatus() *string
	SetTaskId(v string) *AnalysisTask
	GetTaskId() *string
	SetTaskName(v string) *AnalysisTask
	GetTaskName() *string
	SetTaskType(v string) *AnalysisTask
	GetTaskType() *string
}

type AnalysisTask struct {
	// The task completion time as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1789747260000
	CompletedAt *int64 `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// The task creation time as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1789747200000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The error code for the task failure.
	//
	// example:
	//
	// NO_VALID_ATTRIBUTION_LABEL
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message for the task failure.
	//
	// example:
	//
	// No valid attribution label was found
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The generated analysis report. This field is typically not returned by list operations.
	Report map[string]interface{} `json:"report,omitempty" xml:"report,omitempty"`
	// The number of evaluation details selected by the user.
	//
	// example:
	//
	// 20
	SelectedCount *int32 `json:"selectedCount,omitempty" xml:"selectedCount,omitempty"`
	// The task status. Valid values: PENDING, RUNNING, SUCCEEDED, PARTIAL_SUCCEEDED, and FAILED.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The analysis task ID.
	//
	// example:
	//
	// analysis-task-8b36f2e2b1f94f9c91ce7a4b0f6d9c25
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The analysis task name.
	//
	// example:
	//
	// Low-score sample attribution analysis
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The analysis task type.
	//
	// example:
	//
	// ATTRIBUTION_SUMMARY
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s AnalysisTask) String() string {
	return dara.Prettify(s)
}

func (s AnalysisTask) GoString() string {
	return s.String()
}

func (s *AnalysisTask) GetCompletedAt() *int64 {
	return s.CompletedAt
}

func (s *AnalysisTask) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *AnalysisTask) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AnalysisTask) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AnalysisTask) GetReport() map[string]interface{} {
	return s.Report
}

func (s *AnalysisTask) GetSelectedCount() *int32 {
	return s.SelectedCount
}

func (s *AnalysisTask) GetStatus() *string {
	return s.Status
}

func (s *AnalysisTask) GetTaskId() *string {
	return s.TaskId
}

func (s *AnalysisTask) GetTaskName() *string {
	return s.TaskName
}

func (s *AnalysisTask) GetTaskType() *string {
	return s.TaskType
}

func (s *AnalysisTask) SetCompletedAt(v int64) *AnalysisTask {
	s.CompletedAt = &v
	return s
}

func (s *AnalysisTask) SetCreatedAt(v int64) *AnalysisTask {
	s.CreatedAt = &v
	return s
}

func (s *AnalysisTask) SetErrorCode(v string) *AnalysisTask {
	s.ErrorCode = &v
	return s
}

func (s *AnalysisTask) SetErrorMessage(v string) *AnalysisTask {
	s.ErrorMessage = &v
	return s
}

func (s *AnalysisTask) SetReport(v map[string]interface{}) *AnalysisTask {
	s.Report = v
	return s
}

func (s *AnalysisTask) SetSelectedCount(v int32) *AnalysisTask {
	s.SelectedCount = &v
	return s
}

func (s *AnalysisTask) SetStatus(v string) *AnalysisTask {
	s.Status = &v
	return s
}

func (s *AnalysisTask) SetTaskId(v string) *AnalysisTask {
	s.TaskId = &v
	return s
}

func (s *AnalysisTask) SetTaskName(v string) *AnalysisTask {
	s.TaskName = &v
	return s
}

func (s *AnalysisTask) SetTaskType(v string) *AnalysisTask {
	s.TaskType = &v
	return s
}

func (s *AnalysisTask) Validate() error {
	return dara.Validate(s)
}
