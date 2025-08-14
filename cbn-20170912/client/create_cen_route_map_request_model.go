// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsPathMatchMode(v string) *CreateCenRouteMapRequest
	GetAsPathMatchMode() *string
	SetCenId(v string) *CreateCenRouteMapRequest
	GetCenId() *string
	SetCenRegionId(v string) *CreateCenRouteMapRequest
	GetCenRegionId() *string
	SetCidrMatchMode(v string) *CreateCenRouteMapRequest
	GetCidrMatchMode() *string
	SetCommunityMatchMode(v string) *CreateCenRouteMapRequest
	GetCommunityMatchMode() *string
	SetCommunityOperateMode(v string) *CreateCenRouteMapRequest
	GetCommunityOperateMode() *string
	SetDescription(v string) *CreateCenRouteMapRequest
	GetDescription() *string
	SetDestinationChildInstanceTypes(v []*string) *CreateCenRouteMapRequest
	GetDestinationChildInstanceTypes() []*string
	SetDestinationCidrBlocks(v []*string) *CreateCenRouteMapRequest
	GetDestinationCidrBlocks() []*string
	SetDestinationInstanceIds(v []*string) *CreateCenRouteMapRequest
	GetDestinationInstanceIds() []*string
	SetDestinationInstanceIdsReverseMatch(v bool) *CreateCenRouteMapRequest
	GetDestinationInstanceIdsReverseMatch() *bool
	SetDestinationRegionIds(v []*string) *CreateCenRouteMapRequest
	GetDestinationRegionIds() []*string
	SetDestinationRouteTableIds(v []*string) *CreateCenRouteMapRequest
	GetDestinationRouteTableIds() []*string
	SetMapResult(v string) *CreateCenRouteMapRequest
	GetMapResult() *string
	SetMatchAddressType(v string) *CreateCenRouteMapRequest
	GetMatchAddressType() *string
	SetMatchAsns(v []*int64) *CreateCenRouteMapRequest
	GetMatchAsns() []*int64
	SetMatchCommunitySet(v []*string) *CreateCenRouteMapRequest
	GetMatchCommunitySet() []*string
	SetNextPriority(v int32) *CreateCenRouteMapRequest
	GetNextPriority() *int32
	SetOperateCommunitySet(v []*string) *CreateCenRouteMapRequest
	GetOperateCommunitySet() []*string
	SetOwnerAccount(v string) *CreateCenRouteMapRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCenRouteMapRequest
	GetOwnerId() *int64
	SetPreference(v int32) *CreateCenRouteMapRequest
	GetPreference() *int32
	SetPrependAsPath(v []*int64) *CreateCenRouteMapRequest
	GetPrependAsPath() []*int64
	SetPriority(v int32) *CreateCenRouteMapRequest
	GetPriority() *int32
	SetResourceOwnerAccount(v string) *CreateCenRouteMapRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCenRouteMapRequest
	GetResourceOwnerId() *int64
	SetRouteTypes(v []*string) *CreateCenRouteMapRequest
	GetRouteTypes() []*string
	SetSourceChildInstanceTypes(v []*string) *CreateCenRouteMapRequest
	GetSourceChildInstanceTypes() []*string
	SetSourceInstanceIds(v []*string) *CreateCenRouteMapRequest
	GetSourceInstanceIds() []*string
	SetSourceInstanceIdsReverseMatch(v bool) *CreateCenRouteMapRequest
	GetSourceInstanceIdsReverseMatch() *bool
	SetSourceRegionIds(v []*string) *CreateCenRouteMapRequest
	GetSourceRegionIds() []*string
	SetSourceRouteTableIds(v []*string) *CreateCenRouteMapRequest
	GetSourceRouteTableIds() []*string
	SetTransitRouterRouteTableId(v string) *CreateCenRouteMapRequest
	GetTransitRouterRouteTableId() *string
	SetTransmitDirection(v string) *CreateCenRouteMapRequest
	GetTransmitDirection() *string
}

