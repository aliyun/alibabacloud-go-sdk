// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisOrderId(v string) *FlightRefundPreCalShrinkRequest
	GetDisOrderId() *string
	SetIsVoluntary(v string) *FlightRefundPreCalShrinkRequest
	GetIsVoluntary() *string
	SetPassengerSegmentInfoListShrink(v string) *FlightRefundPreCalShrinkRequest
	GetPassengerSegmentInfoListShrink() *string
}

type FlightRefundPreCalShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	// example:
	//
	// 1
	IsVoluntary *string `json:"is_voluntary,omitempty" xml:"is_voluntary,omitempty"`
	// This parameter is required.
	PassengerSegmentInfoListShrink *string `json:"passenger_segment_info_list,omitempty" xml:"passenger_segment_info_list,omitempty"`
}

func (s FlightRefundPreCalShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightRefundPreCalShrinkRequest) GetIsVoluntary() *string {
	return s.IsVoluntary
}

func (s *FlightRefundPreCalShrinkRequest) GetPassengerSegmentInfoListShrink() *string {
	return s.PassengerSegmentInfoListShrink
}

func (s *FlightRefundPreCalShrinkRequest) SetDisOrderId(v string) *FlightRefundPreCalShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightRefundPreCalShrinkRequest) SetIsVoluntary(v string) *FlightRefundPreCalShrinkRequest {
	s.IsVoluntary = &v
	return s
}

func (s *FlightRefundPreCalShrinkRequest) SetPassengerSegmentInfoListShrink(v string) *FlightRefundPreCalShrinkRequest {
	s.PassengerSegmentInfoListShrink = &v
	return s
}

func (s *FlightRefundPreCalShrinkRequest) Validate() error {
	return dara.Validate(s)
}
