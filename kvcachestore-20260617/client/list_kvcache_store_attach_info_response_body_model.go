// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAttachInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachInfos(v []*ListKVCacheStoreAttachInfoResponseBodyAttachInfos) *ListKVCacheStoreAttachInfoResponseBody
	GetAttachInfos() []*ListKVCacheStoreAttachInfoResponseBodyAttachInfos
	SetNextToken(v string) *ListKVCacheStoreAttachInfoResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListKVCacheStoreAttachInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKVCacheStoreAttachInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListKVCacheStoreAttachInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListKVCacheStoreAttachInfoResponseBody
	GetTotalCount() *int64
}

type ListKVCacheStoreAttachInfoResponseBody struct {
	// The list of mount information.
	AttachInfos []*ListKVCacheStoreAttachInfoResponseBodyAttachInfos `json:"AttachInfos,omitempty" xml:"AttachInfos,omitempty" type:"Repeated"`
	// The pagination token used to query the next batch of data.
	//
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3NUkN7qf+fcWj7joK8M6tU
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID. A request ID is returned regardless of whether the call is successful.
	//
	// example:
	//
	// B127704C-ECB1-5B0A-AA9C-8F394A6F179F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned for the paged query.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKVCacheStoreAttachInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAttachInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAttachInfoResponseBody) GetAttachInfos() []*ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	return s.AttachInfos
}

func (s *ListKVCacheStoreAttachInfoResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKVCacheStoreAttachInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKVCacheStoreAttachInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKVCacheStoreAttachInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKVCacheStoreAttachInfoResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListKVCacheStoreAttachInfoResponseBody) SetAttachInfos(v []*ListKVCacheStoreAttachInfoResponseBodyAttachInfos) *ListKVCacheStoreAttachInfoResponseBody {
	s.AttachInfos = v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBody) SetNextToken(v string) *ListKVCacheStoreAttachInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBody) SetPageNumber(v int32) *ListKVCacheStoreAttachInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBody) SetPageSize(v int32) *ListKVCacheStoreAttachInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBody) SetRequestId(v string) *ListKVCacheStoreAttachInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBody) SetTotalCount(v int64) *ListKVCacheStoreAttachInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBody) Validate() error {
	if s.AttachInfos != nil {
		for _, item := range s.AttachInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKVCacheStoreAttachInfoResponseBodyAttachInfos struct {
	// The time of the most recent attach operation, in ISO 8601 format. The value is null if the instance has not been attached.
	//
	// example:
	//
	// 2026-06-20T08:30:00Z
	AttachedAt *string `json:"AttachedAt,omitempty" xml:"AttachedAt,omitempty"`
	// The file system capacity, in GiB.
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// KVCacheStore KvcsId
	//
	// example:
	//
	// kvcs-xxxxx
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The mount point ID at the file system level.
	//
	// example:
	//
	// mp-xxxxx
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// The region where the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The attach status. Valid values:
	//
	// - Attaching: The instance is being mounted.
	//
	// - Attached: The instance is mounted.
	//
	// - Detaching: The instance is being unmounted.
	//
	// After unmounting is complete, the record is deleted and not returned.
	//
	// example:
	//
	// ATTACHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance type. Valid values:
	//
	// - kvcs: KVCacheStore (CPFS).
	//
	// example:
	//
	// preview
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The VSC ID on the compute side.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// The zone where the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListKVCacheStoreAttachInfoResponseBodyAttachInfos) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetAttachedAt() *string {
	return s.AttachedAt
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetCapacity() *int64 {
	return s.Capacity
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetKvcsId() *string {
	return s.KvcsId
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetMountPointId() *string {
	return s.MountPointId
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetStatus() *string {
	return s.Status
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetType() *string {
	return s.Type
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetVscId() *string {
	return s.VscId
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetAttachedAt(v string) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.AttachedAt = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetCapacity(v int64) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.Capacity = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetKvcsId(v string) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.KvcsId = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetMountPointId(v string) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.MountPointId = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetRegionId(v string) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.RegionId = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetStatus(v string) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.Status = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetType(v string) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.Type = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetVscId(v string) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.VscId = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) SetZoneId(v string) *ListKVCacheStoreAttachInfoResponseBodyAttachInfos {
	s.ZoneId = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponseBodyAttachInfos) Validate() error {
	return dara.Validate(s)
}
