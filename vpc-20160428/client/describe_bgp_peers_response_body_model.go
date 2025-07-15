// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpPeersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBgpPeers(v *DescribeBgpPeersResponseBodyBgpPeers) *DescribeBgpPeersResponseBody
	GetBgpPeers() *DescribeBgpPeersResponseBodyBgpPeers
	SetPageNumber(v int32) *DescribeBgpPeersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBgpPeersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBgpPeersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBgpPeersResponseBody
	GetTotalCount() *int32
}

type DescribeBgpPeersResponseBody struct {
	// The information about the BGP peer.
	BgpPeers *DescribeBgpPeersResponseBodyBgpPeers `json:"BgpPeers,omitempty" xml:"BgpPeers,omitempty" type:"Struct"`
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
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBgpPeersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPeersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBgpPeersResponseBody) GetBgpPeers() *DescribeBgpPeersResponseBodyBgpPeers {
	return s.BgpPeers
}

func (s *DescribeBgpPeersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBgpPeersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBgpPeersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBgpPeersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBgpPeersResponseBody) SetBgpPeers(v *DescribeBgpPeersResponseBodyBgpPeers) *DescribeBgpPeersResponseBody {
	s.BgpPeers = v
	return s
}

func (s *DescribeBgpPeersResponseBody) SetPageNumber(v int32) *DescribeBgpPeersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBgpPeersResponseBody) SetPageSize(v int32) *DescribeBgpPeersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBgpPeersResponseBody) SetRequestId(v string) *DescribeBgpPeersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBgpPeersResponseBody) SetTotalCount(v int32) *DescribeBgpPeersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBgpPeersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBgpPeersResponseBodyBgpPeers struct {
	BgpPeer []*DescribeBgpPeersResponseBodyBgpPeersBgpPeer `json:"BgpPeer,omitempty" xml:"BgpPeer,omitempty" type:"Repeated"`
}

func (s DescribeBgpPeersResponseBodyBgpPeers) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPeersResponseBodyBgpPeers) GoString() string {
	return s.String()
}

func (s *DescribeBgpPeersResponseBodyBgpPeers) GetBgpPeer() []*DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	return s.BgpPeer
}

func (s *DescribeBgpPeersResponseBodyBgpPeers) SetBgpPeer(v []*DescribeBgpPeersResponseBodyBgpPeersBgpPeer) *DescribeBgpPeersResponseBodyBgpPeers {
	s.BgpPeer = v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeers) Validate() error {
	return dara.Validate(s)
}

