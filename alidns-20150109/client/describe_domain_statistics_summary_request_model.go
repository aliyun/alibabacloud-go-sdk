// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribeDomainStatisticsSummaryRequest
	GetEndDate() *string
	SetKeyword(v string) *DescribeDomainStatisticsSummaryRequest
	GetKeyword() *string
	SetLang(v string) *DescribeDomainStatisticsSummaryRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeDomainStatisticsSummaryRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainStatisticsSummaryRequest
	GetPageSize() *int64
	SetSearchMode(v string) *DescribeDomainStatisticsSummaryRequest
	GetSearchMode() *string
	SetStartDate(v string) *DescribeDomainStatisticsSummaryRequest
	GetStartDate() *string
	SetThreshold(v int64) *DescribeDomainStatisticsSummaryRequest
	GetThreshold() *int64
}

type DescribeDomainStatisticsSummaryRequest struct {
	// The end of the time range to query. Specify the time in the **YYYY-MM-DD*	- format.
	//
	// The default value is the day when you perform the operation.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The keyword for searches in %KeyWord% mode. The value is not case-sensitive.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language type.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search mode of the keyword. Valid values:
	//
	// 	- **LIKE**: fuzzy match (default).
	//
	// 	- **EXACT**: exact match.
	//
	// example:
	//
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The beginning of the time range to query. Specify the time in the **YYYY-MM-DD*	- format.
	//
	// You can only query DNS records of the last 90 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The threshold of query volume that can be obtained. You can also obtain data about a domain name with the query volume less than or equal to the threshold. For example, if you set this parameter to 100, you can query domain names with less than 100 queries.
	//
	// example:
	//
	// 12
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeDomainStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDomainStatisticsSummaryRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDomainStatisticsSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainStatisticsSummaryRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainStatisticsSummaryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainStatisticsSummaryRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeDomainStatisticsSummaryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDomainStatisticsSummaryRequest) GetThreshold() *int64 {
	return s.Threshold
}

func (s *DescribeDomainStatisticsSummaryRequest) SetEndDate(v string) *DescribeDomainStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetKeyword(v string) *DescribeDomainStatisticsSummaryRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetLang(v string) *DescribeDomainStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetPageNumber(v int64) *DescribeDomainStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetPageSize(v int64) *DescribeDomainStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetSearchMode(v string) *DescribeDomainStatisticsSummaryRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetStartDate(v string) *DescribeDomainStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) SetThreshold(v int64) *DescribeDomainStatisticsSummaryRequest {
	s.Threshold = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
