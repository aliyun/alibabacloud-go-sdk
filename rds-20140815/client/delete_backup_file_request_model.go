// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DeleteBackupFileRequest
	GetBackupId() *string
	SetBackupTime(v string) *DeleteBackupFileRequest
	GetBackupTime() *string
	SetDBInstanceId(v string) *DeleteBackupFileRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DeleteBackupFileRequest
	GetDBName() *string
	SetOwnerId(v int64) *DeleteBackupFileRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteBackupFileRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteBackupFileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteBackupFileRequest
	GetResourceOwnerId() *int64
}

type DeleteBackupFileRequest struct {
	// You can specify only the ID of a backup file whose backup policy is Single-database Backup. You can specify the IDs of up to 100 backup files at a time. Separate the IDs with commas (,). You can call the DescribeBackups operation to query the IDs of data backup files.
	//
	// example:
	//
	// 29304****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The time before which the backup files you want to delete are generated. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2011-06-11T16:00:00Z
	BackupTime *string `json:"BackupTime,omitempty" xml:"BackupTime,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp6wjk5******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb
	DBName  *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeDBInstanceAttribute operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteBackupFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DeleteBackupFileRequest) GetBackupTime() *string {
	return s.BackupTime
}

func (s *DeleteBackupFileRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteBackupFileRequest) GetDBName() *string {
	return s.DBName
}

func (s *DeleteBackupFileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteBackupFileRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteBackupFileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteBackupFileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBackupFileRequest) SetBackupId(v string) *DeleteBackupFileRequest {
	s.BackupId = &v
	return s
}

func (s *DeleteBackupFileRequest) SetBackupTime(v string) *DeleteBackupFileRequest {
	s.BackupTime = &v
	return s
}

func (s *DeleteBackupFileRequest) SetDBInstanceId(v string) *DeleteBackupFileRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteBackupFileRequest) SetDBName(v string) *DeleteBackupFileRequest {
	s.DBName = &v
	return s
}

func (s *DeleteBackupFileRequest) SetOwnerId(v int64) *DeleteBackupFileRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBackupFileRequest) SetRegionId(v string) *DeleteBackupFileRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteBackupFileRequest) SetResourceOwnerAccount(v string) *DeleteBackupFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBackupFileRequest) SetResourceOwnerId(v int64) *DeleteBackupFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBackupFileRequest) Validate() error {
	return dara.Validate(s)
}
