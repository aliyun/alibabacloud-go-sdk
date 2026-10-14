// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstanceSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskAccessProtocol(v string) *AttachInstanceSDGRequest
	GetDiskAccessProtocol() *string
	SetDiskType(v string) *AttachInstanceSDGRequest
	GetDiskType() *string
	SetInstanceIds(v []*string) *AttachInstanceSDGRequest
	GetInstanceIds() []*string
	SetLoadOpt(v *AttachInstanceSDGRequestLoadOpt) *AttachInstanceSDGRequest
	GetLoadOpt() *AttachInstanceSDGRequestLoadOpt
	SetSDGId(v string) *AttachInstanceSDGRequest
	GetSDGId() *string
}

type AttachInstanceSDGRequest struct {
	// The cloud disk access protocol type. Valid values:
	//
	// - iscsi: the default protocol for standard cloud disks. Both rbd cloud disks and standard cloud disks support this protocol.
	//
	//  - rbd: the default protocol for rbd cloud disks. Only rbd cloud disks support this mount protocol.
	//
	// example:
	//
	// iscsi
	DiskAccessProtocol *string `json:"DiskAccessProtocol,omitempty" xml:"DiskAccessProtocol,omitempty"`
	// The type of the shared cloud disk to be created. Valid values:
	//
	// - standard: standard cloud disk.
	//
	// - rbd: rbd cloud disk.
	//
	// >
	//
	// > - If no shared cloud disk type is specified, a standard cloud disk is created by default.
	//
	// > - An rbd shared cloud disk can be created only when the parent disk type is rbd.
	//
	// example:
	//
	// standard
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The list of instance IDs.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The file loading rate optimization parameter.
	LoadOpt *AttachInstanceSDGRequestLoadOpt `json:"LoadOpt,omitempty" xml:"LoadOpt,omitempty" type:"Struct"`
	// The ID of the SDG.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s AttachInstanceSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceSDGRequest) GoString() string {
	return s.String()
}

func (s *AttachInstanceSDGRequest) GetDiskAccessProtocol() *string {
	return s.DiskAccessProtocol
}

func (s *AttachInstanceSDGRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *AttachInstanceSDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AttachInstanceSDGRequest) GetLoadOpt() *AttachInstanceSDGRequestLoadOpt {
	return s.LoadOpt
}

func (s *AttachInstanceSDGRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *AttachInstanceSDGRequest) SetDiskAccessProtocol(v string) *AttachInstanceSDGRequest {
	s.DiskAccessProtocol = &v
	return s
}

func (s *AttachInstanceSDGRequest) SetDiskType(v string) *AttachInstanceSDGRequest {
	s.DiskType = &v
	return s
}

func (s *AttachInstanceSDGRequest) SetInstanceIds(v []*string) *AttachInstanceSDGRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachInstanceSDGRequest) SetLoadOpt(v *AttachInstanceSDGRequestLoadOpt) *AttachInstanceSDGRequest {
	s.LoadOpt = v
	return s
}

func (s *AttachInstanceSDGRequest) SetSDGId(v string) *AttachInstanceSDGRequest {
	s.SDGId = &v
	return s
}

func (s *AttachInstanceSDGRequest) Validate() error {
	if s.LoadOpt != nil {
		if err := s.LoadOpt.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AttachInstanceSDGRequestLoadOpt struct {
	// Specifies whether to enable block device read/write splitting. Default value: false.
	//
	// example:
	//
	// true
	BlockRwSplit *bool `json:"BlockRwSplit,omitempty" xml:"BlockRwSplit,omitempty"`
	// The size for block device read/write splitting, in GiB. A value of 0 indicates that the default size is used.
	//
	// example:
	//
	// 20
	BlockRwSplitSize *int32 `json:"BlockRwSplitSize,omitempty" xml:"BlockRwSplitSize,omitempty"`
	// Specifies whether to enable caching. Block device read/write splitting must be enabled before caching can be used. Default value: false.
	//
	// example:
	//
	// true
	Cache *bool `json:"Cache,omitempty" xml:"Cache,omitempty"`
	// The cache size, in GiB. A value of 0 indicates that the default size is used.
	//
	// example:
	//
	// 20
	CacheSize *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
}

func (s AttachInstanceSDGRequestLoadOpt) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceSDGRequestLoadOpt) GoString() string {
	return s.String()
}

func (s *AttachInstanceSDGRequestLoadOpt) GetBlockRwSplit() *bool {
	return s.BlockRwSplit
}

func (s *AttachInstanceSDGRequestLoadOpt) GetBlockRwSplitSize() *int32 {
	return s.BlockRwSplitSize
}

func (s *AttachInstanceSDGRequestLoadOpt) GetCache() *bool {
	return s.Cache
}

func (s *AttachInstanceSDGRequestLoadOpt) GetCacheSize() *int32 {
	return s.CacheSize
}

func (s *AttachInstanceSDGRequestLoadOpt) SetBlockRwSplit(v bool) *AttachInstanceSDGRequestLoadOpt {
	s.BlockRwSplit = &v
	return s
}

func (s *AttachInstanceSDGRequestLoadOpt) SetBlockRwSplitSize(v int32) *AttachInstanceSDGRequestLoadOpt {
	s.BlockRwSplitSize = &v
	return s
}

func (s *AttachInstanceSDGRequestLoadOpt) SetCache(v bool) *AttachInstanceSDGRequestLoadOpt {
	s.Cache = &v
	return s
}

func (s *AttachInstanceSDGRequestLoadOpt) SetCacheSize(v int32) *AttachInstanceSDGRequestLoadOpt {
	s.CacheSize = &v
	return s
}

func (s *AttachInstanceSDGRequestLoadOpt) Validate() error {
	return dara.Validate(s)
}
