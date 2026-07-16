// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopListSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightReShopListSearchRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopListSearchRequest
	GetOutOrderId() *string
	SetOutWheelSearch(v bool) *IntlFlightReShopListSearchRequest
	GetOutWheelSearch() *bool
	SetPassengerJourneyGroupKey(v string) *IntlFlightReShopListSearchRequest
	GetPassengerJourneyGroupKey() *string
	SetReShopReasonCode(v string) *IntlFlightReShopListSearchRequest
	GetReShopReasonCode() *string
	SetSearchJourneys(v []*IntlFlightReShopListSearchRequestSearchJourneys) *IntlFlightReShopListSearchRequest
	GetSearchJourneys() []*IntlFlightReShopListSearchRequestSearchJourneys
	SetSelectedPassengers(v []*IntlFlightReShopListSearchRequestSelectedPassengers) *IntlFlightReShopListSearchRequest
	GetSelectedPassengers() []*IntlFlightReShopListSearchRequestSelectedPassengers
	SetToken(v string) *IntlFlightReShopListSearchRequest
	GetToken() *string
}

type IntlFlightReShopListSearchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1002027205317939247
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 3985893777358602240
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// false
	OutWheelSearch *bool `json:"out_wheel_search,omitempty" xml:"out_wheel_search,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// edcac4f4c79d40ccb141ddb6da567e65
	PassengerJourneyGroupKey *string `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	// example:
	//
	// 0
	ReShopReasonCode *string `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	// This parameter is required.
	SearchJourneys []*IntlFlightReShopListSearchRequestSearchJourneys `json:"search_journeys,omitempty" xml:"search_journeys,omitempty" type:"Repeated"`
	// This parameter is required.
	SelectedPassengers []*IntlFlightReShopListSearchRequestSelectedPassengers `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty" type:"Repeated"`
	// example:
	//
	// 0305b8203a7767626f911d97a91a9473
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s IntlFlightReShopListSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopListSearchRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopListSearchRequest) GetOutWheelSearch() *bool {
	return s.OutWheelSearch
}

func (s *IntlFlightReShopListSearchRequest) GetPassengerJourneyGroupKey() *string {
	return s.PassengerJourneyGroupKey
}

func (s *IntlFlightReShopListSearchRequest) GetReShopReasonCode() *string {
	return s.ReShopReasonCode
}

func (s *IntlFlightReShopListSearchRequest) GetSearchJourneys() []*IntlFlightReShopListSearchRequestSearchJourneys {
	return s.SearchJourneys
}

func (s *IntlFlightReShopListSearchRequest) GetSelectedPassengers() []*IntlFlightReShopListSearchRequestSelectedPassengers {
	return s.SelectedPassengers
}

func (s *IntlFlightReShopListSearchRequest) GetToken() *string {
	return s.Token
}

func (s *IntlFlightReShopListSearchRequest) SetOrderId(v string) *IntlFlightReShopListSearchRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopListSearchRequest) SetOutOrderId(v string) *IntlFlightReShopListSearchRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopListSearchRequest) SetOutWheelSearch(v bool) *IntlFlightReShopListSearchRequest {
	s.OutWheelSearch = &v
	return s
}

func (s *IntlFlightReShopListSearchRequest) SetPassengerJourneyGroupKey(v string) *IntlFlightReShopListSearchRequest {
	s.PassengerJourneyGroupKey = &v
	return s
}

func (s *IntlFlightReShopListSearchRequest) SetReShopReasonCode(v string) *IntlFlightReShopListSearchRequest {
	s.ReShopReasonCode = &v
	return s
}

func (s *IntlFlightReShopListSearchRequest) SetSearchJourneys(v []*IntlFlightReShopListSearchRequestSearchJourneys) *IntlFlightReShopListSearchRequest {
	s.SearchJourneys = v
	return s
}

func (s *IntlFlightReShopListSearchRequest) SetSelectedPassengers(v []*IntlFlightReShopListSearchRequestSelectedPassengers) *IntlFlightReShopListSearchRequest {
	s.SelectedPassengers = v
	return s
}

func (s *IntlFlightReShopListSearchRequest) SetToken(v string) *IntlFlightReShopListSearchRequest {
	s.Token = &v
	return s
}

func (s *IntlFlightReShopListSearchRequest) Validate() error {
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

type IntlFlightReShopListSearchRequestSearchJourneys struct {
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2023-12-28
	DepDate         *string                                                           `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	SelectedFlights []*IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights `json:"selected_flights,omitempty" xml:"selected_flights,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopListSearchRequestSearchJourneys) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchRequestSearchJourneys) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) GetDepDate() *string {
	return s.DepDate
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) GetSelectedFlights() []*IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights {
	return s.SelectedFlights
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) SetArrCityCode(v string) *IntlFlightReShopListSearchRequestSearchJourneys {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) SetDepCityCode(v string) *IntlFlightReShopListSearchRequestSearchJourneys {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) SetDepDate(v string) *IntlFlightReShopListSearchRequestSearchJourneys {
	s.DepDate = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) SetSelectedFlights(v []*IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) *IntlFlightReShopListSearchRequestSearchJourneys {
	s.SelectedFlights = v
	return s
}

func (s *IntlFlightReShopListSearchRequestSearchJourneys) Validate() error {
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

type IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights struct {
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2023-09-01 07:10
	FlightTime *string `json:"flight_time,omitempty" xml:"flight_time,omitempty"`
	// example:
	//
	// KA5809
	MarketFlightNo *string `json:"market_flight_no,omitempty" xml:"market_flight_no,omitempty"`
}

func (s IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) GetFlightTime() *string {
	return s.FlightTime
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) GetMarketFlightNo() *string {
	return s.MarketFlightNo
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) SetArrCityCode(v string) *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) SetDepCityCode(v string) *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) SetFlightTime(v string) *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights {
	s.FlightTime = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) SetMarketFlightNo(v string) *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights {
	s.MarketFlightNo = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSearchJourneysSelectedFlights) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchRequestSelectedPassengers struct {
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// example:
	//
	// 1000001
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
}

func (s IntlFlightReShopListSearchRequestSelectedPassengers) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchRequestSelectedPassengers) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchRequestSelectedPassengers) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightReShopListSearchRequestSelectedPassengers) GetPassengerId() *int64 {
	return s.PassengerId
}

func (s *IntlFlightReShopListSearchRequestSelectedPassengers) SetFullName(v string) *IntlFlightReShopListSearchRequestSelectedPassengers {
	s.FullName = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSelectedPassengers) SetPassengerId(v int64) *IntlFlightReShopListSearchRequestSelectedPassengers {
	s.PassengerId = &v
	return s
}

func (s *IntlFlightReShopListSearchRequestSelectedPassengers) Validate() error {
	return dara.Validate(s)
}
