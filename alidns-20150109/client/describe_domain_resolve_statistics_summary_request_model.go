// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResolveStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeDomainResolveStatisticsSummaryRequest
	GetDirection() *string
	SetEndDate(v string) *DescribeDomainResolveStatisticsSummaryRequest
	GetEndDate() *string
	SetKeyword(v string) *DescribeDomainResolveStatisticsSummaryRequest
	GetKeyword() *string
	SetLang(v string) *DescribeDomainResolveStatisticsSummaryRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeDomainResolveStatisticsSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDomainResolveStatisticsSummaryRequest
	GetPageSize() *int32
	SetSearchMode(v string) *DescribeDomainResolveStatisticsSummaryRequest
	GetSearchMode() *string
	SetStartDate(v string) *DescribeDomainResolveStatisticsSummaryRequest
	GetStartDate() *string
	SetThreshold(v int64) *DescribeDomainResolveStatisticsSummaryRequest
	GetThreshold() *int64
}

type DescribeDomainResolveStatisticsSummaryRequest struct {
	// The sort direction. Valid values:
	//
	// - DESC: descending
	//
	// - ASC: ascending
	//
	// example:
	//
	// DESC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end date. The format is yyyy-MM-dd. For example, 2023-03-13.
	//
	// example:
	//
	// 2023-03-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The keyword. This parameter is used with SearchMode.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language. Valid values: zh, en, and ja.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The maximum value is 1000. The minimum value is 1.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search mode of the keyword. Valid values:
	//
	// - LIKE: fuzzy search (default)
	//
	// - EXACT: exact match
	//
	// example:
	//
	// EXACT
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The start date. The format is yyyy-MM-dd. For example, 2023-03-01.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The threshold for the number of DNS queries. This parameter filters domain names by query volume.
	//
	// If you do not specify this parameter, the operation returns domain names with more than zero queries.
	//
	// If you specify a value less than 0, the operation returns all domain names.
	//
	// If you specify 0, the operation returns domain names with zero queries.
	//
	// If you specify a value greater than 0, the operation returns domain names with a query volume up to this value.
	//
	// example:
	//
	// -1
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeDomainResolveStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResolveStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) GetThreshold() *int64 {
	return s.Threshold
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetDirection(v string) *DescribeDomainResolveStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetEndDate(v string) *DescribeDomainResolveStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetKeyword(v string) *DescribeDomainResolveStatisticsSummaryRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetLang(v string) *DescribeDomainResolveStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetPageNumber(v int32) *DescribeDomainResolveStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetPageSize(v int32) *DescribeDomainResolveStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetSearchMode(v string) *DescribeDomainResolveStatisticsSummaryRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetStartDate(v string) *DescribeDomainResolveStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) SetThreshold(v int64) *DescribeDomainResolveStatisticsSummaryRequest {
	s.Threshold = &v
	return s
}

func (s *DescribeDomainResolveStatisticsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
