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
	// example:
	//
	// 2
	Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
	// This parameter is required.
	AirLegsShrink *string `json:"air_legs,omitempty" xml:"air_legs,omitempty"`
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
	Infants                    *int32  `json:"infants,omitempty" xml:"infants,omitempty"`
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
