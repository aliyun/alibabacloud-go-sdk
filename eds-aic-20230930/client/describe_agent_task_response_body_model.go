// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAgentTaskResponseBody
	GetCode() *string
	SetCount(v int32) *DescribeAgentTaskResponseBody
	GetCount() *int32
	SetMessage(v string) *DescribeAgentTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAgentTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*DescribeAgentTaskResponseBodyTasks) *DescribeAgentTaskResponseBody
	GetTasks() []*DescribeAgentTaskResponseBodyTasks
}

type DescribeAgentTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// For example, "200" indicates success.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task count.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 310A783E-CC46-5452-A8A3-71AE5DB5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of tasks.
	Tasks []*DescribeAgentTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DescribeAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgentTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAgentTaskResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAgentTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgentTaskResponseBody) GetTasks() []*DescribeAgentTaskResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeAgentTaskResponseBody) SetCode(v string) *DescribeAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAgentTaskResponseBody) SetCount(v int32) *DescribeAgentTaskResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAgentTaskResponseBody) SetMessage(v string) *DescribeAgentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAgentTaskResponseBody) SetRequestId(v string) *DescribeAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgentTaskResponseBody) SetTasks(v []*DescribeAgentTaskResponseBodyTasks) *DescribeAgentTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeAgentTaskResponseBody) Validate() error {
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

type DescribeAgentTaskResponseBodyTasks struct {
	// The task\\"s current status. Valid values:
	//
	// `PENDING`: The task is being created.
	//
	// `RUNNING`: The task is running.
	//
	// `COMPLETED`: The task has completed.
	//
	// `FAILED`: The task failed.
	//
	// `TIMEOUT`: The task timed out.
	//
	// example:
	//
	// COMPLETED
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// The mobile node ID.
	//
	// example:
	//
	// acp-anzzuho371azi44xr
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task\\"s creation time, in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-13T17:42:19Z
	RunningAt *string `json:"RunningAt,omitempty" xml:"RunningAt,omitempty"`
	// The number of steps executed.
	//
	// example:
	//
	// 30
	Steps *string `json:"Steps,omitempty" xml:"Steps,omitempty"`
	// The task duration. This field is returned only when `CurrentStatus` is `FAILED` or `COMPLETED`.
	//
	// example:
	//
	// 50
	TaskDuration *string `json:"TaskDuration,omitempty" xml:"TaskDuration,omitempty"`
	// The globally unique task ID.
	//
	// example:
	//
	// t-imr0fufqd7cle****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task result. This field is returned only when `CurrentStatus` is `COMPLETED` or `FAILED`.
	//
	// example:
	//
	// Download DingTalk succeeded.
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// The user prompt that the Agent uses to perform the task.
	//
	// example:
	//
	// Download DingTalk from App Store
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s DescribeAgentTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgentTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeAgentTaskResponseBodyTasks) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *DescribeAgentTaskResponseBodyTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAgentTaskResponseBodyTasks) GetRunningAt() *string {
	return s.RunningAt
}

func (s *DescribeAgentTaskResponseBodyTasks) GetSteps() *string {
	return s.Steps
}

func (s *DescribeAgentTaskResponseBodyTasks) GetTaskDuration() *string {
	return s.TaskDuration
}

func (s *DescribeAgentTaskResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeAgentTaskResponseBodyTasks) GetTaskResult() *string {
	return s.TaskResult
}

func (s *DescribeAgentTaskResponseBodyTasks) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *DescribeAgentTaskResponseBodyTasks) SetCurrentStatus(v string) *DescribeAgentTaskResponseBodyTasks {
	s.CurrentStatus = &v
	return s
}

func (s *DescribeAgentTaskResponseBodyTasks) SetInstanceId(v string) *DescribeAgentTaskResponseBodyTasks {
	s.InstanceId = &v
	return s
}

func (s *DescribeAgentTaskResponseBodyTasks) SetRunningAt(v string) *DescribeAgentTaskResponseBodyTasks {
	s.RunningAt = &v
	return s
}

func (s *DescribeAgentTaskResponseBodyTasks) SetSteps(v string) *DescribeAgentTaskResponseBodyTasks {
	s.Steps = &v
	return s
}

func (s *DescribeAgentTaskResponseBodyTasks) SetTaskDuration(v string) *DescribeAgentTaskResponseBodyTasks {
	s.TaskDuration = &v
	return s
}

func (s *DescribeAgentTaskResponseBodyTasks) SetTaskId(v string) *DescribeAgentTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeAgentTaskResponseBodyTasks) SetTaskResult(v string) *DescribeAgentTaskResponseBodyTasks {
	s.TaskResult = &v
	return s
}

func (s *DescribeAgentTaskResponseBodyTasks) SetUserPrompt(v string) *DescribeAgentTaskResponseBodyTasks {
	s.UserPrompt = &v
	return s
}

func (s *DescribeAgentTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
