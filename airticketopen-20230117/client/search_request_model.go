// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdults(v int32) *SearchRequest
	GetAdults() *int32
	SetAirLegs(v []*SearchRequestAirLegs) *SearchRequest
	GetAirLegs() []*SearchRequestAirLegs
	SetCabinClass(v string) *SearchRequest
	GetCabinClass() *string
	SetChildren(v int32) *SearchRequest
	GetChildren() *int32
	SetInfants(v int32) *SearchRequest
	GetInfants() *int32
	SetSearchControlOptions(v *SearchRequestSearchControlOptions) *SearchRequest
	GetSearchControlOptions() *SearchRequestSearchControlOptions
}

type SearchRequest struct {
	// adult passenger amount 1-9
	//
	// example:
	//
	// 2
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// itinerary list
	//
	// This parameter is required.
	AirLegs []*SearchRequestAirLegs `json:"air_legs,omitempty" xml:"air_legs,omitempty" type:"Repeated"`
	// cabin class
	//
	// 1. **ALL_CABIN*	- : all cabin class
	//
	// 2. **Y*	- : economy class
	//
	// 3. **FC*	- : first class and business class
	//
	// 4. **S*	- : premium economy class
	//
	// 5. **YS*	- : economy class and premium economy class
	//
	// 6. **YSC*	- : economy class, premium economy class and business class
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// child passenger amount 0-9
	//
	// example:
	//
	// 1
	Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
	// infant passenger amount 0-9
	//
	// example:
	//
	// 1
	Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
	// search controls
	SearchControlOptions *SearchRequestSearchControlOptions `json:"search_control_options,omitempty" xml:"search_control_options,omitempty" type:"Struct"`
}

func (s SearchRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchRequest) GoString() string {
	return s.String()
}

func (s *SearchRequest) GetAdults() *int32 {
	return s.Adults
}

func (s *SearchRequest) GetAirLegs() []*SearchRequestAirLegs {
	return s.AirLegs
}

func (s *SearchRequest) GetCabinClass() *string {
	return s.CabinClass
}

func (s *SearchRequest) GetChildren() *int32 {
	return s.Children
}

func (s *SearchRequest) GetInfants() *int32 {
	return s.Infants
}

func (s *SearchRequest) GetSearchControlOptions() *SearchRequestSearchControlOptions {
	return s.SearchControlOptions
}

func (s *SearchRequest) SetAdults(v int32) *SearchRequest {
	s.Adults = &v
	return s
}

func (s *SearchRequest) SetAirLegs(v []*SearchRequestAirLegs) *SearchRequest {
	s.AirLegs = v
	return s
}

func (s *SearchRequest) SetCabinClass(v string) *SearchRequest {
	s.CabinClass = &v
	return s
}

func (s *SearchRequest) SetChildren(v int32) *SearchRequest {
	s.Children = &v
	return s
}

func (s *SearchRequest) SetInfants(v int32) *SearchRequest {
	s.Infants = &v
	return s
}

func (s *SearchRequest) SetSearchControlOptions(v *SearchRequestSearchControlOptions) *SearchRequest {
	s.SearchControlOptions = v
	return s
}

func (s *SearchRequest) Validate() error {
	if s.AirLegs != nil {
		for _, item := range s.AirLegs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchControlOptions != nil {
		if err := s.SearchControlOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchRequestAirLegs struct {
	// arrival airport [IATA airport code] list
	//
	// example:
	//
	// MFM
	ArrivalAirportList []*string `json:"arrival_airport_list,omitempty" xml:"arrival_airport_list,omitempty" type:"Repeated"`
	// arrival city code
	//
	// This parameter is required.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// departure airport [IATA airport code] list
	//
	// example:
	//
	// PVG
	DepartureAirportList []*string `json:"departure_airport_list,omitempty" xml:"departure_airport_list,omitempty" type:"Repeated"`
	// departure city code
	//
	// This parameter is required.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure date (eg: yyyyMMdd)
	//
	// This parameter is required.
	//
	// example:
	//
	// 20230310
	DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
}

func (s SearchRequestAirLegs) String() string {
	return dara.Prettify(s)
}

func (s SearchRequestAirLegs) GoString() string {
	return s.String()
}

func (s *SearchRequestAirLegs) GetArrivalAirportList() []*string {
	return s.ArrivalAirportList
}

func (s *SearchRequestAirLegs) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *SearchRequestAirLegs) GetDepartureAirportList() []*string {
	return s.DepartureAirportList
}

func (s *SearchRequestAirLegs) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *SearchRequestAirLegs) GetDepartureDate() *string {
	return s.DepartureDate
}

func (s *SearchRequestAirLegs) SetArrivalAirportList(v []*string) *SearchRequestAirLegs {
	s.ArrivalAirportList = v
	return s
}

func (s *SearchRequestAirLegs) SetArrivalCity(v string) *SearchRequestAirLegs {
	s.ArrivalCity = &v
	return s
}

func (s *SearchRequestAirLegs) SetDepartureAirportList(v []*string) *SearchRequestAirLegs {
	s.DepartureAirportList = v
	return s
}

func (s *SearchRequestAirLegs) SetDepartureCity(v string) *SearchRequestAirLegs {
	s.DepartureCity = &v
	return s
}

func (s *SearchRequestAirLegs) SetDepartureDate(v string) *SearchRequestAirLegs {
	s.DepartureDate = &v
	return s
}

func (s *SearchRequestAirLegs) Validate() error {
	return dara.Validate(s)
}

type SearchRequestSearchControlOptions struct {
	// excluded airlines list
	AirlineExcludedList []*string `json:"airline_excluded_list,omitempty" xml:"airline_excluded_list,omitempty" type:"Repeated"`
	// preferred airlines list
	AirlinePreferList []*string `json:"airline_prefer_list,omitempty" xml:"airline_prefer_list,omitempty" type:"Repeated"`
	// service quality
	//
	// example:
	//
	// A1
	ServiceQuality *string `json:"service_quality,omitempty" xml:"service_quality,omitempty"`
}

func (s SearchRequestSearchControlOptions) String() string {
	return dara.Prettify(s)
}

func (s SearchRequestSearchControlOptions) GoString() string {
	return s.String()
}

func (s *SearchRequestSearchControlOptions) GetAirlineExcludedList() []*string {
	return s.AirlineExcludedList
}

func (s *SearchRequestSearchControlOptions) GetAirlinePreferList() []*string {
	return s.AirlinePreferList
}

func (s *SearchRequestSearchControlOptions) GetServiceQuality() *string {
	return s.ServiceQuality
}

func (s *SearchRequestSearchControlOptions) SetAirlineExcludedList(v []*string) *SearchRequestSearchControlOptions {
	s.AirlineExcludedList = v
	return s
}

func (s *SearchRequestSearchControlOptions) SetAirlinePreferList(v []*string) *SearchRequestSearchControlOptions {
	s.AirlinePreferList = v
	return s
}

func (s *SearchRequestSearchControlOptions) SetServiceQuality(v string) *SearchRequestSearchControlOptions {
	s.ServiceQuality = &v
	return s
}

func (s *SearchRequestSearchControlOptions) Validate() error {
	return dara.Validate(s)
}
