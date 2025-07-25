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
	// The order in which you want to sort the returned entries. Valid values:
	//
	// 	- DESC: the descending order
	//
	// 	- ASC: the ascending order
	//
	// example:
	//
	// DESC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end time in the yyyy-MM-dd format, for example, 2023-03-13.
	//
	// example:
	//
	// 2023-03-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The keyword. The Keyword parameter is used together with the SearchMode parameter.
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
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 1000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search mode of the keyword. Valid values:
	//
	// 	- LIKE (default): fuzzy search
	//
	// 	- EXACT: exact search
	//
	// example:
	//
	// EXACT
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The start time in the yyyy-MM-dd format, for example, 2023-03-01.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The threshold for the number of Domain Name System (DNS) requests. You can query the domain names at the specified quantity level of DNS requests and query the number of DNS requests for each domain name.
	//
	// If you do not specify this parameter, the data about the domain names that have DNS requests is obtained.
	//
	// If you set this parameter to a value less than 0, the data about all domain names is obtained.
	//
	// If you set this parameter to 0, the data about the domain names that do not have DNS requests is obtained.
	//
	// If you set this parameter to a value greater than 0, the data about the domain names whose number of DNS requests is less than or equal to the value of this parameter is obtained.
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
