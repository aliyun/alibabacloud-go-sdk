// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenChildInstanceRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenRouteEntries(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) *DescribeCenChildInstanceRouteEntriesResponseBody
	GetCenRouteEntries() *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries
	SetPageNumber(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCenChildInstanceRouteEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeCenChildInstanceRouteEntriesResponseBody struct {
	// The information about the route.
	CenRouteEntries *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries `json:"CenRouteEntries,omitempty" xml:"CenRouteEntries,omitempty" type:"Struct"`
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
	// The ID of the request.
	//
	// example:
	//
	// 17A57456-EF48-419D-9AE6-9B03D9996018
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) GetCenRouteEntries() *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	return s.CenRouteEntries
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetCenRouteEntries(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.CenRouteEntries = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetPageSize(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetRequestId(v string) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeCenChildInstanceRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBody) Validate() error {
	if s.CenRouteEntries != nil {
		if err := s.CenRouteEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries struct {
	CenRouteEntry []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry `json:"CenRouteEntry,omitempty" xml:"CenRouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) GetCenRouteEntry() []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	return s.CenRouteEntry
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) SetCenRouteEntry(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries {
	s.CenRouteEntry = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntries) Validate() error {
	if s.CenRouteEntry != nil {
		for _, item := range s.CenRouteEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry struct {
	// The AS paths of the routes.
	AsPaths *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths `json:"AsPaths,omitempty" xml:"AsPaths,omitempty" type:"Struct"`
	// The routing policy that the routes match.
	CenRouteMapRecords *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords `json:"CenRouteMapRecords,omitempty" xml:"CenRouteMapRecords,omitempty" type:"Struct"`
	// The community attributes of the route entries.
	Communities *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities `json:"Communities,omitempty" xml:"Communities,omitempty" type:"Struct"`
	// A list of overlapping routes.
	Conflicts *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts `json:"Conflicts,omitempty" xml:"Conflicts,omitempty" type:"Struct"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the instance specified as the next hop in the route.
	//
	// example:
	//
	// vbr-bp13gtbhdp0pfqg6s****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	// The region ID of the instance specified as the next hop in the route.
	//
	// example:
	//
	// cn-hangzhou
	NextHopRegionId *string `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	// The type of the instance specified as the next hop in the route. Valid values:
	//
	// 	- **Instance**: an ECS instance
	//
	// 	- **HaVip**: an HAVIP
	//
	// 	- **RouterInterface**: a router interface
	//
	// 	- **NetworkInterface**: an ENI
	//
	// 	- **VpnGateway**: a VPN gateway
	//
	// 	- **IPv6Gateway**: an IPv6 gateway
	//
	// 	- **Ipv4Gateway**: an IPv4 gateway
	//
	// 	- **NatGateway**: a NAT gateway
	//
	// 	- **Attachment**: a network instance connection
	//
	// 	- **service**: a cloud service
	//
	// 	- **VBR**: a VBR
	//
	// 	- **CCN**: a CCN instance
	//
	// 	- **VPC**: a VPC
	//
	// 	- **local**: a system route (no next hop is specified)
	//
	// 	- **TR**: a transit router
	//
	// 	- **BlackHole**: a blackhole route (no next hop is specified)
	//
	// 	- **EcRouterInterface**: a router interface for Express Connect
	//
	// 	- **HealthCheck**: a health check
	//
	// 	- **AS**: an access gateway for CCN
	//
	// 	- **classic**: a classic network-type instance
	//
	// 	- **GatewayEndpoint**: a gateway endpoint
	//
	// 	- **CPE**: a data center connected to a VBR
	//
	// example:
	//
	// VBR
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// Indicates whether the route is allowed to be advertised to or withdrawn from the CEN instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	OperationalMode *bool `json:"OperationalMode,omitempty" xml:"OperationalMode,omitempty"`
	// Indicates whether the route is advertised to the CEN instance. Valid values:
	//
	// 	- **Published**
	//
	// 	- **NonPublished**
	//
	// example:
	//
	// Published
	PublishStatus *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The route table ID.
	//
	// example:
	//
	// vtb-bp1r9pvl4xen8s9ju****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The route status. Valid values:
	//
	// 	- **Active**: available routes
	//
	// 	- **Candidate**: standby routes
	//
	// 	- **Rejected**: rejected routes
	//
	// 	- **Prohibited**: prohibited routes
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The route type. Valid values:
	//
	// 	- **CEN**: route that is advertised through CEN
	//
	// 	- **System**: system route
	//
	// 	- **Custom**: custom route
	//
	// example:
	//
	// CEN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetAsPaths() *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths {
	return s.AsPaths
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetCenRouteMapRecords() *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords {
	return s.CenRouteMapRecords
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetCommunities() *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities {
	return s.Communities
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetConflicts() *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts {
	return s.Conflicts
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetNextHopRegionId() *string {
	return s.NextHopRegionId
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetOperationalMode() *bool {
	return s.OperationalMode
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetType() *string {
	return s.Type
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetAsPaths(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.AsPaths = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenRouteMapRecords(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenRouteMapRecords = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCommunities(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Communities = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetConflicts(v *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Conflicts = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetDestinationCidrBlock(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopInstanceId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopRegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetOperationalMode(v bool) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.OperationalMode = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetPublishStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.PublishStatus = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetRouteTableId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.RouteTableId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Type = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) Validate() error {
	if s.AsPaths != nil {
		if err := s.AsPaths.Validate(); err != nil {
			return err
		}
	}
	if s.CenRouteMapRecords != nil {
		if err := s.CenRouteMapRecords.Validate(); err != nil {
			return err
		}
	}
	if s.Communities != nil {
		if err := s.Communities.Validate(); err != nil {
			return err
		}
	}
	if s.Conflicts != nil {
		if err := s.Conflicts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths struct {
	AsPath []*string `json:"AsPath,omitempty" xml:"AsPath,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) GetAsPath() []*string {
	return s.AsPath
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) SetAsPath(v []*string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths {
	s.AsPath = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) Validate() error {
	return dara.Validate(s)
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords struct {
	CenRouteMapRecord []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord `json:"CenRouteMapRecord,omitempty" xml:"CenRouteMapRecord,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) GetCenRouteMapRecord() []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	return s.CenRouteMapRecord
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) SetCenRouteMapRecord(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords {
	s.CenRouteMapRecord = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) Validate() error {
	if s.CenRouteMapRecord != nil {
		for _, item := range s.CenRouteMapRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord struct {
	// The region ID of the routing policy.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The routing policy ID.
	//
	// example:
	//
	// cenrmap-w4yf7toozfol3q****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRouteMapId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) Validate() error {
	return dara.Validate(s)
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities struct {
	Community []*string `json:"Community,omitempty" xml:"Community,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) GetCommunity() []*string {
	return s.Community
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) SetCommunity(v []*string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities {
	s.Community = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) Validate() error {
	return dara.Validate(s)
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts struct {
	Conflict []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict `json:"Conflict,omitempty" xml:"Conflict,omitempty" type:"Repeated"`
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) GetConflict() []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	return s.Conflict
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) SetConflict(v []*DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts {
	s.Conflict = v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflicts) Validate() error {
	if s.Conflict != nil {
		for _, item := range s.Conflict {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict struct {
	// The destination CIDR block of the overlapping route.
	//
	// example:
	//
	// 192.168.1.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the peer network instance on which the overlapping routes are found.
	//
	// example:
	//
	// ccn-0q3b7oviikmm9h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the peer network instance on which the overlapping routes are found. Valid values:
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// 	- **ECR**: ECR
	//
	// example:
	//
	// CCN
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID of the peer network instance on which the overlapping routes are found.
	//
	// example:
	//
	// ccn-cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cause of the route error. Valid values:
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

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) GoString() string {
	return s.String()
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetDestinationCidrBlock(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetInstanceId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.InstanceId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetInstanceType(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.InstanceType = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetRegionId(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.RegionId = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) SetStatus(v string) *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict {
	s.Status = &v
	return s
}

func (s *DescribeCenChildInstanceRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryConflictsConflict) Validate() error {
	return dara.Validate(s)
}
