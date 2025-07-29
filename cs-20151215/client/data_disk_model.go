// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataDisk interface {
	dara.Model
	String() string
	GoString() string
	SetAutoFormat(v bool) *DataDisk
	GetAutoFormat() *bool
	SetAutoSnapshotPolicyId(v string) *DataDisk
	GetAutoSnapshotPolicyId() *string
	SetBurstingEnabled(v bool) *DataDisk
	GetBurstingEnabled() *bool
	SetCategory(v string) *DataDisk
	GetCategory() *string
	SetDevice(v string) *DataDisk
	GetDevice() *string
	SetDiskName(v string) *DataDisk
	GetDiskName() *string
	SetEncrypted(v string) *DataDisk
	GetEncrypted() *string
	SetFileSystem(v string) *DataDisk
	GetFileSystem() *string
	SetKmsKeyId(v string) *DataDisk
	GetKmsKeyId() *string
	SetMountTarget(v string) *DataDisk
	GetMountTarget() *string
	SetPerformanceLevel(v string) *DataDisk
	GetPerformanceLevel() *string
	SetProvisionedIops(v int64) *DataDisk
	GetProvisionedIops() *int64
	SetSize(v int64) *DataDisk
	GetSize() *int64
	SetSnapshotId(v string) *DataDisk
	GetSnapshotId() *string
}

type DataDisk struct {
	// example:
	//
	// true
	AutoFormat *bool `json:"auto_format,omitempty" xml:"auto_format,omitempty"`
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	AutoSnapshotPolicyId *string `json:"auto_snapshot_policy_id,omitempty" xml:"auto_snapshot_policy_id,omitempty"`
	// example:
	//
	// true
	BurstingEnabled *bool `json:"bursting_enabled,omitempty" xml:"bursting_enabled,omitempty"`
	// example:
	//
	// cloud_ssd
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// /dev/xvdb
	Device *string `json:"device,omitempty" xml:"device,omitempty"`
	// example:
	//
	// DataDiskName
	DiskName *string `json:"disk_name,omitempty" xml:"disk_name,omitempty"`
	// example:
	//
	// true
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// example:
	//
	// ext4
	FileSystem *string `json:"file_system,omitempty" xml:"file_system,omitempty"`
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KmsKeyId *string `json:"kms_key_id,omitempty" xml:"kms_key_id,omitempty"`
	// example:
	//
	// /mnt/path1
	MountTarget *string `json:"mount_target,omitempty" xml:"mount_target,omitempty"`
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performance_level,omitempty" xml:"performance_level,omitempty"`
	// example:
	//
	// 1000
	ProvisionedIops *int64 `json:"provisioned_iops,omitempty" xml:"provisioned_iops,omitempty"`
	// example:
	//
	// 40
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// s-280s7****
	SnapshotId *string `json:"snapshot_id,omitempty" xml:"snapshot_id,omitempty"`
}

func (s DataDisk) String() string {
	return dara.Prettify(s)
}

func (s DataDisk) GoString() string {
	return s.String()
}

func (s *DataDisk) GetAutoFormat() *bool {
	return s.AutoFormat
}

func (s *DataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DataDisk) GetCategory() *string {
	return s.Category
}

func (s *DataDisk) GetDevice() *string {
	return s.Device
}

func (s *DataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *DataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *DataDisk) GetFileSystem() *string {
	return s.FileSystem
}

func (s *DataDisk) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *DataDisk) GetMountTarget() *string {
	return s.MountTarget
}

func (s *DataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DataDisk) GetSize() *int64 {
	return s.Size
}

func (s *DataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DataDisk) SetAutoFormat(v bool) *DataDisk {
	s.AutoFormat = &v
	return s
}

func (s *DataDisk) SetAutoSnapshotPolicyId(v string) *DataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DataDisk) SetBurstingEnabled(v bool) *DataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *DataDisk) SetCategory(v string) *DataDisk {
	s.Category = &v
	return s
}

func (s *DataDisk) SetDevice(v string) *DataDisk {
	s.Device = &v
	return s
}

func (s *DataDisk) SetDiskName(v string) *DataDisk {
	s.DiskName = &v
	return s
}

func (s *DataDisk) SetEncrypted(v string) *DataDisk {
	s.Encrypted = &v
	return s
}

func (s *DataDisk) SetFileSystem(v string) *DataDisk {
	s.FileSystem = &v
	return s
}

func (s *DataDisk) SetKmsKeyId(v string) *DataDisk {
	s.KmsKeyId = &v
	return s
}

func (s *DataDisk) SetMountTarget(v string) *DataDisk {
	s.MountTarget = &v
	return s
}

func (s *DataDisk) SetPerformanceLevel(v string) *DataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DataDisk) SetProvisionedIops(v int64) *DataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *DataDisk) SetSize(v int64) *DataDisk {
	s.Size = &v
	return s
}

func (s *DataDisk) SetSnapshotId(v string) *DataDisk {
	s.SnapshotId = &v
	return s
}

func (s *DataDisk) Validate() error {
	return dara.Validate(s)
}
