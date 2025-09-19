// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeRestoreJobsResponseBodyPageInfo) *DescribeRestoreJobsResponseBody
	GetPageInfo() *DescribeRestoreJobsResponseBodyPageInfo
	SetRequestId(v string) *DescribeRestoreJobsResponseBody
	GetRequestId() *string
	SetRestoreJobs(v []*DescribeRestoreJobsResponseBodyRestoreJobs) *DescribeRestoreJobsResponseBody
	GetRestoreJobs() []*DescribeRestoreJobsResponseBodyRestoreJobs
}

type DescribeRestoreJobsResponseBody struct {
	// The pagination information.
	PageInfo *DescribeRestoreJobsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0ED92280-4363-57D3-A4D3-4D3FBC99B29F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the restoration tasks.
	RestoreJobs []*DescribeRestoreJobsResponseBodyRestoreJobs `json:"RestoreJobs,omitempty" xml:"RestoreJobs,omitempty" type:"Repeated"`
}

func (s DescribeRestoreJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobsResponseBody) GetPageInfo() *DescribeRestoreJobsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeRestoreJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreJobsResponseBody) GetRestoreJobs() []*DescribeRestoreJobsResponseBodyRestoreJobs {
	return s.RestoreJobs
}

func (s *DescribeRestoreJobsResponseBody) SetPageInfo(v *DescribeRestoreJobsResponseBodyPageInfo) *DescribeRestoreJobsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeRestoreJobsResponseBody) SetRequestId(v string) *DescribeRestoreJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBody) SetRestoreJobs(v []*DescribeRestoreJobsResponseBodyRestoreJobs) *DescribeRestoreJobsResponseBody {
	s.RestoreJobs = v
	return s
}

