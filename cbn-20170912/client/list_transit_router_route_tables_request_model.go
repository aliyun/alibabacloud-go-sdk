// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterRouteTablesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterRouteTablesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterRouteTablesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterRouteTablesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListTransitRouterRouteTablesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterRouteTablesRequest
	GetResourceOwnerId() *int64
	SetRouteTableOptions(v *ListTransitRouterRouteTablesRequestRouteTableOptions) *ListTransitRouterRouteTablesRequest
	GetRouteTableOptions() *ListTransitRouterRouteTablesRequestRouteTableOptions
	SetTag(v []*ListTransitRouterRouteTablesRequestTag) *ListTransitRouterRouteTablesRequest
	GetTag() []*ListTransitRouterRouteTablesRequestTag
	SetTransitRouterId(v string) *ListTransitRouterRouteTablesRequest
	GetTransitRouterId() *string
	SetTransitRouterRouteTableIds(v []*string) *ListTransitRouterRouteTablesRequest
	GetTransitRouterRouteTableIds() []*string
	SetTransitRouterRouteTableNames(v []*string) *ListTransitRouterRouteTablesRequest
	GetTransitRouterRouteTableNames() []*string
	SetTransitRouterRouteTableStatus(v string) *ListTransitRouterRouteTablesRequest
	GetTransitRouterRouteTableStatus() *string
	SetTransitRouterRouteTableType(v string) *ListTransitRouterRouteTablesRequest
	GetTransitRouterRouteTableType() *string
}

type ListTransitRouterRouteTablesRequest struct {
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query or no subsequent query is to be sent, ignore this parameter.
	//
	// 	- If a subsequent query is to be sent, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// dd20****
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The features of the route table.
	RouteTableOptions *ListTransitRouterRouteTablesRequestRouteTableOptions `json:"RouteTableOptions,omitempty" xml:"RouteTableOptions,omitempty" type:"Struct"`
	// The information about the tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*ListTransitRouterRouteTablesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-uf654ttymmljlvh2x****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the route table.
	//
	// You can query multiple route tables in each call. Maximum value of **N**: **20**.
	//
	// example:
	//
	// vtb-bp1l8awdb4iuo9uwu****
	TransitRouterRouteTableIds []*string `json:"TransitRouterRouteTableIds,omitempty" xml:"TransitRouterRouteTableIds,omitempty" type:"Repeated"`
	// The name of the route table.
	//
	// You can query multiple route tables in each call. Maximum value of **N**: **20**.
	//
	// > If you set both **TransitRouterRouteTableNames.N*	- and **TransitRouterRouteTableIds.N**, make sure that the specified name and ID belong to the same route table.
	//
	// example:
	//
	// testname
	TransitRouterRouteTableNames []*string `json:"TransitRouterRouteTableNames,omitempty" xml:"TransitRouterRouteTableNames,omitempty" type:"Repeated"`
	// The status of the route table. Valid values:
	//
	// 	- **Creating**: The route table is being created.
	//
	// 	- **Deleting**: The route table is being deleted.
	//
	// 	- **Active**: The route table is available.
	//
	// example:
	//
	// Active
	TransitRouterRouteTableStatus *string `json:"TransitRouterRouteTableStatus,omitempty" xml:"TransitRouterRouteTableStatus,omitempty"`
	// The type of the route table. Valid values:
	//
	// 	- **Custom**: a custom route table
	//
	// 	- **System**: the default route table
	//
	// example:
	//
	// Custom
	TransitRouterRouteTableType *string `json:"TransitRouterRouteTableType,omitempty" xml:"TransitRouterRouteTableType,omitempty"`
}

func (s ListTransitRouterRouteTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterRouteTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterRouteTablesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterRouteTablesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterRouteTablesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterRouteTablesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterRouteTablesRequest) GetRouteTableOptions() *ListTransitRouterRouteTablesRequestRouteTableOptions {
	return s.RouteTableOptions
}

func (s *ListTransitRouterRouteTablesRequest) GetTag() []*ListTransitRouterRouteTablesRequestTag {
	return s.Tag
}

func (s *ListTransitRouterRouteTablesRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterRouteTablesRequest) GetTransitRouterRouteTableIds() []*string {
	return s.TransitRouterRouteTableIds
}

func (s *ListTransitRouterRouteTablesRequest) GetTransitRouterRouteTableNames() []*string {
	return s.TransitRouterRouteTableNames
}

func (s *ListTransitRouterRouteTablesRequest) GetTransitRouterRouteTableStatus() *string {
	return s.TransitRouterRouteTableStatus
}

func (s *ListTransitRouterRouteTablesRequest) GetTransitRouterRouteTableType() *string {
	return s.TransitRouterRouteTableType
}

func (s *ListTransitRouterRouteTablesRequest) SetMaxResults(v int32) *ListTransitRouterRouteTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetNextToken(v string) *ListTransitRouterRouteTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetOwnerAccount(v string) *ListTransitRouterRouteTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetOwnerId(v int64) *ListTransitRouterRouteTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetResourceOwnerAccount(v string) *ListTransitRouterRouteTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetResourceOwnerId(v int64) *ListTransitRouterRouteTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetRouteTableOptions(v *ListTransitRouterRouteTablesRequestRouteTableOptions) *ListTransitRouterRouteTablesRequest {
	s.RouteTableOptions = v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTag(v []*ListTransitRouterRouteTablesRequestTag) *ListTransitRouterRouteTablesRequest {
	s.Tag = v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterId(v string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterRouteTableIds(v []*string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterRouteTableIds = v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterRouteTableNames(v []*string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterRouteTableNames = v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterRouteTableStatus(v string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterRouteTableStatus = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) SetTransitRouterRouteTableType(v string) *ListTransitRouterRouteTablesRequest {
	s.TransitRouterRouteTableType = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequest) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterRouteTablesRequestRouteTableOptions struct {
	// Specifies whether to enable equal-cost multi-path (ECMP) routing. Valid values:
	//
	// 	- **disable**: disables ECMP routing If you disable ECMP routing, routes that are learned from different regions but have the same prefix and attributes select the transit router with the smallest region ID as the next hop. Region IDs are sorted in alphabetic order. The network latency and bandwidth consumption also vary based on the region. Proceed with caution.
	//
	// 	- **enable**: enables ECMP routing. If you enable ECMP routing, routes that are learned from different regions but have the same prefix and attributes form an ECMP route. The network latency and bandwidth consumption also vary based on the region. Proceed with caution.
	//
	// example:
	//
	// disable
	MultiRegionECMP *string `json:"MultiRegionECMP,omitempty" xml:"MultiRegionECMP,omitempty"`
}

func (s ListTransitRouterRouteTablesRequestRouteTableOptions) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablesRequestRouteTableOptions) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesRequestRouteTableOptions) GetMultiRegionECMP() *string {
	return s.MultiRegionECMP
}

func (s *ListTransitRouterRouteTablesRequestRouteTableOptions) SetMultiRegionECMP(v string) *ListTransitRouterRouteTablesRequestRouteTableOptions {
	s.MultiRegionECMP = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequestRouteTableOptions) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterRouteTablesRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterRouteTablesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterRouteTablesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterRouteTablesRequestTag) SetKey(v string) *ListTransitRouterRouteTablesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequestTag) SetValue(v string) *ListTransitRouterRouteTablesRequestTag {
	s.Value = &v
	return s
}

func (s *ListTransitRouterRouteTablesRequestTag) Validate() error {
	return dara.Validate(s)
}
