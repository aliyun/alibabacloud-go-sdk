// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightCreateOrderResponseBody
	GetCode() *string
	SetMessage(v string) *FlightCreateOrderResponseBody
	GetMessage() *string
	SetModule(v *FlightCreateOrderResponseBodyModule) *FlightCreateOrderResponseBody
	GetModule() *FlightCreateOrderResponseBodyModule
	SetRequestId(v string) *FlightCreateOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightCreateOrderResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightCreateOrderResponseBody
	GetTraceId() *string
}

type FlightCreateOrderResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightCreateOrderResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightCreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightCreateOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightCreateOrderResponseBody) GetModule() *FlightCreateOrderResponseBodyModule {
	return s.Module
}

func (s *FlightCreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightCreateOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightCreateOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightCreateOrderResponseBody) SetCode(v string) *FlightCreateOrderResponseBody {
	s.Code = &v
	return s
}

func (s *FlightCreateOrderResponseBody) SetMessage(v string) *FlightCreateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *FlightCreateOrderResponseBody) SetModule(v *FlightCreateOrderResponseBodyModule) *FlightCreateOrderResponseBody {
	s.Module = v
	return s
}

func (s *FlightCreateOrderResponseBody) SetRequestId(v string) *FlightCreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightCreateOrderResponseBody) SetSuccess(v bool) *FlightCreateOrderResponseBody {
	s.Success = &v
	return s
}

func (s *FlightCreateOrderResponseBody) SetTraceId(v string) *FlightCreateOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightCreateOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightCreateOrderResponseBodyModule struct {
	// example:
	//
	// 3287177727711
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// example:
	//
	// 123
	BtripOrderId *int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	LastPayTime *string `json:"last_pay_time,omitempty" xml:"last_pay_time,omitempty"`
	// example:
	//
	// 1
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// example:
	//
	// 1
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 100
	PaymentPrice *int64 `json:"payment_price,omitempty" xml:"payment_price,omitempty"`
	// example:
	//
	// 100
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

func (s FlightCreateOrderResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderResponseBodyModule) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *FlightCreateOrderResponseBodyModule) GetBtripOrderId() *int64 {
	return s.BtripOrderId
}

func (s *FlightCreateOrderResponseBodyModule) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightCreateOrderResponseBodyModule) GetLastPayTime() *string {
	return s.LastPayTime
}

func (s *FlightCreateOrderResponseBodyModule) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *FlightCreateOrderResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *FlightCreateOrderResponseBodyModule) GetPaymentPrice() *int64 {
	return s.PaymentPrice
}

func (s *FlightCreateOrderResponseBodyModule) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *FlightCreateOrderResponseBodyModule) SetAlipayTradeNo(v string) *FlightCreateOrderResponseBodyModule {
	s.AlipayTradeNo = &v
	return s
}

func (s *FlightCreateOrderResponseBodyModule) SetBtripOrderId(v int64) *FlightCreateOrderResponseBodyModule {
	s.BtripOrderId = &v
	return s
}

func (s *FlightCreateOrderResponseBodyModule) SetDisOrderId(v string) *FlightCreateOrderResponseBodyModule {
	s.DisOrderId = &v
	return s
}

func (s *FlightCreateOrderResponseBodyModule) SetLastPayTime(v string) *FlightCreateOrderResponseBodyModule {
	s.LastPayTime = &v
	return s
}

func (s *FlightCreateOrderResponseBodyModule) SetOrderStatus(v int32) *FlightCreateOrderResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *FlightCreateOrderResponseBodyModule) SetPayStatus(v int32) *FlightCreateOrderResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *FlightCreateOrderResponseBodyModule) SetPaymentPrice(v int64) *FlightCreateOrderResponseBodyModule {
	s.PaymentPrice = &v
	return s
}

func (s *FlightCreateOrderResponseBodyModule) SetTotalPrice(v int64) *FlightCreateOrderResponseBodyModule {
	s.TotalPrice = &v
	return s
}

func (s *FlightCreateOrderResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