type CreateCenRouteMapRequest struct {
	// The match method that is used to match routes based on the AS path. Valid values:
	//
	// 	- **Include**: fuzzy match. A route is a match if the AS path of the route overlaps with the AS path in the match conditions.
	//
	// 	- **Complete**: exact match. A route is a match only if the AS path of the route matches the AS path in the match conditions.
	//
	// example:
	//
	// Include
	AsPathMatchMode *string `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the region in which the routing policy is applied.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	CenRegionId *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	// The match method that is used to match routes against the prefix list. Valid values:
	//
	// 	- **Include**: fuzzy match. A route is a match if the route prefix is included in the match conditions.
	//
	// For example, if you set the match condition to 1.1.0.0/16 and fuzzy match is applied, the route whose prefix is 1.1.1.0/24 meets the match condition.
	//
	// 	- **Complete**: exact match. A route is a match only if the route prefix is the same as the prefix specified in the match condition.
	//
	// For example, if you set the match condition to 1.1.0.0/16 and exact match is applied, only the route whose prefix is 1.1.0.0/16 meets the match condition.
	//
	// example:
	//
	// Include
	CidrMatchMode *string `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	// The match method that is used to match routes based on the community. Valid values:
	//
	// 	- **Include**: fuzzy match. A route is a match if the community of the route overlaps with the community in the match conditions.
	//
	// 	- **Complete**: exact match. A route is a match only if the community of the route matches the community in the match conditions.
	//
	// example:
	//
	// Include
	CommunityMatchMode *string `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	// The action to be performed on the community. Valid values:
	//
	// 	- **Additive**: adds the community to the route.
	//
	// 	- **Replace**: replaces the original community of the route.
	//
	// This parameter specifies the action to be performed when a route meets the match condition.
	//
	// example:
	//
	// Additive
	CommunityOperateMode *string `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	// The description of the routing policy.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The types of destination network instance to which the routes belong. The following types of network instances are supported:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// 	- **VPN**: IPsec connection
	//
	//     >This parameter does not take effect if the IPsec-VPN connection or SSL client is associated with a transit router through a VPN gateway and a VPC. This parameter takes effect only if the IPsec connection is directly connected to the transit router.
	//
	// You can specify one or more network instance types.
	//
	// > The destination network instance types are valid only if the routing policy is applied to scenarios where routes are advertised from the gateway in the current region to network instances in the current region.
	//
	// example:
	//
	// VPC
	DestinationChildInstanceTypes []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	// The prefix list against which routes are matched.
	//
	// Specify IP addresses in CIDR notations. You can specify at most 32 CIDR blocks.
	//
	// IPv4 and IPv4 addresses are supported.
	//
	// example:
	//
	// 10.10.10.0/24
	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Repeated"`
	// The IDs of the destination network instances to which the routes belong. The following network instance types are supported:
	//
	// 	- VPC
	//
	// 	- VBR
	//
	// 	- CCN instance
	//
	// 	- SAG instance
	//
	// 	- The ID of the IPsec-VPN connection.
	//
	// You can enter at most 32 IDs.
	//
	// > The destination instance IDs take effect only when Direction is set to Export from Regional Gateway and the destination instances are deployed in the current region.
	//
	// example:
	//
	// vpc-afrfs434465fdf****
	DestinationInstanceIds []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to exclude destination instance IDs. Valid values:
	//
	// 	- **false*	- (default): A route is a match if the destination instance ID is included in the list specified by **SourceInstanceIds.N**.
	//
	// 	- **true**: A route is a match if the destination network instance ID is not in the list specified by **SourceInstanceIds.N**.
	//
	// example:
	//
	// false
	DestinationInstanceIdsReverseMatch *bool `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	// The destination region IDs of the route. You can specify at most 32 region IDs.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// The IDs of the destination route tables to which routes are evaluated. You can enter at most 32 route table IDs.
	//
	// > The destination route table IDs take effect only when Direction is set to Export from Regional Gateway and the destination route tables belong to network instances deployed in the current region.
	//
	// example:
	//
	// vtb-adefrgtr144vf****
	DestinationRouteTableIds []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	// The action to be performed on a route that meets all the match conditions. Valid values:
	//
	// 	- **Permit**: the route is permitted.
	//
	// 	- **Deny**: the route is denied.
	//
	// This parameter is required.
	//
	// example:
	//
	// Permit
	MapResult *string `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	// The type of IP address in the match condition. Valid values:
	//
	// 	- **IPv4**: IPv4 address
	//
	// 	- **IPv6**: IPv6 address
	//
	// This parameter can be empty. If no value is specified, all types of IP address are a match.
	//
	// example:
	//
	// IPv4
	MatchAddressType *string `json:"MatchAddressType,omitempty" xml:"MatchAddressType,omitempty"`
	// The AS paths based on which routes are compared.
	//
	// You can specify at most 32 AS numbers.
	//
	// > Only the AS-SEQUENCE parameter is supported. The AS-SET, AS-CONFED-SEQUENCE, and AS-CONFED-SET parameters are not supported. In other words, only the AS number list is supported. Sets and sub-lists are not supported.
	//
	// example:
	//
	// 65501
	MatchAsns []*int64 `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	// The community set based on which routes are compared.
	//
	// Specify the community in the format of n:m. Valid values of n and m: **1*	- to **65535**. Each community must comply with the RFC 1997 standard. The RFC 8092 standard that defines Border Gateway Protocol (BGP) large communities is not supported.
	//
	// You can specify at most 32 communities.
	//
	// > If the configurations of the communities are incorrect, routes may fail to be advertised to your data center.
	//
	// example:
	//
	// 65501:1
	MatchCommunitySet []*string `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Repeated"`
	// The priority of the routing policy that you want to associate with the current one.
	//
	// 	- This parameter takes effect only when the **MapResult*	- parameter is set to **Permit**. This way, the permitted route is matched against the next routing policy.
	//
	// 	- The region and direction of the routing policy to be associated must be the same as those of the current routing policy.
	//
	// 	- The priority of the next routing policy must be lower than the priority of the current routing policy.
	//
	// example:
	//
	// 20
	NextPriority *int32 `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	// The community set on which actions are performed.
	//
	// Specify the community in the format of n:m. Valid values of n and m: **1*	- to **65535**. Each community must comply with RFC 1997. The RFC 8092 standard that defines BGP large communities is not supported.
	//
	// You can specify at most 32 communities.
	//
	// > If the configurations of the communities are incorrect, routes may fail to be advertised to your data center.
	//
	// example:
	//
	// 65501:1
	OperateCommunitySet []*string `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new priority of the route.
	//
	// Valid values: **1*	- to **100**. The default priority is **50**. A smaller value indicates a higher priority.
	//
	// This parameter specifies the action to be performed when a route meets the match condition.
	//
	// example:
	//
	// 50
	Preference *int32 `json:"Preference,omitempty" xml:"Preference,omitempty"`
	// The AS paths that are prepended by using an action statement when regional gateways receive or advertise routes.
	//
	// The AS paths vary based on the direction in which the routing policy is applied:
	//
	// 	- If AS paths are prepended to a routing policy that is applied in the inbound direction, you must specify source network instance IDs and the source region in the match condition. In addition, the source region must be the same as the region where the routing policy is applied.
	//
	// 	- If AS paths are prepended to a routing policy that is applied in the outbound direction, you must specify destination network instance IDs in the match condition.
	//
	// This parameter specifies the action to be performed when a route meets the match condition. You can specify at most 32 AS numbers.
	//
	// example:
	//
	// 65501
	PrependAsPath []*int64 `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Repeated"`
	// The priority of the routing policy. Valid values: **1*	- to **100**. A smaller value indicates a higher priority.
	//
	// > You cannot specify the same priority for routing policies that apply in the same region and direction. The system matches routes against the match conditions of routing policies in descending order of priority. A smaller value indicates a higher priority. You must set the priorities to proper values.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Priority             *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of route to be compared. Valid values: The following route types are supported:
	//
	// 	- **System**: system routes that are automatically generated by the system.
	//
	// 	- **Custom**: custom routes that are manually added.
	//
	// 	- **BGP**: routes that are advertised over BGP.
	//
	// You can specify multiple route types.
	//
	// example:
	//
	// System
	RouteTypes []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	// The types of source network instance to which the routes belong. The following types of network instances are supported:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// 	- **VPN**: VPN gateway or IPsec connection
	//
	//     	- If the IPsec-VPN connection or SSL client is associated with a VPN gateway, the VPC associated with the VPN gateway must be connected to a transit router, and the VPN gateway must use BGP dynamic routing. Otherwise, this parameter cannot take effect.
	//
	//     	- This parameter takes effect if the IPsec connection is directly connected to a transit router.
	//
	// You can specify one or more network instance types.
	//
	// example:
	//
	// VPC
	SourceChildInstanceTypes []*string `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Repeated"`
	// The IDs of the source network instances to which the routes belong. The following network instance types are supported:
	//
	// 	- Virtual private cloud (VPC)
	//
	// 	- Virtual border router (VBR)
	//
	// 	- Cloud Connect Network (CCN) instance
	//
	// 	- Smart Access Gateway (SAG) instance
	//
	// 	- The ID of the IPsec-VPN connection.
	//
	// You can enter at most 32 IDs.
	//
	// example:
	//
	// vpc-adeg3544fdf34vf****
	SourceInstanceIds []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to exclude source instance IDs. Valid values:
	//
	// 	- **false*	- (default): A route is a match if the source instance ID is included in the list specified by **SourceInstanceIds.N**.
	//
	// 	- **true**: A route is a match if the source network instance ID is not in the list specified by **SourceInstanceIds.N**.
	//
	// example:
	//
	// false
	SourceInstanceIdsReverseMatch *bool `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	// The IDs of the source regions from which routes are evaluated. You can enter at most 32 region IDs.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	SourceRegionIds []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	// The IDs of the source route tables from which routes are evaluated. You can enter at most 32 route table IDs.
	//
	// example:
	//
	// vtb-adfr233vf34rvd4****
	SourceRouteTableIds []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
	// The ID of the route table of the transit router.
	//
	// If you do not specify a route table ID, the routing policy is automatically associated with the default route table of the transit router.
	//
	// example:
	//
	// vtb-gw8nx3515m1mbd1z1****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	// The direction in which the routing policy is applied. Valid values:
	//
	// 	- **RegionIn**: Routes are advertised to the gateways in the regions that are connected by the CEN instance.
	//
	// For example, routes are advertised from network instances deployed in the current region or other regions to the gateway deployed in the current region.
	//
	// 	- **RegionOut**: Routes are advertised from the gateways in the regions that are connected by the CEN instance.
	//
	// For example, routes are advertised from the gateway deployed in the current region to network instances deployed in the same region, or to gateways deployed in other regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// RegionIn
	TransmitDirection *string `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
}

