// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransitVisaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlightSegmentParamListShrink(v string) *TransitVisaShrinkRequest
	GetFlightSegmentParamListShrink() *string
}

type TransitVisaShrinkRequest struct {
	FlightSegmentParamListShrink *string `json:"flight_segment_param_list,omitempty" xml:"flight_segment_param_list,omitempty"`
}

func (s TransitVisaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TransitVisaShrinkRequest) GoString() string {
	return s.String()
}

func (s *TransitVisaShrinkRequest) GetFlightSegmentParamListShrink() *string {
	return s.FlightSegmentParamListShrink
}

func (s *TransitVisaShrinkRequest) SetFlightSegmentParamListShrink(v string) *TransitVisaShrinkRequest {
	s.FlightSegmentParamListShrink = &v
	return s
}

func (s *TransitVisaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
