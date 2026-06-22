// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProxyList(v []*DescribeSecurityProxyResponseBodyProxyList) *DescribeSecurityProxyResponseBody
	GetProxyList() []*DescribeSecurityProxyResponseBodyProxyList
	SetRequestId(v string) *DescribeSecurityProxyResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSecurityProxyResponseBody
	GetTotalCount() *int32
}

type DescribeSecurityProxyResponseBody struct {
	// The list of NAT firewalls.
	ProxyList []*DescribeSecurityProxyResponseBodyProxyList `json:"ProxyList,omitempty" xml:"ProxyList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F0F82705-CFC7-5F83-86C8-A063892F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecurityProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResponseBody) GetProxyList() []*DescribeSecurityProxyResponseBodyProxyList {
	return s.ProxyList
}

func (s *DescribeSecurityProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityProxyResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecurityProxyResponseBody) SetProxyList(v []*DescribeSecurityProxyResponseBodyProxyList) *DescribeSecurityProxyResponseBody {
	s.ProxyList = v
	return s
}

func (s *DescribeSecurityProxyResponseBody) SetRequestId(v string) *DescribeSecurityProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityProxyResponseBody) SetTotalCount(v int32) *DescribeSecurityProxyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityProxyResponseBody) Validate() error {
	if s.ProxyList != nil {
		for _, item := range s.ProxyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityProxyResponseBodyProxyList struct {
	// The IPv4 CIDR block of the VPC.
	//
	// example:
	//
	// 192.168.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The error details.
	//
	// example:
	//
	// “”
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The UID of the Cloud Firewall member accounts.
	//
	// example:
	//
	// 1797733170015112
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-2zex8sf4s5vus8rq3rjqo
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The name of the NAT gateway.
	//
	// example:
	//
	// ecs-slb-eip-waf
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	// The ID of the NAT firewall.
	//
	// example:
	//
	// proxy-nat4921f192b6cf438d93f8
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The name of the NAT firewall.
	//
	// example:
	//
	// nat-idmp-fir
	ProxyName *string `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	// The region ID of the VPC.
	//
	// example:
	//
	// ap-southeast-1
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The list of SNAT IP addresses.
	SnatIpList []*string `json:"SnatIpList,omitempty" xml:"SnatIpList,omitempty" type:"Repeated"`
	// The status of Cloud Firewall. Valid values:
	//
	// - **configuring**: Being created.
	//
	// - **deleting**: Being deleted.
	//
	// - **normal**: Normal.
	//
	// - **abnormal**: Abnormal.
	//
	// - **opening**: Being enabled.
	//
	// - **closing**: Being disabled.
	//
	// - **closed**: Disabled.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether strict mode is enabled. Valid values:
	//
	// - 1: Strict mode is enabled.
	//
	// - 0: Strict mode is disabled.
	//
	// example:
	//
	// 0
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// The ID of the vSwitch to which the Cloud Firewall interface belongs.
	//
	// example:
	//
	// vsw-5gu2qqfmjmwl8ktzgfekl
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC-connected instance ID.
	//
	// example:
	//
	// vpc-wz9xn35tq33hunzvpu0se
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC instance.
	//
	// example:
	//
	// vpc-bp1kw9igsq0yyzeanqamx
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeSecurityProxyResponseBodyProxyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResponseBodyProxyList) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetDetail() *string {
	return s.Detail
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetProxyName() *string {
	return s.ProxyName
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetSnatIpList() []*string {
	return s.SnatIpList
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetStatus() *string {
	return s.Status
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSecurityProxyResponseBodyProxyList) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetCidrBlock(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetDetail(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.Detail = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetMemberUid(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.MemberUid = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetNatGatewayId(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetNatGatewayName(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetProxyId(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.ProxyId = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetProxyName(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.ProxyName = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetRegionNo(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.RegionNo = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetSnatIpList(v []*string) *DescribeSecurityProxyResponseBodyProxyList {
	s.SnatIpList = v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetStatus(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.Status = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetStrictMode(v int32) *DescribeSecurityProxyResponseBodyProxyList {
	s.StrictMode = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetVSwitchId(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.VSwitchId = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetVpcId(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.VpcId = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) SetVpcName(v string) *DescribeSecurityProxyResponseBodyProxyList {
	s.VpcName = &v
	return s
}

func (s *DescribeSecurityProxyResponseBodyProxyList) Validate() error {
	return dara.Validate(s)
}
