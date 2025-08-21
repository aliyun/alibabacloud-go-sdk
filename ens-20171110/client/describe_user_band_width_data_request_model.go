// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBandWidthDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeUserBandWidthDataRequest
	GetEndTime() *string
	SetEnsRegionId(v string) *DescribeUserBandWidthDataRequest
	GetEnsRegionId() *string
	SetInstanceId(v string) *DescribeUserBandWidthDataRequest
	GetInstanceId() *string
	SetIsp(v string) *DescribeUserBandWidthDataRequest
	GetIsp() *string
	SetPeriod(v string) *DescribeUserBandWidthDataRequest
	GetPeriod() *string
	SetStartTime(v string) *DescribeUserBandWidthDataRequest
	GetStartTime() *string
}

type DescribeUserBandWidthDataRequest struct {
	// The end of the time range to query.
	//
	// 	- Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- If the value of the seconds place is not 00, the start time is automatically set to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-21T12:22:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the node. You can specify only one node ID. By default, all nodes are queried.
	//
	// example:
	//
	// cn-taiyuan-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the instance for which you want to query the data. You can specify only one instance ID. By default, all instances are queried.
	//
	// example:
	//
	// i-5inkeimcipxk26yqtzm4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Internet service provider (ISP). Valid values:
	//
	// 	- cmcc: China Mobile
	//
	// 	- telecom: China Telecom
	//
	// 	- unicom: China Unicom
	//
	// 	- multiCarrier: multi-line ISP
	//
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The precision of the monitoring data that you want to obtain. Valid values: 300, 1200, 3600, and 14400. Default value: 300. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The beginning of the time range to query.
	//
	// 	- Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- If the value of the seconds place is not 00, the start time is automatically set to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-21T10:22:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserBandWidthDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBandWidthDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBandWidthDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUserBandWidthDataRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeUserBandWidthDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserBandWidthDataRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeUserBandWidthDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeUserBandWidthDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUserBandWidthDataRequest) SetEndTime(v string) *DescribeUserBandWidthDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetEnsRegionId(v string) *DescribeUserBandWidthDataRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetInstanceId(v string) *DescribeUserBandWidthDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetIsp(v string) *DescribeUserBandWidthDataRequest {
	s.Isp = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetPeriod(v string) *DescribeUserBandWidthDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) SetStartTime(v string) *DescribeUserBandWidthDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUserBandWidthDataRequest) Validate() error {
	return dara.Validate(s)
}
