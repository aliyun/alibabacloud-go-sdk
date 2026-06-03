// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActualTimeGte(v int64) *CreateTaskExportTaskRequest
	GetActualTimeGte() *int64
	SetActualTimeLte(v int64) *CreateTaskExportTaskRequest
	GetActualTimeLte() *int64
	SetCallDurationGte(v int64) *CreateTaskExportTaskRequest
	GetCallDurationGte() *int64
	SetCallDurationLte(v int64) *CreateTaskExportTaskRequest
	GetCallDurationLte() *int64
	SetCalledNumber(v string) *CreateTaskExportTaskRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *CreateTaskExportTaskRequest
	GetCallingNumber() *string
	SetHasAnswered(v bool) *CreateTaskExportTaskRequest
	GetHasAnswered() *bool
	SetHasHangUpByRejection(v bool) *CreateTaskExportTaskRequest
	GetHasHangUpByRejection() *bool
	SetHasReachedEndOfFlow(v bool) *CreateTaskExportTaskRequest
	GetHasReachedEndOfFlow() *bool
	SetInstanceId(v string) *CreateTaskExportTaskRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *CreateTaskExportTaskRequest
	GetJobGroupId() *string
	SetJobGroupNameQuery(v string) *CreateTaskExportTaskRequest
	GetJobGroupNameQuery() *string
	SetJobId(v string) *CreateTaskExportTaskRequest
	GetJobId() *string
	SetJobStatusStringList(v string) *CreateTaskExportTaskRequest
	GetJobStatusStringList() *string
	SetOtherId(v string) *CreateTaskExportTaskRequest
	GetOtherId() *string
	SetPageIndex(v int32) *CreateTaskExportTaskRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *CreateTaskExportTaskRequest
	GetPageSize() *int32
	SetRecordingDurationGte(v int64) *CreateTaskExportTaskRequest
	GetRecordingDurationGte() *int64
	SetRecordingDurationLte(v int64) *CreateTaskExportTaskRequest
	GetRecordingDurationLte() *int64
	SetScriptNameQuery(v string) *CreateTaskExportTaskRequest
	GetScriptNameQuery() *string
	SetSortBy(v string) *CreateTaskExportTaskRequest
	GetSortBy() *string
	SetSortOrder(v string) *CreateTaskExportTaskRequest
	GetSortOrder() *string
	SetTaskCreateTimeGte(v int64) *CreateTaskExportTaskRequest
	GetTaskCreateTimeGte() *int64
	SetTaskCreateTimeLte(v int64) *CreateTaskExportTaskRequest
	GetTaskCreateTimeLte() *int64
	SetTaskId(v string) *CreateTaskExportTaskRequest
	GetTaskId() *string
	SetTaskStatusStringList(v string) *CreateTaskExportTaskRequest
	GetTaskStatusStringList() *string
	SetUserIdMatch(v string) *CreateTaskExportTaskRequest
	GetUserIdMatch() *string
}

