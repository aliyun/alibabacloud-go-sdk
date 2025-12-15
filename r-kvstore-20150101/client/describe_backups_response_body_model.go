// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeBackupsResponseBodyAccessDeniedDetail) *DescribeBackupsResponseBody
	GetAccessDeniedDetail() *DescribeBackupsResponseBodyAccessDeniedDetail
	SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody
	GetBackups() *DescribeBackupsResponseBodyBackups
	SetFreeSize(v int64) *DescribeBackupsResponseBody
	GetFreeSize() *int64
	SetFullStorageSize(v int64) *DescribeBackupsResponseBody
	GetFullStorageSize() *int64
	SetLogStorageSize(v int64) *DescribeBackupsResponseBody
	GetLogStorageSize() *int64
	SetPageNumber(v int32) *DescribeBackupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBackupsResponseBody
	GetTotalCount() *int32
}

type DescribeBackupsResponseBody struct {
	// The following parameters are no longer used. Ignore the parameters.
	AccessDeniedDetail *DescribeBackupsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The queried backup sets.
	Backups *DescribeBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Struct"`
	// This parameter does not take effect. Ignore this parameter.
	//
	// example:
	//
	// 100000
	FreeSize *int64 `json:"FreeSize,omitempty" xml:"FreeSize,omitempty"`
	// The size of the full backup file of the instance. Unit: bytes. Full backups originate from scheduled backups, manual backups, and backups generated during cache analysis.
	//
	// >  The value of this parameter is independent of the number and size of the returned backup sets. Instead, it reflects the total size of all valid full backups of the instance.
	//
	// example:
	//
	// 1000
	FullStorageSize *int64 `json:"FullStorageSize,omitempty" xml:"FullStorageSize,omitempty"`
	// The size of the log backup file of the instance. Unit: bytes. This value is valid only when flashback is enabled.
	//
	// >  The value of this parameter is independent of the number and size of the returned backup sets. Instead, it reflects the total size of all valid log backups of the instance.
	//
	// example:
	//
	// 5000
	LogStorageSize *int64 `json:"LogStorageSize,omitempty" xml:"LogStorageSize,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 963C20F0-7CE1-4591-AAF3-6F3CD1CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of backup files that were returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) GetAccessDeniedDetail() *DescribeBackupsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeBackupsResponseBody) GetBackups() *DescribeBackupsResponseBodyBackups {
	return s.Backups
}

func (s *DescribeBackupsResponseBody) GetFreeSize() *int64 {
	return s.FreeSize
}

func (s *DescribeBackupsResponseBody) GetFullStorageSize() *int64 {
	return s.FullStorageSize
}

func (s *DescribeBackupsResponseBody) GetLogStorageSize() *int64 {
	return s.LogStorageSize
}

func (s *DescribeBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBackupsResponseBody) SetAccessDeniedDetail(v *DescribeBackupsResponseBodyAccessDeniedDetail) *DescribeBackupsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeBackupsResponseBody) SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody {
	s.Backups = v
	return s
}

func (s *DescribeBackupsResponseBody) SetFreeSize(v int64) *DescribeBackupsResponseBody {
	s.FreeSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetFullStorageSize(v int64) *DescribeBackupsResponseBody {
	s.FullStorageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetLogStorageSize(v int64) *DescribeBackupsResponseBody {
	s.LogStorageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v int32) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v int32) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v int32) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Backups != nil {
		if err := s.Backups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupsResponseBodyAccessDeniedDetail struct {
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// This parameter is no longer used. Ignore this parameter.
	//
	// example:
	//
	// _
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeBackupsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeBackupsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeBackupsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeBackupsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeBackupsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeBackupsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeBackupsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeBackupsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeBackupsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupsResponseBodyBackups struct {
	Backup []*DescribeBackupsResponseBodyBackupsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyBackups) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackups) GetBackup() []*DescribeBackupsResponseBodyBackupsBackup {
	return s.Backup
}

func (s *DescribeBackupsResponseBodyBackups) SetBackup(v []*DescribeBackupsResponseBodyBackupsBackup) *DescribeBackupsResponseBodyBackups {
	s.Backup = v
	return s
}

func (s *DescribeBackupsResponseBodyBackups) Validate() error {
	if s.Backup != nil {
		for _, item := range s.Backup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupsResponseBodyBackupsBackup struct {
	// The names of the databases that are backed up. The default value is **all**, which indicates that all databases are backed up.
	//
	// example:
	//
	// all
	BackupDBNames *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
	// The public download URL of the backup file.
	//
	// example:
	//
	// https://rdsbak-hk45-v2.oss-cn-hongkong.aliyuncs.com/********
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// The end time of the backup.
	//
	// example:
	//
	// 2019-03-14T05:31:13Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The ID of the backup file.
	//
	// example:
	//
	// 165*****50
	BackupId *int64 `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The internal download URL of the backup file.
	//
	// >  You can use this URL to download the backup file from an Elastic Compute Service (ECS) instance that is connected to the Tair instance. The ECS instance must belong to the same classic network or reside in the same virtual private cloud (VPC) as the Tair instance.
	//
	// example:
	//
	// https://rdsbak-hk45-v2.oss-cn-hongkong.aliyuncs.com/********
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// The ID of the backup task.
	//
	// example:
	//
	// 24340
	BackupJobID *int64 `json:"BackupJobID,omitempty" xml:"BackupJobID,omitempty"`
	// The backup method. Valid values:
	//
	// 	- **Logical**
	//
	// 	- **Physical**
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The size of the backup file.
	//
	// example:
	//
	// 1024
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup.
	//
	// example:
	//
	// 2019-03-14T05:28:50Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The status of the backup. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **FullBackup**
	//
	// 	- **IncrementalBackup**
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The engine version (major version) of the instance.
	//
	// example:
	//
	// 4.0
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// The node ID.
	//
	// >  If the instance uses the standard architecture, this parameter returns the instance ID.
	//
	// example:
	//
	// r-bp10noxlhcoim2****-db-1
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	// If the backup includes account information, kernel parameters and whitelist details.
	//
	// example:
	//
	// {"whitelist":true,"config":true,"account":true}
	RecoverConfigMode *string `json:"RecoverConfigMode,omitempty" xml:"RecoverConfigMode,omitempty"`
}

func (s DescribeBackupsResponseBodyBackupsBackup) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackupsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupDBNames() *string {
	return s.BackupDBNames
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupDownloadURL() *string {
	return s.BackupDownloadURL
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupId() *int64 {
	return s.BackupId
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupIntranetDownloadURL() *string {
	return s.BackupIntranetDownloadURL
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupJobID() *int64 {
	return s.BackupJobID
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetNodeInstanceId() *string {
	return s.NodeInstanceId
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetRecoverConfigMode() *string {
	return s.RecoverConfigMode
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDBNames(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDBNames = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v int64) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupIntranetDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupJobID(v int64) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupJobID = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupSize(v int64) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetEngineVersion(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.EngineVersion = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetExpectExpireTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetNodeInstanceId(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.NodeInstanceId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetRecoverConfigMode(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.RecoverConfigMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) Validate() error {
	return dara.Validate(s)
}
