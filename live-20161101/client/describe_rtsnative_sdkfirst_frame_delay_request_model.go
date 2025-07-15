// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKFirstFrameDelayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameDelayRequest
	GetDataInterval() *string
	SetDomainNameList(v []*string) *DescribeRTSNativeSDKFirstFrameDelayRequest
	GetDomainNameList() []*string
	SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameDelayRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameDelayRequest
	GetStartTime() *string
}

type DescribeRTSNativeSDKFirstFrameDelayRequest struct {
	// The time granularity. Valid values: 300, 3600, 14400, 28800, and 86400. Unit: seconds. The default value is 300. If you specify an invalid value or do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// From V2.1.0, all domain names are queried by default. You can also specify specific domain names that you want to query. In this case, separate the domain names with commas (,). You can specify up to 500 domain names in each call.
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
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

func (s DescribeRTSNativeSDKFirstFrameDelayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameDelayRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameDelayRequest {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) SetDomainNameList(v []*string) *DescribeRTSNativeSDKFirstFrameDelayRequest {
	s.DomainNameList = v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameDelayRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameDelayRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayRequest) Validate() error {
	return dara.Validate(s)
}
