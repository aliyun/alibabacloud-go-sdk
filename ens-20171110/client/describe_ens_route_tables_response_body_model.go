// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRouteTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeEnsRouteTablesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnsRouteTablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEnsRouteTablesResponseBody
	GetRequestId() *string
	SetRouteTables(v []*DescribeEnsRouteTablesResponseBodyRouteTables) *DescribeEnsRouteTablesResponseBody
	GetRouteTables() []*DescribeEnsRouteTablesResponseBodyRouteTables
	SetTotalCount(v int32) *DescribeEnsRouteTablesResponseBody
	GetTotalCount() *int32
}

type DescribeEnsRouteTablesResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// DC51ACB0-460D-5CA0-BA2D-E1F3B5547AE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the route tables.
	RouteTables []*DescribeEnsRouteTablesResponseBodyRouteTables `json:"RouteTables,omitempty" xml:"RouteTables,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEnsRouteTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteTablesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnsRouteTablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnsRouteTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsRouteTablesResponseBody) GetRouteTables() []*DescribeEnsRouteTablesResponseBodyRouteTables {
	return s.RouteTables
}

func (s *DescribeEnsRouteTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEnsRouteTablesResponseBody) SetPageNumber(v int32) *DescribeEnsRouteTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBody) SetPageSize(v int32) *DescribeEnsRouteTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBody) SetRequestId(v string) *DescribeEnsRouteTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBody) SetRouteTables(v []*DescribeEnsRouteTablesResponseBodyRouteTables) *DescribeEnsRouteTablesResponseBody {
	s.RouteTables = v
	return s
}

func (s *DescribeEnsRouteTablesResponseBody) SetTotalCount(v int32) *DescribeEnsRouteTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsRouteTablesResponseBodyRouteTables struct {
	// The type of the resource with which the route table is associated. Valid values:
	//
	// 	- **VSwitch**
	//
	// 	- **Gateway**
	//
	// example:
	//
	// VSwitch
	AssociateType *string `json:"AssociateType,omitempty" xml:"AssociateType,omitempty"`
	// The time when the route table was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-03-08T08:35:18Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-beijing-15
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// Specifies whether it is the default gateway route table.
	//
	// example:
	//
	// false
	IsDefaultGatewayRouteTable *bool `json:"IsDefaultGatewayRouteTable,omitempty" xml:"IsDefaultGatewayRouteTable,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5v9lufsezl4g****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// rt-5xde2bit9****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The name of the route table that you want to query.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test-tf-vtb7
	RouteTableName *string `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
	// The status. Valid values:
	//
	// 	- Available: The route table is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the route table. Examples:
	//
	// 	- Custom: custom route table.
	//
	// 	- System: system route table.
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vSwitches that are associated with the route table.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s DescribeEnsRouteTablesResponseBodyRouteTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteTablesResponseBodyRouteTables) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetAssociateType() *string {
	return s.AssociateType
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetDescription() *string {
	return s.Description
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetIsDefaultGatewayRouteTable() *bool {
	return s.IsDefaultGatewayRouteTable
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetType() *string {
	return s.Type
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetAssociateType(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.AssociateType = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetCreationTime(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.CreationTime = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetDescription(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.Description = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetEnsRegionId(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetIsDefaultGatewayRouteTable(v bool) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.IsDefaultGatewayRouteTable = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetNetworkId(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.NetworkId = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetRouteTableId(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.RouteTableId = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetRouteTableName(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.RouteTableName = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetStatus(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.Status = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetType(v string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.Type = &v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) SetVSwitchIds(v []*string) *DescribeEnsRouteTablesResponseBodyRouteTables {
	s.VSwitchIds = v
	return s
}

func (s *DescribeEnsRouteTablesResponseBodyRouteTables) Validate() error {
	return dara.Validate(s)
}
