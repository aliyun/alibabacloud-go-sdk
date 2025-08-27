// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelStaticInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotelIds(v []*string) *HotelStaticInfoRequest
	GetHotelIds() []*string
}

type HotelStaticInfoRequest struct {
	// This parameter is required.
	HotelIds []*string `json:"hotel_ids,omitempty" xml:"hotel_ids,omitempty" type:"Repeated"`
}

func (s HotelStaticInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelStaticInfoRequest) GoString() string {
	return s.String()
}

func (s *HotelStaticInfoRequest) GetHotelIds() []*string {
	return s.HotelIds
}

func (s *HotelStaticInfoRequest) SetHotelIds(v []*string) *HotelStaticInfoRequest {
	s.HotelIds = v
	return s
}

func (s *HotelStaticInfoRequest) Validate() error {
	return dara.Validate(s)
}
