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
	// The storage consumed by archived backups. Unit: bytes.
	//
	// example:
	//
	// 0
	ArchiveBackupSize *int64 `json:"ArchiveBackupSize,omitempty" xml:"ArchiveBackupSize,omitempty"`
	// The total storage consumed by data backups, excluding archived backups. Unit: bytes.
	//
	// > For **SQL Server*	- instances, this value indicates the total size of physical backups and snapshot backups.
	//
	// example:
	//
	// 94324736
	BackupDataSize *int64 `json:"BackupDataSize,omitempty" xml:"BackupDataSize,omitempty"`
	// The storage consumed by snapshot backups for **SQL Server instances**. Unit: bytes. A value of 0 indicates no data.
	//
	// example:
	//
	// 0
	BackupEcsSnapshotSize *string `json:"BackupEcsSnapshotSize,omitempty" xml:"BackupEcsSnapshotSize,omitempty"`
	// The total storage consumed by log backups, excluding archived backups. Unit: bytes.
	//
	// example:
	//
	// 45145563
	BackupLogSize *int64 `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	// The size of data files in backup sets stored in OSS. Unit: bytes. A value of 0 indicates no data.
	//
	// > For **SQL Server*	- instances, this value indicates the storage consumed by physical backups.
	//
	// example:
	//
	// 8821760
	BackupOssDataSize *int64 `json:"BackupOssDataSize,omitempty" xml:"BackupOssDataSize,omitempty"`
	// The size of log files in backup sets stored in OSS. Unit: bytes. A value of 0 indicates no data.
	//
	// example:
	//
	// 44180999
	BackupOssLogSize *int64 `json:"BackupOssLogSize,omitempty" xml:"BackupOssLogSize,omitempty"`
	// The storage consumed by backups (data backups + log backups). Unit: bytes. A value of -1 indicates no data.
	//
	// example:
	//
	// 53002759
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The storage consumed by cold backups. Unit: bytes. A value of -1 indicates no data.
	//
	// example:
	//
	// 2337275904
	ColdBackupSize *int64 `json:"ColdBackupSize,omitempty" xml:"ColdBackupSize,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The storage consumed by data files. Unit: bytes. A value of -1 indicates no data.
	//
	// example:
	//
	// 1292094741
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The used storage (DataSize + LogSize). Unit: bytes. A value of -1 indicates no data.
	//
	// example:
	//
	// 2337275904
	DiskUsed *int64 `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The storage consumed by log files. Unit: bytes. A value of -1 indicates no data.
	//
	// example:
	//
	// 1045181163
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The billable storage consumed by backups after the free quota is deducted. Unit: bytes.
	//
	// example:
	//
	// 0
	PaidBackupSize *int64 `json:"PaidBackupSize,omitempty" xml:"PaidBackupSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F937E173-559C-4498-8D90-38D32342B9E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage consumed by SQL data. Unit: bytes. A value of -1 indicates no data.
	//
	// example:
	//
	// 315052751
	SQLSize *int64 `json:"SQLSize,omitempty" xml:"SQLSize,omitempty"`
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
