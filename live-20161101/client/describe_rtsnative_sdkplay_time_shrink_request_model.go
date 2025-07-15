// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKPlayTimeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKPlayTimeShrinkRequest
	GetDataInterval() *string
	SetDomainNameListShrink(v string) *DescribeRTSNativeSDKPlayTimeShrinkRequest
	GetDomainNameListShrink() *string
	SetEndTime(v string) *DescribeRTSNativeSDKPlayTimeShrinkRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeRTSNativeSDKPlayTimeShrinkRequest
	GetStartTime() *string
}

type DescribeRTSNativeSDKPlayTimeShrinkRequest struct {
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

func (s DescribeRTSNativeSDKPlayTimeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayTimeShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) GetDomainNameListShrink() *string {
	return s.DomainNameListShrink
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) SetDataInterval(v string) *DescribeRTSNativeSDKPlayTimeShrinkRequest {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) SetDomainNameListShrink(v string) *DescribeRTSNativeSDKPlayTimeShrinkRequest {
	s.DomainNameListShrink = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) SetEndTime(v string) *DescribeRTSNativeSDKPlayTimeShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) SetStartTime(v string) *DescribeRTSNativeSDKPlayTimeShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
