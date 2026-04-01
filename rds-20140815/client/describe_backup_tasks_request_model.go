// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupJobId(v int32) *DescribeBackupTasksRequest
	GetBackupJobId() *int32
	SetBackupJobStatus(v string) *DescribeBackupTasksRequest
	GetBackupJobStatus() *string
	SetBackupMode(v string) *DescribeBackupTasksRequest
	GetBackupMode() *string
	SetClientToken(v string) *DescribeBackupTasksRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeBackupTasksRequest
	GetDBInstanceId() *string
	SetFlag(v string) *DescribeBackupTasksRequest
	GetFlag() *string
	SetOwnerAccount(v string) *DescribeBackupTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupTasksRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeBackupTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupTasksRequest
	GetResourceOwnerId() *int64
}

type DescribeBackupTasksRequest struct {
	// The ID of the backup task.
	//
	// example:
	//
	// 4762614
	BackupJobId *int32 `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The status of the backup task. Valid values:
	//
	// 	- **NoStart**
	//
	// 	- **Progressing**
	//
	// By default, this operation returns backup tasks in both states.
	//
	// example:
	//
	// NoStart
	BackupJobStatus *string `json:"BackupJobStatus,omitempty" xml:"BackupJobStatus,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Automated
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// Specifies the client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	Flag                 *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeBackupTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksRequest) GetBackupJobId() *int32 {
	return s.BackupJobId
}

func (s *DescribeBackupTasksRequest) GetBackupJobStatus() *string {
	return s.BackupJobStatus
}

func (s *DescribeBackupTasksRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupTasksRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeBackupTasksRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBackupTasksRequest) GetFlag() *string {
	return s.Flag
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

func (s *DescribeBackupTasksRequest) SetBackupJobId(v int32) *DescribeBackupTasksRequest {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetBackupJobStatus(v string) *DescribeBackupTasksRequest {
	s.BackupJobStatus = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetBackupMode(v string) *DescribeBackupTasksRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetClientToken(v string) *DescribeBackupTasksRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetDBInstanceId(v string) *DescribeBackupTasksRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetFlag(v string) *DescribeBackupTasksRequest {
	s.Flag = &v
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

func (s *DescribeBackupTasksRequest) Validate() error {
	return dara.Validate(s)
}
