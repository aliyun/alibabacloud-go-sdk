// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteEntryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestCidrBlockList(v []*string) *DescribeRouteEntryListRequest
	GetDestCidrBlockList() []*string
	SetDestinationCidrBlock(v string) *DescribeRouteEntryListRequest
	GetDestinationCidrBlock() *string
	SetIpVersion(v string) *DescribeRouteEntryListRequest
	GetIpVersion() *string
	SetMaxResult(v int32) *DescribeRouteEntryListRequest
	GetMaxResult() *int32
	SetNextHopId(v string) *DescribeRouteEntryListRequest
	GetNextHopId() *string
	SetNextHopType(v string) *DescribeRouteEntryListRequest
	GetNextHopType() *string
	SetNextToken(v string) *DescribeRouteEntryListRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeRouteEntryListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRouteEntryListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeRouteEntryListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRouteEntryListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouteEntryListRequest
	GetResourceOwnerId() *int64
	SetRouteEntryId(v string) *DescribeRouteEntryListRequest
	GetRouteEntryId() *string
	SetRouteEntryName(v string) *DescribeRouteEntryListRequest
	GetRouteEntryName() *string
	SetRouteEntryType(v string) *DescribeRouteEntryListRequest
	GetRouteEntryType() *string
	SetRouteTableId(v string) *DescribeRouteEntryListRequest
	GetRouteTableId() *string
	SetServiceType(v string) *DescribeRouteEntryListRequest
	GetServiceType() *string
}

type DescribeRouteEntryListRequest struct {
	// The list of destination CIDR blocks of routes.
	DestCidrBlockList []*string `json:"DestCidrBlockList,omitempty" xml:"DestCidrBlockList,omitempty" type:"Repeated"`
	// The destination CIDR block of the route. Both IPv4 and IPv6 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.2.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The version of the IP protocol. Valid values:
	//
	// - **ipv4**: IPv4.
	//
	// - **ipv6**: IPv6.
	//
	// example:
	//
	// ipv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The instance ID of the next hop.
	//
	// example:
	//
	// vpn-bp10zyaph5cc8b7c7****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The next hop type. Valid values:
	//
	// - **Instance*	- (default): ECS instance.
	//
	// - **HaVip**: high-availability virtual IP address.
	//
	// - **VpnGateway**: VPN gateway.
	//
	// - **NatGateway**: NAT gateway.
	//
	// - **NetworkInterface**: secondary elastic network interface controller (NIC).
	//
	// - **RouterInterface**: routing interface.
	//
	// - **IPv6Gateway**: IPv6 gateway.
	//
	// - **Attachment**: transit router.
	//
	// - **Ipv4Gateway**: IPv4 gateway.
	//
	// - **GatewayEndpoint**: gateway endpoint.
	//
	// - **Ecr**: Express Connect Router (ECR).
	//
	// example:
	//
	// Instance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The pagination token. Valid values:
	//
	// - If this is the first request or no subsequent query exists, leave this parameter empty.
	//
	// - If a subsequent query exists, set the value to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the route table to which the route belongs.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the route that you want to query.
	//
	// example:
	//
	// rte-bp1mnnr2al0naomnp****
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The name of the route.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The type of the route. Valid values:
	//
	// - **Custom**: custom route.
	//
	// - **System**: system route.
	//
	// - **BGP**: BGP route.
	//
	// - **CEN**: Cloud Enterprise Network (CEN) route.
	//
	// - **ECR**: Express Connect Router (ECR) route.
	//
	// example:
	//
	// System
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	// The ID of the route table that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1r9pvl4xen8s9ju****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The type of route hosting. If this field is empty, the route is not hosted.
	//
	// Valid values: **TR**, which indicates that the hosting type is transit router.
	//
	// example:
	//
	// TR
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s DescribeRouteEntryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListRequest) GetDestCidrBlockList() []*string {
	return s.DestCidrBlockList
}

func (s *DescribeRouteEntryListRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeRouteEntryListRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeRouteEntryListRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *DescribeRouteEntryListRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeRouteEntryListRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeRouteEntryListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRouteEntryListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRouteEntryListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouteEntryListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteEntryListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouteEntryListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouteEntryListRequest) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DescribeRouteEntryListRequest) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *DescribeRouteEntryListRequest) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *DescribeRouteEntryListRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteEntryListRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeRouteEntryListRequest) SetDestCidrBlockList(v []*string) *DescribeRouteEntryListRequest {
	s.DestCidrBlockList = v
	return s
}

func (s *DescribeRouteEntryListRequest) SetDestinationCidrBlock(v string) *DescribeRouteEntryListRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetIpVersion(v string) *DescribeRouteEntryListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetMaxResult(v int32) *DescribeRouteEntryListRequest {
	s.MaxResult = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetNextHopId(v string) *DescribeRouteEntryListRequest {
	s.NextHopId = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetNextHopType(v string) *DescribeRouteEntryListRequest {
	s.NextHopType = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetNextToken(v string) *DescribeRouteEntryListRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetOwnerAccount(v string) *DescribeRouteEntryListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetOwnerId(v int64) *DescribeRouteEntryListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetRegionId(v string) *DescribeRouteEntryListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetResourceOwnerAccount(v string) *DescribeRouteEntryListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetResourceOwnerId(v int64) *DescribeRouteEntryListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetRouteEntryId(v string) *DescribeRouteEntryListRequest {
	s.RouteEntryId = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetRouteEntryName(v string) *DescribeRouteEntryListRequest {
	s.RouteEntryName = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetRouteEntryType(v string) *DescribeRouteEntryListRequest {
	s.RouteEntryType = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetRouteTableId(v string) *DescribeRouteEntryListRequest {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteEntryListRequest) SetServiceType(v string) *DescribeRouteEntryListRequest {
	s.ServiceType = &v
	return s
}

func (s *DescribeRouteEntryListRequest) Validate() error {
	return dara.Validate(s)
}
