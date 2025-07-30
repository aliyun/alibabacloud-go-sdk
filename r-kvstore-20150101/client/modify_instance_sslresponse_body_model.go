// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceSSLResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *ModifyInstanceSSLResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyInstanceSSLResponseBody
	GetTaskId() *string
}

type ModifyInstanceSSLResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AD7E16AA-6B23-43BF-979C-07D957FB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 32184****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSSLResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceSSLResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyInstanceSSLResponseBody) SetInstanceId(v string) *ModifyInstanceSSLResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSSLResponseBody) SetRequestId(v string) *ModifyInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceSSLResponseBody) SetTaskId(v string) *ModifyInstanceSSLResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