type CreateTaskExportTaskRequest struct {
	// example:
	//
	// 1646496000000
	ActualTimeGte *int64 `json:"ActualTimeGte,omitempty" xml:"ActualTimeGte,omitempty"`
	// example:
	//
	// 1646582400000
	ActualTimeLte *int64 `json:"ActualTimeLte,omitempty" xml:"ActualTimeLte,omitempty"`
	// example:
	//
	// 10
	CallDurationGte *int64 `json:"CallDurationGte,omitempty" xml:"CallDurationGte,omitempty"`
	// example:
	//
	// 20
	CallDurationLte *int64 `json:"CallDurationLte,omitempty" xml:"CallDurationLte,omitempty"`
	// example:
	//
	// 11111111111
	CalledNumber  *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// true
	HasAnswered *bool `json:"HasAnswered,omitempty" xml:"HasAnswered,omitempty"`
	// example:
	//
	// true
	HasHangUpByRejection *bool `json:"HasHangUpByRejection,omitempty" xml:"HasHangUpByRejection,omitempty"`
	// example:
	//
	// true
	HasReachedEndOfFlow *bool `json:"HasReachedEndOfFlow,omitempty" xml:"HasReachedEndOfFlow,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1dcb09c5-d5db-4397-bf65-db854463beea
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cb731aee-0a5b-4c2b-924c-d9e82eb1d8d7
	JobGroupId        *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	JobGroupNameQuery *string `json:"JobGroupNameQuery,omitempty" xml:"JobGroupNameQuery,omitempty"`
	// example:
	//
	// 82097dd5-54df-475f-beba-eec8f4b7a3e1
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// Succeeded
	JobStatusStringList *string `json:"JobStatusStringList,omitempty" xml:"JobStatusStringList,omitempty"`
	// example:
	//
	// 64ebe700-91b4-49cb-b457-0b7c0b598a86
	OtherId *string `json:"OtherId,omitempty" xml:"OtherId,omitempty"`
	// example:
	//
	// 0
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	RecordingDurationGte *int64 `json:"RecordingDurationGte,omitempty" xml:"RecordingDurationGte,omitempty"`
	// example:
	//
	// 20
	RecordingDurationLte *int64  `json:"RecordingDurationLte,omitempty" xml:"RecordingDurationLte,omitempty"`
	ScriptNameQuery      *string `json:"ScriptNameQuery,omitempty" xml:"ScriptNameQuery,omitempty"`
	// example:
	//
	// actualTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// example:
	//
	// 1646496000000
	TaskCreateTimeGte *int64 `json:"TaskCreateTimeGte,omitempty" xml:"TaskCreateTimeGte,omitempty"`
	// example:
	//
	// 1646582400000
	TaskCreateTimeLte *int64 `json:"TaskCreateTimeLte,omitempty" xml:"TaskCreateTimeLte,omitempty"`
	// example:
	//
	// 64ebe700-91b4-49cb-b457-0b7c0b598a86
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// Succeeded,Failed
	TaskStatusStringList *string `json:"TaskStatusStringList,omitempty" xml:"TaskStatusStringList,omitempty"`
	// example:
	//
	// 82097dd5-54df-475f-beba-eec8f4b7a3e1
	UserIdMatch *string `json:"UserIdMatch,omitempty" xml:"UserIdMatch,omitempty"`
}

func (s CreateTaskExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskExportTaskRequest) GetActualTimeGte() *int64 {
	return s.ActualTimeGte
}

func (s *CreateTaskExportTaskRequest) GetActualTimeLte() *int64 {
	return s.ActualTimeLte
}

func (s *CreateTaskExportTaskRequest) GetCallDurationGte() *int64 {
	return s.CallDurationGte
}

func (s *CreateTaskExportTaskRequest) GetCallDurationLte() *int64 {
	return s.CallDurationLte
}

func (s *CreateTaskExportTaskRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *CreateTaskExportTaskRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *CreateTaskExportTaskRequest) GetHasAnswered() *bool {
	return s.HasAnswered
}

func (s *CreateTaskExportTaskRequest) GetHasHangUpByRejection() *bool {
	return s.HasHangUpByRejection
}

func (s *CreateTaskExportTaskRequest) GetHasReachedEndOfFlow() *bool {
	return s.HasReachedEndOfFlow
}

func (s *CreateTaskExportTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTaskExportTaskRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *CreateTaskExportTaskRequest) GetJobGroupNameQuery() *string {
	return s.JobGroupNameQuery
}

func (s *CreateTaskExportTaskRequest) GetJobId() *string {
	return s.JobId
}

func (s *CreateTaskExportTaskRequest) GetJobStatusStringList() *string {
	return s.JobStatusStringList
}

func (s *CreateTaskExportTaskRequest) GetOtherId() *string {
	return s.OtherId
}

func (s *CreateTaskExportTaskRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *CreateTaskExportTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CreateTaskExportTaskRequest) GetRecordingDurationGte() *int64 {
	return s.RecordingDurationGte
}

func (s *CreateTaskExportTaskRequest) GetRecordingDurationLte() *int64 {
	return s.RecordingDurationLte
}

func (s *CreateTaskExportTaskRequest) GetScriptNameQuery() *string {
	return s.ScriptNameQuery
}

func (s *CreateTaskExportTaskRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *CreateTaskExportTaskRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *CreateTaskExportTaskRequest) GetTaskCreateTimeGte() *int64 {
	return s.TaskCreateTimeGte
}

func (s *CreateTaskExportTaskRequest) GetTaskCreateTimeLte() *int64 {
	return s.TaskCreateTimeLte
}

func (s *CreateTaskExportTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTaskExportTaskRequest) GetTaskStatusStringList() *string {
	return s.TaskStatusStringList
}

func (s *CreateTaskExportTaskRequest) GetUserIdMatch() *string {
	return s.UserIdMatch
}

func (s *CreateTaskExportTaskRequest) SetActualTimeGte(v int64) *CreateTaskExportTaskRequest {
	s.ActualTimeGte = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetActualTimeLte(v int64) *CreateTaskExportTaskRequest {
	s.ActualTimeLte = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetCallDurationGte(v int64) *CreateTaskExportTaskRequest {
	s.CallDurationGte = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetCallDurationLte(v int64) *CreateTaskExportTaskRequest {
	s.CallDurationLte = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetCalledNumber(v string) *CreateTaskExportTaskRequest {
	s.CalledNumber = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetCallingNumber(v string) *CreateTaskExportTaskRequest {
	s.CallingNumber = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetHasAnswered(v bool) *CreateTaskExportTaskRequest {
	s.HasAnswered = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetHasHangUpByRejection(v bool) *CreateTaskExportTaskRequest {
	s.HasHangUpByRejection = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetHasReachedEndOfFlow(v bool) *CreateTaskExportTaskRequest {
	s.HasReachedEndOfFlow = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetInstanceId(v string) *CreateTaskExportTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetJobGroupId(v string) *CreateTaskExportTaskRequest {
	s.JobGroupId = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetJobGroupNameQuery(v string) *CreateTaskExportTaskRequest {
	s.JobGroupNameQuery = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetJobId(v string) *CreateTaskExportTaskRequest {
	s.JobId = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetJobStatusStringList(v string) *CreateTaskExportTaskRequest {
	s.JobStatusStringList = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetOtherId(v string) *CreateTaskExportTaskRequest {
	s.OtherId = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetPageIndex(v int32) *CreateTaskExportTaskRequest {
	s.PageIndex = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetPageSize(v int32) *CreateTaskExportTaskRequest {
	s.PageSize = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetRecordingDurationGte(v int64) *CreateTaskExportTaskRequest {
	s.RecordingDurationGte = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetRecordingDurationLte(v int64) *CreateTaskExportTaskRequest {
	s.RecordingDurationLte = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetScriptNameQuery(v string) *CreateTaskExportTaskRequest {
	s.ScriptNameQuery = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetSortBy(v string) *CreateTaskExportTaskRequest {
	s.SortBy = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetSortOrder(v string) *CreateTaskExportTaskRequest {
	s.SortOrder = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetTaskCreateTimeGte(v int64) *CreateTaskExportTaskRequest {
	s.TaskCreateTimeGte = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetTaskCreateTimeLte(v int64) *CreateTaskExportTaskRequest {
	s.TaskCreateTimeLte = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetTaskId(v string) *CreateTaskExportTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetTaskStatusStringList(v string) *CreateTaskExportTaskRequest {
	s.TaskStatusStringList = &v
	return s
}

func (s *CreateTaskExportTaskRequest) SetUserIdMatch(v string) *CreateTaskExportTaskRequest {
	s.UserIdMatch = &v
	return s
}

func (s *CreateTaskExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
