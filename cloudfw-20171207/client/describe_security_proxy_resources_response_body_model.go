// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityProxyResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSecurityProxyResourcesResponseBody
	GetRequestId() *string
	SetResourceList(v []*DescribeSecurityProxyResourcesResponseBodyResourceList) *DescribeSecurityProxyResourcesResponseBody
	GetResourceList() []*DescribeSecurityProxyResourcesResponseBodyResourceList
}

type DescribeSecurityProxyResourcesResponseBody struct {
	// example:
	//
	// A8E8D50E-9F45-5662-B116-A1D0807F****
	RequestId    *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceList []*DescribeSecurityProxyResourcesResponseBodyResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
}

func (s DescribeSecurityProxyResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityProxyResourcesResponseBody) GetResourceList() []*DescribeSecurityProxyResourcesResponseBodyResourceList {
	return s.ResourceList
}

func (s *DescribeSecurityProxyResourcesResponseBody) SetRequestId(v string) *DescribeSecurityProxyResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBody) SetResourceList(v []*DescribeSecurityProxyResourcesResponseBodyResourceList) *DescribeSecurityProxyResourcesResponseBody {
	s.ResourceList = v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBody) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityProxyResourcesResponseBodyResourceList struct {
	// example:
	//
	// cn-beijing
	RegionNo *string                                                          `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	VpcList  []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
}

func (s DescribeSecurityProxyResourcesResponseBodyResourceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResourcesResponseBodyResourceList) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceList) GetVpcList() []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcList {
	return s.VpcList
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceList) SetRegionNo(v string) *DescribeSecurityProxyResourcesResponseBodyResourceList {
	s.RegionNo = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceList) SetVpcList(v []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) *DescribeSecurityProxyResourcesResponseBodyResourceList {
	s.VpcList = v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceList) Validate() error {
	if s.VpcList != nil {
		for _, item := range s.VpcList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityProxyResourcesResponseBodyResourceListVpcList struct {
	// example:
	//
	// 192.168.0.0/16
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// 157862808111****
	MemberUid   *string                                                                     `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	NatGateways []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways `json:"NatGateways,omitempty" xml:"NatGateways,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-8vbuzirdl3w1r7exw****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vpc-wz94a4q37rgl7g****
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) GetNatGateways() []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways {
	return s.NatGateways
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) SetCidrBlock(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList {
	s.CidrBlock = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) SetMemberUid(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList {
	s.MemberUid = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) SetNatGateways(v []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList {
	s.NatGateways = v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) SetVpcId(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList {
	s.VpcId = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) SetVpcName(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList {
	s.VpcName = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcList) Validate() error {
	if s.NatGateways != nil {
		for _, item := range s.NatGateways {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways struct {
	// example:
	//
	// Available
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// ngw-bp1bm0k2t8i6ooxu****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// ngw-test
	NatGatewayName    *string                                                                                      `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	NatRouteEntryList []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList `json:"NatRouteEntryList,omitempty" xml:"NatRouteEntryList,omitempty" type:"Repeated"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) GetDetail() *string {
	return s.Detail
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) GetNatRouteEntryList() []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList {
	return s.NatRouteEntryList
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) GetStatus() *string {
	return s.Status
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) SetDetail(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways {
	s.Detail = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) SetNatGatewayId(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) SetNatGatewayName(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) SetNatRouteEntryList(v []*DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways {
	s.NatRouteEntryList = v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) SetStatus(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways {
	s.Status = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGateways) Validate() error {
	if s.NatRouteEntryList != nil {
		for _, item := range s.NatRouteEntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList struct {
	// example:
	//
	// 10.0.70.XX/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// example:
	//
	// ngw-2zey0w2u02u1x584m****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// example:
	//
	// NatGateway
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// example:
	//
	// vtb-2ze409pp09d994a****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) SetDestinationCidr(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) SetNextHopId(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList {
	s.NextHopId = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) SetNextHopType(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList {
	s.NextHopType = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) SetRouteTableId(v string) *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponseBodyResourceListVpcListNatGatewaysNatRouteEntryList) Validate() error {
	return dara.Validate(s)
}
