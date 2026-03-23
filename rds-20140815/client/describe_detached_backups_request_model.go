// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDetachedBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeDetachedBackupsRequest
	GetBackupId() *string
	SetBackupMode(v string) *DescribeDetachedBackupsRequest
	GetBackupMode() *string
	SetBackupStatus(v string) *DescribeDetachedBackupsRequest
	GetBackupStatus() *string
	SetDBInstanceId(v string) *DescribeDetachedBackupsRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDetachedBackupsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *DescribeDetachedBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDetachedBackupsRequest
	GetPageSize() *int32
	SetRegion(v string) *DescribeDetachedBackupsRequest
	GetRegion() *string
	SetResourceGroupId(v string) *DescribeDetachedBackupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *DescribeDetachedBackupsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDetachedBackupsRequest
	GetStartTime() *string
}

type DescribeDetachedBackupsRequest struct {
	BackupId     *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupMode   *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDetachedBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetachedBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeDetachedBackupsRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeDetachedBackupsRequest) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeDetachedBackupsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDetachedBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDetachedBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDetachedBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDetachedBackupsRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDetachedBackupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDetachedBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDetachedBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDetachedBackupsRequest) SetBackupId(v string) *DescribeDetachedBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupMode(v string) *DescribeDetachedBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupStatus(v string) *DescribeDetachedBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetDBInstanceId(v string) *DescribeDetachedBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetEndTime(v string) *DescribeDetachedBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetPageNumber(v int32) *DescribeDetachedBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetPageSize(v int32) *DescribeDetachedBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetRegion(v string) *DescribeDetachedBackupsRequest {
	s.Region = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetResourceGroupId(v string) *DescribeDetachedBackupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetResourceOwnerId(v int64) *DescribeDetachedBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetStartTime(v string) *DescribeDetachedBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) Validate() error {
	return dara.Validate(s)
}
