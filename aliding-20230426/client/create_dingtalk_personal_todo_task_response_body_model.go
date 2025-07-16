// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDingtalkPersonalTodoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v int64) *CreateDingtalkPersonalTodoTaskResponseBody
	GetCreatedTime() *int64
	SetRequestId(v string) *CreateDingtalkPersonalTodoTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateDingtalkPersonalTodoTaskResponseBody
	GetTaskId() *string
	SetVendorRequestId(v string) *CreateDingtalkPersonalTodoTaskResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateDingtalkPersonalTodoTaskResponseBody
	GetVendorType() *string
}

type CreateDingtalkPersonalTodoTaskResponseBody struct {
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

func (s CreateDingtalkPersonalTodoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) SetCreatedTime(v int64) *CreateDingtalkPersonalTodoTaskResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) SetRequestId(v string) *CreateDingtalkPersonalTodoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) SetTaskId(v string) *CreateDingtalkPersonalTodoTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) SetVendorRequestId(v string) *CreateDingtalkPersonalTodoTaskResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) SetVendorType(v string) *CreateDingtalkPersonalTodoTaskResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
