// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCopilotSessionResponseBody
	GetCode() *string
	SetData(v *CreateCopilotSessionResponseBodyData) *CreateCopilotSessionResponseBody
	GetData() *CreateCopilotSessionResponseBodyData
	SetMessage(v string) *CreateCopilotSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCopilotSessionResponseBody
	GetRequestId() *string
}

type CreateCopilotSessionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *CreateCopilotSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which can be used for end-to-end diagnostics.
	//
	// example:
	//
	// 5BC79EF5-798B-5499-9683-4E14EBBBF712
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateCopilotSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCopilotSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCopilotSessionResponseBody) GetData() *CreateCopilotSessionResponseBodyData {
	return s.Data
}

func (s *CreateCopilotSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCopilotSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCopilotSessionResponseBody) SetCode(v string) *CreateCopilotSessionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCopilotSessionResponseBody) SetData(v *CreateCopilotSessionResponseBodyData) *CreateCopilotSessionResponseBody {
	s.Data = v
	return s
}

func (s *CreateCopilotSessionResponseBody) SetMessage(v string) *CreateCopilotSessionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCopilotSessionResponseBody) SetRequestId(v string) *CreateCopilotSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCopilotSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCopilotSessionResponseBodyData struct {
	// The information about the current active execution.
	ActiveTurn *CreateCopilotSessionResponseBodyDataActiveTurn `json:"activeTurn,omitempty" xml:"activeTurn,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2026-10-29T10:00:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The cursor of the latest persisted event in the session. For a newly created session without events, 0 is returned. This field is a decimal string.
	//
	// example:
	//
	// 0
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// The information about the most recent execution.
	LastTurn *CreateCopilotSessionResponseBodyDataLastTurn `json:"lastTurn,omitempty" xml:"lastTurn,omitempty" type:"Struct"`
	// The session name. If not specified, the value is New Session.
	//
	// example:
	//
	// Troubleshoot CPU load
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The information about the associated scheduled task.
	ScheduledTask *CreateCopilotSessionResponseBodyDataScheduledTask `json:"scheduledTask,omitempty" xml:"scheduledTask,omitempty" type:"Struct"`
	// The session ID. Use this value to send messages, query, update, or delete the session.
	//
	// example:
	//
	// c356bc310e4345c3a602a2912322ad21
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The session source. Valid values:
	//
	// - openapi: proactive call.
	//
	// - cron: scheduled task.
	//
	// - im_dingtalk: DingTalk.
	//
	// - im_feishu: Lark.
	//
	// - alert: alert.
	//
	// The session created by this operation is openapi.
	//
	// example:
	//
	// openapi
	SessionType *string `json:"sessionType,omitempty" xml:"sessionType,omitempty"`
	// The session status. The session created by this operation is active, which indicates that the session is available. It does not indicate that the agent is executing a task.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-10-29T10:00:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateCopilotSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCopilotSessionResponseBodyData) GetActiveTurn() *CreateCopilotSessionResponseBodyDataActiveTurn {
	return s.ActiveTurn
}

func (s *CreateCopilotSessionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCopilotSessionResponseBodyData) GetCursor() *string {
	return s.Cursor
}

func (s *CreateCopilotSessionResponseBodyData) GetLastTurn() *CreateCopilotSessionResponseBodyDataLastTurn {
	return s.LastTurn
}

func (s *CreateCopilotSessionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateCopilotSessionResponseBodyData) GetScheduledTask() *CreateCopilotSessionResponseBodyDataScheduledTask {
	return s.ScheduledTask
}

func (s *CreateCopilotSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateCopilotSessionResponseBodyData) GetSessionType() *string {
	return s.SessionType
}

func (s *CreateCopilotSessionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateCopilotSessionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCopilotSessionResponseBodyData) SetActiveTurn(v *CreateCopilotSessionResponseBodyDataActiveTurn) *CreateCopilotSessionResponseBodyData {
	s.ActiveTurn = v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetCreatedAt(v string) *CreateCopilotSessionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetCursor(v string) *CreateCopilotSessionResponseBodyData {
	s.Cursor = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetLastTurn(v *CreateCopilotSessionResponseBodyDataLastTurn) *CreateCopilotSessionResponseBodyData {
	s.LastTurn = v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetName(v string) *CreateCopilotSessionResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetScheduledTask(v *CreateCopilotSessionResponseBodyDataScheduledTask) *CreateCopilotSessionResponseBodyData {
	s.ScheduledTask = v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetSessionId(v string) *CreateCopilotSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetSessionType(v string) *CreateCopilotSessionResponseBodyData {
	s.SessionType = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetStatus(v string) *CreateCopilotSessionResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) SetUpdatedAt(v string) *CreateCopilotSessionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyData) Validate() error {
	if s.ActiveTurn != nil {
		if err := s.ActiveTurn.Validate(); err != nil {
			return err
		}
	}
	if s.LastTurn != nil {
		if err := s.LastTurn.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduledTask != nil {
		if err := s.ScheduledTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCopilotSessionResponseBodyDataActiveTurn struct {
	// The time when the execution was created. The time is in ISO 8601 format and includes the time zone offset.
	//
	// example:
	//
	// 2026-10-29T10:01:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The time when the execution finished. The time is in ISO 8601 format and includes the time zone offset. The value is null if the execution has not finished.
	//
	// example:
	//
	// 2026-10-29T10:02:00+08:00
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The cumulative number of input tokens consumed by this execution.
	//
	// example:
	//
	// 1200
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// The cumulative number of output tokens generated by this execution.
	//
	// example:
	//
	// 80
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// The ID of the form that is waiting for user input. Use this value when submitting user input. The value is null if no form is pending submission.
	//
	// example:
	//
	// form-001
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// The execution status.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The execution ID. Use this value to submit user input, reconnect to the event stream, or stop the execution.
	//
	// example:
	//
	// d71fc5ceac2949f39a264c8a585f3148
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
	// The time when the execution was updated. The time is in ISO 8601 format and includes the time zone offset.
	//
	// example:
	//
	// 2026-10-29T10:01:30+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateCopilotSessionResponseBodyDataActiveTurn) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotSessionResponseBodyDataActiveTurn) GoString() string {
	return s.String()
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) GetPendingInputId() *string {
	return s.PendingInputId
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) GetStatus() *string {
	return s.Status
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) GetTurnId() *string {
	return s.TurnId
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) SetCreatedAt(v string) *CreateCopilotSessionResponseBodyDataActiveTurn {
	s.CreatedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) SetFinishedAt(v string) *CreateCopilotSessionResponseBodyDataActiveTurn {
	s.FinishedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) SetInputTokens(v int64) *CreateCopilotSessionResponseBodyDataActiveTurn {
	s.InputTokens = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) SetOutputTokens(v int64) *CreateCopilotSessionResponseBodyDataActiveTurn {
	s.OutputTokens = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) SetPendingInputId(v string) *CreateCopilotSessionResponseBodyDataActiveTurn {
	s.PendingInputId = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) SetStatus(v string) *CreateCopilotSessionResponseBodyDataActiveTurn {
	s.Status = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) SetTurnId(v string) *CreateCopilotSessionResponseBodyDataActiveTurn {
	s.TurnId = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) SetUpdatedAt(v string) *CreateCopilotSessionResponseBodyDataActiveTurn {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataActiveTurn) Validate() error {
	return dara.Validate(s)
}

type CreateCopilotSessionResponseBodyDataLastTurn struct {
	// The time when the execution was created. The time is in ISO 8601 format and includes the time zone offset.
	//
	// example:
	//
	// 2026-10-29T10:01:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The time when the execution finished. The time is in ISO 8601 format and includes the time zone offset. The value is null if the execution has not finished.
	//
	// example:
	//
	// 2026-10-29T10:01:00+08:00
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The cumulative number of input tokens consumed by this execution.
	//
	// example:
	//
	// 1200
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// The cumulative number of output tokens consumed by this execution.
	//
	// example:
	//
	// 80
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// The ID of the form that is waiting for user input.
	//
	// example:
	//
	// form-001
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// The execution status.
	//
	// example:
	//
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The execution ID.
	//
	// example:
	//
	// d71fc5ceac2949f39a264c8a585f3148
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
	// The time when the execution was updated. The time is in ISO 8601 format and includes the time zone offset.
	//
	// example:
	//
	// 2026-10-29T10:01:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateCopilotSessionResponseBodyDataLastTurn) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotSessionResponseBodyDataLastTurn) GoString() string {
	return s.String()
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) GetPendingInputId() *string {
	return s.PendingInputId
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) GetStatus() *string {
	return s.Status
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) GetTurnId() *string {
	return s.TurnId
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) SetCreatedAt(v string) *CreateCopilotSessionResponseBodyDataLastTurn {
	s.CreatedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) SetFinishedAt(v string) *CreateCopilotSessionResponseBodyDataLastTurn {
	s.FinishedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) SetInputTokens(v int64) *CreateCopilotSessionResponseBodyDataLastTurn {
	s.InputTokens = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) SetOutputTokens(v int64) *CreateCopilotSessionResponseBodyDataLastTurn {
	s.OutputTokens = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) SetPendingInputId(v string) *CreateCopilotSessionResponseBodyDataLastTurn {
	s.PendingInputId = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) SetStatus(v string) *CreateCopilotSessionResponseBodyDataLastTurn {
	s.Status = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) SetTurnId(v string) *CreateCopilotSessionResponseBodyDataLastTurn {
	s.TurnId = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) SetUpdatedAt(v string) *CreateCopilotSessionResponseBodyDataLastTurn {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataLastTurn) Validate() error {
	return dara.Validate(s)
}

type CreateCopilotSessionResponseBodyDataScheduledTask struct {
	// The ID of this scheduled execution, which is used to distinguish different trigger records of the same scheduled task.
	//
	// example:
	//
	// run-example-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The scheduled trigger time. The time is in ISO 8601 format and includes the time zone offset.
	//
	// example:
	//
	// 2026-10-29T09:00:00+08:00
	ScheduledAt *string `json:"scheduledAt,omitempty" xml:"scheduledAt,omitempty"`
	// The time when Copilot started processing this scheduled execution. The time is in ISO 8601 format and includes the time zone offset.
	//
	// example:
	//
	// 2026-10-29T09:00:00+08:00
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The ID of the associated scheduled task.
	//
	// example:
	//
	// 42
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The task name when the scheduled task is triggered. Subsequent changes to the task name do not affect this value.
	//
	// example:
	//
	// Daily inspection
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The scheduling time zone of the scheduled task.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s CreateCopilotSessionResponseBodyDataScheduledTask) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotSessionResponseBodyDataScheduledTask) GoString() string {
	return s.String()
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) GetRunId() *string {
	return s.RunId
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) GetScheduledAt() *string {
	return s.ScheduledAt
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) GetStartedAt() *string {
	return s.StartedAt
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) SetRunId(v string) *CreateCopilotSessionResponseBodyDataScheduledTask {
	s.RunId = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) SetScheduledAt(v string) *CreateCopilotSessionResponseBodyDataScheduledTask {
	s.ScheduledAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) SetStartedAt(v string) *CreateCopilotSessionResponseBodyDataScheduledTask {
	s.StartedAt = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) SetTaskId(v int64) *CreateCopilotSessionResponseBodyDataScheduledTask {
	s.TaskId = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) SetTaskName(v string) *CreateCopilotSessionResponseBodyDataScheduledTask {
	s.TaskName = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) SetTimezone(v string) *CreateCopilotSessionResponseBodyDataScheduledTask {
	s.Timezone = &v
	return s
}

func (s *CreateCopilotSessionResponseBodyDataScheduledTask) Validate() error {
	return dara.Validate(s)
}
