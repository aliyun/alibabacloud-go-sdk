// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcompany(v string) *QueryDomainListRequest
	GetCcompany() *string
	SetDomainName(v string) *QueryDomainListRequest
	GetDomainName() *string
	SetEndExpirationDate(v int64) *QueryDomainListRequest
	GetEndExpirationDate() *int64
	SetEndRegistrationDate(v int64) *QueryDomainListRequest
	GetEndRegistrationDate() *int64
	SetLang(v string) *QueryDomainListRequest
	GetLang() *string
	SetOrderByType(v string) *QueryDomainListRequest
	GetOrderByType() *string
	SetOrderKeyType(v string) *QueryDomainListRequest
	GetOrderKeyType() *string
	SetPageNum(v int32) *QueryDomainListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryDomainListRequest
	GetPageSize() *int32
	SetProductDomainType(v string) *QueryDomainListRequest
	GetProductDomainType() *string
	SetQueryType(v string) *QueryDomainListRequest
	GetQueryType() *string
	SetStartExpirationDate(v int64) *QueryDomainListRequest
	GetStartExpirationDate() *int64
	SetStartRegistrationDate(v int64) *QueryDomainListRequest
	GetStartRegistrationDate() *int64
	SetUserClientIp(v string) *QueryDomainListRequest
	GetUserClientIp() *string
}

type QueryDomainListRequest struct {
	Ccompany *string `json:"Ccompany,omitempty" xml:"Ccompany,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1522080000000
	EndExpirationDate *int64 `json:"EndExpirationDate,omitempty" xml:"EndExpirationDate,omitempty"`
	// example:
	//
	// 1522080000000
	EndRegistrationDate *int64 `json:"EndRegistrationDate,omitempty" xml:"EndRegistrationDate,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// DESC
	OrderByType *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	// example:
	//
	// ExpirationDate
	OrderKeyType *string `json:"OrderKeyType,omitempty" xml:"OrderKeyType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// gTLD
	ProductDomainType *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	// example:
	//
	// 1
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// example:
	//
	// 1522080000000
	StartExpirationDate *int64 `json:"StartExpirationDate,omitempty" xml:"StartExpirationDate,omitempty"`
	// example:
	//
	// 1522080000000
	StartRegistrationDate *int64 `json:"StartRegistrationDate,omitempty" xml:"StartRegistrationDate,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainListRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainListRequest) GetCcompany() *string {
	return s.Ccompany
}

func (s *QueryDomainListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainListRequest) GetEndExpirationDate() *int64 {
	return s.EndExpirationDate
}

func (s *QueryDomainListRequest) GetEndRegistrationDate() *int64 {
	return s.EndRegistrationDate
}

func (s *QueryDomainListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainListRequest) GetOrderByType() *string {
	return s.OrderByType
}

func (s *QueryDomainListRequest) GetOrderKeyType() *string {
	return s.OrderKeyType
}

func (s *QueryDomainListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryDomainListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDomainListRequest) GetProductDomainType() *string {
	return s.ProductDomainType
}

func (s *QueryDomainListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryDomainListRequest) GetStartExpirationDate() *int64 {
	return s.StartExpirationDate
}

func (s *QueryDomainListRequest) GetStartRegistrationDate() *int64 {
	return s.StartRegistrationDate
}

func (s *QueryDomainListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainListRequest) SetCcompany(v string) *QueryDomainListRequest {
	s.Ccompany = &v
	return s
}

func (s *QueryDomainListRequest) SetDomainName(v string) *QueryDomainListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainListRequest) SetEndExpirationDate(v int64) *QueryDomainListRequest {
	s.EndExpirationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetEndRegistrationDate(v int64) *QueryDomainListRequest {
	s.EndRegistrationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetLang(v string) *QueryDomainListRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainListRequest) SetOrderByType(v string) *QueryDomainListRequest {
	s.OrderByType = &v
	return s
}

func (s *QueryDomainListRequest) SetOrderKeyType(v string) *QueryDomainListRequest {
	s.OrderKeyType = &v
	return s
}

func (s *QueryDomainListRequest) SetPageNum(v int32) *QueryDomainListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryDomainListRequest) SetPageSize(v int32) *QueryDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDomainListRequest) SetProductDomainType(v string) *QueryDomainListRequest {
	s.ProductDomainType = &v
	return s
}

func (s *QueryDomainListRequest) SetQueryType(v string) *QueryDomainListRequest {
	s.QueryType = &v
	return s
}

func (s *QueryDomainListRequest) SetStartExpirationDate(v int64) *QueryDomainListRequest {
	s.StartExpirationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetStartRegistrationDate(v int64) *QueryDomainListRequest {
	s.StartRegistrationDate = &v
	return s
}

func (s *QueryDomainListRequest) SetUserClientIp(v string) *QueryDomainListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainListRequest) Validate() error {
	return dara.Validate(s)
}
