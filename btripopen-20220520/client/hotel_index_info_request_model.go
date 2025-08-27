// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelIndexInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCityCode(v string) *HotelIndexInfoRequest
	GetCityCode() *string
	SetHotelStatus(v int32) *HotelIndexInfoRequest
	GetHotelStatus() *int32
	SetPageSize(v int32) *HotelIndexInfoRequest
	GetPageSize() *int32
	SetPageToken(v string) *HotelIndexInfoRequest
	GetPageToken() *string
}

type HotelIndexInfoRequest struct {
	// example:
	//
	// 330000
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 0
	HotelStatus *int32 `json:"hotel_status,omitempty" xml:"hotel_status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 1038882
	PageToken *string `json:"page_token,omitempty" xml:"page_token,omitempty"`
}

func (s HotelIndexInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelIndexInfoRequest) GoString() string {
	return s.String()
}

func (s *HotelIndexInfoRequest) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelIndexInfoRequest) GetHotelStatus() *int32 {
	return s.HotelStatus
}

func (s *HotelIndexInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotelIndexInfoRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *HotelIndexInfoRequest) SetCityCode(v string) *HotelIndexInfoRequest {
	s.CityCode = &v
	return s
}

func (s *HotelIndexInfoRequest) SetHotelStatus(v int32) *HotelIndexInfoRequest {
	s.HotelStatus = &v
	return s
}

func (s *HotelIndexInfoRequest) SetPageSize(v int32) *HotelIndexInfoRequest {
	s.PageSize = &v
	return s
}

func (s *HotelIndexInfoRequest) SetPageToken(v string) *HotelIndexInfoRequest {
	s.PageToken = &v
	return s
}

func (s *HotelIndexInfoRequest) Validate() error {
	return dara.Validate(s)
}
