// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterBackups(v []*DescribeClusterBackupsResponseBodyClusterBackups) *DescribeClusterBackupsResponseBody
	GetClusterBackups() []*DescribeClusterBackupsResponseBodyClusterBackups
	SetMaxResults(v int32) *DescribeClusterBackupsResponseBody
	GetMaxResults() *int32
	SetPageNumber(v int32) *DescribeClusterBackupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterBackupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeClusterBackupsResponseBody
	GetRequestId() *string
}

type DescribeClusterBackupsResponseBody struct {
	// The cluster backup sets. A cluster backup file contains the backup sets of each node.
	ClusterBackups []*DescribeClusterBackupsResponseBodyClusterBackups `json:"ClusterBackups,omitempty" xml:"ClusterBackups,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2F42BB4E-461F-5B55-A37C-53B1141C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupsResponseBody) GetClusterBackups() []*DescribeClusterBackupsResponseBodyClusterBackups {
	return s.ClusterBackups
}

func (s *DescribeClusterBackupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeClusterBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterBackupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterBackupsResponseBody) SetClusterBackups(v []*DescribeClusterBackupsResponseBodyClusterBackups) *DescribeClusterBackupsResponseBody {
	s.ClusterBackups = v
	return s
}

func (s *DescribeClusterBackupsResponseBody) SetMaxResults(v int32) *DescribeClusterBackupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeClusterBackupsResponseBody) SetPageNumber(v int32) *DescribeClusterBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterBackupsResponseBody) SetPageSize(v int32) *DescribeClusterBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterBackupsResponseBody) SetRequestId(v string) *DescribeClusterBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterBackupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterBackupsResponseBodyClusterBackups struct {
	// The backup status. Valid values:
	//
	// 	- **init**: The backup is being initialized.
	//
	// 	- **No_Need**: Log backup is not performed.
	//
	// 	- **Running**: Log backup is being performed.
	//
	// 	- **Ready**: Log backup is complete.
	//
	// 	- **Failed**: Log backup failed.
	//
	// >  If the **ClusterBackupStatus*	- parameter is set to OK, full backup is successful. If you want to perform point-in-time-restoration on an instance for which log backup is enabled or to implement consistency restoration, make sure that log backup is complete.
	//
	// example:
	//
	// Ready
	AttachLogStatus  *string `json:"AttachLogStatus,omitempty" xml:"AttachLogStatus,omitempty"`
	BackupExpireTime *string `json:"BackupExpireTime,omitempty" xml:"BackupExpireTime,omitempty"`
	// The collection of the backup sets of each child node in a cluster backup set.
	Backups []*DescribeClusterBackupsResponseBodyClusterBackupsBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Repeated"`
	// The end of the time range within which the cluster backup is performed.
	//
	// example:
	//
	// 2023-10-16T19:33:20Z
	ClusterBackupEndTime *string `json:"ClusterBackupEndTime,omitempty" xml:"ClusterBackupEndTime,omitempty"`
	// The backup set ID.
	//
	// example:
	//
	// cb-o8c2ugnxo26kx***
	ClusterBackupId *string `json:"ClusterBackupId,omitempty" xml:"ClusterBackupId,omitempty"`
	// The cluster backup mode.
	//
	// example:
	//
	// Automated
	ClusterBackupMode *string `json:"ClusterBackupMode,omitempty" xml:"ClusterBackupMode,omitempty"`
	// The size of the cluster backup set. Unit: bytes.
	//
	// example:
	//
	// 107374182400
	ClusterBackupSize *string `json:"ClusterBackupSize,omitempty" xml:"ClusterBackupSize,omitempty"`
	// The beginning of the time range within which the cluster backup is performed.
	//
	// example:
	//
	// 2023-10-16T19:33:20Z
	ClusterBackupStartTime *string `json:"ClusterBackupStartTime,omitempty" xml:"ClusterBackupStartTime,omitempty"`
	// The status of the cluster backup set.
	//
	// example:
	//
	// OK
	ClusterBackupStatus *string `json:"ClusterBackupStatus,omitempty" xml:"ClusterBackupStatus,omitempty"`
	// Version of the backuped instance.
	//
	// 	- **6.0**
	//
	// 	- **5.0**
	//
	// 	- **4.4**
	//
	// 	- **4.2**
	//
	// 	- **4.0**
	//
	// 	- **3.4**
	//
	// example:
	//
	// 4.2
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The additional information in the JSON format.
	ExtraInfo *DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	// Indicates whether the cluster backup sets take effect. Valid values:
	//
	// 	- **1**: The cluster backup sets take effect.
	//
	// 	- **0**: The backup sets of child nodes are incomplete or fail.
	//
	// example:
	//
	// 1
	IsAvail *int32 `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	// The progress of the backup task. Unit: %. The progress is returned only for running backup tasks.
	//
	// example:
	//
	// 50
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeClusterBackupsResponseBodyClusterBackups) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupsResponseBodyClusterBackups) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetAttachLogStatus() *string {
	return s.AttachLogStatus
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetBackupExpireTime() *string {
	return s.BackupExpireTime
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetBackups() []*DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	return s.Backups
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetClusterBackupEndTime() *string {
	return s.ClusterBackupEndTime
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetClusterBackupId() *string {
	return s.ClusterBackupId
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetClusterBackupMode() *string {
	return s.ClusterBackupMode
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetClusterBackupSize() *string {
	return s.ClusterBackupSize
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetClusterBackupStartTime() *string {
	return s.ClusterBackupStartTime
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetClusterBackupStatus() *string {
	return s.ClusterBackupStatus
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetExtraInfo() *DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo {
	return s.ExtraInfo
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetIsAvail() *int32 {
	return s.IsAvail
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) GetProgress() *string {
	return s.Progress
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetAttachLogStatus(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.AttachLogStatus = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetBackupExpireTime(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.BackupExpireTime = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetBackups(v []*DescribeClusterBackupsResponseBodyClusterBackupsBackups) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.Backups = v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetClusterBackupEndTime(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.ClusterBackupEndTime = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetClusterBackupId(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.ClusterBackupId = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetClusterBackupMode(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.ClusterBackupMode = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetClusterBackupSize(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.ClusterBackupSize = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetClusterBackupStartTime(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.ClusterBackupStartTime = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetClusterBackupStatus(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.ClusterBackupStatus = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetEngineVersion(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.EngineVersion = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetExtraInfo(v *DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.ExtraInfo = v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetIsAvail(v int32) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.IsAvail = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) SetProgress(v string) *DescribeClusterBackupsResponseBodyClusterBackups {
	s.Progress = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackups) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterBackupsResponseBodyClusterBackupsBackups struct {
	// The URL that is used to download the backup set over the Internet. If the backup set cannot be downloaded, an empty string is returned.
	//
	// example:
	//
	// http://oss.com/xxx
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// The end time of the backup. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-10-16T19:33:20Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 738025367
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The URL that is used to download the backup set over an internal network. If the backup set cannot be downloaded, null is returned.
	//
	// example:
	//
	// http://oss.com/xxx
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// The backup name.
	//
	// example:
	//
	// 12345678.tar.gz
	BackupName *string `json:"BackupName,omitempty" xml:"BackupName,omitempty"`
	// The size of the backup file. Unit: bytes.
	//
	// example:
	//
	// 77544597650
	BackupSize *string `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-10-16T19:33:20Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The status of the backup task. Valid values:
	//
	// 	- **Success**: The backup task is successful.
	//
	// 	- **Failed**: The backup task failed.
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The information of the node associated with the backup.
	ExtraInfo *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	// The shard name.
	//
	// example:
	//
	// d-bp16cb162771****
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
}

