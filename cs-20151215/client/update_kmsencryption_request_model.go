// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKMSEncryptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisableEncryption(v bool) *UpdateKMSEncryptionRequest
	GetDisableEncryption() *bool
	SetKmsKeyId(v string) *UpdateKMSEncryptionRequest
	GetKmsKeyId() *string
}

type UpdateKMSEncryptionRequest struct {
	// Specifies whether to disable the Secret encryption at rest feature.
	//
	// 	- `true`: Disables the encryption at rest feature.
	//
	// 	- `false`: Enables the encryption at rest feature.
	//
	// example:
	//
	// false
	DisableEncryption *bool `json:"disable_encryption,omitempty" xml:"disable_encryption,omitempty"`
	// The key ID of the KMS key used by the Secret encryption at rest feature.
	//
	// 	Notice: You cannot use a service key. Use only a master key or a customer master key (CMK). The key type must be `Aliyun_AES_256`, and the key usage must be `ENCRYPT/DECRYPT`.
	//
	//
	//
	//
	// 	Warning: During the procedure of enabling or shutdown of encryption at rest and after the feature is enabled, do not disable or delete the KMS key in the KMS console or through OpenAPI. Otherwise, the cluster API Server becomes not active, which prevents normal retrieval of objects such as Secrets and ServiceAccounts and affects the normal operation of business applications.
	//
	// example:
	//
	// key-abc***
	KmsKeyId *string `json:"kms_key_id,omitempty" xml:"kms_key_id,omitempty"`
}

func (s UpdateKMSEncryptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKMSEncryptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateKMSEncryptionRequest) GetDisableEncryption() *bool {
	return s.DisableEncryption
}

func (s *UpdateKMSEncryptionRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *UpdateKMSEncryptionRequest) SetDisableEncryption(v bool) *UpdateKMSEncryptionRequest {
	s.DisableEncryption = &v
	return s
}

func (s *UpdateKMSEncryptionRequest) SetKmsKeyId(v string) *UpdateKMSEncryptionRequest {
	s.KmsKeyId = &v
	return s
}

func (s *UpdateKMSEncryptionRequest) Validate() error {
	return dara.Validate(s)
}
