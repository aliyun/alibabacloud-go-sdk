// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*AttachInstancesResponseBodyList) *AttachInstancesResponseBody
	GetList() []*AttachInstancesResponseBodyList
	SetTaskId(v string) *AttachInstancesResponseBody
	GetTaskId() *string
}

type AttachInstancesResponseBody struct {
	// The list of node addition information.
	List []*AttachInstancesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// T-5a544aff80282e39ea00****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s AttachInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBody) GetList() []*AttachInstancesResponseBodyList {
	return s.List
}

func (s *AttachInstancesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *AttachInstancesResponseBody) SetList(v []*AttachInstancesResponseBodyList) *AttachInstancesResponseBody {
	s.List = v
	return s
}

func (s *AttachInstancesResponseBody) SetTaskId(v string) *AttachInstancesResponseBody {
	s.TaskId = &v
	return s
}

func (s *AttachInstancesResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachInstancesResponseBodyList struct {
	// The status code of the node addition result.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ECS instance ID.
	//
	// example:
	//
	// i-2ze0lgm3y6iylcbt****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The description of the node addition result.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s AttachInstancesResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesResponseBodyList) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBodyList) GetCode() *string {
	return s.Code
}

func (s *AttachInstancesResponseBodyList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachInstancesResponseBodyList) GetMessage() *string {
	return s.Message
}

func (s *AttachInstancesResponseBodyList) SetCode(v string) *AttachInstancesResponseBodyList {
	s.Code = &v
	return s
}

func (s *AttachInstancesResponseBodyList) SetInstanceId(v string) *AttachInstancesResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *AttachInstancesResponseBodyList) SetMessage(v string) *AttachInstancesResponseBodyList {
	s.Message = &v
	return s
}

func (s *AttachInstancesResponseBodyList) Validate() error {
	return dara.Validate(s)
}
