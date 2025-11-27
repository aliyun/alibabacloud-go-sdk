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
	// The storage that is occupied by archived backup files on the instance. Unit: bytes.
	//
	// example:
	//
	// 0
	ArchiveBackupSize *int64 `json:"ArchiveBackupSize,omitempty" xml:"ArchiveBackupSize,omitempty"`
	// The storage that is occupied by data backup files, excluding archived backup files, on the instance. Unit: bytes.
	//
	// example:
	//
	// 94324736
	BackupDataSize *int64 `json:"BackupDataSize,omitempty" xml:"BackupDataSize,omitempty"`
	// The storage capacity that is used to store the snapshot backup files of the **RDS for SQL Server*	- instance. Unit: bytes. The value 0 indicates that no snapshot backup files are stored for the instance.
	//
	// example:
	//
	// 0
	BackupEcsSnapshotSize *string `json:"BackupEcsSnapshotSize,omitempty" xml:"BackupEcsSnapshotSize,omitempty"`
	// The storage that is occupied by log backup files, excluding archived backup files, on the instance. Unit: bytes.
	//
	// example:
	//
	// 45145563
	BackupLogSize *int64 `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	// The size of data backup files that are stored in Object Storage Service (OSS) buckets. Unit: bytes. The value 0 indicates no data backup files are stored in OSS buckets.
	//
	// example:
	//
	// 8821760
	BackupOssDataSize *int64 `json:"BackupOssDataSize,omitempty" xml:"BackupOssDataSize,omitempty"`
	// The size of log backup files that are stored in OSS buckets. Unit: bytes. The value 0 indicates no log backup files are stored in OSS buckets.
	//
	// example:
	//
	// 44180999
	BackupOssLogSize *int64 `json:"BackupOssLogSize,omitempty" xml:"BackupOssLogSize,omitempty"`
	// The storage that is used to store backup files. Unit: bytes. The value -1 indicates that no backup files are stored.
	//
	// example:
	//
	// 53002759
	BackupSize *int64 `json:"BackupSize,omitempty" xml:"BackupSize,omitempty"`
	// The storage that is used to store cold backup files. Unit: bytes. The value -1 indicates that no cold backup files are stored.
	//
	// example:
	//
	// 2337275904
	ColdBackupSize *int64 `json:"ColdBackupSize,omitempty" xml:"ColdBackupSize,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The storage that is used to store data files. Unit: bytes. The value -1 indicates that no data files are stored.
	//
	// example:
	//
	// 1292094741
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The total storage that is occupied by data files and log files on the instance. Unit: bytes. The value -1 indicates that no data files or log files are stored on the instance.
	//
	// example:
	//
	// 2337275904
	DiskUsed *int64 `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The storage that is used to store log files. Unit: bytes. The value -1 indicates that no log files are stored.
	//
	// example:
	//
	// 1045181163
	LogSize *int64 `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	// The backup storage for which you must pay. The system provides a free quota on backup storage. You must pay for the backup storage that exceeds the free quota. Unit: bytes.
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
	// The storage that is occupied to execute SQL statements on the instance. Unit: bytes. The value -1 indicates that no SQL statements are executed.
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
