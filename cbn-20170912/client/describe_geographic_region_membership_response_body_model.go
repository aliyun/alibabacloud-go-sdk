// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGeographicRegionMembershipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeGeographicRegionMembershipResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGeographicRegionMembershipResponseBody
	GetPageSize() *int32
	SetRegionIds(v *DescribeGeographicRegionMembershipResponseBodyRegionIds) *DescribeGeographicRegionMembershipResponseBody
	GetRegionIds() *DescribeGeographicRegionMembershipResponseBodyRegionIds
	SetRequestId(v string) *DescribeGeographicRegionMembershipResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGeographicRegionMembershipResponseBody
	GetTotalCount() *int32
}

type DescribeGeographicRegionMembershipResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of regions.
	RegionIds *DescribeGeographicRegionMembershipResponseBodyRegionIds `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DC9EB0C9-60AF-4A09-A36C-608F70130274
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGeographicRegionMembershipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGeographicRegionMembershipResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGeographicRegionMembershipResponseBody) GetRegionIds() *DescribeGeographicRegionMembershipResponseBodyRegionIds {
	return s.RegionIds
}

func (s *DescribeGeographicRegionMembershipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGeographicRegionMembershipResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetPageNumber(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetPageSize(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetRegionIds(v *DescribeGeographicRegionMembershipResponseBodyRegionIds) *DescribeGeographicRegionMembershipResponseBody {
	s.RegionIds = v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetRequestId(v string) *DescribeGeographicRegionMembershipResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) SetTotalCount(v int32) *DescribeGeographicRegionMembershipResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBody) Validate() error {
	if s.RegionIds != nil {
		if err := s.RegionIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGeographicRegionMembershipResponseBodyRegionIds struct {
	RegionId []*DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId `json:"RegionId,omitempty" xml:"RegionId,omitempty" type:"Repeated"`
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIds) GetRegionId() []*DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId {
	return s.RegionId
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIds) SetRegionId(v []*DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) *DescribeGeographicRegionMembershipResponseBodyRegionIds {
	s.RegionId = v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIds) Validate() error {
	if s.RegionId != nil {
		for _, item := range s.RegionId {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId struct {
	// The ID of the region.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// us-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) String() string {
	return dara.Prettify(s)
}

func (s DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) GoString() string {
	return s.String()
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) SetRegionId(v string) *DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId {
	s.RegionId = &v
	return s
}

func (s *DescribeGeographicRegionMembershipResponseBodyRegionIdsRegionId) Validate() error {
	return dara.Validate(s)
}
