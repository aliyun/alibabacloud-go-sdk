// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcChannelMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRtcChannelMetricRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeRtcChannelMetricRequest
	GetChannelId() *string
	SetOwnerId(v int64) *DescribeRtcChannelMetricRequest
	GetOwnerId() *int64
	SetTimePoint(v string) *DescribeRtcChannelMetricRequest
	GetTimePoint() *string
}

type DescribeRtcChannelMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aoe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeRtcChannelMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcChannelMetricRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRtcChannelMetricRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRtcChannelMetricRequest) GetTimePoint() *string {
	return s.TimePoint
}

func (s *DescribeRtcChannelMetricRequest) SetAppId(v string) *DescribeRtcChannelMetricRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetChannelId(v string) *DescribeRtcChannelMetricRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetOwnerId(v int64) *DescribeRtcChannelMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetTimePoint(v string) *DescribeRtcChannelMetricRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) Validate() error {
	return dara.Validate(s)
}
