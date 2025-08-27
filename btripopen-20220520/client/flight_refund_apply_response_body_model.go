// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightRefundApplyResponseBody
	GetCode() *string
	SetMessage(v string) *FlightRefundApplyResponseBody
	GetMessage() *string
	SetModule(v *FlightRefundApplyResponseBodyModule) *FlightRefundApplyResponseBody
	GetModule() *FlightRefundApplyResponseBodyModule
	SetRequestId(v string) *FlightRefundApplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightRefundApplyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightRefundApplyResponseBody
	GetTraceId() *string
}

type FlightRefundApplyResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightRefundApplyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightRefundApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyResponseBody) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightRefundApplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightRefundApplyResponseBody) GetModule() *FlightRefundApplyResponseBodyModule {
	return s.Module
}

func (s *FlightRefundApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightRefundApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightRefundApplyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightRefundApplyResponseBody) SetCode(v string) *FlightRefundApplyResponseBody {
	s.Code = &v
	return s
}

func (s *FlightRefundApplyResponseBody) SetMessage(v string) *FlightRefundApplyResponseBody {
	s.Message = &v
	return s
}

func (s *FlightRefundApplyResponseBody) SetModule(v *FlightRefundApplyResponseBodyModule) *FlightRefundApplyResponseBody {
	s.Module = v
	return s
}

func (s *FlightRefundApplyResponseBody) SetRequestId(v string) *FlightRefundApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightRefundApplyResponseBody) SetSuccess(v bool) *FlightRefundApplyResponseBody {
	s.Success = &v
	return s
}

func (s *FlightRefundApplyResponseBody) SetTraceId(v string) *FlightRefundApplyResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightRefundApplyResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightRefundApplyResponseBodyModule struct {
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// refun1234
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// example:
	//
	// 1231231
	RefundApplyId *int64 `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
	// example:
	//
	// 1000
	RefundFee *int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// example:
	//
	// 100
	RefundMoney *int64 `json:"refund_money,omitempty" xml:"refund_money,omitempty"`
}

func (s FlightRefundApplyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyResponseBodyModule) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightRefundApplyResponseBodyModule) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *FlightRefundApplyResponseBodyModule) GetRefundApplyId() *int64 {
	return s.RefundApplyId
}

func (s *FlightRefundApplyResponseBodyModule) GetRefundFee() *int64 {
	return s.RefundFee
}

func (s *FlightRefundApplyResponseBodyModule) GetRefundMoney() *int64 {
	return s.RefundMoney
}

func (s *FlightRefundApplyResponseBodyModule) SetDisOrderId(v string) *FlightRefundApplyResponseBodyModule {
	s.DisOrderId = &v
	return s
}

func (s *FlightRefundApplyResponseBodyModule) SetDisSubOrderId(v string) *FlightRefundApplyResponseBodyModule {
	s.DisSubOrderId = &v
	return s
}

func (s *FlightRefundApplyResponseBodyModule) SetRefundApplyId(v int64) *FlightRefundApplyResponseBodyModule {
	s.RefundApplyId = &v
	return s
}

func (s *FlightRefundApplyResponseBodyModule) SetRefundFee(v int64) *FlightRefundApplyResponseBodyModule {
	s.RefundFee = &v
	return s
}

func (s *FlightRefundApplyResponseBodyModule) SetRefundMoney(v int64) *FlightRefundApplyResponseBodyModule {
	s.RefundMoney = &v
	return s
}

func (s *FlightRefundApplyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
