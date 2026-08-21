// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeZoneRecordsRequest
	GetKeyword() *string
	SetLang(v string) *DescribeZoneRecordsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeZoneRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeZoneRecordsRequest
	GetPageSize() *int32
	SetSearchMode(v string) *DescribeZoneRecordsRequest
	GetSearchMode() *string
	SetTag(v string) *DescribeZoneRecordsRequest
	GetTag() *string
	SetUserClientIp(v string) *DescribeZoneRecordsRequest
	GetUserClientIp() *string
	SetZoneId(v string) *DescribeZoneRecordsRequest
	GetZoneId() *string
}

type DescribeZoneRecordsRequest struct {
	// The keyword for the hostname. The search is not case-sensitive. Use the SearchMode parameter to switch between a fuzzy search and an exact search. The default is an exact search.
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
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search mode. Valid values:
	//
	// - **LIKE**: fuzzy search
	//
	// - **EXACT**: exact search (default)
	//
	// This parameter is used with the Keyword parameter.
	//
	// example:
	//
	// EXACT
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The tag. Valid values:
	//
	// - ecs: Queries the hostnames that are synchronized from ECS instances to the zone.
	//
	// - If you do not specify this parameter, all DNS records in the zone are queried.
	//
	// example:
	//
	// ecs
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The client IP address of the user.
	//
	// example:
	//
	// 127.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The unique ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// a96d70eb4ab8ef01503dc5486914****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZoneRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneRecordsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeZoneRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeZoneRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeZoneRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeZoneRecordsRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *DescribeZoneRecordsRequest) GetTag() *string {
	return s.Tag
}

func (s *DescribeZoneRecordsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeZoneRecordsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZoneRecordsRequest) SetKeyword(v string) *DescribeZoneRecordsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetLang(v string) *DescribeZoneRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetPageNumber(v int32) *DescribeZoneRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetPageSize(v int32) *DescribeZoneRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetSearchMode(v string) *DescribeZoneRecordsRequest {
	s.SearchMode = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetTag(v string) *DescribeZoneRecordsRequest {
	s.Tag = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetUserClientIp(v string) *DescribeZoneRecordsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeZoneRecordsRequest) SetZoneId(v string) *DescribeZoneRecordsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeZoneRecordsRequest) Validate() error {
	return dara.Validate(s)
}
