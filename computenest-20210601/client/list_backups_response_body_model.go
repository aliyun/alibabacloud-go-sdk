// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackups(v []*ListBackupsResponseBodyBackups) *ListBackupsResponseBody
	GetBackups() []*ListBackupsResponseBodyBackups
	SetMaxResults(v int32) *ListBackupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListBackupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListBackupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBackupsResponseBody
	GetTotalCount() *int32
}

type ListBackupsResponseBody struct {
	// The backups.
	Backups []*ListBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Repeated"`
	// The maximum number of entries returned on the current page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. If this parameter is empty, all results have been returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLYoaeZA6xVdkCrmG9EmGshtmECUGpq9Qm7x5vQkpz9NXH0XzUc9t4Kxaf3UtuPY4a0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BB58FE53-ED8F-5D12-9746-CD3A5F463D95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackupsResponseBody) GetBackups() []*ListBackupsResponseBodyBackups {
	return s.Backups
}

func (s *ListBackupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBackupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBackupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBackupsResponseBody) SetBackups(v []*ListBackupsResponseBodyBackups) *ListBackupsResponseBody {
	s.Backups = v
	return s
}

func (s *ListBackupsResponseBody) SetMaxResults(v int32) *ListBackupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListBackupsResponseBody) SetNextToken(v string) *ListBackupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListBackupsResponseBody) SetRequestId(v string) *ListBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBackupsResponseBody) SetTotalCount(v int32) *ListBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBackupsResponseBody) Validate() error {
	if s.Backups != nil {
		for _, item := range s.Backups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBackupsResponseBodyBackups struct {
	// The backup ID.
	//
	// example:
	//
	// backup-4e13aa8ca6a34150addd
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup mode. The return value is:
	//
	// - **Manual**: The backup is created manually.
	//
	// example:
	//
	// Manual
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The time when the backup was created.
	//
	// example:
	//
	// 2024-09-03T19:54:38+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the backup.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the backup ended.
	//
	// example:
	//
	// 2024-08-15T02:24:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the backup was last updated.
	//
	// example:
	//
	// 2024-09-03T19:54:38+08:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The retention period in days. Backups are deleted after the retention period expires. By default, this parameter is empty, which means that the backups do not expire.
	//
	// example:
	//
	// 1
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-7b6138dfce1e4c41ab71
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The time when the backup started.
	//
	// example:
	//
	// 2025-06-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status.
	//
	// - Creating: The backup is being created.
	//
	// - Created: The backup is created.
	//
	// - CreateFailed: The backup failed to be created.
	//
	// - Deleting: The backup is being deleted.
	//
	// - Deleted: The backup is deleted.
	//
	// - DeleteFailed: The backup failed to be deleted.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the status.
	//
	// example:
	//
	// Disk i-xxxx backup failed, error message: error
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
}

func (s ListBackupsResponseBodyBackups) String() string {
	return dara.Prettify(s)
}

func (s ListBackupsResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *ListBackupsResponseBodyBackups) GetBackupId() *string {
	return s.BackupId
}

func (s *ListBackupsResponseBodyBackups) GetBackupMode() *string {
	return s.BackupMode
}

func (s *ListBackupsResponseBodyBackups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBackupsResponseBodyBackups) GetDescription() *string {
	return s.Description
}

func (s *ListBackupsResponseBodyBackups) GetEndTime() *string {
	return s.EndTime
}

func (s *ListBackupsResponseBodyBackups) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListBackupsResponseBodyBackups) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ListBackupsResponseBodyBackups) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListBackupsResponseBodyBackups) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBackupsResponseBodyBackups) GetStatus() *string {
	return s.Status
}

func (s *ListBackupsResponseBodyBackups) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *ListBackupsResponseBodyBackups) SetBackupId(v string) *ListBackupsResponseBodyBackups {
	s.BackupId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetBackupMode(v string) *ListBackupsResponseBodyBackups {
	s.BackupMode = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetCreateTime(v string) *ListBackupsResponseBodyBackups {
	s.CreateTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetDescription(v string) *ListBackupsResponseBodyBackups {
	s.Description = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetEndTime(v string) *ListBackupsResponseBodyBackups {
	s.EndTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetModifiedTime(v string) *ListBackupsResponseBodyBackups {
	s.ModifiedTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetRetentionDays(v int32) *ListBackupsResponseBodyBackups {
	s.RetentionDays = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetServiceInstanceId(v string) *ListBackupsResponseBodyBackups {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetStartTime(v string) *ListBackupsResponseBodyBackups {
	s.StartTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetStatus(v string) *ListBackupsResponseBodyBackups {
	s.Status = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetStatusDetail(v string) *ListBackupsResponseBodyBackups {
	s.StatusDetail = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) Validate() error {
	return dara.Validate(s)
}
