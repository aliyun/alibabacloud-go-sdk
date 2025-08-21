// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *DescribeNetworkAttributeResponseBody
	GetCidrBlock() *string
	SetCloudResources(v *DescribeNetworkAttributeResponseBodyCloudResources) *DescribeNetworkAttributeResponseBody
	GetCloudResources() *DescribeNetworkAttributeResponseBodyCloudResources
	SetCreatedTime(v string) *DescribeNetworkAttributeResponseBody
	GetCreatedTime() *string
	SetDescription(v string) *DescribeNetworkAttributeResponseBody
	GetDescription() *string
	SetEnsRegionId(v string) *DescribeNetworkAttributeResponseBody
	GetEnsRegionId() *string
	SetGatewayRouteTableId(v string) *DescribeNetworkAttributeResponseBody
	GetGatewayRouteTableId() *string
	SetHaVipIds(v *DescribeNetworkAttributeResponseBodyHaVipIds) *DescribeNetworkAttributeResponseBody
	GetHaVipIds() *DescribeNetworkAttributeResponseBodyHaVipIds
	SetInstanceIds(v *DescribeNetworkAttributeResponseBodyInstanceIds) *DescribeNetworkAttributeResponseBody
	GetInstanceIds() *DescribeNetworkAttributeResponseBodyInstanceIds
	SetLoadBalancerIds(v *DescribeNetworkAttributeResponseBodyLoadBalancerIds) *DescribeNetworkAttributeResponseBody
	GetLoadBalancerIds() *DescribeNetworkAttributeResponseBodyLoadBalancerIds
	SetNatGatewayIds(v *DescribeNetworkAttributeResponseBodyNatGatewayIds) *DescribeNetworkAttributeResponseBody
	GetNatGatewayIds() *DescribeNetworkAttributeResponseBodyNatGatewayIds
	SetNetworkAclId(v string) *DescribeNetworkAttributeResponseBody
	GetNetworkAclId() *string
	SetNetworkId(v string) *DescribeNetworkAttributeResponseBody
	GetNetworkId() *string
	SetNetworkInterfaceIds(v *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds) *DescribeNetworkAttributeResponseBody
	GetNetworkInterfaceIds() *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds
	SetNetworkName(v string) *DescribeNetworkAttributeResponseBody
	GetNetworkName() *string
	SetRequestId(v string) *DescribeNetworkAttributeResponseBody
	GetRequestId() *string
	SetRouteTableId(v string) *DescribeNetworkAttributeResponseBody
	GetRouteTableId() *string
	SetRouteTableIds(v *DescribeNetworkAttributeResponseBodyRouteTableIds) *DescribeNetworkAttributeResponseBody
	GetRouteTableIds() *DescribeNetworkAttributeResponseBodyRouteTableIds
	SetRouterTableId(v string) *DescribeNetworkAttributeResponseBody
	GetRouterTableId() *string
	SetStatus(v string) *DescribeNetworkAttributeResponseBody
	GetStatus() *string
	SetVSwitchIds(v *DescribeNetworkAttributeResponseBodyVSwitchIds) *DescribeNetworkAttributeResponseBody
	GetVSwitchIds() *DescribeNetworkAttributeResponseBodyVSwitchIds
}

