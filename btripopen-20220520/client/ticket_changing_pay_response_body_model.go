// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingPayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TicketChangingPayResponseBody
	GetCode() *string
	SetMessage(v string) *TicketChangingPayResponseBody
	GetMessage() *string
	SetModule(v *TicketChangingPayResponseBodyModule) *TicketChangingPayResponseBody
	GetModule() *TicketChangingPayResponseBodyModule
	SetRequestId(v string) *TicketChangingPayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketChangingPayResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TicketChangingPayResponseBody
	GetTraceId() *string
}

type TicketChangingPayResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TicketChangingPayResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TicketChangingPayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingPayResponseBody) GoString() string {
	return s.String()
}

func (s *TicketChangingPayResponseBody) GetCode() *string {
	return s.Code
}

func (s *TicketChangingPayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TicketChangingPayResponseBody) GetModule() *TicketChangingPayResponseBodyModule {
	return s.Module
}

func (s *TicketChangingPayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketChangingPayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketChangingPayResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TicketChangingPayResponseBody) SetCode(v string) *TicketChangingPayResponseBody {
	s.Code = &v
	return s
}

func (s *TicketChangingPayResponseBody) SetMessage(v string) *TicketChangingPayResponseBody {
	s.Message = &v
	return s
}

func (s *TicketChangingPayResponseBody) SetModule(v *TicketChangingPayResponseBodyModule) *TicketChangingPayResponseBody {
	s.Module = v
	return s
}

func (s *TicketChangingPayResponseBody) SetRequestId(v string) *TicketChangingPayResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketChangingPayResponseBody) SetSuccess(v bool) *TicketChangingPayResponseBody {
	s.Success = &v
	return s
}

func (s *TicketChangingPayResponseBody) SetTraceId(v string) *TicketChangingPayResponseBody {
	s.TraceId = &v
	return s
}

func (s *TicketChangingPayResponseBody) Validate() error {
	return dara.Validate(s)
}

type TicketChangingPayResponseBodyModule struct {
	// example:
	//
	// true
	CanRetry *bool `json:"can_retry,omitempty" xml:"can_retry,omitempty"`
	// example:
	//
	// 1000
	PayPrice  *int64 `json:"pay_price,omitempty" xml:"pay_price,omitempty"`
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 0000-00-00 00:00:00
	PayTime *string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// 23098276578908765
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
}

func (s TicketChangingPayResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingPayResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TicketChangingPayResponseBodyModule) GetCanRetry() *bool {
	return s.CanRetry
}

func (s *TicketChangingPayResponseBodyModule) GetPayPrice() *int64 {
	return s.PayPrice
}

func (s *TicketChangingPayResponseBodyModule) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *TicketChangingPayResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *TicketChangingPayResponseBodyModule) GetTradeNo() *string {
	return s.TradeNo
}

func (s *TicketChangingPayResponseBodyModule) SetCanRetry(v bool) *TicketChangingPayResponseBodyModule {
	s.CanRetry = &v
	return s
}

func (s *TicketChangingPayResponseBodyModule) SetPayPrice(v int64) *TicketChangingPayResponseBodyModule {
	s.PayPrice = &v
	return s
}

func (s *TicketChangingPayResponseBodyModule) SetPayStatus(v int32) *TicketChangingPayResponseBodyModule {
	s.PayStatus = &v
	return s
}

func (s *TicketChangingPayResponseBodyModule) SetPayTime(v string) *TicketChangingPayResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *TicketChangingPayResponseBodyModule) SetTradeNo(v string) *TicketChangingPayResponseBodyModule {
	s.TradeNo = &v
	return s
}

func (s *TicketChangingPayResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
