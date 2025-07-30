// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBackupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterBackups(v []*DescribeClusterBackupListResponseBodyClusterBackups) *DescribeClusterBackupListResponseBody
	GetClusterBackups() []*DescribeClusterBackupListResponseBodyClusterBackups
	SetFreeSize(v int64) *DescribeClusterBackupListResponseBody
	GetFreeSize() *int64
	SetFullStorageSize(v int64) *DescribeClusterBackupListResponseBody
	GetFullStorageSize() *int64
	SetLogStorageSize(v int64) *DescribeClusterBackupListResponseBody
	GetLogStorageSize() *int64
	SetMaxResults(v int32) *DescribeClusterBackupListResponseBody
	GetMaxResults() *int32
	SetPageNumber(v int32) *DescribeClusterBackupListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterBackupListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeClusterBackupListResponseBody
	GetRequestId() *string
}

type DescribeClusterBackupListResponseBody struct {
	// The backup sets of the instance. A backup contains the backup sets of all shards in the instance.
	ClusterBackups []*DescribeClusterBackupListResponseBodyClusterBackups `json:"ClusterBackups,omitempty" xml:"ClusterBackups,omitempty" type:"Repeated"`
	// This parameter does not take effect. Ignore this parameter.
	//
	// example:
	//
	// 100000
	FreeSize *int64 `json:"FreeSize,omitempty" xml:"FreeSize,omitempty"`
	// The size of the full backup file of the instance. Unit: bytes. Full backups originate from scheduled backups, manual backups, and backups generated during cache analysis.
	//
	// >  The value of this parameter is independent of the number and size of returned backup sets. Instead, it represents the size of all valid full backups of the instance.
	//
	// example:
	//
	// 1000
	FullStorageSize *int64 `json:"FullStorageSize,omitempty" xml:"FullStorageSize,omitempty"`
	// The size of the log backup file of the instance. Unit: bytes. This parameter is valid only when flashback is enabled.
	//
	// >  The value of this parameter is independent of the number and size of returned backup sets. Instead, it represents the size of all valid log backups of the instance.
	//
	// example:
	//
	// 5000
	LogStorageSize *int64 `json:"LogStorageSize,omitempty" xml:"LogStorageSize,omitempty"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C009DA42-3B19-5B81-963D-1509DE2408DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterBackupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupListResponseBody) GetClusterBackups() []*DescribeClusterBackupListResponseBodyClusterBackups {
	return s.ClusterBackups
}

func (s *DescribeClusterBackupListResponseBody) GetFreeSize() *int64 {
	return s.FreeSize
}

func (s *DescribeClusterBackupListResponseBody) GetFullStorageSize() *int64 {
	return s.FullStorageSize
}

func (s *DescribeClusterBackupListResponseBody) GetLogStorageSize() *int64 {
	return s.LogStorageSize
}

func (s *DescribeClusterBackupListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeClusterBackupListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterBackupListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterBackupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterBackupListResponseBody) SetClusterBackups(v []*DescribeClusterBackupListResponseBodyClusterBackups) *DescribeClusterBackupListResponseBody {
	s.ClusterBackups = v
	return s
}

func (s *DescribeClusterBackupListResponseBody) SetFreeSize(v int64) *DescribeClusterBackupListResponseBody {
	s.FreeSize = &v
	return s
}

func (s *DescribeClusterBackupListResponseBody) SetFullStorageSize(v int64) *DescribeClusterBackupListResponseBody {
	s.FullStorageSize = &v
	return s
}

func (s *DescribeClusterBackupListResponseBody) SetLogStorageSize(v int64) *DescribeClusterBackupListResponseBody {
	s.LogStorageSize = &v
	return s
}

func (s *DescribeClusterBackupListResponseBody) SetMaxResults(v int32) *DescribeClusterBackupListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeClusterBackupListResponseBody) SetPageNumber(v int32) *DescribeClusterBackupListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterBackupListResponseBody) SetPageSize(v int32) *DescribeClusterBackupListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterBackupListResponseBody) SetRequestId(v string) *DescribeClusterBackupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterBackupListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterBackupListResponseBodyClusterBackups struct {
	// The backup sets of all shards in the instance.
	Backups []*DescribeClusterBackupListResponseBodyClusterBackupsBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Repeated"`
	// The end time of the backup.
	//
	// example:
	//
	// 2024-01-10T17:21:55Z
	ClusterBackupEndTime *string `json:"ClusterBackupEndTime,omitempty" xml:"ClusterBackupEndTime,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// cb-zmdqj2m3xyxjtdt0
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The backup mode.
	//
	// example:
	//
	// Automated
	ClusterBackupMode *string `json:"ClusterBackupMode,omitempty" xml:"ClusterBackupMode,omitempty"`
	// The size of the backup set.
	//
	// example:
	//
	// 2048
	ClusterBackupSize *string `json:"ClusterBackupSize,omitempty" xml:"ClusterBackupSize,omitempty"`
	// The start time of the backup.
	//
	// example:
	//
	// 2024-01-10T17:21:25Z
	ClusterBackupStartTime *string `json:"ClusterBackupStartTime,omitempty" xml:"ClusterBackupStartTime,omitempty"`
	// The status of the backup set.
	//
	// 	- OK
	//
	// 	- RUNNING
	//
	// 	- Failed
	//
	// example:
	//
	// OK
	ClusterBackupStatus *string `json:"ClusterBackupStatus,omitempty" xml:"ClusterBackupStatus,omitempty"`
	ExpectExpireTime    *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// Indicates whether the backup set is valid. A value of 0 indicates that shard-level backups failed or have not been completed.
	//
	// example:
	//
	// 1
	IsAvail *int32 `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	// The backup progress. The system displays only the progress of running backup tasks.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The memory size of a single shard during a full backup. Unit: MB.
	//
	// example:
	//
	// 1024
	ShardClassMemory *int32 `json:"ShardClassMemory,omitempty" xml:"ShardClassMemory,omitempty"`
}

