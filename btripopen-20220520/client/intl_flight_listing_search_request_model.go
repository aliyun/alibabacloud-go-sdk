// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightListingSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightListingSearchRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightListingSearchRequest
	GetBuyerName() *string
	SetCabinType(v int32) *IntlFlightListingSearchRequest
	GetCabinType() *int32
	SetIsvName(v string) *IntlFlightListingSearchRequest
	GetIsvName() *string
	SetOutWheelSearch(v bool) *IntlFlightListingSearchRequest
	GetOutWheelSearch() *bool
	SetQueryRecordId(v string) *IntlFlightListingSearchRequest
	GetQueryRecordId() *string
	SetSearchJourneys(v []*IntlFlightListingSearchRequestSearchJourneys) *IntlFlightListingSearchRequest
	GetSearchJourneys() []*IntlFlightListingSearchRequestSearchJourneys
	SetSearchMode(v int32) *IntlFlightListingSearchRequest
	GetSearchMode() *int32
	SetSearchPassengerList(v []*IntlFlightListingSearchRequestSearchPassengerList) *IntlFlightListingSearchRequest
	GetSearchPassengerList() []*IntlFlightListingSearchRequestSearchPassengerList
	SetToken(v string) *IntlFlightListingSearchRequest
	GetToken() *string
	SetTripType(v int32) *IntlFlightListingSearchRequest
	GetTripType() *int32
}

type IntlFlightListingSearchRequest struct {
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
	SearchJourneys []*IntlFlightListingSearchRequestSearchJourneys `json:"search_journeys,omitempty" xml:"search_journeys,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	SearchMode          *int32                                               `json:"search_mode,omitempty" xml:"search_mode,omitempty"`
	SearchPassengerList []*IntlFlightListingSearchRequestSearchPassengerList `json:"search_passenger_list,omitempty" xml:"search_passenger_list,omitempty" type:"Repeated"`
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

func (s IntlFlightListingSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightListingSearchRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightListingSearchRequest) GetCabinType() *int32 {
	return s.CabinType
}

func (s *IntlFlightListingSearchRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightListingSearchRequest) GetOutWheelSearch() *bool {
	return s.OutWheelSearch
}

func (s *IntlFlightListingSearchRequest) GetQueryRecordId() *string {
	return s.QueryRecordId
}

func (s *IntlFlightListingSearchRequest) GetSearchJourneys() []*IntlFlightListingSearchRequestSearchJourneys {
	return s.SearchJourneys
}

func (s *IntlFlightListingSearchRequest) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *IntlFlightListingSearchRequest) GetSearchPassengerList() []*IntlFlightListingSearchRequestSearchPassengerList {
	return s.SearchPassengerList
}

func (s *IntlFlightListingSearchRequest) GetToken() *string {
	return s.Token
}

func (s *IntlFlightListingSearchRequest) GetTripType() *int32 {
	return s.TripType
}

func (s *IntlFlightListingSearchRequest) SetBtripUserId(v string) *IntlFlightListingSearchRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightListingSearchRequest) SetBuyerName(v string) *IntlFlightListingSearchRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightListingSearchRequest) SetCabinType(v int32) *IntlFlightListingSearchRequest {
	s.CabinType = &v
	return s
}

func (s *IntlFlightListingSearchRequest) SetIsvName(v string) *IntlFlightListingSearchRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightListingSearchRequest) SetOutWheelSearch(v bool) *IntlFlightListingSearchRequest {
	s.OutWheelSearch = &v
	return s
}

func (s *IntlFlightListingSearchRequest) SetQueryRecordId(v string) *IntlFlightListingSearchRequest {
	s.QueryRecordId = &v
	return s
}

func (s *IntlFlightListingSearchRequest) SetSearchJourneys(v []*IntlFlightListingSearchRequestSearchJourneys) *IntlFlightListingSearchRequest {
	s.SearchJourneys = v
	return s
}

func (s *IntlFlightListingSearchRequest) SetSearchMode(v int32) *IntlFlightListingSearchRequest {
	s.SearchMode = &v
	return s
}

func (s *IntlFlightListingSearchRequest) SetSearchPassengerList(v []*IntlFlightListingSearchRequestSearchPassengerList) *IntlFlightListingSearchRequest {
	s.SearchPassengerList = v
	return s
}

func (s *IntlFlightListingSearchRequest) SetToken(v string) *IntlFlightListingSearchRequest {
	s.Token = &v
	return s
}

func (s *IntlFlightListingSearchRequest) SetTripType(v int32) *IntlFlightListingSearchRequest {
	s.TripType = &v
	return s
}

func (s *IntlFlightListingSearchRequest) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchRequestSearchJourneys struct {
	// This parameter is required.
	//
	// example:
	//
	// MEL
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HKG
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-12-28
	DepDate         *string                                                        `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	SelectedFlights []*IntlFlightListingSearchRequestSearchJourneysSelectedFlights `json:"selected_flights,omitempty" xml:"selected_flights,omitempty" type:"Repeated"`
}

