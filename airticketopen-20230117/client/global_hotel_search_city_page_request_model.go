// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelSearchCityPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelSearchCityPageRequest
	GetAccountNo() *int64
	SetCount(v int32) *GlobalHotelSearchCityPageRequest
	GetCount() *int32
	SetCountryCode(v string) *GlobalHotelSearchCityPageRequest
	GetCountryCode() *string
	SetStart(v int32) *GlobalHotelSearchCityPageRequest
	GetStart() *int32
	SetTracerId(v string) *GlobalHotelSearchCityPageRequest
	GetTracerId() *string
}

type GlobalHotelSearchCityPageRequest struct {
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
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelSearchCityPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchCityPageRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchCityPageRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelSearchCityPageRequest) GetCount() *int32 {
	return s.Count
}

func (s *GlobalHotelSearchCityPageRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *GlobalHotelSearchCityPageRequest) GetStart() *int32 {
	return s.Start
}

func (s *GlobalHotelSearchCityPageRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelSearchCityPageRequest) SetAccountNo(v int64) *GlobalHotelSearchCityPageRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelSearchCityPageRequest) SetCount(v int32) *GlobalHotelSearchCityPageRequest {
	s.Count = &v
	return s
}

func (s *GlobalHotelSearchCityPageRequest) SetCountryCode(v string) *GlobalHotelSearchCityPageRequest {
	s.CountryCode = &v
	return s
}

func (s *GlobalHotelSearchCityPageRequest) SetStart(v int32) *GlobalHotelSearchCityPageRequest {
	s.Start = &v
	return s
}

func (s *GlobalHotelSearchCityPageRequest) SetTracerId(v string) *GlobalHotelSearchCityPageRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelSearchCityPageRequest) Validate() error {
	return dara.Validate(s)
}
