// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeDataBackupsRequest
	GetBackupId() *string
	SetBackupMode(v string) *DescribeDataBackupsRequest
	GetBackupMode() *string
	SetBackupStatus(v string) *DescribeDataBackupsRequest
	GetBackupStatus() *string
	SetDBInstanceId(v string) *DescribeDataBackupsRequest
	GetDBInstanceId() *string
	SetDataType(v string) *DescribeDataBackupsRequest
	GetDataType() *string
	SetEndTime(v string) *DescribeDataBackupsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeDataBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataBackupsRequest
	GetPageSize() *int32
	SetStartTime(v string) *DescribeDataBackupsRequest
	GetStartTime() *string
}

type DescribeDataBackupsRequest struct {
	// The ID of the backup set. If you specify BackupId, the details of the backup set are returned.
	//
	// > You can call the [DescribeDataBackups](https://help.aliyun.com/document_detail/210093.html) operation to query the information about all backup sets of an instance, including backup set IDs.
	//
	// example:
	//
	// 327329803
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- Automated
	//
	// 	- Manual
	//
	// If you do not specify this parameter, all backup sets are returned.
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The state of the backup set. Valid values:
	//
	// 	- Success
	//
	// 	- Failed
	//
	// If you do not specify this parameter, all backup sets are returned.
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The backup type. Valid values:
	//
	// 	- **DATA**: full backup.
	//
	// 	- **RESTOREPOI**: point-in-time recovery backup.
	//
	// If you do not specify this parameter, the backup sets of full backup are returned.
	//
	// example:
	//
	// DATA
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2011-06-01T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Pages start from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- 30
	//
	// 	- 50
	//
	// 	- 100
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2011-06-01T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeDataBackupsRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeDataBackupsRequest) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeDataBackupsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDataBackupsRequest) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDataBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDataBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataBackupsRequest) SetBackupId(v string) *DescribeDataBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetBackupMode(v string) *DescribeDataBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetBackupStatus(v string) *DescribeDataBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetDBInstanceId(v string) *DescribeDataBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetDataType(v string) *DescribeDataBackupsRequest {
	s.DataType = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetEndTime(v string) *DescribeDataBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetPageNumber(v int32) *DescribeDataBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetPageSize(v int32) *DescribeDataBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataBackupsRequest) SetStartTime(v string) *DescribeDataBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataBackupsRequest) Validate() error {
	return dara.Validate(s)
}
