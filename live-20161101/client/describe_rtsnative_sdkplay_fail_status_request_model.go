// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKPlayFailStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKPlayFailStatusRequest
	GetDataInterval() *string
	SetDomainNameList(v []*string) *DescribeRTSNativeSDKPlayFailStatusRequest
	GetDomainNameList() []*string
	SetEndTime(v string) *DescribeRTSNativeSDKPlayFailStatusRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeRTSNativeSDKPlayFailStatusRequest
	GetStartTime() *string
}

type DescribeRTSNativeSDKPlayFailStatusRequest struct {
	// The time granularity. Valid values: 300, 3600, 14400, 28800, and 86400. Unit: seconds. If this parameter is not specified or the specified value is not supported, the default value 300 is used.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// By default, all domain names are queried (version 2.1.0 and later). You can also specify domain names. Separate multiple domain names with commas (,). A maximum of 500 domain names can be queried at a time.
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
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

func (s DescribeRTSNativeSDKPlayFailStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayFailStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) SetDataInterval(v string) *DescribeRTSNativeSDKPlayFailStatusRequest {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) SetDomainNameList(v []*string) *DescribeRTSNativeSDKPlayFailStatusRequest {
	s.DomainNameList = v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) SetEndTime(v string) *DescribeRTSNativeSDKPlayFailStatusRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) SetStartTime(v string) *DescribeRTSNativeSDKPlayFailStatusRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusRequest) Validate() error {
	return dara.Validate(s)
}
