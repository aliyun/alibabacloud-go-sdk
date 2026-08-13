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
	SetMessage(v string) *GetScheduledTaskExecutionRecordsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScheduledTaskExecutionRecordsResponseBody
	GetRequestId() *string
	SetTasks(v []*GetScheduledTaskExecutionRecordsResponseBodyTasks) *GetScheduledTaskExecutionRecordsResponseBody
	GetTasks() []*GetScheduledTaskExecutionRecordsResponseBodyTasks
}

type GetScheduledTaskExecutionRecordsResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string                                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Tasks     []*GetScheduledTaskExecutionRecordsResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
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

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) GetTasks() []*GetScheduledTaskExecutionRecordsResponseBodyTasks {
	return s.Tasks
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetCode(v string) *GetScheduledTaskExecutionRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsResponseBody) SetMessage(v string) *GetScheduledTaskExecutionRecordsResponseBody {
	s.Message = &v
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
	// Cron 表达式
	//
	// example:
	//
	// string_value
	CronExpression *string `json:"cronExpression,omitempty" xml:"cronExpression,omitempty"`
	// 任务简述
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 是否公开
	//
	// example:
	//
	// true
	IsOpen *bool `json:"isOpen,omitempty" xml:"isOpen,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 任务 ID
	//
	// example:
	//
	// exampleTaskId
	TaskId   *string                                                      `json:"taskId,omitempty" xml:"taskId,omitempty"`
	Timeline []*GetScheduledTaskExecutionRecordsResponseBodyTasksTimeline `json:"timeline,omitempty" xml:"timeline,omitempty" type:"Repeated"`
	// 时区
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// 触发类型 cron/manual/event
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

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetDescription() *string {
	return s.Description
}

func (s *GetScheduledTaskExecutionRecordsResponseBodyTasks) GetIsOpen() *bool {
	return s.IsOpen
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
	// 实际执行时间（仅历史记录）
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	ActualTime *string `json:"actualTime,omitempty" xml:"actualTime,omitempty"`
	// 执行记录展示名称
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// 错误信息（仅失败记录）
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 执行记录 ID（历史记录才有）
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
	// 执行输出内容（仅历史记录）
	//
	// example:
	//
	// string_value
	OutputContent *string `json:"outputContent,omitempty" xml:"outputContent,omitempty"`
	// 计划执行时间 ISO8601
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	ScheduledTime *string `json:"scheduledTime,omitempty" xml:"scheduledTime,omitempty"`
	// 状态：PENDING/RUNNING/SUCCESS/FAILED/SCHEDULED
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
