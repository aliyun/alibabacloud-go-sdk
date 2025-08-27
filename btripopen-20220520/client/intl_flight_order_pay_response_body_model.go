// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOrderPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightOrderPayResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightOrderPayResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightOrderPayResponseBodyModule) *IntlFlightOrderPayResponseBody
	GetModule() *IntlFlightOrderPayResponseBodyModule
	SetRequestId(v string) *IntlFlightOrderPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightOrderPayResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightOrderPayResponseBody
	GetTraceId() *string
}

type IntlFlightOrderPayResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message   *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Module    *IntlFlightOrderPayResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightOrderPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightOrderPayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightOrderPayResponseBody) GetModule() *IntlFlightOrderPayResponseBodyModule {
	return s.Module
}

func (s *IntlFlightOrderPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightOrderPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightOrderPayResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightOrderPayResponseBody) SetCode(v string) *IntlFlightOrderPayResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightOrderPayResponseBody) SetMessage(v string) *IntlFlightOrderPayResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightOrderPayResponseBody) SetModule(v *IntlFlightOrderPayResponseBodyModule) *IntlFlightOrderPayResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightOrderPayResponseBody) SetRequestId(v string) *IntlFlightOrderPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightOrderPayResponseBody) SetSuccess(v bool) *IntlFlightOrderPayResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightOrderPayResponseBody) SetTraceId(v string) *IntlFlightOrderPayResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightOrderPayResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOrderPayResponseBodyModule struct {
	ActualPayPrice *int64 `json:"actual_pay_price,omitempty" xml:"actual_pay_price,omitempty"`
	PayStatus      *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}

func (s IntlFlightOrderPayResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOrderPayResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightOrderPayResponseBodyModule) GetActualPayPrice() *int64 {
	return s.ActualPayPrice
}

func (s *IntlFlightOrderPayResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *IntlFlightOrderPayResponseBodyModule) SetActualPayPrice(v int64) *IntlFlightOrderPayResponseBodyModule {
	s.ActualPayPrice = &v
	return s
}

func (s *IntlFlightOrderPayResponseBodyModule) SetPayStatus(v int32) *IntlFlightOrderPayResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *IntlFlightOrderPayResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
