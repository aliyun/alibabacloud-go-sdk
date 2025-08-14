// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTransitRouterRouteTablesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterRouteTablesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTransitRouterRouteTablesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTransitRouterRouteTablesResponseBody
	GetTotalCount() *int32
	SetTransitRouterRouteTables(v []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) *ListTransitRouterRouteTablesResponseBody
	GetTransitRouterRouteTables() []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables
}

type ListTransitRouterRouteTablesResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. Valid values:
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// 	- If a value of **NextToken*	- is not returned, it indicates that no additional results exist.
	//
	// example:
	//
	// dd20****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 82678F4A-C9F7-4CC1-8BF0-D619A63BFC57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// A list of route tables.
	TransitRouterRouteTables []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables `json:"TransitRouterRouteTables,omitempty" xml:"TransitRouterRouteTables,omitempty" type:"Repeated"`
}

func (s ListTransitRouterRouteTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterRouteTablesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterRouteTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTransitRouterRouteTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTransitRouterRouteTablesResponseBody) GetTransitRouterRouteTables() []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	return s.TransitRouterRouteTables
}

func (s *ListTransitRouterRouteTablesResponseBody) SetMaxResults(v int32) *ListTransitRouterRouteTablesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) SetNextToken(v string) *ListTransitRouterRouteTablesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) SetRequestId(v string) *ListTransitRouterRouteTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) SetTotalCount(v int32) *ListTransitRouterRouteTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) SetTransitRouterRouteTables(v []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) *ListTransitRouterRouteTablesResponseBody {
	s.TransitRouterRouteTables = v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables struct {
	// The time when the route table was created.
	//
	// The time follows the ISO8601 standard in the YYYY-MM-DDThh:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-15T09:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The region ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The features of the route table.
	RouteTableOptions *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions `json:"RouteTableOptions,omitempty" xml:"RouteTableOptions,omitempty" type:"Struct"`
	// The tags.
	Tags []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The transit router ID.
	//
	// example:
	//
	// tr-8vb8bie2koduo5awz****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The description of the route table.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-bp1dudbh2d5na6b50****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
	// The name of the route table.
	//
	// example:
	//
	// testname
	TransitRouterRouteTableName *string `json:"TransitRouterRouteTableName,omitempty" xml:"TransitRouterRouteTableName,omitempty"`
	// The status of the route table. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Deleting**
	//
	// 	- **Active**
	//
	// example:
	//
	// Active
	TransitRouterRouteTableStatus *string `json:"TransitRouterRouteTableStatus,omitempty" xml:"TransitRouterRouteTableStatus,omitempty"`
	// The type of the route table. Valid values:
	//
	// 	- **Custom**
	//
	// 	- **System**
	//
	// example:
	//
	// System
	TransitRouterRouteTableType *string `json:"TransitRouterRouteTableType,omitempty" xml:"TransitRouterRouteTableType,omitempty"`
}

func (s ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetRouteTableOptions() *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions {
	return s.RouteTableOptions
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetTags() []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags {
	return s.Tags
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetTransitRouterRouteTableDescription() *string {
	return s.TransitRouterRouteTableDescription
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetTransitRouterRouteTableName() *string {
	return s.TransitRouterRouteTableName
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetTransitRouterRouteTableStatus() *string {
	return s.TransitRouterRouteTableStatus
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) GetTransitRouterRouteTableType() *string {
	return s.TransitRouterRouteTableType
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetCreateTime(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.CreateTime = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetRegionId(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.RegionId = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetRouteTableOptions(v *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.RouteTableOptions = v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTags(v []*ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.Tags = v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterId(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterId = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableDescription(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableDescription = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableId(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableName(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableName = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableStatus(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableStatus = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) SetTransitRouterRouteTableType(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables {
	s.TransitRouterRouteTableType = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTables) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions struct {
	// Indicates whether ECMP routing is enabled. Valid values:
	//
	// 	- **disable*	- If ECMP routing is disabled, routes that are learned from different regions but have the same prefix and attributes select the transit router with the smallest region ID as the next hop. Region IDs are sorted in alphabetic order. The network latency and bandwidth consumption also vary based on the region. Proceed with caution.
	//
	// 	- **enable*	- If ECMP routing is enabled, routes that are learned from different regions but have the same prefix and attributes form an ECMP route. The network latency and bandwidth consumption also vary based on the region. Proceed with caution.
	//
	// example:
	//
	// disable
	MultiRegionECMP *string `json:"MultiRegionECMP,omitempty" xml:"MultiRegionECMP,omitempty"`
}

func (s ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions) GetMultiRegionECMP() *string {
	return s.MultiRegionECMP
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions) SetMultiRegionECMP(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions {
	s.MultiRegionECMP = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesRouteTableOptions) Validate() error {
	return dara.Validate(s)
}

type ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags) GetKey() *string {
	return s.Key
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags) GetValue() *string {
	return s.Value
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags) SetKey(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags {
	s.Key = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags) SetValue(v string) *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags {
	s.Value = &v
	return s
}

func (s *ListTransitRouterRouteTablesResponseBodyTransitRouterRouteTablesTags) Validate() error {
	return dara.Validate(s)
}
