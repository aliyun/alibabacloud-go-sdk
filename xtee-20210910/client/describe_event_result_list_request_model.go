// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventResultListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventResultListRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeEventResultListRequest
	GetBeginTime() *int64
	SetCurrentPage(v int64) *DescribeEventResultListRequest
	GetCurrentPage() *int64
	SetEndTime(v int64) *DescribeEventResultListRequest
	GetEndTime() *int64
	SetPageSize(v int64) *DescribeEventResultListRequest
	GetPageSize() *int64
	SetRegId(v string) *DescribeEventResultListRequest
	GetRegId() *string
}

type DescribeEventResultListRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start time, in milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1737101348000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The end time, in milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1683616457000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeEventResultListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventResultListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventResultListRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeEventResultListRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeEventResultListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeEventResultListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventResultListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventResultListRequest) SetLang(v string) *DescribeEventResultListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventResultListRequest) SetBeginTime(v int64) *DescribeEventResultListRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeEventResultListRequest) SetCurrentPage(v int64) *DescribeEventResultListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEventResultListRequest) SetEndTime(v int64) *DescribeEventResultListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventResultListRequest) SetPageSize(v int64) *DescribeEventResultListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventResultListRequest) SetRegId(v string) *DescribeEventResultListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventResultListRequest) Validate() error {
	return dara.Validate(s)
}
