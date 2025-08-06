// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerCollectDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodPlayerCollectDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodPlayerCollectDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodPlayerCollectDataRequest
	GetInterval() *string
	SetMetrics(v string) *DescribeVodPlayerCollectDataRequest
	GetMetrics() *string
	SetOs(v string) *DescribeVodPlayerCollectDataRequest
	GetOs() *string
	SetPeriod(v string) *DescribeVodPlayerCollectDataRequest
	GetPeriod() *string
	SetStartTime(v string) *DescribeVodPlayerCollectDataRequest
	GetStartTime() *string
	SetTerminalType(v string) *DescribeVodPlayerCollectDataRequest
	GetTerminalType() *string
}

type DescribeVodPlayerCollectDataRequest struct {
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
	// web
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribeVodPlayerCollectDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodPlayerCollectDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodPlayerCollectDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodPlayerCollectDataRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *DescribeVodPlayerCollectDataRequest) GetOs() *string {
	return s.Os
}

func (s *DescribeVodPlayerCollectDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeVodPlayerCollectDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodPlayerCollectDataRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribeVodPlayerCollectDataRequest) SetAppId(v string) *DescribeVodPlayerCollectDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetEndTime(v string) *DescribeVodPlayerCollectDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetInterval(v string) *DescribeVodPlayerCollectDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetMetrics(v string) *DescribeVodPlayerCollectDataRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetOs(v string) *DescribeVodPlayerCollectDataRequest {
	s.Os = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetPeriod(v string) *DescribeVodPlayerCollectDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetStartTime(v string) *DescribeVodPlayerCollectDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetTerminalType(v string) *DescribeVodPlayerCollectDataRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) Validate() error {
	return dara.Validate(s)
}
