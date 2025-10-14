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
	DiskAccessProtocol *string `json:"DiskAccessProtocol,omitempty" xml:"DiskAccessProtocol,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The IDs of the instances.
	//
	// This parameter is required.
	InstanceIds []*string                        `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	LoadOpt     *AttachInstanceSDGRequestLoadOpt `json:"LoadOpt,omitempty" xml:"LoadOpt,omitempty" type:"Struct"`
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
	BlockRwSplit     *bool  `json:"BlockRwSplit,omitempty" xml:"BlockRwSplit,omitempty"`
	BlockRwSplitSize *int32 `json:"BlockRwSplitSize,omitempty" xml:"BlockRwSplitSize,omitempty"`
	Cache            *bool  `json:"Cache,omitempty" xml:"Cache,omitempty"`
	CacheSize        *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
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