type DescribeBgpPeersResponseBodyBgpPeersBgpPeer struct {
	// The number of advertised routes.
	//
	// example:
	//
	// 0
	AdvertisedRouteCount *int32 `json:"AdvertisedRouteCount,omitempty" xml:"AdvertisedRouteCount,omitempty"`
	// The authentication key of the BGP group.
	//
	// example:
	//
	// !PWZ****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The Bidirectional Forwarding Detection (BFD) hop count.
	//
	// example:
	//
	// 3
	BfdMultiHop *int32 `json:"BfdMultiHop,omitempty" xml:"BfdMultiHop,omitempty"`
	// The ID of the BGP group.
	//
	// example:
	//
	// bgpg-2zev8h2wo414sfh****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// The ID of the BGP peer.
	//
	// example:
	//
	// bgp-2ze3un0ft1jd1xd****
	BgpPeerId *string `json:"BgpPeerId,omitempty" xml:"BgpPeerId,omitempty"`
	// The status of the BGP connection. Valid values:
	//
	// 	- **Idle**: The BGP connection is not used.
	//
	// 	- **Connect**: The BGP connection is used.
	//
	// 	- **Active**: The BGP connection is available.
	//
	// 	- **Established**: The BGP connection is established.
	//
	// 	- **Down**: The BGP connection is unavailable.
	//
	// example:
	//
	// Connect
	BgpStatus *string `json:"BgpStatus,omitempty" xml:"BgpStatus,omitempty"`
	// The description of the BGP group.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether BFD is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableBfd *bool `json:"EnableBfd,omitempty" xml:"EnableBfd,omitempty"`
	// The time when the BGP peer is modified.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-05-20T07:29:08Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The hold time.
	//
	// example:
	//
	// 30
	Hold *string `json:"Hold,omitempty" xml:"Hold,omitempty"`
	// The version of the IP address.
	//
	// example:
	//
	// ipv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether a fake autonomous system number (ASN) is used. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsFake *bool `json:"IsFake,omitempty" xml:"IsFake,omitempty"`
	// The Keepalive interval.
	//
	// example:
	//
	// 10
	Keepalive *string `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The ASN of the device on the Alibaba Cloud side.
	//
	// example:
	//
	// 45104
	LocalAsn *string `json:"LocalAsn,omitempty" xml:"LocalAsn,omitempty"`
	// The name of the BGP peer.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The autonomous system number (ASN)of the BGP peer.
	//
	// example:
	//
	// 65515
	PeerAsn *string `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The IP address of the BGP peer.
	//
	// example:
	//
	// 116.62.XX.XX
	PeerIpAddress *string `json:"PeerIpAddress,omitempty" xml:"PeerIpAddress,omitempty"`
	// The number of received routes.
	//
	// example:
	//
	// 0
	ReceivedRouteCount *int32 `json:"ReceivedRouteCount,omitempty" xml:"ReceivedRouteCount,omitempty"`
	// The ID of the region to which the BGP group belongs.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum number of routes.
	//
	// example:
	//
	// 110
	RouteLimit *string `json:"RouteLimit,omitempty" xml:"RouteLimit,omitempty"`
	// The Router ID.
	//
	// example:
	//
	// vbr-2zecmmvg5gvu8i4te****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The status of the BGP peer. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBgpPeersResponseBodyBgpPeersBgpPeer) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GoString() string {
	return s.String()
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetAdvertisedRouteCount() *int32 {
	return s.AdvertisedRouteCount
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetBfdMultiHop() *int32 {
	return s.BfdMultiHop
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetBgpPeerId() *string {
	return s.BgpPeerId
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetBgpStatus() *string {
	return s.BgpStatus
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetDescription() *string {
	return s.Description
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetEnableBfd() *bool {
	return s.EnableBfd
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetHold() *string {
	return s.Hold
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetIsFake() *bool {
	return s.IsFake
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetKeepalive() *string {
	return s.Keepalive
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetLocalAsn() *string {
	return s.LocalAsn
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetName() *string {
	return s.Name
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetPeerAsn() *string {
	return s.PeerAsn
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetPeerIpAddress() *string {
	return s.PeerIpAddress
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetReceivedRouteCount() *int32 {
	return s.ReceivedRouteCount
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetRouteLimit() *string {
	return s.RouteLimit
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) GetStatus() *string {
	return s.Status
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetAdvertisedRouteCount(v int32) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.AdvertisedRouteCount = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetAuthKey(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.AuthKey = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetBfdMultiHop(v int32) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.BfdMultiHop = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetBgpGroupId(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.BgpGroupId = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetBgpPeerId(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.BgpPeerId = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetBgpStatus(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.BgpStatus = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetDescription(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.Description = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetEnableBfd(v bool) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.EnableBfd = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetGmtModified(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.GmtModified = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetHold(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.Hold = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetIpVersion(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.IpVersion = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetIsFake(v bool) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.IsFake = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetKeepalive(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.Keepalive = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetLocalAsn(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.LocalAsn = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetName(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.Name = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetPeerAsn(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.PeerAsn = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetPeerIpAddress(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.PeerIpAddress = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetReceivedRouteCount(v int32) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.ReceivedRouteCount = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetRegionId(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.RegionId = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetRouteLimit(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.RouteLimit = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetRouterId(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.RouterId = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) SetStatus(v string) *DescribeBgpPeersResponseBodyBgpPeersBgpPeer {
	s.Status = &v
	return s
}

func (s *DescribeBgpPeersResponseBodyBgpPeersBgpPeer) Validate() error {
	return dara.Validate(s)
}
