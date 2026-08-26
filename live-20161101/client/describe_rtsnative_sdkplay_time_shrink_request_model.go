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
	// The time granularity. Valid values: 300, 3600, 14400, 28800, and 86400. Unit: seconds. If this parameter is not specified or the specified value is not supported, the default value 300 is used.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The list of domain names to query. By default, all domain names are queried (version 2.1.0 and later). You can also specify domain names. Separate multiple domain names with commas (,). A maximum of 500 domain names can be queried at a time.
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