func (s CreateCenRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *CreateCenRouteMapRequest) GetAsPathMatchMode() *string {
	return s.AsPathMatchMode
}

func (s *CreateCenRouteMapRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateCenRouteMapRequest) GetCenRegionId() *string {
	return s.CenRegionId
}

func (s *CreateCenRouteMapRequest) GetCidrMatchMode() *string {
	return s.CidrMatchMode
}

func (s *CreateCenRouteMapRequest) GetCommunityMatchMode() *string {
	return s.CommunityMatchMode
}

func (s *CreateCenRouteMapRequest) GetCommunityOperateMode() *string {
	return s.CommunityOperateMode
}

func (s *CreateCenRouteMapRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCenRouteMapRequest) GetDestinationChildInstanceTypes() []*string {
	return s.DestinationChildInstanceTypes
}

func (s *CreateCenRouteMapRequest) GetDestinationCidrBlocks() []*string {
	return s.DestinationCidrBlocks
}

func (s *CreateCenRouteMapRequest) GetDestinationInstanceIds() []*string {
	return s.DestinationInstanceIds
}

func (s *CreateCenRouteMapRequest) GetDestinationInstanceIdsReverseMatch() *bool {
	return s.DestinationInstanceIdsReverseMatch
}

func (s *CreateCenRouteMapRequest) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *CreateCenRouteMapRequest) GetDestinationRouteTableIds() []*string {
	return s.DestinationRouteTableIds
}

