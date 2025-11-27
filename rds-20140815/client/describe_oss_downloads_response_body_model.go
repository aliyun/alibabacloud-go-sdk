// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssDownloadsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeOssDownloadsResponseBody
	GetDBInstanceId() *string
	SetItems(v *DescribeOssDownloadsResponseBodyItems) *DescribeOssDownloadsResponseBody
	GetItems() *DescribeOssDownloadsResponseBodyItems
	SetMigrateTaskId(v string) *DescribeOssDownloadsResponseBody
	GetMigrateTaskId() *string
	SetRequestId(v string) *DescribeOssDownloadsResponseBody
	GetRequestId() *string
}

type DescribeOssDownloadsResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Details of the backup file.
	Items *DescribeOssDownloadsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The ID of the migration task.
	//
	// example:
	//
	// 562154852
	MigrateTaskId *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5409D02-D661-4BF3-8F3D-0A814D0574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOssDownloadsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssDownloadsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssDownloadsResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeOssDownloadsResponseBody) GetItems() *DescribeOssDownloadsResponseBodyItems {
	return s.Items
}

func (s *DescribeOssDownloadsResponseBody) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *DescribeOssDownloadsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOssDownloadsResponseBody) SetDBInstanceId(v string) *DescribeOssDownloadsResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeOssDownloadsResponseBody) SetItems(v *DescribeOssDownloadsResponseBodyItems) *DescribeOssDownloadsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeOssDownloadsResponseBody) SetMigrateTaskId(v string) *DescribeOssDownloadsResponseBody {
	s.MigrateTaskId = &v
	return s
}

func (s *DescribeOssDownloadsResponseBody) SetRequestId(v string) *DescribeOssDownloadsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssDownloadsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOssDownloadsResponseBodyItems struct {
	OssDownload []*DescribeOssDownloadsResponseBodyItemsOssDownload `json:"OssDownload,omitempty" xml:"OssDownload,omitempty" type:"Repeated"`
}

func (s DescribeOssDownloadsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssDownloadsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeOssDownloadsResponseBodyItems) GetOssDownload() []*DescribeOssDownloadsResponseBodyItemsOssDownload {
	return s.OssDownload
}

func (s *DescribeOssDownloadsResponseBodyItems) SetOssDownload(v []*DescribeOssDownloadsResponseBodyItemsOssDownload) *DescribeOssDownloadsResponseBodyItems {
	s.OssDownload = v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItems) Validate() error {
	if s.OssDownload != nil {
		for _, item := range s.OssDownload {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOssDownloadsResponseBodyItemsOssDownload struct {
	// The backup type. Valid values:
	//
	// 	- **Database**: full backup file
	//
	// 	- **Differential_Database**: incremental backup file
	//
	// 	- **Transaction_Log**: log backup file
	//
	// example:
	//
	// Database
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The time when the backup file was created in the download list. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-08-17T12:45:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the backup file.
	//
	// example:
	//
	// App description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end of the time range during which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-08-27T12:45:15Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the backup file stored in the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// test
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the backup file. Unit: MB
	//
	// example:
	//
	// 2
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// Indicates whether the backup file is available. Valid values: **True and False**.
	//
	// example:
	//
	// True
	IsAvailable *string `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// The state of the backup file. Valid values:
	//
	// 	- **NoStart**
	//
	// 	- **Downloading**
	//
	// 	- **Finished**
	//
	// 	- **DownloadFailed**
	//
	// 	- **VerifyFailed**
	//
	// 	- **Deleted**
	//
	// 	- **DeleteFailed**
	//
	// 	- **CheckSuccess**
	//
	// 	- **CheckFailed**
	//
	// 	- **Restoring**
	//
	// 	- **Restored**
	//
	// 	- **RestoreFailed**
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeOssDownloadsResponseBodyItemsOssDownload) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssDownloadsResponseBodyItemsOssDownload) GoString() string {
	return s.String()
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) GetDescription() *string {
	return s.Description
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) GetFileName() *string {
	return s.FileName
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) GetFileSize() *string {
	return s.FileSize
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) GetIsAvailable() *string {
	return s.IsAvailable
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) GetStatus() *string {
	return s.Status
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) SetBackupMode(v string) *DescribeOssDownloadsResponseBodyItemsOssDownload {
	s.BackupMode = &v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) SetCreateTime(v string) *DescribeOssDownloadsResponseBodyItemsOssDownload {
	s.CreateTime = &v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) SetDescription(v string) *DescribeOssDownloadsResponseBodyItemsOssDownload {
	s.Description = &v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) SetEndTime(v string) *DescribeOssDownloadsResponseBodyItemsOssDownload {
	s.EndTime = &v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) SetFileName(v string) *DescribeOssDownloadsResponseBodyItemsOssDownload {
	s.FileName = &v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) SetFileSize(v string) *DescribeOssDownloadsResponseBodyItemsOssDownload {
	s.FileSize = &v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) SetIsAvailable(v string) *DescribeOssDownloadsResponseBodyItemsOssDownload {
	s.IsAvailable = &v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) SetStatus(v string) *DescribeOssDownloadsResponseBodyItemsOssDownload {
	s.Status = &v
	return s
}

func (s *DescribeOssDownloadsResponseBodyItemsOssDownload) Validate() error {
	return dara.Validate(s)
}
