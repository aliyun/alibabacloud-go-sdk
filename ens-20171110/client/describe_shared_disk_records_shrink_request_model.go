// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSharedDiskRecordsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskIdsShrink(v string) *DescribeSharedDiskRecordsShrinkRequest
	GetDiskIdsShrink() *string
	SetEndTime(v string) *DescribeSharedDiskRecordsShrinkRequest
	GetEndTime() *string
	SetNamespacesShrink(v string) *DescribeSharedDiskRecordsShrinkRequest
	GetNamespacesShrink() *string
	SetPageNumber(v int32) *DescribeSharedDiskRecordsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSharedDiskRecordsShrinkRequest
	GetPageSize() *int32
	SetRegionIdsShrink(v string) *DescribeSharedDiskRecordsShrinkRequest
	GetRegionIdsShrink() *string
	SetSdgIdsShrink(v string) *DescribeSharedDiskRecordsShrinkRequest
	GetSdgIdsShrink() *string
	SetStartTime(v string) *DescribeSharedDiskRecordsShrinkRequest
	GetStartTime() *string
}

type DescribeSharedDiskRecordsShrinkRequest struct {
	// The shared disk IDs.
	DiskIdsShrink *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The billing end time.
	//
	// example:
	//
	// 2019-10-30T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The namespaces.
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
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
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	// The SDG IDs.
	SdgIdsShrink *string `json:"SdgIds,omitempty" xml:"SdgIds,omitempty"`
	// The billing start time.
	//
	// example:
	//
	// 2019-05-21T10:22:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSharedDiskRecordsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedDiskRecordsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSharedDiskRecordsShrinkRequest) GetDiskIdsShrink() *string {
	return s.DiskIdsShrink
}

func (s *DescribeSharedDiskRecordsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSharedDiskRecordsShrinkRequest) GetNamespacesShrink() *string {
	return s.NamespacesShrink
}

func (s *DescribeSharedDiskRecordsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSharedDiskRecordsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSharedDiskRecordsShrinkRequest) GetRegionIdsShrink() *string {
	return s.RegionIdsShrink
}

func (s *DescribeSharedDiskRecordsShrinkRequest) GetSdgIdsShrink() *string {
	return s.SdgIdsShrink
}

func (s *DescribeSharedDiskRecordsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSharedDiskRecordsShrinkRequest) SetDiskIdsShrink(v string) *DescribeSharedDiskRecordsShrinkRequest {
	s.DiskIdsShrink = &v
	return s
}

func (s *DescribeSharedDiskRecordsShrinkRequest) SetEndTime(v string) *DescribeSharedDiskRecordsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSharedDiskRecordsShrinkRequest) SetNamespacesShrink(v string) *DescribeSharedDiskRecordsShrinkRequest {
	s.NamespacesShrink = &v
	return s
}

func (s *DescribeSharedDiskRecordsShrinkRequest) SetPageNumber(v int32) *DescribeSharedDiskRecordsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSharedDiskRecordsShrinkRequest) SetPageSize(v int32) *DescribeSharedDiskRecordsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSharedDiskRecordsShrinkRequest) SetRegionIdsShrink(v string) *DescribeSharedDiskRecordsShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *DescribeSharedDiskRecordsShrinkRequest) SetSdgIdsShrink(v string) *DescribeSharedDiskRecordsShrinkRequest {
	s.SdgIdsShrink = &v
	return s
}

func (s *DescribeSharedDiskRecordsShrinkRequest) SetStartTime(v string) *DescribeSharedDiskRecordsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSharedDiskRecordsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
