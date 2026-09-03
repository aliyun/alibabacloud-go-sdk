// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSnapshotsResponseBody
	GetCode() *string
	SetMaxResults(v int32) *ListSnapshotsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListSnapshotsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListSnapshotsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v []*ListSnapshotsResponseBodySnapshots) *ListSnapshotsResponseBody
	GetSnapshots() []*ListSnapshotsResponseBodySnapshots
	SetSuccess(v bool) *ListSnapshotsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSnapshotsResponseBody
	GetTotalCount() *int32
}

type ListSnapshotsResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The maximum number of results returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The return message. The value "successful" is returned for successful requests. An error message is returned for failed requests.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token for the next page. An empty NextToken indicates that no more pages are available.
	//
	// example:
	//
	// aWQj********MCMy
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17189276-****-****-****-0FF51B5A41A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of backup points.
	Snapshots []*ListSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// Indicates whether the request was successful.
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSnapshotsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSnapshotsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSnapshotsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSnapshotsResponseBody) GetSnapshots() []*ListSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *ListSnapshotsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSnapshotsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSnapshotsResponseBody) SetCode(v string) *ListSnapshotsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetMaxResults(v int32) *ListSnapshotsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetMessage(v string) *ListSnapshotsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetNextToken(v string) *ListSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetSnapshots(v []*ListSnapshotsResponseBodySnapshots) *ListSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsResponseBody) SetSuccess(v bool) *ListSnapshotsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetTotalCount(v int32) *ListSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSnapshotsResponseBody) Validate() error {
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

type ListSnapshotsResponseBodySnapshots struct {
	// The actual amount of data written, in bytes.
	//
	// example:
	//
	// 600
	ActualBytes *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The actual number of nodes.
	//
	// example:
	//
	// 6
	ActualItems *int64 `json:"ActualItems,omitempty" xml:"ActualItems,omitempty"`
	// The time when the backup was archived.
	//
	// example:
	//
	// 1640334062
	ArchiveTime *int64 `json:"ArchiveTime,omitempty" xml:"ArchiveTime,omitempty"`
	// The backup type. The value is **COMPLETE**, which indicates a full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The amount of data that has been backed up, in bytes.
	//
	// example:
	//
	// 800
	BytesDone *int64 `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	// The total amount of data to back up, in bytes.
	//
	// example:
	//
	// 1000
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// When **SourceType*	- is set to **ECS_FILE*	- or **File**, this parameter indicates the backup client ID. In other cases, it indicates the ID of the backup data source.
	//
	// example:
	//
	// c-*********************
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The backup completion time. A UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1642496679
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the backup point was created, in seconds.
	//
	// example:
	//
	// 1607436917
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the backup was created.
	//
	// example:
	//
	// 1642496679
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The file that records backup failure information, including information about partially completed backups.
	//
	// example:
	//
	// temp/report/123456789/job-xxxxxxxxx_failed.zip
	ErrorFile *string `json:"ErrorFile,omitempty" xml:"ErrorFile,omitempty"`
	// The source paths excluded from the backup job.
	//
	// example:
	//
	// [\\"/test/example_cn-hangzhou_7.txt\\", \\"/test/example_cn-hangzhou_1.txt\\", \\"/test/example_cn-hangzhou_3.txt\\", \\"/test/example_cn-hangzhou_9.txt\\", \\"/test/example_cn-hangzhou_6.txt\\"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// The backup expiration time.
	//
	// example:
	//
	// 1771901707
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The source paths included in the backup job.
	//
	// example:
	//
	// [\\"/test/example_cn-huhehaote_3.txt\\", \\"/test/example_cn-huhehaote_9.txt\\", \\"/test/example_cn-huhehaote_5.txt\\", \\"/test/example_cn-huhehaote_1.txt\\", \\"/test/example_cn-huhehaote_7.txt\\"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The instance ID of the backup data source.
	//
	// example:
	//
	// i-*********************
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Tablestore instance name.
	//
	// example:
	//
	// instancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of nodes that have been backed up.
	//
	// example:
	//
	// 8
	ItemsDone *int64 `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	// The total number of nodes to back up.
	//
	// example:
	//
	// 10
	ItemsTotal *int64 `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	// The backup job ID.
	//
	// example:
	//
	// job-00030j3chkt******2
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The source path.
	//
	// example:
	//
	// /home
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The list of source paths.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The backup plan ID.
	//
	// example:
	//
	// po-123***7890
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The protected source data size, in bytes. When SourceType is set to ECS_FILE, this indicates the protected cloud disk capacity.
	//
	// example:
	//
	// 42949672960
	ProtectedDataSize *int64 `json:"ProtectedDataSize,omitempty" xml:"ProtectedDataSize,omitempty"`
	// The end time of the Tablestore backup job execution. A UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 1642521709966
	RangeEnd *int64 `json:"RangeEnd,omitempty" xml:"RangeEnd,omitempty"`
	// The start time of the Tablestore backup job execution. A UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 1642492553038
	RangeStart *int64 `json:"RangeStart,omitempty" xml:"RangeStart,omitempty"`
	// The retention period of the backup, in days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The hash value of the backup point.
	//
	// example:
	//
	// f2ac5fd243**********************bc4451777be019
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The backup point ID.
	//
	// example:
	//
	// s-00047mxg17p26*****b
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The data source type. Valid values:
	//
	// - **ECS_FILE**: ECS file backup
	//
	// - **File**: On-premises file backup
	//
	// - **OSS**: OSS backup
	//
	// - **NAS**: Alibaba Cloud NAS backup
	//
	// - **COMMON_NAS**: On-premises NAS backup
	//
	// - **CONTAINER**: Container backup
	//
	// - **OTS_TABLE**: Tablestore backup
	//
	// - **COMMON_FILE_SYSTEM**: CPFS backup
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The backup point status.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class.
	//
	// example:
	//
	// STANDARD
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The name of the data table in the Tablestore instance.
	//
	// example:
	//
	// table2
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The ID of the vault that stores the backup data.
	//
	// example:
	//
	// v-00030j*******sn
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s ListSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodySnapshots) GetActualBytes() *int64 {
	return s.ActualBytes
}

func (s *ListSnapshotsResponseBodySnapshots) GetActualItems() *int64 {
	return s.ActualItems
}

func (s *ListSnapshotsResponseBodySnapshots) GetArchiveTime() *int64 {
	return s.ArchiveTime
}

func (s *ListSnapshotsResponseBodySnapshots) GetBackupType() *string {
	return s.BackupType
}

func (s *ListSnapshotsResponseBodySnapshots) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *ListSnapshotsResponseBodySnapshots) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *ListSnapshotsResponseBodySnapshots) GetClientId() *string {
	return s.ClientId
}

func (s *ListSnapshotsResponseBodySnapshots) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *ListSnapshotsResponseBodySnapshots) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListSnapshotsResponseBodySnapshots) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListSnapshotsResponseBodySnapshots) GetErrorFile() *string {
	return s.ErrorFile
}

func (s *ListSnapshotsResponseBodySnapshots) GetExclude() *string {
	return s.Exclude
}

func (s *ListSnapshotsResponseBodySnapshots) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListSnapshotsResponseBodySnapshots) GetInclude() *string {
	return s.Include
}

func (s *ListSnapshotsResponseBodySnapshots) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSnapshotsResponseBodySnapshots) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListSnapshotsResponseBodySnapshots) GetItemsDone() *int64 {
	return s.ItemsDone
}

func (s *ListSnapshotsResponseBodySnapshots) GetItemsTotal() *int64 {
	return s.ItemsTotal
}

func (s *ListSnapshotsResponseBodySnapshots) GetJobId() *string {
	return s.JobId
}

func (s *ListSnapshotsResponseBodySnapshots) GetPath() *string {
	return s.Path
}

func (s *ListSnapshotsResponseBodySnapshots) GetPaths() []*string {
	return s.Paths
}

func (s *ListSnapshotsResponseBodySnapshots) GetPlanId() *string {
	return s.PlanId
}

func (s *ListSnapshotsResponseBodySnapshots) GetProtectedDataSize() *int64 {
	return s.ProtectedDataSize
}

func (s *ListSnapshotsResponseBodySnapshots) GetRangeEnd() *int64 {
	return s.RangeEnd
}

func (s *ListSnapshotsResponseBodySnapshots) GetRangeStart() *int64 {
	return s.RangeStart
}

func (s *ListSnapshotsResponseBodySnapshots) GetRetention() *int64 {
	return s.Retention
}

func (s *ListSnapshotsResponseBodySnapshots) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *ListSnapshotsResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ListSnapshotsResponseBodySnapshots) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSnapshotsResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *ListSnapshotsResponseBodySnapshots) GetStorageClass() *string {
	return s.StorageClass
}

func (s *ListSnapshotsResponseBodySnapshots) GetTableName() *string {
	return s.TableName
}

func (s *ListSnapshotsResponseBodySnapshots) GetVaultId() *string {
	return s.VaultId
}

func (s *ListSnapshotsResponseBodySnapshots) SetActualBytes(v int64) *ListSnapshotsResponseBodySnapshots {
	s.ActualBytes = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetActualItems(v int64) *ListSnapshotsResponseBodySnapshots {
	s.ActualItems = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetArchiveTime(v int64) *ListSnapshotsResponseBodySnapshots {
	s.ArchiveTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetBackupType(v string) *ListSnapshotsResponseBodySnapshots {
	s.BackupType = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetBytesDone(v int64) *ListSnapshotsResponseBodySnapshots {
	s.BytesDone = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetBytesTotal(v int64) *ListSnapshotsResponseBodySnapshots {
	s.BytesTotal = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetClientId(v string) *ListSnapshotsResponseBodySnapshots {
	s.ClientId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetCompleteTime(v int64) *ListSnapshotsResponseBodySnapshots {
	s.CompleteTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetCreateTime(v int64) *ListSnapshotsResponseBodySnapshots {
	s.CreateTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetCreatedTime(v int64) *ListSnapshotsResponseBodySnapshots {
	s.CreatedTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetErrorFile(v string) *ListSnapshotsResponseBodySnapshots {
	s.ErrorFile = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetExclude(v string) *ListSnapshotsResponseBodySnapshots {
	s.Exclude = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetExpireTime(v int64) *ListSnapshotsResponseBodySnapshots {
	s.ExpireTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetInclude(v string) *ListSnapshotsResponseBodySnapshots {
	s.Include = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetInstanceId(v string) *ListSnapshotsResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetInstanceName(v string) *ListSnapshotsResponseBodySnapshots {
	s.InstanceName = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetItemsDone(v int64) *ListSnapshotsResponseBodySnapshots {
	s.ItemsDone = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetItemsTotal(v int64) *ListSnapshotsResponseBodySnapshots {
	s.ItemsTotal = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetJobId(v string) *ListSnapshotsResponseBodySnapshots {
	s.JobId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetPath(v string) *ListSnapshotsResponseBodySnapshots {
	s.Path = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetPaths(v []*string) *ListSnapshotsResponseBodySnapshots {
	s.Paths = v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetPlanId(v string) *ListSnapshotsResponseBodySnapshots {
	s.PlanId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetProtectedDataSize(v int64) *ListSnapshotsResponseBodySnapshots {
	s.ProtectedDataSize = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRangeEnd(v int64) *ListSnapshotsResponseBodySnapshots {
	s.RangeEnd = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRangeStart(v int64) *ListSnapshotsResponseBodySnapshots {
	s.RangeStart = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRetention(v int64) *ListSnapshotsResponseBodySnapshots {
	s.Retention = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSnapshotHash(v string) *ListSnapshotsResponseBodySnapshots {
	s.SnapshotHash = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *ListSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSourceType(v string) *ListSnapshotsResponseBodySnapshots {
	s.SourceType = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetStatus(v string) *ListSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetStorageClass(v string) *ListSnapshotsResponseBodySnapshots {
	s.StorageClass = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetTableName(v string) *ListSnapshotsResponseBodySnapshots {
	s.TableName = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetVaultId(v string) *ListSnapshotsResponseBodySnapshots {
	s.VaultId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) Validate() error {
	return dara.Validate(s)
}
