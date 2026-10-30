// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCopilotSessionResponseBody
	GetCode() *string
	SetData(v *GetCopilotSessionResponseBodyData) *GetCopilotSessionResponseBody
	GetData() *GetCopilotSessionResponseBodyData
	SetMessage(v string) *GetCopilotSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCopilotSessionResponseBody
	GetRequestId() *string
}

type GetCopilotSessionResponseBody struct {
	// example:
	//
	// Success
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetCopilotSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5BC79EF5-798B-5499-9683-4E14EBBBF712
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCopilotSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetCopilotSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCopilotSessionResponseBody) GetData() *GetCopilotSessionResponseBodyData {
	return s.Data
}

func (s *GetCopilotSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCopilotSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCopilotSessionResponseBody) SetCode(v string) *GetCopilotSessionResponseBody {
	s.Code = &v
	return s
}

func (s *GetCopilotSessionResponseBody) SetData(v *GetCopilotSessionResponseBodyData) *GetCopilotSessionResponseBody {
	s.Data = v
	return s
}

func (s *GetCopilotSessionResponseBody) SetMessage(v string) *GetCopilotSessionResponseBody {
	s.Message = &v
	return s
}

func (s *GetCopilotSessionResponseBody) SetRequestId(v string) *GetCopilotSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCopilotSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCopilotSessionResponseBodyData struct {
	ActiveTurn *GetCopilotSessionResponseBodyDataActiveTurn `json:"activeTurn,omitempty" xml:"activeTurn,omitempty" type:"Struct"`
	// example:
	//
	// 2026-09-09T01:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 十进制事件序号字符串，范围 0..9223372036854775807；避免 JavaScript 精度丢失
	//
	// example:
	//
	// 12
	Cursor   *string                                    `json:"cursor,omitempty" xml:"cursor,omitempty"`
	LastTurn *GetCopilotSessionResponseBodyDataLastTurn `json:"lastTurn,omitempty" xml:"lastTurn,omitempty" type:"Struct"`
	// example:
	//
	// 排查 CPU 负载
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 定时任务触发快照；非定时会话为空
	ScheduledTask *GetCopilotSessionResponseBodyDataScheduledTask `json:"scheduledTask,omitempty" xml:"scheduledTask,omitempty" type:"Struct"`
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
	// 2026-09-09T01:02:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetCopilotSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCopilotSessionResponseBodyData) GetActiveTurn() *GetCopilotSessionResponseBodyDataActiveTurn {
	return s.ActiveTurn
}

func (s *GetCopilotSessionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCopilotSessionResponseBodyData) GetCursor() *string {
	return s.Cursor
}

func (s *GetCopilotSessionResponseBodyData) GetLastTurn() *GetCopilotSessionResponseBodyDataLastTurn {
	return s.LastTurn
}

func (s *GetCopilotSessionResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetCopilotSessionResponseBodyData) GetScheduledTask() *GetCopilotSessionResponseBodyDataScheduledTask {
	return s.ScheduledTask
}

func (s *GetCopilotSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetCopilotSessionResponseBodyData) GetSessionType() *string {
	return s.SessionType
}

