// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRouteTablesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteTablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRouteTablesResponseBody
	GetRequestId() *string
	SetRouteTables(v *DescribeRouteTablesResponseBodyRouteTables) *DescribeRouteTablesResponseBody
	GetRouteTables() *DescribeRouteTablesResponseBodyRouteTables
	SetTotalCount(v int32) *DescribeRouteTablesResponseBody
	GetTotalCount() *int32
}

type DescribeRouteTablesResponseBody struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DC668356-BCB4-42FD-9BC3-FA2B2E04B634
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed information about the route tables.
	RouteTables *DescribeRouteTablesResponseBodyRouteTables `json:"RouteTables,omitempty" xml:"RouteTables,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouteTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteTablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteTablesResponseBody) GetRouteTables() *DescribeRouteTablesResponseBodyRouteTables {
	return s.RouteTables
}

func (s *DescribeRouteTablesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRouteTablesResponseBody) SetPageNumber(v int32) *DescribeRouteTablesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteTablesResponseBody) SetPageSize(v int32) *DescribeRouteTablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteTablesResponseBody) SetRequestId(v string) *DescribeRouteTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteTablesResponseBody) SetRouteTables(v *DescribeRouteTablesResponseBodyRouteTables) *DescribeRouteTablesResponseBody {
	s.RouteTables = v
	return s
}

func (s *DescribeRouteTablesResponseBody) SetTotalCount(v int32) *DescribeRouteTablesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteTablesResponseBody) Validate() error {
	if s.RouteTables != nil {
		if err := s.RouteTables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTables struct {
	RouteTable []*DescribeRouteTablesResponseBodyRouteTablesRouteTable `json:"RouteTable,omitempty" xml:"RouteTable,omitempty" type:"Repeated"`
}

func (s DescribeRouteTablesResponseBodyRouteTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTables) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTables) GetRouteTable() []*DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	return s.RouteTable
}

func (s *DescribeRouteTablesResponseBodyRouteTables) SetRouteTable(v []*DescribeRouteTablesResponseBodyRouteTablesRouteTable) *DescribeRouteTablesResponseBodyRouteTables {
	s.RouteTable = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTables) Validate() error {
	if s.RouteTable != nil {
		for _, item := range s.RouteTable {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTable struct {
	// The time when the route table was created.
	//
	// The time is displayed in the `YYYY-MM-DDThh:mm:ssZ` format in UTC.
	//
	// example:
	//
	// 2017-08-22T10:40:25Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the resource group to which the route table belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The information about the route.
	RouteEntrys *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys `json:"RouteEntrys,omitempty" xml:"RouteEntrys,omitempty" type:"Struct"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The type of the route table. Valid values:
	//
	// 	- **Custom**
	//
	// 	- **System**
	//
	// example:
	//
	// System
	RouteTableType *string `json:"RouteTableType,omitempty" xml:"RouteTableType,omitempty"`
	// The status of the route table. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vRouter ID.
	//
	// example:
	//
	// vrt-bp1lhl0taikrteen8****
	VRouterId *string `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	// The vSwitch ID.
	VSwitchIds *DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTable) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetRouteEntrys() *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys {
	return s.RouteEntrys
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetRouteTableType() *string {
	return s.RouteTableType
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetVRouterId() *string {
	return s.VRouterId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) GetVSwitchIds() *DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds {
	return s.VSwitchIds
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetCreationTime(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.CreationTime = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetResourceGroupId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetRouteEntrys(v *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.RouteEntrys = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetRouteTableId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetRouteTableType(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.RouteTableType = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetStatus(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.Status = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetVRouterId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.VRouterId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) SetVSwitchIds(v *DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds) *DescribeRouteTablesResponseBodyRouteTablesRouteTable {
	s.VSwitchIds = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTable) Validate() error {
	if s.RouteEntrys != nil {
		if err := s.RouteEntrys.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchIds != nil {
		if err := s.VSwitchIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys struct {
	RouteEntry []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry `json:"RouteEntry,omitempty" xml:"RouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) GetRouteEntry() []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	return s.RouteEntry
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) SetRouteEntry(v []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys {
	s.RouteEntry = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrys) Validate() error {
	if s.RouteEntry != nil {
		for _, item := range s.RouteEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry struct {
	// The description of the route. The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// RouteEntryDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the route. The destination CIDR block supports IPv4 and IPv6. Make sure that the destination CIDR block meets the following requirements:
	//
	// 	- The destination CIDR block is not 100.64.0.0/10 or a subset of 100.64.0.0/10.
	//
	// 	- The destination CIDR block of each route in the route table is unique.
	//
	// example:
	//
	// 192.168.0.1/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the instance associated with the next hop.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urzd****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **Instance*	- (default): an Elastic Compute Service (ECS) instance
	//
	// 	- **HaVip**: a high-availability virtual IP address (HaVip).
	//
	// 	- **VpnGateway**: a VPN gateway
	//
	// 	- **NatGateway**: a NAT gateway
	//
	// 	- **NetworkInterface**: a secondary elastic network interface (ENI)
	//
	// 	- **RouterInterface**: a router interface
	//
	// 	- **IPv6Gateway**: an IPv6 gateway
	//
	// 	- **Attachment**: a transit router
	//
	// example:
	//
	// local
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The information about the next hop.
	NextHops *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops `json:"NextHops,omitempty" xml:"NextHops,omitempty" type:"Struct"`
	// The ID of the route.
	//
	// example:
	//
	// rte-bp1mnnr2al0naomnpxxx
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The route name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The route table ID.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The route status. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The route type. Valid values:
	//
	// 	- **Custom**
	//
	// 	- **System**
	//
	// 	- **BGP**
	//
	// 	- **CEN**
	//
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetNextHops() *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops {
	return s.NextHops
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) GetType() *string {
	return s.Type
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetDescription(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.Description = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetDestinationCidrBlock(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetInstanceId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.InstanceId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetNextHopType(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetNextHops(v *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.NextHops = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetRouteEntryId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.RouteEntryId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetRouteEntryName(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.RouteEntryName = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetRouteTableId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetStatus(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) SetType(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry {
	s.Type = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntry) Validate() error {
	if s.NextHops != nil {
		if err := s.NextHops.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops struct {
	NextHop []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop `json:"NextHop,omitempty" xml:"NextHop,omitempty" type:"Repeated"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) GetNextHop() []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	return s.NextHop
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) SetNextHop(v []*DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops {
	s.NextHop = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHops) Validate() error {
	if s.NextHop != nil {
		for _, item := range s.NextHop {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop struct {
	// Indicates whether the route is available. Valid values:
	//
	// 	- **0**: unavailable
	//
	// 	- **1**: available
	//
	// example:
	//
	// 0
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the next hop.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urzdpvqw
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of the next hop. Valid values:
	//
	// 	- **Instance**: an ECS instance
	//
	// 	- **HaVip**: an HaVip
	//
	// 	- **VpnGateway**: a VPN gateway
	//
	// 	- **NatGateway**: a NAT gateway
	//
	// 	- **NetworkInterface**: a secondary ENI
	//
	// 	- **RouterInterface**: a router interface
	//
	// 	- **IPv6Gateway**: an IPv6 gateway
	//
	// 	- **Attachment**: a transit router
	//
	// example:
	//
	// HaVip
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The weight of the route.
	//
	// example:
	//
	// 80
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GetEnabled() *int32 {
	return s.Enabled
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) SetEnabled(v int32) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	s.Enabled = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) SetNextHopId(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopId = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) SetNextHopType(v string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopType = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) SetWeight(v int32) *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop {
	s.Weight = &v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableRouteEntrysRouteEntryNextHopsNextHop) Validate() error {
	return dara.Validate(s)
}

type DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds) SetVSwitchId(v []*string) *DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds {
	s.VSwitchId = v
	return s
}

func (s *DescribeRouteTablesResponseBodyRouteTablesRouteTableVSwitchIds) Validate() error {
	return dara.Validate(s)
}
