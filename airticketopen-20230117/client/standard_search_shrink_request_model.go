// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStandardSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdults(v int32) *StandardSearchShrinkRequest
	GetAdults() *int32
	SetAirLegsShrink(v string) *StandardSearchShrinkRequest
	GetAirLegsShrink() *string
	SetCabinClass(v string) *StandardSearchShrinkRequest
	GetCabinClass() *string
	SetChildren(v int32) *StandardSearchShrinkRequest
	GetChildren() *int32
	SetInfants(v int32) *StandardSearchShrinkRequest
	GetInfants() *int32
	SetSearchControlOptionsShrink(v string) *StandardSearchShrinkRequest
	GetSearchControlOptionsShrink() *string
}

type StandardSearchShrinkRequest struct {
	// Number of adult passengers, range 1-9
	//
	// example:
	//
	// 2
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// Journey array. At least one of departure_city and departure_airport_list must be non-empty; when departure_airport_list has values, they must belong to the same city. At least one of arrival_city and arrival_airport_list must be non-empty; when arrival_airport_list has values, they must belong to the same city.
	//
	// This parameter is required.
	AirLegsShrink *string `json:"air_legs,omitempty" xml:"air_legs,omitempty"`
	// Defaults to ALL_CABIN if not specified. Cabin class ALL_CABIN: All cabin classes; Y: Economy class; FC: First class and Business class; S: Premium Economy class; YS: Economy class and Premium Economy class; YSC: Economy class, Premium Economy class, and Business class;
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// Number of child passengers, range 0-9
	//
	// example:
	//
	// 1
	Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
	// Number of infant passengers, range 0-9
	//
	// example:
	//
	// 1
	Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
	// Search control options, optional
	SearchControlOptionsShrink *string `json:"search_control_options,omitempty" xml:"search_control_options,omitempty"`
}

func (s StandardSearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *StandardSearchShrinkRequest) GetAdults() *int32 {
	return s.Adults
}

func (s *StandardSearchShrinkRequest) GetAirLegsShrink() *string {
	return s.AirLegsShrink
}

func (s *StandardSearchShrinkRequest) GetCabinClass() *string {
	return s.CabinClass
}

func (s *StandardSearchShrinkRequest) GetChildren() *int32 {
	return s.Children
}

func (s *StandardSearchShrinkRequest) GetInfants() *int32 {
	return s.Infants
}

func (s *StandardSearchShrinkRequest) GetSearchControlOptionsShrink() *string {
	return s.SearchControlOptionsShrink
}

func (s *StandardSearchShrinkRequest) SetAdults(v int32) *StandardSearchShrinkRequest {
	s.Adults = &v
	return s
}

func (s *StandardSearchShrinkRequest) SetAirLegsShrink(v string) *StandardSearchShrinkRequest {
	s.AirLegsShrink = &v
	return s
}

func (s *StandardSearchShrinkRequest) SetCabinClass(v string) *StandardSearchShrinkRequest {
	s.CabinClass = &v
	return s
}

func (s *StandardSearchShrinkRequest) SetChildren(v int32) *StandardSearchShrinkRequest {
	s.Children = &v
	return s
}

func (s *StandardSearchShrinkRequest) SetInfants(v int32) *StandardSearchShrinkRequest {
	s.Infants = &v
	return s
}

func (s *StandardSearchShrinkRequest) SetSearchControlOptionsShrink(v string) *StandardSearchShrinkRequest {
	s.SearchControlOptionsShrink = &v
	return s
}

func (s *StandardSearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
