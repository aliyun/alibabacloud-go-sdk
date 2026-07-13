// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeRecordLogsRequest
	GetDomainName() *string
	SetKeyWord(v string) *DescribeRecordLogsRequest
	GetKeyWord() *string
	SetLang(v string) *DescribeRecordLogsRequest
	GetLang() *string
	SetPageNumber(v int64) *DescribeRecordLogsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRecordLogsRequest
	GetPageSize() *int64
	SetStartDate(v string) *DescribeRecordLogsRequest
	GetStartDate() *string
	SetUserClientIp(v string) *DescribeRecordLogsRequest
	GetUserClientIp() *string
	SetEndDate(v string) *DescribeRecordLogsRequest
	GetEndDate() *string
}

type DescribeRecordLogsRequest struct {
	// The domain name.<props="china">You can call the [DescribeDomains](https://help.aliyun.com/document_detail/29751.html) operation to obtain the domain name.<props="intl">You can call the [DescribeDomains](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describedomains) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The keyword. The system searches for the keyword in the "%KeyWord%" pattern. The search is not case-sensitive.
	//
	// example:
	//
	// test
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start of the time range to query. The format is **YYYY-MM-DD**.
	//
	// example:
	//
	// 2015-12-12
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 192.0.2.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The end of the time range to query. The format is **YYYY-MM-DD**.
	//
	// example:
	//
	// 2015-12-12
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
}

func (s DescribeRecordLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordLogsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeRecordLogsRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeRecordLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRecordLogsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRecordLogsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRecordLogsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeRecordLogsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeRecordLogsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeRecordLogsRequest) SetDomainName(v string) *DescribeRecordLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetKeyWord(v string) *DescribeRecordLogsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetLang(v string) *DescribeRecordLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetPageNumber(v int64) *DescribeRecordLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetPageSize(v int64) *DescribeRecordLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetStartDate(v string) *DescribeRecordLogsRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetUserClientIp(v string) *DescribeRecordLogsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRecordLogsRequest) SetEndDate(v string) *DescribeRecordLogsRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeRecordLogsRequest) Validate() error {
	return dara.Validate(s)
}
