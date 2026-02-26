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
	BackupDownloadLinkByDB    *DescribeBackupsResponseBodyItemsBackupBackupDownloadLinkByDB `json:"BackupDownloadLinkByDB,omitempty" xml:"BackupDownloadLinkByDB,omitempty" type:"Struct"`
	BackupDownloadURL         *string                                                       `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	BackupEndTime             *string                                                       `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId                  *string                                                       `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupInitiator           *string                                                       `json:"BackupInitiator,omitempty" xml:"BackupInitiator,omitempty"`
	BackupIntranetDownloadURL *string                                                       `json:"BackupIntranetDownloadURL,omitempty" xml:"BackupIntranetDownloadURL,omitempty"`
	BackupMethod              *string                                                       `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupMode                *string                                                       `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupSize                *int64                                                        `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStartTime           *string                                                       `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStatus              *string                                                       `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType                *string                                                       `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	Checksum                  *string                                                       `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	ConsistentTime            *int64                                                        `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	CopyOnlyBackup            *string                                                       `json:"CopyOnlyBackup,omitempty" xml:"CopyOnlyBackup,omitempty"`
	DBInstanceId              *string                                                       `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Encryption                *string                                                       `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	Engine                    *string                                                       `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion             *string                                                       `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	ExpectExpireTime          *string                                                       `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	HostInstanceID            *string                                                       `json:"HostInstanceID,omitempty" xml:"HostInstanceID,omitempty"`
	IsAvail                   *int32                                                        `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	MetaStatus                *string                                                       `json:"MetaStatus,omitempty" xml:"MetaStatus,omitempty"`
	StorageClass              *string                                                       `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	StoreStatus               *string                                                       `json:"StoreStatus,omitempty" xml:"StoreStatus,omitempty"`
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
	DataBase             *string `json:"DataBase,omitempty" xml:"DataBase,omitempty"`
	DownloadLink         *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
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
