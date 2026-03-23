// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveBackupSize(v int64) *DescribeResourceUsageResponseBody
	GetArchiveBackupSize() *int64
	SetBackupDataSize(v int64) *DescribeResourceUsageResponseBody
	GetBackupDataSize() *int64
	SetBackupEcsSnapshotSize(v string) *DescribeResourceUsageResponseBody
	GetBackupEcsSnapshotSize() *string
	SetBackupLogSize(v int64) *DescribeResourceUsageResponseBody
	GetBackupLogSize() *int64
	SetBackupOssDataSize(v int64) *DescribeResourceUsageResponseBody
	GetBackupOssDataSize() *int64
	SetBackupOssLogSize(v int64) *DescribeResourceUsageResponseBody
	GetBackupOssLogSize() *int64
	SetBackupSize(v int64) *DescribeResourceUsageResponseBody
	GetBackupSize() *int64
	SetColdBackupSize(v int64) *DescribeResourceUsageResponseBody
	GetColdBackupSize() *int64
	SetDBInstanceId(v string) *DescribeResourceUsageResponseBody
	GetDBInstanceId() *string
	SetDataSize(v int64) *DescribeResourceUsageResponseBody
	GetDataSize() *int64
	SetDiskUsed(v int64) *DescribeResourceUsageResponseBody
	GetDiskUsed() *int64
	SetEngine(v string) *DescribeResourceUsageResponseBody
	GetEngine() *string
	SetLogSize(v int64) *DescribeResourceUsageResponseBody
	GetLogSize() *int64
	SetPaidBackupSize(v int64) *DescribeResourceUsageResponseBody
	GetPaidBackupSize() *int64
	SetRequestId(v string) *DescribeResourceUsageResponseBody
	GetRequestId() *string
	SetSQLSize(v int64) *DescribeResourceUsageResponseBody
	GetSQLSize() *int64
}

type DescribeResourceUsageResponseBody struct {
	ArchiveBackupSize     *int64  `json:"ArchiveBackupSize,omitempty" xml:"ArchiveBackupSize,omitempty"`
	BackupDataSize        *int64  `json:"BackupDataSize,omitempty" xml:"BackupDataSize,omitempty"`
	BackupEcsSnapshotSize *string `json:"BackupEcsSnapshotSize,omitempty" xml:"BackupEcsSnapshotSize,omitempty"`
	BackupLogSize         *int64  `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	BackupOssDataSize     *int64  `json:"BackupOssDataSize,omitempty" xml:"BackupOssDataSize,omitempty"`
	BackupOssLogSize      *int64  `json:"BackupOssLogSize,omitempty" xml:"BackupOssLogSize,omitempty"`
	BackupSize            *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	ColdBackupSize        *int64  `json:"ColdBackupSize,omitempty" xml:"ColdBackupSize,omitempty"`
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DataSize              *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	DiskUsed              *int64  `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	LogSize               *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	PaidBackupSize        *int64  `json:"PaidBackupSize,omitempty" xml:"PaidBackupSize,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLSize               *int64  `json:"SQLSize,omitempty" xml:"SQLSize,omitempty"`
}

func (s DescribeResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponseBody) GetArchiveBackupSize() *int64 {
	return s.ArchiveBackupSize
}

func (s *DescribeResourceUsageResponseBody) GetBackupDataSize() *int64 {
	return s.BackupDataSize
}

func (s *DescribeResourceUsageResponseBody) GetBackupEcsSnapshotSize() *string {
	return s.BackupEcsSnapshotSize
}

func (s *DescribeResourceUsageResponseBody) GetBackupLogSize() *int64 {
	return s.BackupLogSize
}

func (s *DescribeResourceUsageResponseBody) GetBackupOssDataSize() *int64 {
	return s.BackupOssDataSize
}

func (s *DescribeResourceUsageResponseBody) GetBackupOssLogSize() *int64 {
	return s.BackupOssLogSize
}

func (s *DescribeResourceUsageResponseBody) GetBackupSize() *int64 {
	return s.BackupSize
}

func (s *DescribeResourceUsageResponseBody) GetColdBackupSize() *int64 {
	return s.ColdBackupSize
}

func (s *DescribeResourceUsageResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeResourceUsageResponseBody) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DescribeResourceUsageResponseBody) GetDiskUsed() *int64 {
	return s.DiskUsed
}

func (s *DescribeResourceUsageResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeResourceUsageResponseBody) GetLogSize() *int64 {
	return s.LogSize
}

func (s *DescribeResourceUsageResponseBody) GetPaidBackupSize() *int64 {
	return s.PaidBackupSize
}

func (s *DescribeResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceUsageResponseBody) GetSQLSize() *int64 {
	return s.SQLSize
}

func (s *DescribeResourceUsageResponseBody) SetArchiveBackupSize(v int64) *DescribeResourceUsageResponseBody {
	s.ArchiveBackupSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetBackupDataSize(v int64) *DescribeResourceUsageResponseBody {
	s.BackupDataSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetBackupEcsSnapshotSize(v string) *DescribeResourceUsageResponseBody {
	s.BackupEcsSnapshotSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetBackupLogSize(v int64) *DescribeResourceUsageResponseBody {
	s.BackupLogSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetBackupOssDataSize(v int64) *DescribeResourceUsageResponseBody {
	s.BackupOssDataSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetBackupOssLogSize(v int64) *DescribeResourceUsageResponseBody {
	s.BackupOssLogSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetBackupSize(v int64) *DescribeResourceUsageResponseBody {
	s.BackupSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetColdBackupSize(v int64) *DescribeResourceUsageResponseBody {
	s.ColdBackupSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDBInstanceId(v string) *DescribeResourceUsageResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDataSize(v int64) *DescribeResourceUsageResponseBody {
	s.DataSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetDiskUsed(v int64) *DescribeResourceUsageResponseBody {
	s.DiskUsed = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetEngine(v string) *DescribeResourceUsageResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetLogSize(v int64) *DescribeResourceUsageResponseBody {
	s.LogSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetPaidBackupSize(v int64) *DescribeResourceUsageResponseBody {
	s.PaidBackupSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetRequestId(v string) *DescribeResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) SetSQLSize(v int64) *DescribeResourceUsageResponseBody {
	s.SQLSize = &v
	return s
}

func (s *DescribeResourceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
