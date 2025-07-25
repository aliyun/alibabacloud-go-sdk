// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordStatisticsSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeRecordStatisticsSummaryRequest
	GetDomainName() *string
	SetDomainType(v string) *DescribeRecordStatisticsSummaryRequest
	GetDomainType() *string
	SetEndDate(v string) *DescribeRecordStatisticsSummaryRequest
	GetEndDate() *string
	SetKeyword(v string) *DescribeRecordStatisticsSummaryRequest
	GetKeyword() *string
	SetLang(v string) *DescribeRecordStatisticsSummaryRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeRecordStatisticsSummaryRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRecordStatisticsSummaryRequest
	GetPageSize() *int64
	SetSearchMode(v string) *DescribeRecordStatisticsSummaryRequest
	GetSearchMode() *string
	SetStartDate(v string) *DescribeRecordStatisticsSummaryRequest
	GetStartDate() *string
	SetThreshold(v int64) *DescribeRecordStatisticsSummaryRequest
	GetThreshold() *int64
}

type DescribeRecordStatisticsSummaryRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns-example.com
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
	// The end date of the query. Specify the start date in the **YYYY-MM-DD*	- format.
	//
	// The default value is the day when you query the data.
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
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
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
	// 	- **LIKE*	- (default): fuzzy search
	//
	// 	- **EXACT**: exact search
	//
	// example:
	//
	// EXACT
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The start date of the query. Specify the start date in the **YYYY-MM-DD*	- format.
	//
	// You can only query the DNS records within the last 90 days.``
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-07-04
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The maximum number of DNS requests that you can obtain. You can obtain data about a domain name with DNS request volume less than or equal to the maximum number. For example, if you set this parameter to 100, you can query domain names with less than 100 DNS requests.
	//
	// example:
	//
	// 12
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeRecordStatisticsSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordStatisticsSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordStatisticsSummaryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRecordStatisticsSummaryRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeRecordStatisticsSummaryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeRecordStatisticsSummaryRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeRecordStatisticsSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecordStatisticsSummaryRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRecordStatisticsSummaryRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRecordStatisticsSummaryRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeRecordStatisticsSummaryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeRecordStatisticsSummaryRequest) GetThreshold() *int64 {
	return s.Threshold
}

func (s *DescribeRecordStatisticsSummaryRequest) SetDomainName(v string) *DescribeRecordStatisticsSummaryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetDomainType(v string) *DescribeRecordStatisticsSummaryRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetEndDate(v string) *DescribeRecordStatisticsSummaryRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetKeyword(v string) *DescribeRecordStatisticsSummaryRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetLang(v string) *DescribeRecordStatisticsSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetPageNumber(v int64) *DescribeRecordStatisticsSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetPageSize(v int64) *DescribeRecordStatisticsSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetSearchMode(v string) *DescribeRecordStatisticsSummaryRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetStartDate(v string) *DescribeRecordStatisticsSummaryRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) SetThreshold(v int64) *DescribeRecordStatisticsSummaryRequest {
	s.Threshold = &v
	return s
}

func (s *DescribeRecordStatisticsSummaryRequest) Validate() error {
	return dara.Validate(s)
}
