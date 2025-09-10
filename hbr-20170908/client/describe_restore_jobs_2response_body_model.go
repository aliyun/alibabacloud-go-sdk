// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreJobs2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRestoreJobs2ResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeRestoreJobs2ResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeRestoreJobs2ResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRestoreJobs2ResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRestoreJobs2ResponseBody
	GetRequestId() *string
	SetRestoreJobs(v *DescribeRestoreJobs2ResponseBodyRestoreJobs) *DescribeRestoreJobs2ResponseBody
	GetRestoreJobs() *DescribeRestoreJobs2ResponseBodyRestoreJobs
	SetSuccess(v bool) *DescribeRestoreJobs2ResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeRestoreJobs2ResponseBody
	GetTotalCount() *int32
}

type DescribeRestoreJobs2ResponseBody struct {
	// The response status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried restore jobs.
	//
	// example:
	//
	// {\\"RestoreJob\\": []}
	RestoreJobs *DescribeRestoreJobs2ResponseBodyRestoreJobs `json:"RestoreJobs,omitempty" xml:"RestoreJobs,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRestoreJobs2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRestoreJobs2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRestoreJobs2ResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreJobs2ResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreJobs2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreJobs2ResponseBody) GetRestoreJobs() *DescribeRestoreJobs2ResponseBodyRestoreJobs {
	return s.RestoreJobs
}

func (s *DescribeRestoreJobs2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRestoreJobs2ResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRestoreJobs2ResponseBody) SetCode(v string) *DescribeRestoreJobs2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetMessage(v string) *DescribeRestoreJobs2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetPageNumber(v int32) *DescribeRestoreJobs2ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetPageSize(v int32) *DescribeRestoreJobs2ResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetRequestId(v string) *DescribeRestoreJobs2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetRestoreJobs(v *DescribeRestoreJobs2ResponseBodyRestoreJobs) *DescribeRestoreJobs2ResponseBody {
	s.RestoreJobs = v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetSuccess(v bool) *DescribeRestoreJobs2ResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) SetTotalCount(v int32) *DescribeRestoreJobs2ResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreJobs2ResponseBodyRestoreJobs struct {
	RestoreJob []*DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob `json:"RestoreJob,omitempty" xml:"RestoreJob,omitempty" type:"Repeated"`
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobs) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobs) GetRestoreJob() []*DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	return s.RestoreJob
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobs) SetRestoreJob(v []*DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) *DescribeRestoreJobs2ResponseBodyRestoreJobs {
	s.RestoreJob = v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobs) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob struct {
	// The actual amount of data that is restored after duplicates are removed. Unit: bytes.
	//
	// example:
	//
	// 600
	ActualBytes *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// This parameter is valid only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the actual number of objects that are restored by the restore job.
	//
	// example:
	//
	// 6
	ActualItems *int64 `json:"ActualItems,omitempty" xml:"ActualItems,omitempty"`
	// The amount of data that was restored. Unit: bytes.
	//
	// example:
	//
	// 800
	BytesDone *int64 `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	// The total amount of data that was backed up from the data source. Unit: bytes.
	//
	// example:
	//
	// 1000
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// The ID of the client group used for restoration.
	//
	// example:
	//
	// cl-000******hp6
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the restore job was completed. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the restore job was created. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The name of the Resource Access Management (RAM) role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Indicates whether data is backed up within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
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
	// 158975xxxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The files that failed to be restored.
	//
	// example:
	//
	// "test.csv"
	ErrorFile *string `json:"ErrorFile,omitempty" xml:"ErrorFile,omitempty"`
	// The error message that is returned for the restore job.
	//
	// example:
	//
	// PARTIAL_COMPLETE
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// This parameter is valid only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the paths to the files that are excluded from the restore job. The value can be up to 255 characters in length.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// The time when the restore job expires.
	//
	// example:
	//
	// 1634714531
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The details about the VMware failback task.
	//
	// example:
	//
	// {"cpu":4,"extra":"{\\"restoreVMNamePrefix\\":\\"627-\\",\\"dataCenterName\\":\\"SDDC-Datacenter\\",\\"dataStoreId\\":\\"datastore-50\\",\\"folderId\\":\\"group-v49\\",\\"resourcePoolId\\":\\"resgroup-46\\",\\"locationName\\":\\"vcenter.pc-uf600arcwi9td3eyj641.acvs.aliyuncs.com/SDDC-Datacenter/Workloads\\",\\"computeResourceName\\":\\"SDDC-Datacenter/Default_c-uf600arcwi9td3eyj640\\",\\"dataStoreName\\":\\"Default_c-uf600arcwi9td3eyj640/WorkloadDatastore\\",\\"networkMoReference\\":\\"DistributedVirtualPortgroup:dvportgroup-1001\\",\\"useHotAdd\\":false}","instanceId":"i-2vc357i2eannmmotcagz","memoryInMB":8192,"serverId":"0fdc0c86-eb92-4e05-91ab-eeaf9fb6ad01","uefiBoot":false}
	FailbackDetail *string `json:"FailbackDetail,omitempty" xml:"FailbackDetail,omitempty"`
	// The paths to the files that are included in the restore job.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// This parameter is valid only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the number of restored objects.
	//
	// example:
	//
	// 8
	ItemsDone *int64 `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	// This parameter is valid only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates the total number of objects in the data source.
	//
	// example:
	//
	// 10
	ItemsTotal *int64 `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	// The amount of data that was restored. Unit: bytes. This parameter is valid only if the StorageClass parameter is set to ARCHIVE. The minimum billable size of the data stored at the Archive tier is 1 MB.
	//
	// example:
	//
	// 1048576
	MeteringBytesDone *int64 `json:"MeteringBytesDone,omitempty" xml:"MeteringBytesDone,omitempty"`
	// The total amount of data that was backed up from the data source. Unit: bytes. This parameter is valid only if the StorageClass parameter is set to ARCHIVE. The minimum billable size of the data stored at the Archive tier is 1 MB.
	//
	// example:
	//
	// 1048576
	MeteringBytesTotal *int64 `json:"MeteringBytesTotal,omitempty" xml:"MeteringBytesTotal,omitempty"`
	// This parameter is valid only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter indicates whether Windows Volume Shadow Copy Service (VSS) is used to define a restoration path.
	//
	// 	- This parameter is available only for Windows ECS instances.
	//
	// 	- If data changes occur in the backup source, the source data must be the same as the data to be backed up before you can set this parameter to `["UseVSS":true]`.
	//
	// 	- If you use VSS, you cannot restore data from multiple directories.
	//
	// example:
	//
	// {"UseVSS":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The details about the Tablestore instance.
	OtsDetail *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty" type:"Struct"`
	// The ID of the parent job.
	//
	// example:
	//
	// r-0003hd2an3x4dyv0l18b
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The progress of the restore job. Valid values: [0,10000]. For example, 10000 indicates that the progress is 100%.
	//
	// example:
	//
	// 10000
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The report of the restore job.
	Report *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport `json:"Report,omitempty" xml:"Report,omitempty" type:"Struct"`
	// The ID of the restore job.
	//
	// example:
	//
	// r-0003*****8a
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	// The type of the restore job.
	//
	// example:
	//
	// ECS_FILE
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The hash value of the backup snapshot.
	//
	// example:
	//
	// f2fe...
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The ID of the snapshot used for restoration.
	//
	// example:
	//
	// s-0002******ga88
	SnapshotId       *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: ECS files
	//
	// 	- **OSS**: Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: Apsara File Storage NAS (NAS) file systems
	//
	// 	- **OTS_TABLE**: Tablestore instances
	//
	// 	- **UDM_ECS**: ECS instances
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
	// The time when the restore job started. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the restore job. Valid values:
	//
	// 	- **COMPLETE**: The job is completed.
	//
	// 	- **PARTIAL_COMPLETE**: The job is partially completed.
	//
	// 	- **FAILED**: The job failed.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage class of the backup data. Valid values:
	//
	// 	- **STANDARD**
	//
	// 	- **ARCHIVE**
	//
	// example:
	//
	// STANDARD
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// The name of the destination OSS bucket. This parameter is returned only for OSS buckets.
	//
	// example:
	//
	// target-bucket
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	// The ID of the destination client.
	//
	// example:
	//
	// c-000e*****397
	TargetClientId *string `json:"TargetClientId,omitempty" xml:"TargetClientId,omitempty"`
	// The time when the file system was created. This parameter is returned only for NAS file systems.
	//
	// example:
	//
	// 1634714531
	TargetCreateTime *int64 `json:"TargetCreateTime,omitempty" xml:"TargetCreateTime,omitempty"`
	// The ID of the destination data source.
	//
	// example:
	//
	// ds-000*****997
	TargetDataSourceId *string `json:"TargetDataSourceId,omitempty" xml:"TargetDataSourceId,omitempty"`
	// The ID of the destination NAS file system. This parameter is returned only for NAS file systems.
	//
	// example:
	//
	// 0be9****9c9
	TargetFileSystemId *string `json:"TargetFileSystemId,omitempty" xml:"TargetFileSystemId,omitempty"`
	// The ID of the destination instance for the restore job.
	//
	// example:
	//
	// i-2ze3m7ktcgw******cs
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The name of the destination Tablestore instance.
	//
	// example:
	//
	// instancename
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	// The destination file path of the restore job.
	//
	// example:
	//
	// "D:\\\\rebk"
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The prefix of the objects that are restored. This parameter is returned only for OSS buckets.
	//
	// example:
	//
	// "/target"
	TargetPrefix *string `json:"TargetPrefix,omitempty" xml:"TargetPrefix,omitempty"`
	// The name of the destination table in the Tablestore instance.
	//
	// example:
	//
	// tablename
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	// The time when the Tablestore instance was backed up. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642560261
	TargetTime *int64 `json:"TargetTime,omitempty" xml:"TargetTime,omitempty"`
	// The details about Elastic Compute Service (ECS) instance backup.
	//
	// example:
	//
	// {\\"doCopy\\":true,\\"doBackup\\":false,\\"instanceName\\":\\"instance example\\",\\"appConsistent\\":false,\\"destinationRegionId\\":\\"cn-shanghai\\",\\"enableFsFreeze\\":true,\\"osNameEn\\":\\"Windows Server  2019 Data Center Edition 64bit Chinese Edition\\",\\"osName\\":\\"Windows Server  2019 Data Center Edition 64bit Chinese Edition\\",\\"diskIdList\\":[],\\"backupVaultId\\":\\"\\",\\"snapshotGroup\\":true,\\"destinationRetention\\":35,\\"platform\\":\\"Windows Server 2012\\",\\"timeoutInSeconds\\":60,\\"backupRetention\\":1,\\"osType\\":\\"windows\\",\\"preScriptPath\\":\\"\\",\\"postScriptPath\\":\\"\\",\\"enableWriters\\":true,\\"ecsDeleted\\":false}
	UdmDetail *string `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty"`
	// The time when the restore job was updated. This value is a UNIX timestamp. Unit: seconds.
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

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetActualBytes() *int64 {
	return s.ActualBytes
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetActualItems() *int64 {
	return s.ActualItems
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetErrorFile() *string {
	return s.ErrorFile
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetExclude() *string {
	return s.Exclude
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetFailbackDetail() *string {
	return s.FailbackDetail
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetInclude() *string {
	return s.Include
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetItemsDone() *int64 {
	return s.ItemsDone
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetItemsTotal() *int64 {
	return s.ItemsTotal
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetMeteringBytesDone() *int64 {
	return s.MeteringBytesDone
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetMeteringBytesTotal() *int64 {
	return s.MeteringBytesTotal
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetOptions() *string {
	return s.Options
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetOtsDetail() *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail {
	return s.OtsDetail
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetReport() *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport {
	return s.Report
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetRestoreId() *string {
	return s.RestoreId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetRestoreType() *string {
	return s.RestoreType
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetSpeed() *int64 {
	return s.Speed
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetStatus() *string {
	return s.Status
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetBucket() *string {
	return s.TargetBucket
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetClientId() *string {
	return s.TargetClientId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetCreateTime() *int64 {
	return s.TargetCreateTime
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetDataSourceId() *string {
	return s.TargetDataSourceId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetFileSystemId() *string {
	return s.TargetFileSystemId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetPath() *string {
	return s.TargetPath
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetPrefix() *string {
	return s.TargetPrefix
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetTargetTime() *int64 {
	return s.TargetTime
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetUdmDetail() *string {
	return s.UdmDetail
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetActualBytes(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ActualBytes = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetActualItems(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ActualItems = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetBytesDone(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.BytesDone = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetBytesTotal(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.BytesTotal = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetClusterId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCompleteTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CompleteTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCreatedTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CreatedTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCrossAccountRoleName(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCrossAccountType(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetCrossAccountUserId(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetErrorFile(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ErrorFile = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetErrorMessage(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetExclude(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Exclude = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetExpireTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetFailbackDetail(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.FailbackDetail = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetInclude(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Include = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetItemsDone(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ItemsDone = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetItemsTotal(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ItemsTotal = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetMeteringBytesDone(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.MeteringBytesDone = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetMeteringBytesTotal(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.MeteringBytesTotal = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetOptions(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Options = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetOtsDetail(v *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.OtsDetail = v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetParentId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.ParentId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetProgress(v int32) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Progress = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetReport(v *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Report = v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetRestoreId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.RestoreId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetRestoreType(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.RestoreType = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSnapshotHash(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSnapshotId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.SnapshotId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSourceInstanceId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.SourceInstanceId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSourceType(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.SourceType = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetSpeed(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetStartTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetStatus(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.Status = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetStorageClass(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.StorageClass = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetBucket(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetBucket = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetClientId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetClientId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetCreateTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetCreateTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetDataSourceId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetDataSourceId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetFileSystemId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetFileSystemId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetInstanceId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetInstanceId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetInstanceName(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetInstanceName = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetPath(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetPath = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetPrefix(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetPrefix = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetTableName(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetTableName = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetTargetTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.TargetTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetUdmDetail(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.UdmDetail = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetUpdatedTime(v int64) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) SetVaultId(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob {
	s.VaultId = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJob) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail struct {
	// The number of channels processed by each Tablestore restore job.
	//
	// example:
	//
	// 2
	BatchChannelCount *int32 `json:"BatchChannelCount,omitempty" xml:"BatchChannelCount,omitempty"`
	// Indicates whether the existing Tablestore restore job was overwritten.
	//
	// example:
	//
	// false
	OverwriteExisting *bool `json:"OverwriteExisting,omitempty" xml:"OverwriteExisting,omitempty"`
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) GetBatchChannelCount() *int32 {
	return s.BatchChannelCount
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) GetOverwriteExisting() *bool {
	return s.OverwriteExisting
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) SetBatchChannelCount(v int32) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail {
	s.BatchChannelCount = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) SetOverwriteExisting(v bool) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail {
	s.OverwriteExisting = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobOtsDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport struct {
	// The files that failed to be executed.
	//
	// example:
	//
	// /temp/report/158975xxxxxx4625/r-0001hfxxxxxymsspjjtl/job-0001hfxxxxxymsspjjtl_failed.zip
	FailedFiles *string `json:"FailedFiles,omitempty" xml:"FailedFiles,omitempty"`
	// The status of the report generation.
	//
	// example:
	//
	// COMPLETE
	ReportTaskStatus *string `json:"ReportTaskStatus,omitempty" xml:"ReportTaskStatus,omitempty"`
	// The skipped files.
	//
	// example:
	//
	// /temp/report/158975xxxxxx4625/r-0001hfxxxxxymsspjjtl/job-0001hfxxxxxymsspjjtl_skipped.zip
	SkippedFiles *string `json:"SkippedFiles,omitempty" xml:"SkippedFiles,omitempty"`
	// The files that are successfully executed.
	//
	// example:
	//
	// /temp/report/158975xxxxxx4625/r-0001hfxxxxxymsspjjtl/job-0001hfxxxxxymsspjjtl_success.zip
	SuccessFiles *string `json:"SuccessFiles,omitempty" xml:"SuccessFiles,omitempty"`
	// The full files that are restored based on the file list.
	//
	// example:
	//
	// /temp/report/158975xxxxxx4625/job-0001hfxxxxxymsspjjtl/job-0001hfxxxxxymsspjjtl_total.csv
	TotalFiles *string `json:"TotalFiles,omitempty" xml:"TotalFiles,omitempty"`
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) GetFailedFiles() *string {
	return s.FailedFiles
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) GetReportTaskStatus() *string {
	return s.ReportTaskStatus
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) GetSkippedFiles() *string {
	return s.SkippedFiles
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) GetSuccessFiles() *string {
	return s.SuccessFiles
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) GetTotalFiles() *string {
	return s.TotalFiles
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) SetFailedFiles(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport {
	s.FailedFiles = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) SetReportTaskStatus(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport {
	s.ReportTaskStatus = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) SetSkippedFiles(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport {
	s.SkippedFiles = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) SetSuccessFiles(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport {
	s.SuccessFiles = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) SetTotalFiles(v string) *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport {
	s.TotalFiles = &v
	return s
}

func (s *DescribeRestoreJobs2ResponseBodyRestoreJobsRestoreJobReport) Validate() error {
	return dara.Validate(s)
}
