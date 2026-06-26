// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunAgentTaskResponseBody
	GetCode() *string
	SetCount(v int32) *RunAgentTaskResponseBody
	GetCount() *int32
	SetMessage(v string) *RunAgentTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *RunAgentTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*RunAgentTaskResponseBodyTasks) *RunAgentTaskResponseBody
	GetTasks() []*RunAgentTaskResponseBodyTasks
}

type RunAgentTaskResponseBody struct {
	// The status code of the operation.
	//
	// example:
	//
	// For example, "200" indicates success.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of tasks.
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
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tasks.
	Tasks []*RunAgentTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s RunAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RunAgentTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunAgentTaskResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *RunAgentTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunAgentTaskResponseBody) GetTasks() []*RunAgentTaskResponseBodyTasks {
	return s.Tasks
}

func (s *RunAgentTaskResponseBody) SetCode(v string) *RunAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RunAgentTaskResponseBody) SetCount(v int32) *RunAgentTaskResponseBody {
	s.Count = &v
	return s
}

func (s *RunAgentTaskResponseBody) SetMessage(v string) *RunAgentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RunAgentTaskResponseBody) SetRequestId(v string) *RunAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunAgentTaskResponseBody) SetTasks(v []*RunAgentTaskResponseBodyTasks) *RunAgentTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *RunAgentTaskResponseBody) Validate() error {
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

type RunAgentTaskResponseBodyTasks struct {
	// The current status of the task. Valid values:
	//
	// - PENDING: The task is being created.
	//
	// - RUNNING: The task is running.
	//
	// - COMPLETED: The task is completed.
	//
	// - FAILED: The task failed.
	//
	// - TIMEOUT: The task execution timed out.
	//
	// example:
	//
	// COMPLETED
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// The Mobile node ID.
	//
	// example:
	//
	// acp-ek65k51zoxia3x8xz
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the task was created, in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-13T17:42:19Z
	RunningAt *string `json:"RunningAt,omitempty" xml:"RunningAt,omitempty"`
	// The task ID, which is globally unique.
	//
	// example:
	//
	// t-imr0fufqd7cle****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The user instruction in natural language. The Agent performs operations based on this instruction.
	//
	// example:
	//
	// Download DingTalk from App Store
	UserPrompt *string `json:"UserPrompt,omitempty" xml:"UserPrompt,omitempty"`
}

func (s RunAgentTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s RunAgentTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *RunAgentTaskResponseBodyTasks) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *RunAgentTaskResponseBodyTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunAgentTaskResponseBodyTasks) GetRunningAt() *string {
	return s.RunningAt
}

func (s *RunAgentTaskResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *RunAgentTaskResponseBodyTasks) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *RunAgentTaskResponseBodyTasks) SetCurrentStatus(v string) *RunAgentTaskResponseBodyTasks {
	s.CurrentStatus = &v
	return s
}

func (s *RunAgentTaskResponseBodyTasks) SetInstanceId(v string) *RunAgentTaskResponseBodyTasks {
	s.InstanceId = &v
	return s
}

func (s *RunAgentTaskResponseBodyTasks) SetRunningAt(v string) *RunAgentTaskResponseBodyTasks {
	s.RunningAt = &v
	return s
}

func (s *RunAgentTaskResponseBodyTasks) SetTaskId(v string) *RunAgentTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *RunAgentTaskResponseBodyTasks) SetUserPrompt(v string) *RunAgentTaskResponseBodyTasks {
	s.UserPrompt = &v
	return s
}

func (s *RunAgentTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
