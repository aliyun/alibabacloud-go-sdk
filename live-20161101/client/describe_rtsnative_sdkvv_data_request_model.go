// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKVvDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKVvDataRequest
	GetDataInterval() *string
	SetDomainNameList(v []*string) *DescribeRTSNativeSDKVvDataRequest
	GetDomainNameList() []*string
	SetEndTime(v string) *DescribeRTSNativeSDKVvDataRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeRTSNativeSDKVvDataRequest
	GetStartTime() *string
}

type DescribeRTSNativeSDKVvDataRequest struct {
	// The time granularity. Valid values: 300, 3600, 14400, 28800, and 86400. Unit: seconds. The default value is 300. If you specify an invalid value or do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 300
	DataInterval *string `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	// The array of domain names.
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

func (s DescribeRTSNativeSDKVvDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKVvDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKVvDataRequest) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKVvDataRequest) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *DescribeRTSNativeSDKVvDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKVvDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKVvDataRequest) SetDataInterval(v string) *DescribeRTSNativeSDKVvDataRequest {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataRequest) SetDomainNameList(v []*string) *DescribeRTSNativeSDKVvDataRequest {
	s.DomainNameList = v
	return s
}

func (s *DescribeRTSNativeSDKVvDataRequest) SetEndTime(v string) *DescribeRTSNativeSDKVvDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataRequest) SetStartTime(v string) *DescribeRTSNativeSDKVvDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataRequest) Validate() error {
	return dara.Validate(s)
}
