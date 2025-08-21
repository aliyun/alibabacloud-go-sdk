// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRouteEntryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *DescribeEnsRouteEntryListRequest
	GetDestinationCidrBlock() *string
	SetNextHopId(v string) *DescribeEnsRouteEntryListRequest
	GetNextHopId() *string
	SetNextHopType(v string) *DescribeEnsRouteEntryListRequest
	GetNextHopType() *string
	SetPageNumber(v int32) *DescribeEnsRouteEntryListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnsRouteEntryListRequest
	GetPageSize() *int32
	SetRouteEntryId(v string) *DescribeEnsRouteEntryListRequest
	GetRouteEntryId() *string
	SetRouteEntryName(v string) *DescribeEnsRouteEntryListRequest
	GetRouteEntryName() *string
	SetRouteEntryType(v string) *DescribeEnsRouteEntryListRequest
	GetRouteEntryType() *string
	SetRouteTableId(v string) *DescribeEnsRouteEntryListRequest
	GetRouteTableId() *string
}

type DescribeEnsRouteEntryListRequest struct {
	// The destination Classless Inter-Domain Routing (CIDR) block of the route entry.
	//
	// example:
	//
	// 11.0.0.0/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the next hop.
	//
	// example:
	//
	// i-2zecshuv3axtr2gc4noa
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of next hop of the custom route entry. Valid values:
	//
	// 	- Instance (default): an ENS instance.
	//
	// 	- HaVip: a high-availability virtual IP address (HAVIP).
	//
	// 	- NetworkPeer: VPC peering connection.
	//
	// example:
	//
	// Instance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The page number of the returned page. Valid values: integers that are greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the route that you want to query.
	//
	// example:
	//
	// rte-8vbdw66evgv44u2u7v3hx
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// test0
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The route type. Valid values:
	//
	// 	- Custom: custom route
	//
	// 	- System: system route
	//
	// example:
	//
	// Custom
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	// The ID of the route table that you want to query.
	//
	// example:
	//
	// vtb-hp3wdhynneo7fsclox8hs
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeEnsRouteEntryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteEntryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteEntryListRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeEnsRouteEntryListRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeEnsRouteEntryListRequest) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeEnsRouteEntryListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnsRouteEntryListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnsRouteEntryListRequest) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DescribeEnsRouteEntryListRequest) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *DescribeEnsRouteEntryListRequest) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *DescribeEnsRouteEntryListRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeEnsRouteEntryListRequest) SetDestinationCidrBlock(v string) *DescribeEnsRouteEntryListRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) SetNextHopId(v string) *DescribeEnsRouteEntryListRequest {
	s.NextHopId = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) SetNextHopType(v string) *DescribeEnsRouteEntryListRequest {
	s.NextHopType = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) SetPageNumber(v int32) *DescribeEnsRouteEntryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) SetPageSize(v int32) *DescribeEnsRouteEntryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) SetRouteEntryId(v string) *DescribeEnsRouteEntryListRequest {
	s.RouteEntryId = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) SetRouteEntryName(v string) *DescribeEnsRouteEntryListRequest {
	s.RouteEntryName = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) SetRouteEntryType(v string) *DescribeEnsRouteEntryListRequest {
	s.RouteEntryType = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) SetRouteTableId(v string) *DescribeEnsRouteEntryListRequest {
	s.RouteTableId = &v
	return s
}

func (s *DescribeEnsRouteEntryListRequest) Validate() error {
	return dara.Validate(s)
}
