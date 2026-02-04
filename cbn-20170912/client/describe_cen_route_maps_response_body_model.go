// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenRouteMapsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeCenRouteMapsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenRouteMapsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCenRouteMapsResponseBody
	GetRequestId() *string
	SetRouteMaps(v *DescribeCenRouteMapsResponseBodyRouteMaps) *DescribeCenRouteMapsResponseBody
	GetRouteMaps() *DescribeCenRouteMapsResponseBodyRouteMaps
	SetTotalCount(v int32) *DescribeCenRouteMapsResponseBody
	GetTotalCount() *int32
}

type DescribeCenRouteMapsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 24CE1987-D1D1-5324-9BAD-2750B60E6ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the routing policy.
	RouteMaps *DescribeCenRouteMapsResponseBodyRouteMaps `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenRouteMapsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenRouteMapsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenRouteMapsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenRouteMapsResponseBody) GetRouteMaps() *DescribeCenRouteMapsResponseBodyRouteMaps {
	return s.RouteMaps
}

func (s *DescribeCenRouteMapsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenRouteMapsResponseBody) SetPageNumber(v int32) *DescribeCenRouteMapsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetPageSize(v int32) *DescribeCenRouteMapsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetRequestId(v string) *DescribeCenRouteMapsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetRouteMaps(v *DescribeCenRouteMapsResponseBodyRouteMaps) *DescribeCenRouteMapsResponseBody {
	s.RouteMaps = v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) SetTotalCount(v int32) *DescribeCenRouteMapsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBody) Validate() error {
	if s.RouteMaps != nil {
		if err := s.RouteMaps.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenRouteMapsResponseBodyRouteMaps struct {
	RouteMap []*DescribeCenRouteMapsResponseBodyRouteMapsRouteMap `json:"RouteMap,omitempty" xml:"RouteMap,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMaps) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMaps) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) GetRouteMap() []*DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	return s.RouteMap
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) SetRouteMap(v []*DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) *DescribeCenRouteMapsResponseBodyRouteMaps {
	s.RouteMap = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMaps) Validate() error {
	if s.RouteMap != nil {
		for _, item := range s.RouteMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMap struct {
	// The match method that is used to match routes based on the AS path.
	//
	// 	- **Include**: fuzzy match. A route is a match if the AS path of the route overlaps with the AS path specified in the match condition.
	//
	// 	- **Complete**: exact match. A route is a match only if the AS path of the route is the same as an AS path specified in the match condition.
	//
	// example:
	//
	// Include
	AsPathMatchMode *string `json:"AsPathMatchMode,omitempty" xml:"AsPathMatchMode,omitempty"`
	// The CEN instance ID.
	//
	// example:
	//
	// cen-wx12mmlt17ld82****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The region ID of the routing policy.
	//
	// example:
	//
	// cn-hangzhou
	CenRegionId *string `json:"CenRegionId,omitempty" xml:"CenRegionId,omitempty"`
	// The match method that is used to evaluate routes based on the prefix. Valid values:
	//
	// 	- **Include**: fuzzy match. A route is a match if the route prefix is included in the match conditions.
	//
	// For example, if you set the match condition to 10.10.0.0/16 and fuzzy match is applied, the route whose prefix is 10.10.1.0/24 meets the match condition.
	//
	// 	- **Complete**: exact match. A route is a match only if the route prefix is the same as the prefix specified in the match condition.
	//
	// For example, if you set the match condition to 10.10.0.0/16 and exact match is enabled, a route is a match only if the prefix is 10.10.0.0/16.
	//
	// example:
	//
	// Include
	CidrMatchMode *string `json:"CidrMatchMode,omitempty" xml:"CidrMatchMode,omitempty"`
	// The match method that is used to match routes against the community.
	//
	// 	- **Include**: fuzzy match. A route is a match if the community of the route overlaps with the community specified in the match condition.
	//
	// 	- **Complete**: exact match. A route meets the match condition only if the community of the route is the same as the community specified in the match condition.
	//
	// example:
	//
	// Include
	CommunityMatchMode *string `json:"CommunityMatchMode,omitempty" xml:"CommunityMatchMode,omitempty"`
	// The action that is performed on the community of the route.
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
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The types of destination network instances to which the routes belong.
	//
	// 	- **VPC**
	//
	// 	- **VBR**
	//
	// 	- **CCN**
	//
	// 	- **VPN**
	//
	// >  The destination route tables take effect only if the routing policy is applied to the egress gateway direction, and the type of the destination route tables is the same as that of the network instance in the current region.
	DestinationChildInstanceTypes *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes `json:"DestinationChildInstanceTypes,omitempty" xml:"DestinationChildInstanceTypes,omitempty" type:"Struct"`
	// The prefix list against which routes are matched.
	//
	// IPv4 and IPv6 addresses are supported.
	DestinationCidrBlocks *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks `json:"DestinationCidrBlocks,omitempty" xml:"DestinationCidrBlocks,omitempty" type:"Struct"`
	// The IDs of the destination network instances to which the routes point.
	//
	// >  The destination route tables take effect only if the routing policy is applied to the egress gateway direction, and the ID the destination instance is the same as that of the network instance in the current region.
	DestinationInstanceIds *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds `json:"DestinationInstanceIds,omitempty" xml:"DestinationInstanceIds,omitempty" type:"Struct"`
	// Indicates whether the destination network instance IDs are excluded.
	//
	// 	- **false*	- (default): A route is a match if its destination network instance ID is in the list specified by **DestinationInstanceIds.N**.
	//
	// 	- **true**: A route is a match if its destination network instance ID is not in the list specified by **DestinationInstanceIds.N**.
	//
	// example:
	//
	// false
	DestinationInstanceIdsReverseMatch *bool `json:"DestinationInstanceIdsReverseMatch,omitempty" xml:"DestinationInstanceIdsReverseMatch,omitempty"`
	// The IDs of the destination regions for the routing policy.
	DestinationRegionIds *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds `json:"DestinationRegionIds,omitempty" xml:"DestinationRegionIds,omitempty" type:"Struct"`
	// The IDs of the destination route tables to which the routes belong. You can enter at most 32 route table IDs.
	//
	// >  The destination route tables take effect only if the routing policy is applied to the egress gateway direction, and the destination route table IDs are in the current region.
	DestinationRouteTableIds *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds `json:"DestinationRouteTableIds,omitempty" xml:"DestinationRouteTableIds,omitempty" type:"Struct"`
	// The action performed on a route that meets the match conditions.
	//
	// 	- **Permit**: the route is permitted.
	//
	// 	- **Deny**: the route is denied.
	//
	// example:
	//
	// Deny
	MapResult *string `json:"MapResult,omitempty" xml:"MapResult,omitempty"`
	// The type of IP address to be matched against the match condition. Valid values:
	//
	// 	- **IPv4**: IPv4 addresses
	//
	// 	- **IPv6**: IPv6 addresses
	//
	// 	- If no value is returned, both IPv4 and IPv6 addresses are matched against the match condition.
	//
	// example:
	//
	// IPv4
	MatchAddressType *string `json:"MatchAddressType,omitempty" xml:"MatchAddressType,omitempty"`
	// The AS paths against which routes are matched.
	MatchAsns *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns `json:"MatchAsns,omitempty" xml:"MatchAsns,omitempty" type:"Struct"`
	// The community set against which routes are matched.
	MatchCommunitySet *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet `json:"MatchCommunitySet,omitempty" xml:"MatchCommunitySet,omitempty" type:"Struct"`
	// The priority of the routing policy that you want to associate with the current one.
	//
	// example:
	//
	// 33
	NextPriority *int32 `json:"NextPriority,omitempty" xml:"NextPriority,omitempty"`
	// The community set on which actions are performed.
	OperateCommunitySet *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet `json:"OperateCommunitySet,omitempty" xml:"OperateCommunitySet,omitempty" type:"Struct"`
	// The new priority of the route.
	//
	// A smaller value indicates a higher priority.
	//
	// This parameter indicates the action to be performed when a route meets the match condition.
	//
	// example:
	//
	// 20
	Preference *int32 `json:"Preference,omitempty" xml:"Preference,omitempty"`
	// The AS paths that are prepended by using an action statement when regional gateways receive or advertise routes.
	//
	// This parameter indicates the action to be performed when a route meets the match condition.
	PrependAsPath *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath `json:"PrependAsPath,omitempty" xml:"PrependAsPath,omitempty" type:"Struct"`
	// The priority of the routing policy. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 5000
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The routing policy ID.
	//
	// example:
	//
	// cenrmap-y40mxdvf7joc12****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
	// The type of route that is compared. Valid values:
	//
	// 	- **System**: system routes that are automatically generated by the system.
	//
	// 	- **Custom**: custom routes that are manually added.
	//
	// 	- **BGP**: routes that are advertised over BGP.
	RouteTypes *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes `json:"RouteTypes,omitempty" xml:"RouteTypes,omitempty" type:"Struct"`
	// The types of source network instances to which the routes belong.
	//
	// 	- **VPC**
	//
	// 	- **VBR**
	//
	// 	- **CCN**
	//
	// 	- **VPN**
	SourceChildInstanceTypes *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes `json:"SourceChildInstanceTypes,omitempty" xml:"SourceChildInstanceTypes,omitempty" type:"Struct"`
	// The IDs of the source network instances to which the routes belong.
	SourceInstanceIds *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds `json:"SourceInstanceIds,omitempty" xml:"SourceInstanceIds,omitempty" type:"Struct"`
	// Indicates whether the source network instance IDs are excluded.
	//
	// 	- **false*	- (default): A route is a match if its source network instance ID is in the list specified by **SourceInstanceIds.N**.
	//
	// 	- **true**: A route is match if its source network instance ID is not in the list specified by **SourceInstanceIds.N**.
	//
	// example:
	//
	// false
	SourceInstanceIdsReverseMatch *bool `json:"SourceInstanceIdsReverseMatch,omitempty" xml:"SourceInstanceIdsReverseMatch,omitempty"`
	// The IDs of the source regions to which the routes belong.
	SourceRegionIds *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds `json:"SourceRegionIds,omitempty" xml:"SourceRegionIds,omitempty" type:"Struct"`
	// The IDs of the source route tables to which the routes belong.
	SourceRouteTableIds *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds `json:"SourceRouteTableIds,omitempty" xml:"SourceRouteTableIds,omitempty" type:"Struct"`
	// The status of the routing policy. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Active**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The route table ID of the transit router with which the routing policy is associated.
	//
	// example:
	//
	// vtb-gw8nx3515m1mbd1z1****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	// The direction in which the routing policy is applied.
	//
	// example:
	//
	// RegionOut
	TransmitDirection *string `json:"TransmitDirection,omitempty" xml:"TransmitDirection,omitempty"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetAsPathMatchMode() *string {
	return s.AsPathMatchMode
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetCenRegionId() *string {
	return s.CenRegionId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetCidrMatchMode() *string {
	return s.CidrMatchMode
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetCommunityMatchMode() *string {
	return s.CommunityMatchMode
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetCommunityOperateMode() *string {
	return s.CommunityOperateMode
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetDestinationChildInstanceTypes() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes {
	return s.DestinationChildInstanceTypes
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetDestinationCidrBlocks() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks {
	return s.DestinationCidrBlocks
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetDestinationInstanceIds() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds {
	return s.DestinationInstanceIds
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetDestinationInstanceIdsReverseMatch() *bool {
	return s.DestinationInstanceIdsReverseMatch
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetDestinationRegionIds() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds {
	return s.DestinationRegionIds
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetDestinationRouteTableIds() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds {
	return s.DestinationRouteTableIds
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetMapResult() *string {
	return s.MapResult
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetMatchAddressType() *string {
	return s.MatchAddressType
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetMatchAsns() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns {
	return s.MatchAsns
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetMatchCommunitySet() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet {
	return s.MatchCommunitySet
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetNextPriority() *int32 {
	return s.NextPriority
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetOperateCommunitySet() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet {
	return s.OperateCommunitySet
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetPreference() *int32 {
	return s.Preference
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetPrependAsPath() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath {
	return s.PrependAsPath
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetRouteTypes() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes {
	return s.RouteTypes
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetSourceChildInstanceTypes() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes {
	return s.SourceChildInstanceTypes
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetSourceInstanceIds() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds {
	return s.SourceInstanceIds
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetSourceInstanceIdsReverseMatch() *bool {
	return s.SourceInstanceIdsReverseMatch
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetSourceRegionIds() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds {
	return s.SourceRegionIds
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetSourceRouteTableIds() *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds {
	return s.SourceRouteTableIds
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) GetTransmitDirection() *string {
	return s.TransmitDirection
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetAsPathMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.AsPathMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCenId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CenId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCenRegionId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CenRegionId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCidrMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CidrMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCommunityMatchMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CommunityMatchMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetCommunityOperateMode(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.CommunityOperateMode = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDescription(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Description = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationChildInstanceTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationChildInstanceTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationCidrBlocks(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationCidrBlocks = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationInstanceIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationInstanceIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationInstanceIdsReverseMatch(v bool) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationInstanceIdsReverseMatch = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationRegionIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationRegionIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetDestinationRouteTableIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.DestinationRouteTableIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMapResult(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MapResult = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMatchAddressType(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MatchAddressType = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMatchAsns(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MatchAsns = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetMatchCommunitySet(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.MatchCommunitySet = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetNextPriority(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.NextPriority = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetOperateCommunitySet(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.OperateCommunitySet = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPreference(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Preference = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPrependAsPath(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.PrependAsPath = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetPriority(v int32) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Priority = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetRouteMapId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetRouteTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.RouteTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceChildInstanceTypes(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceChildInstanceTypes = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceInstanceIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceInstanceIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceInstanceIdsReverseMatch(v bool) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceInstanceIdsReverseMatch = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceRegionIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceRegionIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetSourceRouteTableIds(v *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.SourceRouteTableIds = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetStatus(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.Status = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetTransitRouterRouteTableId(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) SetTransmitDirection(v string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap {
	s.TransmitDirection = &v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMap) Validate() error {
	if s.DestinationChildInstanceTypes != nil {
		if err := s.DestinationChildInstanceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.DestinationCidrBlocks != nil {
		if err := s.DestinationCidrBlocks.Validate(); err != nil {
			return err
		}
	}
	if s.DestinationInstanceIds != nil {
		if err := s.DestinationInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.DestinationRegionIds != nil {
		if err := s.DestinationRegionIds.Validate(); err != nil {
			return err
		}
	}
	if s.DestinationRouteTableIds != nil {
		if err := s.DestinationRouteTableIds.Validate(); err != nil {
			return err
		}
	}
	if s.MatchAsns != nil {
		if err := s.MatchAsns.Validate(); err != nil {
			return err
		}
	}
	if s.MatchCommunitySet != nil {
		if err := s.MatchCommunitySet.Validate(); err != nil {
			return err
		}
	}
	if s.OperateCommunitySet != nil {
		if err := s.OperateCommunitySet.Validate(); err != nil {
			return err
		}
	}
	if s.PrependAsPath != nil {
		if err := s.PrependAsPath.Validate(); err != nil {
			return err
		}
	}
	if s.RouteTypes != nil {
		if err := s.RouteTypes.Validate(); err != nil {
			return err
		}
	}
	if s.SourceChildInstanceTypes != nil {
		if err := s.SourceChildInstanceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.SourceInstanceIds != nil {
		if err := s.SourceInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRegionIds != nil {
		if err := s.SourceRegionIds.Validate(); err != nil {
			return err
		}
	}
	if s.SourceRouteTableIds != nil {
		if err := s.SourceRouteTableIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes struct {
	DestinationChildInstanceType []*string `json:"DestinationChildInstanceType,omitempty" xml:"DestinationChildInstanceType,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) GetDestinationChildInstanceType() []*string {
	return s.DestinationChildInstanceType
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) SetDestinationChildInstanceType(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes {
	s.DestinationChildInstanceType = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationChildInstanceTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks struct {
	DestinationCidrBlock []*string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) GetDestinationCidrBlock() []*string {
	return s.DestinationCidrBlock
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) SetDestinationCidrBlock(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks {
	s.DestinationCidrBlock = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationCidrBlocks) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds struct {
	DestinationInstanceId []*string `json:"DestinationInstanceId,omitempty" xml:"DestinationInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) GetDestinationInstanceId() []*string {
	return s.DestinationInstanceId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) SetDestinationInstanceId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds {
	s.DestinationInstanceId = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds struct {
	DestinationRegionId []*string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds) GetDestinationRegionId() []*string {
	return s.DestinationRegionId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds) SetDestinationRegionId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds {
	s.DestinationRegionId = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRegionIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds struct {
	DestinationRouteTableId []*string `json:"DestinationRouteTableId,omitempty" xml:"DestinationRouteTableId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) GetDestinationRouteTableId() []*string {
	return s.DestinationRouteTableId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) SetDestinationRouteTableId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds {
	s.DestinationRouteTableId = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapDestinationRouteTableIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns struct {
	MatchAsn []*string `json:"MatchAsn,omitempty" xml:"MatchAsn,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) GetMatchAsn() []*string {
	return s.MatchAsn
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) SetMatchAsn(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns {
	s.MatchAsn = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchAsns) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet struct {
	MatchCommunity []*string `json:"MatchCommunity,omitempty" xml:"MatchCommunity,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) GetMatchCommunity() []*string {
	return s.MatchCommunity
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) SetMatchCommunity(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet {
	s.MatchCommunity = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapMatchCommunitySet) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet struct {
	OperateCommunity []*string `json:"OperateCommunity,omitempty" xml:"OperateCommunity,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) GetOperateCommunity() []*string {
	return s.OperateCommunity
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) SetOperateCommunity(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet {
	s.OperateCommunity = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapOperateCommunitySet) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath struct {
	AsPath []*string `json:"AsPath,omitempty" xml:"AsPath,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) GetAsPath() []*string {
	return s.AsPath
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) SetAsPath(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath {
	s.AsPath = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapPrependAsPath) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes struct {
	RouteType []*string `json:"RouteType,omitempty" xml:"RouteType,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) GetRouteType() []*string {
	return s.RouteType
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) SetRouteType(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes {
	s.RouteType = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapRouteTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes struct {
	SourceChildInstanceType []*string `json:"SourceChildInstanceType,omitempty" xml:"SourceChildInstanceType,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) GetSourceChildInstanceType() []*string {
	return s.SourceChildInstanceType
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) SetSourceChildInstanceType(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes {
	s.SourceChildInstanceType = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceChildInstanceTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds struct {
	SourceInstanceId []*string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) GetSourceInstanceId() []*string {
	return s.SourceInstanceId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) SetSourceInstanceId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds {
	s.SourceInstanceId = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds struct {
	SourceRegionId []*string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) GetSourceRegionId() []*string {
	return s.SourceRegionId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) SetSourceRegionId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds {
	s.SourceRegionId = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRegionIds) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds struct {
	SourceRouteTableId []*string `json:"SourceRouteTableId,omitempty" xml:"SourceRouteTableId,omitempty" type:"Repeated"`
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) GoString() string {
	return s.String()
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) GetSourceRouteTableId() []*string {
	return s.SourceRouteTableId
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) SetSourceRouteTableId(v []*string) *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds {
	s.SourceRouteTableId = v
	return s
}

func (s *DescribeCenRouteMapsResponseBodyRouteMapsRouteMapSourceRouteTableIds) Validate() error {
	return dara.Validate(s)
}
