// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopOtaSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightReShopOtaSearchRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopOtaSearchRequest
	GetOutOrderId() *string
	SetOutWheelSearch(v bool) *IntlFlightReShopOtaSearchRequest
	GetOutWheelSearch() *bool
	SetPassengerJourneyGroupKey(v string) *IntlFlightReShopOtaSearchRequest
	GetPassengerJourneyGroupKey() *string
	SetReShopReasonCode(v string) *IntlFlightReShopOtaSearchRequest
	GetReShopReasonCode() *string
	SetSearchJourneys(v []*IntlFlightReShopOtaSearchRequestSearchJourneys) *IntlFlightReShopOtaSearchRequest
	GetSearchJourneys() []*IntlFlightReShopOtaSearchRequestSearchJourneys
	SetSelectedPassengers(v []*IntlFlightReShopOtaSearchRequestSelectedPassengers) *IntlFlightReShopOtaSearchRequest
	GetSelectedPassengers() []*IntlFlightReShopOtaSearchRequestSelectedPassengers
	SetToken(v string) *IntlFlightReShopOtaSearchRequest
	GetToken() *string
}

type IntlFlightReShopOtaSearchRequest struct {
	// This parameter is required.
	OrderId        *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OutOrderId     *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	OutWheelSearch *bool   `json:"out_wheel_search,omitempty" xml:"out_wheel_search,omitempty"`
	// This parameter is required.
	PassengerJourneyGroupKey *string `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	ReShopReasonCode         *string `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	// This parameter is required.
	SearchJourneys []*IntlFlightReShopOtaSearchRequestSearchJourneys `json:"search_journeys,omitempty" xml:"search_journeys,omitempty" type:"Repeated"`
	// This parameter is required.
	SelectedPassengers []*IntlFlightReShopOtaSearchRequestSelectedPassengers `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty" type:"Repeated"`
	Token              *string                                               `json:"token,omitempty" xml:"token,omitempty"`
}

func (s IntlFlightReShopOtaSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopOtaSearchRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopOtaSearchRequest) GetOutWheelSearch() *bool {
	return s.OutWheelSearch
}

func (s *IntlFlightReShopOtaSearchRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopOtaSearchRequest) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopOtaSearchRequest) GetSearchJourneys() []*IntlFlightReShopOtaSearchRequestSearchJourneys {
	return s.SearchJourneys
}

func (s *IntlFlightReShopOtaSearchRequest) GetSelectedPassengers() []*IntlFlightReShopOtaSearchRequestSelectedPassengers {
	return s.SelectedPassengers
}

func (s *IntlFlightReShopOtaSearchRequest) GetToken() *string {
	return s.Token
}

func (s *IntlFlightReShopOtaSearchRequest) SetOrderId(v string) *IntlFlightReShopOtaSearchRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequest) SetOutOrderId(v string) *IntlFlightReShopOtaSearchRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequest) SetOutWheelSearch(v bool) *IntlFlightReShopOtaSearchRequest {
	s.OutWheelSearch = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopOtaSearchRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequest) SetReShopReasonCode(v string) *IntlFlightReShopOtaSearchRequest {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequest) SetSearchJourneys(v []*IntlFlightReShopOtaSearchRequestSearchJourneys) *IntlFlightReShopOtaSearchRequest {
	s.SearchJourneys = v
	return s
}

func (s *IntlFlightReShopOtaSearchRequest) SetSelectedPassengers(v []*IntlFlightReShopOtaSearchRequestSelectedPassengers) *IntlFlightReShopOtaSearchRequest {
	s.SelectedPassengers = v
	return s
}

func (s *IntlFlightReShopOtaSearchRequest) SetToken(v string) *IntlFlightReShopOtaSearchRequest {
	s.Token = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequest) Validate() error {
	if s.SearchJourneys != nil {
		for _, item := range s.SearchJourneys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SelectedPassengers != nil {
		for _, item := range s.SelectedPassengers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IntlFlightReShopOtaSearchRequestSearchJourneys struct {
	ArrCityCode     *string                                                          `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	DepCityCode     *string                                                          `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepDate         *string                                                          `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	SelectedFlights []*IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights `json:"selected_flights,omitempty" xml:"selected_flights,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopOtaSearchRequestSearchJourneys) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchRequestSearchJourneys) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) GetDepDate() *string {
	return s.DepDate
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) GetSelectedFlights() []*IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights {
	return s.SelectedFlights
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) SetArrCityCode(v string) *IntlFlightReShopOtaSearchRequestSearchJourneys {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) SetDepCityCode(v string) *IntlFlightReShopOtaSearchRequestSearchJourneys {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) SetDepDate(v string) *IntlFlightReShopOtaSearchRequestSearchJourneys {
	s.DepDate = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) SetSelectedFlights(v []*IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) *IntlFlightReShopOtaSearchRequestSearchJourneys {
	s.SelectedFlights = v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneys) Validate() error {
	if s.SelectedFlights != nil {
		for _, item := range s.SelectedFlights {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights struct {
	ArrCityCode    *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	DepCityCode    *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	FlightTime     *string `json:"flight_time,omitempty" xml:"flight_time,omitempty"`
	MarketFlightNo *string `json:"market_flight_no,omitempty" xml:"market_flight_no,omitempty"`
}

func (s IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) GetFlightTime() *string {
	return s.FlightTime
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) GetMarketFlightNo() *string {
	return s.MarketFlightNo
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) SetArrCityCode(v string) *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) SetDepCityCode(v string) *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) SetFlightTime(v string) *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights {
	s.FlightTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) SetMarketFlightNo(v string) *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights {
	s.MarketFlightNo = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSearchJourneysSelectedFlights) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchRequestSelectedPassengers struct {
	FullName    *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	PassengerId *int64  `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
}

func (s IntlFlightReShopOtaSearchRequestSelectedPassengers) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchRequestSelectedPassengers) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchRequestSelectedPassengers) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightReShopOtaSearchRequestSelectedPassengers) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopOtaSearchRequestSelectedPassengers) SetFullName(v string) *IntlFlightReShopOtaSearchRequestSelectedPassengers {
	s.FullName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSelectedPassengers) SetPassengerId(v int64) *IntlFlightReShopOtaSearchRequestSelectedPassengers {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopOtaSearchRequestSelectedPassengers) Validate() error {
	return dara.Validate(s)
}
