// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipSegmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipSegments(v *DescribeEipSegmentResponseBodyEipSegments) *DescribeEipSegmentResponseBody
	GetEipSegments() *DescribeEipSegmentResponseBodyEipSegments
	SetPageNumber(v int32) *DescribeEipSegmentResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEipSegmentResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEipSegmentResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEipSegmentResponseBody
	GetTotalCount() *int32
}

type DescribeEipSegmentResponseBody struct {
	// The details of the contiguous EIP group.
	EipSegments *DescribeEipSegmentResponseBodyEipSegments `json:"EipSegments,omitempty" xml:"EipSegments,omitempty" type:"Struct"`
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
	// F7A6301A-64BA-41EC-8284-8F4838C15D1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEipSegmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipSegmentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipSegmentResponseBody) GetEipSegments() *DescribeEipSegmentResponseBodyEipSegments {
	return s.EipSegments
}

func (s *DescribeEipSegmentResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEipSegmentResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEipSegmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEipSegmentResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEipSegmentResponseBody) SetEipSegments(v *DescribeEipSegmentResponseBodyEipSegments) *DescribeEipSegmentResponseBody {
	s.EipSegments = v
	return s
}

func (s *DescribeEipSegmentResponseBody) SetPageNumber(v int32) *DescribeEipSegmentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEipSegmentResponseBody) SetPageSize(v int32) *DescribeEipSegmentResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEipSegmentResponseBody) SetRequestId(v string) *DescribeEipSegmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEipSegmentResponseBody) SetTotalCount(v int32) *DescribeEipSegmentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEipSegmentResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEipSegmentResponseBodyEipSegments struct {
	EipSegment []*DescribeEipSegmentResponseBodyEipSegmentsEipSegment `json:"EipSegment,omitempty" xml:"EipSegment,omitempty" type:"Repeated"`
}

func (s DescribeEipSegmentResponseBodyEipSegments) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipSegmentResponseBodyEipSegments) GoString() string {
	return s.String()
}

func (s *DescribeEipSegmentResponseBodyEipSegments) GetEipSegment() []*DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	return s.EipSegment
}

func (s *DescribeEipSegmentResponseBodyEipSegments) SetEipSegment(v []*DescribeEipSegmentResponseBodyEipSegmentsEipSegment) *DescribeEipSegmentResponseBodyEipSegments {
	s.EipSegment = v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegments) Validate() error {
	return dara.Validate(s)
}

type DescribeEipSegmentResponseBodyEipSegmentsEipSegment struct {
	// The time when the contiguous EIP group was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-03-06T12:30:07Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the contiguous EIP group.
	//
	// example:
	//
	// MyEipSegment
	Descritpion *string `json:"Descritpion,omitempty" xml:"Descritpion,omitempty"`
	// The ID of the contiguous EIP group.
	//
	// example:
	//
	// eipsg-2zett8ba055tbsxme****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of EIPs in the contiguous EIP group.
	//
	// example:
	//
	// 16
	IpCount *string `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// The name of the contiguous EIP group.
	//
	// example:
	//
	// MyEipSegment
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region to which the contiguous EIP group belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The CIDR block and mask of the contiguous EIP group.
	//
	// example:
	//
	// 161.xx.xx.32/28
	Segment *string `json:"Segment,omitempty" xml:"Segment,omitempty"`
	// The status of the contiguous EIP group. Valid values:
	//
	// 	- **Allocating**
	//
	// 	- **Allocated**
	//
	// 	- **Releasing**
	//
	// example:
	//
	// Allocated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The zone of the contiguous EIP group.
	//
	// example:
	//
	// cn-hangzhou-a
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeEipSegmentResponseBodyEipSegmentsEipSegment) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GoString() string {
	return s.String()
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetDescritpion() *string {
	return s.Descritpion
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetIpCount() *string {
	return s.IpCount
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetName() *string {
	return s.Name
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetSegment() *string {
	return s.Segment
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetStatus() *string {
	return s.Status
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) GetZone() *string {
	return s.Zone
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetCreationTime(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.CreationTime = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetDescritpion(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.Descritpion = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetInstanceId(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.InstanceId = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetIpCount(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.IpCount = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetName(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.Name = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetRegionId(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.RegionId = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetSegment(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.Segment = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetStatus(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.Status = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) SetZone(v string) *DescribeEipSegmentResponseBodyEipSegmentsEipSegment {
	s.Zone = &v
	return s
}

func (s *DescribeEipSegmentResponseBodyEipSegmentsEipSegment) Validate() error {
	return dara.Validate(s)
}
