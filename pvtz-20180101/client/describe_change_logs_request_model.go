// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChangeLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *DescribeChangeLogsRequest
	GetEndTimestamp() *int64
	SetEntityType(v string) *DescribeChangeLogsRequest
	GetEntityType() *string
	SetKeyword(v string) *DescribeChangeLogsRequest
	GetKeyword() *string
	SetLang(v string) *DescribeChangeLogsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeChangeLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeChangeLogsRequest
	GetPageSize() *int32
	SetStartTimestamp(v int64) *DescribeChangeLogsRequest
	GetStartTimestamp() *int64
	SetUserClientIp(v string) *DescribeChangeLogsRequest
	GetUserClientIp() *string
	SetZoneId(v string) *DescribeChangeLogsRequest
	GetZoneId() *string
}

type DescribeChangeLogsRequest struct {
	// The end of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516779348000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The type of operation logs. Valid values:
	//
	// 	- **PV_ZONE**: the logs that record the operations on built-in authoritative zones
	//
	// 	- **PV_RECORD**: the logs that record the operations on DNS records
	//
	// 	- **RESOLVER_RULE**: the logs that record the operations on forwarding rules
	//
	// 	- **CUSTOM_LINE**: the logs that record the operations on user-defined lines
	//
	// 	- **RESOLVER_ENDPOINT**: the logs that record the operations on outbound endpoints
	//
	// 	- **INBOUND_ENDPOINT**: the logs that record the operations on inbound endpoints
	//
	// 	- **CACHE_RESERVE_DOMAIN**: the logs that record the operations on cache retention domain names
	//
	// >  If you set EntityType to other values, all types of logs are queried.
	//
	// example:
	//
	// PV_ZONE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The keyword of the operation or the operation content. Fuzzy search is supported. The value is not case-sensitive.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Set the time to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1516779348000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The zone ID. Valid values:
	//
	// 	- If you set ZoneId to a zone ID, the logs that record the operations on the DNS records of the specified zone are queried.\\
	//
	// 	- If you leave ZoneId empty, the logs that record the operations on all zones and the DNS records of these zones that belong to the current Alibaba Cloud account are queried.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeChangeLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChangeLogsRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeChangeLogsRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeChangeLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeChangeLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeChangeLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeChangeLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeChangeLogsRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeChangeLogsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeChangeLogsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeChangeLogsRequest) SetEndTimestamp(v int64) *DescribeChangeLogsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetEntityType(v string) *DescribeChangeLogsRequest {
	s.EntityType = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetKeyword(v string) *DescribeChangeLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetLang(v string) *DescribeChangeLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetPageNumber(v int32) *DescribeChangeLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetPageSize(v int32) *DescribeChangeLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetStartTimestamp(v int64) *DescribeChangeLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetUserClientIp(v string) *DescribeChangeLogsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeChangeLogsRequest) SetZoneId(v string) *DescribeChangeLogsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeChangeLogsRequest) Validate() error {
	return dara.Validate(s)
}
