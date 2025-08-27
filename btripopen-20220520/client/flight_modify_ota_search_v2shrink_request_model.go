// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyOtaSearchV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCabinClassShrink(v string) *FlightModifyOtaSearchV2ShrinkRequest
	GetCabinClassShrink() *string
	SetDepDateShrink(v string) *FlightModifyOtaSearchV2ShrinkRequest
	GetDepDateShrink() *string
	SetIsvName(v string) *FlightModifyOtaSearchV2ShrinkRequest
	GetIsvName() *string
	SetOrderId(v int64) *FlightModifyOtaSearchV2ShrinkRequest
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyOtaSearchV2ShrinkRequest
	GetOutOrderId() *string
	SetPassengerSegmentRelationsShrink(v string) *FlightModifyOtaSearchV2ShrinkRequest
	GetPassengerSegmentRelationsShrink() *string
	SetSelectedSegmentsShrink(v string) *FlightModifyOtaSearchV2ShrinkRequest
	GetSelectedSegmentsShrink() *string
	SetSessionId(v string) *FlightModifyOtaSearchV2ShrinkRequest
	GetSessionId() *string
	SetVoluntary(v bool) *FlightModifyOtaSearchV2ShrinkRequest
	GetVoluntary() *bool
}

type FlightModifyOtaSearchV2ShrinkRequest struct {
	CabinClassShrink *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	DepDateShrink    *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// example:
	//
	// name
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1017002195370467200
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467200
	OutOrderId                      *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerSegmentRelationsShrink *string `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty"`
	SelectedSegmentsShrink          *string `json:"selected_segments,omitempty" xml:"selected_segments,omitempty"`
	// example:
	//
	// 590f17eca9374f20ac7e8ed8a7db2f35
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s FlightModifyOtaSearchV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetCabinClassShrink() *string {
	return s.CabinClassShrink
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetDepDateShrink() *string {
	return s.DepDateShrink
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetPassengerSegmentRelationsShrink() *string {
	return s.PassengerSegmentRelationsShrink
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetSelectedSegmentsShrink() *string {
	return s.SelectedSegmentsShrink
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetCabinClassShrink(v string) *FlightModifyOtaSearchV2ShrinkRequest {
	s.CabinClassShrink = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetDepDateShrink(v string) *FlightModifyOtaSearchV2ShrinkRequest {
	s.DepDateShrink = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetIsvName(v string) *FlightModifyOtaSearchV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetOrderId(v int64) *FlightModifyOtaSearchV2ShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetOutOrderId(v string) *FlightModifyOtaSearchV2ShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetPassengerSegmentRelationsShrink(v string) *FlightModifyOtaSearchV2ShrinkRequest {
	s.PassengerSegmentRelationsShrink = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetSelectedSegmentsShrink(v string) *FlightModifyOtaSearchV2ShrinkRequest {
	s.SelectedSegmentsShrink = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetSessionId(v string) *FlightModifyOtaSearchV2ShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) SetVoluntary(v bool) *FlightModifyOtaSearchV2ShrinkRequest {
	s.Voluntary = &v
	return s
}

func (s *FlightModifyOtaSearchV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
