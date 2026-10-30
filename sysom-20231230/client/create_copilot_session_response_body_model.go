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
	// example:
	//
	// Success
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateCopilotSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
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
	ActiveTurn *CreateCopilotSessionResponseBodyDataActiveTurn `json:"activeTurn,omitempty" xml:"activeTurn,omitempty" type:"Struct"`
	// example:
	//
	// 2026-10-29T10:00:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 十进制事件序号字符串，范围 0..9223372036854775807；避免 JavaScript 精度丢失
	//
	// example:
	//
	// 0
	Cursor   *string                                       `json:"cursor,omitempty" xml:"cursor,omitempty"`
	LastTurn *CreateCopilotSessionResponseBodyDataLastTurn `json:"lastTurn,omitempty" xml:"lastTurn,omitempty" type:"Struct"`
	// example:
	//
	// 排查 CPU 负载
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 定时任务触发快照；非定时会话为空
	ScheduledTask *CreateCopilotSessionResponseBodyDataScheduledTask `json:"scheduledTask,omitempty" xml:"scheduledTask,omitempty" type:"Struct"`
	// example:
	//
	// c356bc310e4345c3a602a2912322ad21
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 会话初始来源：chat、scheduled_task、alert；后续追问不改变来源
	//
	// example:
	//
	// openapi
	SessionType *string `json:"sessionType,omitempty" xml:"sessionType,omitempty"`
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
	// example:
	//
	// 2026-10-29T10:01:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 2026-10-29T10:02:00+08:00
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// example:
	//
	// 1200
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 80
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// form-001
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// d71fc5ceac2949f39a264c8a585f3148
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
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
	// example:
	//
	// 2026-10-29T10:01:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 2026-10-29T10:01:00+08:00
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// example:
	//
	// 1200
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 80
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// form-001
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// d71fc5ceac2949f39a264c8a585f3148
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
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
	// 本次定时执行的 Run ID
	//
	// example:
	//
	// run-example-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// 计划触发时间，ISO 8601 含时区偏移
	//
	// example:
	//
	// 2026-10-29T09:00:00+08:00
	ScheduledAt *string `json:"scheduledAt,omitempty" xml:"scheduledAt,omitempty"`
	// Copilot 开始处理时间，ISO 8601 含时区偏移
	//
	// example:
	//
	// 2026-10-29T09:00:00+08:00
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// 任务 ID；任务删除后仍保留
	//
	// example:
	//
	// 42
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 触发时的任务名称快照
	//
	// example:
	//
	// 每日巡检
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// 任务的 cron 时区
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
