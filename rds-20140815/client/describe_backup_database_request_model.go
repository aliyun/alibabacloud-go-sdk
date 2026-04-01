// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeBackupDatabaseRequest
	GetBackupId() *string
	SetDBInstanceId(v string) *DescribeBackupDatabaseRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *DescribeBackupDatabaseRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeBackupDatabaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupDatabaseRequest
	GetResourceOwnerId() *int64
}

type DescribeBackupDatabaseRequest struct {
	// The ID of the backup set.
	//
	// example:
	//
	// 90262212
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupDatabaseRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeBackupDatabaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupDatabaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupDatabaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupDatabaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupDatabaseRequest) SetBackupId(v string) *DescribeBackupDatabaseRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupDatabaseRequest) SetDBInstanceId(v string) *DescribeBackupDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupDatabaseRequest) SetOwnerId(v int64) *DescribeBackupDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupDatabaseRequest) SetResourceOwnerAccount(v string) *DescribeBackupDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupDatabaseRequest) SetResourceOwnerId(v int64) *DescribeBackupDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
