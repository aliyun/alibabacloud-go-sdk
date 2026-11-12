// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSharedDiskRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeSharedDiskRecordsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSharedDiskRecordsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeSharedDiskRecordsResponseBody
	GetRequestId() *string
	SetSharedDisks(v []*DescribeSharedDiskRecordsResponseBodySharedDisks) *DescribeSharedDiskRecordsResponseBody
	GetSharedDisks() []*DescribeSharedDiskRecordsResponseBodySharedDisks
	SetTotalCount(v int64) *DescribeSharedDiskRecordsResponseBody
	GetTotalCount() *int64
}

type DescribeSharedDiskRecordsResponseBody struct {
	// The current page number.
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
	// 68B85217-03B8-5141-9216-EA4D7C496B9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The shared disk billing records.
	SharedDisks []*DescribeSharedDiskRecordsResponseBodySharedDisks `json:"SharedDisks,omitempty" xml:"SharedDisks,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSharedDiskRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedDiskRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSharedDiskRecordsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSharedDiskRecordsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSharedDiskRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSharedDiskRecordsResponseBody) GetSharedDisks() []*DescribeSharedDiskRecordsResponseBodySharedDisks {
	return s.SharedDisks
}

func (s *DescribeSharedDiskRecordsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSharedDiskRecordsResponseBody) SetPageNumber(v int64) *DescribeSharedDiskRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBody) SetPageSize(v int64) *DescribeSharedDiskRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBody) SetRequestId(v string) *DescribeSharedDiskRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBody) SetSharedDisks(v []*DescribeSharedDiskRecordsResponseBodySharedDisks) *DescribeSharedDiskRecordsResponseBody {
	s.SharedDisks = v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBody) SetTotalCount(v int64) *DescribeSharedDiskRecordsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBody) Validate() error {
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

type DescribeSharedDiskRecordsResponseBodySharedDisks struct {
	// The creation time of the shared disk.
	//
	// example:
	//
	// 2026-06-01 10:00:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The shared disk ID.
	//
	// example:
	//
	// d-57kvvpuj1rk2ghumlgs6
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The disk size. Unit: GB.
	//
	// example:
	//
	// 10
	DiskSize *int64 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The ID of the instance from which the shared disk was last detached.
	//
	// example:
	//
	// standard
	LastDetachInstanceId *string `json:"LastDetachInstanceId,omitempty" xml:"LastDetachInstanceId,omitempty"`
	// The time when the shared disk was last detached.
	//
	// example:
	//
	// 2026-06-14 18:00:00
	LastDetachTime *string `json:"LastDetachTime,omitempty" xml:"LastDetachTime,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The SDG sharing ratio parameter, which supports cloud disks with different maximum sharing ratios.
	//
	// example:
	//
	// 100
	PerformanceLevel *int64 `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The reclamation time of the shared disk.
	//
	// example:
	//
	// 2026-06-15 14:30:00
	RecycleTime *string `json:"RecycleTime,omitempty" xml:"RecycleTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SDG ID of the shared disk.
	//
	// example:
	//
	// sdg-fqehhfdjv
	SdgId *string `json:"SdgId,omitempty" xml:"SdgId,omitempty"`
}

func (s DescribeSharedDiskRecordsResponseBodySharedDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeSharedDiskRecordsResponseBodySharedDisks) GoString() string {
	return s.String()
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetDiskSize() *int64 {
	return s.DiskSize
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetLastDetachInstanceId() *string {
	return s.LastDetachInstanceId
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetLastDetachTime() *string {
	return s.LastDetachTime
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetPerformanceLevel() *int64 {
	return s.PerformanceLevel
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetRecycleTime() *string {
	return s.RecycleTime
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) GetSdgId() *string {
	return s.SdgId
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetCreationTime(v string) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.CreationTime = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetDiskId(v string) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetDiskSize(v int64) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetLastDetachInstanceId(v string) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.LastDetachInstanceId = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetLastDetachTime(v string) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.LastDetachTime = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetNamespace(v string) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.Namespace = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetPerformanceLevel(v int64) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetRecycleTime(v string) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.RecycleTime = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetRegionId(v string) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.RegionId = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) SetSdgId(v string) *DescribeSharedDiskRecordsResponseBodySharedDisks {
	s.SdgId = &v
	return s
}

func (s *DescribeSharedDiskRecordsResponseBodySharedDisks) Validate() error {
	return dara.Validate(s)
}
