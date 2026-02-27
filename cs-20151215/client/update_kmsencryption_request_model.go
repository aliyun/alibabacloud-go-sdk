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
	// example:
	//
	// false
	DisableEncryption *bool `json:"disable_encryption,omitempty" xml:"disable_encryption,omitempty"`
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
