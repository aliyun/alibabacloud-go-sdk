// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoftwarelibDistributeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSoftwarelibDistributeTaskResponseBody
	GetRequestId() *string
	SetTask(v *CreateSoftwarelibDistributeTaskResponseBodyTask) *CreateSoftwarelibDistributeTaskResponseBody
	GetTask() *CreateSoftwarelibDistributeTaskResponseBodyTask
}

type CreateSoftwarelibDistributeTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B608C6AE-623D-55C4-9454-601B88AE937E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the created task.
	Task *CreateSoftwarelibDistributeTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s CreateSoftwarelibDistributeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSoftwarelibDistributeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSoftwarelibDistributeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSoftwarelibDistributeTaskResponseBody) GetTask() *CreateSoftwarelibDistributeTaskResponseBodyTask {
	return s.Task
}

func (s *CreateSoftwarelibDistributeTaskResponseBody) SetRequestId(v string) *CreateSoftwarelibDistributeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponseBody) SetTask(v *CreateSoftwarelibDistributeTaskResponseBodyTask) *CreateSoftwarelibDistributeTaskResponseBody {
	s.Task = v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSoftwarelibDistributeTaskResponseBodyTask struct {
	// The task creation time as a second-level UNIX timestamp.
	//
	// example:
	//
	// 1782268092
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The task name.
	//
	// example:
	//
	// test_task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The software name.
	//
	// example:
	//
	// test software
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// The task status. Valid values:
	//
	// - **enabled**: enabled.
	//
	// - **disabled**: disabled.
	//
	// The initial status of a task after creation is disabled.
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The operating system to which the task applies. Valid values:
	//
	// - **Windows**: Windows.
	//
	// - **Mac(Apple)**: macOS with Apple silicon.
	//
	// - **Mac(Intel)**: macOS with Intel processors.
	//
	// example:
	//
	// Windows
	SupportOs *string `json:"SupportOs,omitempty" xml:"SupportOs,omitempty"`
	// The task ID, which is used to query the task execution result.
	//
	// example:
	//
	// softwarelib-distribute-task-911dd7898bc2****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSoftwarelibDistributeTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s CreateSoftwarelibDistributeTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) GetName() *string {
	return s.Name
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) GetSoftwareName() *string {
	return s.SoftwareName
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) GetStatus() *string {
	return s.Status
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) GetSupportOs() *string {
	return s.SupportOs
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) SetCreateTime(v string) *CreateSoftwarelibDistributeTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) SetName(v string) *CreateSoftwarelibDistributeTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) SetSoftwareName(v string) *CreateSoftwarelibDistributeTaskResponseBodyTask {
	s.SoftwareName = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) SetStatus(v string) *CreateSoftwarelibDistributeTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) SetSupportOs(v string) *CreateSoftwarelibDistributeTaskResponseBodyTask {
	s.SupportOs = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) SetTaskId(v string) *CreateSoftwarelibDistributeTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *CreateSoftwarelibDistributeTaskResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
