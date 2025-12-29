// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupExpireTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupExpireTime(v string) *ModifyBackupExpireTimeRequest
	GetBackupExpireTime() *string
	SetBackupId(v string) *ModifyBackupExpireTimeRequest
	GetBackupId() *string
	SetDBInstanceId(v string) *ModifyBackupExpireTimeRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyBackupExpireTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyBackupExpireTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyBackupExpireTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyBackupExpireTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyBackupExpireTimeRequest struct {
	// example:
	//
	// 2025-03-29T03:47:12Z
	BackupExpireTime *string `json:"BackupExpireTime,omitempty" xml:"BackupExpireTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 260032xxxx
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dds-bp16cb162771****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyBackupExpireTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupExpireTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupExpireTimeRequest) GetBackupExpireTime() *string {
	return s.BackupExpireTime
}

func (s *ModifyBackupExpireTimeRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ModifyBackupExpireTimeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyBackupExpireTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyBackupExpireTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyBackupExpireTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyBackupExpireTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyBackupExpireTimeRequest) SetBackupExpireTime(v string) *ModifyBackupExpireTimeRequest {
	s.BackupExpireTime = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetBackupId(v string) *ModifyBackupExpireTimeRequest {
	s.BackupId = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetDBInstanceId(v string) *ModifyBackupExpireTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetOwnerAccount(v string) *ModifyBackupExpireTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetOwnerId(v int64) *ModifyBackupExpireTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetResourceOwnerAccount(v string) *ModifyBackupExpireTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) SetResourceOwnerId(v int64) *ModifyBackupExpireTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupExpireTimeRequest) Validate() error {
	return dara.Validate(s)
}