type DescribeNetworkAttributeResponseBody struct {
	// The IPv4 CIDR block of the network.
	//
	// example:
	//
	// 10.0.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The list of resources in the network.
	CloudResources *DescribeNetworkAttributeResponseBodyCloudResources `json:"CloudResources,omitempty" xml:"CloudResources,omitempty" type:"Struct"`
	// The time when the network was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-01T00:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the network.
	//
	// example:
	//
	// abc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-beijing
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the gateway route table.
	//
	// example:
	//
	// rt-539***tbs
	GatewayRouteTableId *string `json:"GatewayRouteTableId,omitempty" xml:"GatewayRouteTableId,omitempty"`
	// List of HaVipIds.
	HaVipIds *DescribeNetworkAttributeResponseBodyHaVipIds `json:"HaVipIds,omitempty" xml:"HaVipIds,omitempty" type:"Struct"`
	// The instance IDs.
	InstanceIds *DescribeNetworkAttributeResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// List of ELB instances.
	LoadBalancerIds *DescribeNetworkAttributeResponseBodyLoadBalancerIds `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Struct"`
	// List of NAT Gateways.
	NatGatewayIds *DescribeNetworkAttributeResponseBodyNatGatewayIds `json:"NatGatewayIds,omitempty" xml:"NatGatewayIds,omitempty" type:"Struct"`
	// The ID of the network access control list (ACL).
	//
	// example:
	//
	// nacl-a2do9e413e0sp****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5***
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// A list of multicast source IDs.
	NetworkInterfaceIds *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Struct"`
	// The name of the network.
	//
	// example:
	//
	// abc
	NetworkName *string `json:"NetworkName,omitempty" xml:"NetworkName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// rt-539***fpu
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// List of routing table IDs.
	RouteTableIds *DescribeNetworkAttributeResponseBodyRouteTableIds `json:"RouteTableIds,omitempty" xml:"RouteTableIds,omitempty" type:"Struct"`
	// The ID of the route table.
	//
	// example:
	//
	// rtb-5***
	RouterTableId *string `json:"RouterTableId,omitempty" xml:"RouterTableId,omitempty"`
	// The status of the network. Valid values:
	//
	// 	- Pending
	//
	// 	- Available
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of vSwitches in the network.
	VSwitchIds *DescribeNetworkAttributeResponseBodyVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
}

func (s DescribeNetworkAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeNetworkAttributeResponseBody) GetCloudResources() *DescribeNetworkAttributeResponseBodyCloudResources {
	return s.CloudResources
}

func (s *DescribeNetworkAttributeResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeNetworkAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAttributeResponseBody) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNetworkAttributeResponseBody) GetGatewayRouteTableId() *string {
	return s.GatewayRouteTableId
}

func (s *DescribeNetworkAttributeResponseBody) GetHaVipIds() *DescribeNetworkAttributeResponseBodyHaVipIds {
	return s.HaVipIds
}

func (s *DescribeNetworkAttributeResponseBody) GetInstanceIds() *DescribeNetworkAttributeResponseBodyInstanceIds {
	return s.InstanceIds
}

func (s *DescribeNetworkAttributeResponseBody) GetLoadBalancerIds() *DescribeNetworkAttributeResponseBodyLoadBalancerIds {
	return s.LoadBalancerIds
}

func (s *DescribeNetworkAttributeResponseBody) GetNatGatewayIds() *DescribeNetworkAttributeResponseBodyNatGatewayIds {
	return s.NatGatewayIds
}

func (s *DescribeNetworkAttributeResponseBody) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeNetworkAttributeResponseBody) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworkAttributeResponseBody) GetNetworkInterfaceIds() *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds {
	return s.NetworkInterfaceIds
}

func (s *DescribeNetworkAttributeResponseBody) GetNetworkName() *string {
	return s.NetworkName
}

func (s *DescribeNetworkAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkAttributeResponseBody) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeNetworkAttributeResponseBody) GetRouteTableIds() *DescribeNetworkAttributeResponseBodyRouteTableIds {
	return s.RouteTableIds
}

func (s *DescribeNetworkAttributeResponseBody) GetRouterTableId() *string {
	return s.RouterTableId
}

func (s *DescribeNetworkAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkAttributeResponseBody) GetVSwitchIds() *DescribeNetworkAttributeResponseBodyVSwitchIds {
	return s.VSwitchIds
}

func (s *DescribeNetworkAttributeResponseBody) SetCidrBlock(v string) *DescribeNetworkAttributeResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetCloudResources(v *DescribeNetworkAttributeResponseBodyCloudResources) *DescribeNetworkAttributeResponseBody {
	s.CloudResources = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetCreatedTime(v string) *DescribeNetworkAttributeResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetDescription(v string) *DescribeNetworkAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetEnsRegionId(v string) *DescribeNetworkAttributeResponseBody {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetGatewayRouteTableId(v string) *DescribeNetworkAttributeResponseBody {
	s.GatewayRouteTableId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetHaVipIds(v *DescribeNetworkAttributeResponseBodyHaVipIds) *DescribeNetworkAttributeResponseBody {
	s.HaVipIds = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetInstanceIds(v *DescribeNetworkAttributeResponseBodyInstanceIds) *DescribeNetworkAttributeResponseBody {
	s.InstanceIds = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetLoadBalancerIds(v *DescribeNetworkAttributeResponseBodyLoadBalancerIds) *DescribeNetworkAttributeResponseBody {
	s.LoadBalancerIds = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetNatGatewayIds(v *DescribeNetworkAttributeResponseBodyNatGatewayIds) *DescribeNetworkAttributeResponseBody {
	s.NatGatewayIds = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetNetworkAclId(v string) *DescribeNetworkAttributeResponseBody {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetNetworkId(v string) *DescribeNetworkAttributeResponseBody {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetNetworkInterfaceIds(v *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds) *DescribeNetworkAttributeResponseBody {
	s.NetworkInterfaceIds = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetNetworkName(v string) *DescribeNetworkAttributeResponseBody {
	s.NetworkName = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetRequestId(v string) *DescribeNetworkAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetRouteTableId(v string) *DescribeNetworkAttributeResponseBody {
	s.RouteTableId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetRouteTableIds(v *DescribeNetworkAttributeResponseBodyRouteTableIds) *DescribeNetworkAttributeResponseBody {
	s.RouteTableIds = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetRouterTableId(v string) *DescribeNetworkAttributeResponseBody {
	s.RouterTableId = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetStatus(v string) *DescribeNetworkAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) SetVSwitchIds(v *DescribeNetworkAttributeResponseBodyVSwitchIds) *DescribeNetworkAttributeResponseBody {
	s.VSwitchIds = v
	return s
}

func (s *DescribeNetworkAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyCloudResources struct {
	CloudResourceSetType []*DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType `json:"CloudResourceSetType,omitempty" xml:"CloudResourceSetType,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyCloudResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyCloudResources) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyCloudResources) GetCloudResourceSetType() []*DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType {
	return s.CloudResourceSetType
}

