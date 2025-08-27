// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightPayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightPayOrderResponseBody
	GetCode() *string
	SetMessage(v string) *FlightPayOrderResponseBody
	GetMessage() *string
	SetModule(v *FlightPayOrderResponseBodyModule) *FlightPayOrderResponseBody
	GetModule() *FlightPayOrderResponseBodyModule
	SetRequestId(v string) *FlightPayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightPayOrderResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightPayOrderResponseBody
	GetTraceId() *string
}

type FlightPayOrderResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightPayOrderResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210e842b16611337974412836dae27
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightPayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *FlightPayOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightPayOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightPayOrderResponseBody) GetModule() *FlightPayOrderResponseBodyModule {
	return s.Module
}

func (s *FlightPayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightPayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightPayOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightPayOrderResponseBody) SetCode(v string) *FlightPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *FlightPayOrderResponseBody) SetMessage(v string) *FlightPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *FlightPayOrderResponseBody) SetModule(v *FlightPayOrderResponseBodyModule) *FlightPayOrderResponseBody {
	s.Module = v
	return s
}

func (s *FlightPayOrderResponseBody) SetRequestId(v string) *FlightPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightPayOrderResponseBody) SetSuccess(v bool) *FlightPayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *FlightPayOrderResponseBody) SetTraceId(v string) *FlightPayOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightPayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightPayOrderResponseBodyModule struct {
	// example:
	//
	// 1000
	ActualPayPrice *int64 `json:"actual_pay_price,omitempty" xml:"actual_pay_price,omitempty"`
	// example:
	//
	// 12989127316726531726
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// example:
	//
	// 0000-00-00 00:00:00
	LastPayTime *string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// example:
	//
	// 0
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}

func (s FlightPayOrderResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightPayOrderResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightPayOrderResponseBodyModule) GetActualPayPrice() *int64 {
	return s.ActualPayPrice
}

func (s *FlightPayOrderResponseBodyModule) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *FlightPayOrderResponseBodyModule) GetLastPayTime() *string {
	return s.LastPayTime
}

func (s *FlightPayOrderResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *FlightPayOrderResponseBodyModule) SetActualPayPrice(v int64) *FlightPayOrderResponseBodyModule {
	s.ActualPayPrice = &v
	return s
}

func (s *FlightPayOrderResponseBodyModule) SetAlipayTradeNo(v string) *FlightPayOrderResponseBodyModule {
	s.AlipayTradeNo = &v
	return s
}

func (s *FlightPayOrderResponseBodyModule) SetLastPayTime(v string) *FlightPayOrderResponseBodyModule {
	s.LastPayTime = &v
	return s
}

func (s *FlightPayOrderResponseBodyModule) SetPayStatus(v int32) *FlightPayOrderResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *FlightPayOrderResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
