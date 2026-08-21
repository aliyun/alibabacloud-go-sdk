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
	// The end time. This value is a UNIX timestamp.
	//
	// example:
	//
	// 2516779348000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The type of log to obtain. Valid values:
	//
	// - **PV_ZONE**: operation logs of built-in authoritative zones.
	//
	// - **PV_RECORD**: operation logs of DNS records.
	//
	// - **RESOLVER_RULE**: operation logs of forwarding rules.
	//
	// - **CUSTOM_LINE**: operation logs of custom lines.
	//
	// - **RESOLVER_ENDPOINT**: operation logs of outbound endpoints.
	//
	// - **INBOUND_ENDPOINT**: operation logs of inbound endpoints.
	//
	// - **CACHE_RESERVE_DOMAIN**: operation logs of domains for which cache is retained.
	//
	// > If you specify another value, this parameter is ignored and logs of all types are returned.
	//
	// example:
	//
	// PV_ZONE
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The keyword, such as a behavior or content. Fuzzy search is supported. The keyword is not case-sensitive.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// Default value: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1516779348000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The IP address of the user.
	//
	// example:
	//
	// 192.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The ID of the zone.
	//
	// - If you specify this parameter, the operation returns the change logs of DNS records for the specified zone.<br>
	//
	// - If you leave this parameter empty, the operation returns the change logs of all zones and DNS records that belong to the current account.
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
