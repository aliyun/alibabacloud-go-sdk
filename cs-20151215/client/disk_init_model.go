// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiskInit interface {
	dara.Model
	String() string
	GoString() string
	SetDiskName(v string) *DiskInit
	GetDiskName() *string
	SetLocalDisk(v bool) *DiskInit
	GetLocalDisk() *bool
	SetMkfsType(v string) *DiskInit
	GetMkfsType() *string
	SetMountForRuntime(v bool) *DiskInit
	GetMountForRuntime() *bool
	SetMountTarget(v string) *DiskInit
	GetMountTarget() *string
}

type DiskInit struct {
	// This parameter is required.
	//
	// example:
	//
	// disk0
	DiskName  *string `json:"disk_name,omitempty" xml:"disk_name,omitempty"`
	LocalDisk *bool   `json:"local_disk,omitempty" xml:"local_disk,omitempty"`
	// example:
	//
	// ext4
	MkfsType        *string `json:"mkfs_type,omitempty" xml:"mkfs_type,omitempty"`
	MountForRuntime *bool   `json:"mount_for_runtime,omitempty" xml:"mount_for_runtime,omitempty"`
	// example:
	//
	// /mnt/disk0
	MountTarget *string `json:"mount_target,omitempty" xml:"mount_target,omitempty"`
}

func (s DiskInit) String() string {
	return dara.Prettify(s)
}

func (s DiskInit) GoString() string {
	return s.String()
}

func (s *DiskInit) GetDiskName() *string {
	return s.DiskName
}

func (s *DiskInit) GetLocalDisk() *bool {
	return s.LocalDisk
}

func (s *DiskInit) GetMkfsType() *string {
	return s.MkfsType
}

func (s *DiskInit) GetMountForRuntime() *bool {
	return s.MountForRuntime
}

func (s *DiskInit) GetMountTarget() *string {
	return s.MountTarget
}

func (s *DiskInit) SetDiskName(v string) *DiskInit {
	s.DiskName = &v
	return s
}

func (s *DiskInit) SetLocalDisk(v bool) *DiskInit {
	s.LocalDisk = &v
	return s
}

func (s *DiskInit) SetMkfsType(v string) *DiskInit {
	s.MkfsType = &v
	return s
}

func (s *DiskInit) SetMountForRuntime(v bool) *DiskInit {
	s.MountForRuntime = &v
	return s
}

func (s *DiskInit) SetMountTarget(v string) *DiskInit {
	s.MountTarget = &v
	return s
}

func (s *DiskInit) Validate() error {
	return dara.Validate(s)
}
