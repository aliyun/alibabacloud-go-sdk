// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody
	GetItems() *DescribeBackupsResponseBodyItems
	SetPageNumber(v string) *DescribeBackupsResponseBody
	GetPageNumber() *string
	SetPageRecordCount(v string) *DescribeBackupsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeBackupsResponseBody
	GetRequestId() *string
	SetTotalEcsSnapshotSize(v int64) *DescribeBackupsResponseBody
	GetTotalEcsSnapshotSize() *int64
	SetTotalRecordCount(v string) *DescribeBackupsResponseBody
	GetTotalRecordCount() *string
}

type DescribeBackupsResponseBody struct {
	// The returned backup sets.
	Items *DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of backup sets on the current page.
	//
	// example:
	//
	// 30
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A6D328C-84B8-40DC-BF49-6C73984D7494
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The size of the snapshot chain of the instance. Unit: bytes.
	//
	// example:
	//
	// 0
	TotalEcsSnapshotSize *int64 `json:"TotalEcsSnapshotSize,omitempty" xml:"TotalEcsSnapshotSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) GetItems() *DescribeBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeBackupsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupsResponseBody) GetTotalEcsSnapshotSize() *int64 {
	return s.TotalEcsSnapshotSize
}

func (s *DescribeBackupsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeBackupsResponseBody) SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageRecordCount(v string) *DescribeBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalEcsSnapshotSize(v int64) *DescribeBackupsResponseBody {
	s.TotalEcsSnapshotSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalRecordCount(v string) *DescribeBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupsResponseBodyItems struct {
	Backup []*DescribeBackupsResponseBodyItemsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItems) GetBackup() []*DescribeBackupsResponseBodyItemsBackup {
	return s.Backup
}

func (s *DescribeBackupsResponseBodyItems) SetBackup(v []*DescribeBackupsResponseBodyItemsBackup) *DescribeBackupsResponseBodyItems {
	s.Backup = v
	return s
}

func (s *DescribeBackupsResponseBodyItems) Validate() error {
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

type DescribeBackupsResponseBodyItemsBackup struct {
	// An array consisting of URLs from which you can download backup sets of individual databases.
	BackupDownloadLinkByDB *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB `json:"BackupDownloadLinkByDB,omitempty" xml:"BackupDownloadLinkByDB,omitempty" type:"Struct"`
	// The URL that is used to download the backup set over the Internet. If the backup set cannot be downloaded, null is returned.
	//
	// >  For example, if BackupMethod of an ApsaraDB RDS for SQL Server instance is set to **Snapshot**, a null string is returned.
	//
	// example:
	//
	// http://rdsbak-hz-v3.oss-cn-hangzhou.aliyuncs.com/xxxxx
	BackupDownloadURL *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	// The end time of the backup task. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
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
	// The initiator of the backup task. Valid values:
	//
	// 	- **System**
	//
	// 	- **User**
	//
	// example:
	//
	// System
	BackupInitiator *string `json:"BackupInitiator,omitempty" xml:"BackupInitiator,omitempty"`
	// The URL that is used to download the backup set over an internal network. If the backup set cannot be downloaded, null is returned.
	//
	// >  For example, if BackupMethod of an ApsaraDB RDS for SQL Server instance is set to **Snapshot**, a null string is returned.
	//
	// example:
	//
	// http://rdsbak-hz-v3.oss-cn-hangzhou-internal.aliyuncs.com/xxxxx
	BackupIntranetDownloadURL *string `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	// The method that is used to generate the backup set. Valid values:
	//
	// 	- **Logical**: logical backup
	//
	// 	- **Physical**: physical backup
	//
	// 	- **Snapshot**: snapshot backup
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
	// The size of the data backup file. Unit: bytes.
	//
	// example:
	//
	// 2167808
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-02-03T12:20:00Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The state of the backup set.
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The backup type of the backup set. Valid values:
	//
	// 	- **FullBackup**
	//
	// 	- **IncrementalBackup**
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The checksum. The value of this parameter is calculated by using the CRC64 algorithm.
	//
	// example:
	//
	// 1835830439**********
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// The point in time at which the data in the backup set is consistent. The return value of this parameter is a timestamp.
	//
	// >  If the instance runs MySQL 5.6, a timestamp is returned. Otherwise, the value 0 is returned.
	//
	// example:
	//
	// 1576506856
	ConsistentTime *int64 `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The backup mode of the backup set. Valid values:
	//
	// 	- 0: the standard mode. This mode supports full backups and incremental backups.
	//
	// 	- 1: the copy-only mode. This mode supports only full backups.
	//
	// >  This parameter is returned only when the instance runs SQL Server.
	//
	// example:
	//
	// 0
	CopyOnlyBackup *string `json:"CopyOnlyBackup,omitempty" xml:"CopyOnlyBackup,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The encryption information about the backup set.
	//
	// example:
	//
	// {}
	Encryption *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	// The type of the database engine. Valid values:
	//
	// 	- MySQL
	//
	// 	- SQLServer
	//
	// 	- PostgreSQL
	//
	// 	- MariaDB
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 8.0
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
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
	// 	- **OK**: The data backup file is normal.
	//
	// 	- **LARGE**: The data backup file contains an abnormally large number of tables. It cannot be used to restore individual databases or tables.
	//
	// 	- **EMPTY**: The data backup file is generated from a failed backup task.
	//
	// >  If an empty string is returned, the data backup file cannot be used to restore individual databases or tables.
	//
	// example:
	//
	// OK
	MetaStatus *string `json:"MetaStatus,omitempty" xml:"MetaStatus,omitempty"`
	// The storage class of the backup set. Valid values:
	//
	// 	- **0**: regular storage
	//
	// 	- **1**: archive storage
	//
	// example:
	//
	// 0
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	// Indicates whether the backup set can be deleted. Valid values:
	//
	// 	- **Enabled**: The backup set can be deleted.
	//
	// 	- **Disabled**: The backup set cannot be deleted.
	//
	// example:
	//
	// Disabled
	StoreStatus *string `json:"StoreStatus,omitempty" xml:"StoreStatus,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackup) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupDownloadLinkByDB() *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB {
	return s.BackupDownloadLinkByDB
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupDownloadURL() *string {
	return s.BackupDownloadURL
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupInitiator() *string {
	return s.BackupInitiator
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupIntranetDownloadURL() *string {
	return s.BackupIntranetDownloadURL
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetChecksum() *string {
	return s.Checksum
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetConsistentTime() *int64 {
	return s.ConsistentTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetCopyOnlyBackup() *string {
	return s.CopyOnlyBackup
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetEncryption() *string {
	return s.Encryption
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetEngine() *string {
	return s.Engine
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetHostInstanceID() *string {
	return s.HostInstanceID
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetIsAvail() *int32 {
	return s.IsAvail
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetMetaStatus() *string {
	return s.MetaStatus
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetStoreStatus() *string {
	return s.StoreStatus
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupDownloadLinkByDB(v *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupDownloadLinkByDB = v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupDownloadURL(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupInitiator(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupInitiator = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupIntranetDownloadURL(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupIntranetDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSize(v int64) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetChecksum(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.Checksum = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetConsistentTime(v int64) *DescribeBackupsResponseBodyItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetCopyOnlyBackup(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.CopyOnlyBackup = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBInstanceId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetEncryption(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.Encryption = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetEngine(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.Engine = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetEngineVersion(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.EngineVersion = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetExpectExpireTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetHostInstanceID(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.HostInstanceID = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetIsAvail(v int32) *DescribeBackupsResponseBodyItemsBackup {
	s.IsAvail = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetMetaStatus(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.MetaStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetStorageClass(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.StorageClass = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetStoreStatus(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.StoreStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) Validate() error {
	if s.BackupDownloadLinkByDB != nil {
		if err := s.BackupDownloadLinkByDB.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB struct {
	BackupDownloadLinkByDB []*DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB `json:"BackupDownloadLinkByDB,omitempty" xml:"BackupDownloadLinkByDB,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB) GetBackupDownloadLinkByDB() []*DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB {
	return s.BackupDownloadLinkByDB
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB) SetBackupDownloadLinkByDB(v []*DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB {
	s.BackupDownloadLinkByDB = v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB) Validate() error {
	if s.BackupDownloadLinkByDB != nil {
		for _, item := range s.BackupDownloadLinkByDB {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB struct {
	// The name of the database.
	//
	// example:
	//
	// dbs
	DataBase *string `json:"DataBase,omitempty" xml:"DataBase,omitempty"`
	// The public URL from which you can download the backup set.
	//
	// example:
	//
	// https://cn-hangzhou.bak.rds.aliyuncs.com/custins53664665/hins18676859_2021072909473127987849.zip?Expires=*****&dbList=tb1
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	// The internal URL from which you can download the backup set.
	//
	// example:
	//
	// https://cn-hangzhou-internal.bak.rds.aliyuncs.com/custins53664665/hins18676859_2021072909473127987849.zip?Expires=*****&dbList=tb1
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) GetDataBase() *string {
	return s.DataBase
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) GetDownloadLink() *string {
	return s.DownloadLink
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) GetIntranetDownloadLink() *string {
	return s.IntranetDownloadLink
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) SetDataBase(v string) *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB {
	s.DataBase = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) SetDownloadLink(v string) *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) SetIntranetDownloadLink(v string) *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDBBackupDownloadLinkByDB) Validate() error {
	return dara.Validate(s)
}
