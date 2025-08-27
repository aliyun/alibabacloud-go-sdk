// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpRefundPrice(v int64) *FlightRefundApplyShrinkRequest
	GetCorpRefundPrice() *int64
	SetDisOrderId(v string) *FlightRefundApplyShrinkRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *FlightRefundApplyShrinkRequest
	GetDisSubOrderId() *string
	SetDisplayRefundMoney(v string) *FlightRefundApplyShrinkRequest
	GetDisplayRefundMoney() *string
	SetExtraShrink(v string) *FlightRefundApplyShrinkRequest
	GetExtraShrink() *string
	SetIsVoluntary(v int32) *FlightRefundApplyShrinkRequest
	GetIsVoluntary() *int32
	SetItemUnitIds(v string) *FlightRefundApplyShrinkRequest
	GetItemUnitIds() *string
	SetPassengerSegmentInfoListShrink(v string) *FlightRefundApplyShrinkRequest
	GetPassengerSegmentInfoListShrink() *string
	SetPersonalRefundPrice(v int64) *FlightRefundApplyShrinkRequest
	GetPersonalRefundPrice() *int64
	SetReasonDetail(v string) *FlightRefundApplyShrinkRequest
	GetReasonDetail() *string
	SetReasonType(v int32) *FlightRefundApplyShrinkRequest
	GetReasonType() *int32
	SetRefundVoucherInfoShrink(v string) *FlightRefundApplyShrinkRequest
	GetRefundVoucherInfoShrink() *string
	SetSessionId(v string) *FlightRefundApplyShrinkRequest
	GetSessionId() *string
	SetTotalRefundPrice(v int64) *FlightRefundApplyShrinkRequest
	GetTotalRefundPrice() *int64
}

type FlightRefundApplyShrinkRequest struct {
	// example:
	//
	// 100
	CorpRefundPrice *int64 `json:"corp_refund_price,omitempty" xml:"corp_refund_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// refu123
	DisSubOrderId *string `json:"dis_sub_order_id,omitempty" xml:"dis_sub_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	DisplayRefundMoney *string `json:"display_refund_money,omitempty" xml:"display_refund_money,omitempty"`
	ExtraShrink        *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	IsVoluntary *int32 `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FlightItem_9966772382
	ItemUnitIds                    *string `json:"item_unit_ids,omitempty" xml:"item_unit_ids,omitempty"`
	PassengerSegmentInfoListShrink *string `json:"passenger_segment_info_list,omitempty" xml:"passenger_segment_info_list,omitempty"`
	// example:
	//
	// 100
	PersonalRefundPrice *int64  `json:"personal_refund_price,omitempty" xml:"personal_refund_price,omitempty"`
	ReasonDetail        *string `json:"reason_detail,omitempty" xml:"reason_detail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReasonType              *int32  `json:"reason_type,omitempty" xml:"reason_type,omitempty"`
	RefundVoucherInfoShrink *string `json:"refund_voucher_info,omitempty" xml:"refund_voucher_info,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0f9418cd2ce34af49ab0de16fea166d1
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// 100
	TotalRefundPrice *int64 `json:"total_refund_price,omitempty" xml:"total_refund_price,omitempty"`
}

func (s FlightRefundApplyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyShrinkRequest) GetCorpRefundPrice() *int64 {
	return s.CorpRefundPrice
}

func (s *FlightRefundApplyShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightRefundApplyShrinkRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *FlightRefundApplyShrinkRequest) GetDisplayRefundMoney() *string {
	return s.DisplayRefundMoney
}

func (s *FlightRefundApplyShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
}

func (s *FlightRefundApplyShrinkRequest) GetIsVoluntary() *int32 {
	return s.IsVoluntary
}

func (s *FlightRefundApplyShrinkRequest) GetItemUnitIds() *string {
	return s.ItemUnitIds
}

func (s *FlightRefundApplyShrinkRequest) GetPassengerSegmentInfoListShrink() *string {
	return s.PassengerSegmentInfoListShrink
}

func (s *FlightRefundApplyShrinkRequest) GetPersonalRefundPrice() *int64 {
	return s.PersonalRefundPrice
}

func (s *FlightRefundApplyShrinkRequest) GetReasonDetail() *string {
	return s.ReasonDetail
}

func (s *FlightRefundApplyShrinkRequest) GetReasonType() *int32 {
	return s.ReasonType
}

func (s *FlightRefundApplyShrinkRequest) GetRefundVoucherInfoShrink() *string {
	return s.RefundVoucherInfoShrink
}

func (s *FlightRefundApplyShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightRefundApplyShrinkRequest) GetTotalRefundPrice() *int64 {
	return s.TotalRefundPrice
}

func (s *FlightRefundApplyShrinkRequest) SetCorpRefundPrice(v int64) *FlightRefundApplyShrinkRequest {
	s.CorpRefundPrice = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetDisOrderId(v string) *FlightRefundApplyShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetDisSubOrderId(v string) *FlightRefundApplyShrinkRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetDisplayRefundMoney(v string) *FlightRefundApplyShrinkRequest {
	s.DisplayRefundMoney = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetExtraShrink(v string) *FlightRefundApplyShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetIsVoluntary(v int32) *FlightRefundApplyShrinkRequest {
	s.IsVoluntary = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetItemUnitIds(v string) *FlightRefundApplyShrinkRequest {
	s.ItemUnitIds = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetPassengerSegmentInfoListShrink(v string) *FlightRefundApplyShrinkRequest {
	s.PassengerSegmentInfoListShrink = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetPersonalRefundPrice(v int64) *FlightRefundApplyShrinkRequest {
	s.PersonalRefundPrice = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetReasonDetail(v string) *FlightRefundApplyShrinkRequest {
	s.ReasonDetail = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetReasonType(v int32) *FlightRefundApplyShrinkRequest {
	s.ReasonType = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetRefundVoucherInfoShrink(v string) *FlightRefundApplyShrinkRequest {
	s.RefundVoucherInfoShrink = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetSessionId(v string) *FlightRefundApplyShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) SetTotalRefundPrice(v int64) *FlightRefundApplyShrinkRequest {
	s.TotalRefundPrice = &v
	return s
}

func (s *FlightRefundApplyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
