// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskIds(v string) *ListDisksRequest
	GetDiskIds() *string
	SetDiskType(v string) *ListDisksRequest
	GetDiskType() *string
	SetInstanceId(v string) *ListDisksRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDisksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDisksRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDisksRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListDisksRequest
	GetResourceGroupId() *string
	SetTag(v []*ListDisksRequestTag) *ListDisksRequest
	GetTag() []*ListDisksRequestTag
}

type ListDisksRequest struct {
	// The IDs of the disks. The value can be a JSON array that consists of up to 100 disk IDs. Separate multiple disk IDs with commas (,).
	//
	// example:
	//
	// ["d-bp14wq0149cpp2x****", "d-bp14wq0149cpp2y****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The disk type. Valid values:
	//
	// - system: system disk
	//
	// - data: data disk
	//
	// By default, system disks and data disks are both queried.
	//
	// example:
	//
	// system
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disks.
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
	// The tags that are added to the disks.
	Tag []*ListDisksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDisksRequest) GoString() string {
	return s.String()
}

func (s *ListDisksRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *ListDisksRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *ListDisksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDisksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDisksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDisksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDisksRequest) GetTag() []*ListDisksRequestTag {
	return s.Tag
}

func (s *ListDisksRequest) SetDiskIds(v string) *ListDisksRequest {
	s.DiskIds = &v
	return s
}

func (s *ListDisksRequest) SetDiskType(v string) *ListDisksRequest {
	s.DiskType = &v
	return s
}

func (s *ListDisksRequest) SetInstanceId(v string) *ListDisksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDisksRequest) SetPageNumber(v int32) *ListDisksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDisksRequest) SetPageSize(v int32) *ListDisksRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisksRequest) SetRegionId(v string) *ListDisksRequest {
	s.RegionId = &v
	return s
}

func (s *ListDisksRequest) SetResourceGroupId(v string) *ListDisksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDisksRequest) SetTag(v []*ListDisksRequestTag) *ListDisksRequest {
	s.Tag = v
	return s
}

func (s *ListDisksRequest) Validate() error {
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

type ListDisksRequestTag struct {
	// The tag key. The tag key can be up to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDisksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListDisksRequestTag) GoString() string {
	return s.String()
}

func (s *ListDisksRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListDisksRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListDisksRequestTag) SetKey(v string) *ListDisksRequestTag {
	s.Key = &v
	return s
}

func (s *ListDisksRequestTag) SetValue(v string) *ListDisksRequestTag {
	s.Value = &v
	return s
}

func (s *ListDisksRequestTag) Validate() error {
	return dara.Validate(s)
}
