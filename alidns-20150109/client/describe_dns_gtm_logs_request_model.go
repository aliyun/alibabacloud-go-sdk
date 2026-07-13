// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *DescribeDnsGtmLogsRequest
	GetEndTimestamp() *int64
	SetInstanceId(v string) *DescribeDnsGtmLogsRequest
	GetInstanceId() *string
	SetKeyword(v string) *DescribeDnsGtmLogsRequest
	GetKeyword() *string
	SetLang(v string) *DescribeDnsGtmLogsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeDnsGtmLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnsGtmLogsRequest
	GetPageSize() *int32
	SetStartTimestamp(v int64) *DescribeDnsGtmLogsRequest
	GetStartTimestamp() *int64
}

type DescribeDnsGtmLogsRequest struct {
	// The end of the time range to query. This is a UNIX timestamp.
	//
	// example:
	//
	// 1516779348000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The instance ID. Call the [DescribeDnsGtmInstances](https://www.alibabacloud.com/help/en/dns/api-alidns-2015-01-09-describednsgtminstances?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The keyword. The search is performed in the \\`%KeyWord%\\` pattern and is not case-sensitive.
	//
	// example:
	//
	// demo
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language of some returned parameters. The default value is en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. The value starts from 1. The default value is 1.
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
	// The start of the time range to query. This is a UNIX timestamp.
	//
	// example:
	//
	// 1516779348000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeDnsGtmLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmLogsRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeDnsGtmLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDnsGtmLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDnsGtmLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnsGtmLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnsGtmLogsRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeDnsGtmLogsRequest) SetEndTimestamp(v int64) *DescribeDnsGtmLogsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetInstanceId(v string) *DescribeDnsGtmLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetKeyword(v string) *DescribeDnsGtmLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetLang(v string) *DescribeDnsGtmLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetPageNumber(v int32) *DescribeDnsGtmLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetPageSize(v int32) *DescribeDnsGtmLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) SetStartTimestamp(v int64) *DescribeDnsGtmLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeDnsGtmLogsRequest) Validate() error {
	return dara.Validate(s)
}
