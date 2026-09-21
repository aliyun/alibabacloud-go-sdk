// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudPhoneNodeSystemImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudPhoneNodeSystemImageResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateCloudPhoneNodeSystemImageResponseBody
	GetTaskId() *string
	SetTasks(v *UpdateCloudPhoneNodeSystemImageResponseBodyTasks) *UpdateCloudPhoneNodeSystemImageResponseBody
	GetTasks() *UpdateCloudPhoneNodeSystemImageResponseBodyTasks
}

type UpdateCloudPhoneNodeSystemImageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-ehs0yoedj0xe9****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task information.
	Tasks *UpdateCloudPhoneNodeSystemImageResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s UpdateCloudPhoneNodeSystemImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPhoneNodeSystemImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBody) GetTasks() *UpdateCloudPhoneNodeSystemImageResponseBodyTasks {
	return s.Tasks
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBody) SetRequestId(v string) *UpdateCloudPhoneNodeSystemImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBody) SetTaskId(v string) *UpdateCloudPhoneNodeSystemImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBody) SetTasks(v *UpdateCloudPhoneNodeSystemImageResponseBodyTasks) *UpdateCloudPhoneNodeSystemImageResponseBody {
	s.Tasks = v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBody) Validate() error {
	if s.Tasks != nil {
		if err := s.Tasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCloudPhoneNodeSystemImageResponseBodyTasks struct {
	// The child tasks.
	ChildTasks []*UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks `json:"ChildTasks,omitempty" xml:"ChildTasks,omitempty" type:"Repeated"`
	// The parent task ID.
	//
	// example:
	//
	// t-xxxx
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
}

func (s UpdateCloudPhoneNodeSystemImageResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPhoneNodeSystemImageResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasks) GetChildTasks() []*UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks {
	return s.ChildTasks
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasks) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasks) SetChildTasks(v []*UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks) *UpdateCloudPhoneNodeSystemImageResponseBodyTasks {
	s.ChildTasks = v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasks) SetParentTaskId(v string) *UpdateCloudPhoneNodeSystemImageResponseBodyTasks {
	s.ParentTaskId = &v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasks) Validate() error {
	if s.ChildTasks != nil {
		for _, item := range s.ChildTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks struct {
	// The node ID.
	//
	// example:
	//
	// cpn-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The child task ID.
	//
	// example:
	//
	// t-xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks) GoString() string {
	return s.String()
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks) SetInstanceId(v string) *UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks) SetTaskId(v string) *UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks {
	s.TaskId = &v
	return s
}

func (s *UpdateCloudPhoneNodeSystemImageResponseBodyTasksChildTasks) Validate() error {
	return dara.Validate(s)
}
