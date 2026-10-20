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
	// Specifies whether to disable the encryption at rest feature for Secrets.
	//
	// 	- `true`: Disables the encryption at rest feature.
	//
	// 	- `false`: Enables the encryption at rest feature.
	//
	// example:
	//
	// false
	DisableEncryption *bool `json:"disable_encryption,omitempty" xml:"disable_encryption,omitempty"`
	// The KMS key ID used for the encryption at rest feature for Secrets.
	//
	// 	Notice: You cannot use a service key. You must use a master key or a customer master key (CMK). The key type must be `Aliyun_AES_256`, and the key usage must be `ENCRYPT/DECRYPT`.
	//
	//
	//
	//
	// 	Warning: Do not use the KMS console or OpenAPI to disable or delete this KMS key during the enable or shutdown procedure for encryption at rest, or after the feature is successfully enabled. Otherwise, the cluster API server becomes inactive, which prevents you from retrieving objects such as Secrets and ServiceAccounts and affects the normal operation of your applications.
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
