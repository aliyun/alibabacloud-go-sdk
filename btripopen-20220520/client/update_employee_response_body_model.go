// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEmployeeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateEmployeeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateEmployeeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEmployeeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEmployeeResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateEmployeeResponseBody
	GetTraceId() *string
}

type UpdateEmployeeResponseBody struct {
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
	// A5009956-1077-52FB-B520-EA8C7E91D722
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

func (s UpdateEmployeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEmployeeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateEmployeeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEmployeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEmployeeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEmployeeResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateEmployeeResponseBody) SetCode(v string) *UpdateEmployeeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEmployeeResponseBody) SetHttpStatusCode(v int32) *UpdateEmployeeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateEmployeeResponseBody) SetMessage(v string) *UpdateEmployeeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEmployeeResponseBody) SetRequestId(v string) *UpdateEmployeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEmployeeResponseBody) SetSuccess(v bool) *UpdateEmployeeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEmployeeResponseBody) SetTraceId(v string) *UpdateEmployeeResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateEmployeeResponseBody) Validate() error {
	return dara.Validate(s)
}
