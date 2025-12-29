// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DeleteBackupRequest
	GetBackupId() *string
	SetDBInstanceId(v string) *DeleteBackupRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DeleteBackupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteBackupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteBackupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteBackupRequest
	GetResourceOwnerId() *int64
}

type DeleteBackupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5664****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dds-bp11483712c1****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DeleteBackupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteBackupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteBackupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteBackupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteBackupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteBackupRequest) SetBackupId(v string) *DeleteBackupRequest {
	s.BackupId = &v
	return s
}

func (s *DeleteBackupRequest) SetDBInstanceId(v string) *DeleteBackupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteBackupRequest) SetOwnerAccount(v string) *DeleteBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBackupRequest) SetOwnerId(v int64) *DeleteBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBackupRequest) SetResourceOwnerAccount(v string) *DeleteBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBackupRequest) SetResourceOwnerId(v int64) *DeleteBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBackupRequest) Validate() error {
	return dara.Validate(s)
}
