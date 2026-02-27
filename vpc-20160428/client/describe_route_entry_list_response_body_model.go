// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteEntryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeRouteEntryListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeRouteEntryListResponseBody
	GetRequestId() *string
	SetRouteEntrys(v *DescribeRouteEntryListResponseBodyRouteEntrys) *DescribeRouteEntryListResponseBody
	GetRouteEntrys() *DescribeRouteEntryListResponseBodyRouteEntrys
}

type DescribeRouteEntryListResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If a value is returned for **NextToken**, the value is used to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteEntrys *DescribeRouteEntryListResponseBodyRouteEntrys `json:"RouteEntrys,omitempty" xml:"RouteEntrys,omitempty" type:"Struct"`
}

func (s DescribeRouteEntryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRouteEntryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteEntryListResponseBody) GetRouteEntrys() *DescribeRouteEntryListResponseBodyRouteEntrys {
	return s.RouteEntrys
}

func (s *DescribeRouteEntryListResponseBody) SetNextToken(v string) *DescribeRouteEntryListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRouteEntryListResponseBody) SetRequestId(v string) *DescribeRouteEntryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBody) SetRouteEntrys(v *DescribeRouteEntryListResponseBodyRouteEntrys) *DescribeRouteEntryListResponseBody {
	s.RouteEntrys = v
	return s
}

func (s *DescribeRouteEntryListResponseBody) Validate() error {
	if s.RouteEntrys != nil {
		if err := s.RouteEntrys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteEntryListResponseBodyRouteEntrys struct {
	RouteEntry []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry `json:"RouteEntry,omitempty" xml:"RouteEntry,omitempty" type:"Repeated"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrys) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrys) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrys) GetRouteEntry() []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	return s.RouteEntry
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrys) SetRouteEntry(v []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) *DescribeRouteEntryListResponseBodyRouteEntrys {
	s.RouteEntry = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrys) Validate() error {
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

type DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	GmtModified    *string                                                          `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IpVersion      *string                                                          `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	NextHops       *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops `json:"NextHops,omitempty" xml:"NextHops,omitempty" type:"Struct"`
	Origin         *string                                                          `json:"Origin,omitempty" xml:"Origin,omitempty"`
	RouteEntryId   *string                                                          `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	RouteEntryName *string                                                          `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	RouteTableId   *string                                                          `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	ServiceType    *string                                                          `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status         *string                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetNextHops() *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops {
	return s.NextHops
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) GetType() *string {
	return s.Type
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetDescription(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.Description = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetDestinationCidrBlock(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetGmtModified(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.GmtModified = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetIpVersion(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.IpVersion = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetNextHops(v *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.NextHops = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetOrigin(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.Origin = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetRouteEntryId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.RouteEntryId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetRouteEntryName(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.RouteEntryName = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetRouteTableId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetServiceType(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.ServiceType = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetStatus(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.Status = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) SetType(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry {
	s.Type = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntry) Validate() error {
	if s.NextHops != nil {
		if err := s.NextHops.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops struct {
	NextHop []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop `json:"NextHop,omitempty" xml:"NextHop,omitempty" type:"Repeated"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) GetNextHop() []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	return s.NextHop
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) SetNextHop(v []*DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops {
	s.NextHop = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHops) Validate() error {
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

type DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop struct {
	Enabled            *int32                                                                                    `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	NextHopId          *string                                                                                   `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopRegionId    *string                                                                                   `json:"NextHopRegionId,omitempty" xml:"NextHopRegionId,omitempty"`
	NextHopRelatedInfo *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo `json:"NextHopRelatedInfo,omitempty" xml:"NextHopRelatedInfo,omitempty" type:"Struct"`
	NextHopType        *string                                                                                   `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	Weight             *int32                                                                                    `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetEnabled() *int32 {
	return s.Enabled
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetNextHopRegionId() *string {
	return s.NextHopRegionId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetNextHopRelatedInfo() *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo {
	return s.NextHopRelatedInfo
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetEnabled(v int32) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.Enabled = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetNextHopId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetNextHopRegionId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopRegionId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetNextHopRelatedInfo(v *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopRelatedInfo = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetNextHopType(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.NextHopType = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) SetWeight(v int32) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop {
	s.Weight = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHop) Validate() error {
	if s.NextHopRelatedInfo != nil {
		if err := s.NextHopRelatedInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) SetInstanceId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) SetInstanceType(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo {
	s.InstanceType = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) SetRegionId(v string) *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntrysRouteEntryNextHopsNextHopNextHopRelatedInfo) Validate() error {
	return dara.Validate(s)
}
