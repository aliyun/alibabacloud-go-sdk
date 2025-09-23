// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEncryptionDBSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyEncryptionDBSecretRequest
	GetDBClusterId() *string
	SetEncryptionDBStatus(v string) *ModifyEncryptionDBSecretRequest
	GetEncryptionDBStatus() *string
	SetEncryptionKey(v string) *ModifyEncryptionDBSecretRequest
	GetEncryptionKey() *string
	SetOwnerAccount(v string) *ModifyEncryptionDBSecretRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyEncryptionDBSecretRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyEncryptionDBSecretRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyEncryptionDBSecretRequest
	GetResourceOwnerId() *int64
	SetRoleArn(v string) *ModifyEncryptionDBSecretRequest
	GetRoleArn() *string
}

type ModifyEncryptionDBSecretRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Enabled
	EncryptionDBStatus *string `json:"EncryptionDBStatus,omitempty" xml:"EncryptionDBStatus,omitempty"`
	// example:
	//
	// 749c1df7-****-****-****-*********
	EncryptionKey        *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// acs:ram::1406926*****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s ModifyEncryptionDBSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEncryptionDBSecretRequest) GoString() string {
	return s.String()
}

func (s *ModifyEncryptionDBSecretRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyEncryptionDBSecretRequest) GetEncryptionDBStatus() *string {
	return s.EncryptionDBStatus
}

func (s *ModifyEncryptionDBSecretRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *ModifyEncryptionDBSecretRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyEncryptionDBSecretRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyEncryptionDBSecretRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyEncryptionDBSecretRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyEncryptionDBSecretRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ModifyEncryptionDBSecretRequest) SetDBClusterId(v string) *ModifyEncryptionDBSecretRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyEncryptionDBSecretRequest) SetEncryptionDBStatus(v string) *ModifyEncryptionDBSecretRequest {
	s.EncryptionDBStatus = &v
	return s
}

func (s *ModifyEncryptionDBSecretRequest) SetEncryptionKey(v string) *ModifyEncryptionDBSecretRequest {
	s.EncryptionKey = &v
	return s
}

func (s *ModifyEncryptionDBSecretRequest) SetOwnerAccount(v string) *ModifyEncryptionDBSecretRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyEncryptionDBSecretRequest) SetOwnerId(v int64) *ModifyEncryptionDBSecretRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyEncryptionDBSecretRequest) SetResourceOwnerAccount(v string) *ModifyEncryptionDBSecretRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyEncryptionDBSecretRequest) SetResourceOwnerId(v int64) *ModifyEncryptionDBSecretRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyEncryptionDBSecretRequest) SetRoleArn(v string) *ModifyEncryptionDBSecretRequest {
	s.RoleArn = &v
	return s
}

func (s *ModifyEncryptionDBSecretRequest) Validate() error {
	return dara.Validate(s)
}
