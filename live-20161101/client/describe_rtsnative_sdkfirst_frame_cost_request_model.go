// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKFirstFrameCostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameCostRequest
	GetDataInterval() *string
	SetDomainNameList(v []*string) *DescribeRTSNativeSDKFirstFrameCostRequest
	GetDomainNameList() []*string
	SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameCostRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameCostRequest
	GetStartTime() *string
}

type DescribeRTSNativeSDKFirstFrameCostRequest struct {
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

func (s DescribeRTSNativeSDKFirstFrameCostRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameCostRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) GetDataInterval() *string {
	return s.DataInterval
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) SetDataInterval(v string) *DescribeRTSNativeSDKFirstFrameCostRequest {
	s.DataInterval = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) SetDomainNameList(v []*string) *DescribeRTSNativeSDKFirstFrameCostRequest {
	s.DomainNameList = v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) SetEndTime(v string) *DescribeRTSNativeSDKFirstFrameCostRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) SetStartTime(v string) *DescribeRTSNativeSDKFirstFrameCostRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostRequest) Validate() error {
	return dara.Validate(s)
}
