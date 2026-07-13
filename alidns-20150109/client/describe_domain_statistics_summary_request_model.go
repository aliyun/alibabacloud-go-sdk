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
	// The end date of the query. The format is **YYYY-MM-DD**.
	//
	// The default value is the current date.
	//
	// example:
	//
	// 2019-07-04
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The keyword. The search is performed in \\`%KeyWord%\\` mode and is case-insensitive.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The start value is **1**. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The maximum value is **100*	- and the minimum value is **1**. The default value is **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search mode for the keyword. Valid values:
	//
	// - **LIKE*	- (default): Fuzzy match
	//
	// - **EXACT**: Exact match
	//
	// example:
	//
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The start date of the query. The format is **YYYY-MM-DD**.
	//
	// You can query data from the last 90 days only. The value of `StartDate` must be greater than or equal to the current date minus 90 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The threshold for the number of requests. This operation returns domain names whose request count is less than or equal to the specified value. For example, if you set this parameter to 100, domain names with 100 or fewer requests are returned. If you do not specify this parameter, all domain names are returned.
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
