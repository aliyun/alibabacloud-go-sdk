// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAirlineCode(v string) *FlightListingSearchV2Request
	GetAirlineCode() *string
	SetCabinTypeList(v []*int32) *FlightListingSearchV2Request
	GetCabinTypeList() []*int32
	SetDirectOnly(v bool) *FlightListingSearchV2Request
	GetDirectOnly() *bool
	SetIsvName(v string) *FlightListingSearchV2Request
	GetIsvName() *string
	SetNeedMultiClassPrice(v bool) *FlightListingSearchV2Request
	GetNeedMultiClassPrice() *bool
	SetNeedQueryServiceFee(v bool) *FlightListingSearchV2Request
	GetNeedQueryServiceFee() *bool
	SetNeedShareFlight(v bool) *FlightListingSearchV2Request
	GetNeedShareFlight() *bool
	SetNeedYCBestPrice(v bool) *FlightListingSearchV2Request
	GetNeedYCBestPrice() *bool
	SetSearchJourneys(v []*FlightListingSearchV2RequestSearchJourneys) *FlightListingSearchV2Request
	GetSearchJourneys() []*FlightListingSearchV2RequestSearchJourneys
	SetSearchMode(v int32) *FlightListingSearchV2Request
	GetSearchMode() *int32
	SetTripType(v int32) *FlightListingSearchV2Request
	GetTripType() *int32
}

type FlightListingSearchV2Request struct {
	// example:
	//
	// CA
	AirlineCode   *string  `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	CabinTypeList []*int32 `json:"cabin_type_list,omitempty" xml:"cabin_type_list,omitempty" type:"Repeated"`
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
	SearchJourneys []*FlightListingSearchV2RequestSearchJourneys `json:"search_journeys,omitempty" xml:"search_journeys,omitempty" type:"Repeated"`
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

func (s FlightListingSearchV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2Request) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2Request) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightListingSearchV2Request) GetCabinTypeList() []*int32 {
	return s.CabinTypeList
}

func (s *FlightListingSearchV2Request) GetDirectOnly() *bool {
	return s.DirectOnly
}

func (s *FlightListingSearchV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightListingSearchV2Request) GetNeedMultiClassPrice() *bool {
	return s.NeedMultiClassPrice
}

func (s *FlightListingSearchV2Request) GetNeedQueryServiceFee() *bool {
	return s.NeedQueryServiceFee
}

func (s *FlightListingSearchV2Request) GetNeedShareFlight() *bool {
	return s.NeedShareFlight
}

func (s *FlightListingSearchV2Request) GetNeedYCBestPrice() *bool {
	return s.NeedYCBestPrice
}

func (s *FlightListingSearchV2Request) GetSearchJourneys() []*FlightListingSearchV2RequestSearchJourneys {
	return s.SearchJourneys
}

func (s *FlightListingSearchV2Request) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *FlightListingSearchV2Request) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightListingSearchV2Request) SetAirlineCode(v string) *FlightListingSearchV2Request {
	s.AirlineCode = &v
	return s
}

func (s *FlightListingSearchV2Request) SetCabinTypeList(v []*int32) *FlightListingSearchV2Request {
	s.CabinTypeList = v
	return s
}

func (s *FlightListingSearchV2Request) SetDirectOnly(v bool) *FlightListingSearchV2Request {
	s.DirectOnly = &v
	return s
}

func (s *FlightListingSearchV2Request) SetIsvName(v string) *FlightListingSearchV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightListingSearchV2Request) SetNeedMultiClassPrice(v bool) *FlightListingSearchV2Request {
	s.NeedMultiClassPrice = &v
	return s
}

func (s *FlightListingSearchV2Request) SetNeedQueryServiceFee(v bool) *FlightListingSearchV2Request {
	s.NeedQueryServiceFee = &v
	return s
}

func (s *FlightListingSearchV2Request) SetNeedShareFlight(v bool) *FlightListingSearchV2Request {
	s.NeedShareFlight = &v
	return s
}

func (s *FlightListingSearchV2Request) SetNeedYCBestPrice(v bool) *FlightListingSearchV2Request {
	s.NeedYCBestPrice = &v
	return s
}

func (s *FlightListingSearchV2Request) SetSearchJourneys(v []*FlightListingSearchV2RequestSearchJourneys) *FlightListingSearchV2Request {
	s.SearchJourneys = v
	return s
}

func (s *FlightListingSearchV2Request) SetSearchMode(v int32) *FlightListingSearchV2Request {
	s.SearchMode = &v
	return s
}

func (s *FlightListingSearchV2Request) SetTripType(v int32) *FlightListingSearchV2Request {
	s.TripType = &v
	return s
}

func (s *FlightListingSearchV2Request) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2RequestSearchJourneys struct {
	// This parameter is required.
	//
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-09-01
	DepDate         *string                                                      `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	SelectedFlights []*FlightListingSearchV2RequestSearchJourneysSelectedFlights `json:"selected_flights,omitempty" xml:"selected_flights,omitempty" type:"Repeated"`
}

