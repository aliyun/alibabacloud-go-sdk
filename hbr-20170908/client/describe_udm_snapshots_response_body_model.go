// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUdmSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeUdmSnapshotsResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeUdmSnapshotsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeUdmSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v []*DescribeUdmSnapshotsResponseBodySnapshots) *DescribeUdmSnapshotsResponseBody
	GetSnapshots() []*DescribeUdmSnapshotsResponseBodySnapshots
	SetSuccess(v bool) *DescribeUdmSnapshotsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeUdmSnapshotsResponseBody
	GetTotalCount() *int64
}

type DescribeUdmSnapshotsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 51CDEECB-7001-51CC-94AC-2A0F2A4B71D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about snapshots.
	Snapshots []*DescribeUdmSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of backup snapshots.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUdmSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdmSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeUdmSnapshotsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeUdmSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUdmSnapshotsResponseBody) GetSnapshots() []*DescribeUdmSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *DescribeUdmSnapshotsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeUdmSnapshotsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeUdmSnapshotsResponseBody) SetCode(v string) *DescribeUdmSnapshotsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetMessage(v string) *DescribeUdmSnapshotsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetRequestId(v string) *DescribeUdmSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetSnapshots(v []*DescribeUdmSnapshotsResponseBodySnapshots) *DescribeUdmSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetSuccess(v bool) *DescribeUdmSnapshotsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) SetTotalCount(v int64) *DescribeUdmSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUdmSnapshotsResponseBodySnapshots struct {
	// The size of the backup snapshot. Unit: bytes.
	//
	// example:
	//
	// 600
	ActualBytes *string `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The special retention type, which is valid only for special backups. Valid values:
	//
	// 	- **WEEKLY**: weekly backups
	//
	// 	- **MONTHLY**: monthly backups
	//
	// 	- **YEARLY**: yearly backups
	//
	// example:
	//
	// WEEKLY
	AdvancedRetentionType *string `json:"AdvancedRetentionType,omitempty" xml:"AdvancedRetentionType,omitempty"`
	ArchiveErrorMessage   *string `json:"ArchiveErrorMessage,omitempty" xml:"ArchiveErrorMessage,omitempty"`
	ArchiveStatus         *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	ArchiveTriggerTime    *int64  `json:"ArchiveTriggerTime,omitempty" xml:"ArchiveTriggerTime,omitempty"`
	// The backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The total amount of data. Unit: bytes.
	//
	// example:
	//
	// 1000
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// Indicates whether the disk backup point can be deleted. This parameter is valid only if the value of SourceType is UDM_ECS_DISK.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	CanBeDeleted *bool `json:"CanBeDeleted,omitempty" xml:"CanBeDeleted,omitempty"`
	// The time when the backup snapshot was completed. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1646895666
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the backup snapshot was created.
	//
	// example:
	//
	// 1607436917
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the backup snapshot was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642496679
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The snapshot details.
	Detail *DescribeUdmSnapshotsResponseBodySnapshotsDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The ID of the cloud disk or local disk.
	//
	// example:
	//
	// d-2ze86h5fga5rfwxxa8ef
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The expiration time of the backup.
	//
	// example:
	//
	// 1640334062
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp1f0pe78dxizrsdcgxd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the backup job.
	//
	// example:
	//
	// job-00030j3chkt******2
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the backup snapshot.
	//
	// example:
	//
	// s-00047mg17p26x*****b
	NativeSnapshotId *string `json:"NativeSnapshotId,omitempty" xml:"NativeSnapshotId,omitempty"`
	// The snapshot information.
	//
	// example:
	//
	// {
	//
	// 					"sourceDiskId":"d-bp17misjuy636t82v7b0",
	//
	// 					"lastModifiedTime":"2022-03-09T11:35:12Z",
	//
	// 					"snapshotSN":"64588-429372675-857161235",
	//
	// 					"snapshotId":"s-bp1fbtwv3e6xr6wpe9e0",
	//
	// 					"creationTime":"2022-03-09T11:31:12Z",
	//
	// 					"snapshotType":"user",
	//
	// 					"usage":"none",
	//
	// 					"description":"",
	//
	// 					"sourceStorageType":"disk",
	//
	// 					"tags":[
	//
	// 						{
	//
	// 							"tagValue":"job-0007e0wqjl0imbrtkmnm",
	//
	// 							"tagKey":"HBR JobId"
	//
	// 						}
	//
	// 					],
	//
	// 					"productCode":"",
	//
	// 					"encrypted":false,
	//
	// 					"sourceDiskType":"system",
	//
	// 					"retentionDays":30,
	//
	// 					"snapshotName":"Created-from-HBR-job:job-0007e0wqjl0imbrtkmnm",
	//
	// 					"kMSKeyId":"",
	//
	// 					"progress":"100%",
	//
	// 					"category":"standard",
	//
	// 					"sourceDiskSize":"20",
	//
	// 					"status":"accomplished"
	//
	// 				}
	NativeSnapshotInfo *string `json:"NativeSnapshotInfo,omitempty" xml:"NativeSnapshotInfo,omitempty"`
	// The hash value of the parent backup snapshot.
	//
	// example:
	//
	// f2fe..
	ParentSnapshotHash *string `json:"ParentSnapshotHash,omitempty" xml:"ParentSnapshotHash,omitempty"`
	// The prefix of the backup snapshot.
	//
	// example:
	//
	// example/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The timestamp of the backup snapshot. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642496679
	RealSnapshotTime *int64 `json:"RealSnapshotTime,omitempty" xml:"RealSnapshotTime,omitempty"`
	// The retention period of the backup snapshot. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The hash value of the backup snapshot.
	//
	// example:
	//
	// f2fe...
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The ID of the backup snapshot.
	//
	// example:
	//
	// s-00047mxg17p26*****b
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **UDM_ECS**: ECS instance backup
	//
	// 	- **UDM_ECS_DISK**: disk backup subtask of ECS instance backup
	//
	// 	- **UDM_DISK**: disk backup
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the backup snapshot was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the backup job. Valid values:
	//
	// 	- **COMPLETE**: The backup job is completed.
	//
	// 	- **PARTIAL_COMPLETE**: The backup job is partially completed.
	//
	// 	- **FAILED**: The backup job has failed.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the backup snapshot was updated. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642496679
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeUdmSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdmSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetActualBytes() *string {
	return s.ActualBytes
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetAdvancedRetentionType() *string {
	return s.AdvancedRetentionType
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetArchiveErrorMessage() *string {
	return s.ArchiveErrorMessage
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetArchiveStatus() *string {
	return s.ArchiveStatus
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetArchiveTriggerTime() *int64 {
	return s.ArchiveTriggerTime
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetCanBeDeleted() *bool {
	return s.CanBeDeleted
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetDetail() *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	return s.Detail
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetJobId() *string {
	return s.JobId
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetNativeSnapshotId() *string {
	return s.NativeSnapshotId
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetNativeSnapshotInfo() *string {
	return s.NativeSnapshotInfo
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetParentSnapshotHash() *string {
	return s.ParentSnapshotHash
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetPrefix() *string {
	return s.Prefix
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetRealSnapshotTime() *int64 {
	return s.RealSnapshotTime
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetActualBytes(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.ActualBytes = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetAdvancedRetentionType(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.AdvancedRetentionType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetArchiveErrorMessage(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.ArchiveErrorMessage = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetArchiveStatus(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.ArchiveStatus = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetArchiveTriggerTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.ArchiveTriggerTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetBackupType(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.BackupType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetBytesTotal(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.BytesTotal = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetCanBeDeleted(v bool) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.CanBeDeleted = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetCompleteTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.CompleteTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetCreateTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.CreateTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetCreatedTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.CreatedTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetDetail(v *DescribeUdmSnapshotsResponseBodySnapshotsDetail) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.Detail = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetDiskId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.DiskId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetExpireTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.ExpireTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetInstanceId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetJobId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.JobId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetNativeSnapshotId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.NativeSnapshotId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetNativeSnapshotInfo(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.NativeSnapshotInfo = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetParentSnapshotHash(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.ParentSnapshotHash = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetPrefix(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.Prefix = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetRealSnapshotTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.RealSnapshotTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetRetention(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.Retention = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetSnapshotHash(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetSourceType(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.SourceType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetStartTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.StartTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) SetUpdatedTime(v int64) *DescribeUdmSnapshotsResponseBodySnapshots {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshots) Validate() error {
	return dara.Validate(s)
}

type DescribeUdmSnapshotsResponseBodySnapshotsDetail struct {
	// The consistency level.
	//
	// example:
	//
	// CRASH
	ConsistentLevel *string `json:"ConsistentLevel,omitempty" xml:"ConsistentLevel,omitempty"`
	// Indicates whether the system disk is included.
	//
	// example:
	//
	// true
	ContainOsDisk *bool `json:"ContainOsDisk,omitempty" xml:"ContainOsDisk,omitempty"`
	// The type of the source disk.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The name of the disk.
	//
	// example:
	//
	// /dev/xvdb
	DiskDevName *string `json:"DiskDevName,omitempty" xml:"DiskDevName,omitempty"`
	// The mapping between the device and the recovery point ID.
	//
	// example:
	//
	// {
	//
	//     "/dev/xvdb":"s-0000u7y6wm3v1e7hxh5a",
	//
	//     "/dev/xvda":"s-0004bl6yr5pt89jjsv5a"
	//
	// }
	DiskHbrSnapshotIdWithDeviceMap map[string]interface{} `json:"DiskHbrSnapshotIdWithDeviceMap,omitempty" xml:"DiskHbrSnapshotIdWithDeviceMap,omitempty"`
	// The IDs of the disks that are backed up at the recovery point.
	DiskIdList []*string `json:"DiskIdList,omitempty" xml:"DiskIdList,omitempty" type:"Repeated"`
	// The reason for the downgrade.
	//
	// example:
	//
	// HBR.NoRamRoleBound
	DowngradeReason *string `json:"DowngradeReason,omitempty" xml:"DowngradeReason,omitempty"`
	// The hostname.
	//
	// example:
	//
	// iZbpxxxxxxxxxxxxxxxxe2Z
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The mapping between the instance ID and the disk ID.
	//
	// example:
	//
	// {
	//
	//     "i-bp1dlp0keohh7ids4uo6":"d-bp1e6427vhd320hifvs",
	//
	//     "i-bp1dlp0keohh7ids4uo6":"d-bp1e6427vhd320hifvd"
	//
	// }
	InstanceIdWithDiskIdListMap map[string]interface{} `json:"InstanceIdWithDiskIdListMap,omitempty" xml:"InstanceIdWithDiskIdListMap,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// swh-hbr
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The specifications of the source instance.
	//
	// example:
	//
	// ecs.c6.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Indicates whether the backup is created by the instant clone feature.
	//
	// example:
	//
	// false
	InstantAccess *bool `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	// The list of snapshot IDs, corresponding to DiskIdList.
	NativeSnapshotIdList []*string `json:"NativeSnapshotIdList,omitempty" xml:"NativeSnapshotIdList,omitempty" type:"Repeated"`
	// The ID of the system disk.
	//
	// example:
	//
	// d-bp1e6427vhd320hifvc
	OsDiskId *string `json:"OsDiskId,omitempty" xml:"OsDiskId,omitempty"`
	// The name of the operating system.
	//
	// example:
	//
	// Debian 10.10 64-bit (UEFI)
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// The English name of the operating system.
	//
	// example:
	//
	// Debian  11.1 64 bit
	OsNameEn *string `json:"OsNameEn,omitempty" xml:"OsNameEn,omitempty"`
	// The type of the operating system. Valid values: linux and windows.
	//
	// example:
	//
	// windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The performance level of the source disk.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The system platform.
	//
	// example:
	//
	// CentOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the snapshot group.
	//
	// example:
	//
	// ssg-uf6856txcaq31uj***
	SnapshotGroupId *string `json:"SnapshotGroupId,omitempty" xml:"SnapshotGroupId,omitempty"`
	// Indicates whether the disk is a system disk.
	//
	// example:
	//
	// true
	SystemDisk *bool `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// BNSHSVR42 IPGUARD
	VmName *string `json:"VmName,omitempty" xml:"VmName,omitempty"`
}

