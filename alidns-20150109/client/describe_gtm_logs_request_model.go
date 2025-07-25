// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *DescribeGtmLogsRequest
	GetEndTimestamp() *int64
	SetInstanceId(v string) *DescribeGtmLogsRequest
	GetInstanceId() *string
	SetKeyword(v string) *DescribeGtmLogsRequest
	GetKeyword() *string
	SetLang(v string) *DescribeGtmLogsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeGtmLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmLogsRequest
	GetPageSize() *int32
	SetStartTimestamp(v int64) *DescribeGtmLogsRequest
	GetStartTimestamp() *int64
}

type DescribeGtmLogsRequest struct {
	// The timestamp that specifies the end of the time range to query.
	//
	// example:
	//
	// 1363453350000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the GTM instance whose logs you want to query.
	//
	// example:
	//
	// gtm-cn-xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The keyword for searching logs, in case-insensitive "%Keyword%" format.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 1363453340000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribeGtmLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmLogsRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeGtmLogsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeGtmLogsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmLogsRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeGtmLogsRequest) SetEndTimestamp(v int64) *DescribeGtmLogsRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetInstanceId(v string) *DescribeGtmLogsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetKeyword(v string) *DescribeGtmLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetLang(v string) *DescribeGtmLogsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetPageNumber(v int32) *DescribeGtmLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetPageSize(v int32) *DescribeGtmLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmLogsRequest) SetStartTimestamp(v int64) *DescribeGtmLogsRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeGtmLogsRequest) Validate() error {
	return dara.Validate(s)
}
