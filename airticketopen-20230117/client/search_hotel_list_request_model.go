// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchHotelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *SearchHotelListRequest
	GetAccountNo() *int64
	SetCityCode(v string) *SearchHotelListRequest
	GetCityCode() *string
	SetPageNo(v int32) *SearchHotelListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchHotelListRequest
	GetPageSize() *int32
	SetTracerId(v string) *SearchHotelListRequest
	GetTracerId() *string
}

type SearchHotelListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// beijing
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s SearchHotelListRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchHotelListRequest) GoString() string {
	return s.String()
}

func (s *SearchHotelListRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *SearchHotelListRequest) GetCityCode() *string {
	return s.CityCode
}

func (s *SearchHotelListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchHotelListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchHotelListRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *SearchHotelListRequest) SetAccountNo(v int64) *SearchHotelListRequest {
	s.AccountNo = &v
	return s
}

func (s *SearchHotelListRequest) SetCityCode(v string) *SearchHotelListRequest {
	s.CityCode = &v
	return s
}

func (s *SearchHotelListRequest) SetPageNo(v int32) *SearchHotelListRequest {
	s.PageNo = &v
	return s
}

func (s *SearchHotelListRequest) SetPageSize(v int32) *SearchHotelListRequest {
	s.PageSize = &v
	return s
}

func (s *SearchHotelListRequest) SetTracerId(v string) *SearchHotelListRequest {
	s.TracerId = &v
	return s
}

func (s *SearchHotelListRequest) Validate() error {
	return dara.Validate(s)
}