func (s *CreateCenRouteMapRequest) GetMapResult() *string {
	return s.MapResult
}

func (s *CreateCenRouteMapRequest) GetMatchAddressType() *string {
	return s.MatchAddressType
}

func (s *CreateCenRouteMapRequest) GetMatchAsns() []*int64 {
	return s.MatchAsns
}

func (s *CreateCenRouteMapRequest) GetMatchCommunitySet() []*string {
	return s.MatchCommunitySet
}

func (s *CreateCenRouteMapRequest) GetNextPriority() *int32 {
	return s.NextPriority
}

func (s *CreateCenRouteMapRequest) GetOperateCommunitySet() []*string {
	return s.OperateCommunitySet
}

func (s *CreateCenRouteMapRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCenRouteMapRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCenRouteMapRequest) GetPreference() *int32 {
	return s.Preference
}

func (s *CreateCenRouteMapRequest) GetPrependAsPath() []*int64 {
	return s.PrependAsPath
}

func (s *CreateCenRouteMapRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateCenRouteMapRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCenRouteMapRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCenRouteMapRequest) GetRouteTypes() []*string {
	return s.RouteTypes
}

func (s *CreateCenRouteMapRequest) GetSourceChildInstanceTypes() []*string {
	return s.SourceChildInstanceTypes
}

