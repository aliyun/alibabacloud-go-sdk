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
	if s.EipSegments != nil {
		if err := s.EipSegments.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.EipSegment != nil {
		for _, item := range s.EipSegment {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEipSegmentResponseBodyEipSegmentsEipSegment struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Descritpion  *string `json:"Descritpion,omitempty" xml:"Descritpion,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IpCount      *string `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Segment      *string `json:"Segment,omitempty" xml:"Segment,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Zone         *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
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
