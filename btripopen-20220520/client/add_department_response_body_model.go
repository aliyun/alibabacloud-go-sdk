// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDepartmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddDepartmentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddDepartmentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddDepartmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddDepartmentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddDepartmentResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AddDepartmentResponseBody
	GetTraceId() *string
}

type AddDepartmentResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId        *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AddDepartmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *AddDepartmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddDepartmentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddDepartmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddDepartmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDepartmentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddDepartmentResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AddDepartmentResponseBody) SetCode(v string) *AddDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *AddDepartmentResponseBody) SetHttpStatusCode(v int32) *AddDepartmentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddDepartmentResponseBody) SetMessage(v string) *AddDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *AddDepartmentResponseBody) SetRequestId(v string) *AddDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDepartmentResponseBody) SetSuccess(v bool) *AddDepartmentResponseBody {
	s.Success = &v
	return s
}

func (s *AddDepartmentResponseBody) SetTraceId(v string) *AddDepartmentResponseBody {
	s.TraceId = &v
	return s
}

func (s *AddDepartmentResponseBody) Validate() error {
	return dara.Validate(s)
}
