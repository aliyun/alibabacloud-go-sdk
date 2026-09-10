// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckTaskHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *ListDataCheckTaskHistoryRequest
	GetBatchId() *int64
	SetCheckResult(v int32) *ListDataCheckTaskHistoryRequest
	GetCheckResult() *int32
	SetCreateEndTime(v string) *ListDataCheckTaskHistoryRequest
	GetCreateEndTime() *string
	SetCreateStartTime(v string) *ListDataCheckTaskHistoryRequest
	GetCreateStartTime() *string
	SetExecEndTime(v string) *ListDataCheckTaskHistoryRequest
	GetExecEndTime() *string
	SetExecStartTime(v string) *ListDataCheckTaskHistoryRequest
	GetExecStartTime() *string
	SetExecStatus(v int32) *ListDataCheckTaskHistoryRequest
	GetExecStatus() *int32
	SetFinishEndTime(v string) *ListDataCheckTaskHistoryRequest
	GetFinishEndTime() *string
	SetFinishStartTime(v string) *ListDataCheckTaskHistoryRequest
	GetFinishStartTime() *string
	SetPageIndex(v int32) *ListDataCheckTaskHistoryRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckTaskHistoryRequest
	GetPageSize() *int32
	SetTaskId(v int64) *ListDataCheckTaskHistoryRequest
	GetTaskId() *int64
}

type ListDataCheckTaskHistoryRequest struct {
	// The ID of the validation job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20001
	BatchId *int64 `json:"batchId,omitempty" xml:"batchId,omitempty"`
	// Filters by validation result. Valid values:
	//
	// - 0: No records.
	//
	// - 1: Passed.
	//
	// - 2: Failed.
	//
	// example:
	//
	// 0
	CheckResult *int32 `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// The end of the job creation time filter range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	CreateEndTime *string `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	// The start of the job creation time filter range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 00:00:00
	CreateStartTime *string `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	// The end of the execution start time filter range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 12:00:00
	ExecEndTime *string `json:"execEndTime,omitempty" xml:"execEndTime,omitempty"`
	// The start of the execution start time filter range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	ExecStartTime *string `json:"execStartTime,omitempty" xml:"execStartTime,omitempty"`
	// Filters by execution status. Valid values:
	//
	// - 0: Pending.
	//
	// - 1: Running.
	//
	// - 2: Stopped.
	//
	// - 3: Failed.
	//
	// - 4: Completed.
	//
	// example:
	//
	// 0
	ExecStatus *int32 `json:"execStatus,omitempty" xml:"execStatus,omitempty"`
	// The end of the execution end time filter range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 12:30:00
	FinishEndTime *string `json:"finishEndTime,omitempty" xml:"finishEndTime,omitempty"`
	// The start of the execution end time filter range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 10:30:00
	FinishStartTime *string `json:"finishStartTime,omitempty" xml:"finishStartTime,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the data validation task.
	//
	// example:
	//
	// 1001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListDataCheckTaskHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckTaskHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListDataCheckTaskHistoryRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *ListDataCheckTaskHistoryRequest) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *ListDataCheckTaskHistoryRequest) GetCreateEndTime() *string {
	return s.CreateEndTime
}

func (s *ListDataCheckTaskHistoryRequest) GetCreateStartTime() *string {
	return s.CreateStartTime
}

func (s *ListDataCheckTaskHistoryRequest) GetExecEndTime() *string {
	return s.ExecEndTime
}

func (s *ListDataCheckTaskHistoryRequest) GetExecStartTime() *string {
	return s.ExecStartTime
}

func (s *ListDataCheckTaskHistoryRequest) GetExecStatus() *int32 {
	return s.ExecStatus
}

func (s *ListDataCheckTaskHistoryRequest) GetFinishEndTime() *string {
	return s.FinishEndTime
}

func (s *ListDataCheckTaskHistoryRequest) GetFinishStartTime() *string {
	return s.FinishStartTime
}

func (s *ListDataCheckTaskHistoryRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckTaskHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckTaskHistoryRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListDataCheckTaskHistoryRequest) SetBatchId(v int64) *ListDataCheckTaskHistoryRequest {
	s.BatchId = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetCheckResult(v int32) *ListDataCheckTaskHistoryRequest {
	s.CheckResult = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetCreateEndTime(v string) *ListDataCheckTaskHistoryRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetCreateStartTime(v string) *ListDataCheckTaskHistoryRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetExecEndTime(v string) *ListDataCheckTaskHistoryRequest {
	s.ExecEndTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetExecStartTime(v string) *ListDataCheckTaskHistoryRequest {
	s.ExecStartTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetExecStatus(v int32) *ListDataCheckTaskHistoryRequest {
	s.ExecStatus = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetFinishEndTime(v string) *ListDataCheckTaskHistoryRequest {
	s.FinishEndTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetFinishStartTime(v string) *ListDataCheckTaskHistoryRequest {
	s.FinishStartTime = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetPageIndex(v int32) *ListDataCheckTaskHistoryRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetPageSize(v int32) *ListDataCheckTaskHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) SetTaskId(v int64) *ListDataCheckTaskHistoryRequest {
	s.TaskId = &v
	return s
}

func (s *ListDataCheckTaskHistoryRequest) Validate() error {
	return dara.Validate(s)
}
