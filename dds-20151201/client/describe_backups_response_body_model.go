// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody
	GetBackups() *DescribeBackupsResponseBodyBackups
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
	// The details of the backup set.
	Backups *DescribeBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 275D43C3-F12F-5224-B375-0C6BF453BD56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of backup sets.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) GetBackups() *DescribeBackupsResponseBodyBackups {
	return s.Backups
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

func (s *DescribeBackupsResponseBody) SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody {
	s.Backups = v
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
	if s.Backups != nil {
		if err := s.Backups.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	// The name of the database that has been backed up.
	//
	// example:
	//
	// database
	BackupDBNames *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
	// The URL that is used to download the backup set over the Internet. If the backup set cannot be downloaded, an empty string is returned.
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// The end time of the backup. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-14T03:57:34Z
	BackupEndTime    *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupExpireTime *string `json:"BackupExpireTime,omitempty" xml:"BackupExpireTime,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 32102****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The internal download URL of the backup set.
	//
	// >  You can use the URL to download the specified backup set on an Elastic Compute Service (ECS) instance that is in the same virtual private cloud (VPC) as the ApsaraDB for MongoDB instance.
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// The ID of the backup task.
	//
	// example:
	//
	// 1123xxxx
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The method that is used to generate the backup set. Valid values:
	//
	// 	- **Snapshot**
	//
	// 	- **Physical**
	//
	// 	- **Logical**
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup mode of the backup set. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The name of the backup set. The parameter is invalid.
	//
	// example:
	//
	// 12345678.tar.gz
	BackupName *string `json:"BackupName,omitempty" xml:"BackupName,omitempty"`
	// The backup granularity. The parameter is invalid.
	//
	// example:
	//
	// DBInstance
	BackupScale *string `json:"BackupScale,omitempty" xml:"BackupScale,omitempty"`
	// The size of the backup set. Unit: bytes.
	//
	// example:
	//
	// 6168576
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-14T03:56:17Z
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
	// Availability of the backup set.
	//
	// - 0: unavailable
	//
	// - 1: available
	//
	// example:
	//
	// 1
	IsAvail *bool `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
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

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupExpireTime() *string {
	return s.BackupExpireTime
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupIntranetDownloadURL() *string {
	return s.BackupIntranetDownloadURL
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupName() *string {
	return s.BackupName
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupScale() *string {
	return s.BackupScale
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

func (s *DescribeBackupsResponseBodyBackupsBackup) GetIsAvail() *bool {
	return s.IsAvail
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

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupExpireTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupExpireTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupIntranetDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupJobId(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupJobId = &v
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

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupName(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupName = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupScale(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupScale = &v
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

func (s *DescribeBackupsResponseBodyBackupsBackup) SetIsAvail(v bool) *DescribeBackupsResponseBodyBackupsBackup {
	s.IsAvail = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) Validate() error {
	return dara.Validate(s)
}
