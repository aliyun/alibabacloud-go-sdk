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
	SetInstanceId(v string) *DescribeBackupTasksRequest
	GetInstanceId() *string
	SetJobMode(v string) *DescribeBackupTasksRequest
	GetJobMode() *string
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
	// >  If you call the [CreateBackup](https://help.aliyun.com/document_detail/473819.html) operation to perform a manual backup task, you can set this parameter to the returned backup ID to query the backup progress of the task.
	//
	// example:
	//
	// 1162****
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The instance ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**: automatic backup. You can call the [DescribeBackupPolicy](https://help.aliyun.com/document_detail/473822.html) operation to query the automatic backup policy.
	//
	// 	- **Manual**: manual backup.
	//
	// >  By default, the information about backup tasks in both modes is returned.
	//
	// example:
	//
	// Manual
	JobMode              *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
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

func (s *DescribeBackupTasksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupTasksRequest) GetJobMode() *string {
	return s.JobMode
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

func (s *DescribeBackupTasksRequest) SetInstanceId(v string) *DescribeBackupTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetJobMode(v string) *DescribeBackupTasksRequest {
	s.JobMode = &v
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
