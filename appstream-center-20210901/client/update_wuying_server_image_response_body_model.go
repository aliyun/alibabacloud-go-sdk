// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWuyingServerImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWuyingServerImageResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateWuyingServerImageResponseBody
	GetTaskId() *string
}

type UpdateWuyingServerImageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the update task.
	//
	// example:
	//
	// ota-be7jzm29wrrz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateWuyingServerImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWuyingServerImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWuyingServerImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWuyingServerImageResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateWuyingServerImageResponseBody) SetRequestId(v string) *UpdateWuyingServerImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWuyingServerImageResponseBody) SetTaskId(v string) *UpdateWuyingServerImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateWuyingServerImageResponseBody) Validate() error {
	return dara.Validate(s)
}