func (s DescribeClusterBackupListResponseBodyClusterBackups) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupListResponseBodyClusterBackups) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetBackups() []*DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	return s.Backups
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetClusterBackupEndTime() *string {
	return s.ClusterBackupEndTime
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetClusterBackupId() *string {
	return s.ClusterBackupId
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetClusterBackupMode() *string {
	return s.ClusterBackupMode
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetClusterBackupSize() *string {
	return s.ClusterBackupSize
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetClusterBackupStartTime() *string {
	return s.ClusterBackupStartTime
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetClusterBackupStatus() *string {
	return s.ClusterBackupStatus
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetIsAvail() *int32 {
	return s.IsAvail
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetProgress() *string {
	return s.Progress
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) GetShardClassMemory() *int32 {
	return s.ShardClassMemory
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetBackups(v []*DescribeClusterBackupListResponseBodyClusterBackupsBackups) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.Backups = v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetClusterBackupEndTime(v string) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.ClusterBackupEndTime = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetClusterBackupId(v string) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.ClusterBackupId = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetClusterBackupMode(v string) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.ClusterBackupMode = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetClusterBackupSize(v string) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.ClusterBackupSize = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetClusterBackupStartTime(v string) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.ClusterBackupStartTime = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetClusterBackupStatus(v string) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.ClusterBackupStatus = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetExpectExpireTime(v string) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetIsAvail(v int32) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.IsAvail = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetProgress(v string) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.Progress = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) SetShardClassMemory(v int32) *DescribeClusterBackupListResponseBodyClusterBackups {
	s.ShardClassMemory = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackups) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterBackupListResponseBodyClusterBackupsBackups struct {
	// The public download URL of the backup file.
	//
	// example:
	//
	// http://rdsbakbucket-huhehaote-v2.oss-cn-huhehaote.aliyuncs.com/custins424747958/hins100322105_data_20240110012135.rdb
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// The end time of the backup. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T17:21:57
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The ID of the backup file.
	//
	// example:
	//
	// 514645788
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The internal download URL of the backup file.
	//
	// >  You can use this URL to download the backup file from an Elastic Compute Service (ECS) instance that is connected to the Tair (Redis OSS-compatible) instance. The ECS instance must reside in the same virtual private cloud (VPC) as the Tair (Redis OSS-compatible) instance.
	//
	// example:
	//
	// http://rdsbakbucket-huhehaote-v2.oss-cn-huhehaote-internal.aliyuncs.com/custins424747958/hins100322105_data_20240110012135.rdb
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// The name of the backup.
	//
	// example:
	//
	// hins100322105_data_20240110012135.rdb
	BackupName *string `json:"BackupName,omitempty" xml:"BackupName,omitempty"`
	// The size of the backup file. Unit: bytes.
	//
	// example:
	//
	// 1024
	BackupSize *string `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T17:21:30Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The status of the backup. Valid values:
	//
	// 	- **OK**
	//
	// 	- **ERROR**
	//
	// example:
	//
	// OK
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The database engine. The return value is **redis**.
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The additional information.
	ExtraInfo *DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	// The instance name.
	//
	// example:
	//
	// hins100322105_data_20240108012127.rdb
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the backup set is available. Valid values:
	//
	// 	- **0**: unavailable
	//
	// 	- **1**: available
	//
	// example:
	//
	// 1
	IsAvail *string `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	// This parameter does not take effect. Ignore this parameter.
	//
	// example:
	//
	// null
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
}

func (s DescribeClusterBackupListResponseBodyClusterBackupsBackups) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupListResponseBodyClusterBackupsBackups) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetBackupDownloadURL() *string {
	return s.BackupDownloadURL
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetBackupIntranetDownloadURL() *string {
	return s.BackupIntranetDownloadURL
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetBackupName() *string {
	return s.BackupName
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetBackupSize() *string {
	return s.BackupSize
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetEngine() *string {
	return s.Engine
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetExtraInfo() *DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo {
	return s.ExtraInfo
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetIsAvail() *string {
	return s.IsAvail
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) GetRecoverConfigMode() *string {
	return s.RecoverConfigMode
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetBackupDownloadURL(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetBackupEndTime(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetBackupId(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.BackupId = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetBackupIntranetDownloadURL(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetBackupName(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.BackupName = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetBackupSize(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.BackupSize = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetBackupStartTime(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetBackupStatus(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.BackupStatus = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetEngine(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.Engine = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetExtraInfo(v *DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.ExtraInfo = v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetInstanceName(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.InstanceName = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetIsAvail(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.IsAvail = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) SetRecoverConfigMode(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackups {
	s.RecoverConfigMode = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackups) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo struct {
	// The engine version.
	//
	// example:
	//
	// 5.0
	CustinsDbVersion *string `json:"CustinsDbVersion,omitempty" xml:"CustinsDbVersion,omitempty"`
}

func (s DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo) GetCustinsDbVersion() *string {
	return s.CustinsDbVersion
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo) SetCustinsDbVersion(v string) *DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo {
	s.CustinsDbVersion = &v
	return s
}

func (s *DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo) Validate() error {
	return dara.Validate(s)
}
