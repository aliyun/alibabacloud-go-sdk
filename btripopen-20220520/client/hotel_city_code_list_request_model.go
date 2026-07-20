// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelCityCodeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCountryCode(v string) *HotelCityCodeListRequest
	GetCountryCode() *string
}

type HotelCityCodeListRequest struct {
	// example:
	//
	// 1
	CountryCode *string `json:"country_code,omitempty" xml:"country_code,omitempty"`
}

func (s HotelCityCodeListRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelCityCodeListRequest) GoString() string {
	return s.String()
}

func (s *HotelCityCodeListRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *HotelCityCodeListRequest) SetCountryCode(v string) *HotelCityCodeListRequest {
	s.CountryCode = &v
	return s
}

func (s *HotelCityCodeListRequest) Validate() error {
	return dara.Validate(s)
}
