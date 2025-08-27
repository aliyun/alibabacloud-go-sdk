// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelStaticInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelIdsShrink(v string) *HotelStaticInfoShrinkRequest
	GetHotelIdsShrink() *string
}

type HotelStaticInfoShrinkRequest struct {
	// This parameter is required.
	HotelIdsShrink *string `json:"hotel_ids,omitempty" xml:"hotel_ids,omitempty"`
}

func (s HotelStaticInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoShrinkRequest) GetHotelIdsShrink() *string {
	return s.HotelIdsShrink
}

func (s *HotelStaticInfoShrinkRequest) SetHotelIdsShrink(v string) *HotelStaticInfoShrinkRequest {
	s.HotelIdsShrink = &v
	return s
}

func (s *HotelStaticInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
