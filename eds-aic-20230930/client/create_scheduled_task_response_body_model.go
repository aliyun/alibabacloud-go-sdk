// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScheduledTaskResponseBody
	GetCode() *string
	SetMessage(v string) *CreateScheduledTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateScheduledTaskResponseBody
	GetRequestId() *string
	SetTasks(v []*CreateScheduledTaskResponseBodyTasks) *CreateScheduledTaskResponseBody
	GetTasks() []*CreateScheduledTaskResponseBodyTasks
	SetTotalCount(v int32) *CreateScheduledTaskResponseBody
	GetTotalCount() *int32
}

type CreateScheduledTaskResponseBody struct {
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
	// 9A51B1DF-96FF-3BCC-B08C-783161D3****
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*CreateScheduledTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScheduledTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledTaskResponseBody) GetTasks() []*CreateScheduledTaskResponseBodyTasks {
	return s.Tasks
}

func (s *CreateScheduledTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CreateScheduledTaskResponseBody) SetCode(v string) *CreateScheduledTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetMessage(v string) *CreateScheduledTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetTasks(v []*CreateScheduledTaskResponseBodyTasks) *CreateScheduledTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetTotalCount(v int32) *CreateScheduledTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) Validate() error {
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

type CreateScheduledTaskResponseBodyTasks struct {
	InstanceResults []*CreateScheduledTaskResponseBodyTasksInstanceResults `json:"InstanceResults,omitempty" xml:"InstanceResults,omitempty" type:"Repeated"`
	// example:
	//
	// sch-260705-agb*****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// example:
	//
	// tsk-260705-0jj*****
	TaskConfigId *string `json:"TaskConfigId,omitempty" xml:"TaskConfigId,omitempty"`
}

func (s CreateScheduledTaskResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBodyTasks) GetInstanceResults() []*CreateScheduledTaskResponseBodyTasksInstanceResults {
	return s.InstanceResults
}

func (s *CreateScheduledTaskResponseBodyTasks) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *CreateScheduledTaskResponseBodyTasks) GetTaskConfigId() *string {
	return s.TaskConfigId
}

func (s *CreateScheduledTaskResponseBodyTasks) SetInstanceResults(v []*CreateScheduledTaskResponseBodyTasksInstanceResults) *CreateScheduledTaskResponseBodyTasks {
	s.InstanceResults = v
	return s
}

func (s *CreateScheduledTaskResponseBodyTasks) SetScheduledId(v string) *CreateScheduledTaskResponseBodyTasks {
	s.ScheduledId = &v
	return s
}

func (s *CreateScheduledTaskResponseBodyTasks) SetTaskConfigId(v string) *CreateScheduledTaskResponseBodyTasks {
	s.TaskConfigId = &v
	return s
}

func (s *CreateScheduledTaskResponseBodyTasks) Validate() error {
	if s.InstanceResults != nil {
		for _, item := range s.InstanceResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateScheduledTaskResponseBodyTasksInstanceResults struct {
	// example:
	//
	// privateAccount not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// acp-iuyb1zv1ap6nb****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScheduledTaskResponseBodyTasksInstanceResults) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskResponseBodyTasksInstanceResults) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBodyTasksInstanceResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateScheduledTaskResponseBodyTasksInstanceResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScheduledTaskResponseBodyTasksInstanceResults) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScheduledTaskResponseBodyTasksInstanceResults) SetErrorMessage(v string) *CreateScheduledTaskResponseBodyTasksInstanceResults {
	s.ErrorMessage = &v
	return s
}

func (s *CreateScheduledTaskResponseBodyTasksInstanceResults) SetInstanceId(v string) *CreateScheduledTaskResponseBodyTasksInstanceResults {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduledTaskResponseBodyTasksInstanceResults) SetSuccess(v bool) *CreateScheduledTaskResponseBodyTasksInstanceResults {
	s.Success = &v
	return s
}

func (s *CreateScheduledTaskResponseBodyTasksInstanceResults) Validate() error {
	return dara.Validate(s)
}
