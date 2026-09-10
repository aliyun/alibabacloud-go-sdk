// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckResult(v int32) *GetDataCheckTaskListRequest
	GetCheckResult() *int32
	SetCheckType(v int32) *GetDataCheckTaskListRequest
	GetCheckType() *int32
	SetCreateEndTime(v string) *GetDataCheckTaskListRequest
	GetCreateEndTime() *string
	SetCreateStartTime(v string) *GetDataCheckTaskListRequest
	GetCreateStartTime() *string
	SetExecStatus(v int32) *GetDataCheckTaskListRequest
	GetExecStatus() *int32
	SetIsScheduled(v int32) *GetDataCheckTaskListRequest
	GetIsScheduled() *int32
	SetPageIndex(v int32) *GetDataCheckTaskListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetDataCheckTaskListRequest
	GetPageSize() *int32
	SetTaskName(v string) *GetDataCheckTaskListRequest
	GetTaskName() *string
	SetTemplateName(v string) *GetDataCheckTaskListRequest
	GetTemplateName() *string
	SetUpdateEndTime(v string) *GetDataCheckTaskListRequest
	GetUpdateEndTime() *string
	SetUpdateStartTime(v string) *GetDataCheckTaskListRequest
	GetUpdateStartTime() *string
}

type GetDataCheckTaskListRequest struct {
	// The validation result filter. Valid values:
	//
	// - 0: no record.
	//
	// - 1: passed.
	//
	// - 2: failed.
	//
	// example:
	//
	// 0
	CheckResult *int32 `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// The validation type filter. Valid values:
	//
	// - 0: row count comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// example:
	//
	// 0
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The end of the creation time range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 10:00:00
	CreateEndTime *string `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	// The start of the creation time range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-16 00:00:00
	CreateStartTime *string `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	// The execution status filter. Valid values:
	//
	// - 0: pending.
	//
	// - 1: running.
	//
	// - 2: stopped.
	//
	// - 3: failed.
	//
	// - 4: completed.
	//
	// example:
	//
	// 0
	ExecStatus *int32 `json:"execStatus,omitempty" xml:"execStatus,omitempty"`
	// Specifies whether scheduling is enabled. Valid values:
	//
	// - 0: disabled.
	//
	// - 1: enabled.
	//
	// example:
	//
	// 0
	IsScheduled *int32 `json:"isScheduled,omitempty" xml:"isScheduled,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The task name. Fuzzy match is supported.
	//
	// example:
	//
	// data_check_task_demo
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The validation template name. Fuzzy match is supported. The server automatically converts the name into a list of template IDs for filtering.
	//
	// example:
	//
	// Row Count Validation Template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The end of the update time range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-14 13:59:03
	UpdateEndTime *string `json:"updateEndTime,omitempty" xml:"updateEndTime,omitempty"`
	// The start of the update time range. Format: YYYY-MM-DD HH:MM:SS.
	//
	// example:
	//
	// 2026-01-14 11:21:53
	UpdateStartTime *string `json:"updateStartTime,omitempty" xml:"updateStartTime,omitempty"`
}

func (s GetDataCheckTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTaskListRequest) GoString() string {
	return s.String()
}

func (s *GetDataCheckTaskListRequest) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *GetDataCheckTaskListRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckTaskListRequest) GetCreateEndTime() *string {
	return s.CreateEndTime
}

func (s *GetDataCheckTaskListRequest) GetCreateStartTime() *string {
	return s.CreateStartTime
}

func (s *GetDataCheckTaskListRequest) GetExecStatus() *int32 {
	return s.ExecStatus
}

func (s *GetDataCheckTaskListRequest) GetIsScheduled() *int32 {
	return s.IsScheduled
}

func (s *GetDataCheckTaskListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetDataCheckTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDataCheckTaskListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *GetDataCheckTaskListRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetDataCheckTaskListRequest) GetUpdateEndTime() *string {
	return s.UpdateEndTime
}

func (s *GetDataCheckTaskListRequest) GetUpdateStartTime() *string {
	return s.UpdateStartTime
}

func (s *GetDataCheckTaskListRequest) SetCheckResult(v int32) *GetDataCheckTaskListRequest {
	s.CheckResult = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetCheckType(v int32) *GetDataCheckTaskListRequest {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetCreateEndTime(v string) *GetDataCheckTaskListRequest {
	s.CreateEndTime = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetCreateStartTime(v string) *GetDataCheckTaskListRequest {
	s.CreateStartTime = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetExecStatus(v int32) *GetDataCheckTaskListRequest {
	s.ExecStatus = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetIsScheduled(v int32) *GetDataCheckTaskListRequest {
	s.IsScheduled = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetPageIndex(v int32) *GetDataCheckTaskListRequest {
	s.PageIndex = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetPageSize(v int32) *GetDataCheckTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetTaskName(v string) *GetDataCheckTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetTemplateName(v string) *GetDataCheckTaskListRequest {
	s.TemplateName = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetUpdateEndTime(v string) *GetDataCheckTaskListRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *GetDataCheckTaskListRequest) SetUpdateStartTime(v string) *GetDataCheckTaskListRequest {
	s.UpdateStartTime = &v
	return s
}

func (s *GetDataCheckTaskListRequest) Validate() error {
	return dara.Validate(s)
}
