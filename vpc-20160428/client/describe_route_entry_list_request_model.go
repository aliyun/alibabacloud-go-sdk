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
	// The destination CIDR blocks of the routes.
	DestCidrBlockList []*string `json:"DestCidrBlockList,omitempty" xml:"DestCidrBlockList,omitempty" type:"Repeated"`
	// The destination CIDR block of the route. IPv4 and IPv6 CIDR blocks are supported.
	//
	// example:
	//
	// 192.168.2.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The IP version. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **IPv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The ID of the next hop.
	//
	// example:
	//
	// vpn-bp10zyaph5cc8b7c7****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The next hop type. Valid values:
	//
	// 	- **Instance**: an Elastic Compute Service (ECS) instance. This is the default value.
	//
	// 	- **HaVip**: a high-availability virtual IP address (HaVip).
	//
	// 	- **VpnGateway**: a VPN gateway.
	//
	// 	- **NatGateway**: a NAT gateway.
	//
	// 	- **NetworkInterface**: a secondary elastic network interface (ENI).
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
	// 	- **Ecr**: Express Connect Router (ECR).
	//
	// example:
	//
	// Instance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the route table.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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
	// The name of the route entry.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The route type. Valid values:
	//
	// 	- **Custom**: custom routes.
	//
	// 	- **System**: system routes.
	//
	// 	- **BGP**: BGP routes.
	//
	// 	- **CEN**: Cloud Enterprise Network (CEN) routes.
	//
	// 	- **ECR**: Express Connect Router (ECR) routes.
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
	// Specifies whether to host the route. If the parameter is empty, the route is not hosted.
	//
	// Set the value to **TR**, which specifies that the route is hosted by a transit router.
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
