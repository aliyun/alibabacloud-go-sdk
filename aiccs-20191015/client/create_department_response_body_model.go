// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDepartmentResponseBody
	GetCode() *string
	SetData(v int64) *CreateDepartmentResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateDepartmentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDepartmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDepartmentResponseBody
	GetSuccess() *bool
}

type CreateDepartmentResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Department ID.
	//
	// example:
	//
	// 123456
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 98B032F5-6473-4EAC-8BA8-C28993513A1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API invocation succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDepartmentResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDepartmentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDepartmentResponseBody) SetCode(v string) *CreateDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetData(v int64) *CreateDepartmentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetHttpStatusCode(v int32) *CreateDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetMessage(v string) *CreateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetRequestId(v string) *CreateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetSuccess(v bool) *CreateDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
