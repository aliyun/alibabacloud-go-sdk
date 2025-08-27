// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightOrderCancelResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightOrderCancelResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightOrderCancelResponseBodyModule) *IntlFlightOrderCancelResponseBody
	GetModule() *IntlFlightOrderCancelResponseBodyModule
	SetRequestId(v string) *IntlFlightOrderCancelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightOrderCancelResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightOrderCancelResponseBody
	GetTraceId() *string
}

type IntlFlightOrderCancelResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module    *IntlFlightOrderCancelResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightOrderCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderCancelResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderCancelResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightOrderCancelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightOrderCancelResponseBody) GetModule() *IntlFlightOrderCancelResponseBodyModule {
	return s.Module
}

func (s *IntlFlightOrderCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightOrderCancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightOrderCancelResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightOrderCancelResponseBody) SetCode(v string) *IntlFlightOrderCancelResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightOrderCancelResponseBody) SetMessage(v string) *IntlFlightOrderCancelResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightOrderCancelResponseBody) SetModule(v *IntlFlightOrderCancelResponseBodyModule) *IntlFlightOrderCancelResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightOrderCancelResponseBody) SetRequestId(v string) *IntlFlightOrderCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightOrderCancelResponseBody) SetSuccess(v bool) *IntlFlightOrderCancelResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightOrderCancelResponseBody) SetTraceId(v string) *IntlFlightOrderCancelResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightOrderCancelResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderCancelResponseBodyModule struct {
	// example:
	//
	// 2023-08-10 17:45:32
	CancelTime *string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
}

func (s IntlFlightOrderCancelResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderCancelResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderCancelResponseBodyModule) GetCancelTime() *string {
	return s.CancelTime
}

func (s *IntlFlightOrderCancelResponseBodyModule) SetCancelTime(v string) *IntlFlightOrderCancelResponseBodyModule {
	s.CancelTime = &v
	return s
}

func (s *IntlFlightOrderCancelResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
