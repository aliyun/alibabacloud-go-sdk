// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRouteTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociateType(v string) *DescribeEnsRouteTablesRequest
	GetAssociateType() *string
	SetEnsRegionId(v string) *DescribeEnsRouteTablesRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeEnsRouteTablesRequest
	GetEnsRegionIds() []*string
	SetNetworkId(v string) *DescribeEnsRouteTablesRequest
	GetNetworkId() *string
	SetPageNumber(v int32) *DescribeEnsRouteTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnsRouteTablesRequest
	GetPageSize() *int32
	SetRouteTableId(v string) *DescribeEnsRouteTablesRequest
	GetRouteTableId() *string
	SetRouteTableName(v string) *DescribeEnsRouteTablesRequest
	GetRouteTableName() *string
	SetType(v string) *DescribeEnsRouteTablesRequest
	GetType() *string
}

type DescribeEnsRouteTablesRequest struct {
	// The type of the resource with which the route table is associated. Valid values:
	//
	// 	- **VSwitch**
	//
	// 	- **Gateway**
	//
	// example:
	//
	// Gateway
	AssociateType *string `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-xian-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of edge nodes. You can specify 1 to 100 IDs.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The ID of the network.
	//
	// example:
	//
	// n-257gqcdfvx6n****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-5p1cifr72di****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The name of the route table.
	//
	// example:
	//
	// tftest-nat04
	RouteTableName *string `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
	// The SNAT type.
	//
	// 	- FullCone: full cone NAT.
	//
	// example:
	//
	// FullCone
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeEnsRouteTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteTablesRequest) GetAssociateType() *string {
	return s.AssociateType
}

func (s *DescribeEnsRouteTablesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsRouteTablesRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeEnsRouteTablesRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeEnsRouteTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnsRouteTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnsRouteTablesRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeEnsRouteTablesRequest) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *DescribeEnsRouteTablesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeEnsRouteTablesRequest) SetAssociateType(v string) *DescribeEnsRouteTablesRequest {
	s.AssociateType = &v
	return s
}

func (s *DescribeEnsRouteTablesRequest) SetEnsRegionId(v string) *DescribeEnsRouteTablesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRouteTablesRequest) SetEnsRegionIds(v []*string) *DescribeEnsRouteTablesRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeEnsRouteTablesRequest) SetNetworkId(v string) *DescribeEnsRouteTablesRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeEnsRouteTablesRequest) SetPageNumber(v int32) *DescribeEnsRouteTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRouteTablesRequest) SetPageSize(v int32) *DescribeEnsRouteTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRouteTablesRequest) SetRouteTableId(v string) *DescribeEnsRouteTablesRequest {
	s.RouteTableId = &v
	return s
}

func (s *DescribeEnsRouteTablesRequest) SetRouteTableName(v string) *DescribeEnsRouteTablesRequest {
	s.RouteTableName = &v
	return s
}

func (s *DescribeEnsRouteTablesRequest) SetType(v string) *DescribeEnsRouteTablesRequest {
	s.Type = &v
	return s
}

func (s *DescribeEnsRouteTablesRequest) Validate() error {
	return dara.Validate(s)
}
