// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLuggageDirectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlightSegmentParamListShrink(v string) *LuggageDirectShrinkRequest
	GetFlightSegmentParamListShrink() *string
}

type LuggageDirectShrinkRequest struct {
	// The list of flight segments that constitute an itinerary. Maximum size: 2.
	FlightSegmentParamListShrink *string `json:"flight_segment_param_list,omitempty" xml:"flight_segment_param_list,omitempty"`
}

func (s LuggageDirectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s LuggageDirectShrinkRequest) GoString() string {
	return s.String()
}

func (s *LuggageDirectShrinkRequest) GetFlightSegmentParamListShrink() *string {
	return s.FlightSegmentParamListShrink
}

func (s *LuggageDirectShrinkRequest) SetFlightSegmentParamListShrink(v string) *LuggageDirectShrinkRequest {
	s.FlightSegmentParamListShrink = &v
	return s
}

func (s *LuggageDirectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
