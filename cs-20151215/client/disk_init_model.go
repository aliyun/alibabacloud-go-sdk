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
	// The name of the disk, which uniquely identifies a disk in the node pool.
	//
	// - If local_disk is set to false, the value must match a disk_name in scaling_group.data_disks.
	//
	// - If local_disk is set to true, the value must be different from any disk_name in scaling_group.data_disks.
	//
	// This parameter is required.
	//
	// example:
	//
	// disk0
	DiskName *string `json:"disk_name,omitempty" xml:"disk_name,omitempty"`
	// The configuration for the local disk. If this parameter is set to true:
	//
	// - If the disk cannot be matched by using the ECS DiskName and symbolic links, the local disk is matched.
	//
	// - If the disk is detected as uninitialized during boot, it is reinitialized.
	//
	// - Both mount_target and mkfs_type must be specified or left empty at the same time.
	//
	// - If this parameter is used in combination with mount_for_runtime and the local disk is reinitialized, the node must be removed from and then re-imported to the node pool.
	LocalDisk *bool `json:"local_disk,omitempty" xml:"local_disk,omitempty"`
	// The type of the file system for automatic formatting. If this parameter is left empty, the disk is not formatted.
	//
	// example:
	//
	// ext4
	MkfsType *string `json:"mkfs_type,omitempty" xml:"mkfs_type,omitempty"`
	// The disk used to store container runtime data.
	MountForRuntime *bool `json:"mount_for_runtime,omitempty" xml:"mount_for_runtime,omitempty"`
	// The mount directory. If this parameter is left empty, the disk is not mounted.
	//
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
