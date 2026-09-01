// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVcoRouteEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVcoRouteEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVcoRouteEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVcoRouteEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVcoRouteEntriesResponseBody
	GetTotalCount() *int32
	SetVcoRouteEntries(v []*DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) *DescribeVcoRouteEntriesResponseBody
	GetVcoRouteEntries() []*DescribeVcoRouteEntriesResponseBodyVcoRouteEntries
	SetVpnRouteCounts(v []*DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) *DescribeVcoRouteEntriesResponseBody
	GetVpnRouteCounts() []*DescribeVcoRouteEntriesResponseBodyVpnRouteCounts
}

type DescribeVcoRouteEntriesResponseBody struct {
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for paging queries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E18980E8-C8C2-31BD-8156-AE2BBDEC87E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of routes.
	VcoRouteEntries []*DescribeVcoRouteEntriesResponseBodyVcoRouteEntries `json:"VcoRouteEntries,omitempty" xml:"VcoRouteEntries,omitempty" type:"Repeated"`
	// The route statistics for the IPsec-VPN connection in dual-tunnel mode.
	//
	// > This information is returned only for IPsec-VPN connections in dual-tunnel mode.
	VpnRouteCounts []*DescribeVcoRouteEntriesResponseBodyVpnRouteCounts `json:"VpnRouteCounts,omitempty" xml:"VpnRouteCounts,omitempty" type:"Repeated"`
}

func (s DescribeVcoRouteEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVcoRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVcoRouteEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVcoRouteEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVcoRouteEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVcoRouteEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVcoRouteEntriesResponseBody) GetVcoRouteEntries() []*DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	return s.VcoRouteEntries
}

func (s *DescribeVcoRouteEntriesResponseBody) GetVpnRouteCounts() []*DescribeVcoRouteEntriesResponseBodyVpnRouteCounts {
	return s.VpnRouteCounts
}

func (s *DescribeVcoRouteEntriesResponseBody) SetPageNumber(v int32) *DescribeVcoRouteEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBody) SetPageSize(v int32) *DescribeVcoRouteEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBody) SetRequestId(v string) *DescribeVcoRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeVcoRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBody) SetVcoRouteEntries(v []*DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) *DescribeVcoRouteEntriesResponseBody {
	s.VcoRouteEntries = v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBody) SetVpnRouteCounts(v []*DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) *DescribeVcoRouteEntriesResponseBody {
	s.VpnRouteCounts = v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBody) Validate() error {
	if s.VcoRouteEntries != nil {
		for _, item := range s.VcoRouteEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VpnRouteCounts != nil {
		for _, item := range s.VpnRouteCounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVcoRouteEntriesResponseBodyVcoRouteEntries struct {
	// The list of autonomous system (AS) numbers that the BGP route passes through.
	//
	// example:
	//
	// [12000]
	AsPath *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	// The community value carried by the BGP route.
	//
	// example:
	//
	// 65535:65510
	Community *string `json:"Community,omitempty" xml:"Community,omitempty"`
	// The timestamp when the destination route was created.
	//
	// The timestamp is in the UNIX format and represents the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1658217008000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The next hop of the route.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// The list of next-hop tunnels.
	//
	// > - This information is returned only for IPsec-VPN connections in dual-tunnel mode.
	//
	// > - Tunnel information is returned only when the tunnel status is **Phase 2 negotiations succeeded**.
	NextHopTunnelIdList []*string `json:"NextHopTunnelIdList,omitempty" xml:"NextHopTunnelIdList,omitempty" type:"Repeated"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 192.168.10.0/24
	RouteDest *string `json:"RouteDest,omitempty" xml:"RouteDest,omitempty"`
	// The type of the route.
	//
	// - **custom**: The route is a destination route.
	//
	// - **bgp**: The route is a BGP route.
	//
	// example:
	//
	// custom
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	// The source of the BGP route.
	//
	// - **CLOUD**: The BGP route is learned by the IPsec-VPN connection from the transit router.
	//
	// - **VPN_BGP**: The BGP route is learned by the IPsec-VPN connection from the on-premises data center.
	//
	// example:
	//
	// CLOUD
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the route.
	//
	// - **published**: The destination route has been published to the transit router instance.
	//
	// - **Active**: The BGP route is available.
	//
	// example:
	//
	// published
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
	// The weight of the destination route.
	//
	// > This parameter is not in use.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetAsPath() *string {
	return s.AsPath
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetCommunity() *string {
	return s.Community
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetNextHop() *string {
	return s.NextHop
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetNextHopTunnelIdList() []*string {
	return s.NextHopTunnelIdList
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetRouteDest() *string {
	return s.RouteDest
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetSource() *string {
	return s.Source
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetState() *string {
	return s.State
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetAsPath(v string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.AsPath = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetCommunity(v string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.Community = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetCreateTime(v int64) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.CreateTime = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetNextHop(v string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.NextHop = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetNextHopTunnelIdList(v []*string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.NextHopTunnelIdList = v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetRouteDest(v string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.RouteDest = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetRouteEntryType(v string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.RouteEntryType = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetSource(v string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.Source = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetState(v string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.State = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetVpnConnectionId(v string) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) SetWeight(v int32) *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries {
	s.Weight = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVcoRouteEntries) Validate() error {
	return dara.Validate(s)
}

type DescribeVcoRouteEntriesResponseBodyVpnRouteCounts struct {
	// The number of routes.
	//
	// example:
	//
	// 3
	RouteCount *int32 `json:"RouteCount,omitempty" xml:"RouteCount,omitempty"`
	// The type of the route.
	//
	// - **custom**: destination route.
	//
	// - **bgp**: BGP route.
	//
	// example:
	//
	// bgp
	RouteEntryType *string `json:"RouteEntryType,omitempty" xml:"RouteEntryType,omitempty"`
	// The source of the BGP route.
	//
	// - **CLOUD**: The BGP route is learned by the IPsec-VPN connection from the transit router.
	//
	// - **VPN_BGP**: The BGP route is learned by the IPsec-VPN connection from the on-premises data center.
	//
	// example:
	//
	// VPN_BGP
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) GoString() string {
	return s.String()
}

func (s *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) GetRouteCount() *int32 {
	return s.RouteCount
}

func (s *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) GetRouteEntryType() *string {
	return s.RouteEntryType
}

func (s *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) GetSource() *string {
	return s.Source
}

func (s *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) SetRouteCount(v int32) *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts {
	s.RouteCount = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) SetRouteEntryType(v string) *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts {
	s.RouteEntryType = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) SetSource(v string) *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts {
	s.Source = &v
	return s
}

func (s *DescribeVcoRouteEntriesResponseBodyVpnRouteCounts) Validate() error {
	return dara.Validate(s)
}
