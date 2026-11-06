// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeApplicationBackupsResponseBodyItems) *DescribeApplicationBackupsResponseBody
	GetItems() []*DescribeApplicationBackupsResponseBodyItems
	SetPageNumbers(v string) *DescribeApplicationBackupsResponseBody
	GetPageNumbers() *string
	SetPageRecordCount(v string) *DescribeApplicationBackupsResponseBody
	GetPageRecordCount() *string
	SetRequestId(v string) *DescribeApplicationBackupsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v string) *DescribeApplicationBackupsResponseBody
	GetTotalRecordCount() *string
}

type DescribeApplicationBackupsResponseBody struct {
	// The list of application backup sets.
	Items []*DescribeApplicationBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumbers *string `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty"`
	// The number of backup sets returned on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *string `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of backup sets that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalRecordCount *string `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeApplicationBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationBackupsResponseBody) GetItems() []*DescribeApplicationBackupsResponseBodyItems {
	return s.Items
}

func (s *DescribeApplicationBackupsResponseBody) GetPageNumbers() *string {
	return s.PageNumbers
}

func (s *DescribeApplicationBackupsResponseBody) GetPageRecordCount() *string {
	return s.PageRecordCount
}

func (s *DescribeApplicationBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationBackupsResponseBody) GetTotalRecordCount() *string {
	return s.TotalRecordCount
}

func (s *DescribeApplicationBackupsResponseBody) SetItems(v []*DescribeApplicationBackupsResponseBodyItems) *DescribeApplicationBackupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeApplicationBackupsResponseBody) SetPageNumbers(v string) *DescribeApplicationBackupsResponseBody {
	s.PageNumbers = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBody) SetPageRecordCount(v string) *DescribeApplicationBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBody) SetRequestId(v string) *DescribeApplicationBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBody) SetTotalRecordCount(v string) *DescribeApplicationBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBody) Validate() error {
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

type DescribeApplicationBackupsResponseBodyItems struct {
	// The end time of the backup. The time is displayed in UTC in the YYYY-MM-DDTHH:mm:ssZ format.
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
	// The start time of the backup. The time is displayed in UTC in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2020-11-15T07:30:05Z
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The backup status. Valid values: Success: The backup is completed. Failed: The backup failed.
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The backup type. Valid values: FullBackup: full backup. IncrementalBackup: incremental backup.
	//
	// example:
	//
	// FullBackup
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The backup mode. Valid values: Automated: automatic backup. Manual: manual backup.
	//
	// example:
	//
	// Manual
	JobMode *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
}

func (s DescribeApplicationBackupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationBackupsResponseBodyItems) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeApplicationBackupsResponseBodyItems) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeApplicationBackupsResponseBodyItems) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeApplicationBackupsResponseBodyItems) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeApplicationBackupsResponseBodyItems) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeApplicationBackupsResponseBodyItems) GetJobMode() *string {
	return s.JobMode
}

func (s *DescribeApplicationBackupsResponseBodyItems) SetBackupEndTime(v string) *DescribeApplicationBackupsResponseBodyItems {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBodyItems) SetBackupId(v string) *DescribeApplicationBackupsResponseBodyItems {
	s.BackupId = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBodyItems) SetBackupStartTime(v string) *DescribeApplicationBackupsResponseBodyItems {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBodyItems) SetBackupStatus(v string) *DescribeApplicationBackupsResponseBodyItems {
	s.BackupStatus = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBodyItems) SetBackupType(v string) *DescribeApplicationBackupsResponseBodyItems {
	s.BackupType = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBodyItems) SetJobMode(v string) *DescribeApplicationBackupsResponseBodyItems {
	s.JobMode = &v
	return s
}

func (s *DescribeApplicationBackupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
