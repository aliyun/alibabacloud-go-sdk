// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddEmployeeResponseBody
	GetCode() *string
	SetMessage(v string) *AddEmployeeResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddEmployeeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddEmployeeResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AddEmployeeResponseBody
	GetTraceId() *string
}

type AddEmployeeResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s AddEmployeeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeeResponseBody) GoString() string {
	return s.String()
}

func (s *AddEmployeeResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddEmployeeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddEmployeeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEmployeeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddEmployeeResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AddEmployeeResponseBody) SetCode(v string) *AddEmployeeResponseBody {
	s.Code = &v
	return s
}

func (s *AddEmployeeResponseBody) SetMessage(v string) *AddEmployeeResponseBody {
	s.Message = &v
	return s
}

func (s *AddEmployeeResponseBody) SetRequestId(v string) *AddEmployeeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEmployeeResponseBody) SetSuccess(v bool) *AddEmployeeResponseBody {
	s.Success = &v
	return s
}

func (s *AddEmployeeResponseBody) SetTraceId(v string) *AddEmployeeResponseBody {
	s.TraceId = &v
	return s
}

func (s *AddEmployeeResponseBody) Validate() error {
	return dara.Validate(s)
}
