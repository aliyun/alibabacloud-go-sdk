// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCabinTypeListShrink(v string) *FlightOtaSearchV2ShrinkRequest
	GetCabinTypeListShrink() *string
	SetDirectOnly(v bool) *FlightOtaSearchV2ShrinkRequest
	GetDirectOnly() *bool
	SetIsvName(v string) *FlightOtaSearchV2ShrinkRequest
	GetIsvName() *string
	SetNeedShareFlight(v bool) *FlightOtaSearchV2ShrinkRequest
	GetNeedShareFlight() *bool
	SetSearchJourneysShrink(v string) *FlightOtaSearchV2ShrinkRequest
	GetSearchJourneysShrink() *string
	SetSearchMode(v int32) *FlightOtaSearchV2ShrinkRequest
	GetSearchMode() *int32
	SetTripType(v int32) *FlightOtaSearchV2ShrinkRequest
	GetTripType() *int32
}

type FlightOtaSearchV2ShrinkRequest struct {
	CabinTypeListShrink *string `json:"cabin_type_list,omitempty" xml:"cabin_type_list,omitempty"`
	// example:
	//
	// true
	DirectOnly *bool `json:"direct_only,omitempty" xml:"direct_only,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cheshi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// true
	NeedShareFlight *bool `json:"need_share_flight,omitempty" xml:"need_share_flight,omitempty"`
	// This parameter is required.
	SearchJourneysShrink *string `json:"search_journeys,omitempty" xml:"search_journeys,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	SearchMode *int32 `json:"search_mode,omitempty" xml:"search_mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightOtaSearchV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ShrinkRequest) GetCabinTypeListShrink() *string {
	return s.CabinTypeListShrink
}

func (s *FlightOtaSearchV2ShrinkRequest) GetDirectOnly() *bool {
	return s.DirectOnly
}

func (s *FlightOtaSearchV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightOtaSearchV2ShrinkRequest) GetNeedShareFlight() *bool {
	return s.NeedShareFlight
}

func (s *FlightOtaSearchV2ShrinkRequest) GetSearchJourneysShrink() *string {
	return s.SearchJourneysShrink
}

func (s *FlightOtaSearchV2ShrinkRequest) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *FlightOtaSearchV2ShrinkRequest) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightOtaSearchV2ShrinkRequest) SetCabinTypeListShrink(v string) *FlightOtaSearchV2ShrinkRequest {
	s.CabinTypeListShrink = &v
	return s
}

func (s *FlightOtaSearchV2ShrinkRequest) SetDirectOnly(v bool) *FlightOtaSearchV2ShrinkRequest {
	s.DirectOnly = &v
	return s
}

func (s *FlightOtaSearchV2ShrinkRequest) SetIsvName(v string) *FlightOtaSearchV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightOtaSearchV2ShrinkRequest) SetNeedShareFlight(v bool) *FlightOtaSearchV2ShrinkRequest {
	s.NeedShareFlight = &v
	return s
}

func (s *FlightOtaSearchV2ShrinkRequest) SetSearchJourneysShrink(v string) *FlightOtaSearchV2ShrinkRequest {
	s.SearchJourneysShrink = &v
	return s
}

func (s *FlightOtaSearchV2ShrinkRequest) SetSearchMode(v int32) *FlightOtaSearchV2ShrinkRequest {
	s.SearchMode = &v
	return s
}

func (s *FlightOtaSearchV2ShrinkRequest) SetTripType(v int32) *FlightOtaSearchV2ShrinkRequest {
	s.TripType = &v
	return s
}

func (s *FlightOtaSearchV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
