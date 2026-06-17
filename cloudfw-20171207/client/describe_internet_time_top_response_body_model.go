// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTimeTopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataCount(v int32) *DescribeInternetTimeTopResponseBody
	GetDataCount() *int32
	SetDataList(v []*DescribeInternetTimeTopResponseBodyDataList) *DescribeInternetTimeTopResponseBody
	GetDataList() []*DescribeInternetTimeTopResponseBodyDataList
	SetRequestId(v string) *DescribeInternetTimeTopResponseBody
	GetRequestId() *string
	SetTrafficTime(v int32) *DescribeInternetTimeTopResponseBody
	GetTrafficTime() *int32
}

type DescribeInternetTimeTopResponseBody struct {
	// The number of entries returned.
	//
	// example:
	//
	// 19
	DataCount *int32 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// The list of data entries.
	DataList []*DescribeInternetTimeTopResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7F2D5C04-731F-50B0-ADE1-01637B3C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The timestamp of the traffic data. Unit: seconds.
	//
	// example:
	//
	// 1734399660
	TrafficTime *int32 `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
}

func (s DescribeInternetTimeTopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTimeTopResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetTimeTopResponseBody) GetDataCount() *int32 {
	return s.DataCount
}

func (s *DescribeInternetTimeTopResponseBody) GetDataList() []*DescribeInternetTimeTopResponseBodyDataList {
	return s.DataList
}

func (s *DescribeInternetTimeTopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInternetTimeTopResponseBody) GetTrafficTime() *int32 {
	return s.TrafficTime
}

func (s *DescribeInternetTimeTopResponseBody) SetDataCount(v int32) *DescribeInternetTimeTopResponseBody {
	s.DataCount = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBody) SetDataList(v []*DescribeInternetTimeTopResponseBodyDataList) *DescribeInternetTimeTopResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetTimeTopResponseBody) SetRequestId(v string) *DescribeInternetTimeTopResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBody) SetTrafficTime(v int32) *DescribeInternetTimeTopResponseBody {
	s.TrafficTime = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInternetTimeTopResponseBodyDataList struct {
	// The public IP address.
	//
	// example:
	//
	// 183.60.164.XXX
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 187
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The inbound packet forwarding rate. Unit: pps.
	//
	// example:
	//
	// 2
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-wz98eedr5l5hkb8****e7
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The name of the NAT gateway.
	//
	// example:
	//
	// ngw-test
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	// The public IP address of the NAT gateway.
	//
	// example:
	//
	// 47.97.66.XXX
	NatIP *string `json:"NatIP,omitempty" xml:"NatIP,omitempty"`
	// The number of new connections.
	//
	// example:
	//
	// 27
	NewConn *int64 `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// The outbound traffic. Unit: bit/s.
	//
	// example:
	//
	// 45
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The outbound packet forwarding rate. Unit: pps.
	//
	// example:
	//
	// 2
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 10.21.186.XXX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the asset instance.
	//
	// example:
	//
	// lb-bp14ue2rgktunncq****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// test
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// The type of the public IP address.
	//
	// example:
	//
	// EcsPublicIP
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of sessions.
	//
	// example:
	//
	// 27
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The total bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 232
	TotalBps *int64 `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
	// The total number of packets.
	//
	// example:
	//
	// 88
	TotalPps *int64 `json:"TotalPps,omitempty" xml:"TotalPps,omitempty"`
	// The ID of the virtual private cloud (VPC) instance.
	//
	// example:
	//
	// vpc-wz9o0uzfjuj81fx7m****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeInternetTimeTopResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTimeTopResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetIP() *string {
	return s.IP
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetInBps() *int64 {
	return s.InBps
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetInPps() *int64 {
	return s.InPps
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetNatIP() *string {
	return s.NatIP
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetNewConn() *int64 {
	return s.NewConn
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetOutBps() *int64 {
	return s.OutBps
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetOutPps() *int64 {
	return s.OutPps
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetTotalBps() *int64 {
	return s.TotalBps
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetTotalPps() *int64 {
	return s.TotalPps
}

func (s *DescribeInternetTimeTopResponseBodyDataList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetIP(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.IP = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetInBps(v int64) *DescribeInternetTimeTopResponseBodyDataList {
	s.InBps = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetInPps(v int64) *DescribeInternetTimeTopResponseBodyDataList {
	s.InPps = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetNatGatewayId(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetNatGatewayName(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetNatIP(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.NatIP = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetNewConn(v int64) *DescribeInternetTimeTopResponseBodyDataList {
	s.NewConn = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetOutBps(v int64) *DescribeInternetTimeTopResponseBodyDataList {
	s.OutBps = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetOutPps(v int64) *DescribeInternetTimeTopResponseBodyDataList {
	s.OutPps = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetPrivateIP(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.PrivateIP = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetRegionNo(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetResourceInstanceId(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetResourceInstanceName(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetResourceType(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.ResourceType = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetSessionCount(v int64) *DescribeInternetTimeTopResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetTotalBps(v int64) *DescribeInternetTimeTopResponseBodyDataList {
	s.TotalBps = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetTotalPps(v int64) *DescribeInternetTimeTopResponseBodyDataList {
	s.TotalPps = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) SetVpcId(v string) *DescribeInternetTimeTopResponseBodyDataList {
	s.VpcId = &v
	return s
}

func (s *DescribeInternetTimeTopResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
