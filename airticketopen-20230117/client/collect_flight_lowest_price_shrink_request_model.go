// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectFlightLowestPriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLowestPriceFlightInfoListShrink(v string) *CollectFlightLowestPriceShrinkRequest
	GetLowestPriceFlightInfoListShrink() *string
}

type CollectFlightLowestPriceShrinkRequest struct {
	// This parameter is required.
	LowestPriceFlightInfoListShrink *string `json:"lowest_price_flight_info_list,omitempty" xml:"lowest_price_flight_info_list,omitempty"`
}

func (s CollectFlightLowestPriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CollectFlightLowestPriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CollectFlightLowestPriceShrinkRequest) GetLowestPriceFlightInfoListShrink() *string {
	return s.LowestPriceFlightInfoListShrink
}

func (s *CollectFlightLowestPriceShrinkRequest) SetLowestPriceFlightInfoListShrink(v string) *CollectFlightLowestPriceShrinkRequest {
	s.LowestPriceFlightInfoListShrink = &v
	return s
}

func (s *CollectFlightLowestPriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