func (s DescribeClusterBackupsResponseBodyClusterBackupsBackups) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupsResponseBodyClusterBackupsBackups) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetBackupDownloadURL() *string {
	return s.BackupDownloadURL
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetBackupIntranetDownloadURL() *string {
	return s.BackupIntranetDownloadURL
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetBackupName() *string {
	return s.BackupName
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetBackupSize() *string {
	return s.BackupSize
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetExtraInfo() *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo {
	return s.ExtraInfo
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) GetIsAvail() *string {
	return s.IsAvail
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetBackupDownloadURL(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetBackupEndTime(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetBackupId(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.BackupId = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetBackupIntranetDownloadURL(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetBackupName(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.BackupName = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetBackupSize(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.BackupSize = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetBackupStartTime(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetBackupStatus(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.BackupStatus = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetExtraInfo(v *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.ExtraInfo = v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetInstanceName(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.InstanceName = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) SetIsAvail(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackups {
	s.IsAvail = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackups) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo struct {
	// The instance type of the node.
	//
	// example:
	//
	// mdb.shard.4x.large.d
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The node ID.
	//
	// example:
	//
	// d-2ze75ab1fa5d****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node type.
	//
	// example:
	//
	// db
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The storage capacity of the node. Unit: MB.
	//
	// example:
	//
	// 20480
	StorageSize *string `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) GetStorageSize() *string {
	return s.StorageSize
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) SetInstanceClass(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo {
	s.InstanceClass = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) SetNodeId(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) SetNodeType(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) SetStorageSize(v string) *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo {
	s.StorageSize = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsBackupsExtraInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo struct {
	// Indicates whether the cluster backups are migrated from the historical backup sets. If the value of this parameter is **1**, the cluster backups are migrated from the historical backup sets.
	//
	// example:
	//
	// 1
	RegistryFromHistory *string `json:"RegistryFromHistory,omitempty" xml:"RegistryFromHistory,omitempty"`
}

func (s DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo) GetRegistryFromHistory() *string {
	return s.RegistryFromHistory
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo) SetRegistryFromHistory(v string) *DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo {
	s.RegistryFromHistory = &v
	return s
}

func (s *DescribeClusterBackupsResponseBodyClusterBackupsExtraInfo) Validate() error {
	return dara.Validate(s)
}
