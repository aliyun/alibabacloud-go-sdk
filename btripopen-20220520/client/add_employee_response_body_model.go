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
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
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
