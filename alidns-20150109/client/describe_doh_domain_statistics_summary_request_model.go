// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohDomainStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDohDomainStatisticsSummaryRequest
	GetDomainName() *string
	SetEndDate(v string) *DescribeDohDomainStatisticsSummaryRequest
	GetEndDate() *string
	SetLang(v string) *DescribeDohDomainStatisticsSummaryRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeDohDomainStatisticsSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDohDomainStatisticsSummaryRequest
	GetPageSize() *int32
	SetStartDate(v string) *DescribeDohDomainStatisticsSummaryRequest
	GetStartDate() *string
}

type DescribeDohDomainStatisticsSummaryRequest struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// The default value is the day when you perform the operation.
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
}

func (s DescribeDohDomainStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohDomainStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsSummaryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDohDomainStatisticsSummaryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDohDomainStatisticsSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDohDomainStatisticsSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDohDomainStatisticsSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDohDomainStatisticsSummaryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetDomainName(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetEndDate(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetLang(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetPageNumber(v int32) *DescribeDohDomainStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetPageSize(v int32) *DescribeDohDomainStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) SetStartDate(v string) *DescribeDohDomainStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
