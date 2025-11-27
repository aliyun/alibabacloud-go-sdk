// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDetachedBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeDetachedBackupsResponseBodyItems) *DescribeDetachedBackupsResponseBody
	GetItems() *DescribeDetachedBackupsResponseBodyItems
	SetPageNumber(v string) *DescribeDetachedBackupsResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeDetachedBackupsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeDetachedBackupsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeDetachedBackupsResponseBody
	GetTotalRecordCount() *string
}

type DescribeDetachedBackupsResponseBody struct {
	// The queried backup sets.
	Items *DescribeDetachedBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A6D328C-84B8-40DC-BF49-6C73984D7494
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDetachedBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetachedBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseBody) GetItems() *DescribeDetachedBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeDetachedBackupsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDetachedBackupsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeDetachedBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDetachedBackupsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeDetachedBackupsResponseBody) SetItems(v *DescribeDetachedBackupsResponseBodyItems) *DescribeDetachedBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) SetPageNumber(v string) *DescribeDetachedBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) SetPageRecordCount(v string) *DescribeDetachedBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) SetRequestId(v string) *DescribeDetachedBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) SetTotalRecordCount(v string) *DescribeDetachedBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDetachedBackupsResponseBodyItems struct {
	Backup []*DescribeDetachedBackupsResponseBodyItemsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeDetachedBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetachedBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseBodyItems) GetBackup() []*DescribeDetachedBackupsResponseBodyItemsBackup {
	return s.Backup
}

func (s *DescribeDetachedBackupsResponseBodyItems) SetBackup(v []*DescribeDetachedBackupsResponseBodyItemsBackup) *DescribeDetachedBackupsResponseBodyItems {
	s.Backup = v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItems) Validate() error {
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

type DescribeDetachedBackupsResponseBodyItemsBackup struct {
	// The URL that is used to download the diagnostic report over the Internet. If the diagnostic report cannot be downloaded, an empty string is returned.
	//
	// example:
	//
	// http://rdsbak-hz-v3.oss-cn-hangzhou.aliyuncs.com/xxxxx
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// The end time of the backup task.
	//
	// The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-13T12:20:00Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 321020562
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The URL that is used to download the log file over an internal network. If the log file cannot be downloaded, an empty string is returned.
	//
	// example:
	//
	// http://rdsbak-hz-v3.oss-cn-hangzhou-internal.aliyuncs.com/xxxxx
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// The method that is used to generate the data backup file. Valid values:
	//
	// 	- **Logical**: logical backup
	//
	// 	- **Physical**: physical backup
	//
	// example:
	//
	// Physical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup method. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The backup size. Unit: bytes.
	//
	// example:
	//
	// 2167808
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup task.
	//
	// The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-03T12:20:00Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The status of the backup set. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The backup type of the backup file. Valid values:
	//
	// 	- **FullBackup**
	//
	// 	- **IncrementalBackup**
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The point in time at which the data in the backup set is consistent. The return value of this parameter is a timestamp.
	//
	// >  If the instance runs MySQL 5.6, a timestamp is returned. Otherwise, the value 0 is returned.
	//
	// example:
	//
	// 1576506856
	ConsistentTime *int64 `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test
	DBInstanceComment *string `json:"DBInstanceComment,omitempty" xml:"DBInstanceComment,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the instance that generates the backup set. This parameter is used to indicate whether the instance that generates the backup set is a primary instance or a secondary instance.
	//
	// example:
	//
	// 5882781
	HostInstanceID *string `json:"HostInstanceID,omitempty" xml:"HostInstanceID,omitempty"`
	// Indicates whether the backup set is available. Valid values:
	//
	// 	- **0**: The backup set is unavailable.
	//
	// 	- **1**: The backup set is available.
	//
	// example:
	//
	// 1
	IsAvail *int32 `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	// The status of the backup set that is used to restore individual databases or tables. Valid values:
	//
	// 	- **OK**: The backup set is normal.
	//
	// 	- **LARGE**: The backup set contains an abnormally large number of tables. It cannot be used to restore individual databases or tables.
	//
	// 	- **EMPTY**: The backup set is generated from a failed backup task.
	//
	// example:
	//
	// OK
	MetaStatus *string `json:"MetaStatus,omitempty" xml:"MetaStatus,omitempty"`
	// Indicates whether the data backup file can be deleted. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Disabled
	StoreStatus *string `json:"StoreStatus,omitempty" xml:"StoreStatus,omitempty"`
}

func (s DescribeDetachedBackupsResponseBodyItemsBackup) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetachedBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupDownloadURL() *string {
	return s.BackupDownloadURL
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupIntranetDownloadURL() *string {
	return s.BackupIntranetDownloadURL
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetConsistentTime() *int64 {
	return s.ConsistentTime
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetDBInstanceComment() *string {
	return s.DBInstanceComment
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetHostInstanceID() *string {
	return s.HostInstanceID
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetIsAvail() *int32 {
	return s.IsAvail
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetMetaStatus() *string {
	return s.MetaStatus
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetStoreStatus() *string {
	return s.StoreStatus
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupDownloadURL(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupIntranetDownloadURL(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupMethod(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupMode(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupSize(v int64) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupStatus(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetConsistentTime(v int64) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetDBInstanceComment(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.DBInstanceComment = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetDBInstanceId(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetHostInstanceID(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.HostInstanceID = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetIsAvail(v int32) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.IsAvail = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetMetaStatus(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.MetaStatus = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetStoreStatus(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.StoreStatus = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) Validate() error {
	return dara.Validate(s)
}