func (s *DescribeNetworkAttributeResponseBodyCloudResources) SetCloudResourceSetType(v []*DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) *DescribeNetworkAttributeResponseBodyCloudResources {
	s.CloudResourceSetType = v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyCloudResources) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType struct {
	// The number of resources in the network.
	//
	// example:
	//
	// 3
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The resource type. VSwitch.
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) GetResourceCount() *int32 {
	return s.ResourceCount
}

func (s *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) SetResourceCount(v int32) *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType {
	s.ResourceCount = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) SetResourceType(v string) *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType {
	s.ResourceType = &v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyCloudResourcesCloudResourceSetType) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyHaVipIds struct {
	HaVipId []*string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyHaVipIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyHaVipIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyHaVipIds) GetHaVipId() []*string {
	return s.HaVipId
}

func (s *DescribeNetworkAttributeResponseBodyHaVipIds) SetHaVipId(v []*string) *DescribeNetworkAttributeResponseBodyHaVipIds {
	s.HaVipId = v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyHaVipIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeNetworkAttributeResponseBodyInstanceIds) SetInstanceId(v []*string) *DescribeNetworkAttributeResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyLoadBalancerIds struct {
	LoadBalancerId []*string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyLoadBalancerIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyLoadBalancerIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyLoadBalancerIds) GetLoadBalancerId() []*string {
	return s.LoadBalancerId
}

func (s *DescribeNetworkAttributeResponseBodyLoadBalancerIds) SetLoadBalancerId(v []*string) *DescribeNetworkAttributeResponseBodyLoadBalancerIds {
	s.LoadBalancerId = v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyLoadBalancerIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyNatGatewayIds struct {
	NatGatewayId []*string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyNatGatewayIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyNatGatewayIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyNatGatewayIds) GetNatGatewayId() []*string {
	return s.NatGatewayId
}

func (s *DescribeNetworkAttributeResponseBodyNatGatewayIds) SetNatGatewayId(v []*string) *DescribeNetworkAttributeResponseBodyNatGatewayIds {
	s.NatGatewayId = v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyNatGatewayIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyNetworkInterfaceIds struct {
	NetworkInterfaceId []*string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyNetworkInterfaceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyNetworkInterfaceIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds) GetNetworkInterfaceId() []*string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds) SetNetworkInterfaceId(v []*string) *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds {
	s.NetworkInterfaceId = v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyNetworkInterfaceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyRouteTableIds struct {
	RouteTableId []*string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyRouteTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyRouteTableIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyRouteTableIds) GetRouteTableId() []*string {
	return s.RouteTableId
}

func (s *DescribeNetworkAttributeResponseBodyRouteTableIds) SetRouteTableId(v []*string) *DescribeNetworkAttributeResponseBodyRouteTableIds {
	s.RouteTableId = v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyRouteTableIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAttributeResponseBodyVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAttributeResponseBodyVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAttributeResponseBodyVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAttributeResponseBodyVSwitchIds) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *DescribeNetworkAttributeResponseBodyVSwitchIds) SetVSwitchId(v []*string) *DescribeNetworkAttributeResponseBodyVSwitchIds {
	s.VSwitchId = v
	return s
}

func (s *DescribeNetworkAttributeResponseBodyVSwitchIds) Validate() error {
	return dara.Validate(s)
}
