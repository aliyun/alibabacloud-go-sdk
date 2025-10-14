// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRouteEntryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeEnsRouteEntryListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEnsRouteEntryListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEnsRouteEntryListResponseBody
	GetRequestId() *string
	SetRouteEntrys(v []*DescribeEnsRouteEntryListResponseBodyRouteEntrys) *DescribeEnsRouteEntryListResponseBody
	GetRouteEntrys() []*DescribeEnsRouteEntryListResponseBodyRouteEntrys
	SetTotalCount(v int32) *DescribeEnsRouteEntryListResponseBody
	GetTotalCount() *int32
}

type DescribeEnsRouteEntryListResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the routes.
	RouteEntrys []*DescribeEnsRouteEntryListResponseBodyRouteEntrys `json:"RouteEntrys,omitempty" xml:"RouteEntrys,omitempty" type:"Repeated"`
	// The number of returned entries.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEnsRouteEntryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteEntryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteEntryListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEnsRouteEntryListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEnsRouteEntryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsRouteEntryListResponseBody) GetRouteEntrys() []*DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	return s.RouteEntrys
}

func (s *DescribeEnsRouteEntryListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEnsRouteEntryListResponseBody) SetPageNumber(v int32) *DescribeEnsRouteEntryListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBody) SetPageSize(v int32) *DescribeEnsRouteEntryListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBody) SetRequestId(v string) *DescribeEnsRouteEntryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBody) SetRouteEntrys(v []*DescribeEnsRouteEntryListResponseBodyRouteEntrys) *DescribeEnsRouteEntryListResponseBody {
	s.RouteEntrys = v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBody) SetTotalCount(v int32) *DescribeEnsRouteEntryListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBody) Validate() error {
	if s.RouteEntrys != nil {
		for _, item := range s.RouteEntrys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsRouteEntryListResponseBodyRouteEntrys struct {
	// The time when the IP address was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-16T03:50:05Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Enter a description for the route.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block of the route.
	//
	// example:
	//
	// 101.0.45.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The information about the next hop.
	NextHops []*DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops `json:"NextHops,omitempty" xml:"NextHops,omitempty" type:"Repeated"`
	// The ID of the route.
	//
	// example:
	//
	// rte-2zeksx7h436f5unk349m1
	RouteEntryId *string `json:"RouteEntryId,omitempty" xml:"RouteEntryId,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// test0
	RouteEntryName *string `json:"RouteEntryName,omitempty" xml:"RouteEntryName,omitempty"`
	// The ID of the route table.
	//
	// example:
	//
	// vtb-uf62p9o5cn35fi8xgurnm
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	// The source CIDR block. This field is used when you configure a route entry in the gateway route table. This field is not supported in the vSwitch route table.
	//
	// example:
	//
	// 10.XXX.XXX.0/24
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
	// The status of the route entry. Valid values:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the route entry.
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeEnsRouteEntryListResponseBodyRouteEntrys) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteEntryListResponseBodyRouteEntrys) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetDescription() *string {
	return s.Description
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetNextHops() []*DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops {
	return s.NextHops
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetRouteEntryId() *string {
	return s.RouteEntryId
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetRouteEntryName() *string {
	return s.RouteEntryName
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) GetType() *string {
	return s.Type
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetCreationTime(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.CreationTime = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetDescription(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.Description = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetDestinationCidrBlock(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetNextHops(v []*DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.NextHops = v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetRouteEntryId(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.RouteEntryId = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetRouteEntryName(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.RouteEntryName = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetRouteTableId(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.RouteTableId = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetSourceCidrBlock(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.SourceCidrBlock = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetStatus(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.Status = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) SetType(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrys {
	s.Type = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrys) Validate() error {
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

type DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops struct {
	// The ID of the next hop.
	//
	// example:
	//
	// i-bp1111yup9991890woxq
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The instance ID of the next hop.
	//
	// example:
	//
	// testInstance
	NextHopName *string `json:"NextHopName,omitempty" xml:"NextHopName,omitempty"`
	// The type of the next hop. Valid values:
	//
	// example:
	//
	// Instance
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
}

func (s DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) GoString() string {
	return s.String()
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) GetNextHopName() *string {
	return s.NextHopName
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) SetNextHopId(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops {
	s.NextHopId = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) SetNextHopName(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops {
	s.NextHopName = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) SetNextHopType(v string) *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops {
	s.NextHopType = &v
	return s
}

func (s *DescribeEnsRouteEntryListResponseBodyRouteEntrysNextHops) Validate() error {
	return dara.Validate(s)
}
