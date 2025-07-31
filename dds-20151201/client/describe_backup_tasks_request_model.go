// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobId(v string) *DescribeBackupTasksRequest
	GetBackupJobId() *string
	SetDBInstanceId(v string) *DescribeBackupTasksRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeBackupTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupTasksRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeBackupTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeBackupTasksRequest
	GetSecurityToken() *string
}

type DescribeBackupTasksRequest struct {
	// The ID of the backup task.
	//
	// >  If you call the [CreateBackup](https://help.aliyun.com/document_detail/468439.html) operation to perform a manual backup task, you can set this parameter to the returned backup ID to query the backup progress of the task.
	//
	// example:
	//
	// 170034
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-2zeb2d64cb46xxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeBackupTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksRequest) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *DescribeBackupTasksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeBackupTasksRequest) SetBackupJobId(v string) *DescribeBackupTasksRequest {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetDBInstanceId(v string) *DescribeBackupTasksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetOwnerAccount(v string) *DescribeBackupTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetOwnerId(v int64) *DescribeBackupTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetResourceOwnerAccount(v string) *DescribeBackupTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetResourceOwnerId(v int64) *DescribeBackupTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetSecurityToken(v string) *DescribeBackupTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeBackupTasksRequest) Validate() error {
	return dara.Validate(s)
}
