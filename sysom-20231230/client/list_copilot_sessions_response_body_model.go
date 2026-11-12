// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCopilotSessionsResponseBody
	GetCode() *string
	SetData(v *ListCopilotSessionsResponseBodyData) *ListCopilotSessionsResponseBody
	GetData() *ListCopilotSessionsResponseBodyData
	SetMaxResults(v int32) *ListCopilotSessionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListCopilotSessionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCopilotSessionsResponseBody
	GetRequestId() *string
}

type ListCopilotSessionsResponseBody struct {
	// The returned status code. Success indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListCopilotSessionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The page size used for the current request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The error message. If code is Success, this field is empty. Otherwise, this field contains the request error message.
	//
	// example:
	//
	// Service temporarily unavailable
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which can be used for end-to-end diagnostics.
	//
	// example:
	//
	// 5BC79EF5-798B-5499-9683-4E14EBBBF712
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListCopilotSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCopilotSessionsResponseBody) GetData() *ListCopilotSessionsResponseBodyData {
	return s.Data
}

func (s *ListCopilotSessionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCopilotSessionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCopilotSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCopilotSessionsResponseBody) SetCode(v string) *ListCopilotSessionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCopilotSessionsResponseBody) SetData(v *ListCopilotSessionsResponseBodyData) *ListCopilotSessionsResponseBody {
	s.Data = v
	return s
}

func (s *ListCopilotSessionsResponseBody) SetMaxResults(v int32) *ListCopilotSessionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCopilotSessionsResponseBody) SetMessage(v string) *ListCopilotSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCopilotSessionsResponseBody) SetRequestId(v string) *ListCopilotSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCopilotSessionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCopilotSessionsResponseBodyData struct {
	// The page size used for the current request.
	//
	// example:
	//
	// 50
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. Pass this value back as is in the next request. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// 50
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The session list.
	Sessions []*ListCopilotSessionsResponseBodyDataSessions `json:"sessions,omitempty" xml:"sessions,omitempty" type:"Repeated"`
}

func (s ListCopilotSessionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionsResponseBodyData) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCopilotSessionsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCopilotSessionsResponseBodyData) GetSessions() []*ListCopilotSessionsResponseBodyDataSessions {
	return s.Sessions
}

func (s *ListCopilotSessionsResponseBodyData) SetMaxResults(v int64) *ListCopilotSessionsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyData) SetNextToken(v string) *ListCopilotSessionsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyData) SetSessions(v []*ListCopilotSessionsResponseBodyDataSessions) *ListCopilotSessionsResponseBodyData {
	s.Sessions = v
	return s
}

