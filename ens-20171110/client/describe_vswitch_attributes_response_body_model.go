// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableIpAddressCount(v int64) *DescribeVSwitchAttributesResponseBody
	GetAvailableIpAddressCount() *int64
	SetCidrBlock(v string) *DescribeVSwitchAttributesResponseBody
	GetCidrBlock() *string
	SetCreatedTime(v string) *DescribeVSwitchAttributesResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeVSwitchAttributesResponseBody
	GetDescription() *string
	SetEnsRegionId(v string) *DescribeVSwitchAttributesResponseBody
	GetEnsRegionId() *string
	SetHaVipIds(v *DescribeVSwitchAttributesResponseBodyHaVipIds) *DescribeVSwitchAttributesResponseBody
	GetHaVipIds() *DescribeVSwitchAttributesResponseBodyHaVipIds
	SetInstanceIds(v *DescribeVSwitchAttributesResponseBodyInstanceIds) *DescribeVSwitchAttributesResponseBody
	GetInstanceIds() *DescribeVSwitchAttributesResponseBodyInstanceIds
	SetLoadBalancerIds(v *DescribeVSwitchAttributesResponseBodyLoadBalancerIds) *DescribeVSwitchAttributesResponseBody
	GetLoadBalancerIds() *DescribeVSwitchAttributesResponseBodyLoadBalancerIds
	SetNatGatewayIds(v *DescribeVSwitchAttributesResponseBodyNatGatewayIds) *DescribeVSwitchAttributesResponseBody
	GetNatGatewayIds() *DescribeVSwitchAttributesResponseBodyNatGatewayIds
	SetNetworkId(v string) *DescribeVSwitchAttributesResponseBody
	GetNetworkId() *string
	SetNetworkInterfaceIds(v *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds) *DescribeVSwitchAttributesResponseBody
	GetNetworkInterfaceIds() *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds
	SetRequestId(v string) *DescribeVSwitchAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeVSwitchAttributesResponseBody
	GetStatus() *string
	SetVSwitchId(v string) *DescribeVSwitchAttributesResponseBody
	GetVSwitchId() *string
	SetVSwitchName(v string) *DescribeVSwitchAttributesResponseBody
	GetVSwitchName() *string
}

type DescribeVSwitchAttributesResponseBody struct {
	// example:
	//
	// 253
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// example:
	//
	// 10.0.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// example:
	//
	// 2019-06-01T00:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// This is my vswitch.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// cn-xian-unicom
	EnsRegionId     *string                                               `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	HaVipIds        *DescribeVSwitchAttributesResponseBodyHaVipIds        `json:"HaVipIds,omitempty" xml:"HaVipIds,omitempty" type:"Struct"`
	InstanceIds     *DescribeVSwitchAttributesResponseBodyInstanceIds     `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	LoadBalancerIds *DescribeVSwitchAttributesResponseBodyLoadBalancerIds `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Struct"`
	NatGatewayIds   *DescribeVSwitchAttributesResponseBodyNatGatewayIds   `json:"NatGatewayIds,omitempty" xml:"NatGatewayIds,omitempty" type:"Struct"`
	// example:
	//
	// n-257gqcdfvx6n****
	NetworkId           *string                                                   `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	NetworkInterfaceIds *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Struct"`
	// example:
	//
	// C0003****2A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// Test-switch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s DescribeVSwitchAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBody) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeVSwitchAttributesResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVSwitchAttributesResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeVSwitchAttributesResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeVSwitchAttributesResponseBody) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeVSwitchAttributesResponseBody) GetHaVipIds() *DescribeVSwitchAttributesResponseBodyHaVipIds {
	return s.HaVipIds
}

func (s *DescribeVSwitchAttributesResponseBody) GetInstanceIds() *DescribeVSwitchAttributesResponseBodyInstanceIds {
	return s.InstanceIds
}

func (s *DescribeVSwitchAttributesResponseBody) GetLoadBalancerIds() *DescribeVSwitchAttributesResponseBodyLoadBalancerIds {
	return s.LoadBalancerIds
}

func (s *DescribeVSwitchAttributesResponseBody) GetNatGatewayIds() *DescribeVSwitchAttributesResponseBodyNatGatewayIds {
	return s.NatGatewayIds
}

