// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteEntryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeRouteEntryListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRouteEntryListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeRouteEntryListResponseBody
	GetRequestId() *string
	SetRouteEntries(v []*DescribeRouteEntryListResponseBodyRouteEntries) *DescribeRouteEntryListResponseBody
	GetRouteEntries() []*DescribeRouteEntryListResponseBodyRouteEntries
}

type DescribeRouteEntryListResponseBody struct {
	MaxResults   *int32                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteEntries []*DescribeRouteEntryListResponseBodyRouteEntries `json:"RouteEntries,omitempty" xml:"RouteEntries,omitempty" type:"Repeated"`
}

func (s DescribeRouteEntryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRouteEntryListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRouteEntryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteEntryListResponseBody) GetRouteEntries() []*DescribeRouteEntryListResponseBodyRouteEntries {
	return s.RouteEntries
}

func (s *DescribeRouteEntryListResponseBody) SetMaxResults(v int32) *DescribeRouteEntryListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeRouteEntryListResponseBody) SetNextToken(v string) *DescribeRouteEntryListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRouteEntryListResponseBody) SetRequestId(v string) *DescribeRouteEntryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBody) SetRouteEntries(v []*DescribeRouteEntryListResponseBodyRouteEntries) *DescribeRouteEntryListResponseBody {
	s.RouteEntries = v
	return s
}

func (s *DescribeRouteEntryListResponseBody) Validate() error {
	if s.RouteEntries != nil {
		for _, item := range s.RouteEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteEntryListResponseBodyRouteEntries struct {
	Description          *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock *string                                                   `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	IpVersion            *string                                                   `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	NextHops             []*DescribeRouteEntryListResponseBodyRouteEntriesNextHops `json:"NextHops,omitempty" xml:"NextHops,omitempty" type:"Repeated"`
	RouteEntryId         *string                                                   `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	RouteEntryName       *string                                                   `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	RouteTableId         *string                                                   `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	Status               *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                 *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntries) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetNextHops() []*DescribeRouteEntryListResponseBodyRouteEntriesNextHops {
	return s.NextHops
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) GetType() *string {
	return s.Type
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetDescription(v string) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.Description = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetDestinationCidrBlock(v string) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetIpVersion(v string) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.IpVersion = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetNextHops(v []*DescribeRouteEntryListResponseBodyRouteEntriesNextHops) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.NextHops = v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetRouteEntryId(v string) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.RouteEntryId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetRouteEntryName(v string) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.RouteEntryName = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetRouteTableId(v string) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetStatus(v string) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.Status = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) SetType(v string) *DescribeRouteEntryListResponseBodyRouteEntries {
	s.Type = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntries) Validate() error {
	if s.NextHops != nil {
		for _, item := range s.NextHops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteEntryListResponseBodyRouteEntriesNextHops struct {
	NextHopId   *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
}

func (s DescribeRouteEntryListResponseBodyRouteEntriesNextHops) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteEntryListResponseBodyRouteEntriesNextHops) GoString() string {
	return s.String()
}

func (s *DescribeRouteEntryListResponseBodyRouteEntriesNextHops) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeRouteEntryListResponseBodyRouteEntriesNextHops) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeRouteEntryListResponseBodyRouteEntriesNextHops) SetNextHopId(v string) *DescribeRouteEntryListResponseBodyRouteEntriesNextHops {
	s.NextHopId = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntriesNextHops) SetNextHopType(v string) *DescribeRouteEntryListResponseBodyRouteEntriesNextHops {
	s.NextHopType = &v
	return s
}

func (s *DescribeRouteEntryListResponseBodyRouteEntriesNextHops) Validate() error {
	return dara.Validate(s)
}
