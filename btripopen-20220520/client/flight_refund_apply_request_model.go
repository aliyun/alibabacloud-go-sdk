// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpRefundPrice(v int64) *FlightRefundApplyRequest
	GetCorpRefundPrice() *int64
	SetDisOrderId(v string) *FlightRefundApplyRequest
	GetDisOrderId() *string
	SetDisSubOrderId(v string) *FlightRefundApplyRequest
	GetDisSubOrderId() *string
	SetDisplayRefundMoney(v string) *FlightRefundApplyRequest
	GetDisplayRefundMoney() *string
	SetExtra(v map[string]*string) *FlightRefundApplyRequest
	GetExtra() map[string]*string
	SetIsVoluntary(v int32) *FlightRefundApplyRequest
	GetIsVoluntary() *int32
	SetItemUnitIds(v string) *FlightRefundApplyRequest
	GetItemUnitIds() *string
	SetPassengerSegmentInfoList(v []*FlightRefundApplyRequestPassengerSegmentInfoList) *FlightRefundApplyRequest
	GetPassengerSegmentInfoList() []*FlightRefundApplyRequestPassengerSegmentInfoList
	SetPersonalRefundPrice(v int64) *FlightRefundApplyRequest
	GetPersonalRefundPrice() *int64
	SetReasonDetail(v string) *FlightRefundApplyRequest
	GetReasonDetail() *string
	SetReasonType(v int32) *FlightRefundApplyRequest
	GetReasonType() *int32
	SetRefundVoucherInfo(v []*string) *FlightRefundApplyRequest
	GetRefundVoucherInfo() []*string
	SetSessionId(v string) *FlightRefundApplyRequest
	GetSessionId() *string
	SetTotalRefundPrice(v int64) *FlightRefundApplyRequest
	GetTotalRefundPrice() *int64
}

type FlightRefundApplyRequest struct {
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
	DisplayRefundMoney *string            `json:"display_refund_money,omitempty" xml:"display_refund_money,omitempty"`
	Extra              map[string]*string `json:"extra,omitempty" xml:"extra,omitempty"`
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
	ItemUnitIds              *string                                             `json:"item_unit_ids,omitempty" xml:"item_unit_ids,omitempty"`
	PassengerSegmentInfoList []*FlightRefundApplyRequestPassengerSegmentInfoList `json:"passenger_segment_info_list,omitempty" xml:"passenger_segment_info_list,omitempty" type:"Repeated"`
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
	ReasonType        *int32    `json:"reason_type,omitempty" xml:"reason_type,omitempty"`
	RefundVoucherInfo []*string `json:"refund_voucher_info,omitempty" xml:"refund_voucher_info,omitempty" type:"Repeated"`
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

func (s FlightRefundApplyRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyRequest) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyRequest) GetCorpRefundPrice() *int64 {
	return s.CorpRefundPrice
}

func (s *FlightRefundApplyRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightRefundApplyRequest) GetDisSubOrderId() *string {
	return s.DisSubOrderId
}

func (s *FlightRefundApplyRequest) GetDisplayRefundMoney() *string {
	return s.DisplayRefundMoney
}

func (s *FlightRefundApplyRequest) GetExtra() map[string]*string {
	return s.Extra
}

func (s *FlightRefundApplyRequest) GetIsVoluntary() *int32 {
	return s.IsVoluntary
}

func (s *FlightRefundApplyRequest) GetItemUnitIds() *string {
	return s.ItemUnitIds
}

func (s *FlightRefundApplyRequest) GetPassengerSegmentInfoList() []*FlightRefundApplyRequestPassengerSegmentInfoList {
	return s.PassengerSegmentInfoList
}

func (s *FlightRefundApplyRequest) GetPersonalRefundPrice() *int64 {
	return s.PersonalRefundPrice
}

func (s *FlightRefundApplyRequest) GetReasonDetail() *string {
	return s.ReasonDetail
}

func (s *FlightRefundApplyRequest) GetReasonType() *int32 {
	return s.ReasonType
}