func (s *CreateCenRouteMapRequest) GetSourceInstanceIds() []*string {
	return s.SourceInstanceIds
}

func (s *CreateCenRouteMapRequest) GetSourceInstanceIdsReverseMatch() *bool {
	return s.SourceInstanceIdsReverseMatch
}

func (s *CreateCenRouteMapRequest) GetSourceRegionIds() []*string {
	return s.SourceRegionIds
}

func (s *CreateCenRouteMapRequest) GetSourceRouteTableIds() []*string {
	return s.SourceRouteTableIds
}

func (s *CreateCenRouteMapRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *CreateCenRouteMapRequest) GetTransmitDirection() *string {
	return s.TransmitDirection
}

func (s *CreateCenRouteMapRequest) SetAsPathMatchMode(v string) *CreateCenRouteMapRequest {
	s.AsPathMatchMode = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCenId(v string) *CreateCenRouteMapRequest {
	s.CenId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCenRegionId(v string) *CreateCenRouteMapRequest {
	s.CenRegionId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCidrMatchMode(v string) *CreateCenRouteMapRequest {
	s.CidrMatchMode = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCommunityMatchMode(v string) *CreateCenRouteMapRequest {
	s.CommunityMatchMode = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetCommunityOperateMode(v string) *CreateCenRouteMapRequest {
	s.CommunityOperateMode = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetDescription(v string) *CreateCenRouteMapRequest {
	s.Description = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationChildInstanceTypes(v []*string) *CreateCenRouteMapRequest {
	s.DestinationChildInstanceTypes = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationCidrBlocks(v []*string) *CreateCenRouteMapRequest {
	s.DestinationCidrBlocks = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationInstanceIds(v []*string) *CreateCenRouteMapRequest {
	s.DestinationInstanceIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationInstanceIdsReverseMatch(v bool) *CreateCenRouteMapRequest {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationRegionIds(v []*string) *CreateCenRouteMapRequest {
	s.DestinationRegionIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetDestinationRouteTableIds(v []*string) *CreateCenRouteMapRequest {
	s.DestinationRouteTableIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetMapResult(v string) *CreateCenRouteMapRequest {
	s.MapResult = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetMatchAddressType(v string) *CreateCenRouteMapRequest {
	s.MatchAddressType = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetMatchAsns(v []*int64) *CreateCenRouteMapRequest {
	s.MatchAsns = v
	return s
}

func (s *CreateCenRouteMapRequest) SetMatchCommunitySet(v []*string) *CreateCenRouteMapRequest {
	s.MatchCommunitySet = v
	return s
}

func (s *CreateCenRouteMapRequest) SetNextPriority(v int32) *CreateCenRouteMapRequest {
	s.NextPriority = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetOperateCommunitySet(v []*string) *CreateCenRouteMapRequest {
	s.OperateCommunitySet = v
	return s
}

func (s *CreateCenRouteMapRequest) SetOwnerAccount(v string) *CreateCenRouteMapRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetOwnerId(v int64) *CreateCenRouteMapRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetPreference(v int32) *CreateCenRouteMapRequest {
	s.Preference = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetPrependAsPath(v []*int64) *CreateCenRouteMapRequest {
	s.PrependAsPath = v
	return s
}

func (s *CreateCenRouteMapRequest) SetPriority(v int32) *CreateCenRouteMapRequest {
	s.Priority = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetResourceOwnerAccount(v string) *CreateCenRouteMapRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetResourceOwnerId(v int64) *CreateCenRouteMapRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetRouteTypes(v []*string) *CreateCenRouteMapRequest {
	s.RouteTypes = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceChildInstanceTypes(v []*string) *CreateCenRouteMapRequest {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceInstanceIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceInstanceIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceInstanceIdsReverseMatch(v bool) *CreateCenRouteMapRequest {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceRegionIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceRegionIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetSourceRouteTableIds(v []*string) *CreateCenRouteMapRequest {
	s.SourceRouteTableIds = v
	return s
}

func (s *CreateCenRouteMapRequest) SetTransitRouterRouteTableId(v string) *CreateCenRouteMapRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *CreateCenRouteMapRequest) SetTransmitDirection(v string) *CreateCenRouteMapRequest {
	s.TransmitDirection = &v
	return s
}

func (s *CreateCenRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
