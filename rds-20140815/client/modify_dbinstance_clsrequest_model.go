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
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n8t18o******6d5
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Encryption algorithm to use. Valid values:
	//
	// - AES_128_CBC
	//
	// - AES_128_GCM
	//
	// - AES_128_CTR
	//
	// - AES_128_ECB
	//
	// - AES_256_CBC
	//
	// - AES_256_GCM
	//
	// - AES_256_CTR
	//
	// - AES_256_ECB
	//
	// - SM4_128_CBC
	//
	// - SM4_128_GCM
	//
	// - SM4_128_CTR
	//
	// - SM4_128_ECB
	//
	// example:
	//
	// AES_256_GCM
	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" xml:"EncryptionAlgorithm,omitempty"`
	// Encryption key ID. This parameter is required when using a KMS key.
	//
	// example:
	//
	// 749c1df7-****-****-****-****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// Column encryption key mode. Valid values:
	//
	// - client_key (configure a randomly generated user key on the client side)
	//
	// - kms_key (use a custom key configured via Alibaba Cloud KMS)
	//
	// Note:
	//
	// Once an instance is configured to use a KMS key, it can no longer use the client-side random key configuration method.
	//
	// example:
	//
	// kms_key
	EncryptionKeyMode *string `json:"EncryptionKeyMode,omitempty" xml:"EncryptionKeyMode,omitempty"`
	// Column encryption status. Valid values:
	//
	// - 1 (Enabled)
	//
	// - 0 (Disabled)
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	EncryptionStatus *string `json:"EncryptionStatus,omitempty" xml:"EncryptionStatus,omitempty"`
	// Whether to rotate the key
	//
	// example:
	//
	// true
	IsRotate             *bool   `json:"IsRotate,omitempty" xml:"IsRotate,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Global Resource Descriptor (GRD) of the role used to specify the exact role. For more information, see RAM Role Overview.
	//
	// Note:
	//
	// This parameter takes effect only when the column encryption key mode is set to kms_key. If not provided, the system uses an internal default value.
	//
	// example:
	//
	// acs:ram::1406926****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// Whether to enable whitelist mode. true indicates that only columns in the whitelist are encrypted; false indicates that all columns are encrypted.
	//
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
