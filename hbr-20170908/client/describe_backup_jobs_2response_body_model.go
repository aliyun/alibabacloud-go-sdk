// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupJobs2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobs(v *DescribeBackupJobs2ResponseBodyBackupJobs) *DescribeBackupJobs2ResponseBody
	GetBackupJobs() *DescribeBackupJobs2ResponseBodyBackupJobs
	SetCode(v string) *DescribeBackupJobs2ResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeBackupJobs2ResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeBackupJobs2ResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupJobs2ResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupJobs2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupJobs2ResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeBackupJobs2ResponseBody
	GetTotalCount() *int64
}

type DescribeBackupJobs2ResponseBody struct {
	// The returned backup jobs that meet the specified conditions.
	BackupJobs *DescribeBackupJobs2ResponseBodyBackupJobs `json:"BackupJobs,omitempty" xml:"BackupJobs,omitempty" type:"Struct"`
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
	// The page number of the returned page. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned backup jobs that meet the specified conditions.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupJobs2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBody) GetBackupJobs() *DescribeBackupJobs2ResponseBodyBackupJobs {
	return s.BackupJobs
}

func (s *DescribeBackupJobs2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBackupJobs2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBackupJobs2ResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupJobs2ResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupJobs2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupJobs2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupJobs2ResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeBackupJobs2ResponseBody) SetBackupJobs(v *DescribeBackupJobs2ResponseBodyBackupJobs) *DescribeBackupJobs2ResponseBody {
	s.BackupJobs = v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetCode(v string) *DescribeBackupJobs2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetMessage(v string) *DescribeBackupJobs2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetPageNumber(v int32) *DescribeBackupJobs2ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetPageSize(v int32) *DescribeBackupJobs2ResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetRequestId(v string) *DescribeBackupJobs2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetSuccess(v bool) *DescribeBackupJobs2ResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) SetTotalCount(v int64) *DescribeBackupJobs2ResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupJobs2ResponseBodyBackupJobs struct {
	BackupJob []*DescribeBackupJobs2ResponseBodyBackupJobsBackupJob `json:"BackupJob,omitempty" xml:"BackupJob,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobs) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobs) GetBackupJob() []*DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	return s.BackupJob
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobs) SetBackupJob(v []*DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) *DescribeBackupJobs2ResponseBodyBackupJobs {
	s.BackupJob = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobs) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJob struct {
	// The actual amount of data that is backed up after duplicates are removed. Unit: bytes.
	//
	// example:
	//
	// 600
	ActualBytes *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The number of files that are actually processed.
	//
	// example:
	//
	// 8
	ActualFiles *int64 `json:"ActualFiles,omitempty" xml:"ActualFiles,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the actual number of objects that are backed up by the backup job.
	//
	// example:
	//
	// 6
	ActualItems *int64 `json:"ActualItems,omitempty" xml:"ActualItems,omitempty"`
	// The backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **OSS**. This parameter indicates the name of the OSS bucket that is backed up.
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
	// The total amount of data that is backed up from the data source. Unit: bytes.
	//
	// example:
	//
	// 1000
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// The data source details at the destination. Thisparameter is returned only for data synchronization.
	//
	// example:
	//
	// {"dataSourceId": "ds-123456789", "path": "/changelist"}
	ChangeListPath *string `json:"ChangeListPath,omitempty" xml:"ChangeListPath,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the ID of the backup client.
	//
	// example:
	//
	// c-*********************
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The time when the backup job was completed. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **NAS**. This parameter indicates the time when the file system was created. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1607436917
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the backup job was created. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether data is backed up within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- SELF_ACCOUNT: Data is backed up within the same Alibaba Cloud account.
	//
	// 	- CROSS_ACCOUNT: Data is backed up across Alibaba Cloud accounts.
	//
	// example:
	//
	// SELF_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 158975xxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// Destination data source details. (Required only for synchronization)
	//
	// example:
	//
	// {\\"prefix\\":\\"/\\"}
	DestDataSourceDetail *string `json:"DestDataSourceDetail,omitempty" xml:"DestDataSourceDetail,omitempty"`
	// Destination data source ID. (Required only for synchronization)
	//
	// example:
	//
	// ds-000cov4ufudxklj24zdk
	DestDataSourceId *string `json:"DestDataSourceId,omitempty" xml:"DestDataSourceId,omitempty"`
	// Destination data source type. (Required only for synchronization)
	//
	// example:
	//
	// OSS
	DestSourceType *string `json:"DestSourceType,omitempty" xml:"DestSourceType,omitempty"`
	// The udm backup job detail.
	Detail *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	// The error message that is returned for the backup job.
	//
	// example:
	//
	// PARTIAL_COMPLETE
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the paths to the files that are excluded from the backup job. The value must be 1 to 255 characters in length.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **NAS**. This parameter indicates the ID of the NAS file system.
	//
	// example:
	//
	// 005494
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of files that have been processed.
	//
	// example:
	//
	// 9
	FilesDone *int64 `json:"FilesDone,omitempty" xml:"FilesDone,omitempty"`
	// The total number of files to be processed.
	//
	// example:
	//
	// 10
	FilesTotal *int64 `json:"FilesTotal,omitempty" xml:"FilesTotal,omitempty"`
	// The identifier of the container cluster. For a Container Service for Kubernetes (ACK) cluster, specify the cluster ID.
	//
	// example:
	//
	// c83**************************b76
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The paths to the files that are included in the backup job.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **NAS**. This parameter indicates the ID of the ECS instance.
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
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the number of objects that are backed up.
	//
	// example:
	//
	// 8
	ItemsDone *int64 `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the total number of objects in the data source.
	//
	// example:
	//
	// 10
	ItemsTotal *int64 `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	// The ID of the backup job.
	//
	// example:
	//
	// job-000g********w7
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the backup job.
	//
	// example:
	//
	// jobname
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates whether Windows VSS is used to define a backup path.
	//
	// 	- This parameter is available only for Windows ECS instances.
	//
	// 	- If data changes occur in the backup source, the source data must be the same as the data to be backed up before the system sets this parameter to `["UseVSS":true]`.
	//
	// 	- If you use VSS, you cannot back up data from multiple directories.
	//
	// example:
	//
	// {"UseVSS":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The details about the Tablestore instance.
	OtsDetail *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty" type:"Struct"`
	// The backup paths.
	Paths *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Struct"`
	// The ID of the backup plan.
	//
	// example:
	//
	// plan-20********35
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **OSS**. This parameter indicates the prefix of objects that are backed up.
	//
	// example:
	//
	// example/
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The backup progress. For example, 10000 indicates that the progress is 100%.
	//
	// example:
	//
	// 10000
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Task Report
	Report *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport `json:"Report,omitempty" xml:"Report,omitempty" type:"Struct"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: ECS files
	//
	// 	- **OSS**: OSS buckets
	//
	// 	- **NAS**: NAS file systems
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The average speed at which data is backed up. Unit: KB/s.
	//
	// example:
	//
	// 500
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// This parameter is returned only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the throttling rules. Format: `{start}{end}{bandwidth}`. Multiple throttling rules are separated with vertical bars (`{start}|{end}|{bandwidth}`). A specified time range cannot overlap with another one.
	//
	// 	- **start**: the start hour
	//
	// 	- **end**: the end hour
	//
	// 	- **bandwidth**: the bandwidth. Unit: KB/s.
	//
	// example:
	//
	// 0:24:5120
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// The time when the backup job started. This value is a UNIX timestamp. Unit: seconds.
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
	// 	- **FAILED**: The restore job has failed.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of a destination table in the Tablestore instance.
	//
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The time when the backup job was updated. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-0006******q
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetActualBytes() *int64 {
	return s.ActualBytes
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetActualFiles() *int64 {
	return s.ActualFiles
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetActualItems() *int64 {
	return s.ActualItems
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetChangeListPath() *string {
	return s.ChangeListPath
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetDestDataSourceDetail() *string {
	return s.DestDataSourceDetail
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetDestDataSourceId() *string {
	return s.DestDataSourceId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetDestSourceType() *string {
	return s.DestSourceType
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetDetail() *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	return s.Detail
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetExclude() *string {
	return s.Exclude
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetFilesDone() *int64 {
	return s.FilesDone
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetFilesTotal() *int64 {
	return s.FilesTotal
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetInclude() *string {
	return s.Include
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetItemsDone() *int64 {
	return s.ItemsDone
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetItemsTotal() *int64 {
	return s.ItemsTotal
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetJobId() *string {
	return s.JobId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetJobName() *string {
	return s.JobName
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetOptions() *string {
	return s.Options
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetOtsDetail() *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail {
	return s.OtsDetail
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetPaths() *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths {
	return s.Paths
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetPrefix() *string {
	return s.Prefix
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetReport() *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport {
	return s.Report
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetSpeed() *int64 {
	return s.Speed
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetTableName() *string {
	return s.TableName
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetActualBytes(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ActualBytes = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetActualFiles(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ActualFiles = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetActualItems(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ActualItems = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetBackupType(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetBucket(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Bucket = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetBytesDone(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.BytesDone = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetBytesTotal(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.BytesTotal = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetChangeListPath(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ChangeListPath = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetClientId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ClientId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCompleteTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CompleteTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCreateTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCreatedTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CreatedTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCrossAccountRoleName(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCrossAccountType(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetCrossAccountUserId(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetDestDataSourceDetail(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.DestDataSourceDetail = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetDestDataSourceId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.DestDataSourceId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetDestSourceType(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.DestSourceType = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetDetail(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Detail = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetErrorMessage(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetExclude(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Exclude = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetFileSystemId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.FileSystemId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetFilesDone(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.FilesDone = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetFilesTotal(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.FilesTotal = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetIdentifier(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Identifier = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetInclude(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Include = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetInstanceId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetInstanceName(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetItemsDone(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ItemsDone = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetItemsTotal(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.ItemsTotal = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetJobId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.JobId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetJobName(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.JobName = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetOptions(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Options = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetOtsDetail(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.OtsDetail = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetPaths(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Paths = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetPlanId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.PlanId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetPrefix(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Prefix = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetProgress(v int32) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Progress = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetReport(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Report = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetSourceType(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.SourceType = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetSpeed(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Speed = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetSpeedLimit(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.SpeedLimit = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetStartTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetStatus(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.Status = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetTableName(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.TableName = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetUpdatedTime(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) SetVaultId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob {
	s.VaultId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJob) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail struct {
	// The information about the remote replication failure.
	//
	// example:
	//
	// ECS.CreatingSnapshot
	DestinationNativeSnapshotErrorMessage *string `json:"DestinationNativeSnapshotErrorMessage,omitempty" xml:"DestinationNativeSnapshotErrorMessage,omitempty"`
	// The ID of the remote replication snapshot.
	//
	// example:
	//
	// s-******************
	DestinationNativeSnapshotId *string `json:"DestinationNativeSnapshotId,omitempty" xml:"DestinationNativeSnapshotId,omitempty"`
	// The progress of the remote replication.
	//
	// example:
	//
	// 85
	DestinationNativeSnapshotProgress *int32 `json:"DestinationNativeSnapshotProgress,omitempty" xml:"DestinationNativeSnapshotProgress,omitempty"`
	// The state of the remote replication.
	//
	// example:
	//
	// COMPLETE
	DestinationNativeSnapshotStatus *string `json:"DestinationNativeSnapshotStatus,omitempty" xml:"DestinationNativeSnapshotStatus,omitempty"`
	// The retention period of the remote replication backup.
	//
	// example:
	//
	// 30
	DestinationRetention *int64 `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	// The ID of the remote replication backup.
	//
	// example:
	//
	// s-******************
	DestinationSnapshotId *string `json:"DestinationSnapshotId,omitempty" xml:"DestinationSnapshotId,omitempty"`
	// The mapping between snapshots and disks.
	DiskNativeSnapshotIdList *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList `json:"DiskNativeSnapshotIdList,omitempty" xml:"DiskNativeSnapshotIdList,omitempty" type:"Struct"`
	// Indicates whether remote replication is enabled.
	//
	// example:
	//
	// true
	DoCopy *bool `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	// The ecs instance infos.
	//
	// example:
	//
	// {
	//
	//   "i-xxxxxxxx": {
	//
	//     "hostName": "test",
	//
	//     "instanceName": "test",
	//
	//     "instanceType": "ecs.c7.xlarge",
	//
	//     "osType": "linux",
	//
	//     "diskIds": [
	//
	//       "d-xxxxxxxx01",
	//
	//       "d-xxxxxxxx02"
	//
	//     ],
	//
	//     "osNameEn": "Rocky Linux 8.8 64 bit",
	//
	//     "osName": "Rocky Linux 8.8 64位",
	//
	//     "platform": "Rocky Linux"
	//
	//   }
	//
	// }
	InstanceInfos map[string]interface{} `json:"InstanceInfos,omitempty" xml:"InstanceInfos,omitempty"`
	// The ID of the backup snapshot.
	//
	// example:
	//
	// s-******************
	NativeSnapshotId *string `json:"NativeSnapshotId,omitempty" xml:"NativeSnapshotId,omitempty"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetDestinationNativeSnapshotErrorMessage() *string {
	return s.DestinationNativeSnapshotErrorMessage
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetDestinationNativeSnapshotId() *string {
	return s.DestinationNativeSnapshotId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetDestinationNativeSnapshotProgress() *int32 {
	return s.DestinationNativeSnapshotProgress
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetDestinationNativeSnapshotStatus() *string {
	return s.DestinationNativeSnapshotStatus
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetDestinationRetention() *int64 {
	return s.DestinationRetention
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetDestinationSnapshotId() *string {
	return s.DestinationSnapshotId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetDiskNativeSnapshotIdList() *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList {
	return s.DiskNativeSnapshotIdList
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetDoCopy() *bool {
	return s.DoCopy
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetInstanceInfos() map[string]interface{} {
	return s.InstanceInfos
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) GetNativeSnapshotId() *string {
	return s.NativeSnapshotId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetDestinationNativeSnapshotErrorMessage(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.DestinationNativeSnapshotErrorMessage = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetDestinationNativeSnapshotId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.DestinationNativeSnapshotId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetDestinationNativeSnapshotProgress(v int32) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.DestinationNativeSnapshotProgress = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetDestinationNativeSnapshotStatus(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.DestinationNativeSnapshotStatus = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetDestinationRetention(v int64) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.DestinationRetention = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetDestinationSnapshotId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.DestinationSnapshotId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetDiskNativeSnapshotIdList(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.DiskNativeSnapshotIdList = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetDoCopy(v bool) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.DoCopy = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetInstanceInfos(v map[string]interface{}) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.InstanceInfos = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) SetNativeSnapshotId(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail {
	s.NativeSnapshotId = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList struct {
	DiskNativeSnapshotId []*string `json:"DiskNativeSnapshotId,omitempty" xml:"DiskNativeSnapshotId,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList) GetDiskNativeSnapshotId() []*string {
	return s.DiskNativeSnapshotId
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList) SetDiskNativeSnapshotId(v []*string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList {
	s.DiskNativeSnapshotId = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobDetailDiskNativeSnapshotIdList) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail struct {
	// The names of the destination tables in the Tablestore instance.
	TableNames *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Struct"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) GetTableNames() *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames {
	return s.TableNames
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) SetTableNames(v *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail {
	s.TableNames = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames struct {
	TableName []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) GetTableName() []*string {
	return s.TableName
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) SetTableName(v []*string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames {
	s.TableName = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobOtsDetailTableNames) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths struct {
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) GetPath() []*string {
	return s.Path
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) SetPath(v []*string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths {
	s.Path = v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobPaths) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport struct {
	// List of failed files
	//
	// example:
	//
	// /temp/report/158975xxxxxx4625/r-0001hfxxxxxymsspjjtl/job-0001hfxxxxxymsspjjtl_failed.zip
	FailedFiles *string `json:"FailedFiles,omitempty" xml:"FailedFiles,omitempty"`
	// Report generation status.
	//
	// example:
	//
	// COMPLETE
	ReportTaskStatus *string `json:"ReportTaskStatus,omitempty" xml:"ReportTaskStatus,omitempty"`
	// List of skipped files
	//
	// example:
	//
	// /temp/report/158975xxxxxx4625/r-0001hfxxxxxymsspjjtl/job-0001hfxxxxxymsspjjtl_skipped.zip
	SkippedFiles *string `json:"SkippedFiles,omitempty" xml:"SkippedFiles,omitempty"`
	// List of successful files.
	//
	// example:
	//
	// /temp/report/158975xxxxxx4625/r-0001hfxxxxxymsspjjtl/job-0001hfxxxxxymsspjjtl_success.zip
	SuccessFiles *string `json:"SuccessFiles,omitempty" xml:"SuccessFiles,omitempty"`
	// List of all files. (This field is not returned for data synchronization)
	//
	// example:
	//
	// /temp/report/158975xxxxxx4625/job-0001hfxxxxxymsspjjtl/job-0001hfxxxxxymsspjjtl_total.csv
	TotalFiles *string `json:"TotalFiles,omitempty" xml:"TotalFiles,omitempty"`
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) GetFailedFiles() *string {
	return s.FailedFiles
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) GetReportTaskStatus() *string {
	return s.ReportTaskStatus
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) GetSkippedFiles() *string {
	return s.SkippedFiles
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) GetSuccessFiles() *string {
	return s.SuccessFiles
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) GetTotalFiles() *string {
	return s.TotalFiles
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) SetFailedFiles(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport {
	s.FailedFiles = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) SetReportTaskStatus(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport {
	s.ReportTaskStatus = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) SetSkippedFiles(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport {
	s.SkippedFiles = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) SetSuccessFiles(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport {
	s.SuccessFiles = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) SetTotalFiles(v string) *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport {
	s.TotalFiles = &v
	return s
}

func (s *DescribeBackupJobs2ResponseBodyBackupJobsBackupJobReport) Validate() error {
	return dara.Validate(s)
}
