// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohSubDomainStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDohSubDomainStatisticsSummaryRequest
	GetDomainName() *string
	SetEndDate(v string) *DescribeDohSubDomainStatisticsSummaryRequest
	GetEndDate() *string
	SetLang(v string) *DescribeDohSubDomainStatisticsSummaryRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeDohSubDomainStatisticsSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDohSubDomainStatisticsSummaryRequest
	GetPageSize() *int32
	SetStartDate(v string) *DescribeDohSubDomainStatisticsSummaryRequest
	GetStartDate() *string
	SetSubDomain(v string) *DescribeDohSubDomainStatisticsSummaryRequest
	GetSubDomain() *string
}

type DescribeDohSubDomainStatisticsSummaryRequest struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// The default value is the day when you query the required data.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The language type.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// You can query only the DNS records of the last 90 days. `The value of StartDate must be greater than or equal to the difference between the current date and 90`.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The subdomain.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeDohSubDomainStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetDomainName(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetEndDate(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetLang(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetPageNumber(v int32) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetPageSize(v int32) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetStartDate(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) SetSubDomain(v string) *DescribeDohSubDomainStatisticsSummaryRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
