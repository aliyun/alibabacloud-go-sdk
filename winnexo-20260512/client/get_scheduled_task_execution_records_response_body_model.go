// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskExecutionRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScheduledTaskExecutionRecordsResponseBody
	GetCode() *string
	SetHasMore(v bool) *GetScheduledTaskExecutionRecordsResponseBody
	GetHasMore() *bool
	SetMessage(v string) *GetScheduledTaskExecutionRecordsResponseBody
	GetMessage() *string
	SetPage(v int32) *GetScheduledTaskExecutionRecordsResponseBody
	GetPage() *int32
	SetPageSize(v int32) *GetScheduledTaskExecutionRecordsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetScheduledTaskExecutionRecordsResponseBody
	GetRequestId() *string
	SetTasks(v []*GetScheduledTaskExecutionRecordsResponseBodyTasks) *GetScheduledTaskExecutionRecordsResponseBody
	GetTasks() []*GetScheduledTaskExecutionRecordsResponseBodyTasks
	SetTotal(v int64) *GetScheduledTaskExecutionRecordsResponseBody
	GetTotal() *int64
}

type GetScheduledTaskExecutionRecordsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether more data is available.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of tasks per page.
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
	// The task list.
	Tasks []*GetScheduledTaskExecutionRecordsResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// The total number of tasks.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetScheduledTaskExecutionRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetTasks() []*GetScheduledTaskExecutionRecordsResponseBodyTasks {
	return s.Tasks
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetCode(v string) *GetScheduledTaskExecutionRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetHasMore(v bool) *GetScheduledTaskExecutionRecordsResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetMessage(v string) *GetScheduledTaskExecutionRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetPage(v int32) *GetScheduledTaskExecutionRecordsResponseBody {
	s.Page = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetPageSize(v int32) *GetScheduledTaskExecutionRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetRequestId(v string) *GetScheduledTaskExecutionRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetTasks(v []*GetScheduledTaskExecutionRecordsResponseBodyTasks) *GetScheduledTaskExecutionRecordsResponseBody {
	s.Tasks = v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetTotal(v int64) *GetScheduledTaskExecutionRecordsResponseBody {
	s.Total = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) Validate() error {
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

type GetScheduledTaskExecutionRecordsResponseBodyTasks struct {
	// The ID of the collaboration group to which the task belongs. If empty, the task is a personal task.
	//
	// example:
	//
	// exampleCollaborationGroupId
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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
	// The execution model tier. Valid values:
	//
	// - flagship: flagship.
	//
	// - standard: standard.
	//
	// - quick: lightweight.
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
	// The timeline.
	Timeline []*GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline `json:"timeline,omitempty" xml:"timeline,omitempty" type:"Repeated"`
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
	// - Manual: manually executed.
	//
	// - Cron: triggered by a schedule.
	//
	// example:
	//
	// string_value
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetScheduledTaskExecutionRecordsResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionRecordsResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetIsOpen() *bool {
	return s.IsOpen
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetModel() *string {
	return s.Model
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetName() *string {
	return s.Name
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetTimeline() []*GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	return s.Timeline
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetTimezone() *string {
	return s.Timezone
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetCollaborationGroupId(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetCronExpression(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.CronExpression = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetDescription(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetIsOpen(v bool) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.IsOpen = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetModel(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.Model = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetName(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetTaskId(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetTimeline(v []*GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.Timeline = v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetTimezone(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.Timezone = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) SetTriggerType(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasks {
	s.TriggerType = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) Validate() error {
	if s.Timeline != nil {
		for _, item := range s.Timeline {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline struct {
	// The actual working hours, in hours.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	ActualTime *string `json:"actualTime,omitempty" xml:"actualTime,omitempty"`
	// The name of the schedule location.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The error message.
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The execution record ID.
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
	// Indicates whether the execution record has been archived due to expiration.
	//
	// example:
	//
	// false
	IsExpired *bool `json:"isExpired,omitempty" xml:"isExpired,omitempty"`
	// The execution output content (historical records only).
	//
	// example:
	//
	// string_value
	OutputContent *string `json:"outputContent,omitempty" xml:"outputContent,omitempty"`
	// The timed scheduling time.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	ScheduledTime *string `json:"scheduledTime,omitempty" xml:"scheduledTime,omitempty"`
	// The final status of the message.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GetActualTime() *string {
	return s.ActualTime
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GetIsExpired() *bool {
	return s.IsExpired
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GetOutputContent() *string {
	return s.OutputContent
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GetScheduledTime() *string {
	return s.ScheduledTime
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) GetStatus() *string {
	return s.Status
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) SetActualTime(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	s.ActualTime = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) SetDisplayName(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	s.DisplayName = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) SetErrorMessage(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	s.ErrorMessage = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) SetExecutionId(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	s.ExecutionId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) SetIsExpired(v bool) *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	s.IsExpired = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) SetOutputContent(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	s.OutputContent = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) SetScheduledTime(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	s.ScheduledTime = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) SetStatus(v string) *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline {
	s.Status = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline) Validate() error {
	return dara.Validate(s)
}
