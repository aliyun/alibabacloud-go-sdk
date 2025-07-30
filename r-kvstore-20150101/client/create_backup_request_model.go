// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRetentionPeriod(v int64) *CreateBackupRequest
	GetBackupRetentionPeriod() *int64
	SetInstanceId(v string) *CreateBackupRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CreateBackupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateBackupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateBackupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateBackupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateBackupRequest
	GetSecurityToken() *string
}

type CreateBackupRequest struct {
	BackupRetentionPeriod *int64 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) GetBackupRetentionPeriod() *int64 {
	return s.BackupRetentionPeriod
}

func (s *CreateBackupRequest) GetInstanceId() *string {
	return s.InstanceId
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

func (s *CreateBackupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateBackupRequest) SetBackupRetentionPeriod(v int64) *CreateBackupRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *CreateBackupRequest) SetInstanceId(v string) *CreateBackupRequest {
	s.InstanceId = &v
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

func (s *CreateBackupRequest) SetSecurityToken(v string) *CreateBackupRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateBackupRequest) Validate() error {
	return dara.Validate(s)
}
