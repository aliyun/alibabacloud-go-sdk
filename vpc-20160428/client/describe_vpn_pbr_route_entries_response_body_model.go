// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnPbrRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpnPbrRouteEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnPbrRouteEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnPbrRouteEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpnPbrRouteEntriesResponseBody
	GetTotalCount() *int32
	SetVpnPbrRouteEntries(v *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries) *DescribeVpnPbrRouteEntriesResponseBody
	GetVpnPbrRouteEntries() *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries
}

type DescribeVpnPbrRouteEntriesResponseBody struct {
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
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount         *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpnPbrRouteEntries *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries `json:"VpnPbrRouteEntries,omitempty" xml:"VpnPbrRouteEntries,omitempty" type:"Struct"`
}

func (s DescribeVpnPbrRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnPbrRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) GetVpnPbrRouteEntries() *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries {
	return s.VpnPbrRouteEntries
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeVpnPbrRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) SetPageSize(v int32) *DescribeVpnPbrRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) SetRequestId(v string) *DescribeVpnPbrRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeVpnPbrRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) SetVpnPbrRouteEntries(v *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries) *DescribeVpnPbrRouteEntriesResponseBody {
	s.VpnPbrRouteEntries = v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBody) Validate() error {
	if s.VpnPbrRouteEntries != nil {
		if err := s.VpnPbrRouteEntries.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries struct {
	VpnPbrRouteEntry []*DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry `json:"VpnPbrRouteEntry,omitempty" xml:"VpnPbrRouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries) GetVpnPbrRouteEntry() []*DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	return s.VpnPbrRouteEntry
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries) SetVpnPbrRouteEntry(v []*DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries {
	s.VpnPbrRouteEntry = v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntries) Validate() error {
	if s.VpnPbrRouteEntry != nil {
		for _, item := range s.VpnPbrRouteEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry struct {
	CreateTime      *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	NextHop         *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	NextHopTunnelId *string `json:"NextHopTunnelId,omitempty" xml:"NextHopTunnelId,omitempty"`
	Priority        *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RouteDest       *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	RouteSource     *string `json:"RouteSource,omitempty" xml:"RouteSource,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	VpnInstanceId   *string `json:"VpnInstanceId,omitempty" xml:"VpnInstanceId,omitempty"`
	Weight          *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetNextHop() *string {
	return s.NextHop
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetNextHopTunnelId() *string {
	return s.NextHopTunnelId
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetRouteDest() *string {
	return s.RouteDest
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetRouteSource() *string {
	return s.RouteSource
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetState() *string {
	return s.State
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetVpnInstanceId() *string {
	return s.VpnInstanceId
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetCreateTime(v int64) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetNextHop(v string) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.NextHop = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetNextHopTunnelId(v string) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.NextHopTunnelId = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetPriority(v int32) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.Priority = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetRouteDest(v string) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.RouteDest = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetRouteSource(v string) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.RouteSource = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetState(v string) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.State = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetVpnInstanceId(v string) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.VpnInstanceId = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) SetWeight(v int32) *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry {
	s.Weight = &v
	return s
}

func (s *DescribeVpnPbrRouteEntriesResponseBodyVpnPbrRouteEntriesVpnPbrRouteEntry) Validate() error {
	return dara.Validate(s)
}
