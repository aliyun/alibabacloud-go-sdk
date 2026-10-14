// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstanceSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskAccessProtocol(v string) *AttachInstanceSDGShrinkRequest
	GetDiskAccessProtocol() *string
	SetDiskType(v string) *AttachInstanceSDGShrinkRequest
	GetDiskType() *string
	SetInstanceIdsShrink(v string) *AttachInstanceSDGShrinkRequest
	GetInstanceIdsShrink() *string
	SetLoadOptShrink(v string) *AttachInstanceSDGShrinkRequest
	GetLoadOptShrink() *string
	SetSDGId(v string) *AttachInstanceSDGShrinkRequest
	GetSDGId() *string
}

type AttachInstanceSDGShrinkRequest struct {
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
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The file loading rate optimization parameter.
	LoadOptShrink *string `json:"LoadOpt,omitempty" xml:"LoadOpt,omitempty"`
	// The ID of the SDG.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s AttachInstanceSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachInstanceSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachInstanceSDGShrinkRequest) GetDiskAccessProtocol() *string {
	return s.DiskAccessProtocol
}

func (s *AttachInstanceSDGShrinkRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *AttachInstanceSDGShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *AttachInstanceSDGShrinkRequest) GetLoadOptShrink() *string {
	return s.LoadOptShrink
}

func (s *AttachInstanceSDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *AttachInstanceSDGShrinkRequest) SetDiskAccessProtocol(v string) *AttachInstanceSDGShrinkRequest {
	s.DiskAccessProtocol = &v
	return s
}

func (s *AttachInstanceSDGShrinkRequest) SetDiskType(v string) *AttachInstanceSDGShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *AttachInstanceSDGShrinkRequest) SetInstanceIdsShrink(v string) *AttachInstanceSDGShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *AttachInstanceSDGShrinkRequest) SetLoadOptShrink(v string) *AttachInstanceSDGShrinkRequest {
	s.LoadOptShrink = &v
	return s
}

func (s *AttachInstanceSDGShrinkRequest) SetSDGId(v string) *AttachInstanceSDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *AttachInstanceSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
