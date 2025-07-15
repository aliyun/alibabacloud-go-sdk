// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteEntryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeRouteEntryListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeRouteEntryListResponseBody
	GetRequestId() *string
	SetRouteEntrys(v *DescribeRouteEntryListResponseBodyRouteEntrys) *DescribeRouteEntryListResponseBody
	GetRouteEntrys() *DescribeRouteEntryListResponseBodyRouteEntrys
}

type DescribeRouteEntryListResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If a value is returned for **NextToken**, the value is used to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the routes.
	RouteEntrys *DescribeRouteEntryListResponseBodyRouteEntrys `json:"RouteEntrys,omitempty" xml:"RouteEntrys,omitempty" type:"Struct"`
}

func (s DescribeRouteEntryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRouteEntryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteEntryListResponseBody) GetRouteEntrys() *DescribeRouteEntryListResponseBodyRouteEntrys {
	return s.RouteEntrys
}

func (s *DescribeRouteEntryListResponseBody) SetNextToken(v string) *DescribeRouteEntryListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRouteEntryListResponseBody) SetRequestId(v string) *DescribeRouteEntryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBody) SetRouteEntrys(v *DescribeRouteEntryListResponseBodyRouteEntrys) *DescribeRouteEntryListResponseBody {
	s.RouteEntrys = v
	return s
}

func (s *DescribeRouteEntryListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteEntryListResponseBodyRouteEntrys struct {
	RouteEntry []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry `json:"RouteEntry,omitempty" xml:"RouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrys) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrys) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrys) GetRouteEntry() []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	return s.RouteEntry
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrys) SetRouteEntry(v []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) *DescribeRouteEntryListResponseBodyRouteEntrys {
	s.RouteEntry = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrys) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry struct {
	// The description of the route.
	//
	// example:
	//
	// RouteEntryDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 192.168.2.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The time when the route was modified. The time follows the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-05-09T03:00:07Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The IP version. Valid values: Valid values:
	//
	// 	- **ipv4**
	//
	// 	- **ipv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The information about the next hops.
	NextHops *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops `json:"NextHops,omitempty" xml:"NextHops,omitempty" type:"Struct"`
	// The route origin. Valid values:
	//
	// 	- **RoutePropagation**: The route is created by a dynamic propagation source.
	//
	// 	- **SystemCreate**: The route is created by the system.
	//
	// 	- **CustomCreate**: The route is created by a user.
	//
	// example:
	//
	// RoutePropagation
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The ID of the route.
	//
	// example:
	//
	// rte-bp1mnnr2al0naomnp****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// aaa
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-bp15w5q90d2rk3bww****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// Indicates whether the route is hosted. If the parameter is empty, the route is not hosted.
	//
	// If **TR*	- is returned, the route is hosted by a transit router.
	//
	// example:
	//
	// TR
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The status of the route entry. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The route type. Valid values:
	//
	// 	- **Custom**: custom routes.
	//
	// 	- **System**: system routes.
	//
	// 	- **BGP**: BGP routes.
	//
	// 	- **CEN**: CEN routes.
	//
	// 	- **ECR**: ECR routes.
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetNextHops() *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops {
	return s.NextHops
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetType() *string {
	return s.Type
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetDescription(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.Description = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetDestinationCidrBlock(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetGmtModified(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.GmtModified = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetIpVersion(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.IpVersion = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetNextHops(v *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.NextHops = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetOrigin(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.Origin = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetRouteEntryId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.RouteEntryId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetRouteEntryName(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.RouteEntryName = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetRouteTableId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetServiceType(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.ServiceType = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetStatus(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetType(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.Type = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops struct {
	NextHop []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop `json:"NextHop,omitempty" xml:"NextHop,omitempty" type:"Repeated"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) GetNextHop() []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	return s.NextHop
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) SetNextHop(v []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops {
	s.NextHop = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop struct {
	// Indicates whether the route is available. Valid values:
	//
	// 	- **0**: unavailable
	//
	// 	- **1**: available
	//
	// >  This parameter is returned when the next hop type is set to **RouterInterface**.
	//
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the next hop.
	//
	// example:
	//
	// vpn-bp10zyaph5cc8b7c7****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The ID of the region where the next hop is deployed.
	//
	// >  This parameter is returned when the next hop type is set to **RouterInterface**.
	//
	// example:
	//
	// cn-hangzhou
	NextHopRegionId *string `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	// The information about the next hop.
	NextHopRelatedInfo *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo `json:"NextHopRelatedInfo,omitempty" xml:"NextHopRelatedInfo,omitempty" type:"Struct"`
	// The next hop type. Valid values:
	//
	// 	- **Instance**: an ECS instance.
	//
	// 	- **HaVip**: an HaVip.
	//
	// 	- **VpnGateway**: a VPN gateway.
	//
	// 	- **NatGateway**: a NAT gateway.
	//
	// 	- **NetworkInterface**: a secondary ENI.
	//
	// 	- **RouterInterface**: a router interface.
	//
	// 	- **IPv6Gateway**: an IPv6 gateway.
	//
	// 	- **Attachment**: a transit router.
	//
	// 	- **Ipv4Gateway**: an IPv4 gateway.
	//
	// 	- **GatewayEndpoint**: a gateway endpoint.
	//
	// 	- **CenBasic**: CEN does not support transit routers.
	//
	// 	- **Ecr**: ECR.
	//
	// example:
	//
	// Instance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The weight of the route.
	//
	// >  This parameter is returned when the next hop type is set to **RouterInterface**.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetEnabled() *int32 {
	return s.Enabled
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetNextHopRegionId() *string {
	return s.NextHopRegionId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetNextHopRelatedInfo() *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo {
	return s.NextHopRelatedInfo
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetEnabled(v int32) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.Enabled = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetNextHopId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetNextHopRegionId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopRegionId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetNextHopRelatedInfo(v *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopRelatedInfo = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetNextHopType(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopType = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetWeight(v int32) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.Weight = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo struct {
	// The ID of the instance that is associated with the next hop.
	//
	// example:
	//
	// vpc-bp1t36rn9l53iwbsf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance associated with the next hop. Valid values:
	//
	// 	- **VPC**: a VPC
	//
	// 	- **VBR**: a VBR
	//
	// 	- **PCONN**: an Express Connect circuit
	//
	// example:
	//
	// VPC
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID of the instance associated with the next hop. Valid values:
	//
	// example:
	//
	// ch-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) SetInstanceId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) SetInstanceType(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo {
	s.InstanceType = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) SetRegionId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) Validate() error {
	return dara.Validate(s)
}
