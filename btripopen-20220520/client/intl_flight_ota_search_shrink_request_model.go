// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightOtaSearchShrinkRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightOtaSearchShrinkRequest
	GetBuyerName() *string
	SetCabinType(v int32) *IntlFlightOtaSearchShrinkRequest
	GetCabinType() *int32
	SetIsvName(v string) *IntlFlightOtaSearchShrinkRequest
	GetIsvName() *string
	SetSearchJourneysShrink(v string) *IntlFlightOtaSearchShrinkRequest
	GetSearchJourneysShrink() *string
	SetSearchPassengerListShrink(v string) *IntlFlightOtaSearchShrinkRequest
	GetSearchPassengerListShrink() *string
	SetTripType(v int32) *IntlFlightOtaSearchShrinkRequest
	GetTripType() *int32
}

type IntlFlightOtaSearchShrinkRequest struct {
	// example:
	//
	// 10023
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// example:
	//
	// 0
	CabinType *int32 `json:"cabin_type,omitempty" xml:"cabin_type,omitempty"`
	// example:
	//
	// open12igetbis4o07v10B1TlOWcM00
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// This parameter is required.
	SearchJourneysShrink      *string `json:"search_journeys,omitempty" xml:"search_journeys,omitempty"`
	SearchPassengerListShrink *string `json:"search_passenger_list,omitempty" xml:"search_passenger_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s IntlFlightOtaSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightOtaSearchShrinkRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightOtaSearchShrinkRequest) GetCabinType() *int32 {
	return s.CabinType
}

func (s *IntlFlightOtaSearchShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightOtaSearchShrinkRequest) GetSearchJourneysShrink() *string {
	return s.SearchJourneysShrink
}

func (s *IntlFlightOtaSearchShrinkRequest) GetSearchPassengerListShrink() *string {
	return s.SearchPassengerListShrink
}

func (s *IntlFlightOtaSearchShrinkRequest) GetTripType() *int32 {
	return s.TripType
}

func (s *IntlFlightOtaSearchShrinkRequest) SetBtripUserId(v string) *IntlFlightOtaSearchShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightOtaSearchShrinkRequest) SetBuyerName(v string) *IntlFlightOtaSearchShrinkRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightOtaSearchShrinkRequest) SetCabinType(v int32) *IntlFlightOtaSearchShrinkRequest {
	s.CabinType = &v
	return s
}

func (s *IntlFlightOtaSearchShrinkRequest) SetIsvName(v string) *IntlFlightOtaSearchShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightOtaSearchShrinkRequest) SetSearchJourneysShrink(v string) *IntlFlightOtaSearchShrinkRequest {
	s.SearchJourneysShrink = &v
	return s
}

func (s *IntlFlightOtaSearchShrinkRequest) SetSearchPassengerListShrink(v string) *IntlFlightOtaSearchShrinkRequest {
	s.SearchPassengerListShrink = &v
	return s
}

func (s *IntlFlightOtaSearchShrinkRequest) SetTripType(v int32) *IntlFlightOtaSearchShrinkRequest {
	s.TripType = &v
	return s
}

func (s *IntlFlightOtaSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
