// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnBandWidthDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *DescribeEpnBandWidthDataRequest
	GetEPNInstanceId() *string
	SetEndTime(v string) *DescribeEpnBandWidthDataRequest
	GetEndTime() *string
	SetEnsRegionId(v string) *DescribeEpnBandWidthDataRequest
	GetEnsRegionId() *string
	SetInstanceId(v string) *DescribeEpnBandWidthDataRequest
	GetInstanceId() *string
	SetIsp(v string) *DescribeEpnBandWidthDataRequest
	GetIsp() *string
	SetNetworkingModel(v string) *DescribeEpnBandWidthDataRequest
	GetNetworkingModel() *string
	SetPeriod(v string) *DescribeEpnBandWidthDataRequest
	GetPeriod() *string
	SetStartTime(v string) *DescribeEpnBandWidthDataRequest
	GetStartTime() *string
}

type DescribeEpnBandWidthDataRequest struct {
	// The ID of the EPN instance.
	//
	// example:
	//
	// epn-20200825134537VyK81T
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
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
	// 2021-12-16T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5sg1owx0g4ojy66ab2tez77r2
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
	// The networking mode. Valid values:
	//
	// 	- **SpeedUp**: intelligent acceleration network (Internet)
	//
	// 	- **Connection**: internal network
	//
	// 	- **SpeedUpAndConnection**: intelligent acceleration network and internal network
	//
	// example:
	//
	// SpeedUp
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
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
	// 2021-12-15T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEpnBandWidthDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnBandWidthDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandWidthDataRequest) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *DescribeEpnBandWidthDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEpnBandWidthDataRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEpnBandWidthDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEpnBandWidthDataRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeEpnBandWidthDataRequest) GetNetworkingModel() *string {
	return s.NetworkingModel
}

func (s *DescribeEpnBandWidthDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeEpnBandWidthDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEpnBandWidthDataRequest) SetEPNInstanceId(v string) *DescribeEpnBandWidthDataRequest {
	s.EPNInstanceId = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetEndTime(v string) *DescribeEpnBandWidthDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetEnsRegionId(v string) *DescribeEpnBandWidthDataRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetInstanceId(v string) *DescribeEpnBandWidthDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetIsp(v string) *DescribeEpnBandWidthDataRequest {
	s.Isp = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetNetworkingModel(v string) *DescribeEpnBandWidthDataRequest {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetPeriod(v string) *DescribeEpnBandWidthDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) SetStartTime(v string) *DescribeEpnBandWidthDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnBandWidthDataRequest) Validate() error {
	return dara.Validate(s)
}
