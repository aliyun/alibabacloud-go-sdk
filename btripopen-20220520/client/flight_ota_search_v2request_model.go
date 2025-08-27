// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCabinTypeList(v []*int32) *FlightOtaSearchV2Request
	GetCabinTypeList() []*int32
	SetDirectOnly(v bool) *FlightOtaSearchV2Request
	GetDirectOnly() *bool
	SetIsvName(v string) *FlightOtaSearchV2Request
	GetIsvName() *string
	SetNeedShareFlight(v bool) *FlightOtaSearchV2Request
	GetNeedShareFlight() *bool
	SetSearchJourneys(v []*FlightOtaSearchV2RequestSearchJourneys) *FlightOtaSearchV2Request
	GetSearchJourneys() []*FlightOtaSearchV2RequestSearchJourneys
	SetSearchMode(v int32) *FlightOtaSearchV2Request
	GetSearchMode() *int32
	SetTripType(v int32) *FlightOtaSearchV2Request
	GetTripType() *int32
}

type FlightOtaSearchV2Request struct {
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
	// true
	NeedShareFlight *bool `json:"need_share_flight,omitempty" xml:"need_share_flight,omitempty"`
	// This parameter is required.
	SearchJourneys []*FlightOtaSearchV2RequestSearchJourneys `json:"search_journeys,omitempty" xml:"search_journeys,omitempty" type:"Repeated"`
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

func (s FlightOtaSearchV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2Request) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2Request) GetCabinTypeList() []*int32 {
	return s.CabinTypeList
}

func (s *FlightOtaSearchV2Request) GetDirectOnly() *bool {
	return s.DirectOnly
}

func (s *FlightOtaSearchV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightOtaSearchV2Request) GetNeedShareFlight() *bool {
	return s.NeedShareFlight
}

func (s *FlightOtaSearchV2Request) GetSearchJourneys() []*FlightOtaSearchV2RequestSearchJourneys {
	return s.SearchJourneys
}

func (s *FlightOtaSearchV2Request) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *FlightOtaSearchV2Request) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightOtaSearchV2Request) SetCabinTypeList(v []*int32) *FlightOtaSearchV2Request {
	s.CabinTypeList = v
	return s
}

func (s *FlightOtaSearchV2Request) SetDirectOnly(v bool) *FlightOtaSearchV2Request {
	s.DirectOnly = &v
	return s
}

func (s *FlightOtaSearchV2Request) SetIsvName(v string) *FlightOtaSearchV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightOtaSearchV2Request) SetNeedShareFlight(v bool) *FlightOtaSearchV2Request {
	s.NeedShareFlight = &v
	return s
}

func (s *FlightOtaSearchV2Request) SetSearchJourneys(v []*FlightOtaSearchV2RequestSearchJourneys) *FlightOtaSearchV2Request {
	s.SearchJourneys = v
	return s
}

func (s *FlightOtaSearchV2Request) SetSearchMode(v int32) *FlightOtaSearchV2Request {
	s.SearchMode = &v
	return s
}

func (s *FlightOtaSearchV2Request) SetTripType(v int32) *FlightOtaSearchV2Request {
	s.TripType = &v
	return s
}

func (s *FlightOtaSearchV2Request) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2RequestSearchJourneys struct {
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
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	SelectedFlights []*FlightOtaSearchV2RequestSearchJourneysSelectedFlights `json:"selected_flights,omitempty" xml:"selected_flights,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchV2RequestSearchJourneys) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2RequestSearchJourneys) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2RequestSearchJourneys) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOtaSearchV2RequestSearchJourneys) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOtaSearchV2RequestSearchJourneys) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightOtaSearchV2RequestSearchJourneys) GetSelectedFlights() []*FlightOtaSearchV2RequestSearchJourneysSelectedFlights {
	return s.SelectedFlights
}

func (s *FlightOtaSearchV2RequestSearchJourneys) SetArrCityCode(v string) *FlightOtaSearchV2RequestSearchJourneys {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneys) SetDepCityCode(v string) *FlightOtaSearchV2RequestSearchJourneys {
	s.DepCityCode = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneys) SetDepDate(v string) *FlightOtaSearchV2RequestSearchJourneys {
	s.DepDate = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneys) SetSelectedFlights(v []*FlightOtaSearchV2RequestSearchJourneysSelectedFlights) *FlightOtaSearchV2RequestSearchJourneys {
	s.SelectedFlights = v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneys) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2RequestSearchJourneysSelectedFlights struct {
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
	// This parameter is required.
	//
	// example:
	//
	// HO3925
	MarketFlightNo *string `json:"market_flight_no,omitempty" xml:"market_flight_no,omitempty"`
	// example:
	//
	// CX601
	OperateFlightNo *string `json:"operate_flight_no,omitempty" xml:"operate_flight_no,omitempty"`
}

func (s FlightOtaSearchV2RequestSearchJourneysSelectedFlights) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2RequestSearchJourneysSelectedFlights) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) GetFlightTime() *string {
	return s.FlightTime
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) GetMarketFlightNo() *string {
	return s.MarketFlightNo
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) GetOperateFlightNo() *string {
	return s.OperateFlightNo
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) SetArrAirportCode(v string) *FlightOtaSearchV2RequestSearchJourneysSelectedFlights {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) SetArrCityCode(v string) *FlightOtaSearchV2RequestSearchJourneysSelectedFlights {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) SetDepAirportCode(v string) *FlightOtaSearchV2RequestSearchJourneysSelectedFlights {
	s.DepAirportCode = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) SetDepCityCode(v string) *FlightOtaSearchV2RequestSearchJourneysSelectedFlights {
	s.DepCityCode = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) SetFlightTime(v string) *FlightOtaSearchV2RequestSearchJourneysSelectedFlights {
	s.FlightTime = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) SetMarketFlightNo(v string) *FlightOtaSearchV2RequestSearchJourneysSelectedFlights {
	s.MarketFlightNo = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) SetOperateFlightNo(v string) *FlightOtaSearchV2RequestSearchJourneysSelectedFlights {
	s.OperateFlightNo = &v
	return s
}

func (s *FlightOtaSearchV2RequestSearchJourneysSelectedFlights) Validate() error {
	return dara.Validate(s)
}
