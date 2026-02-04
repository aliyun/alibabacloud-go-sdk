// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenRegionDomainRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenRouteEntries(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) *DescribeCenRegionDomainRouteEntriesResponseBody
	GetCenRouteEntries() *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries
	SetPageNumber(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCenRegionDomainRouteEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeCenRegionDomainRouteEntriesResponseBody struct {
	// A list of route entries.
	CenRouteEntries *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries `json:"CenRouteEntries,omitempty" xml:"CenRouteEntries,omitempty" type:"Struct"`
	// The page number of the returned page.
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
	// 004E99FB-E996-5777-888E-BA1D8F215407
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) GetCenRouteEntries() *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	return s.CenRouteEntries
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetCenRouteEntries(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.CenRouteEntries = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetPageSize(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetRequestId(v string) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeCenRegionDomainRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBody) Validate() error {
	if s.CenRouteEntries != nil {
		if err := s.CenRouteEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries struct {
	CenRouteEntry []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry `json:"CenRouteEntry,omitempty" xml:"CenRouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) GetCenRouteEntry() []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	return s.CenRouteEntry
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) SetCenRouteEntry(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries {
	s.CenRouteEntry = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntries) Validate() error {
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

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry struct {
	// The AS paths of the route.
	AsPaths *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths `json:"AsPaths,omitempty" xml:"AsPaths,omitempty" type:"Struct"`
	// The routing policy that the routes match in the outbound direction.
	CenOutRouteMapRecords *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords `json:"CenOutRouteMapRecords,omitempty" xml:"CenOutRouteMapRecords,omitempty" type:"Struct"`
	// The routing policy that the routes match in the inbound direction.
	CenRouteMapRecords *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords `json:"CenRouteMapRecords,omitempty" xml:"CenRouteMapRecords,omitempty" type:"Struct"`
	// The communities of the route.
	Communities *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities `json:"Communities,omitempty" xml:"Communities,omitempty" type:"Struct"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 192.168.1.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the instance specified as the next hop in the route.
	//
	// example:
	//
	// vpc-bp1j8728mm6pweeod****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
	// The ID of the region to which the network instance specified as the next hop in the route belongs.
	//
	// example:
	//
	// cn-hangzhou
	NextHopRegionId *string `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	// The type of the instance specified as the next hop in the route. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VBR**: virtual border router (VBR)
	//
	// 	- **CCN**: Cloud Connect Network (CCN) instance
	//
	// 	- **local_service**: system route. No next hop is specified.
	//
	// example:
	//
	// VPC
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The route priority.
	//
	// A smaller value indicates a higher priority.
	//
	// example:
	//
	// 50
	Preference *int32 `json:"Preference,omitempty" xml:"Preference,omitempty"`
	// The route status. Valid values:
	//
	// 	- **Active**: available
	//
	// 	- **Candidate**: standby
	//
	// 	- **Rejected**: rejected
	//
	// 	- **Prohibited**: prohibited
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the route can be advertised to other regions. Valid values:
	//
	// 	- **Active**: The route can be advertised to other regions.
	//
	// 	- **Prohibited**: The route cannot be advertised to other regions.
	//
	// example:
	//
	// Active
	ToOtherRegionStatus *string `json:"ToOtherRegionStatus,omitempty" xml:"ToOtherRegionStatus,omitempty"`
	// The route type. Valid values:
	//
	// 	- **CEN**: route that is advertised through CEN
	//
	// 	- **Custom**: custom route
	//
	// 	- **System**: system route
	//
	// example:
	//
	// CEN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetAsPaths() *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths {
	return s.AsPaths
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetCenOutRouteMapRecords() *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords {
	return s.CenOutRouteMapRecords
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetCenRouteMapRecords() *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords {
	return s.CenRouteMapRecords
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetCommunities() *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities {
	return s.Communities
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetNextHopRegionId() *string {
	return s.NextHopRegionId
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetPreference() *int32 {
	return s.Preference
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetToOtherRegionStatus() *string {
	return s.ToOtherRegionStatus
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) GetType() *string {
	return s.Type
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetAsPaths(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.AsPaths = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenOutRouteMapRecords(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenOutRouteMapRecords = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCenRouteMapRecords(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.CenRouteMapRecords = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetCommunities(v *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Communities = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetDestinationCidrBlock(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopInstanceId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopRegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetNextHopType(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.NextHopType = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetPreference(v int32) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Preference = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetStatus(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetToOtherRegionStatus(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.ToOtherRegionStatus = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) SetType(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry {
	s.Type = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntry) Validate() error {
	if s.AsPaths != nil {
		if err := s.AsPaths.Validate(); err != nil {
			return err
		}
	}
	if s.CenOutRouteMapRecords != nil {
		if err := s.CenOutRouteMapRecords.Validate(); err != nil {
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
	return nil
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths struct {
	AsPath []*string `json:"AsPath,omitempty" xml:"AsPath,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) GetAsPath() []*string {
	return s.AsPath
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) SetAsPath(v []*string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths {
	s.AsPath = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryAsPaths) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords struct {
	CenOutRouteMapRecord []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord `json:"CenOutRouteMapRecord,omitempty" xml:"CenOutRouteMapRecord,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) GetCenOutRouteMapRecord() []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord {
	return s.CenOutRouteMapRecord
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) SetCenOutRouteMapRecord(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords {
	s.CenOutRouteMapRecord = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecords) Validate() error {
	if s.CenOutRouteMapRecord != nil {
		for _, item := range s.CenOutRouteMapRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord struct {
	// The ID of the region where the routing policy is applied.
	//
	// example:
	//
	// ccn-cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the routing policy.
	//
	// example:
	//
	// cenrmap-dbarzidzp7ek4k****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) SetRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord {
	s.RegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) SetRouteMapId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenOutRouteMapRecordsCenOutRouteMapRecord) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords struct {
	CenRouteMapRecord []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord `json:"CenRouteMapRecord,omitempty" xml:"CenRouteMapRecord,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) GetCenRouteMapRecord() []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	return s.CenRouteMapRecord
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) SetCenRouteMapRecord(v []*DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords {
	s.CenRouteMapRecord = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecords) Validate() error {
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

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord struct {
	// The ID of the region where the routing policy is applied.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the routing policy.
	//
	// example:
	//
	// cenrmap-cz5axczdxb7yfu****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRegionId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RegionId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) SetRouteMapId(v string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord {
	s.RouteMapId = &v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCenRouteMapRecordsCenRouteMapRecord) Validate() error {
	return dara.Validate(s)
}

type DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities struct {
	Community []*string `json:"Community,omitempty" xml:"Community,omitempty" type:"Repeated"`
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) GoString() string {
	return s.String()
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) GetCommunity() []*string {
	return s.Community
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) SetCommunity(v []*string) *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities {
	s.Community = v
	return s
}

func (s *DescribeCenRegionDomainRouteEntriesResponseBodyCenRouteEntriesCenRouteEntryCommunities) Validate() error {
	return dara.Validate(s)
}
