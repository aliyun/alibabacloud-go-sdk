// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerMetricDataDemoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodPlayerMetricDataDemoRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodPlayerMetricDataDemoRequest
	GetEndTime() *string
	SetFilters(v string) *DescribeVodPlayerMetricDataDemoRequest
	GetFilters() *string
	SetInterval(v string) *DescribeVodPlayerMetricDataDemoRequest
	GetInterval() *string
	SetMetrics(v string) *DescribeVodPlayerMetricDataDemoRequest
	GetMetrics() *string
	SetOs(v string) *DescribeVodPlayerMetricDataDemoRequest
	GetOs() *string
	SetPageNumber(v int64) *DescribeVodPlayerMetricDataDemoRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodPlayerMetricDataDemoRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeVodPlayerMetricDataDemoRequest
	GetStartTime() *string
	SetTerminalType(v string) *DescribeVodPlayerMetricDataDemoRequest
	GetTerminalType() *string
	SetTop(v int64) *DescribeVodPlayerMetricDataDemoRequest
	GetTop() *int64
}

type DescribeVodPlayerMetricDataDemoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-05T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// [
	//
	//   {
	//
	//     "Field": "codec",
	//
	//     "Op": "=",
	//
	//     "Value": "h265#_#h264"
	//
	//   },
	//
	//   {
	//
	//     "Field": "os",
	//
	//     "Op": "=",
	//
	//     "Value": "Android#_#iOS"
	//
	//   }
	//
	// ]
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1d
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Vv,Uv,AvgPerVv
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// Android、iOS、Windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5000
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-24T00:55:06Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// web
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
	// example:
	//
	// 5
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DescribeVodPlayerMetricDataDemoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataDemoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetFilters() *string {
	return s.Filters
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetOs() *string {
	return s.Os
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribeVodPlayerMetricDataDemoRequest) GetTop() *int64 {
	return s.Top
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetAppId(v string) *DescribeVodPlayerMetricDataDemoRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetEndTime(v string) *DescribeVodPlayerMetricDataDemoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetFilters(v string) *DescribeVodPlayerMetricDataDemoRequest {
	s.Filters = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetInterval(v string) *DescribeVodPlayerMetricDataDemoRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetMetrics(v string) *DescribeVodPlayerMetricDataDemoRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetOs(v string) *DescribeVodPlayerMetricDataDemoRequest {
	s.Os = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetPageNumber(v int64) *DescribeVodPlayerMetricDataDemoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetPageSize(v int64) *DescribeVodPlayerMetricDataDemoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetStartTime(v string) *DescribeVodPlayerMetricDataDemoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetTerminalType(v string) *DescribeVodPlayerMetricDataDemoRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) SetTop(v int64) *DescribeVodPlayerMetricDataDemoRequest {
	s.Top = &v
	return s
}

func (s *DescribeVodPlayerMetricDataDemoRequest) Validate() error {
	return dara.Validate(s)
}
