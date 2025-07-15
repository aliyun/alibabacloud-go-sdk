// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKPlayTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKPlayTimeRequest
	GetDataInterval() *string
	SetDomainNameList(v []*string) *DescribeRTSNativeSDKPlayTimeRequest
	GetDomainNameList() []*string
	SetEndTime(v string) *DescribeRTSNativeSDKPlayTimeRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeRTSNativeSDKPlayTimeRequest
	GetStartTime() *string
}

type DescribeRTSNativeSDKPlayTimeRequest struct {
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

func (s DescribeRTSNativeSDKPlayTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) SetDataInterval(v string) *DescribeRTSNativeSDKPlayTimeRequest {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) SetDomainNameList(v []*string) *DescribeRTSNativeSDKPlayTimeRequest {
	s.DomainNameList = v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) SetEndTime(v string) *DescribeRTSNativeSDKPlayTimeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) SetStartTime(v string) *DescribeRTSNativeSDKPlayTimeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeRequest) Validate() error {
	return dara.Validate(s)
}
