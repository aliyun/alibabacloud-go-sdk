// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchHistoricalSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchHistoricalSnapshotsResponseBody
	GetCode() *string
	SetLimit(v int32) *SearchHistoricalSnapshotsResponseBody
	GetLimit() *int32
	SetMessage(v string) *SearchHistoricalSnapshotsResponseBody
	GetMessage() *string
	SetNextToken(v string) *SearchHistoricalSnapshotsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SearchHistoricalSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v *SearchHistoricalSnapshotsResponseBodySnapshots) *SearchHistoricalSnapshotsResponseBody
	GetSnapshots() *SearchHistoricalSnapshotsResponseBodySnapshots
	SetSuccess(v bool) *SearchHistoricalSnapshotsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *SearchHistoricalSnapshotsResponseBody
	GetTotalCount() *int32
}

type SearchHistoricalSnapshotsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of historical backup snapshots that are displayed on the current page.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that is required to obtain the next page of backup snapshots.
	//
	// example:
	//
	// BE
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The historical backup snapshots.
	Snapshots *SearchHistoricalSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
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
	// The total number of returned backup snapshots that meet the specified conditions.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchHistoricalSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchHistoricalSnapshotsResponseBody) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchHistoricalSnapshotsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchHistoricalSnapshotsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchHistoricalSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchHistoricalSnapshotsResponseBody) GetSnapshots() *SearchHistoricalSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *SearchHistoricalSnapshotsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchHistoricalSnapshotsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchHistoricalSnapshotsResponseBody) SetCode(v string) *SearchHistoricalSnapshotsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetLimit(v int32) *SearchHistoricalSnapshotsResponseBody {
	s.Limit = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetMessage(v string) *SearchHistoricalSnapshotsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetNextToken(v string) *SearchHistoricalSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetRequestId(v string) *SearchHistoricalSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetSnapshots(v *SearchHistoricalSnapshotsResponseBodySnapshots) *SearchHistoricalSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetSuccess(v bool) *SearchHistoricalSnapshotsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) SetTotalCount(v int32) *SearchHistoricalSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchHistoricalSnapshotsResponseBodySnapshots struct {
	Snapshot []*SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s SearchHistoricalSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshots) GetSnapshot() []*SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	return s.Snapshot
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshots) SetSnapshot(v []*SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) *SearchHistoricalSnapshotsResponseBodySnapshots {
	s.Snapshot = v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshots) Validate() error {
	return dara.Validate(s)
}

type SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot struct {
	// The actual data amount of backup snapshots after duplicates are removed. Unit: bytes.
	//
	// example:
	//
	// 600
	ActualBytes *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The actual number of backup snapshots.
	//
	// >  This parameter is available only for file backup.
	//
	// example:
	//
	// 6
	ActualItems *int64 `json:"ActualItems,omitempty" xml:"ActualItems,omitempty"`
	// Time to archive
	//
	// example:
	//
	// 1640334062
	ArchiveTime *int64 `json:"ArchiveTime,omitempty" xml:"ArchiveTime,omitempty"`
	// The backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **OSS**. This parameter indicates the name of the OSS bucket.
	//
	// example:
	//
	// hbr-backup-oss
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The actual amount of data that is generated by incremental backups. Unit: bytes.
	//
	// example:
	//
	// 800
	BytesDone *int64 `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	// The total amount of data. Unit: bytes.
	//
	// example:
	//
	// 1000
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the ID of the HBR client.
	//
	// example:
	//
	// c-*********************
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The time when the backup snapshot was completed. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **NAS**. This parameter indicates the time when the file system was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1607436917
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the backup snapshot was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The files that record the information about backup failures, including the information about partially completed backups.
	//
	// example:
	//
	// Item	Error Message C:\\Program Files (x86)\\Symantec\\Symantec Endpoint Protection\\14.3.558.0000.105\\Bin\\service.dat	Open: open \\\\?\\C:\\Program Files (x86)\\Symantec\\Symantec Endpoint Protection\\14.3.558.0000.105\\Bin\\service.dat: The process cannot access the file because it is being used by another process. C:\\ProgramData\\McAfee\\Agent\\data\\InstallerFiles\\172e8a3b04b7ab0fd0215f4fb7707e3744b37d83b6743b3eacb94447c74dc9af_contrib.ini	Open: open \\\\?\\C:\\ProgramData\\McAfee\\Agent\\data\\InstallerFiles\\172e8a3b04b7ab0fd0215f4fb7707e3744b37d83b6743b3eacb94447c74dc9af_contrib.ini: Access is denied.
	ErrorFile *string `json:"ErrorFile,omitempty" xml:"ErrorFile,omitempty"`
	// Backup paths not included in the backup job.
	//
	// example:
	//
	// [\\"/test/example_cn-hangzhou_7.txt\\", \\"/test/example_cn-hangzhou_1.txt\\", \\"/test/example_cn-hangzhou_3.txt\\", \\"/test/example_cn-hangzhou_9.txt\\", \\"/test/example_cn-hangzhou_6.txt\\"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// The time when the snapshot expired. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1640334062
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **NAS**. This parameter indicates the ID of the NAS file system.
	//
	// example:
	//
	// 005494
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Backup paths included in the backup job.
	//
	// example:
	//
	// [\\"/test/example_cn-huhehaote_3.txt\\", \\"/test/example_cn-huhehaote_9.txt\\", \\"/test/example_cn-huhehaote_5.txt\\", \\"/test/example_cn-huhehaote_1.txt\\", \\"/test/example_cn-huhehaote_7.txt\\"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// This parameter is valid only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the ID of the ECS instance.
	//
	// example:
	//
	// i-*********************
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Tablestore instance.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of objects that are backed up.
	//
	// >  This parameter is available only for file backup.
	//
	// example:
	//
	// 8
	ItemsDone *int64 `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	// The total number of objects in the data source.
	//
	// >  This parameter is available only for file backup.
	//
	// example:
	//
	// 10
	ItemsTotal *int64 `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	// The ID of the backup job.
	//
	// example:
	//
	// v-*********************
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The hash value of the parent backup snapshot.
	//
	// example:
	//
	// f2fe..
	ParentSnapshotHash *string `json:"ParentSnapshotHash,omitempty" xml:"ParentSnapshotHash,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the path to the files that are backed up.
	//
	// example:
	//
	// ["/home"]
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The source paths.
	//
	// example:
	//
	// "/home"
	Paths *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Struct"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **OSS**. This parameter indicates the prefix of objects that are backed up.
	//
	// example:
	//
	// example/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The time when the backup job ended. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1642521709966
	RangeEnd *int64 `json:"RangeEnd,omitempty" xml:"RangeEnd,omitempty"`
	// The time when the backup job started. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1642492553038
	RangeStart *int64 `json:"RangeStart,omitempty" xml:"RangeStart,omitempty"`
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
	// s-*********************
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// Parent snapshot HASH value before archiving.
	//
	// example:
	//
	// qwer***
	SourceParentSnapshotHash *string `json:"SourceParentSnapshotHash,omitempty" xml:"SourceParentSnapshotHash,omitempty"`
	// Snapshot HASH value before archiving
	//
	// example:
	//
	// qwer***
	SourceSnapshotHash *string `json:"SourceSnapshotHash,omitempty" xml:"SourceSnapshotHash,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: backup snapshots for ECS files
	//
	// 	- **OSS**: backup snapshots for OSS buckets
	//
	// 	- **NAS**: backup snapshots for NAS file systems
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the backup snapshot started. The value is a UNIX timestamp. Unit: seconds.
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
	// Storage type. Values:
	//
	// - **Standard**: Standard.
	//
	// - **Archive**: Archive.
	//
	// - **ColdArchive**: Cold Archive.
	//
	// example:
	//
	// STANDARD
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The name of a table in the Tablestore instance.
	//
	// example:
	//
	// table2
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The time when the backup snapshot was updated. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// Whether to use local NAS.
	//
	// example:
	//
	// false
	UseCommonNas *bool `json:"UseCommonNas,omitempty" xml:"UseCommonNas,omitempty"`
	// The ID of the backup vault that stores the backup snapshot.
	//
	// example:
	//
	// v-0003rf9m17pap3ltpqx5
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) String() string {
	return dara.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetActualBytes() *int64 {
	return s.ActualBytes
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetActualItems() *int64 {
	return s.ActualItems
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetArchiveTime() *int64 {
	return s.ArchiveTime
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetBackupType() *string {
	return s.BackupType
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetBucket() *string {
	return s.Bucket
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetClientId() *string {
	return s.ClientId
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetErrorFile() *string {
	return s.ErrorFile
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetExclude() *string {
	return s.Exclude
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetInclude() *string {
	return s.Include
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetItemsDone() *int64 {
	return s.ItemsDone
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetItemsTotal() *int64 {
	return s.ItemsTotal
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetJobId() *string {
	return s.JobId
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetParentSnapshotHash() *string {
	return s.ParentSnapshotHash
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetPath() *string {
	return s.Path
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetPaths() *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths {
	return s.Paths
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetPrefix() *string {
	return s.Prefix
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetRangeEnd() *int64 {
	return s.RangeEnd
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetRangeStart() *int64 {
	return s.RangeStart
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetRetention() *int64 {
	return s.Retention
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetSourceParentSnapshotHash() *string {
	return s.SourceParentSnapshotHash
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetSourceSnapshotHash() *string {
	return s.SourceSnapshotHash
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetSourceType() *string {
	return s.SourceType
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetStatus() *string {
	return s.Status
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetStorageClass() *string {
	return s.StorageClass
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetTableName() *string {
	return s.TableName
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetUseCommonNas() *bool {
	return s.UseCommonNas
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) GetVaultId() *string {
	return s.VaultId
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetActualBytes(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ActualBytes = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetActualItems(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ActualItems = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetArchiveTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ArchiveTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetBackupType(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.BackupType = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetBucket(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Bucket = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetBytesDone(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.BytesDone = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetBytesTotal(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.BytesTotal = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetClientId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ClientId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetCompleteTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.CompleteTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetCreateTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.CreateTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetCreatedTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.CreatedTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetErrorFile(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ErrorFile = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetExclude(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Exclude = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetExpireTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ExpireTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetFileSystemId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.FileSystemId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetInclude(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Include = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetInstanceId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.InstanceId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetInstanceName(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.InstanceName = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetItemsDone(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ItemsDone = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetItemsTotal(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ItemsTotal = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetJobId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.JobId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetParentSnapshotHash(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.ParentSnapshotHash = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetPath(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Path = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetPaths(v *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Paths = v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetPrefix(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Prefix = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetRangeEnd(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.RangeEnd = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetRangeStart(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.RangeStart = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetRetention(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Retention = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotHash(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotHash = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetSnapshotId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.SnapshotId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetSourceParentSnapshotHash(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceParentSnapshotHash = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetSourceSnapshotHash(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceSnapshotHash = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetSourceType(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.SourceType = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetStartTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.StartTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetStatus(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.Status = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetStorageClass(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.StorageClass = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetTableName(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.TableName = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetUpdatedTime(v int64) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.UpdatedTime = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetUseCommonNas(v bool) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.UseCommonNas = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) SetVaultId(v string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot {
	s.VaultId = &v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshot) Validate() error {
	return dara.Validate(s)
}

type SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths struct {
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) String() string {
	return dara.Prettify(s)
}

func (s SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) GoString() string {
	return s.String()
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) GetPath() []*string {
	return s.Path
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) SetPath(v []*string) *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths {
	s.Path = v
	return s
}

func (s *SearchHistoricalSnapshotsResponseBodySnapshotsSnapshotPaths) Validate() error {
	return dara.Validate(s)
}
