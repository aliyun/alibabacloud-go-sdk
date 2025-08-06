// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodPlayerMetricDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodPlayerMetricDataRequest
	GetEndTime() *string
	SetFilters(v string) *DescribeVodPlayerMetricDataRequest
	GetFilters() *string
	SetInterval(v string) *DescribeVodPlayerMetricDataRequest
	GetInterval() *string
	SetLanguage(v string) *DescribeVodPlayerMetricDataRequest
	GetLanguage() *string
	SetMetrics(v string) *DescribeVodPlayerMetricDataRequest
	GetMetrics() *string
	SetOs(v string) *DescribeVodPlayerMetricDataRequest
	GetOs() *string
	SetPageNumber(v int64) *DescribeVodPlayerMetricDataRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodPlayerMetricDataRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeVodPlayerMetricDataRequest
	GetStartTime() *string
	SetTerminalType(v string) *DescribeVodPlayerMetricDataRequest
	GetTerminalType() *string
	SetTop(v int64) *DescribeVodPlayerMetricDataRequest
	GetTop() *int64
}

type DescribeVodPlayerMetricDataRequest struct {
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
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
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

func (s DescribeVodPlayerMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodPlayerMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodPlayerMetricDataRequest) GetFilters() *string {
	return s.Filters
}

func (s *DescribeVodPlayerMetricDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodPlayerMetricDataRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeVodPlayerMetricDataRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *DescribeVodPlayerMetricDataRequest) GetOs() *string {
	return s.Os
}

func (s *DescribeVodPlayerMetricDataRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodPlayerMetricDataRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodPlayerMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodPlayerMetricDataRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribeVodPlayerMetricDataRequest) GetTop() *int64 {
	return s.Top
}

func (s *DescribeVodPlayerMetricDataRequest) SetAppId(v string) *DescribeVodPlayerMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetEndTime(v string) *DescribeVodPlayerMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetFilters(v string) *DescribeVodPlayerMetricDataRequest {
	s.Filters = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetInterval(v string) *DescribeVodPlayerMetricDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetLanguage(v string) *DescribeVodPlayerMetricDataRequest {
	s.Language = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetMetrics(v string) *DescribeVodPlayerMetricDataRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetOs(v string) *DescribeVodPlayerMetricDataRequest {
	s.Os = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetPageNumber(v int64) *DescribeVodPlayerMetricDataRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetPageSize(v int64) *DescribeVodPlayerMetricDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetStartTime(v string) *DescribeVodPlayerMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetTerminalType(v string) *DescribeVodPlayerMetricDataRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetTop(v int64) *DescribeVodPlayerMetricDataRequest {
	s.Top = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
