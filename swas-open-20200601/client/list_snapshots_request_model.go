// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *ListSnapshotsRequest
	GetDiskId() *string
	SetInstanceId(v string) *ListSnapshotsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListSnapshotsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSnapshotsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListSnapshotsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListSnapshotsRequest
	GetResourceGroupId() *string
	SetSnapshotIds(v string) *ListSnapshotsRequest
	GetSnapshotIds() *string
	SetSourceDiskType(v string) *ListSnapshotsRequest
	GetSourceDiskType() *string
	SetTag(v []*ListSnapshotsRequestTag) *ListSnapshotsRequest
	GetTag() []*ListSnapshotsRequestTag
}

type ListSnapshotsRequest struct {
	// The ID of the disk.
	//
	// example:
	//
	// d-bp14wq0149cpp2xy****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The instance ID of the simple application server.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the snapshot list.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server to which the snapshot belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot IDs. The value can be a JSON array that consists of up to 100 snapshot IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// ["s-bp16oazlsold4dks****", "s-bp16oazlsold4abc****"]
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	// The type of the source cloud disk for the snapshot. Valid values:
	//
	// - System: system cloud disk.
	//
	// - Data: data cloud disk.
	//
	// example:
	//
	// Data
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The tags of the snapshot.
	Tag []*ListSnapshotsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ListSnapshotsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSnapshotsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSnapshotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSnapshotsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSnapshotsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSnapshotsRequest) GetSnapshotIds() *string {
	return s.SnapshotIds
}

func (s *ListSnapshotsRequest) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *ListSnapshotsRequest) GetTag() []*ListSnapshotsRequestTag {
	return s.Tag
}

func (s *ListSnapshotsRequest) SetDiskId(v string) *ListSnapshotsRequest {
	s.DiskId = &v
	return s
}

func (s *ListSnapshotsRequest) SetInstanceId(v string) *ListSnapshotsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageNumber(v int32) *ListSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageSize(v int32) *ListSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsRequest) SetRegionId(v string) *ListSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSnapshotsRequest) SetResourceGroupId(v string) *ListSnapshotsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSnapshotsRequest) SetSnapshotIds(v string) *ListSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *ListSnapshotsRequest) SetSourceDiskType(v string) *ListSnapshotsRequest {
	s.SourceDiskType = &v
	return s
}

func (s *ListSnapshotsRequest) SetTag(v []*ListSnapshotsRequestTag) *ListSnapshotsRequest {
	s.Tag = v
	return s
}

func (s *ListSnapshotsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSnapshotsRequestTag struct {
	// The tag key of the snapshot. The tag key must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the snapshot. The tag value must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSnapshotsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsRequestTag) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListSnapshotsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListSnapshotsRequestTag) SetKey(v string) *ListSnapshotsRequestTag {
	s.Key = &v
	return s
}

func (s *ListSnapshotsRequestTag) SetValue(v string) *ListSnapshotsRequestTag {
	s.Value = &v
	return s
}

func (s *ListSnapshotsRequestTag) Validate() error {
	return dara.Validate(s)
}
