// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundApplyV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightRefundApplyV2Request
	GetIsvName() *string
	SetOrderId(v string) *FlightRefundApplyV2Request
	GetOrderId() *string
	SetOutOrderId(v string) *FlightRefundApplyV2Request
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *FlightRefundApplyV2Request
	GetOutSubOrderId() *string
	SetPassengerSegmentRelations(v []*FlightRefundApplyV2RequestPassengerSegmentRelations) *FlightRefundApplyV2Request
	GetPassengerSegmentRelations() []*FlightRefundApplyV2RequestPassengerSegmentRelations
	SetPreCalType(v int32) *FlightRefundApplyV2Request
	GetPreCalType() *int32
	SetRefundReason(v string) *FlightRefundApplyV2Request
	GetRefundReason() *string
	SetRefundReasonType(v int32) *FlightRefundApplyV2Request
	GetRefundReasonType() *int32
	SetTicketNos(v []*string) *FlightRefundApplyV2Request
	GetTicketNos() []*string
	SetTotalRefundPrice(v int64) *FlightRefundApplyV2Request
	GetTotalRefundPrice() *int64
	SetUploadPictUrls(v string) *FlightRefundApplyV2Request
	GetUploadPictUrls() *string
	SetVoluntary(v bool) *FlightRefundApplyV2Request
	GetVoluntary() *bool
}

type FlightRefundApplyV2Request struct {
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
	OutSubOrderId             *string                                                `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	PassengerSegmentRelations []*FlightRefundApplyV2RequestPassengerSegmentRelations `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PreCalType   *int32  `json:"pre_cal_type,omitempty" xml:"pre_cal_type,omitempty"`
	RefundReason *string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// example:
	//
	// 2
	RefundReasonType *int32    `json:"refund_reason_type,omitempty" xml:"refund_reason_type,omitempty"`
	TicketNos        []*string `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty" type:"Repeated"`
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

func (s FlightRefundApplyV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyV2Request) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightRefundApplyV2Request) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightRefundApplyV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightRefundApplyV2Request) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *FlightRefundApplyV2Request) GetPassengerSegmentRelations() []*FlightRefundApplyV2RequestPassengerSegmentRelations {
	return s.PassengerSegmentRelations
}

func (s *FlightRefundApplyV2Request) GetPreCalType() *int32 {
	return s.PreCalType
}

func (s *FlightRefundApplyV2Request) GetRefundReason() *string {
	return s.RefundReason
}

func (s *FlightRefundApplyV2Request) GetRefundReasonType() *int32 {
	return s.RefundReasonType
}

func (s *FlightRefundApplyV2Request) GetTicketNos() []*string {
	return s.TicketNos
}

func (s *FlightRefundApplyV2Request) GetTotalRefundPrice() *int64 {
	return s.TotalRefundPrice
}

func (s *FlightRefundApplyV2Request) GetUploadPictUrls() *string {
	return s.UploadPictUrls
}

func (s *FlightRefundApplyV2Request) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightRefundApplyV2Request) SetIsvName(v string) *FlightRefundApplyV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetOrderId(v string) *FlightRefundApplyV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetOutOrderId(v string) *FlightRefundApplyV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetOutSubOrderId(v string) *FlightRefundApplyV2Request {
	s.OutSubOrderId = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetPassengerSegmentRelations(v []*FlightRefundApplyV2RequestPassengerSegmentRelations) *FlightRefundApplyV2Request {
	s.PassengerSegmentRelations = v
	return s
}

func (s *FlightRefundApplyV2Request) SetPreCalType(v int32) *FlightRefundApplyV2Request {
	s.PreCalType = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetRefundReason(v string) *FlightRefundApplyV2Request {
	s.RefundReason = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetRefundReasonType(v int32) *FlightRefundApplyV2Request {
	s.RefundReasonType = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetTicketNos(v []*string) *FlightRefundApplyV2Request {
	s.TicketNos = v
	return s
}

func (s *FlightRefundApplyV2Request) SetTotalRefundPrice(v int64) *FlightRefundApplyV2Request {
	s.TotalRefundPrice = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetUploadPictUrls(v string) *FlightRefundApplyV2Request {
	s.UploadPictUrls = &v
	return s
}

func (s *FlightRefundApplyV2Request) SetVoluntary(v bool) *FlightRefundApplyV2Request {
	s.Voluntary = &v
	return s
}

func (s *FlightRefundApplyV2Request) Validate() error {
	return dara.Validate(s)
}

type FlightRefundApplyV2RequestPassengerSegmentRelations struct {
	// example:
	//
	// 1075004
	PassengerId   *string   `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s FlightRefundApplyV2RequestPassengerSegmentRelations) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundApplyV2RequestPassengerSegmentRelations) GoString() string {
	return s.String()
}

func (s *FlightRefundApplyV2RequestPassengerSegmentRelations) GetPassengerId() *string {
	return s.PassengerId
}

func (s *FlightRefundApplyV2RequestPassengerSegmentRelations) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *FlightRefundApplyV2RequestPassengerSegmentRelations) SetPassengerId(v string) *FlightRefundApplyV2RequestPassengerSegmentRelations {
	s.PassengerId = &v
	return s
}

func (s *FlightRefundApplyV2RequestPassengerSegmentRelations) SetSegmentIdList(v []*string) *FlightRefundApplyV2RequestPassengerSegmentRelations {
	s.SegmentIdList = v
	return s
}

func (s *FlightRefundApplyV2RequestPassengerSegmentRelations) Validate() error {
	return dara.Validate(s)
}
