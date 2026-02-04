// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceCLSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceCLSRequest
	GetDBInstanceId() *string
	SetEncryptionAlgorithm(v string) *ModifyDBInstanceCLSRequest
	GetEncryptionAlgorithm() *string
	SetEncryptionKey(v string) *ModifyDBInstanceCLSRequest
	GetEncryptionKey() *string
	SetEncryptionKeyMode(v string) *ModifyDBInstanceCLSRequest
	GetEncryptionKeyMode() *string
	SetEncryptionStatus(v string) *ModifyDBInstanceCLSRequest
	GetEncryptionStatus() *string
	SetIsRotate(v bool) *ModifyDBInstanceCLSRequest
	GetIsRotate() *bool
	SetOwnerAccount(v string) *ModifyDBInstanceCLSRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceCLSRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceCLSRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceCLSRequest
	GetResourceOwnerId() *int64
	SetRoleArn(v string) *ModifyDBInstanceCLSRequest
	GetRoleArn() *string
	SetWhiteListMode(v bool) *ModifyDBInstanceCLSRequest
	GetWhiteListMode() *bool
}

type ModifyDBInstanceCLSRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n8t18o******6d5
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// AES_256_GCM
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" xml:"EncryptionAlgorithm,omitempty"`
	// example:
	//
	// acs:kms:cn-hangzhou:123456789:key/xxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// example:
	//
	// KMS
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	EncryptionStatus *string `json:"EncryptionStatus,omitempty" xml:"EncryptionStatus,omitempty"`
	// example:
	//
	// true
	IsRotate             *bool   `json:"IsRotate,omitempty" xml:"IsRotate,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// acs:123456789:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// example:
	//
	// true
	WhiteListMode *bool `json:"WhiteListMode,omitempty" xml:"WhiteListMode,omitempty"`
}

func (s ModifyDBInstanceCLSRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceCLSRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceCLSRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceCLSRequest) GetEncryptionAlgorithm() *string {
	return s.EncryptionAlgorithm
}

func (s *ModifyDBInstanceCLSRequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *ModifyDBInstanceCLSRequest) GetEncryptionKeyMode() *string {
	return s.EncryptionKeyMode
}

func (s *ModifyDBInstanceCLSRequest) GetEncryptionStatus() *string {
	return s.EncryptionStatus
}

func (s *ModifyDBInstanceCLSRequest) GetIsRotate() *bool {
	return s.IsRotate
}

func (s *ModifyDBInstanceCLSRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceCLSRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceCLSRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceCLSRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceCLSRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ModifyDBInstanceCLSRequest) GetWhiteListMode() *bool {
	return s.WhiteListMode
}

func (s *ModifyDBInstanceCLSRequest) SetDBInstanceId(v string) *ModifyDBInstanceCLSRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetEncryptionAlgorithm(v string) *ModifyDBInstanceCLSRequest {
	s.EncryptionAlgorithm = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetEncryptionKey(v string) *ModifyDBInstanceCLSRequest {
	s.EncryptionKey = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetEncryptionKeyMode(v string) *ModifyDBInstanceCLSRequest {
	s.EncryptionKeyMode = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetEncryptionStatus(v string) *ModifyDBInstanceCLSRequest {
	s.EncryptionStatus = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetIsRotate(v bool) *ModifyDBInstanceCLSRequest {
	s.IsRotate = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetOwnerAccount(v string) *ModifyDBInstanceCLSRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetOwnerId(v int64) *ModifyDBInstanceCLSRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceCLSRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceCLSRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetRoleArn(v string) *ModifyDBInstanceCLSRequest {
	s.RoleArn = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) SetWhiteListMode(v bool) *ModifyDBInstanceCLSRequest {
	s.WhiteListMode = &v
	return s
}

func (s *ModifyDBInstanceCLSRequest) Validate() error {
	return dara.Validate(s)
}
