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
	// The sorting direction. Valid values:
	//
	// - DESC: descending order (default)
	//
	// - ASC: ascending order
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
	// The type of the domain name. This parameter is not case-sensitive. Valid values:
	//
	// - PUBLIC: an authoritative domain name (default)
	//
	// - CACHE: an authoritative proxy domain name
	//
	// example:
	//
	// PUBLIC
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The end date. Specify the date in the YYYY-MM-DD format. For example: 2023-03-13.
	//
	// example:
	//
	// 2023-03-29
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
	// The number of entries to return on each page. The maximum value is 1000 and the minimum value is 1.
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
	// LIKE
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The start date. Specify the date in the YYYY-MM-DD format. For example: 2023-03-01.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-03-29
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The resolution count threshold. This parameter lets you query subdomains based on their resolution counts.
	//
	// If you do not specify this parameter, data for subdomains with a resolution count greater than 0 is returned.
	//
	// If you specify a value less than 0, all data is returned.
	//
	// If you specify 0, data for subdomains with a resolution count of 0 is returned.
	//
	// If you specify a value greater than 0, data for subdomains with a resolution count less than or equal to the specified value is returned.
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
