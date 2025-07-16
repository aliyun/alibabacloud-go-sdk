// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTodoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v int64) *CreatePersonalTodoTaskResponseBody
	GetCreatedTime() *int64
	SetRequestId(v string) *CreatePersonalTodoTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreatePersonalTodoTaskResponseBody
	GetTaskId() *string
	SetVendorRequestId(v string) *CreatePersonalTodoTaskResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreatePersonalTodoTaskResponseBody
	GetVendorType() *string
}

type CreatePersonalTodoTaskResponseBody struct {
	// example:
	//
	// 1703750708595
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// task123abc
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CreatePersonalTodoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *CreatePersonalTodoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalTodoTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreatePersonalTodoTaskResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreatePersonalTodoTaskResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreatePersonalTodoTaskResponseBody) SetCreatedTime(v int64) *CreatePersonalTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreatePersonalTodoTaskResponseBody) SetRequestId(v string) *CreatePersonalTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalTodoTaskResponseBody) SetTaskId(v string) *CreatePersonalTodoTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreatePersonalTodoTaskResponseBody) SetVendorRequestId(v string) *CreatePersonalTodoTaskResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreatePersonalTodoTaskResponseBody) SetVendorType(v string) *CreatePersonalTodoTaskResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreatePersonalTodoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
