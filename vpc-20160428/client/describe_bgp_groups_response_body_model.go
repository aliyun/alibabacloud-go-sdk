// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBgpGroups(v *DescribeBgpGroupsResponseBodyBgpGroups) *DescribeBgpGroupsResponseBody
	GetBgpGroups() *DescribeBgpGroupsResponseBodyBgpGroups
	SetPageNumber(v int32) *DescribeBgpGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBgpGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBgpGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBgpGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeBgpGroupsResponseBody struct {
	// The detailed information about the BGP group.
	BgpGroups *DescribeBgpGroupsResponseBodyBgpGroups `json:"BgpGroups,omitempty" xml:"BgpGroups,omitempty" type:"Struct"`
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
	// 1D0971B2-A35A-42C1-A44C-E91360C36C0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBgpGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBgpGroupsResponseBody) GetBgpGroups() *DescribeBgpGroupsResponseBodyBgpGroups {
	return s.BgpGroups
}

func (s *DescribeBgpGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBgpGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBgpGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBgpGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBgpGroupsResponseBody) SetBgpGroups(v *DescribeBgpGroupsResponseBodyBgpGroups) *DescribeBgpGroupsResponseBody {
	s.BgpGroups = v
	return s
}

func (s *DescribeBgpGroupsResponseBody) SetPageNumber(v int32) *DescribeBgpGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBgpGroupsResponseBody) SetPageSize(v int32) *DescribeBgpGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBgpGroupsResponseBody) SetRequestId(v string) *DescribeBgpGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBgpGroupsResponseBody) SetTotalCount(v int32) *DescribeBgpGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBgpGroupsResponseBody) Validate() error {
	if s.BgpGroups != nil {
		if err := s.BgpGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBgpGroupsResponseBodyBgpGroups struct {
	BgpGroup []*DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup `json:"BgpGroup,omitempty" xml:"BgpGroup,omitempty" type:"Repeated"`
}

func (s DescribeBgpGroupsResponseBodyBgpGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpGroupsResponseBodyBgpGroups) GoString() string {
	return s.String()
}

func (s *DescribeBgpGroupsResponseBodyBgpGroups) GetBgpGroup() []*DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	return s.BgpGroup
}

func (s *DescribeBgpGroupsResponseBodyBgpGroups) SetBgpGroup(v []*DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) *DescribeBgpGroupsResponseBodyBgpGroups {
	s.BgpGroup = v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroups) Validate() error {
	if s.BgpGroup != nil {
		for _, item := range s.BgpGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup struct {
	// The key used by the BGP group.
	//
	// example:
	//
	// !PWZ****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The ID of the BGP group.
	//
	// example:
	//
	// bgpg-bp1k25cyp26cllath****
	BgpGroupId *string `json:"BgpGroupId,omitempty" xml:"BgpGroupId,omitempty"`
	// The description of the BGP group.
	//
	// example:
	//
	// The description of the BGP group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hold time to receive BGP messages. Unit: seconds.
	//
	// >  If no message is received within the hold time, the BGP peer is considered disconnected.
	//
	// example:
	//
	// 30
	Hold *string `json:"Hold,omitempty" xml:"Hold,omitempty"`
	// The IP version of the BGP group. Valid values:
	//
	// 	- **ipv4**: IPv4
	//
	// 	- **ipv6**: IPv6. IPv6 is supported only if the VBR of the BGP group has IPv6 enabled.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// Indicates whether the ASN is fake. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	IsFake *string `json:"IsFake,omitempty" xml:"IsFake,omitempty"`
	// The keepalive time. Unit: seconds.
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
	// The name of the BGP group.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The autonomous system number (ASN) of the on-premises device in the data center.
	//
	// example:
	//
	// 1****
	PeerAsn *string `json:"PeerAsn,omitempty" xml:"PeerAsn,omitempty"`
	// The ID of the region to which the BGP group belongs.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum number of route entries for BGP dynamic route learning.
	//
	// example:
	//
	// 99
	RouteLimit *string `json:"RouteLimit,omitempty" xml:"RouteLimit,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vbr-bp1ctxy813985gkuk****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The status of the BGP group.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GoString() string {
	return s.String()
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetBgpGroupId() *string {
	return s.BgpGroupId
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetDescription() *string {
	return s.Description
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetHold() *string {
	return s.Hold
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetIsFake() *string {
	return s.IsFake
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetKeepalive() *string {
	return s.Keepalive
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetLocalAsn() *string {
	return s.LocalAsn
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetName() *string {
	return s.Name
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetPeerAsn() *string {
	return s.PeerAsn
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetRouteLimit() *string {
	return s.RouteLimit
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) GetStatus() *string {
	return s.Status
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetAuthKey(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.AuthKey = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetBgpGroupId(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.BgpGroupId = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetDescription(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.Description = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetHold(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.Hold = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetIpVersion(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.IpVersion = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetIsFake(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.IsFake = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetKeepalive(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.Keepalive = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetLocalAsn(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.LocalAsn = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetName(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.Name = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetPeerAsn(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.PeerAsn = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetRegionId(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.RegionId = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetRouteLimit(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.RouteLimit = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetRouterId(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.RouterId = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) SetStatus(v string) *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup {
	s.Status = &v
	return s
}

func (s *DescribeBgpGroupsResponseBodyBgpGroupsBgpGroup) Validate() error {
	return dara.Validate(s)
}
