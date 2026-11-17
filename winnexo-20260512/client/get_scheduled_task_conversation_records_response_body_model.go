// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskConversationRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScheduledTaskConversationRecordsResponseBody
	GetCode() *string
	SetHasMore(v bool) *GetScheduledTaskConversationRecordsResponseBody
	GetHasMore() *bool
	SetMessage(v string) *GetScheduledTaskConversationRecordsResponseBody
	GetMessage() *string
	SetPage(v int32) *GetScheduledTaskConversationRecordsResponseBody
	GetPage() *int32
	SetPageSize(v int32) *GetScheduledTaskConversationRecordsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetScheduledTaskConversationRecordsResponseBody
	GetRequestId() *string
	SetTasks(v []*GetScheduledTaskConversationRecordsResponseBodyTasks) *GetScheduledTaskConversationRecordsResponseBody
	GetTasks() []*GetScheduledTaskConversationRecordsResponseBodyTasks
	SetTotal(v int64) *GetScheduledTaskConversationRecordsResponseBody
	GetTotal() *int64
}

type GetScheduledTaskConversationRecordsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether a next page exists.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number. Default value: 1. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The collaboration group ID, such as cg_101. If you specify this parameter, the data is returned by group dimension, and the caller must be a valid group member. If you do not specify this parameter, the data is returned by personal dimension, excluding group tasks.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of tasks.
	//
	// example:
	//
	// []
	Tasks []*GetScheduledTaskConversationRecordsResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// The total number of Logtail pipeline configurations in the current project.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetScheduledTaskConversationRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskConversationRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskConversationRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScheduledTaskConversationRecordsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetScheduledTaskConversationRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduledTaskConversationRecordsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *GetScheduledTaskConversationRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetScheduledTaskConversationRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledTaskConversationRecordsResponseBody) GetTasks() []*GetScheduledTaskConversationRecordsResponseBodyTasks {
	return s.Tasks
}

func (s *GetScheduledTaskConversationRecordsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetScheduledTaskConversationRecordsResponseBody) SetCode(v string) *GetScheduledTaskConversationRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBody) SetHasMore(v bool) *GetScheduledTaskConversationRecordsResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBody) SetMessage(v string) *GetScheduledTaskConversationRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBody) SetPage(v int32) *GetScheduledTaskConversationRecordsResponseBody {
	s.Page = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBody) SetPageSize(v int32) *GetScheduledTaskConversationRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBody) SetRequestId(v string) *GetScheduledTaskConversationRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBody) SetTasks(v []*GetScheduledTaskConversationRecordsResponseBodyTasks) *GetScheduledTaskConversationRecordsResponseBody {
	s.Tasks = v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBody) SetTotal(v int64) *GetScheduledTaskConversationRecordsResponseBody {
	s.Total = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBody) Validate() error {
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

type GetScheduledTaskConversationRecordsResponseBodyTasks struct {
	// Pre-filters by a single task. If you do not specify this parameter, the execution records of all visible tasks are returned.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The list of sessions.
	//
	// example:
	//
	// []
	Conversations []*GetScheduledTaskConversationRecordsResponseBodyTasksConversations `json:"conversations,omitempty" xml:"conversations,omitempty" type:"Repeated"`
	// The cron expression.
	//
	// example:
	//
	// string_value
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether public access is enabled.
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// The execution model tier. Valid values: flagship (Flagship), standard (Standard), and quick (Lightweight).
	//
	// example:
	//
	// standard
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The task ID.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The time zone.
	//
	// > Default value: UTC+8.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The trigger type. Valid values:
	//
	// - Manual: Manual execution.
	//
	// - Cron: Scheduled trigger.
	//
	// example:
	//
	// string_value
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetScheduledTaskConversationRecordsResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskConversationRecordsResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetConversations() []*GetScheduledTaskConversationRecordsResponseBodyTasksConversations {
	return s.Conversations
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetModel() *string {
	return s.Model
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetName() *string {
	return s.Name
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetTimezone() *string {
	return s.Timezone
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetCollaborationGroupId(v string) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetConversations(v []*GetScheduledTaskConversationRecordsResponseBodyTasksConversations) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.Conversations = v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetCronExpression(v string) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.CronExpression = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetDescription(v string) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetIsOpen(v bool) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.IsOpen = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetModel(v string) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.Model = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetName(v string) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetTaskId(v string) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetTimezone(v string) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.Timezone = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) SetTriggerType(v string) *GetScheduledTaskConversationRecordsResponseBodyTasks {
	s.TriggerType = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasks) Validate() error {
	if s.Conversations != nil {
		for _, item := range s.Conversations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScheduledTaskConversationRecordsResponseBodyTasksConversations struct {
	// The actual group topic ID, which can be associated in queryWorkItemFeed. This ID cannot be used as a Workflow workItemId.
	//
	// example:
	//
	// exampleConversationId
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// The name of the schedule location.
	//
	// example:
	//
	// Daily sales weekly report
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Indicates whether unread messages exist.
	//
	// example:
	//
	// false
	HasUnread *bool `json:"hasUnread,omitempty" xml:"hasUnread,omitempty"`
	// The ID of the latest execution record.
	//
	// example:
	//
	// exampleExecutionId
	LatestExecutionId *string `json:"latestExecutionId,omitempty" xml:"latestExecutionId,omitempty"`
	// The latest execution status.
	//
	// example:
	//
	// success
	LatestStatus *string `json:"latestStatus,omitempty" xml:"latestStatus,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-11-16T09:30:00+08:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetScheduledTaskConversationRecordsResponseBodyTasksConversations) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskConversationRecordsResponseBodyTasksConversations) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) GetHasUnread() *bool {
	return s.HasUnread
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) GetLatestExecutionId() *string {
	return s.LatestExecutionId
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) GetLatestStatus() *string {
	return s.LatestStatus
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) SetConversationId(v string) *GetScheduledTaskConversationRecordsResponseBodyTasksConversations {
	s.ConversationId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) SetDisplayName(v string) *GetScheduledTaskConversationRecordsResponseBodyTasksConversations {
	s.DisplayName = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) SetHasUnread(v bool) *GetScheduledTaskConversationRecordsResponseBodyTasksConversations {
	s.HasUnread = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) SetLatestExecutionId(v string) *GetScheduledTaskConversationRecordsResponseBodyTasksConversations {
	s.LatestExecutionId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) SetLatestStatus(v string) *GetScheduledTaskConversationRecordsResponseBodyTasksConversations {
	s.LatestStatus = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) SetUpdatedAt(v string) *GetScheduledTaskConversationRecordsResponseBodyTasksConversations {
	s.UpdatedAt = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsResponseBodyTasksConversations) Validate() error {
	return dara.Validate(s)
}
