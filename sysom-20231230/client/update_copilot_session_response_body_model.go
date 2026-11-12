// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCopilotSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCopilotSessionResponseBody
	GetCode() *string
	SetData(v *UpdateCopilotSessionResponseBodyData) *UpdateCopilotSessionResponseBody
	GetData() *UpdateCopilotSessionResponseBodyData
	SetMessage(v string) *UpdateCopilotSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCopilotSessionResponseBody
	GetRequestId() *string
}

type UpdateCopilotSessionResponseBody struct {
	// The returned status code. Success indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *UpdateCopilotSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message. If the code is Success, this field is empty. Otherwise, this field contains the request error message.
	//
	// example:
	//
	// Service temporarily unavailable
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. It can be used for end-to-end diagnostics.
	//
	// example:
	//
	// 5BC79EF5-798B-5499-9683-4E14EBBBF712
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateCopilotSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCopilotSessionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCopilotSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCopilotSessionResponseBody) GetData() *UpdateCopilotSessionResponseBodyData {
	return s.Data
}

func (s *UpdateCopilotSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCopilotSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCopilotSessionResponseBody) SetCode(v string) *UpdateCopilotSessionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCopilotSessionResponseBody) SetData(v *UpdateCopilotSessionResponseBodyData) *UpdateCopilotSessionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCopilotSessionResponseBody) SetMessage(v string) *UpdateCopilotSessionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCopilotSessionResponseBody) SetRequestId(v string) *UpdateCopilotSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCopilotSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCopilotSessionResponseBodyData struct {
	// The summary of the execution that currently occupies the session. This field is empty when no active execution is in progress. An execution with an unknown outcome may still occupy the session.
	ActiveTurn *UpdateCopilotSessionResponseBodyDataActiveTurn `json:"activeTurn,omitempty" xml:"activeTurn,omitempty" type:"Struct"`
	// The creation time of the session. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:00:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The latest persisted event cursor for the current session, returned as a decimal string. The value is 0 when no events exist. This value does not represent the position consumed by the client.
	//
	// example:
	//
	// 12
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// The summary of the most recent execution. This field is empty if no execution has ever occurred. If the most recent execution has not ended, this field may point to the same execution as activeTurn.
	LastTurn *UpdateCopilotSessionResponseBodyDataLastTurn `json:"lastTurn,omitempty" xml:"lastTurn,omitempty" type:"Struct"`
	// The session title.
	//
	// example:
	//
	// Troubleshoot CPU load
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The snapshot information when the scheduled task is triggered. This is returned only for scheduled task sessions and is empty for sessions from other sources. Subsequent renaming or deletion of the task does not change the historical snapshot.
	ScheduledTask *UpdateCopilotSessionResponseBodyDataScheduledTask `json:"scheduledTask,omitempty" xml:"scheduledTask,omitempty" type:"Struct"`
	// The session ID.
	//
	// example:
	//
	// s001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The source of the session creation. Valid values: openapi (proactive call), cron (scheduled task), im_dingtalk (DingTalk), im_feishu (Lark), and alert (alert). This is determined by the server, and subsequent follow-up questions do not change the source.
	//
	// example:
	//
	// openapi
	SessionType *string `json:"sessionType,omitempty" xml:"sessionType,omitempty"`
	// The session status. Valid values: active (not archived) and archived (archived). This field does not indicate whether the execution has ended.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The last update time of the session. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:02:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateCopilotSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCopilotSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCopilotSessionResponseBodyData) GetActiveTurn() *UpdateCopilotSessionResponseBodyDataActiveTurn {
	return s.ActiveTurn
}

func (s *UpdateCopilotSessionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateCopilotSessionResponseBodyData) GetCursor() *string {
	return s.Cursor
}

func (s *UpdateCopilotSessionResponseBodyData) GetLastTurn() *UpdateCopilotSessionResponseBodyDataLastTurn {
	return s.LastTurn
}

func (s *UpdateCopilotSessionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateCopilotSessionResponseBodyData) GetScheduledTask() *UpdateCopilotSessionResponseBodyDataScheduledTask {
	return s.ScheduledTask
}

func (s *UpdateCopilotSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateCopilotSessionResponseBodyData) GetSessionType() *string {
	return s.SessionType
}

