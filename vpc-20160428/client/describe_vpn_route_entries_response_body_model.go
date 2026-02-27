// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpnRouteEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnRouteEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnRouteEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpnRouteEntriesResponseBody
	GetTotalCount() *int32
	SetVpnRouteCounts(v *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts) *DescribeVpnRouteEntriesResponseBody
	GetVpnRouteCounts() *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts
	SetVpnRouteEntries(v *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries) *DescribeVpnRouteEntriesResponseBody
	GetVpnRouteEntries() *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries
}

type DescribeVpnRouteEntriesResponseBody struct {
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
	// BF3995A6-FA4F-4C74-B90F-89ECF4BFF4D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount      *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpnRouteCounts  *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts  `json:"VpnRouteCounts,omitempty" xml:"VpnRouteCounts,omitempty" type:"Struct"`
	VpnRouteEntries *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries `json:"VpnRouteEntries,omitempty" xml:"VpnRouteEntries,omitempty" type:"Struct"`
}

func (s DescribeVpnRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnRouteEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnRouteEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnRouteEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpnRouteEntriesResponseBody) GetVpnRouteCounts() *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts {
	return s.VpnRouteCounts
}

func (s *DescribeVpnRouteEntriesResponseBody) GetVpnRouteEntries() *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries {
	return s.VpnRouteEntries
}

func (s *DescribeVpnRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeVpnRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBody) SetPageSize(v int32) *DescribeVpnRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBody) SetRequestId(v string) *DescribeVpnRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeVpnRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBody) SetVpnRouteCounts(v *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts) *DescribeVpnRouteEntriesResponseBody {
	s.VpnRouteCounts = v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBody) SetVpnRouteEntries(v *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries) *DescribeVpnRouteEntriesResponseBody {
	s.VpnRouteEntries = v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBody) Validate() error {
	if s.VpnRouteCounts != nil {
		if err := s.VpnRouteCounts.Validate(); err != nil {
			return err
		}
	}
	if s.VpnRouteEntries != nil {
		if err := s.VpnRouteEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnRouteEntriesResponseBodyVpnRouteCounts struct {
	VpnRouteCount []*DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount `json:"VpnRouteCount,omitempty" xml:"VpnRouteCount,omitempty" type:"Repeated"`
}

func (s DescribeVpnRouteEntriesResponseBodyVpnRouteCounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnRouteEntriesResponseBodyVpnRouteCounts) GoString() string {
	return s.String()
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts) GetVpnRouteCount() []*DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount {
	return s.VpnRouteCount
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts) SetVpnRouteCount(v []*DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts {
	s.VpnRouteCount = v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCounts) Validate() error {
	if s.VpnRouteCount != nil {
		for _, item := range s.VpnRouteCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount struct {
	RouteCount     *int32  `json:"RouteCount,omitempty" xml:"RouteCount,omitempty"`
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) GoString() string {
	return s.String()
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) GetRouteCount() *int32 {
	return s.RouteCount
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) GetSource() *string {
	return s.Source
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) SetRouteCount(v int32) *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount {
	s.RouteCount = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) SetRouteEntryType(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount {
	s.RouteEntryType = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) SetSource(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount {
	s.Source = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteCountsVpnRouteCount) Validate() error {
	return dara.Validate(s)
}

type DescribeVpnRouteEntriesResponseBodyVpnRouteEntries struct {
	VpnRouteEntry []*DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry `json:"VpnRouteEntry,omitempty" xml:"VpnRouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeVpnRouteEntriesResponseBodyVpnRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnRouteEntriesResponseBodyVpnRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries) GetVpnRouteEntry() []*DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	return s.VpnRouteEntry
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries) SetVpnRouteEntry(v []*DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries {
	s.VpnRouteEntry = v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntries) Validate() error {
	if s.VpnRouteEntry != nil {
		for _, item := range s.VpnRouteEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry struct {
	AsPath          *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	Community       *string `json:"Community,omitempty" xml:"Community,omitempty"`
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NextHop         *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	NextHopTunnelId *string `json:"NextHopTunnelId,omitempty" xml:"NextHopTunnelId,omitempty"`
	RouteDest       *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	RouteEntryType  *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	VpnInstanceId   *string `json:"VpnInstanceId,omitempty" xml:"VpnInstanceId,omitempty"`
	Weight          *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetAsPath() *string {
	return s.AsPath
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetCommunity() *string {
	return s.Community
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetNextHop() *string {
	return s.NextHop
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetNextHopTunnelId() *string {
	return s.NextHopTunnelId
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetRouteDest() *string {
	return s.RouteDest
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetSource() *string {
	return s.Source
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetState() *string {
	return s.State
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetVpnInstanceId() *string {
	return s.VpnInstanceId
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetAsPath(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.AsPath = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetCommunity(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.Community = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetCreateTime(v int64) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetNextHop(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.NextHop = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetNextHopTunnelId(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.NextHopTunnelId = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetRouteDest(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.RouteDest = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetRouteEntryType(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.RouteEntryType = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetSource(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.Source = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetState(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.State = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetVpnInstanceId(v string) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.VpnInstanceId = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) SetWeight(v int32) *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry {
	s.Weight = &v
	return s
}

func (s *DescribeVpnRouteEntriesResponseBodyVpnRouteEntriesVpnRouteEntry) Validate() error {
	return dara.Validate(s)
}
