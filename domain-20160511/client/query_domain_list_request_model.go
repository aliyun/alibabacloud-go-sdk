// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeadEndDate(v int64) *QueryDomainListRequest
	GetDeadEndDate() *int64
	SetDeadStartDate(v int64) *QueryDomainListRequest
	GetDeadStartDate() *int64
	SetDomainName(v string) *QueryDomainListRequest
	GetDomainName() *string
	SetDomainType(v string) *QueryDomainListRequest
	GetDomainType() *string
	SetEndDate(v string) *QueryDomainListRequest
	GetEndDate() *string
	SetGroupId(v string) *QueryDomainListRequest
	GetGroupId() *string
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
	SetRegEndDate(v int64) *QueryDomainListRequest
	GetRegEndDate() *int64
	SetRegStartDate(v int64) *QueryDomainListRequest
	GetRegStartDate() *int64
	SetStartDate(v string) *QueryDomainListRequest
	GetStartDate() *string
	SetUserClientIp(v string) *QueryDomainListRequest
	GetUserClientIp() *string
}

type QueryDomainListRequest struct {
	// The end of the time range to query based on the time when domain names expire.
	//
	// example:
	//
	// 1696435200000
	DeadEndDate *int64 `json:"DeadEndDate,omitempty" xml:"DeadEndDate,omitempty"`
	// The beginning of the time range to query based on the time when domain names expire.
	//
	// example:
	//
	// 1694016000000
	DeadStartDate *int64 `json:"DeadStartDate,omitempty" xml:"DeadStartDate,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test003.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the domain name. Valid values:
	//
	// GUOJI, TONGYONG, GUONEI, NAME, and WEIBO.
	//
	// example:
	//
	// GUONEI
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The end of the time range to query based on the time when domain names expire.
	//
	// example:
	//
	// 2023-01-11 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The group ID.
	//
	// example:
	//
	// 123
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The language of the error message to return if the request fails. Valid values:
	//
	// zh: Chinese. en: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The order in which you want to sort the queried domain names. Valid values:
	//
	// ASC: ascending order. DESC: descending order.
	//
	// example:
	//
	// DESC
	OrderByType *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	// The field by which domain names to be queried are sorted. Valid values:
	//
	// REGDATE: registration time. DEADDATE: expiration time. CREATEDATE: creation time.
	//
	// example:
	//
	// REGDATE
	OrderKeyType *string `json:"OrderKeyType,omitempty" xml:"OrderKeyType,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type of the domain name. Valid values:
	//
	// New gTLD, gTLD, ccTLD, and other.
	//
	// example:
	//
	// ccTLD
	ProductDomainType *string `json:"ProductDomainType,omitempty" xml:"ProductDomainType,omitempty"`
	// The type of the query. Valid values:
	//
	// 1: renewal. 2: redemption. 4: transfer.
	//
	// example:
	//
	// 1
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The end of the time range to query based on the time when domain names were registered.
	//
	// example:
	//
	// 1696435200000
	RegEndDate *int64 `json:"RegEndDate,omitempty" xml:"RegEndDate,omitempty"`
	// The beginning of the time range to query based on the time when domain names were registered.
	//
	// example:
	//
	// 1694016000000
	RegStartDate *int64 `json:"RegStartDate,omitempty" xml:"RegStartDate,omitempty"`
	// The beginning of the time range to query based on the time when domain names expire.
	//
	// example:
	//
	// 2023-01-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The IP address of the client. Set the value to 127.0.0.1.
	//
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

func (s *QueryDomainListRequest) GetDeadEndDate() *int64 {
	return s.DeadEndDate
}

func (s *QueryDomainListRequest) GetDeadStartDate() *int64 {
	return s.DeadStartDate
}

func (s *QueryDomainListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainListRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *QueryDomainListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryDomainListRequest) GetGroupId() *string {
	return s.GroupId
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

func (s *QueryDomainListRequest) GetRegEndDate() *int64 {
	return s.RegEndDate
}

func (s *QueryDomainListRequest) GetRegStartDate() *int64 {
	return s.RegStartDate
}

func (s *QueryDomainListRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryDomainListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainListRequest) SetDeadEndDate(v int64) *QueryDomainListRequest {
	s.DeadEndDate = &v
	return s
}

func (s *QueryDomainListRequest) SetDeadStartDate(v int64) *QueryDomainListRequest {
	s.DeadStartDate = &v
	return s
}

func (s *QueryDomainListRequest) SetDomainName(v string) *QueryDomainListRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainListRequest) SetDomainType(v string) *QueryDomainListRequest {
	s.DomainType = &v
	return s
}

func (s *QueryDomainListRequest) SetEndDate(v string) *QueryDomainListRequest {
	s.EndDate = &v
	return s
}

func (s *QueryDomainListRequest) SetGroupId(v string) *QueryDomainListRequest {
	s.GroupId = &v
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

func (s *QueryDomainListRequest) SetRegEndDate(v int64) *QueryDomainListRequest {
	s.RegEndDate = &v
	return s
}

func (s *QueryDomainListRequest) SetRegStartDate(v int64) *QueryDomainListRequest {
	s.RegStartDate = &v
	return s
}

func (s *QueryDomainListRequest) SetStartDate(v string) *QueryDomainListRequest {
	s.StartDate = &v
	return s
}

func (s *QueryDomainListRequest) SetUserClientIp(v string) *QueryDomainListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainListRequest) Validate() error {
	return dara.Validate(s)
}
