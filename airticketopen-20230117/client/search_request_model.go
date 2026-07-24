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
	// The number of adult passengers. Valid values: 1 to 9.
	//
	// example:
	//
	// 2
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// The journey array.
	//
	// This parameter is required.
	AirLegs []*SearchRequestAirLegs `json:"air_legs,omitempty" xml:"air_legs,omitempty" type:"Repeated"`
	// The cabin class. Valid values: ALL_CABIN: all cabin classes. Y: economy class. FC: first class and business class. S: premium economy class. YS: economy class and premium economy class. YSC: economy class, premium economy class, and business class.
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// The number of child passengers. Valid values: 0 to 9.
	//
	// example:
	//
	// 1
	Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
	// The number of infant passengers. Valid values: 0 to 9.
	//
	// example:
	//
	// 1
	Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
	// The search control options. This parameter is optional.
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
	// The list of three-letter codes of arrival airports.
	//
	// example:
	//
	// MFM
	ArrivalAirportList []*string `json:"arrival_airport_list,omitempty" xml:"arrival_airport_list,omitempty" type:"Repeated"`
	// The three-letter code of the arrival city.
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// The list of three-letter codes of departure airports.
	//
	// example:
	//
	// PVG
	DepartureAirportList []*string `json:"departure_airport_list,omitempty" xml:"departure_airport_list,omitempty" type:"Repeated"`
	// The three-letter code of the departure city.
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// The departure date (for example, yyyyMMdd).
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
	// The list of excluded airlines.
	//
	// example:
	//
	// 7C
	AirlineExcludedList []*string `json:"airline_excluded_list,omitempty" xml:"airline_excluded_list,omitempty" type:"Repeated"`
	// The list of preferred airlines.
	//
	// example:
	//
	// FD
	AirlinePreferList []*string `json:"airline_prefer_list,omitempty" xml:"airline_prefer_list,omitempty" type:"Repeated"`
	// The service quality identifier.
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
