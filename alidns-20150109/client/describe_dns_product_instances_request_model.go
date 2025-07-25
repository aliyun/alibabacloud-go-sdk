// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsProductInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeDnsProductInstancesRequest
	GetDirection() *string
	SetDomainType(v string) *DescribeDnsProductInstancesRequest
	GetDomainType() *string
	SetLang(v string) *DescribeDnsProductInstancesRequest
	GetLang() *string
	SetOrderBy(v string) *DescribeDnsProductInstancesRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *DescribeDnsProductInstancesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDnsProductInstancesRequest
	GetPageSize() *int64
	SetUserClientIp(v string) *DescribeDnsProductInstancesRequest
	GetUserClientIp() *string
	SetVersionCode(v string) *DescribeDnsProductInstancesRequest
	GetVersionCode() *string
}

type DescribeDnsProductInstancesRequest struct {
	// The order in which you want to sort returned entries. Valid values:
	//
	// 	- DESC: Returned entries are sorted in descending order. If this parameter is left empty, returned entries are sorted in descending order by default.
	//
	// 	- ASC: Returned entries are sorted in ascending order.
	//
	// example:
	//
	// DESC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- PUBLIC (default): hosted public domain name
	//
	// 	- CACHE: cached public domain name
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The method that is used to sort returned entries. Valid values:
	//
	// 	- createDate: sorts returned entries by creation time. If this parameter is left empty, returned entries are sorted by creation time by default.
	//
	// 	- expireDate: sorts returned entries by expiration time.
	//
	// example:
	//
	// createDate
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The version code of the Alibaba Cloud DNS instance.
	//
	// example:
	//
	// version1
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s DescribeDnsProductInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsProductInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsProductInstancesRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeDnsProductInstancesRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDnsProductInstancesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsProductInstancesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeDnsProductInstancesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDnsProductInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDnsProductInstancesRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeDnsProductInstancesRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDnsProductInstancesRequest) SetDirection(v string) *DescribeDnsProductInstancesRequest {
	s.Direction = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetDomainType(v string) *DescribeDnsProductInstancesRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetLang(v string) *DescribeDnsProductInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetOrderBy(v string) *DescribeDnsProductInstancesRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetPageNumber(v int64) *DescribeDnsProductInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetPageSize(v int64) *DescribeDnsProductInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetUserClientIp(v string) *DescribeDnsProductInstancesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) SetVersionCode(v string) *DescribeDnsProductInstancesRequest {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsProductInstancesRequest) Validate() error {
	return dara.Validate(s)
}
