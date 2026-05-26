// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEncryptionConfig interface {
  dara.Model
  String() string
  GoString() string
  SetKeyType(v string) *EncryptionConfig
  GetKeyType() *string 
  SetKmsKeyArn(v string) *EncryptionConfig
  GetKmsKeyArn() *string 
}

type EncryptionConfig struct {
  KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
  KmsKeyArn *string `json:"KmsKeyArn,omitempty" xml:"KmsKeyArn,omitempty"`
}

func (s EncryptionConfig) String() string {
  return dara.Prettify(s)
}

func (s EncryptionConfig) GoString() string {
  return s.String()
}

func (s *EncryptionConfig) GetKeyType() *string  {
  return s.KeyType
}

func (s *EncryptionConfig) GetKmsKeyArn() *string  {
  return s.KmsKeyArn
}

func (s *EncryptionConfig) SetKeyType(v string) *EncryptionConfig {
  s.KeyType = &v
  return s
}

func (s *EncryptionConfig) SetKmsKeyArn(v string) *EncryptionConfig {
  s.KmsKeyArn = &v
  return s
}

func (s *EncryptionConfig) Validate() error {
  return dara.Validate(s)
}

