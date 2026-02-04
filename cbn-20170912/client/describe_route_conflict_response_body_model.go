// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteConflictResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRouteConflictResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteConflictResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRouteConflictResponseBody
	GetRequestId() *string
	SetRouteConflicts(v *DescribeRouteConflictResponseBodyRouteConflicts) *DescribeRouteConflictResponseBody
	GetRouteConflicts() *DescribeRouteConflictResponseBodyRouteConflicts
	SetTotalCount(v int32) *DescribeRouteConflictResponseBody
	GetTotalCount() *int32
}

type DescribeRouteConflictResponseBody struct {
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
	// EE3A2CC7-41F1-58DB-8306-CFC99D9C747B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of overlapping routes.
	RouteConflicts *DescribeRouteConflictResponseBodyRouteConflicts `json:"RouteConflicts,omitempty" xml:"RouteConflicts,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouteConflictResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteConflictResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteConflictResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteConflictResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouteConflictResponseBody) GetRouteConflicts() *DescribeRouteConflictResponseBodyRouteConflicts {
	return s.RouteConflicts
}

func (s *DescribeRouteConflictResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRouteConflictResponseBody) SetPageNumber(v int32) *DescribeRouteConflictResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetPageSize(v int32) *DescribeRouteConflictResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetRequestId(v string) *DescribeRouteConflictResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetRouteConflicts(v *DescribeRouteConflictResponseBodyRouteConflicts) *DescribeRouteConflictResponseBody {
	s.RouteConflicts = v
	return s
}

func (s *DescribeRouteConflictResponseBody) SetTotalCount(v int32) *DescribeRouteConflictResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouteConflictResponseBody) Validate() error {
	if s.RouteConflicts != nil {
		if err := s.RouteConflicts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouteConflictResponseBodyRouteConflicts struct {
	RouteConflict []*DescribeRouteConflictResponseBodyRouteConflictsRouteConflict `json:"RouteConflict,omitempty" xml:"RouteConflict,omitempty" type:"Repeated"`
}

func (s DescribeRouteConflictResponseBodyRouteConflicts) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteConflictResponseBodyRouteConflicts) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) GetRouteConflict() []*DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	return s.RouteConflict
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) SetRouteConflict(v []*DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) *DescribeRouteConflictResponseBodyRouteConflicts {
	s.RouteConflict = v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflicts) Validate() error {
	if s.RouteConflict != nil {
		for _, item := range s.RouteConflict {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouteConflictResponseBodyRouteConflictsRouteConflict struct {
	// The destination CIDR block of the overlapping route.
	//
	// example:
	//
	// 172.16.0.0/16
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the peer network instance on which the overlapping routes are found.
	//
	// example:
	//
	// ccn-0q3b7oviikmm9h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the peer network instance on which the overlapping routes are found.
	//
	// 	- **VPC**: VPC
	//
	// 	- **VBR**: VBR
	//
	// 	- **CCN**: CCN instance
	//
	// example:
	//
	// CCN
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The region ID of the peer network instance on which the overlapping routes are found is deployed.
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

func (s DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) GoString() string {
	return s.String()
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetDestinationCidrBlock(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetInstanceId(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.InstanceId = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetInstanceType(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.InstanceType = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetRegionId(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) SetStatus(v string) *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict {
	s.Status = &v
	return s
}

func (s *DescribeRouteConflictResponseBodyRouteConflictsRouteConflict) Validate() error {
	return dara.Validate(s)
}