func (s *DescribeRestoreJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreJobsResponseBodyPageInfo struct {
	// The number of restoration tasks returned on the current page.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of restoration tasks returned.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRestoreJobsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) SetCount(v int32) *DescribeRestoreJobsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeRestoreJobsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) SetPageSize(v int32) *DescribeRestoreJobsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) SetTotalCount(v int32) *DescribeRestoreJobsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreJobsResponseBodyRestoreJobs struct {
	// The size of backup data. Unit: bytes.
	//
	// example:
	//
	// 20
	ActualBytes *int64 `json:"ActualBytes,omitempty" xml:"ActualBytes,omitempty"`
	// The total size of data that is restored. Unit: bytes.
	//
	// example:
	//
	// 20
	BytesDone *int64 `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	// The total size of data that you want to restore. Unit: bytes.
	//
	// example:
	//
	// 20
	BytesTotal *int64 `json:"BytesTotal,omitempty" xml:"BytesTotal,omitempty"`
	// The ID of the anti-ransomware agent that is used to perform the restoration task.
	//
	// example:
	//
	// c-000frxwusjauhp9ajpu6
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The timestamp when the restoration task is complete. Unit: milliseconds.
	//
	// example:
	//
	// 1583289054000
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The timestamp when the restoration task is created. Unit: milliseconds.
	//
	// example:
	//
	// 1583289052000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The duration of the restoration task. Unit: seconds.
	//
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of the restoration tasks on which errors occur.
	//
	// example:
	//
	// 0
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// The name of the CSV file. The CSV file contains the files that fail to be restored.
	//
	// example:
	//
	// s-000f4wxm8f7gur6g2otm.csv
	ErrorFile *string `json:"ErrorFile,omitempty" xml:"ErrorFile,omitempty"`
	// The URL to download the CSV file. The CSV file contains the files that fail to be restored.
	//
	// example:
	//
	// ["/home/user"]
	ErrorFileUrl *string `json:"ErrorFileUrl,omitempty" xml:"ErrorFileUrl,omitempty"`
	// The error code that is returned for the restoration task.
	//
	// example:
	//
	// NONE
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// The timestamp when the in-progress restoration task is expected to be complete. Unit: seconds.
	//
	// example:
	//
	// 1583299054
	Eta *int64 `json:"Eta,omitempty" xml:"Eta,omitempty"`
	// The directory excluded from the anti-ransomware policy. The value is the directory that you specify to skip protection when you create the anti-ransomware policy.
	//
	// example:
	//
	// ["/home/user"]
	Excludes *string `json:"Excludes,omitempty" xml:"Excludes,omitempty"`
	// The return value of the restoration task.
	//
	// example:
	//
	// 0
	ExitCode *string `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the restoration task is created.
	//
	// example:
	//
	// 2021-04-25T19:11Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the restoration task is updated.
	//
	// example:
	//
	// 2021-04-25T19:11Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The directory in which the restored file is stored. The value is the directory that you specify for protection when you create the anti-ransomware policy
	//
	// example:
	//
	// ["/root/disk-uuid-test","/root/install.sh"]
	Includes *string `json:"Includes,omitempty" xml:"Includes,omitempty"`
	// The ID of the server whose data you want to restore.
	//
	// example:
	//
	// i-bp12xnvdax6307gw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server whose data you want to restore.
	//
	// example:
	//
	// win2012-01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server whose data you want to restore.
	//
	// example:
	//
	// 1.1.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The internal IP address of the server whose data you want to restore.
	//
	// example:
	//
	// 2.1.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The number of files that are restored.
	//
	// example:
	//
	// 0
	ItemsDone *int64 `json:"ItemsDone,omitempty" xml:"ItemsDone,omitempty"`
	// The total number of files that need to be restored.
	//
	// example:
	//
	// 0
	ItemsTotal *int64 `json:"ItemsTotal,omitempty" xml:"ItemsTotal,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The progress of the restoration task in percentage.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED92280-4363-57D3-A4D3-4D3FBC99B29F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the restoration task.
	//
	// example:
	//
	// r-000gmcypy5dyf9ey3uv7
	RestoreId *string `json:"RestoreId,omitempty" xml:"RestoreId,omitempty"`
	// The name of the restoration task.
	//
	// example:
	//
	// Restore
	RestoreName *string `json:"RestoreName,omitempty" xml:"RestoreName,omitempty"`
	// The type of the file that is restored. Valid values:
	//
	// 	- **ECS_FILE**: files on Elastic Compute Service (ECS) instances
	//
	// 	- **FILE**: files on servers in data centers
	//
	// example:
	//
	// ECS_FILE
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The hash value of the snapshot that stores backup data when the data is backed up.
	//
	// example:
	//
	// a3992de83f529b844135fe795d949181735a7d20e0ac8539485c61b7983e618f
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The hash value ID of the snapshot that stores backup data when the data is backed up.
	//
	// example:
	//
	// s-000gmcypy5dy54e39yny
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The version of the backup data.
	//
	// example:
	//
	// 2020-03-03 18:00
	SnapshotVersion *string `json:"SnapshotVersion,omitempty" xml:"SnapshotVersion,omitempty"`
	// The restored content.
	//
	// example:
	//
	// ["/home/admin","\\\\\\\\servername\\\\sharename"]
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the anti-ransomware agent that is used to back up data.
	//
	// example:
	//
	// c-000gmcypy5dyf9ey3uv7
	SourceClientId *string `json:"SourceClientId,omitempty" xml:"SourceClientId,omitempty"`
	// The speed at which data is restored. Unit: byte/s.
	//
	// example:
	//
	// 25766558
	Speed *int64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The status of the restoration task. Valid values:
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **COMPLETE**: The task is complete.
	//
	// 	- **FAILED**: The task fails.
	//
	// 	- **CANCELING**: The task is being canceled.
	//
	// 	- **CANCELED**: The task is canceled.
	//
	// 	- **PARTIAL_COMPLETE**: The task is partially successful.
	//
	// 	- **CREATED**: The task was created but is not run.
	//
	// 	- **EXPIRED**: The task is not updated.
	//
	// 	- **QUEUED**: The task is waiting to be run.
	//
	// 	- **CLIENT_DELETED**: The task fails because the anti-ransomware agent is uninstalled.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The folder to which the backup data is restored. After you create the restoration task, the backup data is restored to the specified folder.
	//
	// example:
	//
	// /home
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The timestamp when the restoration task was last updated. Unit: milliseconds.
	//
	// example:
	//
	// 1583289054000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The UUID of the server whose data you want to restore.
	//
	// example:
	//
	// 6E3DABB6-3F6A-40DB-9492-2C8B59C****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the backup vault in which the backup data is stored.
	//
	// example:
	//
	// v-000b0v0jqzmse2yz06zw
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	// The ID of the region where the backup vault resides.
	//
	// example:
	//
	// cn-hangzhou
	VaultRegionId *string `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
}

func (s DescribeRestoreJobsResponseBodyRestoreJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobsResponseBodyRestoreJobs) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetActualBytes() *int64 {
	return s.ActualBytes
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetBytesTotal() *int64 {
	return s.BytesTotal
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetErrorFile() *string {
	return s.ErrorFile
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetErrorFileUrl() *string {
	return s.ErrorFileUrl
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetErrorType() *string {
	return s.ErrorType
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetEta() *int64 {
	return s.Eta
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetExcludes() *string {
	return s.Excludes
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetExitCode() *string {
	return s.ExitCode
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetIncludes() *string {
	return s.Includes
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetItemsDone() *int64 {
	return s.ItemsDone
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetItemsTotal() *int64 {
	return s.ItemsTotal
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetMessage() *string {
	return s.Message
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetPercentage() *int32 {
	return s.Percentage
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetRestoreId() *string {
	return s.RestoreId
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetRestoreName() *string {
	return s.RestoreName
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetRestoreType() *string {
	return s.RestoreType
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetSnapshotVersion() *string {
	return s.SnapshotVersion
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetSource() *string {
	return s.Source
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetSourceClientId() *string {
	return s.SourceClientId
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetSpeed() *int64 {
	return s.Speed
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetStatus() *string {
	return s.Status
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetTarget() *string {
	return s.Target
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) GetVaultRegionId() *string {
	return s.VaultRegionId
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetActualBytes(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ActualBytes = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetBytesDone(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.BytesDone = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetBytesTotal(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.BytesTotal = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetClientId(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ClientId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetCompleteTime(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.CompleteTime = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetCreatedTime(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.CreatedTime = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetDuration(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Duration = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetErrorCount(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ErrorCount = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetErrorFile(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ErrorFile = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetErrorFileUrl(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ErrorFileUrl = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetErrorType(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ErrorType = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetEta(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Eta = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetExcludes(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Excludes = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetExitCode(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ExitCode = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetGmtCreate(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetGmtModified(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.GmtModified = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetIncludes(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Includes = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetInstanceId(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.InstanceId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetInstanceName(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.InstanceName = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetInternetIp(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.InternetIp = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetIntranetIp(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.IntranetIp = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetItemsDone(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ItemsDone = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetItemsTotal(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.ItemsTotal = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetMessage(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Message = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetPercentage(v int32) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Percentage = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetRequestId(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetRestoreId(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.RestoreId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetRestoreName(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.RestoreName = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetRestoreType(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.RestoreType = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetSnapshotHash(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.SnapshotHash = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetSnapshotId(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.SnapshotId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetSnapshotVersion(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.SnapshotVersion = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetSource(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Source = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetSourceClientId(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.SourceClientId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetSpeed(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Speed = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetStatus(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Status = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetTarget(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Target = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetUpdatedTime(v int64) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetUuid(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.Uuid = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetVaultId(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.VaultId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) SetVaultRegionId(v string) *DescribeRestoreJobsResponseBodyRestoreJobs {
	s.VaultRegionId = &v
	return s
}

func (s *DescribeRestoreJobsResponseBodyRestoreJobs) Validate() error {
	return dara.Validate(s)
}
