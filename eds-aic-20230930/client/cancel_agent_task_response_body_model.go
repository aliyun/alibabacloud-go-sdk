// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelAgentTaskResponseBody
	GetCode() *string
	SetMessage(v string) *CancelAgentTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelAgentTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*CancelAgentTaskResponseBodyTasks) *CancelAgentTaskResponseBody
	GetTasks() []*CancelAgentTaskResponseBodyTasks
}

type CancelAgentTaskResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B886792-2051-5DB4-8AE6-C8E45D3B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of task objects.
	Tasks []*CancelAgentTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s CancelAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAgentTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelAgentTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAgentTaskResponseBody) GetTasks() []*CancelAgentTaskResponseBodyTasks {
	return s.Tasks
}

func (s *CancelAgentTaskResponseBody) SetCode(v string) *CancelAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAgentTaskResponseBody) SetMessage(v string) *CancelAgentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAgentTaskResponseBody) SetRequestId(v string) *CancelAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAgentTaskResponseBody) SetTasks(v []*CancelAgentTaskResponseBodyTasks) *CancelAgentTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *CancelAgentTaskResponseBody) Validate() error {
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

type CancelAgentTaskResponseBodyTasks struct {
	// The task\\"s cancellation time, in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-13T17:42:19Z
	CancelAt *string `json:"CancelAt,omitempty" xml:"CancelAt,omitempty"`
	// The task status after the cancellation request. Possible values include:
	//
	// `CANCELLING` or `CANCELED`: The task is being or has been canceled.
	//
	// `COMPLETED`: The task was already complete and could not be canceled.
	//
	// example:
	//
	// COMPLETED
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// Task status [COMPLETED] does not support cancellation, only PENDING/RUNNING/CANCELLING tasks can be canceled.
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// The mobile node ID.
	//
	// example:
	//
	// acp-ek65k51zoxia3x8xz
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task status before cancellation. Valid values include:
	//
	// PENDING: The task is being created.
	//
	// RUNNING: The task is running.
	//
	// example:
	//
	// RUNNING
	PreviousStatus *string `json:"PreviousStatus,omitempty" xml:"PreviousStatus,omitempty"`
	// The globally unique task ID.
	//
	// example:
	//
	// t-imr0fufqd7cle****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelAgentTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *CancelAgentTaskResponseBodyTasks) GetCancelAt() *string {
	return s.CancelAt
}

func (s *CancelAgentTaskResponseBodyTasks) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *CancelAgentTaskResponseBodyTasks) GetFailedReason() *string {
	return s.FailedReason
}

func (s *CancelAgentTaskResponseBodyTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelAgentTaskResponseBodyTasks) GetPreviousStatus() *string {
	return s.PreviousStatus
}

func (s *CancelAgentTaskResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelAgentTaskResponseBodyTasks) SetCancelAt(v string) *CancelAgentTaskResponseBodyTasks {
	s.CancelAt = &v
	return s
}

func (s *CancelAgentTaskResponseBodyTasks) SetCurrentStatus(v string) *CancelAgentTaskResponseBodyTasks {
	s.CurrentStatus = &v
	return s
}

func (s *CancelAgentTaskResponseBodyTasks) SetFailedReason(v string) *CancelAgentTaskResponseBodyTasks {
	s.FailedReason = &v
	return s
}

func (s *CancelAgentTaskResponseBodyTasks) SetInstanceId(v string) *CancelAgentTaskResponseBodyTasks {
	s.InstanceId = &v
	return s
}

func (s *CancelAgentTaskResponseBodyTasks) SetPreviousStatus(v string) *CancelAgentTaskResponseBodyTasks {
	s.PreviousStatus = &v
	return s
}

func (s *CancelAgentTaskResponseBodyTasks) SetTaskId(v string) *CancelAgentTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *CancelAgentTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
