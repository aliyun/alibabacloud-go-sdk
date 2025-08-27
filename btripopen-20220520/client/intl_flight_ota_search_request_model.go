// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightOtaSearchRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightOtaSearchRequest
	GetBuyerName() *string
	SetCabinType(v int32) *IntlFlightOtaSearchRequest
	GetCabinType() *int32
	SetIsvName(v string) *IntlFlightOtaSearchRequest
	GetIsvName() *string
	SetSearchJourneys(v []*IntlFlightOtaSearchRequestSearchJourneys) *IntlFlightOtaSearchRequest
	GetSearchJourneys() []*IntlFlightOtaSearchRequestSearchJourneys
	SetSearchPassengerList(v []*IntlFlightOtaSearchRequestSearchPassengerList) *IntlFlightOtaSearchRequest
	GetSearchPassengerList() []*IntlFlightOtaSearchRequestSearchPassengerList
	SetTripType(v int32) *IntlFlightOtaSearchRequest
	GetTripType() *int32
}

type IntlFlightOtaSearchRequest struct {
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
	SearchJourneys      []*IntlFlightOtaSearchRequestSearchJourneys      `json:"search_journeys,omitempty" xml:"search_journeys,omitempty" type:"Repeated"`
	SearchPassengerList []*IntlFlightOtaSearchRequestSearchPassengerList `json:"search_passenger_list,omitempty" xml:"search_passenger_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s IntlFlightOtaSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightOtaSearchRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightOtaSearchRequest) GetCabinType() *int32 {
	return s.CabinType
}

func (s *IntlFlightOtaSearchRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightOtaSearchRequest) GetSearchJourneys() []*IntlFlightOtaSearchRequestSearchJourneys {
	return s.SearchJourneys
}

func (s *IntlFlightOtaSearchRequest) GetSearchPassengerList() []*IntlFlightOtaSearchRequestSearchPassengerList {
	return s.SearchPassengerList
}

func (s *IntlFlightOtaSearchRequest) GetTripType() *int32 {
	return s.TripType
}

func (s *IntlFlightOtaSearchRequest) SetBtripUserId(v string) *IntlFlightOtaSearchRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightOtaSearchRequest) SetBuyerName(v string) *IntlFlightOtaSearchRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightOtaSearchRequest) SetCabinType(v int32) *IntlFlightOtaSearchRequest {
	s.CabinType = &v
	return s
}

func (s *IntlFlightOtaSearchRequest) SetIsvName(v string) *IntlFlightOtaSearchRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightOtaSearchRequest) SetSearchJourneys(v []*IntlFlightOtaSearchRequestSearchJourneys) *IntlFlightOtaSearchRequest {
	s.SearchJourneys = v
	return s
}

func (s *IntlFlightOtaSearchRequest) SetSearchPassengerList(v []*IntlFlightOtaSearchRequestSearchPassengerList) *IntlFlightOtaSearchRequest {
	s.SearchPassengerList = v
	return s
}

func (s *IntlFlightOtaSearchRequest) SetTripType(v int32) *IntlFlightOtaSearchRequest {
	s.TripType = &v
	return s
}

func (s *IntlFlightOtaSearchRequest) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchRequestSearchJourneys struct {
	// This parameter is required.
	//
	// example:
	//
	// ZQZ
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-08-15
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	SelectedFlights []*IntlFlightOtaSearchRequestSearchJourneysSelectedFlights `json:"selected_flights,omitempty" xml:"selected_flights,omitempty" type:"Repeated"`
}

func (s IntlFlightOtaSearchRequestSearchJourneys) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchRequestSearchJourneys) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) GetDepDate() *string {
	return s.DepDate
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) GetSelectedFlights() []*IntlFlightOtaSearchRequestSearchJourneysSelectedFlights {
	return s.SelectedFlights
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) SetArrCityCode(v string) *IntlFlightOtaSearchRequestSearchJourneys {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) SetDepCityCode(v string) *IntlFlightOtaSearchRequestSearchJourneys {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) SetDepDate(v string) *IntlFlightOtaSearchRequestSearchJourneys {
	s.DepDate = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) SetSelectedFlights(v []*IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) *IntlFlightOtaSearchRequestSearchJourneys {
	s.SelectedFlights = v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneys) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchRequestSearchJourneysSelectedFlights struct {
	// example:
	//
	// HGH
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NNG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// KOW
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CTU
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-08-15 12:00:00
	FlightTime *string `json:"flight_time,omitempty" xml:"flight_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// KA5809
	MarketFlightNo *string `json:"market_flight_no,omitempty" xml:"market_flight_no,omitempty"`
	// example:
	//
	// CX601
	OperateFlightNo *string `json:"operate_flight_no,omitempty" xml:"operate_flight_no,omitempty"`
}

func (s IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) GetFlightTime() *string {
	return s.FlightTime
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) GetMarketFlightNo() *string {
	return s.MarketFlightNo
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) GetOperateFlightNo() *string {
	return s.OperateFlightNo
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) SetArrAirportCode(v string) *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights {
	s.ArrAirportCode = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) SetArrCityCode(v string) *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) SetDepAirportCode(v string) *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights {
	s.DepAirportCode = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) SetDepCityCode(v string) *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) SetFlightTime(v string) *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights {
	s.FlightTime = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) SetMarketFlightNo(v string) *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights {
	s.MarketFlightNo = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) SetOperateFlightNo(v string) *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights {
	s.OperateFlightNo = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchJourneysSelectedFlights) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchRequestSearchPassengerList struct {
	// This parameter is required.
	//
	// example:
	//
	// 13412341234
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	CertType *int32 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZHANGSAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IntlFlightOtaSearchRequestSearchPassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchRequestSearchPassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) GetCertNo() *string {
	return s.CertNo
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) GetCertType() *int32 {
	return s.CertType
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) GetType() *int32 {
	return s.Type
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) SetCertNo(v string) *IntlFlightOtaSearchRequestSearchPassengerList {
	s.CertNo = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) SetCertType(v int32) *IntlFlightOtaSearchRequestSearchPassengerList {
	s.CertType = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) SetFullName(v string) *IntlFlightOtaSearchRequestSearchPassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) SetType(v int32) *IntlFlightOtaSearchRequestSearchPassengerList {
	s.Type = &v
	return s
}

func (s *IntlFlightOtaSearchRequestSearchPassengerList) Validate() error {
	return dara.Validate(s)
}
