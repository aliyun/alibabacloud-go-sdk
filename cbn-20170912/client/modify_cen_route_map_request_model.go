// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenRouteMapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsPathMatchMode(v string) *ModifyCenRouteMapRequest
	GetAsPathMatchMode() *string
	SetCenId(v string) *ModifyCenRouteMapRequest
	GetCenId() *string
	SetCenRegionId(v string) *ModifyCenRouteMapRequest
	GetCenRegionId() *string
	SetCidrMatchMode(v string) *ModifyCenRouteMapRequest
	GetCidrMatchMode() *string
	SetCommunityMatchMode(v string) *ModifyCenRouteMapRequest
	GetCommunityMatchMode() *string
	SetCommunityOperateMode(v string) *ModifyCenRouteMapRequest
	GetCommunityOperateMode() *string
	SetDescription(v string) *ModifyCenRouteMapRequest
	GetDescription() *string
	SetDestinationChildInstanceTypes(v []*string) *ModifyCenRouteMapRequest
	GetDestinationChildInstanceTypes() []*string
	SetDestinationCidrBlocks(v []*string) *ModifyCenRouteMapRequest
	GetDestinationCidrBlocks() []*string
	SetDestinationInstanceIds(v []*string) *ModifyCenRouteMapRequest
	GetDestinationInstanceIds() []*string
	SetDestinationInstanceIdsReverseMatch(v bool) *ModifyCenRouteMapRequest
	GetDestinationInstanceIdsReverseMatch() *bool
	SetDestinationRegionIds(v []*string) *ModifyCenRouteMapRequest
	GetDestinationRegionIds() []*string
	SetDestinationRouteTableIds(v []*string) *ModifyCenRouteMapRequest
	GetDestinationRouteTableIds() []*string
	SetMapResult(v string) *ModifyCenRouteMapRequest
	GetMapResult() *string
	SetMatchAddressType(v string) *ModifyCenRouteMapRequest
	GetMatchAddressType() *string
	SetMatchAsns(v []*int32) *ModifyCenRouteMapRequest
	GetMatchAsns() []*int32
	SetMatchCommunitySet(v []*string) *ModifyCenRouteMapRequest
	GetMatchCommunitySet() []*string
	SetNextPriority(v int32) *ModifyCenRouteMapRequest
	GetNextPriority() *int32
	SetOperateCommunitySet(v []*string) *ModifyCenRouteMapRequest
	GetOperateCommunitySet() []*string
	SetOwnerAccount(v string) *ModifyCenRouteMapRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCenRouteMapRequest
	GetOwnerId() *int64
	SetPreference(v int32) *ModifyCenRouteMapRequest
	GetPreference() *int32
	SetPrependAsPath(v []*int64) *ModifyCenRouteMapRequest
	GetPrependAsPath() []*int64
	SetPriority(v int32) *ModifyCenRouteMapRequest
	GetPriority() *int32
	SetResourceOwnerAccount(v string) *ModifyCenRouteMapRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCenRouteMapRequest
	GetResourceOwnerId() *int64
	SetRouteMapId(v string) *ModifyCenRouteMapRequest
	GetRouteMapId() *string
	SetRouteTypes(v []*string) *ModifyCenRouteMapRequest
	GetRouteTypes() []*string
	SetSourceChildInstanceTypes(v []*string) *ModifyCenRouteMapRequest
	GetSourceChildInstanceTypes() []*string
	SetSourceInstanceIds(v []*string) *ModifyCenRouteMapRequest
	GetSourceInstanceIds() []*string
	SetSourceInstanceIdsReverseMatch(v bool) *ModifyCenRouteMapRequest
	GetSourceInstanceIdsReverseMatch() *bool
	SetSourceRegionIds(v []*string) *ModifyCenRouteMapRequest
	GetSourceRegionIds() []*string
	SetSourceRouteTableIds(v []*string) *ModifyCenRouteMapRequest
	GetSourceRouteTableIds() []*string
}

