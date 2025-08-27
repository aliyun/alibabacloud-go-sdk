// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelSuggestV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelSuggestV2Request
	GetBtripUserId() *string
	SetCheckIn(v string) *HotelSuggestV2Request
	GetCheckIn() *string
	SetCheckOut(v string) *HotelSuggestV2Request
	GetCheckOut() *string
	SetCityCode(v string) *HotelSuggestV2Request
	GetCityCode() *string
	SetKeyword(v string) *HotelSuggestV2Request
	GetKeyword() *string
	SetSearchType(v int32) *HotelSuggestV2Request
	GetSearchType() *int32
}

type HotelSuggestV2Request struct {
	// example:
	//
	// 1000
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// 2024-05-15
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 2024-06-04
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// example:
	//
	// 330100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	Keyword  *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SearchType *int32 `json:"search_type,omitempty" xml:"search_type,omitempty"`
}

func (s HotelSuggestV2Request) String() string {
	return dara.Prettify(s)
}

func (s HotelSuggestV2Request) GoString() string {
	return s.String()
}

func (s *HotelSuggestV2Request) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelSuggestV2Request) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelSuggestV2Request) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelSuggestV2Request) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelSuggestV2Request) GetKeyword() *string {
	return s.Keyword
}

func (s *HotelSuggestV2Request) GetSearchType() *int32 {
	return s.SearchType
}

func (s *HotelSuggestV2Request) SetBtripUserId(v string) *HotelSuggestV2Request {
	s.BtripUserId = &v
	return s
}

func (s *HotelSuggestV2Request) SetCheckIn(v string) *HotelSuggestV2Request {
	s.CheckIn = &v
	return s
}

func (s *HotelSuggestV2Request) SetCheckOut(v string) *HotelSuggestV2Request {
	s.CheckOut = &v
	return s
}

func (s *HotelSuggestV2Request) SetCityCode(v string) *HotelSuggestV2Request {
	s.CityCode = &v
	return s
}

func (s *HotelSuggestV2Request) SetKeyword(v string) *HotelSuggestV2Request {
	s.Keyword = &v
	return s
}

func (s *HotelSuggestV2Request) SetSearchType(v int32) *HotelSuggestV2Request {
	s.SearchType = &v
	return s
}

func (s *HotelSuggestV2Request) Validate() error {
	return dara.Validate(s)
}
