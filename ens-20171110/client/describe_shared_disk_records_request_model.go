// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSharedDiskRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskIds(v []*string) *DescribeSharedDiskRecordsRequest
	GetDiskIds() []*string
	SetEndTime(v string) *DescribeSharedDiskRecordsRequest
	GetEndTime() *string
	SetNamespaces(v []*string) *DescribeSharedDiskRecordsRequest
	GetNamespaces() []*string
	SetPageNumber(v int32) *DescribeSharedDiskRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSharedDiskRecordsRequest
	GetPageSize() *int32
	SetRegionIds(v []*string) *DescribeSharedDiskRecordsRequest
	GetRegionIds() []*string
	SetSdgIds(v []*string) *DescribeSharedDiskRecordsRequest
	GetSdgIds() []*string
	SetStartTime(v string) *DescribeSharedDiskRecordsRequest
	GetStartTime() *string
}

type DescribeSharedDiskRecordsRequest struct {
	// The shared disk IDs.
	DiskIds []*string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty" type:"Repeated"`
	// The billing end time.
	//
	// example:
	//
	// 2019-10-30T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The page number for the paging query.
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
	// The region IDs.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The SDG IDs.
	SdgIds []*string `json:"SdgIds,omitempty" xml:"SdgIds,omitempty" type:"Repeated"`
	// The billing start time.
	//
	// example:
	//
	// 2019-05-21T10:22:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSharedDiskRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedDiskRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSharedDiskRecordsRequest) GetDiskIds() []*string {
	return s.DiskIds
}

func (s *DescribeSharedDiskRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSharedDiskRecordsRequest) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *DescribeSharedDiskRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSharedDiskRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSharedDiskRecordsRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *DescribeSharedDiskRecordsRequest) GetSdgIds() []*string {
	return s.SdgIds
}

func (s *DescribeSharedDiskRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSharedDiskRecordsRequest) SetDiskIds(v []*string) *DescribeSharedDiskRecordsRequest {
	s.DiskIds = v
	return s
}

func (s *DescribeSharedDiskRecordsRequest) SetEndTime(v string) *DescribeSharedDiskRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSharedDiskRecordsRequest) SetNamespaces(v []*string) *DescribeSharedDiskRecordsRequest {
	s.Namespaces = v
	return s
}

func (s *DescribeSharedDiskRecordsRequest) SetPageNumber(v int32) *DescribeSharedDiskRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSharedDiskRecordsRequest) SetPageSize(v int32) *DescribeSharedDiskRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSharedDiskRecordsRequest) SetRegionIds(v []*string) *DescribeSharedDiskRecordsRequest {
	s.RegionIds = v
	return s
}

func (s *DescribeSharedDiskRecordsRequest) SetSdgIds(v []*string) *DescribeSharedDiskRecordsRequest {
	s.SdgIds = v
	return s
}

func (s *DescribeSharedDiskRecordsRequest) SetStartTime(v string) *DescribeSharedDiskRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSharedDiskRecordsRequest) Validate() error {
	return dara.Validate(s)
}
