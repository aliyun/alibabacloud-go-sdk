// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDeliveryTaskStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserDeliveryTaskStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateUserDeliveryTaskStatusResponseBody
	GetStatus() *string
	SetTaskName(v string) *UpdateUserDeliveryTaskStatusResponseBody
	GetTaskName() *string
}

type UpdateUserDeliveryTaskStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the delivery task.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the delivery task.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateUserDeliveryTaskStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDeliveryTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) SetRequestId(v string) *UpdateUserDeliveryTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) SetStatus(v string) *UpdateUserDeliveryTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) SetTaskName(v string) *UpdateUserDeliveryTaskStatusResponseBody {
	s.TaskName = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