func (s *FlightRefundApplyRequest) GetRefundVoucherInfo() []*string {
	return s.RefundVoucherInfo
}

func (s *FlightRefundApplyRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightRefundApplyRequest) GetTotalRefundPrice() *int64 {
	return s.TotalRefundPrice
}

func (s *FlightRefundApplyRequest) SetCorpRefundPrice(v int64) *FlightRefundApplyRequest {
	s.CorpRefundPrice = &v
	return s
}

func (s *FlightRefundApplyRequest) SetDisOrderId(v string) *FlightRefundApplyRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightRefundApplyRequest) SetDisSubOrderId(v string) *FlightRefundApplyRequest {
	s.DisSubOrderId = &v
	return s
}

func (s *FlightRefundApplyRequest) SetDisplayRefundMoney(v string) *FlightRefundApplyRequest {
	s.DisplayRefundMoney = &v
	return s
}

func (s *FlightRefundApplyRequest) SetExtra(v map[string]*string) *FlightRefundApplyRequest {
	s.Extra = v
	return s
}

func (s *FlightRefundApplyRequest) SetIsVoluntary(v int32) *FlightRefundApplyRequest {
	s.IsVoluntary = &v
	return s
}

func (s *FlightRefundApplyRequest) SetItemUnitIds(v string) *FlightRefundApplyRequest {
	s.ItemUnitIds = &v
	return s
}

func (s *FlightRefundApplyRequest) SetPassengerSegmentInfoList(v []*FlightRefundApplyRequestPassengerSegmentInfoList) *FlightRefundApplyRequest {
	s.PassengerSegmentInfoList = v
	return s
}

func (s *FlightRefundApplyRequest) SetPersonalRefundPrice(v int64) *FlightRefundApplyRequest {
	s.PersonalRefundPrice = &v
	return s
}

func (s *FlightRefundApplyRequest) SetReasonDetail(v string) *FlightRefundApplyRequest {
	s.ReasonDetail = &v
	return s
}

func (s *FlightRefundApplyRequest) SetReasonType(v int32) *FlightRefundApplyRequest {
	s.ReasonType = &v
	return s
}

func (s *FlightRefundApplyRequest) SetRefundVoucherInfo(v []*string) *FlightRefundApplyRequest {
	s.RefundVoucherInfo = v
	return s
}

func (s *FlightRefundApplyRequest) SetSessionId(v string) *FlightRefundApplyRequest {
	s.SessionId = &v
	return s
}

func (s *FlightRefundApplyRequest) SetTotalRefundPrice(v int64) *FlightRefundApplyRequest {
	s.TotalRefundPrice = &v
	return s
}

func (s *FlightRefundApplyRequest) Validate() error {
	return dara.Validate(s)
}

type FlightRefundApplyRequestPassengerSegmentInfoList struct {
	FlightNo      *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 1245
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FlightRefundApplyRequestPassengerSegmentInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyRequestPassengerSegmentInfoList) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyRequestPassengerSegmentInfoList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightRefundApplyRequestPassengerSegmentInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightRefundApplyRequestPassengerSegmentInfoList) GetUserId() *string {
	return s.UserId
}

func (s *FlightRefundApplyRequestPassengerSegmentInfoList) SetFlightNo(v string) *FlightRefundApplyRequestPassengerSegmentInfoList {
	s.FlightNo = &v
	return s
}

func (s *FlightRefundApplyRequestPassengerSegmentInfoList) SetPassengerName(v string) *FlightRefundApplyRequestPassengerSegmentInfoList {
	s.PassengerName = &v
	return s
}

func (s *FlightRefundApplyRequestPassengerSegmentInfoList) SetUserId(v string) *FlightRefundApplyRequestPassengerSegmentInfoList {
	s.UserId = &v
	return s
}

func (s *FlightRefundApplyRequestPassengerSegmentInfoList) Validate() error {
	return dara.Validate(s)
}