func (s *ListCopilotSessionsResponseBodyData) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCopilotSessionsResponseBodyDataSessions struct {
	// The information about the execution that currently occupies the session.
	ActiveTurn *ListCopilotSessionsResponseBodyDataSessionsActiveTurn `json:"activeTurn,omitempty" xml:"activeTurn,omitempty" type:"Struct"`
	// The ID of the execution that currently occupies the session. This field is empty when no active execution exists.
	//
	// example:
	//
	// t001
	ActiveTurnId *string `json:"activeTurnId,omitempty" xml:"activeTurnId,omitempty"`
	// The Agent tag of the session. The current executor is subject to activeTurn.agentId.
	//
	// example:
	//
	// sysagent
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The time when the session was created, in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:00:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The session ID.
	//
	// example:
	//
	// s001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The session title.
	//
	// example:
	//
	// Troubleshoot CPU load
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The snapshot information when the scheduled task is triggered. This field is returned only for scheduled task sessions and is empty for sessions from other sources.
	ScheduledTask *ListCopilotSessionsResponseBodyDataSessionsScheduledTask `json:"scheduledTask,omitempty" xml:"scheduledTask,omitempty" type:"Struct"`
	// The session ID. Use this value when you send questions, query history, and delete sessions later.
	//
	// example:
	//
	// s001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The source from which the session was created. Valid values:
	//
	// - openapi: API call.
	//
	// - cron: Scheduled task.
	//
	// - im_dingtalk: DingTalk.
	//
	// - im_feishu: Lark.
	//
	// - alert: Alert.
	//
	// This value is determined by the server and is not changed by subsequent follow-up questions.
	//
	// example:
	//
	// openapi
	SessionType *string `json:"sessionType,omitempty" xml:"sessionType,omitempty"`
	// The session status. Valid values:
	//
	// - active: Not archived.
	//
	// - archived: Archived.
	//
	// This field does not indicate whether the execution has ended.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the session was last updated, in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:02:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The ID of the user to whom the session belongs. This value is determined by the server based on the caller identity and is not used as a parameter for the client to specify the identity.
	//
	// example:
	//
	// 10001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListCopilotSessionsResponseBodyDataSessions) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionsResponseBodyDataSessions) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetActiveTurn() *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	return s.ActiveTurn
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetActiveTurnId() *string {
	return s.ActiveTurnId
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetAgentId() *string {
	return s.AgentId
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetId() *string {
	return s.Id
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetName() *string {
	return s.Name
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetScheduledTask() *ListCopilotSessionsResponseBodyDataSessionsScheduledTask {
	return s.ScheduledTask
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetSessionId() *string {
	return s.SessionId
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetSessionType() *string {
	return s.SessionType
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetStatus() *string {
	return s.Status
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListCopilotSessionsResponseBodyDataSessions) GetUserId() *string {
	return s.UserId
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetActiveTurn(v *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) *ListCopilotSessionsResponseBodyDataSessions {
	s.ActiveTurn = v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetActiveTurnId(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.ActiveTurnId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetAgentId(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.AgentId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetCreatedAt(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.CreatedAt = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetId(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.Id = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetName(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.Name = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetScheduledTask(v *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) *ListCopilotSessionsResponseBodyDataSessions {
	s.ScheduledTask = v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetSessionId(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.SessionId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetSessionType(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.SessionType = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetStatus(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.Status = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetUpdatedAt(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.UpdatedAt = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) SetUserId(v string) *ListCopilotSessionsResponseBodyDataSessions {
	s.UserId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessions) Validate() error {
	if s.ActiveTurn != nil {
		if err := s.ActiveTurn.Validate(); err != nil {
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

type ListCopilotSessionsResponseBodyDataSessionsActiveTurn struct {
	// The Agent ID bound to the current execution round.
	//
	// example:
	//
	// sysagent
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The time when the current execution round was created, in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:00:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The time when the execution ended, in ISO 8601 format. This field is empty if the execution has not ended.
	//
	// example:
	//
	// 2026-09-09T09:03:00+08:00
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The cumulative number of input tokens in the current round, which is not the incremental usage of the current event.
	//
	// example:
	//
	// 1200
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// The cumulative number of output tokens in the current round.
	//
	// example:
	//
	// 80
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// The ID of the form that is currently waiting for the user to fill out. This field is used when user input is submitted. This field is empty when no form needs to be filled out.
	//
	// example:
	//
	// iso-boot-validation
	PendingInputId *string `json:"pendingInputId,omitempty" xml:"pendingInputId,omitempty"`
	// The execution status. Valid values:
	//
	// - running: The execution is in progress.
	//
	// - waiting_input: Waiting for user input.
	//
	// - cancelling: The execution is being canceled.
	//
	// - completed: The execution is completed.
	//
	// - failed: The execution failed.
	//
	// - cancelled: The execution is canceled.
	//
	// - outcome_unknown: The result is unknown, and the session may still be occupied. The execution should not be automatically rerun.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the current execution round, which is used to submit user input, reconnect, or stop the execution.
	//
	// example:
	//
	// t001
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
	// The time when the current execution round was last updated, in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:02:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListCopilotSessionsResponseBodyDataSessionsActiveTurn) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetAgentId() *string {
	return s.AgentId
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetPendingInputId() *string {
	return s.PendingInputId
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetStatus() *string {
	return s.Status
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetTurnId() *string {
	return s.TurnId
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetAgentId(v string) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.AgentId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetCreatedAt(v string) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.CreatedAt = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetFinishedAt(v string) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.FinishedAt = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetInputTokens(v int64) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.InputTokens = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetOutputTokens(v int64) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.OutputTokens = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetPendingInputId(v string) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.PendingInputId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetStatus(v string) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.Status = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetTurnId(v string) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.TurnId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) SetUpdatedAt(v string) *ListCopilotSessionsResponseBodyDataSessionsActiveTurn {
	s.UpdatedAt = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsActiveTurn) Validate() error {
	return dara.Validate(s)
}

type ListCopilotSessionsResponseBodyDataSessionsScheduledTask struct {
	// The ID of the current scheduled execution. Different runIds are used for multiple triggers of the same scheduled task.
	//
	// example:
	//
	// run-example-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The scheduled trigger time, in ISO 8601 format with time zone. This field is empty when no record exists.
	//
	// example:
	//
	// 2026-09-09T01:00:00+00:00
	ScheduledAt *string `json:"scheduledAt,omitempty" xml:"scheduledAt,omitempty"`
	// The time when Copilot started processing the current schedule, in ISO 8601 format with time zone. This is not the time when the model outputs the first token.
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
	// The scheduling time zone of the scheduled task.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListCopilotSessionsResponseBodyDataSessionsScheduledTask) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionsResponseBodyDataSessionsScheduledTask) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) GetRunId() *string {
	return s.RunId
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) GetScheduledAt() *string {
	return s.ScheduledAt
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) GetStartedAt() *string {
	return s.StartedAt
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) GetTaskName() *string {
	return s.TaskName
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) GetTimezone() *string {
	return s.Timezone
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) SetRunId(v string) *ListCopilotSessionsResponseBodyDataSessionsScheduledTask {
	s.RunId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) SetScheduledAt(v string) *ListCopilotSessionsResponseBodyDataSessionsScheduledTask {
	s.ScheduledAt = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) SetStartedAt(v string) *ListCopilotSessionsResponseBodyDataSessionsScheduledTask {
	s.StartedAt = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) SetTaskId(v int64) *ListCopilotSessionsResponseBodyDataSessionsScheduledTask {
	s.TaskId = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) SetTaskName(v string) *ListCopilotSessionsResponseBodyDataSessionsScheduledTask {
	s.TaskName = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) SetTimezone(v string) *ListCopilotSessionsResponseBodyDataSessionsScheduledTask {
	s.Timezone = &v
	return s
}

func (s *ListCopilotSessionsResponseBodyDataSessionsScheduledTask) Validate() error {
	return dara.Validate(s)
}
