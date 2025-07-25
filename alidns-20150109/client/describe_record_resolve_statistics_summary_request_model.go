// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordResolveStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeRecordResolveStatisticsSummaryRequest
	GetDirection() *string
	SetDomainName(v string) *DescribeRecordResolveStatisticsSummaryRequest
	GetDomainName() *string
	SetDomainType(v string) *DescribeRecordResolveStatisticsSummaryRequest
	GetDomainType() *string
	SetEndDate(v string) *DescribeRecordResolveStatisticsSummaryRequest
	GetEndDate() *string
	SetKeyword(v string) *DescribeRecordResolveStatisticsSummaryRequest
	GetKeyword() *string
	SetLang(v string) *DescribeRecordResolveStatisticsSummaryRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeRecordResolveStatisticsSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRecordResolveStatisticsSummaryRequest
	GetPageSize() *int32
	SetSearchMode(v string) *DescribeRecordResolveStatisticsSummaryRequest
	GetSearchMode() *string
	SetStartDate(v string) *DescribeRecordResolveStatisticsSummaryRequest
	GetStartDate() *string
	SetThreshold(v int64) *DescribeRecordResolveStatisticsSummaryRequest
	GetThreshold() *int64
}

type DescribeRecordResolveStatisticsSummaryRequest struct {
	// The order in which the returned entries are sorted. Valid values:
	//
	// 	- DESC (default): descending order
	//
	// 	- ASC: ascending order
	//
	// example:
	//
	// DESC
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The type of the domain name. The parameter value is not case-sensitive. Valid values:
	//
	// 	- PUBLIC (default): hosted public domain name
	//
	// 	- CACHE: cache-accelerated domain name
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The end date of the time range to be queried. Specify the time in the yyyy-MM-dd format, such as 2023-03-13.
	//
	// example:
	//
	// 2023-03-29
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The keyword. Keyword is used together with SearchMode.
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
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The start date of the time range to be queried. Specify the time in the yyyy-MM-dd format, such as 2023-03-01.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-29
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The threshold for the number of Domain Name System (DNS) requests. You can query the subdomain names at the specified quantity level of DNS requests and query the number of DNS requests for each subdomain name.
	//
	// If you do not specify this parameter, the data about the subdomain names that have DNS requests is obtained.
	//
	// If you set this parameter to a value less than 0, the data about all subdomain names is obtained.
	//
	// If you set this parameter to 0, the data about the subdomain names that do not have DNS requests is obtained.
	//
	// If you set this parameter to a value greater than 0, the data about the subdomain names whose number of DNS requests is less than or equal to the value of this parameter is obtained.
	//
	// example:
	//
	// -1
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeRecordResolveStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordResolveStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) GetThreshold() *int64 {
	return s.Threshold
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetDirection(v string) *DescribeRecordResolveStatisticsSummaryRequest {
	s.Direction = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetDomainName(v string) *DescribeRecordResolveStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetDomainType(v string) *DescribeRecordResolveStatisticsSummaryRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetEndDate(v string) *DescribeRecordResolveStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetKeyword(v string) *DescribeRecordResolveStatisticsSummaryRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetLang(v string) *DescribeRecordResolveStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetPageNumber(v int32) *DescribeRecordResolveStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetPageSize(v int32) *DescribeRecordResolveStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetSearchMode(v string) *DescribeRecordResolveStatisticsSummaryRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetStartDate(v string) *DescribeRecordResolveStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) SetThreshold(v int64) *DescribeRecordResolveStatisticsSummaryRequest {
	s.Threshold = &v
	return s
}

func (s *DescribeRecordResolveStatisticsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
