// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterRouteEntriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterRouteEntriesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterRouteEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterRouteEntriesRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *ListTransitRouterRouteEntriesRequest
	GetPrefixListId() *string
	SetResourceOwnerAccount(v string) *ListTransitRouterRouteEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterRouteEntriesRequest
	GetResourceOwnerId() *int64
	SetRouteFilter(v []*ListTransitRouterRouteEntriesRequestRouteFilter) *ListTransitRouterRouteEntriesRequest
	GetRouteFilter() []*ListTransitRouterRouteEntriesRequestRouteFilter
	SetTransitRouterRouteEntryDestinationCidrBlock(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryDestinationCidrBlock() *string
	SetTransitRouterRouteEntryIds(v []*string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryIds() []*string
	SetTransitRouterRouteEntryNames(v []*string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryNames() []*string
	SetTransitRouterRouteEntryNextHopId(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryNextHopId() *string
	SetTransitRouterRouteEntryNextHopResourceId(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryNextHopResourceId() *string
	SetTransitRouterRouteEntryNextHopResourceType(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryNextHopResourceType() *string
	SetTransitRouterRouteEntryNextHopType(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryNextHopType() *string
	SetTransitRouterRouteEntryOriginResourceId(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryOriginResourceId() *string
	SetTransitRouterRouteEntryOriginResourceType(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryOriginResourceType() *string
	SetTransitRouterRouteEntryStatus(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryStatus() *string
	SetTransitRouterRouteEntryType(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteEntryType() *string
	SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteEntriesRequest
	GetTransitRouterRouteTableId() *string
}

type ListTransitRouterRouteEntriesRequest struct {
	// The number of entries to return per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If this is your first query or no next query is to be sent, ignore this parameter.
	//
	// - If a next query is to be sent, set the value to the **NextToken*	- value returned from the last API call.
	//
	// example:
	//
	// fce19****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-6ehtn5kqxgeyy08fi****
	PrefixListId         *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The filter conditions for the CIDR blocks of the route entry in the routing table.
	RouteFilter []*ListTransitRouterRouteEntriesRequestRouteFilter `json:"RouteFilter,omitempty" xml:"RouteFilter,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The destination CIDR block of the route entry (**This parameter is deprecated. Use the RouteFilter parameter instead**).
	//
	// example:
	//
	// 192.168.0.0/24
	TransitRouterRouteEntryDestinationCidrBlock *string `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	// The IDs of the route entry in the routing table.
	//
	// example:
	//
	// rte-oklkgwmj97z6dn****
	TransitRouterRouteEntryIds []*string `json:"TransitRouterRouteEntryIds,omitempty" xml:"TransitRouterRouteEntryIds,omitempty" type:"Repeated"`
	// The names of the route entry in the routing table.
	//
	// example:
	//
	// testname
	TransitRouterRouteEntryNames []*string `json:"TransitRouterRouteEntryNames,omitempty" xml:"TransitRouterRouteEntryNames,omitempty" type:"Repeated"`
	// The ID of the network instance connection associated with the next hop of the route entry.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterRouteEntryNextHopId *string `json:"TransitRouterRouteEntryNextHopId,omitempty" xml:"TransitRouterRouteEntryNextHopId,omitempty"`
	// The next hop instance ID of the route entry in the routing table.
	//
	// example:
	//
	// vpc-m5ent6du8deaq5*****
	TransitRouterRouteEntryNextHopResourceId *string `json:"TransitRouterRouteEntryNextHopResourceId,omitempty" xml:"TransitRouterRouteEntryNextHopResourceId,omitempty"`
	// The next hop instance type of the route entry. Valid values:
	//
	// - **VPC**: VPC.
	//
	// - **VBR**: VBR.
	//
	// - **TR**: transit router.
	//
	// - **VPN**: IPsec connection.
	//
	// example:
	//
	// VPC
	TransitRouterRouteEntryNextHopResourceType *string `json:"TransitRouterRouteEntryNextHopResourceType,omitempty" xml:"TransitRouterRouteEntryNextHopResourceType,omitempty"`
	// The next hop type. Valid values:
	//
	// - **BlackHole**: specifies that the route entry is a blackhole route.
	//
	// - **Attachment**: specifies that the next hop of the route entry is a network instance connection.
	//
	// example:
	//
	// Attachment
	TransitRouterRouteEntryNextHopType *string `json:"TransitRouterRouteEntryNextHopType,omitempty" xml:"TransitRouterRouteEntryNextHopType,omitempty"`
	// The origin instance ID of the route entry in the routing table.
	//
	// example:
	//
	// vpc-m5ent6du8deaq5*****
	TransitRouterRouteEntryOriginResourceId *string `json:"TransitRouterRouteEntryOriginResourceId,omitempty" xml:"TransitRouterRouteEntryOriginResourceId,omitempty"`
	// The origin instance type of the route entry. Valid values:
	//
	// - **VPC**: VPC.
	//
	// - **VBR**: VBR.
	//
	// - **TR**: transit router.
	//
	// - **VPN**: IPsec connection.
	//
	// example:
	//
	// VPC
	TransitRouterRouteEntryOriginResourceType *string `json:"TransitRouterRouteEntryOriginResourceType,omitempty" xml:"TransitRouterRouteEntryOriginResourceType,omitempty"`
	// The status of the route entry. Valid values:
	//
	// - **All**: queries route entries in all states.
	//
	// - **Active*	- (default): queries only active route entries.
	//
	// - **Rejected**: queries only route entries that are rejected due to route conflicts.
	//
	// - **Prohibited**: queries only route entries that are prohibited by route maps.
	//
	// - **Standby**: queries only standby route entries.
	//
	// - **Candidate**: queries only candidate route entries.
	//
	// If you do not specify this parameter, the system queries only active route entries by default.
	//
	// example:
	//
	// Active
	TransitRouterRouteEntryStatus *string `json:"TransitRouterRouteEntryStatus,omitempty" xml:"TransitRouterRouteEntryStatus,omitempty"`
	// The destination type of the route entry. Valid values:
	//
	// - **Propagated**: generated by automatic learning of the current routing table.
	//
	// - **Static**: statically configured in the current routing table.
	//
	// example:
	//
	// Propagated
	TransitRouterRouteEntryType *string `json:"TransitRouterRouteEntryType,omitempty" xml:"TransitRouterRouteEntryType,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s ListTransitRouterRouteEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteEntriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterRouteEntriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterRouteEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterRouteEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterRouteEntriesRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *ListTransitRouterRouteEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterRouteEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterRouteEntriesRequest) GetRouteFilter() []*ListTransitRouterRouteEntriesRequestRouteFilter {
	return s.RouteFilter
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryDestinationCidrBlock() *string {
	return s.TransitRouterRouteEntryDestinationCidrBlock
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryIds() []*string {
	return s.TransitRouterRouteEntryIds
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryNames() []*string {
	return s.TransitRouterRouteEntryNames
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryNextHopId() *string {
	return s.TransitRouterRouteEntryNextHopId
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryNextHopResourceId() *string {
	return s.TransitRouterRouteEntryNextHopResourceId
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryNextHopResourceType() *string {
	return s.TransitRouterRouteEntryNextHopResourceType
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryNextHopType() *string {
	return s.TransitRouterRouteEntryNextHopType
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryOriginResourceId() *string {
	return s.TransitRouterRouteEntryOriginResourceId
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryOriginResourceType() *string {
	return s.TransitRouterRouteEntryOriginResourceType
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryStatus() *string {
	return s.TransitRouterRouteEntryStatus
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteEntryType() *string {
	return s.TransitRouterRouteEntryType
}

func (s *ListTransitRouterRouteEntriesRequest) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *ListTransitRouterRouteEntriesRequest) SetMaxResults(v int32) *ListTransitRouterRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetNextToken(v string) *ListTransitRouterRouteEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetOwnerAccount(v string) *ListTransitRouterRouteEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetOwnerId(v int64) *ListTransitRouterRouteEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetPrefixListId(v string) *ListTransitRouterRouteEntriesRequest {
	s.PrefixListId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetResourceOwnerAccount(v string) *ListTransitRouterRouteEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetResourceOwnerId(v int64) *ListTransitRouterRouteEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetRouteFilter(v []*ListTransitRouterRouteEntriesRequestRouteFilter) *ListTransitRouterRouteEntriesRequest {
	s.RouteFilter = v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryDestinationCidrBlock(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryDestinationCidrBlock = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryIds(v []*string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryIds = v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryNames(v []*string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryNames = v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryNextHopId(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryNextHopId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryNextHopResourceId(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryNextHopResourceId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryNextHopResourceType(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryNextHopResourceType = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryNextHopType(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryNextHopType = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryOriginResourceId(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryOriginResourceId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryOriginResourceType(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryOriginResourceType = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryStatus(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryStatus = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteEntryType(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteEntryType = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteEntriesRequest {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequest) Validate() error {
	if s.RouteFilter != nil {
		for _, item := range s.RouteFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTransitRouterRouteEntriesRequestRouteFilter struct {
	// The filter condition. Valid values:
	//
	// - **PrefixExactMatchCidrs**: exact match.
	//
	// - **LongestPrefixMatchCidrs**: longest prefix match. Both IP addresses and CIDR blocks are supported.
	//
	// - **SubnetOfMatchCidrs**: subnet match. Matches the subnets of the specified CIDR block, including the specified CIDR block.
	//
	// - **SupernetOfMatchCidrs**: supernet match. Matches the supernets of the specified CIDR block, including the specified CIDR block.
	//
	// Multiple filter conditions are connected by the **AND*	- logical operator by default. A route entry must meet all filter conditions to be considered a match. You cannot specify duplicate filter conditions.
	//
	// example:
	//
	// PrefixExactMatchCidrs
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of filter condition values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListTransitRouterRouteEntriesRequestRouteFilter) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteEntriesRequestRouteFilter) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteEntriesRequestRouteFilter) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterRouteEntriesRequestRouteFilter) GetValue() []*string {
	return s.Value
}

func (s *ListTransitRouterRouteEntriesRequestRouteFilter) SetKey(v string) *ListTransitRouterRouteEntriesRequestRouteFilter {
	s.Key = &v
	return s
}

func (s *ListTransitRouterRouteEntriesRequestRouteFilter) SetValue(v []*string) *ListTransitRouterRouteEntriesRequestRouteFilter {
	s.Value = v
	return s
}

func (s *ListTransitRouterRouteEntriesRequestRouteFilter) Validate() error {
	return dara.Validate(s)
}
