// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RestartInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *RestartInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RestartInstanceResponseBody
	GetTaskId() *string
}

type RestartInstanceResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EFC9161F-15E3-4A6E-8A99-C33331****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 11111****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RestartInstanceResponseBody) SetInstanceId(v string) *RestartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetTaskId(v string) *RestartInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *RestartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
