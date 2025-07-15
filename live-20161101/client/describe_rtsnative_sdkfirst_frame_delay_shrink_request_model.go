// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKFirstFrameDelayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest
	GetDataInterval() *string
	SetDomainNameListShrink(v string) *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest
	GetDomainNameListShrink() *string
	SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest
	GetStartTime() *string
}

type DescribeRTSNativeSDKFirstFrameDelayShrinkRequest struct {
	// The time granularity. Valid values: 300, 3600, 14400, 28800, and 86400. Unit: seconds. The default value is 300. If you specify an invalid value or do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// From V2.1.0, all domain names are queried by default. You can also specify specific domain names that you want to query. In this case, separate the domain names with commas (,). You can specify up to 500 domain names in each call.
	DomainNameListShrink *string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) GetDomainNameListShrink() *string {
	return s.DomainNameListShrink
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) SetDomainNameListShrink(v string) *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest {
	s.DomainNameListShrink = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayShrinkRequest) Validate() error {
	return dara.Validate(s)
}
