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
	// The details of the added nodes.
	List []*AttachInstancesResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// T-5a544aff80282e39ea000039
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
	return dara.Validate(s)
}

type AttachInstancesResponseBodyList struct {
	// The code that indicates the task result.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-2ze0lgm3y6iylcbt****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Indicates whether the ECS instance is successfully added to the ACK cluster.
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
