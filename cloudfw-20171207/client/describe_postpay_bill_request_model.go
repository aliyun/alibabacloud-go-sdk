// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *DescribePostpayBillRequest
	GetCurrentPage() *int64
	SetEndTime(v string) *DescribePostpayBillRequest
	GetEndTime() *string
	SetInterval(v int32) *DescribePostpayBillRequest
	GetInterval() *int32
	SetLang(v string) *DescribePostpayBillRequest
	GetLang() *string
	SetPageSize(v int64) *DescribePostpayBillRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribePostpayBillRequest
	GetStartTime() *string
}

type DescribePostpayBillRequest struct {
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query, expressed as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1646063922
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval for querying data. This is an enumerated value. Valid values:
	//
	// - 3600: queries hourly data.
	//
	// - 86400: queries daily data.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The language. This is an enumerated value.
	//
	// Default value: zh.
	//
	// Valid values: en.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page in a paged query. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query, expressed as a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656664560
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePostpayBillRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayBillRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayBillRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribePostpayBillRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePostpayBillRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribePostpayBillRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePostpayBillRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePostpayBillRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePostpayBillRequest) SetCurrentPage(v int64) *DescribePostpayBillRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePostpayBillRequest) SetEndTime(v string) *DescribePostpayBillRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePostpayBillRequest) SetInterval(v int32) *DescribePostpayBillRequest {
	s.Interval = &v
	return s
}

func (s *DescribePostpayBillRequest) SetLang(v string) *DescribePostpayBillRequest {
	s.Lang = &v
	return s
}

func (s *DescribePostpayBillRequest) SetPageSize(v int64) *DescribePostpayBillRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePostpayBillRequest) SetStartTime(v string) *DescribePostpayBillRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePostpayBillRequest) Validate() error {
	return dara.Validate(s)
}
