// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouteTableAggregationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeTransitRouteTableAggregationResponseBody
	GetCount() *int32
	SetData(v []*DescribeTransitRouteTableAggregationResponseBodyData) *DescribeTransitRouteTableAggregationResponseBody
	GetData() []*DescribeTransitRouteTableAggregationResponseBodyData
	SetNextToken(v string) *DescribeTransitRouteTableAggregationResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTransitRouteTableAggregationResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeTransitRouteTableAggregationResponseBody
	GetTotal() *int32
}

type DescribeTransitRouteTableAggregationResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// A list of aggregate routes.
	Data []*DescribeTransitRouteTableAggregationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeTransitRouteTableAggregationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouteTableAggregationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouteTableAggregationResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeTransitRouteTableAggregationResponseBody) GetData() []*DescribeTransitRouteTableAggregationResponseBodyData {
	return s.Data
}

func (s *DescribeTransitRouteTableAggregationResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTransitRouteTableAggregationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTransitRouteTableAggregationResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeTransitRouteTableAggregationResponseBody) SetCount(v int32) *DescribeTransitRouteTableAggregationResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBody) SetData(v []*DescribeTransitRouteTableAggregationResponseBodyData) *DescribeTransitRouteTableAggregationResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBody) SetNextToken(v string) *DescribeTransitRouteTableAggregationResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBody) SetRequestId(v string) *DescribeTransitRouteTableAggregationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBody) SetTotal(v int32) *DescribeTransitRouteTableAggregationResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTransitRouteTableAggregationResponseBodyData struct {
	// The description of the aggregate route.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the aggregate route.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the aggregate route.
	//
	// The valid value is **Static**, which indicates a static route. By default, aggregate routes advertised to a VPC are considered custom routes.
	//
	// example:
	//
	// Static
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The scope of networks that you want to advertise the aggregate route.
	//
	// The valid value is **VPC**, which indicates that the aggregate route is advertised to all virtual private clouds (VPCs) that are in associated forwarding correlation with the Enterprise Edition transit router and have route synchronization enabled.
	//
	// example:
	//
	// VPC
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The list of propagation ranges of the aggregation route.
	//
	// >  You must specify at least one of the following attributes: Aggregation Scope and Aggregate Scope List. We recommend that you specify the latter. The elements in the two attributes cannot be duplicate.
	ScopeList []*string `json:"ScopeList,omitempty" xml:"ScopeList,omitempty" type:"Repeated"`
	// The status of the advertisement of the aggregate route. Valid values:
	//
	// 	- **AllConfigured**: The aggregate route is advertised to all VPCs.
	//
	// 	- **Configuring**: The aggregate route is being advertised.
	//
	// 	- **ConfigFailed**: The aggregate route failed to be advertised.
	//
	// 	- **PartialConfigured**: Failed to advertise the aggregate route to some VPCs.
	//
	// 	- **Deleting**: The aggregate route is being deleted.
	//
	// example:
	//
	// AllConfigured
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the route table of the Enterprise Edition transit router.
	//
	// example:
	//
	// vtb-6ehgc262hr170qgyc****
	TrRouteTableId *string `json:"TrRouteTableId,omitempty" xml:"TrRouteTableId,omitempty"`
	// The destination CIDR block of the aggregate route.
	//
	// example:
	//
	// 192.168.10.0/24
	TransitRouteTableAggregationCidr *string `json:"TransitRouteTableAggregationCidr,omitempty" xml:"TransitRouteTableAggregationCidr,omitempty"`
}

func (s DescribeTransitRouteTableAggregationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouteTableAggregationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) GetRouteType() *string {
	return s.RouteType
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) GetScopeList() []*string {
	return s.ScopeList
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) GetTrRouteTableId() *string {
	return s.TrRouteTableId
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) GetTransitRouteTableAggregationCidr() *string {
	return s.TransitRouteTableAggregationCidr
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) SetDescription(v string) *DescribeTransitRouteTableAggregationResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) SetName(v string) *DescribeTransitRouteTableAggregationResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) SetRouteType(v string) *DescribeTransitRouteTableAggregationResponseBodyData {
	s.RouteType = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) SetScope(v string) *DescribeTransitRouteTableAggregationResponseBodyData {
	s.Scope = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) SetScopeList(v []*string) *DescribeTransitRouteTableAggregationResponseBodyData {
	s.ScopeList = v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) SetStatus(v string) *DescribeTransitRouteTableAggregationResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) SetTrRouteTableId(v string) *DescribeTransitRouteTableAggregationResponseBodyData {
	s.TrRouteTableId = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) SetTransitRouteTableAggregationCidr(v string) *DescribeTransitRouteTableAggregationResponseBodyData {
	s.TransitRouteTableAggregationCidr = &v
	return s
}

func (s *DescribeTransitRouteTableAggregationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
