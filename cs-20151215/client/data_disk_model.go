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
	// Deprecated
	//
	// Specifies whether to automatically mount the data disk.
	//
	// example:
	//
	// true
	AutoFormat *bool `json:"auto_format,omitempty" xml:"auto_format,omitempty"`
	// The ID of the automatic snapshot policy. The system performs automatic backup for a cloud disk based on the specified automatic snapshot policy.
	//
	// By default, this parameter is empty, which indicates that automatic backup is disabled.
	//
	// example:
	//
	// sp-2zej1nogjvovnz4z****
	AutoSnapshotPolicyId *string `json:"auto_snapshot_policy_id,omitempty" xml:"auto_snapshot_policy_id,omitempty"`
	// Specifies whether to enable the burst feature for the data disk. Valid values:
	//
	// 	- `true`: enables the burst feature.
	//
	// 	- `false`: disables the burst feature for the data disk.
	//
	// This parameter is available only if `DiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"bursting_enabled,omitempty" xml:"bursting_enabled,omitempty"`
	// The category of data disk. Valid values:
	//
	// 	- `cloud`: basic disk.
	//
	// 	- `cloud_efficiency`: ultra disk
	//
	// 	- `cloud_ssd`: standard SSD.
	//
	// 	- `cloud_essd`: Enterprise ESSD (ESSD).
	//
	// 	- `cloud_auto`: ESSD AutoPL disk.
	//
	// 	- `cloud_essd_entry`: ESSD Entry disk.
	//
	// 	- `elastic_ephemeral_disk_premium`: premium elastic ephemeral disk.
	//
	// 	- `elastic_ephemeral_disk_standard`: standard elastic ephemeral disk.
	//
	// Default value: `cloud_efficiency`.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The mount target of the data disk. If you do not specify this parameter, the system automatically assigns a mount target when you create an Elastic Compute Service (ECS) instance. Valid values: /dev/xvdb to /dev/xvdz.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"device,omitempty" xml:"device,omitempty"`
	// The data disk name. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// DataDiskName
	DiskName *string `json:"disk_name,omitempty" xml:"disk_name,omitempty"`
	// Specifies whether to encrypt the data disk. Valid values:
	//
	// 	- `true`: encrypts the data disk.
	//
	// 	- `false`: does not encrypt the data disk.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	Encrypted *string `json:"encrypted,omitempty" xml:"encrypted,omitempty"`
	// Deprecated
	//
	// The file system that is mounted. This parameter takes effect only if auto_format is set to true. Valid values: ext4 and xfs.
	//
	// example:
	//
	// ext4
	FileSystem *string `json:"file_system,omitempty" xml:"file_system,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used to encrypt the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KmsKeyId *string `json:"kms_key_id,omitempty" xml:"kms_key_id,omitempty"`
	// Deprecated
	//
	// The path to which the data disk is mounted. You must specify a valid path.
	//
	// example:
	//
	// /mnt/path1
	MountTarget *string `json:"mount_target,omitempty" xml:"mount_target,omitempty"`
	// The performance level (PL) of the data disk. This parameter takes effect only for an ESSD. This parameter is related to the disk size. For more information, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"performance_level,omitempty" xml:"performance_level,omitempty"`
	// The preset IOPS of the data disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter is available only if `DiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 1000
	ProvisionedIops *int64 `json:"provisioned_iops,omitempty" xml:"provisioned_iops,omitempty"`
	// The size of the data disk. Unit: GiB.
	//
	// Valid values: 40 to 32768.
	//
	// Default value: `120`.
	//
	// example:
	//
	// 40
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The ID of the snapshot that you want to use to create the data disk. If this parameter is specified, the specified size of the data disk is ignored. The size of the disk equals the size of the specified snapshot. If the snapshot was created on or before July 15, 2013, the API request is rejected and the InvalidSnapshot.TooOld message is returned.
	//
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
