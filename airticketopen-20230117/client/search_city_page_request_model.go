// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCityPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *SearchCityPageRequest
	GetAccountNo() *int64
	SetCount(v int32) *SearchCityPageRequest
	GetCount() *int32
	SetCountryCode(v string) *SearchCityPageRequest
	GetCountryCode() *string
	SetStart(v int32) *SearchCityPageRequest
	GetStart() *int32
	SetTracerId(v string) *SearchCityPageRequest
	GetTracerId() *string
}

type SearchCityPageRequest struct {
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

func (s SearchCityPageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchCityPageRequest) GoString() string {
	return s.String()
}

func (s *SearchCityPageRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *SearchCityPageRequest) GetCount() *int32 {
	return s.Count
}

func (s *SearchCityPageRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *SearchCityPageRequest) GetStart() *int32 {
	return s.Start
}

func (s *SearchCityPageRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *SearchCityPageRequest) SetAccountNo(v int64) *SearchCityPageRequest {
	s.AccountNo = &v
	return s
}

func (s *SearchCityPageRequest) SetCount(v int32) *SearchCityPageRequest {
	s.Count = &v
	return s
}

func (s *SearchCityPageRequest) SetCountryCode(v string) *SearchCityPageRequest {
	s.CountryCode = &v
	return s
}

func (s *SearchCityPageRequest) SetStart(v int32) *SearchCityPageRequest {
	s.Start = &v
	return s
}

func (s *SearchCityPageRequest) SetTracerId(v string) *SearchCityPageRequest {
	s.TracerId = &v
	return s
}

func (s *SearchCityPageRequest) Validate() error {
	return dara.Validate(s)
}
