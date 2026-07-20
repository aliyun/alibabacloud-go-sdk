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
	// The business travel order ID. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1017089206331988332
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// The external order ID.
	//
	// example:
	//
	// JP2024071800000002
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// Specifies whether to enable external polling. When enabled, the external frontend controls polling, which reduces the response time (RT) of a single search. Default value: false.
	//
	// example:
	//
	// true
	OutWheelSearch *bool `json:"out_wheel_search,omitempty" xml:"out_wheel_search,omitempty"`
	// The rebooking group key returned by the consultation operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// edcac4f4c79d40ccb141ddb6da567e65
	PassengerJourneyGroupKey *string `json:"passenger_journey_group_key,omitempty" xml:"passenger_journey_group_key,omitempty"`
	// The rebooking reason code.
	//
	// example:
	//
	// 0
	ReShopReasonCode *string `json:"re_shop_reason_code,omitempty" xml:"re_shop_reason_code,omitempty"`
	// The search journeys. Only pass the journeys to be rebooked.
	//
	// This parameter is required.
	SearchJourneys []*IntlFlightReShopOtaSearchRequestSearchJourneys `json:"search_journeys,omitempty" xml:"search_journeys,omitempty" type:"Repeated"`
	// The list of selected passengers for rebooking.
	//
	// This parameter is required.
	SelectedPassengers []*IntlFlightReShopOtaSearchRequestSelectedPassengers `json:"selected_passengers,omitempty" xml:"selected_passengers,omitempty" type:"Repeated"`
	// The query record token used for external polling.
	//
	// example:
	//
	// 9960b412-cc05-4d10-b570-93372d816807
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
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
	// The three-letter code of the arrival city.
	//
	// example:
	//
	// MEL
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// The three-letter code of the departure city.
	//
	// example:
	//
	// HKG
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// The departure date. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2025-12-28
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// The selected flight information for rebooking.
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
	// The three-letter code of the arrival city.
	//
	// example:
	//
	// MEL
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// The three-letter code of the departure city.
	//
	// example:
	//
	// HKG
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// The departure time. Format: yyyy-MM-dd HH:mm.
	//
	// example:
	//
	// 2025-12-28 12:00
	FlightTime *string `json:"flight_time,omitempty" xml:"flight_time,omitempty"`
	// The marketing flight number.
	//
	// example:
	//
	// HO3925
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
	// The full name of the passenger.
	//
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// The passenger ID.
	//
	// example:
	//
	// 1000001
	PassengerId *int64 `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
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