type ModifyCenRouteMapRequest struct {
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
	// For example, if you set the match condition to 10.10.0.0/16 and fuzzy match is applied, the route whose prefix is 10.10.1.0/24 meets the match condition.
	//
	// 	- **Complete**: exact match. A route is a match only if the route prefix is the same as the prefix specified in the match condition.
	//
	// For example, if you set the match condition to 10.10.0.0/16 and exact match is applied, only the route whose prefix is 10.10.0.0/16 meets the match condition.
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
	//     > This parameter does not take effect if the IPsec-VPN connection or SSL client is associated with a transit router through a VPN gateway and a VPC. This parameter takes effect only if the IPsec connection is directly connected to the transit router.
	//
	// The destination network instance types are valid only if the routing policy is applied to scenarios where routes are advertised from the gateway in the current region to network instances in the current region.
	//
	// example:
	//
	// VPC
	DestinationChildInstanceTypes []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	// The prefix list against which routes are matched.
	//
	// You must specify the IP addresses in CIDR notation. You can enter at most 32 CIDR blocks.
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
	// vpc-avcdsg34ds****
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
	// The IDs of the destination route tables to which the routes belong. You can enter at most 32 route table IDs.
	//
	// > The destination route table IDs take effect only when Direction is set to Export from Regional Gateway and the destination route tables belong to network instances deployed in the current region.
	//
	// example:
	//
	// vtb-adfg53c322v****
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
	// The AS paths against which routes are matched.
	//
	// > Only the AS-SEQUENCE parameter is supported. The AS-SET, AS-CONFED-SEQUENCE, and AS-CONFED-SET parameters are not supported. In other words, only the AS number list is supported. Sets and sub-lists are not supported.
	//
	// example:
	//
	// 65501
	MatchAsns []*int32 `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	// The community against which routes are matched.
	//
	// Specify the community in the format of n:m. Valid values of n and m: **1*	- to **65535**. Each community must comply with the RFC 1997 standard. The RFC 8092 standard that defines BGP large communities is not supported.
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
	// 	- The priority of the routing policy to be associated must be lower than the priority of the current routing policy.
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
	// 22
	Preference *int32 `json:"Preference,omitempty" xml:"Preference,omitempty"`
	// The AS paths that are prepended by using an action statement when regional gateways receive or advertise routes.
	//
	// The AS paths vary based on the direction in which the routing policy is applied:
	//
	// 	- If AS paths are prepended to a routing policy that is applied in the inbound direction, you must specify source network instance IDs and the source region in the match condition. In addition, the source region must be the same as the region where the routing policy is applied.
	//
	// 	- If AS paths are prepended to a routing policy that is applied in the outbound direction, you must specify destination network instance IDs in the match condition.
	//
	// This parameter specifies the action to be performed when a route meets the match condition.
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
	// 10
	Priority             *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the routing policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cenrmap-abcdedfghij****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	// The type of route to be matched against the match condition. The following route types are supported:
	//
	// 	- **System**: system routes that are automatically generated by the system.
	//
	// 	- **Custom**: custom routes that are manually added.
	//
	// 	- **BGP**: routes that are advertised over BGP.
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
	// 	- **VPN*	- :VPN gateway or IPsec-VPN connection
	//
	//     	- If the IPsec-VPN connection or SSL client is associated with a VPN gateway, the VPC associated with the VPN gateway must be connected to a transit router, and the VPN gateway must use Border Gateway Protocol (BGP) dynamic routing. Otherwise, this parameter cannot take effect.
	//
	//     	- This parameter takes effect if the IPsec connection is directly connected to a transit router.
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
	// vpc-afsfdf5435vcvc****
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
	// The IDs of the source regions to which the routes belong. You can enter at most 32 region IDs.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	SourceRegionIds []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	// The IDs of the source route tables to which the routes belong. You can enter at most 32 route table IDs.
	//
	// example:
	//
	// vtb-acdbvtbr342cd****
	SourceRouteTableIds []*string `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Repeated"`
}

func (s ModifyCenRouteMapRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenRouteMapRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenRouteMapRequest) GetAsPathMatchMode() *string {
	return s.AsPathMatchMode
}

func (s *ModifyCenRouteMapRequest) GetCenId() *string {
	return s.CenId
}

func (s *ModifyCenRouteMapRequest) GetCenRegionId() *string {
	return s.CenRegionId
}

func (s *ModifyCenRouteMapRequest) GetCidrMatchMode() *string {
	return s.CidrMatchMode
}

func (s *ModifyCenRouteMapRequest) GetCommunityMatchMode() *string {
	return s.CommunityMatchMode
}

func (s *ModifyCenRouteMapRequest) GetCommunityOperateMode() *string {
	return s.CommunityOperateMode
}

func (s *ModifyCenRouteMapRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCenRouteMapRequest) GetDestinationChildInstanceTypes() []*string {
	return s.DestinationChildInstanceTypes
}

func (s *ModifyCenRouteMapRequest) GetDestinationCidrBlocks() []*string {
	return s.DestinationCidrBlocks
}

func (s *ModifyCenRouteMapRequest) GetDestinationInstanceIds() []*string {
	return s.DestinationInstanceIds
}

func (s *ModifyCenRouteMapRequest) GetDestinationInstanceIdsReverseMatch() *bool {
	return s.DestinationInstanceIdsReverseMatch
}

func (s *ModifyCenRouteMapRequest) GetDestinationRegionIds() []*string {
	return s.DestinationRegionIds
}

func (s *ModifyCenRouteMapRequest) GetDestinationRouteTableIds() []*string {
	return s.DestinationRouteTableIds
}

func (s *ModifyCenRouteMapRequest) GetMapResult() *string {
	return s.MapResult
}

