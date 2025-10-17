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
	// Details about the backup sets.
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
	// 15
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 717B2382-BB14-4DCB-BBC2-32DBE0******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 50
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
	// The end time of the backup task. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-12T03:25:55Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup method. Only snapshot backups are supported. The value is set to **Snapshot**.
	//
	// example:
	//
	// Snapshot
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Manual
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The size of the backup set. Unit: bytes.
	//
	// > After you delete the target snapshot backups, the storage space consumed by the backups is released. The released storage space is smaller than the size of the backup file, because the snapshots share certain data blocks.
	//
	// example:
	//
	// 7251056
	BackupSetSize *string `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	// The start time of the backup task. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-12T03:25:44Z
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
	// The type of the backup. Only full backups are supported. The value is set to **FullBackup**.
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The level of the backup set. Valid values:
	//
	// 	- **Level-1**: level-1 backup set
	//
	// 	- **Level-2**: level-2 backup set
	//
	// example:
	//
	// Level-2
	BackupsLevel *string `json:"BackupsLevel,omitempty" xml:"BackupsLevel,omitempty"`
	// The snapshot checkpoint time. The value is a Unix timestamp.
	//
	// example:
	//
	// 1589253947
	ConsistentTime *string `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the backup set is available. Valid values:
	//
	// 	- **0**: The data backup set is unavailable.
	//
	// 	- **1**: The data backup set is available.
	//
	// example:
	//
	// 1
	IsAvail *string `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
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

func (s DescribeDetachedBackupsResponseBodyItemsBackup) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetachedBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupSetSize() *string {
	return s.BackupSetSize
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

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetBackupsLevel() *string {
	return s.BackupsLevel
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetConsistentTime() *string {
	return s.ConsistentTime
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetIsAvail() *string {
	return s.IsAvail
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) GetStoreStatus() *string {
	return s.StoreStatus
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupId = &v
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

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupSetSize(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupSetSize = &v
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

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupsLevel(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupsLevel = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetConsistentTime(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetIsAvail(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.IsAvail = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetStoreStatus(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.StoreStatus = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) Validate() error {
	return dara.Validate(s)
}