func (s *DescribeVSwitchAttributesResponseBody) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeVSwitchAttributesResponseBody) GetNetworkInterfaceIds() *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds {
	return s.NetworkInterfaceIds
}

func (s *DescribeVSwitchAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVSwitchAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVSwitchAttributesResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchAttributesResponseBody) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeVSwitchAttributesResponseBody) SetAvailableIpAddressCount(v int64) *DescribeVSwitchAttributesResponseBody {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetCidrBlock(v string) *DescribeVSwitchAttributesResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetCreatedTime(v string) *DescribeVSwitchAttributesResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetDescription(v string) *DescribeVSwitchAttributesResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetEnsRegionId(v string) *DescribeVSwitchAttributesResponseBody {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetHaVipIds(v *DescribeVSwitchAttributesResponseBodyHaVipIds) *DescribeVSwitchAttributesResponseBody {
	s.HaVipIds = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetInstanceIds(v *DescribeVSwitchAttributesResponseBodyInstanceIds) *DescribeVSwitchAttributesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetLoadBalancerIds(v *DescribeVSwitchAttributesResponseBodyLoadBalancerIds) *DescribeVSwitchAttributesResponseBody {
	s.LoadBalancerIds = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetNatGatewayIds(v *DescribeVSwitchAttributesResponseBodyNatGatewayIds) *DescribeVSwitchAttributesResponseBody {
	s.NatGatewayIds = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetNetworkId(v string) *DescribeVSwitchAttributesResponseBody {
	s.NetworkId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetNetworkInterfaceIds(v *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds) *DescribeVSwitchAttributesResponseBody {
	s.NetworkInterfaceIds = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetRequestId(v string) *DescribeVSwitchAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetStatus(v string) *DescribeVSwitchAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetVSwitchId(v string) *DescribeVSwitchAttributesResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) SetVSwitchName(v string) *DescribeVSwitchAttributesResponseBody {
	s.VSwitchName = &v
	return s
}

func (s *DescribeVSwitchAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVSwitchAttributesResponseBodyHaVipIds struct {
	HaVipId []*string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchAttributesResponseBodyHaVipIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBodyHaVipIds) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBodyHaVipIds) GetHaVipId() []*string {
	return s.HaVipId
}

func (s *DescribeVSwitchAttributesResponseBodyHaVipIds) SetHaVipId(v []*string) *DescribeVSwitchAttributesResponseBodyHaVipIds {
	s.HaVipId = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyHaVipIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVSwitchAttributesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchAttributesResponseBodyInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBodyInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeVSwitchAttributesResponseBodyInstanceIds) SetInstanceId(v []*string) *DescribeVSwitchAttributesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVSwitchAttributesResponseBodyLoadBalancerIds struct {
	LoadBalancerId []*string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchAttributesResponseBodyLoadBalancerIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBodyLoadBalancerIds) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBodyLoadBalancerIds) GetLoadBalancerId() []*string {
	return s.LoadBalancerId
}

func (s *DescribeVSwitchAttributesResponseBodyLoadBalancerIds) SetLoadBalancerId(v []*string) *DescribeVSwitchAttributesResponseBodyLoadBalancerIds {
	s.LoadBalancerId = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyLoadBalancerIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVSwitchAttributesResponseBodyNatGatewayIds struct {
	NatGatewayId []*string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchAttributesResponseBodyNatGatewayIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBodyNatGatewayIds) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBodyNatGatewayIds) GetNatGatewayId() []*string {
	return s.NatGatewayId
}

func (s *DescribeVSwitchAttributesResponseBodyNatGatewayIds) SetNatGatewayId(v []*string) *DescribeVSwitchAttributesResponseBodyNatGatewayIds {
	s.NatGatewayId = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyNatGatewayIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds struct {
	NetworkInterfaceId []*string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" type:"Repeated"`
}

func (s DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds) GetNetworkInterfaceId() []*string {
	return s.NetworkInterfaceId
}

func (s *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds) SetNetworkInterfaceId(v []*string) *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds {
	s.NetworkInterfaceId = v
	return s
}

func (s *DescribeVSwitchAttributesResponseBodyNetworkInterfaceIds) Validate() error {
	return dara.Validate(s)
}
