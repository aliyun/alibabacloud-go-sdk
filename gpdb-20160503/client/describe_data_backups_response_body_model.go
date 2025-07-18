// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDataBackupsResponseBodyItems) *DescribeDataBackupsResponseBody
	GetItems() []*DescribeDataBackupsResponseBodyItems
	SetPageNumber(v int32) *DescribeDataBackupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataBackupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataBackupsResponseBody
	GetRequestId() *string
	SetTotalBackupSize(v int64) *DescribeDataBackupsResponseBody
	GetTotalBackupSize() *int64
	SetTotalCount(v int32) *DescribeDataBackupsResponseBody
	GetTotalCount() *int32
}

type DescribeDataBackupsResponseBody struct {
	// The instance ID.
	Items []*DescribeDataBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3E387971-33A5-5019-AD7F-DC**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total backup set size. Unit: Byte.
	//
	// example:
	//
	// 1111111111
	TotalBackupSize *int64 `json:"TotalBackupSize,omitempty" xml:"TotalBackupSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponseBody) GetItems() []*DescribeDataBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeDataBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataBackupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataBackupsResponseBody) GetTotalBackupSize() *int64 {
	return s.TotalBackupSize
}

func (s *DescribeDataBackupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataBackupsResponseBody) SetItems(v []*DescribeDataBackupsResponseBodyItems) *DescribeDataBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetPageNumber(v int32) *DescribeDataBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetPageSize(v int32) *DescribeDataBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetRequestId(v string) *DescribeDataBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetTotalBackupSize(v int64) *DescribeDataBackupsResponseBody {
	s.TotalBackupSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) SetTotalCount(v int32) *DescribeDataBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataBackupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDataBackupsResponseBodyItems struct {
	// The UTC time when the backup ended. The time is in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-22T12:01:43Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The local time when the backup ended. The time is in the yyyy-MM-dd HH:mm:ss format. The time is your local time.
	//
	// example:
	//
	// 2021-12-22 20:00:25
	BackupEndTimeLocal *string `json:"BackupEndTimeLocal,omitempty" xml:"BackupEndTimeLocal,omitempty"`
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
	// The backup mode.
	//
	// Valid values for full backup:
	//
	// 	- Automated: automatic backup
	//
	// 	- Manual: manual backup
	//
	// Valid values for point-in-time backup:
	//
	// 	- Automated: point-in-time backup after full backup
	//
	// 	- Manual: manual point-in-time backup
	//
	// 	- Period: point-in-time backup that is triggered by a backup policy
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 1111111111
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The size of the backup file. Unit: bytes.
	//
	// example:
	//
	// 2167808
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The UTC time when the backup started. The time is in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-22T12:00:25Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The local time when the backup started. The time is in the yyyy-MM-dd HH:mm:ss format. The time is your local time.
	//
	// example:
	//
	// 2011-05-30 03:29:00
	BackupStartTimeLocal *string `json:"BackupStartTimeLocal,omitempty" xml:"BackupStartTimeLocal,omitempty"`
	// The status of the backup set. Valid values:
	//
	// 	- Success
	//
	// 	- Failure
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The name of a point-in-time backup set or the full backup set.
	//
	// example:
	//
	// adbpgbackup_555*****_20211222200019
	BaksetName *string `json:"BaksetName,omitempty" xml:"BaksetName,omitempty"`
	// 	- For full backup, this parameter indicates the point in time at which the data in the data backup file is consistent.
	//
	// 	- For point-in-time backup, this parameter indicates that the returned point in time is a timestamp.
	//
	// example:
	//
	// 1576506856
	ConsistentTime *int64 `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp**************-master
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the backup. Valid values:
	//
	// 	- DATA: full backup
	//
	// 	- RESTOREPOI: point-in-time backup
	//
	// example:
	//
	// DATA
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s DescribeDataBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupEndTimeLocal() *string {
	return s.BackupEndTimeLocal
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupStartTimeLocal() *string {
	return s.BackupStartTimeLocal
}

func (s *DescribeDataBackupsResponseBodyItems) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeDataBackupsResponseBodyItems) GetBaksetName() *string {
	return s.BaksetName
}

func (s *DescribeDataBackupsResponseBodyItems) GetConsistentTime() *int64 {
	return s.ConsistentTime
}

func (s *DescribeDataBackupsResponseBodyItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDataBackupsResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupEndTime(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupEndTimeLocal(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupEndTimeLocal = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupMethod(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupMethod = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupMode(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupMode = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupSetId(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupSetId = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupSize(v int64) *DescribeDataBackupsResponseBodyItems {
	s.BackupSize = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStartTime(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStartTimeLocal(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStartTimeLocal = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBackupStatus(v string) *DescribeDataBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetBaksetName(v string) *DescribeDataBackupsResponseBodyItems {
	s.BaksetName = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetConsistentTime(v int64) *DescribeDataBackupsResponseBodyItems {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetDBInstanceId(v string) *DescribeDataBackupsResponseBodyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) SetDataType(v string) *DescribeDataBackupsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *DescribeDataBackupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
