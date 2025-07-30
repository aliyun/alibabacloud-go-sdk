// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateInstancePublicConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *AllocateInstancePublicConnectionResponseBody
	GetInstanceName() *string
	SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *AllocateInstancePublicConnectionResponseBody
	GetTaskId() *int64
}

type AllocateInstancePublicConnectionResponseBody struct {
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5ED62C81-9948-5612-81E1-EA3853752306
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 498115273
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AllocateInstancePublicConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *AllocateInstancePublicConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateInstancePublicConnectionResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *AllocateInstancePublicConnectionResponseBody) SetInstanceName(v string) *AllocateInstancePublicConnectionResponseBody {
	s.InstanceName = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetRequestId(v string) *AllocateInstancePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) SetTaskId(v int64) *AllocateInstancePublicConnectionResponseBody {
	s.TaskId = &v
	return s
}

func (s *AllocateInstancePublicConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
