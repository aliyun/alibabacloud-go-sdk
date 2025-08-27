// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightRefundPreCalV2ShrinkRequest
	GetIsvName() *string
	SetOrderId(v string) *FlightRefundPreCalV2ShrinkRequest
	GetOrderId() *string
	SetOutOrderId(v string) *FlightRefundPreCalV2ShrinkRequest
	GetOutOrderId() *string
	SetPassengerSegmentRelationsShrink(v string) *FlightRefundPreCalV2ShrinkRequest
	GetPassengerSegmentRelationsShrink() *string
	SetPreCalType(v int32) *FlightRefundPreCalV2ShrinkRequest
	GetPreCalType() *int32
	SetTicketNosShrink(v string) *FlightRefundPreCalV2ShrinkRequest
	GetTicketNosShrink() *string
	SetVoluntary(v bool) *FlightRefundPreCalV2ShrinkRequest
	GetVoluntary() *bool
}

type FlightRefundPreCalV2ShrinkRequest struct {
	// example:
	//
	// cheshiapi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 3454043907950204159
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId                      *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerSegmentRelationsShrink *string `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty"`
	// example:
	//
	// 2
	PreCalType      *int32  `json:"pre_cal_type,omitempty" xml:"pre_cal_type,omitempty"`
	TicketNosShrink *string `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s FlightRefundPreCalV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightRefundPreCalV2ShrinkRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightRefundPreCalV2ShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightRefundPreCalV2ShrinkRequest) GetPassengerSegmentRelationsShrink() *string {
	return s.PassengerSegmentRelationsShrink
}

func (s *FlightRefundPreCalV2ShrinkRequest) GetPreCalType() *int32 {
	return s.PreCalType
}

func (s *FlightRefundPreCalV2ShrinkRequest) GetTicketNosShrink() *string {
	return s.TicketNosShrink
}

func (s *FlightRefundPreCalV2ShrinkRequest) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightRefundPreCalV2ShrinkRequest) SetIsvName(v string) *FlightRefundPreCalV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightRefundPreCalV2ShrinkRequest) SetOrderId(v string) *FlightRefundPreCalV2ShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *FlightRefundPreCalV2ShrinkRequest) SetOutOrderId(v string) *FlightRefundPreCalV2ShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *FlightRefundPreCalV2ShrinkRequest) SetPassengerSegmentRelationsShrink(v string) *FlightRefundPreCalV2ShrinkRequest {
	s.PassengerSegmentRelationsShrink = &v
	return s
}

func (s *FlightRefundPreCalV2ShrinkRequest) SetPreCalType(v int32) *FlightRefundPreCalV2ShrinkRequest {
	s.PreCalType = &v
	return s
}

func (s *FlightRefundPreCalV2ShrinkRequest) SetTicketNosShrink(v string) *FlightRefundPreCalV2ShrinkRequest {
	s.TicketNosShrink = &v
	return s
}

func (s *FlightRefundPreCalV2ShrinkRequest) SetVoluntary(v bool) *FlightRefundPreCalV2ShrinkRequest {
	s.Voluntary = &v
	return s
}

func (s *FlightRefundPreCalV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
