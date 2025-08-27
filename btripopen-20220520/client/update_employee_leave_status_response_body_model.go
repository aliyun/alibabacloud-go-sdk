// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeLeaveStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEmployeeLeaveStatusResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateEmployeeLeaveStatusResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateEmployeeLeaveStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEmployeeLeaveStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEmployeeLeaveStatusResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateEmployeeLeaveStatusResponseBody
	GetTraceId() *string
}

type UpdateEmployeeLeaveStatusResponseBody struct {
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
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s UpdateEmployeeLeaveStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeLeaveStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeLeaveStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEmployeeLeaveStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateEmployeeLeaveStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEmployeeLeaveStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEmployeeLeaveStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEmployeeLeaveStatusResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateEmployeeLeaveStatusResponseBody) SetCode(v string) *UpdateEmployeeLeaveStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponseBody) SetHttpStatusCode(v int32) *UpdateEmployeeLeaveStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponseBody) SetMessage(v string) *UpdateEmployeeLeaveStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponseBody) SetRequestId(v string) *UpdateEmployeeLeaveStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponseBody) SetSuccess(v bool) *UpdateEmployeeLeaveStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponseBody) SetTraceId(v string) *UpdateEmployeeLeaveStatusResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
