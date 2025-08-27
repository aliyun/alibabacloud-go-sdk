// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirlineCode(v string) *FlightListingSearchRequest
	GetAirlineCode() *string
	SetArrCityCode(v string) *FlightListingSearchRequest
	GetArrCityCode() *string
	SetCabinClass(v string) *FlightListingSearchRequest
	GetCabinClass() *string
	SetDepCityCode(v string) *FlightListingSearchRequest
	GetDepCityCode() *string
	SetDepDate(v string) *FlightListingSearchRequest
	GetDepDate() *string
}

type FlightListingSearchRequest struct {
	// example:
	//
	// MU
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
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
	// 2023-02-26
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
}

func (s FlightListingSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchRequest) GoString() string {
	return s.String()
}

func (s *FlightListingSearchRequest) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightListingSearchRequest) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightListingSearchRequest) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightListingSearchRequest) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightListingSearchRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightListingSearchRequest) SetAirlineCode(v string) *FlightListingSearchRequest {
	s.AirlineCode = &v
	return s
}

func (s *FlightListingSearchRequest) SetArrCityCode(v string) *FlightListingSearchRequest {
	s.ArrCityCode = &v
	return s
}

func (s *FlightListingSearchRequest) SetCabinClass(v string) *FlightListingSearchRequest {
	s.CabinClass = &v
	return s
}

func (s *FlightListingSearchRequest) SetDepCityCode(v string) *FlightListingSearchRequest {
	s.DepCityCode = &v
	return s
}

func (s *FlightListingSearchRequest) SetDepDate(v string) *FlightListingSearchRequest {
	s.DepDate = &v
	return s
}

func (s *FlightListingSearchRequest) Validate() error {
	return dara.Validate(s)
}
