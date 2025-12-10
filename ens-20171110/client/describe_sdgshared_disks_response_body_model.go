// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGSharedDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeSDGSharedDisksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSDGSharedDisksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeSDGSharedDisksResponseBody
	GetRequestId() *string
	SetSharedDisks(v []*DescribeSDGSharedDisksResponseBodySharedDisks) *DescribeSDGSharedDisksResponseBody
	GetSharedDisks() []*DescribeSDGSharedDisksResponseBodySharedDisks
	SetTotalCount(v int64) *DescribeSDGSharedDisksResponseBody
	GetTotalCount() *int64
}

type DescribeSDGSharedDisksResponseBody struct {
	// Current page number when paging
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxsxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Shared disk list
	SharedDisks []*DescribeSDGSharedDisksResponseBodySharedDisks `json:"SharedDisks,omitempty" xml:"SharedDisks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSDGSharedDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGSharedDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDGSharedDisksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSDGSharedDisksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSDGSharedDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSDGSharedDisksResponseBody) GetSharedDisks() []*DescribeSDGSharedDisksResponseBodySharedDisks {
	return s.SharedDisks
}

func (s *DescribeSDGSharedDisksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSDGSharedDisksResponseBody) SetPageNumber(v int64) *DescribeSDGSharedDisksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBody) SetPageSize(v int64) *DescribeSDGSharedDisksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBody) SetRequestId(v string) *DescribeSDGSharedDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBody) SetSharedDisks(v []*DescribeSDGSharedDisksResponseBodySharedDisks) *DescribeSDGSharedDisksResponseBody {
	s.SharedDisks = v
	return s
}

func (s *DescribeSDGSharedDisksResponseBody) SetTotalCount(v int64) *DescribeSDGSharedDisksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBody) Validate() error {
	if s.SharedDisks != nil {
		for _, item := range s.SharedDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSDGSharedDisksResponseBodySharedDisks struct {
	// The time when the shared disk was created.
	//
	// example:
	//
	// 2025-10-09T15:13:21+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// shared disk id
	//
	// example:
	//
	// d-57kvvpuj1rk2ghumlgs6
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Shared disk type
	//
	// example:
	//
	// standard
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The namespace of the service.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The node ID.
	//
	// example:
	//
	// cn-hangzhou-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// SdgId of the shared disk
	//
	// example:
	//
	// sdg-fqehhfdjv
	SdgId *string `json:"SdgId,omitempty" xml:"SdgId,omitempty"`
	// Number of shared mounts
	//
	// example:
	//
	// 10
	SharedNum *int32 `json:"SharedNum,omitempty" xml:"SharedNum,omitempty"`
	// Shared disk status
	//
	// example:
	//
	// avaliable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSDGSharedDisksResponseBodySharedDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGSharedDisksResponseBodySharedDisks) GoString() string {
	return s.String()
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) GetSdgId() *string {
	return s.SdgId
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) GetSharedNum() *int32 {
	return s.SharedNum
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) SetCreationTime(v string) *DescribeSDGSharedDisksResponseBodySharedDisks {
	s.CreationTime = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) SetDiskId(v string) *DescribeSDGSharedDisksResponseBodySharedDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) SetDiskType(v string) *DescribeSDGSharedDisksResponseBodySharedDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) SetNamespace(v string) *DescribeSDGSharedDisksResponseBodySharedDisks {
	s.Namespace = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) SetRegionId(v string) *DescribeSDGSharedDisksResponseBodySharedDisks {
	s.RegionId = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) SetSdgId(v string) *DescribeSDGSharedDisksResponseBodySharedDisks {
	s.SdgId = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) SetSharedNum(v int32) *DescribeSDGSharedDisksResponseBodySharedDisks {
	s.SharedNum = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) SetStatus(v string) *DescribeSDGSharedDisksResponseBodySharedDisks {
	s.Status = &v
	return s
}

func (s *DescribeSDGSharedDisksResponseBodySharedDisks) Validate() error {
	return dara.Validate(s)
}
