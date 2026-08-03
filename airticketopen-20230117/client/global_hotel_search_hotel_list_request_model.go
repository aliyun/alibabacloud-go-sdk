// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelSearchHotelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelSearchHotelListRequest
	GetAccountNo() *int64
	SetCityCode(v string) *GlobalHotelSearchHotelListRequest
	GetCityCode() *string
	SetPageNo(v int32) *GlobalHotelSearchHotelListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GlobalHotelSearchHotelListRequest
	GetPageSize() *int32
	SetTracerId(v string) *GlobalHotelSearchHotelListRequest
	GetTracerId() *string
}

type GlobalHotelSearchHotelListRequest struct {
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

func (s GlobalHotelSearchHotelListRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchHotelListRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchHotelListRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelSearchHotelListRequest) GetCityCode() *string {
	return s.CityCode
}

func (s *GlobalHotelSearchHotelListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GlobalHotelSearchHotelListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GlobalHotelSearchHotelListRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelSearchHotelListRequest) SetAccountNo(v int64) *GlobalHotelSearchHotelListRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelSearchHotelListRequest) SetCityCode(v string) *GlobalHotelSearchHotelListRequest {
	s.CityCode = &v
	return s
}

func (s *GlobalHotelSearchHotelListRequest) SetPageNo(v int32) *GlobalHotelSearchHotelListRequest {
	s.PageNo = &v
	return s
}

func (s *GlobalHotelSearchHotelListRequest) SetPageSize(v int32) *GlobalHotelSearchHotelListRequest {
	s.PageSize = &v
	return s
}

func (s *GlobalHotelSearchHotelListRequest) SetTracerId(v string) *GlobalHotelSearchHotelListRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelSearchHotelListRequest) Validate() error {
	return dara.Validate(s)
}
