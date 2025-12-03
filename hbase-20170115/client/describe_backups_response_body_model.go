// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody
	GetBackups() *DescribeBackupsResponseBodyBackups
	SetEnableStatus(v string) *DescribeBackupsResponseBody
	GetEnableStatus() *string
	SetPageNumber(v int32) *DescribeBackupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBackupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBackupsResponseBody
	GetTotalCount() *int32
}

type DescribeBackupsResponseBody struct {
	Backups      *DescribeBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Struct"`
	EnableStatus *string                             `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	PageNumber   *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) GetBackups() *DescribeBackupsResponseBodyBackups {
	return s.Backups
}

func (s *DescribeBackupsResponseBody) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeBackupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBackupsResponseBody) SetBackups(v *DescribeBackupsResponseBodyBackups) *DescribeBackupsResponseBody {
	s.Backups = v
	return s
}

func (s *DescribeBackupsResponseBody) SetEnableStatus(v string) *DescribeBackupsResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v int32) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageSize(v int32) *DescribeBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotalCount(v int32) *DescribeBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) Validate() error {
	if s.Backups != nil {
		if err := s.Backups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupsResponseBodyBackups struct {
	Backup []*DescribeBackupsResponseBodyBackupsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyBackups) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackups) GetBackup() []*DescribeBackupsResponseBodyBackupsBackup {
	return s.Backup
}

func (s *DescribeBackupsResponseBodyBackups) SetBackup(v []*DescribeBackupsResponseBodyBackupsBackup) *DescribeBackupsResponseBodyBackups {
	s.Backup = v
	return s
}

func (s *DescribeBackupsResponseBodyBackups) Validate() error {
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

type DescribeBackupsResponseBodyBackupsBackup struct {
	BackupDBNames      *string `json:"BackupDBNames,omitempty" xml:"BackupDBNames,omitempty"`
	BackupDownloadURL  *string `json:"BackupDownloadURL,omitempty" xml:"BackupDownloadURL,omitempty"`
	BackupEndTime      *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupEndTimeUTC   *string `json:"BackupEndTimeUTC,omitempty" xml:"BackupEndTimeUTC,omitempty"`
	BackupId           *int32  `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupMethod       *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	BackupMode         *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupSize         *string `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	BackupStartTime    *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	BackupStartTimeUTC *string `json:"BackupStartTimeUTC,omitempty" xml:"BackupStartTimeUTC,omitempty"`
	BackupStatus       *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType         *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
}

func (s DescribeBackupsResponseBodyBackupsBackup) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyBackupsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupDBNames() *string {
	return s.BackupDBNames
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupDownloadURL() *string {
	return s.BackupDownloadURL
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupEndTime() *string {
	return s.BackupEndTime
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupEndTimeUTC() *string {
	return s.BackupEndTimeUTC
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupId() *int32 {
	return s.BackupId
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupSize() *string {
	return s.BackupSize
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupStartTime() *string {
	return s.BackupStartTime
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupStartTimeUTC() *string {
	return s.BackupStartTimeUTC
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeBackupsResponseBodyBackupsBackup) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDBNames(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDBNames = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupDownloadURL(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupDownloadURL = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupEndTimeUTC(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupEndTimeUTC = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupId(v int32) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMethod = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupSize(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStartTimeUTC(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStartTimeUTC = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyBackupsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyBackupsBackup) Validate() error {
	return dara.Validate(s)
}
