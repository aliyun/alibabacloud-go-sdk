// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketChangingApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TicketChangingApplyResponseBody
	GetCode() *string
	SetMessage(v string) *TicketChangingApplyResponseBody
	GetMessage() *string
	SetModule(v *TicketChangingApplyResponseBodyModule) *TicketChangingApplyResponseBody
	GetModule() *TicketChangingApplyResponseBodyModule
	SetRequestId(v string) *TicketChangingApplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketChangingApplyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TicketChangingApplyResponseBody
	GetTraceId() *string
}

type TicketChangingApplyResponseBody struct {
	// example:
	//
	// 0
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TicketChangingApplyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s TicketChangingApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingApplyResponseBody) GoString() string {
	return s.String()
}

func (s *TicketChangingApplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *TicketChangingApplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TicketChangingApplyResponseBody) GetModule() *TicketChangingApplyResponseBodyModule {
	return s.Module
}

func (s *TicketChangingApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketChangingApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketChangingApplyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TicketChangingApplyResponseBody) SetCode(v string) *TicketChangingApplyResponseBody {
	s.Code = &v
	return s
}

func (s *TicketChangingApplyResponseBody) SetMessage(v string) *TicketChangingApplyResponseBody {
	s.Message = &v
	return s
}

func (s *TicketChangingApplyResponseBody) SetModule(v *TicketChangingApplyResponseBodyModule) *TicketChangingApplyResponseBody {
	s.Module = v
	return s
}

func (s *TicketChangingApplyResponseBody) SetRequestId(v string) *TicketChangingApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketChangingApplyResponseBody) SetSuccess(v bool) *TicketChangingApplyResponseBody {
	s.Success = &v
	return s
}

func (s *TicketChangingApplyResponseBody) SetTraceId(v string) *TicketChangingApplyResponseBody {
	s.TraceId = &v
	return s
}

func (s *TicketChangingApplyResponseBody) Validate() error {
	return dara.Validate(s)
}

type TicketChangingApplyResponseBodyModule struct {
	// example:
	//
	// 1000
	BookingChangedTotalFee *int32 `json:"booking_changed_total_fee,omitempty" xml:"booking_changed_total_fee,omitempty"`
	// example:
	//
	// 1000
	BookingOriginTotalFee *int32 `json:"booking_origin_total_fee,omitempty" xml:"booking_origin_total_fee,omitempty"`
	// example:
	//
	// true
	BookingPriceChanged *bool `json:"booking_price_changed,omitempty" xml:"booking_price_changed,omitempty"`
	// example:
	//
	// 1234
	BtripOrderId *int64 `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	// example:
	//
	// 1234
	BtripSubOrderId *int64 `json:"btrip_sub_order_id,omitempty" xml:"btrip_sub_order_id,omitempty"`
	// example:
	//
	// true
	CanPay *bool `json:"can_pay,omitempty" xml:"can_pay,omitempty"`
	// example:
	//
	// 1000
	ChangeFee *int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// example:
	//
	// 0000-00-00 00:00:00
	DeadlineTime *string `json:"deadline_time,omitempty" xml:"deadline_time,omitempty"`
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// mid112
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// example:
	//
	// 1
	MaxRetryTimes *int32 `json:"max_retry_times,omitempty" xml:"max_retry_times,omitempty"`
	// example:
	//
	// 1000
	NextRetryInterval *int64 `json:"next_retry_interval,omitempty" xml:"next_retry_interval,omitempty"`
	// example:
	//
	// true
	Retry           *bool   `json:"retry,omitempty" xml:"retry,omitempty"`
	RetryClientTips *string `json:"retry_client_tips,omitempty" xml:"retry_client_tips,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1000
	UpgradeFee *int64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

func (s TicketChangingApplyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TicketChangingApplyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TicketChangingApplyResponseBodyModule) GetBookingChangedTotalFee() *int32 {
	return s.BookingChangedTotalFee
}

func (s *TicketChangingApplyResponseBodyModule) GetBookingOriginTotalFee() *int32 {
	return s.BookingOriginTotalFee
}

func (s *TicketChangingApplyResponseBodyModule) GetBookingPriceChanged() *bool {
	return s.BookingPriceChanged
}

func (s *TicketChangingApplyResponseBodyModule) GetBtripOrderId() *int64 {
	return s.BtripOrderId
}

func (s *TicketChangingApplyResponseBodyModule) GetBtripSubOrderId() *int64 {
	return s.BtripSubOrderId
}

func (s *TicketChangingApplyResponseBodyModule) GetCanPay() *bool {
	return s.CanPay
}

func (s *TicketChangingApplyResponseBodyModule) GetChangeFee() *int64 {
	return s.ChangeFee
}

func (s *TicketChangingApplyResponseBodyModule) GetDeadlineTime() *string {
	return s.DeadlineTime
}

func (s *TicketChangingApplyResponseBodyModule) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *TicketChangingApplyResponseBodyModule) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *TicketChangingApplyResponseBodyModule) GetMaxRetryTimes() *int32 {
	return s.MaxRetryTimes
}

func (s *TicketChangingApplyResponseBodyModule) GetNextRetryInterval() *int64 {
	return s.NextRetryInterval
}

func (s *TicketChangingApplyResponseBodyModule) GetRetry() *bool {
	return s.Retry
}

func (s *TicketChangingApplyResponseBodyModule) GetRetryClientTips() *string {
	return s.RetryClientTips
}

func (s *TicketChangingApplyResponseBodyModule) GetStatus() *int32 {
	return s.Status
}

func (s *TicketChangingApplyResponseBodyModule) GetUpgradeFee() *int64 {
	return s.UpgradeFee
}

func (s *TicketChangingApplyResponseBodyModule) SetBookingChangedTotalFee(v int32) *TicketChangingApplyResponseBodyModule {
	s.BookingChangedTotalFee = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetBookingOriginTotalFee(v int32) *TicketChangingApplyResponseBodyModule {
	s.BookingOriginTotalFee = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetBookingPriceChanged(v bool) *TicketChangingApplyResponseBodyModule {
	s.BookingPriceChanged = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetBtripOrderId(v int64) *TicketChangingApplyResponseBodyModule {
	s.BtripOrderId = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetBtripSubOrderId(v int64) *TicketChangingApplyResponseBodyModule {
	s.BtripSubOrderId = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetCanPay(v bool) *TicketChangingApplyResponseBodyModule {
	s.CanPay = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetChangeFee(v int64) *TicketChangingApplyResponseBodyModule {
	s.ChangeFee = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetDeadlineTime(v string) *TicketChangingApplyResponseBodyModule {
	s.DeadlineTime = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetDisOrderId(v string) *TicketChangingApplyResponseBodyModule {
	s.DisOrderId = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetDisSubOrderId(v string) *TicketChangingApplyResponseBodyModule {
	s.DisSubOrderId = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetMaxRetryTimes(v int32) *TicketChangingApplyResponseBodyModule {
	s.MaxRetryTimes = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetNextRetryInterval(v int64) *TicketChangingApplyResponseBodyModule {
	s.NextRetryInterval = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetRetry(v bool) *TicketChangingApplyResponseBodyModule {
	s.Retry = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetRetryClientTips(v string) *TicketChangingApplyResponseBodyModule {
	s.RetryClientTips = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetStatus(v int32) *TicketChangingApplyResponseBodyModule {
	s.Status = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) SetUpgradeFee(v int64) *TicketChangingApplyResponseBodyModule {
	s.UpgradeFee = &v
	return s
}

func (s *TicketChangingApplyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
