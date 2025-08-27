// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightRefundApplyV2ShrinkRequest
	GetIsvName() *string
	SetOrderId(v string) *FlightRefundApplyV2ShrinkRequest
	GetOrderId() *string
	SetOutOrderId(v string) *FlightRefundApplyV2ShrinkRequest
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *FlightRefundApplyV2ShrinkRequest
	GetOutSubOrderId() *string
	SetPassengerSegmentRelationsShrink(v string) *FlightRefundApplyV2ShrinkRequest
	GetPassengerSegmentRelationsShrink() *string
	SetPreCalType(v int32) *FlightRefundApplyV2ShrinkRequest
	GetPreCalType() *int32
	SetRefundReason(v string) *FlightRefundApplyV2ShrinkRequest
	GetRefundReason() *string
	SetRefundReasonType(v int32) *FlightRefundApplyV2ShrinkRequest
	GetRefundReasonType() *int32
	SetTicketNosShrink(v string) *FlightRefundApplyV2ShrinkRequest
	GetTicketNosShrink() *string
	SetTotalRefundPrice(v int64) *FlightRefundApplyV2ShrinkRequest
	GetTotalRefundPrice() *int64
	SetUploadPictUrls(v string) *FlightRefundApplyV2ShrinkRequest
	GetUploadPictUrls() *string
	SetVoluntary(v bool) *FlightRefundApplyV2ShrinkRequest
	GetVoluntary() *bool
}

type FlightRefundApplyV2ShrinkRequest struct {
	// example:
	//
	// cheshiapi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1683901850297448082
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1019195836916039
	OutSubOrderId                   *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	PassengerSegmentRelationsShrink *string `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty"`
	// example:
	//
	// 2
	PreCalType   *int32  `json:"pre_cal_type,omitempty" xml:"pre_cal_type,omitempty"`
	RefundReason *string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// example:
	//
	// 2
	RefundReasonType *int32  `json:"refund_reason_type,omitempty" xml:"refund_reason_type,omitempty"`
	TicketNosShrink  *string `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty"`
	// example:
	//
	// 10000
	TotalRefundPrice *int64  `json:"total_refund_price,omitempty" xml:"total_refund_price,omitempty"`
	UploadPictUrls   *string `json:"upload_pict_urls,omitempty" xml:"upload_pict_urls,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s FlightRefundApplyV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightRefundApplyV2ShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightRefundApplyV2ShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightRefundApplyV2ShrinkRequest) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightRefundApplyV2ShrinkRequest) GetPassengerSegmentRelationsShrink() *string {
	return s.PassengerSegmentRelationsShrink
}

func (s *FlightRefundApplyV2ShrinkRequest) GetPreCalType() *int32 {
	return s.PreCalType
}

func (s *FlightRefundApplyV2ShrinkRequest) GetRefundReason() *string {
	return s.RefundReason
}

func (s *FlightRefundApplyV2ShrinkRequest) GetRefundReasonType() *int32 {
	return s.RefundReasonType
}

func (s *FlightRefundApplyV2ShrinkRequest) GetTicketNosShrink() *string {
	return s.TicketNosShrink
}

func (s *FlightRefundApplyV2ShrinkRequest) GetTotalRefundPrice() *int64 {
	return s.TotalRefundPrice
}

func (s *FlightRefundApplyV2ShrinkRequest) GetUploadPictUrls() *string {
	return s.UploadPictUrls
}

func (s *FlightRefundApplyV2ShrinkRequest) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightRefundApplyV2ShrinkRequest) SetIsvName(v string) *FlightRefundApplyV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetOrderId(v string) *FlightRefundApplyV2ShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetOutOrderId(v string) *FlightRefundApplyV2ShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetOutSubOrderId(v string) *FlightRefundApplyV2ShrinkRequest {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetPassengerSegmentRelationsShrink(v string) *FlightRefundApplyV2ShrinkRequest {
	s.PassengerSegmentRelationsShrink = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetPreCalType(v int32) *FlightRefundApplyV2ShrinkRequest {
	s.PreCalType = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetRefundReason(v string) *FlightRefundApplyV2ShrinkRequest {
	s.RefundReason = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetRefundReasonType(v int32) *FlightRefundApplyV2ShrinkRequest {
	s.RefundReasonType = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetTicketNosShrink(v string) *FlightRefundApplyV2ShrinkRequest {
	s.TicketNosShrink = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetTotalRefundPrice(v int64) *FlightRefundApplyV2ShrinkRequest {
	s.TotalRefundPrice = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetUploadPictUrls(v string) *FlightRefundApplyV2ShrinkRequest {
	s.UploadPictUrls = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) SetVoluntary(v bool) *FlightRefundApplyV2ShrinkRequest {
	s.Voluntary = &v
	return s
}

func (s *FlightRefundApplyV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
