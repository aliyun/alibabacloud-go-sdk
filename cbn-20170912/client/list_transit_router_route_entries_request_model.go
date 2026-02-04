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
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// fce19****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The prefix list ID.
	//
	// example:
	//
	// pl-6ehtn5kqxgeyy08fi****
	PrefixListId         *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The filter conditions.
	RouteFilter []*ListTransitRouterRouteEntriesRequestRouteFilter `json:"RouteFilter,omitempty" xml:"RouteFilter,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The destination CIDR block of the route. **This parameter is to be deprecated. We recommend that you use the RouteFilter parameter**.
	//
	// example:
	//
	// 192.168.0.0/24
	TransitRouterRouteEntryDestinationCidrBlock *string `json:"TransitRouterRouteEntryDestinationCidrBlock,omitempty" xml:"TransitRouterRouteEntryDestinationCidrBlock,omitempty"`
	// The route ID.
	//
	// example:
	//
	// rte-oklkgwmj97z6dn****
	TransitRouterRouteEntryIds []*string `json:"TransitRouterRouteEntryIds,omitempty" xml:"TransitRouterRouteEntryIds,omitempty" type:"Repeated"`
	// The route name.
	//
	// example:
	//
	// testname
	TransitRouterRouteEntryNames []*string `json:"TransitRouterRouteEntryNames,omitempty" xml:"TransitRouterRouteEntryNames,omitempty" type:"Repeated"`
	// The ID of the network instance connection that you want to specify as the next hop.
	//
	// example:
	//
	// tr-attach-nls9fzkfat8934****
	TransitRouterRouteEntryNextHopId *string `json:"TransitRouterRouteEntryNextHopId,omitempty" xml:"TransitRouterRouteEntryNextHopId,omitempty"`
	// The next hop ID.
	//
	// example:
	//
	// vpc-m5ent6du8deaq5*****
	TransitRouterRouteEntryNextHopResourceId *string `json:"TransitRouterRouteEntryNextHopResourceId,omitempty" xml:"TransitRouterRouteEntryNextHopResourceId,omitempty"`
	// The next hop type. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VBR**
	//
	// 	- **TR**
	//
	// 	- **VPN**
	//
	// example:
	//
	// VPC
	TransitRouterRouteEntryNextHopResourceType *string `json:"TransitRouterRouteEntryNextHopResourceType,omitempty" xml:"TransitRouterRouteEntryNextHopResourceType,omitempty"`
	// The next hop type. Valid values:
	//
	// 	- **BlackHole**: routes network traffic to a black hole.
	//
	// 	- **Attachment**: routes network traffic to a network instance connection.
	//
	// example:
	//
	// Attachment
	TransitRouterRouteEntryNextHopType *string `json:"TransitRouterRouteEntryNextHopType,omitempty" xml:"TransitRouterRouteEntryNextHopType,omitempty"`
	// The source instance ID.
	//
	// example:
	//
	// vpc-m5ent6du8deaq5*****
	TransitRouterRouteEntryOriginResourceId *string `json:"TransitRouterRouteEntryOriginResourceId,omitempty" xml:"TransitRouterRouteEntryOriginResourceId,omitempty"`
	// The source instance type. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **VBR**
	//
	// 	- **TR**
	//
	// 	- **VPN**
	//
	// example:
	//
	// VPC
	TransitRouterRouteEntryOriginResourceType *string `json:"TransitRouterRouteEntryOriginResourceType,omitempty" xml:"TransitRouterRouteEntryOriginResourceType,omitempty"`
	// The status of the route. Valid values:
	//
	// 	- **All**
	//
	// 	- **Active*	- (default)
	//
	// 	- **Rejected**
	//
	// 	- **Prohibited**
	//
	// 	- **Standby**
	//
	// 	- **Candidate**
	//
	// If you do not specify a value, routes in the active state are queried.
	//
	// example:
	//
	// Active
	TransitRouterRouteEntryStatus *string `json:"TransitRouterRouteEntryStatus,omitempty" xml:"TransitRouterRouteEntryStatus,omitempty"`
	// The route type. Valid values:
	//
	// 	- **Propagated**: automatically learned by the route table.
	//
	// 	- **Static**: static routes.
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
	// The match pattern for filtering CIDR blocks. Valid values:
	//
	// 	- **PrefixExactMatchCidrs**: exact matching.
	//
	// 	- **LongestPrefixMatchCidrs**: longest prefix matching. You can specify IP addresses and CIDR blocks.
	//
	// 	- **SubnetOfMatchCidrs**: subnet matching. The subnets of the specified CIDR blocks, including the CIDR block, are matches against the match conditions.
	//
	// 	- **SupernetOfMatchCidrs**: supernet matching. The supernets of the CIDR block, including the CIDR block, are matched against the match conditions.
	//
	// By default, the logical operator among filter conditions is **AND**. Information about a route entry is returned only if the route entry matches all filter conditions. Filter conditions must be unique.
	//
	// example:
	//
	// PrefixExactMatchCidrs
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value.
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
