// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAndroidInstancesInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ResetAndroidInstancesInGroupResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ResetAndroidInstancesInGroupResponseBody
	GetRequestId() *string
	SetTasks(v *ResetAndroidInstancesInGroupResponseBodyTasks) *ResetAndroidInstancesInGroupResponseBody
	GetTasks() *ResetAndroidInstancesInGroupResponseBodyTasks
}

type ResetAndroidInstancesInGroupResponseBody struct {
	// The order ID for storage reduction. This parameter is returned only when you reduce storage while resetting instances in a cloud phone matrix. It is not returned in other scenarios.
	//
	// example:
	//
	// 22365781890****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Tasks *ResetAndroidInstancesInGroupResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s ResetAndroidInstancesInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAndroidInstancesInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ResetAndroidInstancesInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAndroidInstancesInGroupResponseBody) GetTasks() *ResetAndroidInstancesInGroupResponseBodyTasks {
	return s.Tasks
}

func (s *ResetAndroidInstancesInGroupResponseBody) SetOrderId(v string) *ResetAndroidInstancesInGroupResponseBody {
	s.OrderId = &v
	return s
}

func (s *ResetAndroidInstancesInGroupResponseBody) SetRequestId(v string) *ResetAndroidInstancesInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAndroidInstancesInGroupResponseBody) SetTasks(v *ResetAndroidInstancesInGroupResponseBodyTasks) *ResetAndroidInstancesInGroupResponseBody {
	s.Tasks = v
	return s
}

func (s *ResetAndroidInstancesInGroupResponseBody) Validate() error {
	if s.Tasks != nil {
		if err := s.Tasks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetAndroidInstancesInGroupResponseBodyTasks struct {
	// The child tasks.
	ChildTasks []*ResetAndroidInstancesInGroupResponseBodyTasksChildTasks `json:"ChildTasks,omitempty" xml:"ChildTasks,omitempty" type:"Repeated"`
	// The parent task ID.
	//
	// example:
	//
	// t-xxxx
	ParentTaskId *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
}

func (s ResetAndroidInstancesInGroupResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ResetAndroidInstancesInGroupResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasks) GetChildTasks() []*ResetAndroidInstancesInGroupResponseBodyTasksChildTasks {
	return s.ChildTasks
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasks) GetParentTaskId() *string {
	return s.ParentTaskId
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasks) SetChildTasks(v []*ResetAndroidInstancesInGroupResponseBodyTasksChildTasks) *ResetAndroidInstancesInGroupResponseBodyTasks {
	s.ChildTasks = v
	return s
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasks) SetParentTaskId(v string) *ResetAndroidInstancesInGroupResponseBodyTasks {
	s.ParentTaskId = &v
	return s
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasks) Validate() error {
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

type ResetAndroidInstancesInGroupResponseBodyTasksChildTasks struct {
	// The instance ID.
	//
	// example:
	//
	// acp-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The child task ID.
	//
	// example:
	//
	// t-xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ResetAndroidInstancesInGroupResponseBodyTasksChildTasks) String() string {
	return dara.Prettify(s)
}

func (s ResetAndroidInstancesInGroupResponseBodyTasksChildTasks) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasksChildTasks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasksChildTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasksChildTasks) SetInstanceId(v string) *ResetAndroidInstancesInGroupResponseBodyTasksChildTasks {
	s.InstanceId = &v
	return s
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasksChildTasks) SetTaskId(v string) *ResetAndroidInstancesInGroupResponseBodyTasksChildTasks {
	s.TaskId = &v
	return s
}

func (s *ResetAndroidInstancesInGroupResponseBodyTasksChildTasks) Validate() error {
	return dara.Validate(s)
}
