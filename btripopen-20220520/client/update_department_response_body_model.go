// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateDepartmentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateDepartmentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateDepartmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDepartmentResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateDepartmentResponseBody
	GetTraceId() *string
}

type UpdateDepartmentResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s UpdateDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateDepartmentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDepartmentResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateDepartmentResponseBody) SetCode(v string) *UpdateDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetHttpStatusCode(v int32) *UpdateDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetMessage(v string) *UpdateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetRequestId(v string) *UpdateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetSuccess(v bool) *UpdateDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetTraceId(v string) *UpdateDepartmentResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