func (s *ModifyCenRouteMapRequest) GetMatchAddressType() *string {
	return s.MatchAddressType
}

func (s *ModifyCenRouteMapRequest) GetMatchAsns() []*int32 {
	return s.MatchAsns
}

func (s *ModifyCenRouteMapRequest) GetMatchCommunitySet() []*string {
	return s.MatchCommunitySet
}

func (s *ModifyCenRouteMapRequest) GetNextPriority() *int32 {
	return s.NextPriority
}

func (s *ModifyCenRouteMapRequest) GetOperateCommunitySet() []*string {
	return s.OperateCommunitySet
}

func (s *ModifyCenRouteMapRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCenRouteMapRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCenRouteMapRequest) GetPreference() *int32 {
	return s.Preference
}

func (s *ModifyCenRouteMapRequest) GetPrependAsPath() []*int64 {
	return s.PrependAsPath
}

func (s *ModifyCenRouteMapRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyCenRouteMapRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCenRouteMapRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCenRouteMapRequest) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *ModifyCenRouteMapRequest) GetRouteTypes() []*string {
	return s.RouteTypes
}

func (s *ModifyCenRouteMapRequest) GetSourceChildInstanceTypes() []*string {
	return s.SourceChildInstanceTypes
}

func (s *ModifyCenRouteMapRequest) GetSourceInstanceIds() []*string {
	return s.SourceInstanceIds
}

func (s *ModifyCenRouteMapRequest) GetSourceInstanceIdsReverseMatch() *bool {
	return s.SourceInstanceIdsReverseMatch
}

func (s *ModifyCenRouteMapRequest) GetSourceRegionIds() []*string {
	return s.SourceRegionIds
}

func (s *ModifyCenRouteMapRequest) GetSourceRouteTableIds() []*string {
	return s.SourceRouteTableIds
}

func (s *ModifyCenRouteMapRequest) SetAsPathMatchMode(v string) *ModifyCenRouteMapRequest {
	s.AsPathMatchMode = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCenId(v string) *ModifyCenRouteMapRequest {
	s.CenId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCenRegionId(v string) *ModifyCenRouteMapRequest {
	s.CenRegionId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCidrMatchMode(v string) *ModifyCenRouteMapRequest {
	s.CidrMatchMode = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCommunityMatchMode(v string) *ModifyCenRouteMapRequest {
	s.CommunityMatchMode = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetCommunityOperateMode(v string) *ModifyCenRouteMapRequest {
	s.CommunityOperateMode = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDescription(v string) *ModifyCenRouteMapRequest {
	s.Description = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationChildInstanceTypes(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationChildInstanceTypes = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationCidrBlocks(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationCidrBlocks = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationInstanceIds(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationInstanceIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationInstanceIdsReverseMatch(v bool) *ModifyCenRouteMapRequest {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationRegionIds(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationRegionIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetDestinationRouteTableIds(v []*string) *ModifyCenRouteMapRequest {
	s.DestinationRouteTableIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMapResult(v string) *ModifyCenRouteMapRequest {
	s.MapResult = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMatchAddressType(v string) *ModifyCenRouteMapRequest {
	s.MatchAddressType = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMatchAsns(v []*int32) *ModifyCenRouteMapRequest {
	s.MatchAsns = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetMatchCommunitySet(v []*string) *ModifyCenRouteMapRequest {
	s.MatchCommunitySet = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetNextPriority(v int32) *ModifyCenRouteMapRequest {
	s.NextPriority = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOperateCommunitySet(v []*string) *ModifyCenRouteMapRequest {
	s.OperateCommunitySet = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOwnerAccount(v string) *ModifyCenRouteMapRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetOwnerId(v int64) *ModifyCenRouteMapRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPreference(v int32) *ModifyCenRouteMapRequest {
	s.Preference = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPrependAsPath(v []*int64) *ModifyCenRouteMapRequest {
	s.PrependAsPath = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetPriority(v int32) *ModifyCenRouteMapRequest {
	s.Priority = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetResourceOwnerAccount(v string) *ModifyCenRouteMapRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetResourceOwnerId(v int64) *ModifyCenRouteMapRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetRouteMapId(v string) *ModifyCenRouteMapRequest {
	s.RouteMapId = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetRouteTypes(v []*string) *ModifyCenRouteMapRequest {
	s.RouteTypes = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceChildInstanceTypes(v []*string) *ModifyCenRouteMapRequest {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceInstanceIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceInstanceIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceInstanceIdsReverseMatch(v bool) *ModifyCenRouteMapRequest {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceRegionIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceRegionIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) SetSourceRouteTableIds(v []*string) *ModifyCenRouteMapRequest {
	s.SourceRouteTableIds = v
	return s
}

func (s *ModifyCenRouteMapRequest) Validate() error {
	return dara.Validate(s)
}
