// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMethod(v string) *CreateBackupRequest
	GetBackupMethod() *string
	SetBackupRetentionPeriod(v int64) *CreateBackupRequest
	GetBackupRetentionPeriod() *int64
	SetDBInstanceId(v string) *CreateBackupRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *CreateBackupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateBackupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateBackupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateBackupRequest
	GetResourceOwnerId() *int64
}

type CreateBackupRequest struct {
	// The backup method of the instance. Valid values:
	//
	// - **Snapshot**: snapshot backup
	//
	// - **Logical**: logical backup
	//
	// - **Physical**: physical backup (default)
	//
	// > - Cloud disk instances support only snapshot backup.
	//
	// > - Only replica set and sharded cluster instances support selecting the backup method. You do not need to specify this parameter for standalone instances. The backup method is fixed to snapshot backup.
	//
	// example:
	//
	// Logical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The retention period of the backup. This parameter is applicable only to replica set and sharded cluster instances with cloud disks.
	//
	// > If you do not specify this parameter, the default backup policy applies. Valid values:
	//
	// > - 7 to 730: 7 to 730 days
	//
	// > - -1: long-term retention
	//
	// example:
	//
	// 30
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp2235****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *CreateBackupRequest) GetBackupRetentionPeriod() *int64 {
	return s.BackupRetentionPeriod
}

func (s *CreateBackupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateBackupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateBackupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateBackupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateBackupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateBackupRequest) SetBackupMethod(v string) *CreateBackupRequest {
	s.BackupMethod = &v
	return s
}

func (s *CreateBackupRequest) SetBackupRetentionPeriod(v int64) *CreateBackupRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *CreateBackupRequest) SetDBInstanceId(v string) *CreateBackupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerAccount(v string) *CreateBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerId(v int64) *CreateBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerAccount(v string) *CreateBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerId(v int64) *CreateBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBackupRequest) Validate() error {
	return dara.Validate(s)
}