func (s *GetCopilotSessionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCopilotSessionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCopilotSessionResponseBodyData) SetActiveTurn(v *GetCopilotSessionResponseBodyDataActiveTurn) *GetCopilotSessionResponseBodyData {
	s.ActiveTurn = v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetCreatedAt(v string) *GetCopilotSessionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetCursor(v string) *GetCopilotSessionResponseBodyData {
	s.Cursor = &v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetLastTurn(v *GetCopilotSessionResponseBodyDataLastTurn) *GetCopilotSessionResponseBodyData {
	s.LastTurn = v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetName(v string) *GetCopilotSessionResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetScheduledTask(v *GetCopilotSessionResponseBodyDataScheduledTask) *GetCopilotSessionResponseBodyData {
	s.ScheduledTask = v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetSessionId(v string) *GetCopilotSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetSessionType(v string) *GetCopilotSessionResponseBodyData {
	s.SessionType = &v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetStatus(v string) *GetCopilotSessionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCopilotSessionResponseBodyData) SetUpdatedAt(v string) *GetCopilotSessionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyData) Validate() error {
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

type GetCopilotSessionResponseBodyDataActiveTurn struct {
	// example:
	//
	// 2026-09-09T01:01:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 2026-09-09T01:03:00Z
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
	// iso-boot-validation
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// e499819f351242dba59ab8af31bc8b0f
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
	// example:
	//
	// 2026-09-09T01:02:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetCopilotSessionResponseBodyDataActiveTurn) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotSessionResponseBodyDataActiveTurn) GoString() string {
	return s.String()
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) GetPendingInputId() *string {
	return s.PendingInputId
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) GetStatus() *string {
	return s.Status
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) GetTurnId() *string {
	return s.TurnId
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) SetCreatedAt(v string) *GetCopilotSessionResponseBodyDataActiveTurn {
	s.CreatedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) SetFinishedAt(v string) *GetCopilotSessionResponseBodyDataActiveTurn {
	s.FinishedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) SetInputTokens(v int64) *GetCopilotSessionResponseBodyDataActiveTurn {
	s.InputTokens = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) SetOutputTokens(v int64) *GetCopilotSessionResponseBodyDataActiveTurn {
	s.OutputTokens = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) SetPendingInputId(v string) *GetCopilotSessionResponseBodyDataActiveTurn {
	s.PendingInputId = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) SetStatus(v string) *GetCopilotSessionResponseBodyDataActiveTurn {
	s.Status = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) SetTurnId(v string) *GetCopilotSessionResponseBodyDataActiveTurn {
	s.TurnId = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) SetUpdatedAt(v string) *GetCopilotSessionResponseBodyDataActiveTurn {
	s.UpdatedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataActiveTurn) Validate() error {
	return dara.Validate(s)
}

type GetCopilotSessionResponseBodyDataLastTurn struct {
	// example:
	//
	// 2026-09-09T01:01:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 2026-09-09T01:03:00Z
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
	// iso-boot-validation
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// e499819f351242dba59ab8af31bc8b0f
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
	// example:
	//
	// 2026-09-09T01:02:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetCopilotSessionResponseBodyDataLastTurn) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotSessionResponseBodyDataLastTurn) GoString() string {
	return s.String()
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) GetPendingInputId() *string {
	return s.PendingInputId
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) GetStatus() *string {
	return s.Status
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) GetTurnId() *string {
	return s.TurnId
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) SetCreatedAt(v string) *GetCopilotSessionResponseBodyDataLastTurn {
	s.CreatedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) SetFinishedAt(v string) *GetCopilotSessionResponseBodyDataLastTurn {
	s.FinishedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) SetInputTokens(v int64) *GetCopilotSessionResponseBodyDataLastTurn {
	s.InputTokens = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) SetOutputTokens(v int64) *GetCopilotSessionResponseBodyDataLastTurn {
	s.OutputTokens = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) SetPendingInputId(v string) *GetCopilotSessionResponseBodyDataLastTurn {
	s.PendingInputId = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) SetStatus(v string) *GetCopilotSessionResponseBodyDataLastTurn {
	s.Status = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) SetTurnId(v string) *GetCopilotSessionResponseBodyDataLastTurn {
	s.TurnId = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) SetUpdatedAt(v string) *GetCopilotSessionResponseBodyDataLastTurn {
	s.UpdatedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataLastTurn) Validate() error {
	return dara.Validate(s)
}

type GetCopilotSessionResponseBodyDataScheduledTask struct {
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
	// 2026-09-09T01:00:00+00:00
	ScheduledAt *string `json:"scheduledAt,omitempty" xml:"scheduledAt,omitempty"`
	// Copilot 开始处理时间，ISO 8601 含时区偏移
	//
	// example:
	//
	// 2026-09-09T01:00:03+00:00
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

func (s GetCopilotSessionResponseBodyDataScheduledTask) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotSessionResponseBodyDataScheduledTask) GoString() string {
	return s.String()
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) GetRunId() *string {
	return s.RunId
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) GetScheduledAt() *string {
	return s.ScheduledAt
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) GetStartedAt() *string {
	return s.StartedAt
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) GetTaskName() *string {
	return s.TaskName
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) GetTimezone() *string {
	return s.Timezone
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) SetRunId(v string) *GetCopilotSessionResponseBodyDataScheduledTask {
	s.RunId = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) SetScheduledAt(v string) *GetCopilotSessionResponseBodyDataScheduledTask {
	s.ScheduledAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) SetStartedAt(v string) *GetCopilotSessionResponseBodyDataScheduledTask {
	s.StartedAt = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) SetTaskId(v int64) *GetCopilotSessionResponseBodyDataScheduledTask {
	s.TaskId = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) SetTaskName(v string) *GetCopilotSessionResponseBodyDataScheduledTask {
	s.TaskName = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) SetTimezone(v string) *GetCopilotSessionResponseBodyDataScheduledTask {
	s.Timezone = &v
	return s
}

func (s *GetCopilotSessionResponseBodyDataScheduledTask) Validate() error {
	return dara.Validate(s)
}
