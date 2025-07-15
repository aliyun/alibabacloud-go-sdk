// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKVvDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKVvDataShrinkRequest
	GetDataInterval() *string
	SetDomainNameListShrink(v string) *DescribeRTSNativeSDKVvDataShrinkRequest
	GetDomainNameListShrink() *string
	SetEndTime(v string) *DescribeRTSNativeSDKVvDataShrinkRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeRTSNativeSDKVvDataShrinkRequest
	GetStartTime() *string
}

type DescribeRTSNativeSDKVvDataShrinkRequest struct {
	// The time granularity. Valid values: 300, 3600, 14400, 28800, and 86400. Unit: seconds. The default value is 300. If you specify an invalid value or do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The array of domain names.
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

func (s DescribeRTSNativeSDKVvDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKVvDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) GetDomainNameListShrink() *string {
	return s.DomainNameListShrink
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) SetDataInterval(v string) *DescribeRTSNativeSDKVvDataShrinkRequest {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) SetDomainNameListShrink(v string) *DescribeRTSNativeSDKVvDataShrinkRequest {
	s.DomainNameListShrink = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) SetEndTime(v string) *DescribeRTSNativeSDKVvDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) SetStartTime(v string) *DescribeRTSNativeSDKVvDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
