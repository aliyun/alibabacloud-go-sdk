// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsRequestStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribePdnsRequestStatisticsRequest
	GetDomainName() *string
	SetEndDate(v string) *DescribePdnsRequestStatisticsRequest
	GetEndDate() *string
	SetLang(v string) *DescribePdnsRequestStatisticsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribePdnsRequestStatisticsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribePdnsRequestStatisticsRequest
	GetPageSize() *int64
	SetStartDate(v string) *DescribePdnsRequestStatisticsRequest
	GetStartDate() *string
	SetSubDomain(v string) *DescribePdnsRequestStatisticsRequest
	GetSubDomain() *string
	SetType(v string) *DescribePdnsRequestStatisticsRequest
	GetType() *string
}

type DescribePdnsRequestStatisticsRequest struct {
	// The primary domain name to query.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end date of the query. Use the YYYY-MM-DD format.
	//
	// The default value is the current date.
	//
	// example:
	//
	// 2024-07-14
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language of the request and response. The default value is **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from **1**. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start date of the query. Use the YYYY-MM-DD format.
	//
	// You can query data from the last 90 days only.
	//
	// example:
	//
	// 2024-06-14
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The subdomain for which to view statistics.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The type of request statistics to query. Valid values:
	//
	// - DOMAIN: queries statistics by domain name.
	//
	// - SUB_DOMAIN: queries statistics by subdomain.
	//
	// example:
	//
	// DOMAIN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePdnsRequestStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribePdnsRequestStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribePdnsRequestStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePdnsRequestStatisticsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePdnsRequestStatisticsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePdnsRequestStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribePdnsRequestStatisticsRequest) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribePdnsRequestStatisticsRequest) GetType() *string {
	return s.Type
}

func (s *DescribePdnsRequestStatisticsRequest) SetDomainName(v string) *DescribePdnsRequestStatisticsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribePdnsRequestStatisticsRequest) SetEndDate(v string) *DescribePdnsRequestStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePdnsRequestStatisticsRequest) SetLang(v string) *DescribePdnsRequestStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePdnsRequestStatisticsRequest) SetPageNumber(v int64) *DescribePdnsRequestStatisticsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePdnsRequestStatisticsRequest) SetPageSize(v int64) *DescribePdnsRequestStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePdnsRequestStatisticsRequest) SetStartDate(v string) *DescribePdnsRequestStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribePdnsRequestStatisticsRequest) SetSubDomain(v string) *DescribePdnsRequestStatisticsRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribePdnsRequestStatisticsRequest) SetType(v string) *DescribePdnsRequestStatisticsRequest {
	s.Type = &v
	return s
}

func (s *DescribePdnsRequestStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
