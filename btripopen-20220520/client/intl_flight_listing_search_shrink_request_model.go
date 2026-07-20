// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightListingSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightListingSearchShrinkRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightListingSearchShrinkRequest
	GetBuyerName() *string
	SetCabinType(v int32) *IntlFlightListingSearchShrinkRequest
	GetCabinType() *int32
	SetIsvName(v string) *IntlFlightListingSearchShrinkRequest
	GetIsvName() *string
	SetOutWheelSearch(v bool) *IntlFlightListingSearchShrinkRequest
	GetOutWheelSearch() *bool
	SetQueryRecordId(v string) *IntlFlightListingSearchShrinkRequest
	GetQueryRecordId() *string
	SetSearchJourneysShrink(v string) *IntlFlightListingSearchShrinkRequest
	GetSearchJourneysShrink() *string
	SetSearchMode(v int32) *IntlFlightListingSearchShrinkRequest
	GetSearchMode() *int32
	SetSearchPassengerListShrink(v string) *IntlFlightListingSearchShrinkRequest
	GetSearchPassengerListShrink() *string
	SetToken(v string) *IntlFlightListingSearchShrinkRequest
	GetToken() *string
	SetTripType(v int32) *IntlFlightListingSearchShrinkRequest
	GetTripType() *int32
}

type IntlFlightListingSearchShrinkRequest struct {
	// example:
	//
	// 10001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// ZHANGSAN
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	CabinType *int32 `json:"cabin_type,omitempty" xml:"cabin_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TEST
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	OutWheelSearch *bool `json:"out_wheel_search,omitempty" xml:"out_wheel_search,omitempty"`
	// example:
	//
	// 60b412-cc05-4d10-b570-
	QueryRecordId *string `json:"query_record_id,omitempty" xml:"query_record_id,omitempty"`
	// This parameter is required.
	SearchJourneysShrink *string `json:"search_journeys,omitempty" xml:"search_journeys,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	SearchMode                *int32  `json:"search_mode,omitempty" xml:"search_mode,omitempty"`
	SearchPassengerListShrink *string `json:"search_passenger_list,omitempty" xml:"search_passenger_list,omitempty"`
	// example:
	//
	// 9960b412-cc05-4d10-b570-93372d816807
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s IntlFlightListingSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightListingSearchShrinkRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightListingSearchShrinkRequest) GetCabinType() *int32 {
	return s.CabinType
}

func (s *IntlFlightListingSearchShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightListingSearchShrinkRequest) GetOutWheelSearch() *bool {
	return s.OutWheelSearch
}

func (s *IntlFlightListingSearchShrinkRequest) GetQueryRecordId() *string {
	return s.QueryRecordId
}

func (s *IntlFlightListingSearchShrinkRequest) GetSearchJourneysShrink() *string {
	return s.SearchJourneysShrink
}

func (s *IntlFlightListingSearchShrinkRequest) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *IntlFlightListingSearchShrinkRequest) GetSearchPassengerListShrink() *string {
	return s.SearchPassengerListShrink
}

func (s *IntlFlightListingSearchShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *IntlFlightListingSearchShrinkRequest) GetTripType() *int32 {
	return s.TripType
}

func (s *IntlFlightListingSearchShrinkRequest) SetBtripUserId(v string) *IntlFlightListingSearchShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetBuyerName(v string) *IntlFlightListingSearchShrinkRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetCabinType(v int32) *IntlFlightListingSearchShrinkRequest {
	s.CabinType = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetIsvName(v string) *IntlFlightListingSearchShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetOutWheelSearch(v bool) *IntlFlightListingSearchShrinkRequest {
	s.OutWheelSearch = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetQueryRecordId(v string) *IntlFlightListingSearchShrinkRequest {
	s.QueryRecordId = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetSearchJourneysShrink(v string) *IntlFlightListingSearchShrinkRequest {
	s.SearchJourneysShrink = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetSearchMode(v int32) *IntlFlightListingSearchShrinkRequest {
	s.SearchMode = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetSearchPassengerListShrink(v string) *IntlFlightListingSearchShrinkRequest {
	s.SearchPassengerListShrink = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetToken(v string) *IntlFlightListingSearchShrinkRequest {
	s.Token = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) SetTripType(v int32) *IntlFlightListingSearchShrinkRequest {
	s.TripType = &v
	return s
}

func (s *IntlFlightListingSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
