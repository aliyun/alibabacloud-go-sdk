// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerCollectDataDemoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodPlayerCollectDataDemoRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodPlayerCollectDataDemoRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodPlayerCollectDataDemoRequest
	GetInterval() *string
	SetMetrics(v string) *DescribeVodPlayerCollectDataDemoRequest
	GetMetrics() *string
	SetOs(v string) *DescribeVodPlayerCollectDataDemoRequest
	GetOs() *string
	SetPeriod(v string) *DescribeVodPlayerCollectDataDemoRequest
	GetPeriod() *string
	SetStartTime(v string) *DescribeVodPlayerCollectDataDemoRequest
	GetStartTime() *string
	SetTerminalType(v string) *DescribeVodPlayerCollectDataDemoRequest
	GetTerminalType() *string
}

type DescribeVodPlayerCollectDataDemoRequest struct {
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
	// 1d
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
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
	// Web
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribeVodPlayerCollectDataDemoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataDemoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataDemoRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodPlayerCollectDataDemoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodPlayerCollectDataDemoRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodPlayerCollectDataDemoRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *DescribeVodPlayerCollectDataDemoRequest) GetOs() *string {
	return s.Os
}

func (s *DescribeVodPlayerCollectDataDemoRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeVodPlayerCollectDataDemoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodPlayerCollectDataDemoRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribeVodPlayerCollectDataDemoRequest) SetAppId(v string) *DescribeVodPlayerCollectDataDemoRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoRequest) SetEndTime(v string) *DescribeVodPlayerCollectDataDemoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoRequest) SetInterval(v string) *DescribeVodPlayerCollectDataDemoRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoRequest) SetMetrics(v string) *DescribeVodPlayerCollectDataDemoRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoRequest) SetOs(v string) *DescribeVodPlayerCollectDataDemoRequest {
	s.Os = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoRequest) SetPeriod(v string) *DescribeVodPlayerCollectDataDemoRequest {
	s.Period = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoRequest) SetStartTime(v string) *DescribeVodPlayerCollectDataDemoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoRequest) SetTerminalType(v string) *DescribeVodPlayerCollectDataDemoRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribeVodPlayerCollectDataDemoRequest) Validate() error {
	return dara.Validate(s)
}
