// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseDataBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListSupabaseDataBackupsResponseBodyItems) *ListSupabaseDataBackupsResponseBody
	GetItems() []*ListSupabaseDataBackupsResponseBodyItems
	SetMaxResults(v int32) *ListSupabaseDataBackupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupabaseDataBackupsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListSupabaseDataBackupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSupabaseDataBackupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSupabaseDataBackupsResponseBody
	GetRequestId() *string
	SetTotalBackupSize(v int64) *ListSupabaseDataBackupsResponseBody
	GetTotalBackupSize() *int64
	SetTotalCount(v int32) *ListSupabaseDataBackupsResponseBody
	GetTotalCount() *int32
}

type ListSupabaseDataBackupsResponseBody struct {
	// The list of backup sets.
	Items []*ListSupabaseDataBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The maximum number of entries to return for the current request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. You can use this value as the NextToken parameter in the next request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of backup sets on the current page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total size of the backup sets. Unit: bytes.
	//
	// example:
	//
	// 1111111111
	TotalBackupSize *int64 `json:"TotalBackupSize,omitempty" xml:"TotalBackupSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSupabaseDataBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseDataBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupabaseDataBackupsResponseBody) GetItems() []*ListSupabaseDataBackupsResponseBodyItems {
	return s.Items
}

func (s *ListSupabaseDataBackupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupabaseDataBackupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupabaseDataBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSupabaseDataBackupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSupabaseDataBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupabaseDataBackupsResponseBody) GetTotalBackupSize() *int64 {
	return s.TotalBackupSize
}

func (s *ListSupabaseDataBackupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSupabaseDataBackupsResponseBody) SetItems(v []*ListSupabaseDataBackupsResponseBodyItems) *ListSupabaseDataBackupsResponseBody {
	s.Items = v
	return s
}

func (s *ListSupabaseDataBackupsResponseBody) SetMaxResults(v int32) *ListSupabaseDataBackupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBody) SetNextToken(v string) *ListSupabaseDataBackupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBody) SetPageNumber(v int32) *ListSupabaseDataBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBody) SetPageSize(v int32) *ListSupabaseDataBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBody) SetRequestId(v string) *ListSupabaseDataBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBody) SetTotalBackupSize(v int64) *ListSupabaseDataBackupsResponseBody {
	s.TotalBackupSize = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBody) SetTotalCount(v int32) *ListSupabaseDataBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSupabaseDataBackupsResponseBodyItems struct {
	// The end time of the backup. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 2026-10-09T01:24:44Z
	BackupEndTime *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The local time representation of the backup end time. Format: yyyy-MM-ddTHH:mm:ssZ. The current return value is in Beijing time (UTC+8). The trailing Z is a fixed character in the compatibility format and does not indicate the zero time zone. To parse the time in a standard format, use BackupEndTime.
	//
	// example:
	//
	// 2026-10-09T09:24:44Z
	BackupEndTimeLocal *string `json:"BackupEndTimeLocal,omitempty" xml:"BackupEndTimeLocal,omitempty"`
	// The backup method. Valid values: Physical: physical backup; Snapshot: snapshot backup.
	//
	// example:
	//
	// Snapshot
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The backup mode.
	//
	// Valid values for automatic backups:
	//
	// - **Automated**: automatic system backup.
	//
	// - **Manual**: manual backup.
	//
	// Valid values for restorable points:
	//
	// - **Automated**: the restorable point after a automatic backup.
	//
	// - **Manual**: the restorable point manually triggered by the user.
	//
	// - **Period**: the restorable point triggered periodically based on the backup policy.
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
	// 10737418240
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The start time of the backup. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 2026-10-09T01:23:02Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The local time representation of the backup start time. Format: yyyy-MM-ddTHH:mm:ssZ. The current return value is in Beijing time (UTC+8). The trailing Z is a fixed character in the compatibility format and does not indicate the zero time zone. To parse the time in a standard format, use BackupStartTime.
	//
	// example:
	//
	// 2026-10-09T09:23:02Z
	BackupStartTimeLocal *string `json:"BackupStartTimeLocal,omitempty" xml:"BackupStartTimeLocal,omitempty"`
	// The status of the backup set. Valid values:
	//
	// - **Success**: successful.
	//
	// - **Failure**: failed.
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The name of the restorable point or the full backup set.
	//
	// example:
	//
	// logic_backup
	BaksetName *string `json:"BaksetName,omitempty" xml:"BaksetName,omitempty"`
	// The consistency point in time. The value is a UNIX timestamp in seconds. For a full backup, this parameter indicates the consistency point in time of the backup. For a restorable point, this parameter indicates the point in time to which data can be restored.
	//
	// example:
	//
	// 1791508983
	ConsistentTime *int64 `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	// The backup type. Valid values:
	//
	// - **DATA**: full backup.
	//
	// - **RESTOREPOI**: restorable point.
	//
	// example:
	//
	// DATA
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
}

func (s ListSupabaseDataBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseDataBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupEndTimeLocal() *string {
	return s.BackupEndTimeLocal
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupMode() *string {
	return s.BackupMode
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupStartTimeLocal() *string {
	return s.BackupStartTimeLocal
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetBaksetName() *string {
	return s.BaksetName
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetConsistentTime() *int64 {
	return s.ConsistentTime
}

func (s *ListSupabaseDataBackupsResponseBodyItems) GetDataType() *string {
	return s.DataType
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupEndTime(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupEndTimeLocal(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupEndTimeLocal = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupMethod(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupMethod = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupMode(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupMode = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupSetId(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupSetId = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupSize(v int64) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupSize = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupStartTime(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupStartTimeLocal(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupStartTimeLocal = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBackupStatus(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetBaksetName(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.BaksetName = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetConsistentTime(v int64) *ListSupabaseDataBackupsResponseBodyItems {
	s.ConsistentTime = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) SetDataType(v string) *ListSupabaseDataBackupsResponseBodyItems {
	s.DataType = &v
	return s
}

func (s *ListSupabaseDataBackupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
