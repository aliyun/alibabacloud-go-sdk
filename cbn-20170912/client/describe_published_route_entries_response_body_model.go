// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePublishedRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePublishedRouteEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePublishedRouteEntriesResponseBody
	GetPageSize() *int32
	SetPublishedRouteEntries(v *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) *DescribePublishedRouteEntriesResponseBody
	GetPublishedRouteEntries() *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries
	SetRequestId(v string) *DescribePublishedRouteEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePublishedRouteEntriesResponseBody
	GetTotalCount() *int32
}

type DescribePublishedRouteEntriesResponseBody struct {
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
	// A list of routes.
	PublishedRouteEntries *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries `json:"PublishedRouteEntries,omitempty" xml:"PublishedRouteEntries,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FF1A7B2A-677F-4F71-96EA-6002B329F437
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePublishedRouteEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePublishedRouteEntriesResponseBody) GetPublishedRouteEntries() *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	return s.PublishedRouteEntries
}

func (s *DescribePublishedRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePublishedRouteEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPageNumber(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPageSize(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetPublishedRouteEntries(v *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) *DescribePublishedRouteEntriesResponseBody {
	s.PublishedRouteEntries = v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetRequestId(v string) *DescribePublishedRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) SetTotalCount(v int32) *DescribePublishedRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries struct {
	PublishedRouteEntry []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry `json:"PublishedRouteEntry,omitempty" xml:"PublishedRouteEntry,omitempty" type:"Repeated"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) GetPublishedRouteEntry() []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	return s.PublishedRouteEntry
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) SetPublishedRouteEntry(v []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries {
	s.PublishedRouteEntry = v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntries) Validate() error {
	return dara.Validate(s)
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry struct {
	// The ID of the route table configured on the network instance.
	//
	// example:
	//
	// vtb-il7qut3mjgtlcbpk2****
	ChildInstanceRouteTableId *string `json:"ChildInstanceRouteTableId,omitempty" xml:"ChildInstanceRouteTableId,omitempty"`
	// A list of conflicting routes.
	Conflicts *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts `json:"Conflicts,omitempty" xml:"Conflicts,omitempty" type:"Struct"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 172.16.0.0/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the instance specified as the next hop in the route.
	//
	// example:
	//
	// ecs-bp18sth14qii3pn****
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The type of the instance specified as the next hop in the route.
	//
	// 	- **Instance**: ECS instance
	//
	// 	- **HaVip**: high-availability virtual IP address (HAVIP).
	//
	// 	- **RouterInterface**: router interface.
	//
	// 	- **NetworkInterface**: elastic network interface (ENI).
	//
	// 	- **VpnGateway**: VPN gateway.
	//
	// 	- **IPv6Gateway**: IPv6 gateway.
	//
	// 	- **NatGateway**: NAT gateway.
	//
	// 	- **Attachment**: network instance connection
	//
	// 	- **service**: cloud service
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// 	- **VPC**: VPC
	//
	// 	- **local**: system route. No next hop is specified.
	//
	// 	- **TR**: transit router
	//
	// 	- **BlackHole**: blackhole route. No next hop is specified.
	//
	// 	- **EcRouterInterface**: router interface for Express Connect
	//
	// 	- **HealthCheck**: health check
	//
	// 	- **AS**: access gateway for CCN
	//
	// 	- **classicLink**: classic network-type instance
	//
	// 	- **GatewayEndpoint**: gateway endpoint
	//
	// 	- **CPE**: data center connected to the VBR
	//
	// example:
	//
	// Instance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// Indicates whether the route is allowed to be advertised to or withdrawn from the CEN instance. Valid values:
	//
	// 	- **true**: The route is allowed to be advertised to or withdrawn from the CEN instance.
	//
	// 	- **false**: The route is not allowed to be advertised to or withdrawn from the CEN instance.
	//
	// example:
	//
	// true
	OperationalMode *bool `json:"OperationalMode,omitempty" xml:"OperationalMode,omitempty"`
	// Indicates whether the route is advertised to the CEN instance. Valid values:
	//
	// 	- **Published**: The route is advertised to the CEN instance.
	//
	// 	- **NonPublished**: The route is not advertised to the CEN instance.
	//
	// example:
	//
	// Published
	PublishStatus *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The type of the route. Valid values:
	//
	// 	- **CEN**: route that is advertised through CEN
	//
	// 	- **System**: system route
	//
	// 	- **Custom**: custom route
	//
	// example:
	//
	// System
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GetChildInstanceRouteTableId() *string {
	return s.ChildInstanceRouteTableId
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GetConflicts() *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts {
	return s.Conflicts
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GetOperationalMode() *bool {
	return s.OperationalMode
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) GetRouteType() *string {
	return s.RouteType
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetChildInstanceRouteTableId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.ChildInstanceRouteTableId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetConflicts(v *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.Conflicts = v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetNextHopId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.NextHopId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetNextHopType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetOperationalMode(v bool) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.OperationalMode = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetPublishStatus(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.PublishStatus = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) SetRouteType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry {
	s.RouteType = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntry) Validate() error {
	return dara.Validate(s)
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts struct {
	Conflict []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict `json:"Conflict,omitempty" xml:"Conflict,omitempty" type:"Repeated"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) GetConflict() []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	return s.Conflict
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) SetConflict(v []*DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts {
	s.Conflict = v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflicts) Validate() error {
	return dara.Validate(s)
}

type DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict struct {
	// The destination CIDR block of the conflicting route.
	//
	// example:
	//
	// 192.168.20.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the peer network instance on which conflicting routes are found.
	//
	// example:
	//
	// ccn-0q3b7oviikmm9h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the peer network instance on which the conflicting routes are found. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// example:
	//
	// CCN
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the region where the peer network instance on which the conflicting routes are found is deployed.
	//
	// example:
	//
	// ccn-cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cause of the route confliction. Valid values:
	//
	// 	- **conflict**: The routes have the same destination CIDR block.
	//
	// 	- **overflow**: The number of routes in the route table configured on another network instance has reached the upper limit.
	//
	// example:
	//
	// conflict
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) String() string {
	return dara.Prettify(s)
}

func (s DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) GoString() string {
	return s.String()
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) GetStatus() *string {
	return s.Status
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetDestinationCidrBlock(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetInstanceId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.InstanceId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetInstanceType(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.InstanceType = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetRegionId(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.RegionId = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) SetStatus(v string) *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict {
	s.Status = &v
	return s
}

func (s *DescribePublishedRouteEntriesResponseBodyPublishedRouteEntriesPublishedRouteEntryConflictsConflict) Validate() error {
	return dara.Validate(s)
}
