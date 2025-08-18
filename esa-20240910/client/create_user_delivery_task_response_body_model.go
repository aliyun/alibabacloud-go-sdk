// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserDeliveryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCenter(v string) *CreateUserDeliveryTaskResponseBody
	GetDataCenter() *string
	SetRequestId(v string) *CreateUserDeliveryTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateUserDeliveryTaskResponseBody
	GetStatus() *string
	SetTaskName(v string) *CreateUserDeliveryTaskResponseBody
	GetTaskName() *string
}

type CreateUserDeliveryTaskResponseBody struct {
	// The data center. Valid values:
	//
	// 	- cn: the Chinese mainland.
	//
	// 	- sg: outside the Chinese mainland.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2CCD40B1-3F20-5FF0-8A67-E3F34B87744F
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
	// er-http
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateUserDeliveryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserDeliveryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserDeliveryTaskResponseBody) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CreateUserDeliveryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserDeliveryTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateUserDeliveryTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateUserDeliveryTaskResponseBody) SetDataCenter(v string) *CreateUserDeliveryTaskResponseBody {
	s.DataCenter = &v
	return s
}

func (s *CreateUserDeliveryTaskResponseBody) SetRequestId(v string) *CreateUserDeliveryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserDeliveryTaskResponseBody) SetStatus(v string) *CreateUserDeliveryTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateUserDeliveryTaskResponseBody) SetTaskName(v string) *CreateUserDeliveryTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *CreateUserDeliveryTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
