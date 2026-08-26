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
	// The time granularity. Valid values: 300, 3600, 14400, 28800, and 86400. Unit: seconds. If you do not specify this parameter or specify an unsupported value, the default value 300 is used.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The domain names to query. By default, all domain names are queried (version 2.1.0 and later). You can also specify domain names. Separate multiple domain names with commas (,). You can specify up to 500 domain names at a time.
	DomainNameListShrink *string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty"`
	// The end time. This parameter is required. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-12-10T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time. This parameter is required. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
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
