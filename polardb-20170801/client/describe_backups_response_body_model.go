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
	SetTotalLevel2BackupSize(v string) *DescribeBackupsResponseBody
	GetTotalLevel2BackupSize() *string
	SetTotalRecordCount(v string) *DescribeBackupsResponseBody
	GetTotalRecordCount() *string
}

type DescribeBackupsResponseBody struct {
	// The queried backup sets.
	Items *DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
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
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 4639948800
	TotalLevel2BackupSize *string `json:"TotalLevel2BackupSize,omitempty" xml:"TotalLevel2BackupSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
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

func (s *DescribeBackupsResponseBody) GetTotalLevel2BackupSize() *string {
	return s.TotalLevel2BackupSize
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

func (s *DescribeBackupsResponseBody) SetTotalLevel2BackupSize(v string) *DescribeBackupsResponseBody {
	s.TotalLevel2BackupSize = &v
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
	// The end time of the backup task. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-11-15T07:30:20Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 61*******
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup method. Only **Snapshot*	- may be returned.
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
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The size of the backup set. Unit: bytes.
	//
	// > After you delete the target snapshot backups, the storage space that is consumed by the backups is released. The released storage space is smaller than the size of the backup file, because the snapshots share specific data blocks. For more information, see [FAQ about backup](https://help.aliyun.com/document_detail/164881.html).
	//
	// example:
	//
	// 4639948800
	BackupSetSize *string `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	// The start time of the backup task. The time is displayed in UTC. Unit: seconds.
	//
	// example:
	//
	// 2020-11-15T07:30:05Z
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
	// The type of the backup. Only **FullBackup*	- may be returned.
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The level of the backup set. Valid values:
	//
	// 	- **Level-1**
	//
	// 	- **Level-2**
	//
	// example:
	//
	// Level-1
	BackupsLevel *string `json:"BackupsLevel,omitempty" xml:"BackupsLevel,omitempty"`
	// The snapshot checkpoint time. The value follows the Unix time format. Unit: seconds.
	//
	// example:
	//
	// 1605425407
	ConsistentTime *string `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The expected expiration time of the backup set (This parameter is supported only for clusters for which sparse backup is enabled).
	//
	// example:
	//
	// 2022-10-24T08:13:23Z
	ExpectExpireTime *string `json:"ExpectExpireTime,omitempty" xml:"ExpectExpireTime,omitempty"`
	// The expected expiration type of the backup set (This parameter is supported only for instances that are enabled with sparse backup).
	//
	// Valid values:
	//
	// 	- NEVER
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- EXPIRED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DELAY
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// EXPIRED
	ExpectExpireType *string `json:"ExpectExpireType,omitempty" xml:"ExpectExpireType,omitempty"`
	// Indicates whether the backup set is available. Valid values:
	//
	// 	- **0**: The backup set is unavailable.
	//
	// 	- **1**: The backup set is available.
	//
	// example:
	//
	// 0
	IsAvail *string `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackup) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupSetSize() *string {
	return s.BackupSetSize
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

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupsLevel() *string {
	return s.BackupsLevel
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetConsistentTime() *string {
	return s.ConsistentTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetExpectExpireTime() *string {
	return s.ExpectExpireTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetExpectExpireType() *string {
	return s.ExpectExpireType
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetIsAvail() *string {
	return s.IsAvail
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupId = &v
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

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSetSize(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupSetSize = &v
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

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupsLevel(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupsLevel = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetConsistentTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetExpectExpireTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ExpectExpireTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetExpectExpireType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ExpectExpireType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetIsAvail(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.IsAvail = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) Validate() error {
	return dara.Validate(s)
}