func (s DescribeUdmSnapshotsResponseBodySnapshotsDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdmSnapshotsResponseBodySnapshotsDetail) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetConsistentLevel() *string {
	return s.ConsistentLevel
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetContainOsDisk() *bool {
	return s.ContainOsDisk
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetDiskDevName() *string {
	return s.DiskDevName
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetDiskHbrSnapshotIdWithDeviceMap() map[string]interface{} {
	return s.DiskHbrSnapshotIdWithDeviceMap
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetDiskIdList() []*string {
	return s.DiskIdList
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetDowngradeReason() *string {
	return s.DowngradeReason
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetHostName() *string {
	return s.HostName
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetInstanceIdWithDiskIdListMap() map[string]interface{} {
	return s.InstanceIdWithDiskIdListMap
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetNativeSnapshotIdList() []*string {
	return s.NativeSnapshotIdList
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetOsDiskId() *string {
	return s.OsDiskId
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetOsName() *string {
	return s.OsName
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetOsNameEn() *string {
	return s.OsNameEn
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetOsType() *string {
	return s.OsType
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetSnapshotGroupId() *string {
	return s.SnapshotGroupId
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetSystemDisk() *bool {
	return s.SystemDisk
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) GetVmName() *string {
	return s.VmName
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetConsistentLevel(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.ConsistentLevel = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetContainOsDisk(v bool) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.ContainOsDisk = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDiskCategory(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DiskCategory = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDiskDevName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DiskDevName = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDiskHbrSnapshotIdWithDeviceMap(v map[string]interface{}) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DiskHbrSnapshotIdWithDeviceMap = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDiskIdList(v []*string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DiskIdList = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetDowngradeReason(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.DowngradeReason = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetHostName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.HostName = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetInstanceIdWithDiskIdListMap(v map[string]interface{}) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.InstanceIdWithDiskIdListMap = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetInstanceName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.InstanceName = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetInstanceType(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.InstanceType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetInstantAccess(v bool) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.InstantAccess = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetNativeSnapshotIdList(v []*string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.NativeSnapshotIdList = v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetOsDiskId(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.OsDiskId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetOsName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.OsName = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetOsNameEn(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.OsNameEn = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetOsType(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.OsType = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetPerformanceLevel(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetPlatform(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.Platform = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetSnapshotGroupId(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.SnapshotGroupId = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetSystemDisk(v bool) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.SystemDisk = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) SetVmName(v string) *DescribeUdmSnapshotsResponseBodySnapshotsDetail {
	s.VmName = &v
	return s
}

func (s *DescribeUdmSnapshotsResponseBodySnapshotsDetail) Validate() error {
	return dara.Validate(s)
}
