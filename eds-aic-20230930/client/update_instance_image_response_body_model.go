// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceImageResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateInstanceImageResponseBody
	GetTaskId() *string
	SetTasks(v *UpdateInstanceImageResponseBodyTasks) *UpdateInstanceImageResponseBody
	GetTasks() *UpdateInstanceImageResponseBodyTasks
}

type UpdateInstanceImageResponseBody struct {
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// t-1ljew7on6ay0j****
	TaskId *string                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Tasks  *UpdateInstanceImageResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s UpdateInstanceImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceImageResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateInstanceImageResponseBody) GetTasks() *UpdateInstanceImageResponseBodyTasks {
	return s.Tasks
}

func (s *UpdateInstanceImageResponseBody) SetRequestId(v string) *UpdateInstanceImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceImageResponseBody) SetTaskId(v string) *UpdateInstanceImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateInstanceImageResponseBody) SetTasks(v *UpdateInstanceImageResponseBodyTasks) *UpdateInstanceImageResponseBody {
	s.Tasks = v
	return s
}

func (s *UpdateInstanceImageResponseBody) Validate() error {
	if s.Tasks != nil {
		if err := s.Tasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceImageResponseBodyTasks struct {
	ChildTasks []*UpdateInstanceImageResponseBodyTasksChildTasks `json:"ChildTasks,omitempty" xml:"ChildTasks,omitempty" type:"Repeated"`
	// example:
	//
	// t-xxxx
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
}

func (s UpdateInstanceImageResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceImageResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageResponseBodyTasks) GetChildTasks() []*UpdateInstanceImageResponseBodyTasksChildTasks {
	return s.ChildTasks
}

func (s *UpdateInstanceImageResponseBodyTasks) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *UpdateInstanceImageResponseBodyTasks) SetChildTasks(v []*UpdateInstanceImageResponseBodyTasksChildTasks) *UpdateInstanceImageResponseBodyTasks {
	s.ChildTasks = v
	return s
}

func (s *UpdateInstanceImageResponseBodyTasks) SetParentTaskId(v string) *UpdateInstanceImageResponseBodyTasks {
	s.ParentTaskId = &v
	return s
}

func (s *UpdateInstanceImageResponseBodyTasks) Validate() error {
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

type UpdateInstanceImageResponseBodyTasksChildTasks struct {
	// example:
	//
	// acp-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// t-xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateInstanceImageResponseBodyTasksChildTasks) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceImageResponseBodyTasksChildTasks) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageResponseBodyTasksChildTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceImageResponseBodyTasksChildTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateInstanceImageResponseBodyTasksChildTasks) SetInstanceId(v string) *UpdateInstanceImageResponseBodyTasksChildTasks {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceImageResponseBodyTasksChildTasks) SetTaskId(v string) *UpdateInstanceImageResponseBodyTasksChildTasks {
	s.TaskId = &v
	return s
}

func (s *UpdateInstanceImageResponseBodyTasksChildTasks) Validate() error {
	return dara.Validate(s)
}
