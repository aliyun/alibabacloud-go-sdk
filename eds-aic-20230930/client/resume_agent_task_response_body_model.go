// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeAgentTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeAgentTaskResponseBody
	GetCode() *string
	SetMessage(v string) *ResumeAgentTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeAgentTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*ResumeAgentTaskResponseBodyTasks) *ResumeAgentTaskResponseBody
	GetTasks() []*ResumeAgentTaskResponseBodyTasks
}

type ResumeAgentTaskResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ResumeAgentTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ResumeAgentTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeAgentTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeAgentTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeAgentTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeAgentTaskResponseBody) GetTasks() []*ResumeAgentTaskResponseBodyTasks {
	return s.Tasks
}

func (s *ResumeAgentTaskResponseBody) SetCode(v string) *ResumeAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeAgentTaskResponseBody) SetMessage(v string) *ResumeAgentTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeAgentTaskResponseBody) SetRequestId(v string) *ResumeAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeAgentTaskResponseBody) SetTasks(v []*ResumeAgentTaskResponseBodyTasks) *ResumeAgentTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *ResumeAgentTaskResponseBody) Validate() error {
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

type ResumeAgentTaskResponseBodyTasks struct {
	// example:
	//
	// RUNNING
	CurrentStatus *string `json:"CurrentStatus,omitempty" xml:"CurrentStatus,omitempty"`
	// example:
	//
	// Task status [COMPLETED] does not support resume, only PAUSED tasks can be resumed.
	FailedReason *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
	// example:
	//
	// acp-ek65k51zoxia3x8xz
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2026-04-13T17:42:19Z
	ResumingAt *string `json:"ResumingAt,omitempty" xml:"ResumingAt,omitempty"`
	// example:
	//
	// t-imr0fufqd7cle****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResumeAgentTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ResumeAgentTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ResumeAgentTaskResponseBodyTasks) GetCurrentStatus() *string {
	return s.CurrentStatus
}

func (s *ResumeAgentTaskResponseBodyTasks) GetFailedReason() *string {
	return s.FailedReason
}

func (s *ResumeAgentTaskResponseBodyTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResumeAgentTaskResponseBodyTasks) GetResumingAt() *string {
	return s.ResumingAt
}

func (s *ResumeAgentTaskResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ResumeAgentTaskResponseBodyTasks) SetCurrentStatus(v string) *ResumeAgentTaskResponseBodyTasks {
	s.CurrentStatus = &v
	return s
}

func (s *ResumeAgentTaskResponseBodyTasks) SetFailedReason(v string) *ResumeAgentTaskResponseBodyTasks {
	s.FailedReason = &v
	return s
}

func (s *ResumeAgentTaskResponseBodyTasks) SetInstanceId(v string) *ResumeAgentTaskResponseBodyTasks {
	s.InstanceId = &v
	return s
}

func (s *ResumeAgentTaskResponseBodyTasks) SetResumingAt(v string) *ResumeAgentTaskResponseBodyTasks {
	s.ResumingAt = &v
	return s
}

func (s *ResumeAgentTaskResponseBodyTasks) SetTaskId(v string) *ResumeAgentTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ResumeAgentTaskResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
