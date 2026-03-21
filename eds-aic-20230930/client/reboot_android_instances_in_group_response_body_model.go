// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootAndroidInstancesInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootAndroidInstancesInGroupResponseBody
	GetRequestId() *string
	SetTasks(v *RebootAndroidInstancesInGroupResponseBodyTasks) *RebootAndroidInstancesInGroupResponseBody
	GetTasks() *RebootAndroidInstancesInGroupResponseBodyTasks
}

type RebootAndroidInstancesInGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 227CBB4C-F5DC-589D-A667-C5CA3D52****
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     *RebootAndroidInstancesInGroupResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s RebootAndroidInstancesInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootAndroidInstancesInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootAndroidInstancesInGroupResponseBody) GetTasks() *RebootAndroidInstancesInGroupResponseBodyTasks {
	return s.Tasks
}

func (s *RebootAndroidInstancesInGroupResponseBody) SetRequestId(v string) *RebootAndroidInstancesInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootAndroidInstancesInGroupResponseBody) SetTasks(v *RebootAndroidInstancesInGroupResponseBodyTasks) *RebootAndroidInstancesInGroupResponseBody {
	s.Tasks = v
	return s
}

func (s *RebootAndroidInstancesInGroupResponseBody) Validate() error {
	if s.Tasks != nil {
		if err := s.Tasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RebootAndroidInstancesInGroupResponseBodyTasks struct {
	ChildTasks []*RebootAndroidInstancesInGroupResponseBodyTasksChildTasks `json:"ChildTasks,omitempty" xml:"ChildTasks,omitempty" type:"Repeated"`
	// example:
	//
	// t-xxxx
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
}

func (s RebootAndroidInstancesInGroupResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s RebootAndroidInstancesInGroupResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasks) GetChildTasks() []*RebootAndroidInstancesInGroupResponseBodyTasksChildTasks {
	return s.ChildTasks
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasks) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasks) SetChildTasks(v []*RebootAndroidInstancesInGroupResponseBodyTasksChildTasks) *RebootAndroidInstancesInGroupResponseBodyTasks {
	s.ChildTasks = v
	return s
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasks) SetParentTaskId(v string) *RebootAndroidInstancesInGroupResponseBodyTasks {
	s.ParentTaskId = &v
	return s
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasks) Validate() error {
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

type RebootAndroidInstancesInGroupResponseBodyTasksChildTasks struct {
	// example:
	//
	// acp-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// t-xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RebootAndroidInstancesInGroupResponseBodyTasksChildTasks) String() string {
	return dara.Prettify(s)
}

func (s RebootAndroidInstancesInGroupResponseBodyTasksChildTasks) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasksChildTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasksChildTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasksChildTasks) SetInstanceId(v string) *RebootAndroidInstancesInGroupResponseBodyTasksChildTasks {
	s.InstanceId = &v
	return s
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasksChildTasks) SetTaskId(v string) *RebootAndroidInstancesInGroupResponseBodyTasksChildTasks {
	s.TaskId = &v
	return s
}

func (s *RebootAndroidInstancesInGroupResponseBodyTasksChildTasks) Validate() error {
	return dara.Validate(s)
}
