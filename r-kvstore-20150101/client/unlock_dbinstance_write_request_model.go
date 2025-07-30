// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockDBInstanceWriteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UnlockDBInstanceWriteRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *UnlockDBInstanceWriteRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnlockDBInstanceWriteRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UnlockDBInstanceWriteRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnlockDBInstanceWriteRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UnlockDBInstanceWriteRequest
	GetSecurityToken() *string
}

type UnlockDBInstanceWriteRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UnlockDBInstanceWriteRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockDBInstanceWriteRequest) GoString() string {
	return s.String()
}

func (s *UnlockDBInstanceWriteRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UnlockDBInstanceWriteRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnlockDBInstanceWriteRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnlockDBInstanceWriteRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnlockDBInstanceWriteRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnlockDBInstanceWriteRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UnlockDBInstanceWriteRequest) SetDBInstanceId(v string) *UnlockDBInstanceWriteRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UnlockDBInstanceWriteRequest) SetOwnerAccount(v string) *UnlockDBInstanceWriteRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnlockDBInstanceWriteRequest) SetOwnerId(v int64) *UnlockDBInstanceWriteRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockDBInstanceWriteRequest) SetResourceOwnerAccount(v string) *UnlockDBInstanceWriteRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnlockDBInstanceWriteRequest) SetResourceOwnerId(v int64) *UnlockDBInstanceWriteRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnlockDBInstanceWriteRequest) SetSecurityToken(v string) *UnlockDBInstanceWriteRequest {
	s.SecurityToken = &v
	return s
}

func (s *UnlockDBInstanceWriteRequest) Validate() error {
	return dara.Validate(s)
}
