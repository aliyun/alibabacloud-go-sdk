// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveInteractionMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeLiveInteractionMetricDataRequest
	GetAppId() *string
	SetBeginTs(v int64) *DescribeLiveInteractionMetricDataRequest
	GetBeginTs() *int64
	SetEndTs(v int64) *DescribeLiveInteractionMetricDataRequest
	GetEndTs() *int64
	SetMetricType(v string) *DescribeLiveInteractionMetricDataRequest
	GetMetricType() *string
	SetOs(v string) *DescribeLiveInteractionMetricDataRequest
	GetOs() *string
	SetTerminalType(v string) *DescribeLiveInteractionMetricDataRequest
	GetTerminalType() *string
}

type DescribeLiveInteractionMetricDataRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// e4d7f08a-01fe-41b5-a091-fe41060a****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1698195600000
	BeginTs *int64 `json:"BeginTs,omitempty" xml:"BeginTs,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1698201013000
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The metric. Valid values:
	//
	// 	- JoinChannelSucRate: the success rate of joining a channel within 5 seconds.
	//
	// 	- VideoStuckRate: the video stuttering rate.
	//
	// 	- AudioStuckRate: the audio stuttering rate.
	//
	// 	- FirstFrameCost: the time to first frame.
	//
	// This parameter is required.
	//
	// example:
	//
	// FirstFrameCost
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The operating system. Valid values: iOS and Android.
	//
	// example:
	//
	// Android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The terminal type. Valid values: web and mobile.
	//
	// example:
	//
	// mobile
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribeLiveInteractionMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveInteractionMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveInteractionMetricDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeLiveInteractionMetricDataRequest) GetBeginTs() *int64 {
	return s.BeginTs
}

func (s *DescribeLiveInteractionMetricDataRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeLiveInteractionMetricDataRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeLiveInteractionMetricDataRequest) GetOs() *string {
	return s.Os
}

func (s *DescribeLiveInteractionMetricDataRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribeLiveInteractionMetricDataRequest) SetAppId(v string) *DescribeLiveInteractionMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataRequest) SetBeginTs(v int64) *DescribeLiveInteractionMetricDataRequest {
	s.BeginTs = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataRequest) SetEndTs(v int64) *DescribeLiveInteractionMetricDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataRequest) SetMetricType(v string) *DescribeLiveInteractionMetricDataRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataRequest) SetOs(v string) *DescribeLiveInteractionMetricDataRequest {
	s.Os = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataRequest) SetTerminalType(v string) *DescribeLiveInteractionMetricDataRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribeLiveInteractionMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