func (s FlightListingSearchV2RequestSearchJourneys) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2RequestSearchJourneys) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2RequestSearchJourneys) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightListingSearchV2RequestSearchJourneys) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightListingSearchV2RequestSearchJourneys) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightListingSearchV2RequestSearchJourneys) GetSelectedFlights() []*FlightListingSearchV2RequestSearchJourneysSelectedFlights {
	return s.SelectedFlights
}

func (s *FlightListingSearchV2RequestSearchJourneys) SetArrCityCode(v string) *FlightListingSearchV2RequestSearchJourneys {
	s.ArrCityCode = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneys) SetDepCityCode(v string) *FlightListingSearchV2RequestSearchJourneys {
	s.DepCityCode = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneys) SetDepDate(v string) *FlightListingSearchV2RequestSearchJourneys {
	s.DepDate = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneys) SetSelectedFlights(v []*FlightListingSearchV2RequestSearchJourneysSelectedFlights) *FlightListingSearchV2RequestSearchJourneys {
	s.SelectedFlights = v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneys) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2RequestSearchJourneysSelectedFlights struct {
	// example:
	//
	// HGH
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// PKX
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2023-09-01 07:10:00
	FlightTime *string `json:"flight_time,omitempty" xml:"flight_time,omitempty"`
	// example:
	//
	// HO3925
	MarketFlightNo *string `json:"market_flight_no,omitempty" xml:"market_flight_no,omitempty"`
	// example:
	//
	// CX601
	OperateFlightNo *string `json:"operate_flight_no,omitempty" xml:"operate_flight_no,omitempty"`
}

func (s FlightListingSearchV2RequestSearchJourneysSelectedFlights) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2RequestSearchJourneysSelectedFlights) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) GetFlightTime() *string {
	return s.FlightTime
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) GetMarketFlightNo() *string {
	return s.MarketFlightNo
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) GetOperateFlightNo() *string {
	return s.OperateFlightNo
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) SetArrAirportCode(v string) *FlightListingSearchV2RequestSearchJourneysSelectedFlights {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) SetArrCityCode(v string) *FlightListingSearchV2RequestSearchJourneysSelectedFlights {
	s.ArrCityCode = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) SetDepAirportCode(v string) *FlightListingSearchV2RequestSearchJourneysSelectedFlights {
	s.DepAirportCode = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) SetDepCityCode(v string) *FlightListingSearchV2RequestSearchJourneysSelectedFlights {
	s.DepCityCode = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) SetFlightTime(v string) *FlightListingSearchV2RequestSearchJourneysSelectedFlights {
	s.FlightTime = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) SetMarketFlightNo(v string) *FlightListingSearchV2RequestSearchJourneysSelectedFlights {
	s.MarketFlightNo = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) SetOperateFlightNo(v string) *FlightListingSearchV2RequestSearchJourneysSelectedFlights {
	s.OperateFlightNo = &v
	return s
}

func (s *FlightListingSearchV2RequestSearchJourneysSelectedFlights) Validate() error {
	return dara.Validate(s)
}
