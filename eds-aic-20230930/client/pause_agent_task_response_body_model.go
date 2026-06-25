// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PauseAgentTaskResponseBody
	GetCode() *string
	SetMessage(v string) *PauseAgentTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *PauseAgentTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*PauseAgentTaskResponseBodyTasks) *PauseAgentTaskResponseBody
	GetTasks() []*PauseAgentTaskResponseBodyTasks
}

type PauseAgentTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// E38B41A8-8E00-5AE4-A957-6636ACB8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of tasks.
	Tasks []*PauseAgentTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s PauseAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PauseAgentTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *PauseAgentTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PauseAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseAgentTaskResponseBody) GetTasks() []*PauseAgentTaskResponseBodyTasks {
	return s.Tasks
}

func (s *PauseAgentTaskResponseBody) SetCode(v string) *PauseAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *PauseAgentTaskResponseBody) SetMessage(v string) *PauseAgentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *PauseAgentTaskResponseBody) SetRequestId(v string) *PauseAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseAgentTaskResponseBody) SetTasks(v []*PauseAgentTaskResponseBodyTasks) *PauseAgentTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *PauseAgentTaskResponseBody) Validate() error {
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

type PauseAgentTaskResponseBodyTasks struct {
	// The current status of the task. The following are possible values:
	//
	// PENDING: The task is being created.
	//
	// RUNNING: The task is running.
	//
	// COMPLETED: The task is completed.
	//
	// FAILED: The task failed.
	//
	// TIMEOUT: The task timed out.
	//
	// PAUSING: The task is being paused.
	//
	// PAUSED: The task is paused.
	//
	// example:
	//
	// PAUSING
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// The reason the task failed to pause.
	//
	// example:
	//
	// Task status [COMPLETED] does not support pause, only RUNNING tasks can be paused.
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The ID of the Mobile node.
	//
	// example:
	//
	// acp-anzzuho371azi44xr
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the pause request was initiated, in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-13T17:42:19Z
	PausingAt *string `json:"PausingAt,omitempty" xml:"PausingAt,omitempty"`
	// The status of the task before the pause request. The only valid value is:
	//
	// RUNNING: The task is running.
	//
	// example:
	//
	// RUNNING
	PreviousStatus *string `json:"PreviousStatus,omitempty" xml:"PreviousStatus,omitempty"`
	// The unique ID of the task.
	//
	// example:
	//
	// t-imr0fufqd7cle****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PauseAgentTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s PauseAgentTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *PauseAgentTaskResponseBodyTasks) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *PauseAgentTaskResponseBodyTasks) GetFailedReason() *string {
	return s.FailedReason
}

func (s *PauseAgentTaskResponseBodyTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PauseAgentTaskResponseBodyTasks) GetPausingAt() *string {
	return s.PausingAt
}

func (s *PauseAgentTaskResponseBodyTasks) GetPreviousStatus() *string {
	return s.PreviousStatus
}

func (s *PauseAgentTaskResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *PauseAgentTaskResponseBodyTasks) SetCurrentStatus(v string) *PauseAgentTaskResponseBodyTasks {
	s.CurrentStatus = &v
	return s
}

func (s *PauseAgentTaskResponseBodyTasks) SetFailedReason(v string) *PauseAgentTaskResponseBodyTasks {
	s.FailedReason = &v
	return s
}

func (s *PauseAgentTaskResponseBodyTasks) SetInstanceId(v string) *PauseAgentTaskResponseBodyTasks {
	s.InstanceId = &v
	return s
}

func (s *PauseAgentTaskResponseBodyTasks) SetPausingAt(v string) *PauseAgentTaskResponseBodyTasks {
	s.PausingAt = &v
	return s
}

func (s *PauseAgentTaskResponseBodyTasks) SetPreviousStatus(v string) *PauseAgentTaskResponseBodyTasks {
	s.PreviousStatus = &v
	return s
}

func (s *PauseAgentTaskResponseBodyTasks) SetTaskId(v string) *PauseAgentTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *PauseAgentTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
