// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotSessionMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCopilotSessionMessagesResponseBody
	GetCode() *string
	SetData(v *ListCopilotSessionMessagesResponseBodyData) *ListCopilotSessionMessagesResponseBody
	GetData() *ListCopilotSessionMessagesResponseBodyData
	SetMaxResults(v int32) *ListCopilotSessionMessagesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListCopilotSessionMessagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCopilotSessionMessagesResponseBody
	GetRequestId() *string
}

type ListCopilotSessionMessagesResponseBody struct {
	// The returned status code. Success indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListCopilotSessionMessagesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// A field retained for historical compatibility, which may be omitted. Use data.maxResults instead.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The error message. If the code is Success, this field is empty. Otherwise, this field contains the error message of the request.
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

func (s ListCopilotSessionMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionMessagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCopilotSessionMessagesResponseBody) GetData() *ListCopilotSessionMessagesResponseBodyData {
	return s.Data
}

func (s *ListCopilotSessionMessagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCopilotSessionMessagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCopilotSessionMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCopilotSessionMessagesResponseBody) SetCode(v string) *ListCopilotSessionMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBody) SetData(v *ListCopilotSessionMessagesResponseBodyData) *ListCopilotSessionMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListCopilotSessionMessagesResponseBody) SetMaxResults(v int32) *ListCopilotSessionMessagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBody) SetMessage(v string) *ListCopilotSessionMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBody) SetRequestId(v string) *ListCopilotSessionMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCopilotSessionMessagesResponseBodyData struct {
	// The cursor of the last event on the current page, returned as a decimal string. The value is 0 for an empty page. This is not the pagination nextToken and does not indicate that a round of execution has finished.
	//
	// example:
	//
	// 12
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// Indicates whether there are subsequent historical events. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The page size used for the current request.
	//
	// example:
	//
	// 50
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The history event records on the current page, sorted in ascending order by eventId. A record does not necessarily represent a complete chat message and may be a chunk of text or a tool calling (invoke) event.
	Messages []*ListCopilotSessionMessagesResponseBodyDataMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// The token for the next page. Pass this value back as is in the next request. An empty value indicates that there is no next page.
	//
	// example:
	//
	// 50
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The session summary of the current user. This object does not contain the active execution summary. To query the execution status, call GetCopilotSession.
	Session *ListCopilotSessionMessagesResponseBodyDataSession `json:"session,omitempty" xml:"session,omitempty" type:"Struct"`
}

func (s ListCopilotSessionMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionMessagesResponseBodyData) GetCursor() *string {
	return s.Cursor
}

func (s *ListCopilotSessionMessagesResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListCopilotSessionMessagesResponseBodyData) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCopilotSessionMessagesResponseBodyData) GetMessages() []*ListCopilotSessionMessagesResponseBodyDataMessages {
	return s.Messages
}

func (s *ListCopilotSessionMessagesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCopilotSessionMessagesResponseBodyData) GetSession() *ListCopilotSessionMessagesResponseBodyDataSession {
	return s.Session
}

