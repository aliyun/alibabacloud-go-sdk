// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirlineCode(v string) *FlightOtaSearchRequest
	GetAirlineCode() *string
	SetArrCityCode(v string) *FlightOtaSearchRequest
	GetArrCityCode() *string
	SetCabinClass(v string) *FlightOtaSearchRequest
	GetCabinClass() *string
	SetCarrierFlightNo(v string) *FlightOtaSearchRequest
	GetCarrierFlightNo() *string
	SetDepCityCode(v string) *FlightOtaSearchRequest
	GetDepCityCode() *string
	SetDepDate(v string) *FlightOtaSearchRequest
	GetDepDate() *string
	SetFlightNo(v string) *FlightOtaSearchRequest
	GetFlightNo() *string
}

type FlightOtaSearchRequest struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// CA2324
	CarrierFlightNo *string `json:"carrier_flight_no,omitempty" xml:"carrier_flight_no,omitempty"`
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
	// 2023-08-15 19:30:00
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CA2323
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s FlightOtaSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchRequest) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchRequest) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOtaSearchRequest) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOtaSearchRequest) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightOtaSearchRequest) GetCarrierFlightNo() *string {
	return s.CarrierFlightNo
}

func (s *FlightOtaSearchRequest) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOtaSearchRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightOtaSearchRequest) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOtaSearchRequest) SetAirlineCode(v string) *FlightOtaSearchRequest {
	s.AirlineCode = &v
	return s
}

func (s *FlightOtaSearchRequest) SetArrCityCode(v string) *FlightOtaSearchRequest {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOtaSearchRequest) SetCabinClass(v string) *FlightOtaSearchRequest {
	s.CabinClass = &v
	return s
}

func (s *FlightOtaSearchRequest) SetCarrierFlightNo(v string) *FlightOtaSearchRequest {
	s.CarrierFlightNo = &v
	return s
}

func (s *FlightOtaSearchRequest) SetDepCityCode(v string) *FlightOtaSearchRequest {
	s.DepCityCode = &v
	return s
}

func (s *FlightOtaSearchRequest) SetDepDate(v string) *FlightOtaSearchRequest {
	s.DepDate = &v
	return s
}

func (s *FlightOtaSearchRequest) SetFlightNo(v string) *FlightOtaSearchRequest {
	s.FlightNo = &v
	return s
}

func (s *FlightOtaSearchRequest) Validate() error {
	return dara.Validate(s)
}
