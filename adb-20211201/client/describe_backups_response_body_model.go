// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFreeBackupSize(v int64) *DescribeBackupsResponseBody
	GetFreeBackupSize() *int64
	SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody
	GetItems() *DescribeBackupsResponseBodyItems
	SetPageNumber(v string) *DescribeBackupsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeBackupsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeBackupsResponseBody
	GetRequestId() *string
	SetTotalBackupSize(v int64) *DescribeBackupsResponseBody
	GetTotalBackupSize() *int64
	SetTotalCount(v string) *DescribeBackupsResponseBody
	GetTotalCount() *string
}

type DescribeBackupsResponseBody struct {
	FreeBackupSize *int64 `json:"FreeBackupSize,omitempty" xml:"FreeBackupSize,omitempty"`
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
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CE17270B-F8F8-5A31-9DB4-DADDFDAD7940
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalBackupSize *int64  `json:"TotalBackupSize,omitempty" xml:"TotalBackupSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 300
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) GetFreeBackupSize() *int64 {
	return s.FreeBackupSize
}

func (s *DescribeBackupsResponseBody) GetItems() *DescribeBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeBackupsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeBackupsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupsResponseBody) GetTotalBackupSize() *int64 {
	return s.TotalBackupSize
}

func (s *DescribeBackupsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeBackupsResponseBody) SetFreeBackupSize(v int64) *DescribeBackupsResponseBody {
	s.FreeBackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v string) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalBackupSize(v int64) *DescribeBackupsResponseBody {
	s.TotalBackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v string) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeBackupsResponseBodyItemsBackup struct {
	// The end time of the backup.
	//
	// example:
	//
	// 2022-06-02T16:00Z
	BackupEndTime     *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupExpiredTime *string `json:"BackupExpiredTime,omitempty" xml:"BackupExpiredTime,omitempty"`
	// The backup set ID.
	//
	// example:
	//
	// 32732****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup method. Snapshot is returned.
	//
	// example:
	//
	// Snapshot
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupRegion *string `json:"BackupRegion,omitempty" xml:"BackupRegion,omitempty"`
	// The size of the backup set. Unit: bytes.
	//
	// example:
	//
	// 2167808
	BackupSize *int32 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup.
	//
	// example:
	//
	// 2022-06-01T16:00Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStatus    *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
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
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId    *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ParentBackupId *string `json:"ParentBackupId,omitempty" xml:"ParentBackupId,omitempty"`
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

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupExpiredTime() *string {
	return s.BackupExpiredTime
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupRegion() *string {
	return s.BackupRegion
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetBackupSize() *int32 {
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

func (s *DescribeBackupsResponseBodyItemsBackup) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBackupsResponseBodyItemsBackup) GetParentBackupId() *string {
	return s.ParentBackupId
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupExpiredTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupExpiredTime = &v
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

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupRegion(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupRegion = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSize(v int32) *DescribeBackupsResponseBodyItemsBackup {
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

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetParentBackupId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ParentBackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) Validate() error {
	return dara.Validate(s)
}
