// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKMSConfig interface {
	dara.Model
	String() string
	GoString() string
	SetKmsInstanceId(v string) *KMSConfig
	GetKmsInstanceId() *string
	SetKmsKeyId(v string) *KMSConfig
	GetKmsKeyId() *string
}

type KMSConfig struct {
	// The KMS instance ID
	//
	// example:
	//
	// kst-xxxxxxxx
	KmsInstanceId *string `json:"kmsInstanceId,omitempty" xml:"kmsInstanceId,omitempty"`
	// The KMS key ID
	//
	// example:
	//
	// key-xxxxxxxx
	KmsKeyId *string `json:"kmsKeyId,omitempty" xml:"kmsKeyId,omitempty"`
}

func (s KMSConfig) String() string {
	return dara.Prettify(s)
}

func (s KMSConfig) GoString() string {
	return s.String()
}

func (s *KMSConfig) GetKmsInstanceId() *string {
	return s.KmsInstanceId
}

func (s *KMSConfig) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *KMSConfig) SetKmsInstanceId(v string) *KMSConfig {
	s.KmsInstanceId = &v
	return s
}

func (s *KMSConfig) SetKmsKeyId(v string) *KMSConfig {
	s.KmsKeyId = &v
	return s
}

func (s *KMSConfig) Validate() error {
	return dara.Validate(s)
}
