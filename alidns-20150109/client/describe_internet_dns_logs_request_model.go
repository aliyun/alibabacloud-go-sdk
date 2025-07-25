// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetDnsLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *DescribeInternetDnsLogsRequest
	GetAccountId() *int64
	SetDomainName(v string) *DescribeInternetDnsLogsRequest
	GetDomainName() *string
	SetEndTimestamp(v int64) *DescribeInternetDnsLogsRequest
	GetEndTimestamp() *int64
	SetLang(v string) *DescribeInternetDnsLogsRequest
	GetLang() *string
	SetModule(v string) *DescribeInternetDnsLogsRequest
	GetModule() *string
	SetPageNumber(v int32) *DescribeInternetDnsLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInternetDnsLogsRequest
	GetPageSize() *int32
	SetQueryCondition(v string) *DescribeInternetDnsLogsRequest
	GetQueryCondition() *string
	SetRecursionProtocolType(v string) *DescribeInternetDnsLogsRequest
	GetRecursionProtocolType() *string
	SetStartTimestamp(v int64) *DescribeInternetDnsLogsRequest
	GetStartTimestamp() *int64
}

type DescribeInternetDnsLogsRequest struct {
	// The account ID displayed on the Recursive Resolution (Public DNS) page after you activate Alibaba Cloud Public DNS.
	//
	// example:
	//
	// 51**4
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the query (timestamp, unit: milliseconds). 	Warning: If the query time span is too large and the amount of resolution logs for the queried domain is excessive, it may lead to a query timeout or inaccurate query results.
	//
	// example:
	//
	// 1709196299999
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// Return value language, options:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Module type
	//
	// - AUTHORITY (default): Public Authoritative DNS
	//
	// - RECURSION: Public Recursive DNS
	//
	// example:
	//
	// AUTHORITY
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// Page number, default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size for query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Query parameters
	//
	// - sourceIp: Source IP address
	//
	// - queryNameFuzzy: Domain name (fuzzy value)
	//
	// - queryType: Record type
	//
	// - value: Resolution result
	//
	// - status: Status
	//
	// - serverIp: Resolution server IP
	//
	// example:
	//
	// {"sourceIp":"59.82.XX.XX","queryType":"A"}
	QueryCondition        *string `json:"QueryCondition,omitempty" xml:"QueryCondition,omitempty"`
	RecursionProtocolType *string `json:"RecursionProtocolType,omitempty" xml:"RecursionProtocolType,omitempty"`
	// The start time of the query (timestamp, unit: milliseconds).
	//
	// example:
	//
	// 1709192640000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeInternetDnsLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDnsLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetDnsLogsRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *DescribeInternetDnsLogsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeInternetDnsLogsRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeInternetDnsLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetDnsLogsRequest) GetModule() *string {
	return s.Module
}

func (s *DescribeInternetDnsLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInternetDnsLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInternetDnsLogsRequest) GetQueryCondition() *string {
	return s.QueryCondition
}

func (s *DescribeInternetDnsLogsRequest) GetRecursionProtocolType() *string {
	return s.RecursionProtocolType
}

func (s *DescribeInternetDnsLogsRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeInternetDnsLogsRequest) SetAccountId(v int64) *DescribeInternetDnsLogsRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetDomainName(v string) *DescribeInternetDnsLogsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetEndTimestamp(v int64) *DescribeInternetDnsLogsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetLang(v string) *DescribeInternetDnsLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetModule(v string) *DescribeInternetDnsLogsRequest {
	s.Module = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetPageNumber(v int32) *DescribeInternetDnsLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetPageSize(v int32) *DescribeInternetDnsLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetQueryCondition(v string) *DescribeInternetDnsLogsRequest {
	s.QueryCondition = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetRecursionProtocolType(v string) *DescribeInternetDnsLogsRequest {
	s.RecursionProtocolType = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) SetStartTimestamp(v int64) *DescribeInternetDnsLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeInternetDnsLogsRequest) Validate() error {
	return dara.Validate(s)
}
