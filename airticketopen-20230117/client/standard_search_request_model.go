// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStandardSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdults(v int32) *StandardSearchRequest
	GetAdults() *int32
	SetAirLegs(v []*StandardSearchRequestAirLegs) *StandardSearchRequest
	GetAirLegs() []*StandardSearchRequestAirLegs
	SetCabinClass(v string) *StandardSearchRequest
	GetCabinClass() *string
	SetChildren(v int32) *StandardSearchRequest
	GetChildren() *int32
	SetInfants(v int32) *StandardSearchRequest
	GetInfants() *int32
	SetSearchControlOptions(v *StandardSearchRequestSearchControlOptions) *StandardSearchRequest
	GetSearchControlOptions() *StandardSearchRequestSearchControlOptions
}

type StandardSearchRequest struct {
	// example:
	//
	// 2
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// This parameter is required.
	AirLegs []*StandardSearchRequestAirLegs `json:"air_legs,omitempty" xml:"air_legs,omitempty" type:"Repeated"`
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// 1
	Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
	// example:
	//
	// 1
	Infants              *int32                                     `json:"infants,omitempty" xml:"infants,omitempty"`
	SearchControlOptions *StandardSearchRequestSearchControlOptions `json:"search_control_options,omitempty" xml:"search_control_options,omitempty" type:"Struct"`
}

func (s StandardSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchRequest) GoString() string {
	return s.String()
}

func (s *StandardSearchRequest) GetAdults() *int32 {
	return s.Adults
}

func (s *StandardSearchRequest) GetAirLegs() []*StandardSearchRequestAirLegs {
	return s.AirLegs
}

func (s *StandardSearchRequest) GetCabinClass() *string {
	return s.CabinClass
}

func (s *StandardSearchRequest) GetChildren() *int32 {
	return s.Children
}

func (s *StandardSearchRequest) GetInfants() *int32 {
	return s.Infants
}

func (s *StandardSearchRequest) GetSearchControlOptions() *StandardSearchRequestSearchControlOptions {
	return s.SearchControlOptions
}

func (s *StandardSearchRequest) SetAdults(v int32) *StandardSearchRequest {
	s.Adults = &v
	return s
}

func (s *StandardSearchRequest) SetAirLegs(v []*StandardSearchRequestAirLegs) *StandardSearchRequest {
	s.AirLegs = v
	return s
}

func (s *StandardSearchRequest) SetCabinClass(v string) *StandardSearchRequest {
	s.CabinClass = &v
	return s
}

func (s *StandardSearchRequest) SetChildren(v int32) *StandardSearchRequest {
	s.Children = &v
	return s
}

func (s *StandardSearchRequest) SetInfants(v int32) *StandardSearchRequest {
	s.Infants = &v
	return s
}

func (s *StandardSearchRequest) SetSearchControlOptions(v *StandardSearchRequestSearchControlOptions) *StandardSearchRequest {
	s.SearchControlOptions = v
	return s
}

func (s *StandardSearchRequest) Validate() error {
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

type StandardSearchRequestAirLegs struct {
	// example:
	//
	// MFM
	ArrivalAirportList []*string `json:"arrival_airport_list,omitempty" xml:"arrival_airport_list,omitempty" type:"Repeated"`
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// PVG
	DepartureAirportList []*string `json:"departure_airport_list,omitempty" xml:"departure_airport_list,omitempty" type:"Repeated"`
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20230320
	DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
}

func (s StandardSearchRequestAirLegs) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchRequestAirLegs) GoString() string {
	return s.String()
}

func (s *StandardSearchRequestAirLegs) GetArrivalAirportList() []*string {
	return s.ArrivalAirportList
}

func (s *StandardSearchRequestAirLegs) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *StandardSearchRequestAirLegs) GetDepartureAirportList() []*string {
	return s.DepartureAirportList
}

func (s *StandardSearchRequestAirLegs) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *StandardSearchRequestAirLegs) GetDepartureDate() *string {
	return s.DepartureDate
}

func (s *StandardSearchRequestAirLegs) SetArrivalAirportList(v []*string) *StandardSearchRequestAirLegs {
	s.ArrivalAirportList = v
	return s
}

func (s *StandardSearchRequestAirLegs) SetArrivalCity(v string) *StandardSearchRequestAirLegs {
	s.ArrivalCity = &v
	return s
}

func (s *StandardSearchRequestAirLegs) SetDepartureAirportList(v []*string) *StandardSearchRequestAirLegs {
	s.DepartureAirportList = v
	return s
}

func (s *StandardSearchRequestAirLegs) SetDepartureCity(v string) *StandardSearchRequestAirLegs {
	s.DepartureCity = &v
	return s
}

func (s *StandardSearchRequestAirLegs) SetDepartureDate(v string) *StandardSearchRequestAirLegs {
	s.DepartureDate = &v
	return s
}

func (s *StandardSearchRequestAirLegs) Validate() error {
	return dara.Validate(s)
}

type StandardSearchRequestSearchControlOptions struct {
	// example:
	//
	// 7C
	AirlineExcludedList []*string `json:"airline_excluded_list,omitempty" xml:"airline_excluded_list,omitempty" type:"Repeated"`
	// example:
	//
	// FD
	AirlinePreferList []*string `json:"airline_prefer_list,omitempty" xml:"airline_prefer_list,omitempty" type:"Repeated"`
	// example:
	//
	// A1
	ServiceQuality *string `json:"service_quality,omitempty" xml:"service_quality,omitempty"`
}

func (s StandardSearchRequestSearchControlOptions) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchRequestSearchControlOptions) GoString() string {
	return s.String()
}

func (s *StandardSearchRequestSearchControlOptions) GetAirlineExcludedList() []*string {
	return s.AirlineExcludedList
}

func (s *StandardSearchRequestSearchControlOptions) GetAirlinePreferList() []*string {
	return s.AirlinePreferList
}

func (s *StandardSearchRequestSearchControlOptions) GetServiceQuality() *string {
	return s.ServiceQuality
}

func (s *StandardSearchRequestSearchControlOptions) SetAirlineExcludedList(v []*string) *StandardSearchRequestSearchControlOptions {
	s.AirlineExcludedList = v
	return s
}

func (s *StandardSearchRequestSearchControlOptions) SetAirlinePreferList(v []*string) *StandardSearchRequestSearchControlOptions {
	s.AirlinePreferList = v
	return s
}

func (s *StandardSearchRequestSearchControlOptions) SetServiceQuality(v string) *StandardSearchRequestSearchControlOptions {
	s.ServiceQuality = &v
	return s
}

func (s *StandardSearchRequestSearchControlOptions) Validate() error {
	return dara.Validate(s)
}
