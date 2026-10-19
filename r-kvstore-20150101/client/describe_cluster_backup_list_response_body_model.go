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
	// The list of cluster backup sets. A cluster backup set contains the backup sets of all nodes in the cluster.
	ClusterBackups []*DescribeClusterBackupListResponseBodyClusterBackups `json:"ClusterBackups,omitempty" xml:"ClusterBackups,omitempty" type:"Repeated"`
	// <props="china">The free backup quota of the instance. Unit: bytes. The free quota is equal to the default memory capacity of the instance type. For more information, see the notice on backup commercialization. Full backups and log backups share this free quota. When you change the instance type, the free quota of the instance changes accordingly.
	//
	// <props="intl">This parameter is not in effect. Ignore this parameter.
	//
	// example:
	//
	// 100000
	FreeSize *int64 `json:"FreeSize,omitempty" xml:"FreeSize,omitempty"`
	// The size of full backup files for the current instance. Unit: bytes. Full backups are generated from daily scheduled backups, manual backups, and backups generated during cache analysis.
	//
	// > This value is independent of the number and size of backup sets returned in the current request. It represents the total size of all valid full backups for the instance.
	//
	// example:
	//
	// 1000
	FullStorageSize *int64 `json:"FullStorageSize,omitempty" xml:"FullStorageSize,omitempty"`
	// The size of the instance log backup files for the current instance. Unit: bytes. This value is valid only when the flashback feature is enabled.
	//
	// > This value is independent of the number and size of backup sets in the current response. It represents the total size of all valid instance log backups for the current instance.
	//
	// example:
	//
	// 5000
	LogStorageSize *int64 `json:"LogStorageSize,omitempty" xml:"LogStorageSize,omitempty"`
	// The maximum number of records returned by the current request.
	//
	// example:
	//
	// 4
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number, which is the same as the request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of records per page, which is the same as the request parameter.
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
	if s.ClusterBackups != nil {
		for _, item := range s.ClusterBackups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterBackupListResponseBodyClusterBackups struct {
	// The collection of backup sets for all child nodes under a single cluster backup set.
	Backups []*DescribeClusterBackupListResponseBodyClusterBackupsBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Repeated"`
	// The end time of the cluster backup. The time is in the YYYY-MM-DDTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2024-01-10T17:21:55Z
	ClusterBackupEndTime *string `json:"ClusterBackupEndTime,omitempty" xml:"ClusterBackupEndTime,omitempty"`
	// The cluster backup set ID.
	//
	// example:
	//
	// cb-zmdqj2m3xyxjtdt0
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The cluster backup mode.
	//
	// example:
	//
	// Automated
	ClusterBackupMode *string `json:"ClusterBackupMode,omitempty" xml:"ClusterBackupMode,omitempty"`
	// The size of the cluster backup set.
	//
	// example:
	//
	// 2048
	ClusterBackupSize *string `json:"ClusterBackupSize,omitempty" xml:"ClusterBackupSize,omitempty"`
	// The start time of the cluster backup. The time is in the YYYY-MM-DDTHH:mm:ssZ format and displayed in UTC.
	//
	// example:
	//
	// 2024-01-10T17:21:25Z
	ClusterBackupStartTime *string `json:"ClusterBackupStartTime,omitempty" xml:"ClusterBackupStartTime,omitempty"`
	// The status of the cluster backup set. Valid values:
	//
	// - OK: The backup is successful.
	//
	// - RUNNING: The backup is in progress.
	//
	// - Other values: The backup failed.
	//
	// example:
	//
	// OK
	ClusterBackupStatus *string `json:"ClusterBackupStatus,omitempty" xml:"ClusterBackupStatus,omitempty"`
	// The expiration time of the backup set. The time is in the yyyy-MM-ddTHH:mmZ format and displayed in UTC.
	//
	// example:
	//
	// 2025-06-01T21:51:22Z
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// Indicates whether the cluster backup set is valid. The value is 0 if the backup sets of child nodes are incomplete or failed.
	//
	// example:
	//
	// 1
	IsAvail *int32 `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	// The backup progress. This parameter is returned only for backups that are in progress.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The memory size of a single DB node when the full backup is created. Unit: MB.
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
	if s.Backups != nil {
		for _, item := range s.Backups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterBackupListResponseBodyClusterBackupsBackups struct {
	// The public download URL of the backup file.
	//
	// example:
	//
	// http://rdsbakbucket-huhehaote-v2.oss-cn-huhehaote.aliyuncs.com/custins424747958/hins100322105_data_20240110012135.rdb
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// The end time of the backup. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format and displayed in UTC.
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
	// > You can use this URL to download the backup file from an ECS instance that is connected to the current instance. Both instances must reside in the same VPC.
	//
	// example:
	//
	// http://rdsbakbucket-huhehaote-v2.oss-cn-huhehaote-internal.aliyuncs.com/custins424747958/hins100322105_data_20240110012135.rdb
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// The backup name.
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
	// The start time of the backup. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format and displayed in UTC.
	//
	// example:
	//
	// 2024-01-09T17:21:30Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The backup status. Valid values:
	//
	// 	- **OK**: The backup is successful.
	//
	// 	- **ERROR**: The backup failed.
	//
	// example:
	//
	// OK
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The database engine type. The value is fixed as **redis**.
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The additional information.
	ExtraInfo *DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	// The name of the instance.
	//
	// example:
	//
	// r-bp1y4is8svonly****-db-0
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the backup set is available. Valid values:
	//
	// 	- **0**: unavailable.
	//
	// 	- **1**: available.
	//
	// example:
	//
	// 1
	IsAvail *string `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	// This parameter is not in effect. Ignore this parameter.
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
	if s.ExtraInfo != nil {
		if err := s.ExtraInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterBackupListResponseBodyClusterBackupsBackupsExtraInfo struct {
	// The database engine version.
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
