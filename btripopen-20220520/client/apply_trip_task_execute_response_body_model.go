// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTripTaskExecuteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ApplyTripTaskExecuteResponseBody
	GetCode() *string
	SetMessage(v string) *ApplyTripTaskExecuteResponseBody
	GetMessage() *string
	SetModule(v bool) *ApplyTripTaskExecuteResponseBody
	GetModule() *bool
	SetRequestId(v string) *ApplyTripTaskExecuteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyTripTaskExecuteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ApplyTripTaskExecuteResponseBody
	GetTraceId() *string
}

type ApplyTripTaskExecuteResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Module *bool `json:"module,omitempty" xml:"module,omitempty"`
	// example:
	//
	// 2FB0D7A8-BA41-5D04-BEFC-CADA5481AC53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210bc56016876728084104176d2c35
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ApplyTripTaskExecuteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyTripTaskExecuteResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyTripTaskExecuteResponseBody) GetCode() *string {
	return s.Code
}

func (s *ApplyTripTaskExecuteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ApplyTripTaskExecuteResponseBody) GetModule() *bool {
	return s.Module
}

func (s *ApplyTripTaskExecuteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyTripTaskExecuteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyTripTaskExecuteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ApplyTripTaskExecuteResponseBody) SetCode(v string) *ApplyTripTaskExecuteResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyTripTaskExecuteResponseBody) SetMessage(v string) *ApplyTripTaskExecuteResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyTripTaskExecuteResponseBody) SetModule(v bool) *ApplyTripTaskExecuteResponseBody {
	s.Module = &v
	return s
}

func (s *ApplyTripTaskExecuteResponseBody) SetRequestId(v string) *ApplyTripTaskExecuteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyTripTaskExecuteResponseBody) SetSuccess(v bool) *ApplyTripTaskExecuteResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyTripTaskExecuteResponseBody) SetTraceId(v string) *ApplyTripTaskExecuteResponseBody {
	s.TraceId = &v
	return s
}

func (s *ApplyTripTaskExecuteResponseBody) Validate() error {
	return dara.Validate(s)
}