func (s IntlFlightListingSearchRequestSearchJourneys) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchRequestSearchJourneys) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchRequestSearchJourneys) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightListingSearchRequestSearchJourneys) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightListingSearchRequestSearchJourneys) GetDepDate() *string {
	return s.DepDate
}

func (s *IntlFlightListingSearchRequestSearchJourneys) GetSelectedFlights() []*IntlFlightListingSearchRequestSearchJourneysSelectedFlights {
	return s.SelectedFlights
}

func (s *IntlFlightListingSearchRequestSearchJourneys) SetArrCityCode(v string) *IntlFlightListingSearchRequestSearchJourneys {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneys) SetDepCityCode(v string) *IntlFlightListingSearchRequestSearchJourneys {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneys) SetDepDate(v string) *IntlFlightListingSearchRequestSearchJourneys {
	s.DepDate = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneys) SetSelectedFlights(v []*IntlFlightListingSearchRequestSearchJourneysSelectedFlights) *IntlFlightListingSearchRequestSearchJourneys {
	s.SelectedFlights = v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneys) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchRequestSearchJourneysSelectedFlights struct {
	// example:
	//
	// HNY
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MEL
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// SZX
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HKG
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-12-28 12:00:00
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

func (s IntlFlightListingSearchRequestSearchJourneysSelectedFlights) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchRequestSearchJourneysSelectedFlights) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) GetFlightTime() *string {
	return s.FlightTime
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) GetMarketFlightNo() *string {
	return s.MarketFlightNo
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) GetOperateFlightNo() *string {
	return s.OperateFlightNo
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) SetArrAirportCode(v string) *IntlFlightListingSearchRequestSearchJourneysSelectedFlights {
	s.ArrAirportCode = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) SetArrCityCode(v string) *IntlFlightListingSearchRequestSearchJourneysSelectedFlights {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) SetDepAirportCode(v string) *IntlFlightListingSearchRequestSearchJourneysSelectedFlights {
	s.DepAirportCode = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) SetDepCityCode(v string) *IntlFlightListingSearchRequestSearchJourneysSelectedFlights {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) SetFlightTime(v string) *IntlFlightListingSearchRequestSearchJourneysSelectedFlights {
	s.FlightTime = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) SetMarketFlightNo(v string) *IntlFlightListingSearchRequestSearchJourneysSelectedFlights {
	s.MarketFlightNo = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) SetOperateFlightNo(v string) *IntlFlightListingSearchRequestSearchJourneysSelectedFlights {
	s.OperateFlightNo = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchJourneysSelectedFlights) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchRequestSearchPassengerList struct {
	// This parameter is required.
	//
	// example:
	//
	// 10012301201
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

func (s IntlFlightListingSearchRequestSearchPassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchRequestSearchPassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) GetCertNo() *string {
	return s.CertNo
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) GetCertType() *int32 {
	return s.CertType
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) GetType() *int32 {
	return s.Type
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) SetCertNo(v string) *IntlFlightListingSearchRequestSearchPassengerList {
	s.CertNo = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) SetCertType(v int32) *IntlFlightListingSearchRequestSearchPassengerList {
	s.CertType = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) SetFullName(v string) *IntlFlightListingSearchRequestSearchPassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) SetType(v int32) *IntlFlightListingSearchRequestSearchPassengerList {
	s.Type = &v
	return s
}

func (s *IntlFlightListingSearchRequestSearchPassengerList) Validate() error {
	return dara.Validate(s)
}
