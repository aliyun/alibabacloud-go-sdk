// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEncryptionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SetEncryptionConfigRequest
	GetClientToken() *string
	SetKmsKeyId(v string) *SetEncryptionConfigRequest
	GetKmsKeyId() *string
	SetKmsRegionId(v string) *SetEncryptionConfigRequest
	GetKmsRegionId() *string
}

type SetEncryptionConfigRequest struct {
	// The idempotence token. Format: [0-9a-zA-Z-]{1,64}. Use a UUID.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The ID of the KMS key used for encryption.
	//
	// example:
	//
	// 09d0641c-e96c-495a-925e-9b50xxxxxxxx
	KmsKeyId *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
	// The region ID of the KMS key.
	//
	// example:
	//
	// cn-beijing
	KmsRegionId *string `json:"kmsRegionId,omitempty" xml:"kmsRegionId,omitempty"`
}

func (s SetEncryptionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetEncryptionConfigRequest) GoString() string {
	return s.String()
}

func (s *SetEncryptionConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetEncryptionConfigRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *SetEncryptionConfigRequest) GetKmsRegionId() *string {
	return s.KmsRegionId
}

func (s *SetEncryptionConfigRequest) SetClientToken(v string) *SetEncryptionConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *SetEncryptionConfigRequest) SetKmsKeyId(v string) *SetEncryptionConfigRequest {
	s.KmsKeyId = &v
	return s
}

func (s *SetEncryptionConfigRequest) SetKmsRegionId(v string) *SetEncryptionConfigRequest {
	s.KmsRegionId = &v
	return s
}

func (s *SetEncryptionConfigRequest) Validate() error {
	return dara.Validate(s)
}
