// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnBandwitdhByInternetChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest
	GetEndTime() *string
	SetEnsRegionId(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest
	GetEnsRegionId() *string
	SetIsp(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest
	GetIsp() *string
	SetNetworkingModel(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest
	GetNetworkingModel() *string
	SetStartTime(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest
	GetStartTime() *string
}

type DescribeEpnBandwitdhByInternetChargeTypeRequest struct {
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
	// 2021-12-06T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// example:
	//
	// cn-changsha-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
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
	// telecom
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
	// Connection
	NetworkingModel *string `json:"NetworkingModel,omitempty" xml:"NetworkingModel,omitempty"`
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
	// 2021-12-02T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEpnBandwitdhByInternetChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnBandwitdhByInternetChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) GetNetworkingModel() *string {
	return s.NetworkingModel
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetEndTime(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetEnsRegionId(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetIsp(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.Isp = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetNetworkingModel(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.NetworkingModel = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) SetStartTime(v string) *DescribeEpnBandwitdhByInternetChargeTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEpnBandwitdhByInternetChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
