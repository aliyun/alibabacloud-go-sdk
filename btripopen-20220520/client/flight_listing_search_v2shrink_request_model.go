// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirlineCode(v string) *FlightListingSearchV2ShrinkRequest
	GetAirlineCode() *string
	SetCabinTypeListShrink(v string) *FlightListingSearchV2ShrinkRequest
	GetCabinTypeListShrink() *string
	SetDirectOnly(v bool) *FlightListingSearchV2ShrinkRequest
	GetDirectOnly() *bool
	SetIsvName(v string) *FlightListingSearchV2ShrinkRequest
	GetIsvName() *string
	SetNeedMultiClassPrice(v bool) *FlightListingSearchV2ShrinkRequest
	GetNeedMultiClassPrice() *bool
	SetNeedQueryServiceFee(v bool) *FlightListingSearchV2ShrinkRequest
	GetNeedQueryServiceFee() *bool
	SetNeedShareFlight(v bool) *FlightListingSearchV2ShrinkRequest
	GetNeedShareFlight() *bool
	SetNeedYCBestPrice(v bool) *FlightListingSearchV2ShrinkRequest
	GetNeedYCBestPrice() *bool
	SetSearchJourneysShrink(v string) *FlightListingSearchV2ShrinkRequest
	GetSearchJourneysShrink() *string
	SetSearchMode(v int32) *FlightListingSearchV2ShrinkRequest
	GetSearchMode() *int32
	SetTripType(v int32) *FlightListingSearchV2ShrinkRequest
	GetTripType() *int32
}

type FlightListingSearchV2ShrinkRequest struct {
	// example:
	//
	// CA
	AirlineCode         *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
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
	// false
	NeedMultiClassPrice *bool `json:"need_multi_class_price,omitempty" xml:"need_multi_class_price,omitempty"`
	// example:
	//
	// true
	NeedQueryServiceFee *bool `json:"need_query_service_fee,omitempty" xml:"need_query_service_fee,omitempty"`
	// example:
	//
	// true
	NeedShareFlight *bool `json:"need_share_flight,omitempty" xml:"need_share_flight,omitempty"`
	// example:
	//
	// false
	NeedYCBestPrice *bool `json:"need_y_c_best_price,omitempty" xml:"need_y_c_best_price,omitempty"`
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

func (s FlightListingSearchV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ShrinkRequest) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightListingSearchV2ShrinkRequest) GetCabinTypeListShrink() *string {
	return s.CabinTypeListShrink
}

func (s *FlightListingSearchV2ShrinkRequest) GetDirectOnly() *bool {
	return s.DirectOnly
}

func (s *FlightListingSearchV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightListingSearchV2ShrinkRequest) GetNeedMultiClassPrice() *bool {
	return s.NeedMultiClassPrice
}

func (s *FlightListingSearchV2ShrinkRequest) GetNeedQueryServiceFee() *bool {
	return s.NeedQueryServiceFee
}

func (s *FlightListingSearchV2ShrinkRequest) GetNeedShareFlight() *bool {
	return s.NeedShareFlight
}

func (s *FlightListingSearchV2ShrinkRequest) GetNeedYCBestPrice() *bool {
	return s.NeedYCBestPrice
}

func (s *FlightListingSearchV2ShrinkRequest) GetSearchJourneysShrink() *string {
	return s.SearchJourneysShrink
}

func (s *FlightListingSearchV2ShrinkRequest) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *FlightListingSearchV2ShrinkRequest) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightListingSearchV2ShrinkRequest) SetAirlineCode(v string) *FlightListingSearchV2ShrinkRequest {
	s.AirlineCode = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetCabinTypeListShrink(v string) *FlightListingSearchV2ShrinkRequest {
	s.CabinTypeListShrink = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetDirectOnly(v bool) *FlightListingSearchV2ShrinkRequest {
	s.DirectOnly = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetIsvName(v string) *FlightListingSearchV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetNeedMultiClassPrice(v bool) *FlightListingSearchV2ShrinkRequest {
	s.NeedMultiClassPrice = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetNeedQueryServiceFee(v bool) *FlightListingSearchV2ShrinkRequest {
	s.NeedQueryServiceFee = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetNeedShareFlight(v bool) *FlightListingSearchV2ShrinkRequest {
	s.NeedShareFlight = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetNeedYCBestPrice(v bool) *FlightListingSearchV2ShrinkRequest {
	s.NeedYCBestPrice = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetSearchJourneysShrink(v string) *FlightListingSearchV2ShrinkRequest {
	s.SearchJourneysShrink = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetSearchMode(v int32) *FlightListingSearchV2ShrinkRequest {
	s.SearchMode = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) SetTripType(v int32) *FlightListingSearchV2ShrinkRequest {
	s.TripType = &v
	return s
}

func (s *FlightListingSearchV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
