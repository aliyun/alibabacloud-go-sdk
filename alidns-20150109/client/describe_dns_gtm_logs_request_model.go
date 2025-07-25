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
	// The timestamp that specifies the end of the time range to query.
	//
	// example:
	//
	// 1516779348000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The keyword for searches in "%KeyWord%" mode. The value is not case-sensitive.
	//
	// example:
	//
	// demo
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The language to return some response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The timestamp that specifies the beginning of the time range to query.
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
