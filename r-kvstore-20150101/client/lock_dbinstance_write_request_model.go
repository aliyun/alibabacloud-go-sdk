// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockDBInstanceWriteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *LockDBInstanceWriteRequest
	GetDBInstanceId() *string
	SetLockReason(v string) *LockDBInstanceWriteRequest
	GetLockReason() *string
	SetOwnerAccount(v string) *LockDBInstanceWriteRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *LockDBInstanceWriteRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *LockDBInstanceWriteRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *LockDBInstanceWriteRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *LockDBInstanceWriteRequest
	GetSecurityToken() *string
}

type LockDBInstanceWriteRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The reason why write operations on the instance are locked.
	//
	// This parameter is required.
	//
	// example:
	//
	// lock reason
	LockReason           *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s LockDBInstanceWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s LockDBInstanceWriteRequest) GoString() string {
	return s.String()
}

func (s *LockDBInstanceWriteRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *LockDBInstanceWriteRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *LockDBInstanceWriteRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *LockDBInstanceWriteRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LockDBInstanceWriteRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *LockDBInstanceWriteRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *LockDBInstanceWriteRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *LockDBInstanceWriteRequest) SetDBInstanceId(v string) *LockDBInstanceWriteRequest {
	s.DBInstanceId = &v
	return s
}

func (s *LockDBInstanceWriteRequest) SetLockReason(v string) *LockDBInstanceWriteRequest {
	s.LockReason = &v
	return s
}

func (s *LockDBInstanceWriteRequest) SetOwnerAccount(v string) *LockDBInstanceWriteRequest {
	s.OwnerAccount = &v
	return s
}

func (s *LockDBInstanceWriteRequest) SetOwnerId(v int64) *LockDBInstanceWriteRequest {
	s.OwnerId = &v
	return s
}

func (s *LockDBInstanceWriteRequest) SetResourceOwnerAccount(v string) *LockDBInstanceWriteRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LockDBInstanceWriteRequest) SetResourceOwnerId(v int64) *LockDBInstanceWriteRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LockDBInstanceWriteRequest) SetSecurityToken(v string) *LockDBInstanceWriteRequest {
	s.SecurityToken = &v
	return s
}

func (s *LockDBInstanceWriteRequest) Validate() error {
	return dara.Validate(s)
}
