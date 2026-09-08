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
	// The match mode of the AS path list. Valid values:
	//
	// - **Include**: fuzzy match. A match is successful if the AS path in the match condition overlaps with the AS path of the route to be matched.
	//
	// - **Complete**: exact match. A match is successful only if the AS path in the match condition is the same as the AS path of the route to be matched.
	//
	// example:
	//
	// Include
	AsPathMatchMode *string `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	// The instance ID of the Cloud Enterprise Network (CEN) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the region where the routing policy is applied.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	CenRegionId *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	// The match mode of the prefix list. Valid values:
	//
	// - **Include**: fuzzy match. A match is successful if the route prefix in the match condition contains the route prefix of the route to be matched.
	//
	//  For example, a policy that defines 10.10.0.0/16 can fuzzy match the route 10.10.1.0/24.
	//
	// - **Complete**: exact match. A match is successful only if the route prefix in the match condition is the same as the route prefix of the route to be matched.
	//
	//  For example, a policy that defines 10.10.0.0/16 can exact match only the route 10.10.0.0/16.
	//
	// example:
	//
	// Include
	CidrMatchMode *string `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	// The match mode of the Community. Valid values:
	//
	// - **Include**: fuzzy match. A match is successful if the Community in the match condition overlaps with the Community of the route to be matched.
	//
	// - **Complete**: exact match. A match is successful only if the Community in the match condition is the same as the Community of the route to be matched.
	//
	// - **Contain**: contains match. A match is successful only if the Community of the route to be matched contains all the Communities specified in the match condition.
	//
	// example:
	//
	// Include
	CommunityMatchMode *string `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	// The action to perform on the Community. Valid values:
	//
	// - **Additive**: adds the Community to the route.
	//
	// - **Replace**: replaces the original Community of the route.
	//
	// This parameter specifies the action to perform on a route after the route matches the match conditions.
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
	//     > If the IPsec connection or SSL server is bound to a VPN gateway instance and is connected to a transit router instance through the VPC associated with the VPN gateway instance, this parameter does not take effect. This parameter takes effect only when the IPsec connection is directly bound to a transit router instance.
	//
	// The destination instance type list takes effect only when the direction of the routing policy is Export from Regional Gateway and the destination instance types are instance types in the local region.
	//
	// example:
	//
	// VPC
	DestinationChildInstanceTypes []*string `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Repeated"`
	// The prefix list that the route must match.
	//
	// IP address ranges in the prefix list are in CIDR format. A maximum of 64 IP address ranges can be specified.
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
	// A maximum of 64 instance IDs can be specified.
	//
	// > The destination instance ID list takes effect only when the direction of the routing policy is Export from Regional Gateway and the destination instance IDs are in the local region.
	//
	// example:
	//
	// vpc-avcdsg34ds****
	DestinationInstanceIds []*string `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the reverse match mode for the destination instance ID list. Valid values:
	//
	// - **false*	- (default): No. A match is successful if the destination instance ID of the route is in **DestinationInstanceIds.N**.
	//
	// - **true**: Yes. A match is successful if the destination instance ID of the route is not in **DestinationInstanceIds.N**.
	//
	// example:
	//
	// false
	DestinationInstanceIdsReverseMatch *bool `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	// The list of destination region IDs that the route must match. A maximum of 64 region IDs can be specified.
	DestinationRegionIds []*string `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Repeated"`
	// The list of destination route table IDs that the route must match. A maximum of 64 route table IDs can be specified.
	//
	// > The destination route table ID list takes effect only when the direction of the routing policy is Export from Regional Gateway and the destination route table IDs are route table IDs of network instances in the local region.
	//
	// example:
	//
	// vtb-adfg53c322v****
	DestinationRouteTableIds []*string `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Repeated"`
	// The action to perform on a route that matches all the match conditions. Valid values:
	//
	// - **Permit**: The route is permitted to pass.
	//
	// - **Deny**: The route is denied from passing.
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
	// > Only AS SEQUENCE is supported. AS SET, AS CONFED SEQUENCE, and AS CONFED SET are not supported. Specifically, only AS number lists are supported. Sets and sublists are not supported.
	//
	// example:
	//
	// 65501
	MatchAsns []*int32 `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Repeated"`
	// The Community set that the route must match.
	//
	// Each Community is in the n:m format, where the value ranges of n and m are **1*	- to **65535**. Communities must comply with RFC 1997. Large Communities (RFC 8092) are not supported.
	//
	// A maximum of 64 Communities can be specified.
	//
	// > Incorrect Community configurations may cause routes to fail to be advertised to on-premises data centers.
	//
	// example:
	//
	// 65501:1
	MatchCommunitySet []*string `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Repeated"`
	// Policy priority of the next associated routing policy.
	//
	// - You can set policy priority of the next associated routing policy only when **MapResult*	- is set to **Permit**. Only routes that are permitted to pass continue to match the next associated routing policy.
	//
	// - The next associated routing policy must have the same region and direction as the current routing policy.
	//
	// - Policy priority of the next associated routing policy must be lower than (a number greater than) policy priority of the current routing policy.
	//
	// example:
	//
	// 20
	NextPriority *int32 `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	// The Community set to be executed.
	//
	// Each Community is in the n:m format, where the value ranges of n and m are **1*	- to **65535**. Communities must comply with RFC 1997. Large Communities (RFC 8092) are not supported.
	//
	// A maximum of 32 Communities can be specified.
	//
	// > Incorrect Community configurations may cause routes to fail to be advertised to on-premises data centers.
	//
	// example:
	//
	// 65501:1
	OperateCommunitySet []*string `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of the route to be modified.
	//
	// Valid values: **1*	- to **100**. The default priority of a route is **50**. A smaller value indicates a higher priority.
	//
	// This parameter specifies the action to perform on a route after the route matches the match conditions.
	//
	// example:
	//
	// 22
	Preference *int32 `json:"Preference,omitempty" xml:"Preference,omitempty"`
	// The AS path that is prepended when the regional gateway receives or publishes route entries.
	//
	// The requirements for configuring the prepended AS path vary based on the direction of the routing policy:
	//
	// - When configuring the prepended AS path for the Import to Regional Gateway direction, you must configure the source instance ID list and source region in the match conditions, and the source region must be the same as the region where the routing policy is applied.
	//
	// - When configuring the prepended AS path for the Export from Regional Gateway direction, you must configure the destination instance ID list in the match conditions.
	//
	//
	// This parameter specifies the action to execute on a route after the route matches the match conditions.
	//
	// example:
	//
	// 65501
	PrependAsPath []*int64 `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Repeated"`
	// Policy priority of the routing policy. Valid values: **1*	- to **100**. A smaller value indicates a higher priority.
	//
	// > Policy priority of routing policies in the same region and with the same direction must be unique. When the system executes routing policies, it starts matching conditional statements from the routing policy with the smallest priority number. Specify policy priority based on the expected matching order.
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
	// The list of routing types that the route must match. The following routing types are supported:
	//
	// - **System**: system routes that are automatically generated by the system.
	//
	// - **Custom**: custom routes that are manually added by users.
	//
	// - **BGP**: BGP routes that are propagated through the BGP routing protocol.
	//
	// example:
	//
	// System
	RouteTypes []*string `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Repeated"`
	// The list of source instance types that the route must match. The following instance types are supported:
	//
	// - **VPC**: VPC instance.
	//
	// - **VBR**: border router instance.
	//
	// - **CCN**: CCN instance.
	//
	// - **VPN**: VPN gateway instance or IPsec connection.
	//
	//
	//
	//     - If the IPsec connection or SSL server is attached to a VPN gateway instance, the VPC associated with the VPN gateway instance must be connected to a transit router instance, and the VPN gateway instance must run the BGP dynamic routing protocol for this parameter to take effect.
	//
	//     - If the IPsec connection is directly attached to a transit router instance, this parameter takes effect.
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
	// A maximum of 64 instance IDs can be specified.
	//
	// example:
	//
	// vpc-afsfdf5435vcvc****
	SourceInstanceIds []*string `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to use the reverse match mode for the source instance ID list. Valid values:
	//
	// - **false*	- (default): No. A match is successful if the source instance ID of the route is in **SourceInstanceIds.N**.
	//
	// - **true**: Yes. A match is successful if the source instance ID of the route is not in **SourceInstanceIds.N**.
	//
	// example:
	//
	// false
	SourceInstanceIdsReverseMatch *bool `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	// The list of source region IDs that the route must match. A maximum of 64 region IDs can be specified.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// example:
	//
	// cn-beijing
	SourceRegionIds []*string `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Repeated"`
	// The list of source route table IDs that the route must match. A maximum of 64 route table IDs can be specified.
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
