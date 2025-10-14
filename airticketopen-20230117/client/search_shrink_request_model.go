// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdults(v int32) *SearchShrinkRequest
	GetAdults() *int32
	SetAirLegsShrink(v string) *SearchShrinkRequest
	GetAirLegsShrink() *string
	SetCabinClass(v string) *SearchShrinkRequest
	GetCabinClass() *string
	SetChildren(v int32) *SearchShrinkRequest
	GetChildren() *int32
	SetInfants(v int32) *SearchShrinkRequest
	GetInfants() *int32
	SetSearchControlOptionsShrink(v string) *SearchShrinkRequest
	GetSearchControlOptionsShrink() *string
}

type SearchShrinkRequest struct {
	// adult passenger amount 1-9
	//
	// example:
	//
	// 2
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// itinerary list
	//
	// This parameter is required.
	AirLegsShrink *string `json:"air_legs,omitempty" xml:"air_legs,omitempty"`
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
	SearchControlOptionsShrink *string `json:"search_control_options,omitempty" xml:"search_control_options,omitempty"`
}

func (s SearchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchShrinkRequest) GetAdults() *int32 {
	return s.Adults
}

func (s *SearchShrinkRequest) GetAirLegsShrink() *string {
	return s.AirLegsShrink
}

func (s *SearchShrinkRequest) GetCabinClass() *string {
	return s.CabinClass
}

func (s *SearchShrinkRequest) GetChildren() *int32 {
	return s.Children
}

func (s *SearchShrinkRequest) GetInfants() *int32 {
	return s.Infants
}

func (s *SearchShrinkRequest) GetSearchControlOptionsShrink() *string {
	return s.SearchControlOptionsShrink
}

func (s *SearchShrinkRequest) SetAdults(v int32) *SearchShrinkRequest {
	s.Adults = &v
	return s
}

func (s *SearchShrinkRequest) SetAirLegsShrink(v string) *SearchShrinkRequest {
	s.AirLegsShrink = &v
	return s
}

func (s *SearchShrinkRequest) SetCabinClass(v string) *SearchShrinkRequest {
	s.CabinClass = &v
	return s
}

func (s *SearchShrinkRequest) SetChildren(v int32) *SearchShrinkRequest {
	s.Children = &v
	return s
}

func (s *SearchShrinkRequest) SetInfants(v int32) *SearchShrinkRequest {
	s.Infants = &v
	return s
}

func (s *SearchShrinkRequest) SetSearchControlOptionsShrink(v string) *SearchShrinkRequest {
	s.SearchControlOptionsShrink = &v
	return s
}

func (s *SearchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