func (s *ListCopilotSessionMessagesResponseBodyData) SetCursor(v string) *ListCopilotSessionMessagesResponseBodyData {
	s.Cursor = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyData) SetHasMore(v bool) *ListCopilotSessionMessagesResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyData) SetMaxResults(v int64) *ListCopilotSessionMessagesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyData) SetMessages(v []*ListCopilotSessionMessagesResponseBodyDataMessages) *ListCopilotSessionMessagesResponseBodyData {
	s.Messages = v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyData) SetNextToken(v string) *ListCopilotSessionMessagesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyData) SetSession(v *ListCopilotSessionMessagesResponseBodyDataSession) *ListCopilotSessionMessagesResponseBodyData {
	s.Session = v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyData) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCopilotSessionMessagesResponseBodyDataMessages struct {
	// The human-readable content summary of the event. To restore the complete interactive interface, use the original event in metadata.
	Content []*ListCopilotSessionMessagesResponseBodyDataMessagesContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The time when the historical event record was created, in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-09T09:00:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The position of the event within the session, used for sorting and deduplication.
	//
	// example:
	//
	// 12
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// Currently, an empty string is returned. This field is not used to determine whether the execution has finished.
	//
	// example:
	//
	// ""
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The ID of the historical event record.
	//
	// example:
	//
	// record-001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The complete original event object. The internal fields retain their original protocol names, without recursive naming conversion or the need for further JSON string parsing.
	//
	// example:
	//
	// {"type":"TEXT_MESSAGE_CONTENT","messageId":"m001","delta":"Checking."}
	Metadata interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// Currently, an empty string is returned. This field is not used to identify the event type.
	//
	// example:
	//
	// ""
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The role in the original event. An empty string is returned if the event does not contain a role.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The ID of the associated execution. This field is empty if there is no associated execution.
	//
	// example:
	//
	// t001
	TurnId *string `json:"turnId,omitempty" xml:"turnId,omitempty"`
	// The token usage per event. This data is not currently returned. To query the cumulative usage of an execution, call GetCopilotSession.
	//
	// example:
	//
	// null
	Usage interface{} `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ListCopilotSessionMessagesResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionMessagesResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetContent() []*ListCopilotSessionMessagesResponseBodyDataMessagesContent {
	return s.Content
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetEventId() *string {
	return s.EventId
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetId() *string {
	return s.Id
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetName() *string {
	return s.Name
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetRole() *string {
	return s.Role
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetTurnId() *string {
	return s.TurnId
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) GetUsage() interface{} {
	return s.Usage
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetContent(v []*ListCopilotSessionMessagesResponseBodyDataMessagesContent) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.Content = v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetCreatedAt(v string) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.CreatedAt = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetEventId(v string) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.EventId = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetFinishedAt(v string) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.FinishedAt = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetId(v string) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.Id = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetMetadata(v interface{}) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.Metadata = v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetName(v string) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.Name = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetRole(v string) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.Role = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetTurnId(v string) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.TurnId = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) SetUsage(v interface{}) *ListCopilotSessionMessagesResponseBodyDataMessages {
	s.Usage = v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessages) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCopilotSessionMessagesResponseBodyDataMessagesContent struct {
	// The readable text content in the event. This field is empty if no readable string is available.
	//
	// example:
	//
	// Checking
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The event type, such as USER_MESSAGE, TEXT_MESSAGE_CONTENT, or TOOL_CALL_ARGS. This is not fixed to text.
	//
	// example:
	//
	// TEXT_MESSAGE_CONTENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListCopilotSessionMessagesResponseBodyDataMessagesContent) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionMessagesResponseBodyDataMessagesContent) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessagesContent) GetText() *string {
	return s.Text
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessagesContent) GetType() *string {
	return s.Type
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessagesContent) SetText(v string) *ListCopilotSessionMessagesResponseBodyDataMessagesContent {
	s.Text = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessagesContent) SetType(v string) *ListCopilotSessionMessagesResponseBodyDataMessagesContent {
	s.Type = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataMessagesContent) Validate() error {
	return dara.Validate(s)
}

type ListCopilotSessionMessagesResponseBodyDataSession struct {
	// The agent identifier in the session summary, which does not represent the current executor. The current executor is subject to activeTurn.agentId.
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
	// The session ID, which is the same as sessionId. We recommend using sessionId to initiate subsequent requests.
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
	// The snapshot information when the scheduled task is triggered. This field is returned only for scheduled task sessions and is empty for sessions from other sources. Subsequent renaming or deletion of the task does not change the historical snapshot.
	ScheduledTask *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask `json:"scheduledTask,omitempty" xml:"scheduledTask,omitempty" type:"Struct"`
	// The session ID. Use this value when you send follow-up questions, query history, or delete the session.
	//
	// example:
	//
	// s001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The source from which the session was created. Valid values:
	//
	// - openapi: Active call.
	//
	// - cron: Scheduled task.
	//
	// - im_dingtalk: DingTalk.
	//
	// - im_feishu: Lark.
	//
	// - alert: Alert.
	//
	// This value is determined by the server and is not changed by follow-up questions.
	//
	// example:
	//
	// openapi
	SessionType *string `json:"sessionType,omitempty" xml:"sessionType,omitempty"`
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

func (s ListCopilotSessionMessagesResponseBodyDataSession) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionMessagesResponseBodyDataSession) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetAgentId() *string {
	return s.AgentId
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetId() *string {
	return s.Id
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetName() *string {
	return s.Name
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetScheduledTask() *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask {
	return s.ScheduledTask
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetSessionId() *string {
	return s.SessionId
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetSessionType() *string {
	return s.SessionType
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) GetUserId() *string {
	return s.UserId
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetAgentId(v string) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.AgentId = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetCreatedAt(v string) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.CreatedAt = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetId(v string) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.Id = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetName(v string) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.Name = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetScheduledTask(v *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.ScheduledTask = v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetSessionId(v string) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.SessionId = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetSessionType(v string) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.SessionType = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetUpdatedAt(v string) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.UpdatedAt = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) SetUserId(v string) *ListCopilotSessionMessagesResponseBodyDataSession {
	s.UserId = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSession) Validate() error {
	if s.ScheduledTask != nil {
		if err := s.ScheduledTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask struct {
	// The ID of this scheduled execution. Different runIds are used for multiple triggers of the same scheduled task.
	//
	// example:
	//
	// run-example-001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The scheduled trigger time, in ISO 8601 format with time zone. This field is empty if there is no record.
	//
	// example:
	//
	// 2026-09-09T01:00:00+00:00
	ScheduledAt *string `json:"scheduledAt,omitempty" xml:"scheduledAt,omitempty"`
	// The time when Copilot started processing this schedule, in ISO 8601 format with time zone. This is not the time when the model outputs the first token.
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

func (s ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) GoString() string {
	return s.String()
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) GetRunId() *string {
	return s.RunId
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) GetScheduledAt() *string {
	return s.ScheduledAt
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) GetStartedAt() *string {
	return s.StartedAt
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) GetTaskName() *string {
	return s.TaskName
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) GetTimezone() *string {
	return s.Timezone
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) SetRunId(v string) *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask {
	s.RunId = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) SetScheduledAt(v string) *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask {
	s.ScheduledAt = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) SetStartedAt(v string) *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask {
	s.StartedAt = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) SetTaskId(v int64) *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask {
	s.TaskId = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) SetTaskName(v string) *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask {
	s.TaskName = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) SetTimezone(v string) *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask {
	s.Timezone = &v
	return s
}

func (s *ListCopilotSessionMessagesResponseBodyDataSessionScheduledTask) Validate() error {
	return dara.Validate(s)
}
