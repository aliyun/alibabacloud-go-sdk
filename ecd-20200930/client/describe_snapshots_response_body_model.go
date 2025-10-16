// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeSnapshotsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody
	GetSnapshots() []*DescribeSnapshotsResponseBodySnapshots
}

type DescribeSnapshotsResponseBody struct {
	// If the NextToken parameter is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The snapshots.
	Snapshots []*DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotsResponseBody) GetSnapshots() []*DescribeSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *DescribeSnapshotsResponseBody) SetNextToken(v string) *DescribeSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeSnapshotsResponseBody) Validate() error {
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnapshotsResponseBodySnapshots struct {
	// The point in time at which the snapshot was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-mm-ddthh:mm:ssz` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-20T14:52:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The user who creates the snapshot.
	//
	// example:
	//
	// Administrator
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The time when the snapshot was deleted. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-mm-ddthh:mm:ssz` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-12-20T14:52:28Z
	DeletionTime *string `json:"DeletionTime,omitempty" xml:"DeletionTime,omitempty"`
	// The description of the snapshot.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the cloud computer to which the snapshot belongs.
	//
	// example:
	//
	// ecd-g03l3tlm8djoj****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// test
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The status of the cloud computer.
	//
	// Valid values:
	//
	// 	- Stopped
	//
	// 	- Starting
	//
	// 	- Rebuilding
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Expired
	//
	// 	- Deleted
	//
	// 	- Pending
	//
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DiskStatus    *string `json:"DiskStatus,omitempty" xml:"DiskStatus,omitempty"`
	OsType        *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The progress of creating the snapshot. Unit: %.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The protocol type.
	//
	// Valid values:
	//
	// 	- HDX: High-definition Experience (HDX) protocol
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ASP: in-house Adaptive Streaming Protocol (ASP)
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The remaining time that is required to complete the snapshot creation. Unit: seconds.
	//
	// >  When the `Status` value is `PROGRESSING`, the `RemainTime` value is `-1`. A value of -1 indicates that the system is calculating the remaining time.
	//
	// example:
	//
	// 30
	RemainTime *int32 `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	// The ID of the restore point.
	//
	// example:
	//
	// rp-btgmaa20wkcju****
	RestorePointId *string `json:"RestorePointId,omitempty" xml:"RestorePointId,omitempty"`
	// The name of the restore point.
	//
	// example:
	//
	// DataDiskBackup
	RestorePointName *string `json:"RestorePointName,omitempty" xml:"RestorePointName,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-2zeipxmnhej803x7****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The name of the snapshot.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The type of the snapshot.
	//
	// Valid values:
	//
	// 	- AUTO: automatic snapshot
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- USER: manual snapshot
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// USER
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The capacity of the source disk. Unit: GiB.
	//
	// example:
	//
	// 150
	SourceDiskSize *string `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty"`
	// The type of the source disk.
	//
	// Valid values:
	//
	// 	- SYSTEM: system disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DATA: data disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SYSTEM
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The status of the snapshot.
	//
	// Valid values:
	//
	// 	- PROGRESSING: The snapshot is being created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- FAILED: The snapshot fails to be created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ACCOMPLISHED: The snapshot is created.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// ACCOMPLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether disk encryption is enabled.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used when disk encryption is enabled. You can call the [ListKeys](https://help.aliyun.com/document_detail/28951.html) operation to query the list of KMS keys.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetCreator() *string {
	return s.Creator
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetDeletionTime() *string {
	return s.DeletionTime
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetDescription() *string {
	return s.Description
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetDiskStatus() *string {
	return s.DiskStatus
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetOsType() *string {
	return s.OsType
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetProgress() *string {
	return s.Progress
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetRemainTime() *int32 {
	return s.RemainTime
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetRestorePointId() *string {
	return s.RestorePointId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetRestorePointName() *string {
	return s.RestorePointName
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshotType() *string {
	return s.SnapshotType
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSourceDiskSize() *string {
	return s.SourceDiskSize
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetCreationTime(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetCreator(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Creator = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDeletionTime(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DeletionTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDesktopId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDesktopName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DesktopName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDesktopStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDiskStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DiskStatus = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetOsType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.OsType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetProtocolType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.ProtocolType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshots {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRestorePointId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.RestorePointId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRestorePointName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.RestorePointName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskSize(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetVolumeEncryptionEnabled(v bool) *DescribeSnapshotsResponseBodySnapshots {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetVolumeEncryptionKey(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) Validate() error {
	return dara.Validate(s)
}
