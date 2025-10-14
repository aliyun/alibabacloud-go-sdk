// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectFlightLowestPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLowestPriceFlightInfoList(v []*CollectFlightLowestPriceRequestLowestPriceFlightInfoList) *CollectFlightLowestPriceRequest
	GetLowestPriceFlightInfoList() []*CollectFlightLowestPriceRequestLowestPriceFlightInfoList
}

type CollectFlightLowestPriceRequest struct {
	// This parameter is required.
	LowestPriceFlightInfoList []*CollectFlightLowestPriceRequestLowestPriceFlightInfoList `json:"lowest_price_flight_info_list,omitempty" xml:"lowest_price_flight_info_list,omitempty" type:"Repeated"`
}

func (s CollectFlightLowestPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s CollectFlightLowestPriceRequest) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceRequest) GetLowestPriceFlightInfoList() []*CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	return s.LowestPriceFlightInfoList
}

func (s *CollectFlightLowestPriceRequest) SetLowestPriceFlightInfoList(v []*CollectFlightLowestPriceRequestLowestPriceFlightInfoList) *CollectFlightLowestPriceRequest {
	s.LowestPriceFlightInfoList = v
	return s
}

func (s *CollectFlightLowestPriceRequest) Validate() error {
	if s.LowestPriceFlightInfoList != nil {
		for _, item := range s.LowestPriceFlightInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CollectFlightLowestPriceRequestLowestPriceFlightInfoList struct {
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-11-11
	DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CA123,CA456
	DepartureFlightNumber *string `json:"departure_flight_number,omitempty" xml:"departure_flight_number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100.0
	MarketTotalPrice *float64 `json:"market_total_price,omitempty" xml:"market_total_price,omitempty"`
	// example:
	//
	// 123456789dacd
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 2024-11-11
	ReturnDate *string `json:"return_date,omitempty" xml:"return_date,omitempty"`
	// example:
	//
	// CA123,CA456
	ReturnFlightNumber *string `json:"return_flight_number,omitempty" xml:"return_flight_number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100.1
	SuezTotalPrice *float64 `json:"suez_total_price,omitempty" xml:"suez_total_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s CollectFlightLowestPriceRequestLowestPriceFlightInfoList) String() string {
	return dara.Prettify(s)
}

func (s CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetDepartureDate() *string {
	return s.DepartureDate
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetDepartureFlightNumber() *string {
	return s.DepartureFlightNumber
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetMarketTotalPrice() *float64 {
	return s.MarketTotalPrice
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetRequestId() *string {
	return s.RequestId
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetReturnDate() *string {
	return s.ReturnDate
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetReturnFlightNumber() *string {
	return s.ReturnFlightNumber
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetSolutionId() *string {
	return s.SolutionId
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetSuezTotalPrice() *float64 {
	return s.SuezTotalPrice
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) GetTripType() *int32 {
	return s.TripType
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetArrivalCity(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.ArrivalCity = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetDepartureCity(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.DepartureCity = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetDepartureDate(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.DepartureDate = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetDepartureFlightNumber(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.DepartureFlightNumber = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetMarketTotalPrice(v float64) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.MarketTotalPrice = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetRequestId(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.RequestId = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetReturnDate(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.ReturnDate = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetReturnFlightNumber(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.ReturnFlightNumber = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetSolutionId(v string) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.SolutionId = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetSuezTotalPrice(v float64) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.SuezTotalPrice = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) SetTripType(v int32) *CollectFlightLowestPriceRequestLowestPriceFlightInfoList {
	s.TripType = &v
	return s
}

func (s *CollectFlightLowestPriceRequestLowestPriceFlightInfoList) Validate() error {
	return dara.Validate(s)
}
