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
	// The match mode of the AS path list. Valid values:
	//
	// - **Include**: fuzzy match. A match is successful if the AS path in the match condition overlaps with the AS path of the route being matched.
	//
	// - **Complete**: exact match. A match is successful only if the AS path in the match condition is the same as the AS path of the route being matched.
	//
	// example:
	//
	// Include
	AsPathMatchMode *string `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	// The instance ID of the Cloud Enterprise Network (CEN).
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the region to which the routing policy is applied.
	//
	// You can call [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	CenRegionId *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	// The match mode of the prefix list. Valid values:
	//
	// - **Include**: fuzzy match. A match is successful if the route prefix in the match condition contains the route prefix of the route being matched.
	//
	//  For example, a policy that defines 10.10.0.0/16 can fuzzy match the route 10.10.1.0/24.
	//
	// - **Complete**: exact match. A match is successful only if the route prefix in the match condition is the same as the route prefix of the route being matched.
	//
	//  For example, a policy that defines 10.10.0.0/16 can only exact match the route 10.10.0.0/16.
	//
	// example:
	//
	// Include
	CidrMatchMode *string `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	// The match mode of the Community. Valid values:
	//
	// - **Include**: fuzzy match. A match is successful if the Community in the match condition overlaps with the Community of the route being matched.
	//
	// - **Complete**: exact match. A match is successful only if the Community in the match condition is the same as the Community of the route being matched.
	//
	// - **Contain**: inclusive match. A match is successful only if the Community of the route being matched contains all the Communities specified in the match condition.
	//
	// example:
	//
	// Include
	CommunityMatchMode *string `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	// The action to perform on the Community. Valid values:
	//
	// - **Additive**: adds a Community to the route.
	//
	// - **Replace**: replaces the existing Community of the route.
	//
	// This parameter specifies the action to perform after a route matches the condition.
	//
	// example:
	//
	// Additive
	CommunityOperateMode *string `json:"CommunityOperateMode,omitempty" xml:"CommunityOperateMode,omitempty"`
	// The description of the routing policy.
	//
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of destination instance types that the route must match. The following instance types are supported:
	//
	// - **VPC**: VPC instance.
	//
	// - **VBR**: VBR instance.
	//
	// - **CCN**: CCN instance.
	//
	// - **VPN**: IPsec connection.
	//
	//     > If an IPsec connection or SSL server is bound to a VPN gateway instance and is connected to a transit router instance through the VPC associated with the VPN gateway instance, this parameter does not take effect. This parameter takes effect only when an IPsec connection is directly bound to a transit router instance.
	//
	// You can specify multiple instance types.
	//
	// >The destination instance type list takes effect only when the routing policy direction is outbound from the regional gateway and the destination instance types are instance types in the local region.
	//
	// example:
	//
	// VPC
	DestinationChildInstanceTypes []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	// The prefix list that the route must match.
	//
	// IP address ranges in the prefix list are in CIDR format. You can specify up to 64 IP address ranges.
	//
	// Both IPv4 and IPv6 formats are supported.
	//
	// example:
	//
	// 10.10.10.0/24
	DestinationCidrBlocks []*string `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Repeated"`
	// The list of destination instance IDs that the route must match. The following types of instance IDs are supported:
	//
	// - Virtual Private Cloud (VPC) instance ID
	//
	// - Virtual Border Router (VBR) instance ID
	//
	// - Cloud Connect Network (CCN) instance ID
	//
	// - Smart Access Gateway instance ID
	//
	// - IPsec connection ID
	//
	// You can specify up to 64 instance IDs.
	//
	// >The destination instance ID list takes effect only when the routing policy direction is outbound from the regional gateway and the destination instance IDs are instance IDs in the local region.
	//
	// example:
	//
	// vpc-afrfs434465fdf****
	DestinationInstanceIds []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the exclude matching mode for the destination instance ID list. Valid values:
	//
	// - **false*	- (default): no. A match is successful if the destination instance ID of the route is in the **DestinationInstanceIds.N*	- list.
	//
	// - **true**: yes. A match is successful if the destination instance ID of the route is not in the **DestinationInstanceIds.N*	- list.
	//
	// example:
	//
	// false
	DestinationInstanceIdsReverseMatch *bool `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	// The list of destination region IDs that the route must match. You can specify up to 64 region IDs.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// The list of destination route table IDs that the route must match. You can specify up to 64 route table IDs.
	//
	// >The destination route table ID list takes effect only when the routing policy direction is outbound from the regional gateway and the destination route table IDs are route table IDs of network instances in the local region.
	//
	// example:
	//
	// vtb-adefrgtr144vf****
	DestinationRouteTableIds []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	// The action to perform after all conditions are matched. Valid values:
	//
	// - **Permit**: permits the matched routes.
	//
	// - **Deny**: denies the matched routes.
	//
	// This parameter is required.
	//
	// example:
	//
	// Permit
	MapResult *string `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	// The IP address type that the route must match. Valid values:
	//
	// - **IPv4**: matches only IPv4 routes.
	//
	// - **IPv6**: matches only IPv6 routes.
	//
	// This parameter can be left empty, which indicates that all types of routes are matched.
	//
	// example:
	//
	// IPv4
	MatchAddressType *string `json:"MatchAddressType,omitempty" xml:"MatchAddressType,omitempty"`
	// The AS path list that the route must match.
	//
	// You can specify up to 64 AS numbers.
	//
	// > Only AS SEQUENCE is supported. AS SET, AS CONFED SEQUENCE, and AS CONFED SET are not supported. This means that only AS number lists are supported, not sets or sublists.
	//
	// example:
	//
	// 65501
	MatchAsns []*int64 `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	// The Community set that the route must match.
	//
	// Each Community is in the n:m format, where the value ranges of n and m are **1*	- to **65535**. Communities must comply with RFC 1997. Large Communities (RFC 8092) are not supported.
	//
	// You can specify up to 64 Communities.
	//
	// > Incorrect Community configurations may cause routes to fail to be advertised to on-premises data centers.
	//
	// example:
	//
	// 65501:1
	MatchCommunitySet []*string `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Repeated"`
	// Policy priority of the next associated routing policy.
	//
	// - You can set policy priority of the next associated routing policy only when **MapResult*	- is set to **Permit**. Only routes that are permitted continue to match the next associated routing policy.
	//
	// - The next associated routing policy must have the same region and direction as the current routing policy.
	//
	// - Policy priority of the next associated routing policy must be lower than policy priority of the current routing policy.
	//
	// example:
	//
	// 20
	NextPriority *int32 `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	// The Community set to be executed.
	//
	// Each Community is in the n:m format, where the value ranges of n and m are **1*	- to **65535**. Communities must comply with RFC 1997. Large Communities (RFC 8092) are not supported.
	//
	// You can specify up to 32 Communities.
	//
	// > Incorrect Community configurations may cause routes to fail to be advertised to on-premises data centers.
	//
	// example:
	//
	// 65501:1
	OperateCommunitySet []*string `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The modified priority of the route.
	//
	// Valid values: **1*	- to **100**. The default priority of a route is **50**. A smaller value indicates a higher priority.
	//
	// This parameter specifies the action to perform after a route matches the condition.
	//
	// example:
	//
	// 50
	Preference *int32 `json:"Preference,omitempty" xml:"Preference,omitempty"`
	// The AS path that is prepended when the regional gateway receives or publishes route entries.
	//
	// The requirements for configuring the prepended AS path vary based on the routing policy direction:
	//
	// - When the direction is inbound to the regional gateway, the match condition must include the source instance ID list and source region, and the source region must be the same as the region to which the routing policy is applied.
	//
	// - When the direction is outbound from the regional gateway, the match condition must include the destination instance ID list.
	//
	//
	// This parameter specifies the action to execute after a route matches the condition. You can specify up to 32 AS numbers.
	//
	// example:
	//
	// 65501
	PrependAsPath []*int64 `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Repeated"`
	// Policy priority of the routing policy. Valid values: **1*	- to **100**. A smaller value indicates a higher priority.
	//
	// > Policy priority of routing policies in the same region and with the same direction must be unique. When a routing policy is executed, the system starts matching conditional statements from the routing policy with the smallest priority value. Specify policy priority based on the expected matching order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Priority             *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of routing types that the route must match. The following routing types are supported:
	//
	// - **System**: system routes that are automatically generated by the system.
	//
	// - **Custom**: custom routes that are manually added by users.
	//
	// - **BGP**: BGP routes that are propagated through the BGP routing protocol.
	//
	// You can specify multiple routing types.
	//
	// example:
	//
	// System
	RouteTypes []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	// The list of source instance types that the route must match. The following instance types are supported:
	//
	// - **VPC**: VPC instance.
	//
	// - **VBR**: virtual border router instance.
	//
	// - **CCN**: CCN instance.
	//
	// - **VPN**: VPN gateway instance or IPsec connection.
	//
	//     - If an IPsec connection or SSL server is attached to a VPN gateway instance, the VPC associated with the VPN gateway instance must be connected to a transit router instance, and the VPN gateway instance must run the BGP dynamic routing protocol for this parameter to take effect.
	//
	//     - If an IPsec connection is directly attached to a transit router instance, this parameter takes effect.
	//
	// You can specify multiple instance types.
	//
	// example:
	//
	// VPC
	SourceChildInstanceTypes []*string `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Repeated"`
	// The list of source instance IDs that the route must match. The following types of instance IDs are supported:
	//
	// - Virtual Private Cloud (VPC) instance ID
	//
	// - Virtual Border Router (VBR) instance ID
	//
	// - Cloud Connect Network (CCN) instance ID
	//
	// - Smart Access Gateway instance ID
	//
	// - IPsec connection ID
	//
	// You can specify up to 64 instance IDs.
	//
	// example:
	//
	// vpc-adeg3544fdf34vf****
	SourceInstanceIds []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the exclude matching mode for the source instance ID list. Valid values:
	//
	// - **false*	- (default): no. A match is successful if the source instance ID of the route is in the **SourceInstanceIds.N*	- list.
	//
	// - **true**: yes. A match is successful if the source instance ID of the route is not in the **SourceInstanceIds.N*	- list.
	//
	// example:
	//
	// false
	SourceInstanceIdsReverseMatch *bool `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	// The list of source region IDs that the route must match. You can specify up to 64 region IDs.
	//
	// You can call [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) to query region IDs.
	//
	// example:
	//
	// cn-beijing
	SourceRegionIds []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	// The list of source route table IDs that the route must match. You can specify up to 64 route table IDs.
	//
	// example:
	//
	// vtb-adfr233vf34rvd4****
	SourceRouteTableIds []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
	// The route table ID of the transit router.
	//
	// If you do not specify a route table ID, the routing policy is automatically associated with the default route table of the transit router.
	//
	// example:
	//
	// vtb-gw8nx3515m1mbd1z1****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	// The direction in which the routing policy is applied. Valid values:
	//
	// - **RegionIn**: the inbound direction of the regional gateway. Routes are transmitted to the CEN regional gateway.
	//
	//  For example, a route is advertised from a network instance in the local region to the local regional gateway, or a route is advertised from another region to the local regional gateway.
	//
	// - **RegionOut**: the outbound direction of the regional gateway. Routes are transmitted from the CEN regional gateway.
	//
	//  For example, a route is advertised from the local regional gateway to a network instance in the local region, or to a regional gateway in another region.
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
