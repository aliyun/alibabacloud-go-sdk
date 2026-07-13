// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRecursionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *SearchRecursionRecordsRequest
	GetDirection() *string
	SetEnableStatus(v string) *SearchRecursionRecordsRequest
	GetEnableStatus() *string
	SetMaxResults(v int32) *SearchRecursionRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchRecursionRecordsRequest
	GetNextToken() *string
	SetOrderBy(v string) *SearchRecursionRecordsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *SearchRecursionRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchRecursionRecordsRequest
	GetPageSize() *int32
	SetRemark(v string) *SearchRecursionRecordsRequest
	GetRemark() *string
	SetRequestSource(v string) *SearchRecursionRecordsRequest
	GetRequestSource() *string
	SetRr(v string) *SearchRecursionRecordsRequest
	GetRr() *string
	SetTtl(v int32) *SearchRecursionRecordsRequest
	GetTtl() *int32
	SetType(v string) *SearchRecursionRecordsRequest
	GetType() *string
	SetValue(v string) *SearchRecursionRecordsRequest
	GetValue() *string
	SetWeight(v int32) *SearchRecursionRecordsRequest
	GetWeight() *int32
	SetZoneId(v string) *SearchRecursionRecordsRequest
	GetZoneId() *string
}

type SearchRecursionRecordsRequest struct {
	// The sort order. Valid values: asc for ascending and dsc for descending.
	//
	// example:
	//
	// asc
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The status of the DNS record. Valid values: enable and **disable**.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The maximum number of records to return for the current request.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query.
	//
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sorting method. Valid values: rr, type, value, requestSource, weight, ttl, and enableStatus.
	//
	// example:
	//
	// rr
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The current page number. The value starts from 1. The default value is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. The maximum value is 100. The default value is 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The DNS resolution line. The default value is **default**. For more information, see [DNS resolution lines](https://www.alibabacloud.com/help/en/doc-detail/29807.htm).
	//
	// <props="china">
	//
	// [DNS Line Enumeration](https://help.aliyun.com/document_detail/29807.html)
	//
	//
	//
	// <props="intl">
	//
	// [Resolution Line Enumeration](https://www.alibabacloud.com/help/zh/doc-detail/29807.htm)
	//
	// example:
	//
	// default
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// The host record.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The time to live (TTL) in the local cache, in seconds. Valid values are 5, 30, 60, 3600 (1 hour), 43200 (12 hours), and 86400 (1 day). The default value is 60.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. The following types are supported:A: Maps a domain name to an IPv4 address.AAAA: Maps a domain name to an IPv6 address.CNAME: An alias record that points a domain name to another domain name.MX: A mail exchanger record that points a domain name to a mail server.TXT: A text record that contains arbitrary human-readable text.SRV: A service record that identifies a server for a specific service. This is common in directory management for Microsoft systems.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The record value.
	//
	// example:
	//
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the DNS record. The value ranges from 0 to 100.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The zone ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 169438909000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SearchRecursionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionRecordsRequest) GoString() string {
	return s.String()
}

func (s *SearchRecursionRecordsRequest) GetDirection() *string {
	return s.Direction
}

func (s *SearchRecursionRecordsRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchRecursionRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchRecursionRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchRecursionRecordsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchRecursionRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRecursionRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRecursionRecordsRequest) GetRemark() *string {
	return s.Remark
}

func (s *SearchRecursionRecordsRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *SearchRecursionRecordsRequest) GetRr() *string {
	return s.Rr
}

func (s *SearchRecursionRecordsRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *SearchRecursionRecordsRequest) GetType() *string {
	return s.Type
}

func (s *SearchRecursionRecordsRequest) GetValue() *string {
	return s.Value
}

func (s *SearchRecursionRecordsRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *SearchRecursionRecordsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *SearchRecursionRecordsRequest) SetDirection(v string) *SearchRecursionRecordsRequest {
	s.Direction = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetEnableStatus(v string) *SearchRecursionRecordsRequest {
	s.EnableStatus = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetMaxResults(v int32) *SearchRecursionRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetNextToken(v string) *SearchRecursionRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetOrderBy(v string) *SearchRecursionRecordsRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetPageNumber(v int32) *SearchRecursionRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetPageSize(v int32) *SearchRecursionRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetRemark(v string) *SearchRecursionRecordsRequest {
	s.Remark = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetRequestSource(v string) *SearchRecursionRecordsRequest {
	s.RequestSource = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetRr(v string) *SearchRecursionRecordsRequest {
	s.Rr = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetTtl(v int32) *SearchRecursionRecordsRequest {
	s.Ttl = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetType(v string) *SearchRecursionRecordsRequest {
	s.Type = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetValue(v string) *SearchRecursionRecordsRequest {
	s.Value = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetWeight(v int32) *SearchRecursionRecordsRequest {
	s.Weight = &v
	return s
}

func (s *SearchRecursionRecordsRequest) SetZoneId(v string) *SearchRecursionRecordsRequest {
	s.ZoneId = &v
	return s
}

func (s *SearchRecursionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
