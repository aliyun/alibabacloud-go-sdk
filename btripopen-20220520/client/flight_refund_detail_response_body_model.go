// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightRefundDetailResponseBody
	GetCode() *string
	SetMessage(v string) *FlightRefundDetailResponseBody
	GetMessage() *string
	SetModule(v *FlightRefundDetailResponseBodyModule) *FlightRefundDetailResponseBody
	GetModule() *FlightRefundDetailResponseBodyModule
	SetRequestId(v string) *FlightRefundDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightRefundDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightRefundDetailResponseBody
	GetTraceId() *string
}

type FlightRefundDetailResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightRefundDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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

func (s FlightRefundDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailResponseBody) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightRefundDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightRefundDetailResponseBody) GetModule() *FlightRefundDetailResponseBodyModule {
	return s.Module
}

func (s *FlightRefundDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightRefundDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightRefundDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightRefundDetailResponseBody) SetCode(v string) *FlightRefundDetailResponseBody {
	s.Code = &v
	return s
}

func (s *FlightRefundDetailResponseBody) SetMessage(v string) *FlightRefundDetailResponseBody {
	s.Message = &v
	return s
}

func (s *FlightRefundDetailResponseBody) SetModule(v *FlightRefundDetailResponseBodyModule) *FlightRefundDetailResponseBody {
	s.Module = v
	return s
}

func (s *FlightRefundDetailResponseBody) SetRequestId(v string) *FlightRefundDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightRefundDetailResponseBody) SetSuccess(v bool) *FlightRefundDetailResponseBody {
	s.Success = &v
	return s
}

func (s *FlightRefundDetailResponseBody) SetTraceId(v string) *FlightRefundDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightRefundDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightRefundDetailResponseBodyModule struct {
	// example:
	//
	// 123
	BtripOrderId *int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// example:
	//
	// 123
	BtripSubOrderId *int64 `json:"btrip_sub_order_id,omitempty" xml:"btrip_sub_order_id,omitempty"`
	// example:
	//
	// dis1234
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// refun123
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// example:
	//
	// 1
	IsVoluntary *int32  `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	Reason      *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// 100
	RefundFee     *int64                                               `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	RefundFeeList []*FlightRefundDetailResponseBodyModuleRefundFeeList `json:"refund_fee_list,omitempty" xml:"refund_fee_list,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	RefundPrice *int64 `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
	// example:
	//
	// 0
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s FlightRefundDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailResponseBodyModule) GetBtripOrderId() *int64 {
	return s.BtripOrderId
}

func (s *FlightRefundDetailResponseBodyModule) GetBtripSubOrderId() *int64 {
	return s.BtripSubOrderId
}

func (s *FlightRefundDetailResponseBodyModule) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightRefundDetailResponseBodyModule) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *FlightRefundDetailResponseBodyModule) GetIsVoluntary() *int32 {
	return s.IsVoluntary
}

func (s *FlightRefundDetailResponseBodyModule) GetReason() *string {
	return s.Reason
}

func (s *FlightRefundDetailResponseBodyModule) GetRefundFee() *int64 {
	return s.RefundFee
}

func (s *FlightRefundDetailResponseBodyModule) GetRefundFeeList() []*FlightRefundDetailResponseBodyModuleRefundFeeList {
	return s.RefundFeeList
}

func (s *FlightRefundDetailResponseBodyModule) GetRefundPrice() *int64 {
	return s.RefundPrice
}

func (s *FlightRefundDetailResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *FlightRefundDetailResponseBodyModule) SetBtripOrderId(v int64) *FlightRefundDetailResponseBodyModule {
	s.BtripOrderId = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetBtripSubOrderId(v int64) *FlightRefundDetailResponseBodyModule {
	s.BtripSubOrderId = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetDisOrderId(v string) *FlightRefundDetailResponseBodyModule {
	s.DisOrderId = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetDisSubOrderId(v string) *FlightRefundDetailResponseBodyModule {
	s.DisSubOrderId = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetIsVoluntary(v int32) *FlightRefundDetailResponseBodyModule {
	s.IsVoluntary = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetReason(v string) *FlightRefundDetailResponseBodyModule {
	s.Reason = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetRefundFee(v int64) *FlightRefundDetailResponseBodyModule {
	s.RefundFee = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetRefundFeeList(v []*FlightRefundDetailResponseBodyModuleRefundFeeList) *FlightRefundDetailResponseBodyModule {
	s.RefundFeeList = v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetRefundPrice(v int64) *FlightRefundDetailResponseBodyModule {
	s.RefundPrice = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) SetStatus(v string) *FlightRefundDetailResponseBodyModule {
	s.Status = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightRefundDetailResponseBodyModuleRefundFeeList struct {
	// example:
	//
	// 293982882881999
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// example:
	//
	// 100
	RefundFee *int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// example:
	//
	// 100
	RefundPrice *int64  `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s FlightRefundDetailResponseBodyModuleRefundFeeList) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundDetailResponseBodyModuleRefundFeeList) GoString() string {
	return s.String()
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) GetRefundFee() *int64 {
	return s.RefundFee
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) GetRefundPrice() *int64 {
	return s.RefundPrice
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) GetStatus() *string {
	return s.Status
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) SetAlipayTradeNo(v string) *FlightRefundDetailResponseBodyModuleRefundFeeList {
	s.AlipayTradeNo = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) SetRefundFee(v int64) *FlightRefundDetailResponseBodyModuleRefundFeeList {
	s.RefundFee = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) SetRefundPrice(v int64) *FlightRefundDetailResponseBodyModuleRefundFeeList {
	s.RefundPrice = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) SetStatus(v string) *FlightRefundDetailResponseBodyModuleRefundFeeList {
	s.Status = &v
	return s
}

func (s *FlightRefundDetailResponseBodyModuleRefundFeeList) Validate() error {
	return dara.Validate(s)
}
