// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationChatMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *KopilotListConversationChatMessagesResponseBody
	GetCode() *int64
	SetData(v *KopilotListConversationChatMessagesResponseBodyData) *KopilotListConversationChatMessagesResponseBody
	GetData() *KopilotListConversationChatMessagesResponseBodyData
	SetRequestId(v string) *KopilotListConversationChatMessagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *KopilotListConversationChatMessagesResponseBody
	GetSuccess() *bool
}

type KopilotListConversationChatMessagesResponseBody struct {
	// The response code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned when the call is successful.
	Data *KopilotListConversationChatMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F69385B9-2139-5A07-AE64-37C4B6ED308E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *KopilotListConversationChatMessagesResponseBody) GetData() *KopilotListConversationChatMessagesResponseBodyData {
	return s.Data
}

func (s *KopilotListConversationChatMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KopilotListConversationChatMessagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *KopilotListConversationChatMessagesResponseBody) SetCode(v int64) *KopilotListConversationChatMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBody) SetData(v *KopilotListConversationChatMessagesResponseBodyData) *KopilotListConversationChatMessagesResponseBody {
	s.Data = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBody) SetRequestId(v string) *KopilotListConversationChatMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBody) SetSuccess(v bool) *KopilotListConversationChatMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationChatMessagesResponseBodyData struct {
	// Indicates whether more data is available.
	//
	// example:
	//
	// true
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The list of messages.
	Messages []*KopilotListConversationChatMessagesResponseBodyDataMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The cursor for the next page.
	//
	// example:
	//
	// 1
	NextBeforeTurnId *int64 `json:"NextBeforeTurnId,omitempty" xml:"NextBeforeTurnId,omitempty"`
	// The details of scheduled tasks associated with the current session. Only tasks in the enabled, paused, or pending authorization state are counted.
	ScheduledTaskInfo *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo `json:"ScheduledTaskInfo,omitempty" xml:"ScheduledTaskInfo,omitempty" type:"Struct"`
	// The scheduled task quota for the current Alibaba Cloud account in this environment, counted across regions.
	ScheduledTaskQuota *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota `json:"ScheduledTaskQuota,omitempty" xml:"ScheduledTaskQuota,omitempty" type:"Struct"`
	// The session ID.
	//
	// example:
	//
	// 87ce9505-7dec-4fd7-bc7c-e66d949bfdc9
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The total number of turn IDs.
	//
	// example:
	//
	// 2
	TotalTurns *int64 `json:"TotalTurns,omitempty" xml:"TotalTurns,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetMessages() []*KopilotListConversationChatMessagesResponseBodyDataMessages {
	return s.Messages
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetNextBeforeTurnId() *int64 {
	return s.NextBeforeTurnId
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetScheduledTaskInfo() *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo {
	return s.ScheduledTaskInfo
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetScheduledTaskQuota() *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota {
	return s.ScheduledTaskQuota
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *KopilotListConversationChatMessagesResponseBodyData) GetTotalTurns() *int64 {
	return s.TotalTurns
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetHasMore(v bool) *KopilotListConversationChatMessagesResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetMessages(v []*KopilotListConversationChatMessagesResponseBodyDataMessages) *KopilotListConversationChatMessagesResponseBodyData {
	s.Messages = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetNextBeforeTurnId(v int64) *KopilotListConversationChatMessagesResponseBodyData {
	s.NextBeforeTurnId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetScheduledTaskInfo(v *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) *KopilotListConversationChatMessagesResponseBodyData {
	s.ScheduledTaskInfo = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetScheduledTaskQuota(v *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) *KopilotListConversationChatMessagesResponseBodyData {
	s.ScheduledTaskQuota = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetSessionId(v string) *KopilotListConversationChatMessagesResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) SetTotalTurns(v int64) *KopilotListConversationChatMessagesResponseBodyData {
	s.TotalTurns = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyData) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduledTaskInfo != nil {
		if err := s.ScheduledTaskInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduledTaskQuota != nil {
		if err := s.ScheduledTaskQuota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationChatMessagesResponseBodyDataMessages struct {
	// The actual content of the message.
	//
	// example:
	//
	// test
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The UNIX timestamp when the message was created, in milliseconds.
	//
	// example:
	//
	// 17575885545677
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user satisfaction level.
	//
	// example:
	//
	// 1
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The role identifier.
	//
	// example:
	//
	// assistant
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 2345
	TurnId *string `json:"TurnId,omitempty" xml:"TurnId,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBodyDataMessages) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyDataMessages) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetContent() *string {
	return s.Content
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetFeedback() *string {
	return s.Feedback
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetRole() *string {
	return s.Role
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) GetTurnId() *string {
	return s.TurnId
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetContent(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.Content = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetCreateTime(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.CreateTime = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetFeedback(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.Feedback = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetRole(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.Role = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) SetTurnId(v string) *KopilotListConversationChatMessagesResponseBodyDataMessages {
	s.TurnId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataMessages) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo struct {
	// The time when the overview was generated, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-09-17T12:00:00Z
	AsOf *string `json:"AsOf,omitempty" xml:"AsOf,omitempty"`
	// The number of associated tasks in the ENABLED state.
	//
	// example:
	//
	// 1
	EnabledCount *int64 `json:"EnabledCount,omitempty" xml:"EnabledCount,omitempty"`
	// Indicates whether there is a next page of associated tasks.
	//
	// example:
	//
	// false
	HasMoreTasks *bool `json:"HasMoreTasks,omitempty" xml:"HasMoreTasks,omitempty"`
	// Indicates whether the current session has associated scheduled tasks in the enabled, paused, or pending authorization state.
	//
	// example:
	//
	// true
	HasScheduledTask *bool `json:"HasScheduledTask,omitempty" xml:"HasScheduledTask,omitempty"`
	// The cursor for the next page. This value is empty if there is no next page.
	//
	// example:
	//
	// 123
	NextTaskCursor *string `json:"NextTaskCursor,omitempty" xml:"NextTaskCursor,omitempty"`
	// The total number of associated tasks. Only tasks in the ENABLED, PAUSED, or NEEDS_AUTH state are counted.
	//
	// example:
	//
	// 1
	TaskCount *int64 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
	// The list of associated tasks on the current page.
	Tasks []*KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) GetAsOf() *string {
	return s.AsOf
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) GetEnabledCount() *int64 {
	return s.EnabledCount
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) GetHasMoreTasks() *bool {
	return s.HasMoreTasks
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) GetHasScheduledTask() *bool {
	return s.HasScheduledTask
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) GetNextTaskCursor() *string {
	return s.NextTaskCursor
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) GetTaskCount() *int64 {
	return s.TaskCount
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) GetTasks() []*KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks {
	return s.Tasks
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) SetAsOf(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo {
	s.AsOf = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) SetEnabledCount(v int64) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo {
	s.EnabledCount = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) SetHasMoreTasks(v bool) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo {
	s.HasMoreTasks = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) SetHasScheduledTask(v bool) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo {
	s.HasScheduledTask = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) SetNextTaskCursor(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo {
	s.NextTaskCursor = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) SetTaskCount(v int64) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo {
	s.TaskCount = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) SetTasks(v []*KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo {
	s.Tasks = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfo) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks struct {
	// The run record that is currently queued or running. This value is empty if there is no active run.
	ActiveRun *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun `json:"ActiveRun,omitempty" xml:"ActiveRun,omitempty" type:"Struct"`
	// The most recent completed run record, including failed runs. This value is empty if no record exists.
	LastCompletedRun *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun `json:"LastCompletedRun,omitempty" xml:"LastCompletedRun,omitempty" type:"Struct"`
	// The name of the scheduled task.
	//
	// example:
	//
	// Kafka Resource Inspection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next scheduled execution time, in UTC ISO 8601 format. This value is empty if no next execution is scheduled.
	//
	// example:
	//
	// 2026-09-17T12:15:00Z
	NextRunAt *string `json:"NextRunAt,omitempty" xml:"NextRunAt,omitempty"`
	// The human-readable description of the execution schedule.
	//
	// example:
	//
	// Every 900 seconds
	ScheduleDescription *string `json:"ScheduleDescription,omitempty" xml:"ScheduleDescription,omitempty"`
	// The status of the scheduled task. Valid values:
	//
	// - DRAFT: The task is a draft.
	//
	// - ENABLED: The task is enabled.
	//
	// - PAUSED: The task is paused.
	//
	// - NEEDS_AUTH: The task is pending authorization.
	//
	// - COMPLETED: The task is completed.
	//
	// This status is independent of the run status.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unique identifier of the scheduled task.
	//
	// example:
	//
	// task_0123456789abcdef0123456789abcdef
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) GetActiveRun() *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun {
	return s.ActiveRun
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) GetLastCompletedRun() *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun {
	return s.LastCompletedRun
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) GetName() *string {
	return s.Name
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) GetScheduleDescription() *string {
	return s.ScheduleDescription
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) GetStatus() *string {
	return s.Status
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) SetActiveRun(v *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks {
	s.ActiveRun = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) SetLastCompletedRun(v *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks {
	s.LastCompletedRun = v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) SetName(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks {
	s.Name = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) SetNextRunAt(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks {
	s.NextRunAt = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) SetScheduleDescription(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks {
	s.ScheduleDescription = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) SetStatus(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks {
	s.Status = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) SetTaskId(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks {
	s.TaskId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasks) Validate() error {
	if s.ActiveRun != nil {
		if err := s.ActiveRun.Validate(); err != nil {
			return err
		}
	}
	if s.LastCompletedRun != nil {
		if err := s.LastCompletedRun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun struct {
	// The time when the run ended, in UTC ISO 8601 format. This value is typically empty for queued or running tasks.
	//
	// example:
	//
	// 2026-09-17T12:01:00Z
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The unique identifier of a single run.
	//
	// example:
	//
	// run_0123456789abcdef0123456789abcdef
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The status of a single run. A value of QUEUED indicates that the run is queued. A value of RUNNING indicates that the run is in progress.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) GetRunId() *string {
	return s.RunId
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) GetStatus() *string {
	return s.Status
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) SetFinishedAt(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun {
	s.FinishedAt = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) SetRunId(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun {
	s.RunId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) SetStatus(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun {
	s.Status = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksActiveRun) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun struct {
	// The time when the run ended, in UTC ISO 8601 format. This value is empty if the run has not ended.
	//
	// example:
	//
	// 2026-09-17T12:01:00Z
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The unique identifier of a single run.
	//
	// example:
	//
	// run_0123456789abcdef0123456789abcdef
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The status of the most recent completed run. For example, SUCCEEDED indicates success and FAILED indicates failure.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) GetRunId() *string {
	return s.RunId
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) GetStatus() *string {
	return s.Status
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) SetFinishedAt(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun {
	s.FinishedAt = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) SetRunId(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun {
	s.RunId = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) SetStatus(v string) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun {
	s.Status = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskInfoTasksLastCompletedRun) Validate() error {
	return dara.Validate(s)
}

type KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota struct {
	// The maximum number of tasks or channels allowed, subject to the actual configuration.
	//
	// example:
	//
	// 3
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The remaining quota, calculated as the limit minus the used quota. The minimum value is 0.
	//
	// example:
	//
	// 2
	Remaining *int64 `json:"Remaining,omitempty" xml:"Remaining,omitempty"`
	// The used task quota. Tasks in the DRAFT, ENABLED, PAUSED, or NEEDS_AUTH state are counted. Completed or deleted tasks do not consume the quota.
	//
	// example:
	//
	// 1
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) GetLimit() *int32 {
	return s.Limit
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) GetRemaining() *int64 {
	return s.Remaining
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) GetUsed() *int64 {
	return s.Used
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) SetLimit(v int32) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota {
	s.Limit = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) SetRemaining(v int64) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota {
	s.Remaining = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) SetUsed(v int64) *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota {
	s.Used = &v
	return s
}

func (s *KopilotListConversationChatMessagesResponseBodyDataScheduledTaskQuota) Validate() error {
	return dara.Validate(s)
}
