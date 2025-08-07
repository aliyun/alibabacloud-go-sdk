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
	SetBackupMode(v string) *DescribeBackupTasksRequest
	GetBackupMode() *string
	SetDBClusterId(v string) *DescribeBackupTasksRequest
	GetDBClusterId() *string
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
	// 11111111
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Manual
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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

func (s *DescribeBackupTasksRequest) GetBackupJobId() *string {
	return s.BackupJobId
}

func (s *DescribeBackupTasksRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DescribeBackupTasksRequest) SetBackupJobId(v string) *DescribeBackupTasksRequest {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetBackupMode(v string) *DescribeBackupTasksRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetDBClusterId(v string) *DescribeBackupTasksRequest {
	s.DBClusterId = &v
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