func (s *UpdateCopilotSessionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateCopilotSessionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCopilotSessionResponseBodyData) SetActiveTurn(v *UpdateCopilotSessionResponseBodyDataActiveTurn) *UpdateCopilotSessionResponseBodyData {
	s.ActiveTurn = v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetCreatedAt(v string) *UpdateCopilotSessionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetCursor(v string) *UpdateCopilotSessionResponseBodyData {
	s.Cursor = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetLastTurn(v *UpdateCopilotSessionResponseBodyDataLastTurn) *UpdateCopilotSessionResponseBodyData {
	s.LastTurn = v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetName(v string) *UpdateCopilotSessionResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetScheduledTask(v *UpdateCopilotSessionResponseBodyDataScheduledTask) *UpdateCopilotSessionResponseBodyData {
	s.ScheduledTask = v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetSessionId(v string) *UpdateCopilotSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetSessionType(v string) *UpdateCopilotSessionResponseBodyData {
	s.SessionType = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetStatus(v string) *UpdateCopilotSessionResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) SetUpdatedAt(v string) *UpdateCopilotSessionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyData) Validate() error {
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

type UpdateCopilotSessionResponseBodyDataActiveTurn struct {
	// The creation time of the current execution turn. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:00:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The end time of the execution. The time is in the ISO 8601 format. This field is empty if the execution has not ended.
	//
	// example:
	//
	// 2026-09-09T09:03:00+08:00
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The cumulative number of input tokens for the current turn. This is not the incremental usage for the current event.
	//
	// example:
	//
	// 1200
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// The cumulative number of output tokens for the current turn.
	//
	// example:
	//
	// 80
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// The ID of the form that is currently waiting for user input. It is used when submitting user input. This field is empty when no form is pending.
	//
	// example:
	//
	// iso-boot-validation
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// The execution status. Valid values: running (in progress), waiting_input (waiting for user input), cancelling (cancelling), completed (completed), failed (failed), cancelled (cancelled), and outcome_unknown (outcome unknown, may still occupy the session, and should not be automatically rerun).
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the current execution turn. It is used to submit user input, reconnect, or stop the execution.
	//
	// example:
	//
	// t001
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
	// The last update time of the current execution turn. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:02:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateCopilotSessionResponseBodyDataActiveTurn) String() string {
	return dara.Prettify(s)
}

func (s UpdateCopilotSessionResponseBodyDataActiveTurn) GoString() string {
	return s.String()
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) GetPendingInputId() *string {
	return s.PendingInputId
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) GetStatus() *string {
	return s.Status
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) GetTurnId() *string {
	return s.TurnId
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) SetCreatedAt(v string) *UpdateCopilotSessionResponseBodyDataActiveTurn {
	s.CreatedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) SetFinishedAt(v string) *UpdateCopilotSessionResponseBodyDataActiveTurn {
	s.FinishedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) SetInputTokens(v int64) *UpdateCopilotSessionResponseBodyDataActiveTurn {
	s.InputTokens = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) SetOutputTokens(v int64) *UpdateCopilotSessionResponseBodyDataActiveTurn {
	s.OutputTokens = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) SetPendingInputId(v string) *UpdateCopilotSessionResponseBodyDataActiveTurn {
	s.PendingInputId = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) SetStatus(v string) *UpdateCopilotSessionResponseBodyDataActiveTurn {
	s.Status = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) SetTurnId(v string) *UpdateCopilotSessionResponseBodyDataActiveTurn {
	s.TurnId = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) SetUpdatedAt(v string) *UpdateCopilotSessionResponseBodyDataActiveTurn {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataActiveTurn) Validate() error {
	return dara.Validate(s)
}

type UpdateCopilotSessionResponseBodyDataLastTurn struct {
	// The creation time of the current execution turn. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:00:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The end time of the execution. The time is in the ISO 8601 format. This field is empty if the execution has not ended.
	//
	// example:
	//
	// 2026-09-09T09:03:00+08:00
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The cumulative number of input tokens for the current turn. This is not the incremental usage for the current event.
	//
	// example:
	//
	// 1200
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// The cumulative number of output tokens for the current turn.
	//
	// example:
	//
	// 80
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// The ID of the form that is currently waiting for user input. It is used when submitting user input. This field is empty when no form is pending.
	//
	// example:
	//
	// iso-boot-validation
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// The execution status. Valid values: running (in progress), waiting_input (waiting for user input), cancelling (cancelling), completed (completed), failed (failed), cancelled (cancelled), and outcome_unknown (outcome unknown, may still occupy the session, and should not be automatically rerun).
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the current execution turn. It is used to submit user input, reconnect, or stop the execution.
	//
	// example:
	//
	// t001
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
	// The last update time of the current execution turn. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:02:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateCopilotSessionResponseBodyDataLastTurn) String() string {
	return dara.Prettify(s)
}

func (s UpdateCopilotSessionResponseBodyDataLastTurn) GoString() string {
	return s.String()
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) GetPendingInputId() *string {
	return s.PendingInputId
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) GetStatus() *string {
	return s.Status
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) GetTurnId() *string {
	return s.TurnId
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) SetCreatedAt(v string) *UpdateCopilotSessionResponseBodyDataLastTurn {
	s.CreatedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) SetFinishedAt(v string) *UpdateCopilotSessionResponseBodyDataLastTurn {
	s.FinishedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) SetInputTokens(v int64) *UpdateCopilotSessionResponseBodyDataLastTurn {
	s.InputTokens = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) SetOutputTokens(v int64) *UpdateCopilotSessionResponseBodyDataLastTurn {
	s.OutputTokens = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) SetPendingInputId(v string) *UpdateCopilotSessionResponseBodyDataLastTurn {
	s.PendingInputId = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) SetStatus(v string) *UpdateCopilotSessionResponseBodyDataLastTurn {
	s.Status = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) SetTurnId(v string) *UpdateCopilotSessionResponseBodyDataLastTurn {
	s.TurnId = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) SetUpdatedAt(v string) *UpdateCopilotSessionResponseBodyDataLastTurn {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataLastTurn) Validate() error {
	return dara.Validate(s)
}

type UpdateCopilotSessionResponseBodyDataScheduledTask struct {
	// The ID of this scheduled execution. Multiple triggers of the same scheduled task use different runIds.
	//
	// example:
	//
	// run-example-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The scheduled trigger time. The time is in the ISO 8601 format with a time zone. This field is empty when no record exists.
	//
	// example:
	//
	// 2026-09-09T01:00:00+00:00
	ScheduledAt *string `json:"scheduledAt,omitempty" xml:"scheduledAt,omitempty"`
	// The time when Copilot starts processing this schedule. The time is in the ISO 8601 format with a time zone. This is not the time when the model outputs the first token.
	//
	// example:
	//
	// 2026-09-09T01:00:03+00:00
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The ID of the associated scheduled task.
	//
	// example:
	//
	// 42
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The name of the scheduled task at the time of this trigger.
	//
	// example:
	//
	// Daily inspection
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The scheduling time zone of the scheduled task, such as Asia/Shanghai.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s UpdateCopilotSessionResponseBodyDataScheduledTask) String() string {
	return dara.Prettify(s)
}

func (s UpdateCopilotSessionResponseBodyDataScheduledTask) GoString() string {
	return s.String()
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) GetRunId() *string {
	return s.RunId
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) GetScheduledAt() *string {
	return s.ScheduledAt
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) GetStartedAt() *string {
	return s.StartedAt
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) SetRunId(v string) *UpdateCopilotSessionResponseBodyDataScheduledTask {
	s.RunId = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) SetScheduledAt(v string) *UpdateCopilotSessionResponseBodyDataScheduledTask {
	s.ScheduledAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) SetStartedAt(v string) *UpdateCopilotSessionResponseBodyDataScheduledTask {
	s.StartedAt = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) SetTaskId(v int64) *UpdateCopilotSessionResponseBodyDataScheduledTask {
	s.TaskId = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) SetTaskName(v string) *UpdateCopilotSessionResponseBodyDataScheduledTask {
	s.TaskName = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) SetTimezone(v string) *UpdateCopilotSessionResponseBodyDataScheduledTask {
	s.Timezone = &v
	return s
}

func (s *UpdateCopilotSessionResponseBodyDataScheduledTask) Validate() error {
	return dara.Validate(s)
}
