// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenVaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTokenVaultResponseBody
	GetRequestId() *string
	SetTokenVault(v *GetTokenVaultResponseBodyTokenVault) *GetTokenVaultResponseBody
	GetTokenVault() *GetTokenVaultResponseBodyTokenVault
}

type GetTokenVaultResponseBody struct {
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TokenVault *GetTokenVaultResponseBodyTokenVault `json:"TokenVault,omitempty" xml:"TokenVault,omitempty" type:"Struct"`
}

func (s GetTokenVaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenVaultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenVaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenVaultResponseBody) GetTokenVault() *GetTokenVaultResponseBodyTokenVault {
	return s.TokenVault
}

func (s *GetTokenVaultResponseBody) SetRequestId(v string) *GetTokenVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenVaultResponseBody) SetTokenVault(v *GetTokenVaultResponseBodyTokenVault) *GetTokenVaultResponseBody {
	s.TokenVault = v
	return s
}

func (s *GetTokenVaultResponseBody) Validate() error {
	if s.TokenVault != nil {
		if err := s.TokenVault.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTokenVaultResponseBodyTokenVault struct {
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// example description
	Description      *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptionConfig *GetTokenVaultResponseBodyTokenVaultEncryptionConfig `json:"EncryptionConfig,omitempty" xml:"EncryptionConfig,omitempty" type:"Struct"`
	// example:
	//
	// acs:ram::123456:role/AliyunAgentIdentityVaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:tokenvault/default
	TokenVaultArn *string `json:"TokenVaultArn,omitempty" xml:"TokenVaultArn,omitempty"`
	// example:
	//
	// default
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
	// example:
	//
	// 2026-05-08T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTokenVaultResponseBodyTokenVault) String() string {
	return dara.Prettify(s)
}

func (s GetTokenVaultResponseBodyTokenVault) GoString() string {
	return s.String()
}

func (s *GetTokenVaultResponseBodyTokenVault) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTokenVaultResponseBodyTokenVault) GetDescription() *string {
	return s.Description
}

func (s *GetTokenVaultResponseBodyTokenVault) GetEncryptionConfig() *GetTokenVaultResponseBodyTokenVaultEncryptionConfig {
	return s.EncryptionConfig
}

func (s *GetTokenVaultResponseBodyTokenVault) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetTokenVaultResponseBodyTokenVault) GetTokenVaultArn() *string {
	return s.TokenVaultArn
}

func (s *GetTokenVaultResponseBodyTokenVault) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *GetTokenVaultResponseBodyTokenVault) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetTokenVaultResponseBodyTokenVault) SetCreateTime(v string) *GetTokenVaultResponseBodyTokenVault {
	s.CreateTime = &v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVault) SetDescription(v string) *GetTokenVaultResponseBodyTokenVault {
	s.Description = &v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVault) SetEncryptionConfig(v *GetTokenVaultResponseBodyTokenVaultEncryptionConfig) *GetTokenVaultResponseBodyTokenVault {
	s.EncryptionConfig = v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVault) SetRoleArn(v string) *GetTokenVaultResponseBodyTokenVault {
	s.RoleArn = &v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVault) SetTokenVaultArn(v string) *GetTokenVaultResponseBodyTokenVault {
	s.TokenVaultArn = &v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVault) SetTokenVaultName(v string) *GetTokenVaultResponseBodyTokenVault {
	s.TokenVaultName = &v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVault) SetUpdateTime(v string) *GetTokenVaultResponseBodyTokenVault {
	s.UpdateTime = &v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVault) Validate() error {
	if s.EncryptionConfig != nil {
		if err := s.EncryptionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTokenVaultResponseBodyTokenVaultEncryptionConfig struct {
	// example:
	//
	// SERVICE_MANAGED_KEY
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// example:
	//
	// acs:kms:cn-beijing:123456:key/key-bjxxxxxxxx
	KmsKeyArn *string `json:"KmsKeyArn,omitempty" xml:"KmsKeyArn,omitempty"`
}

func (s GetTokenVaultResponseBodyTokenVaultEncryptionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTokenVaultResponseBodyTokenVaultEncryptionConfig) GoString() string {
	return s.String()
}

func (s *GetTokenVaultResponseBodyTokenVaultEncryptionConfig) GetKeyType() *string {
	return s.KeyType
}

func (s *GetTokenVaultResponseBodyTokenVaultEncryptionConfig) GetKmsKeyArn() *string {
	return s.KmsKeyArn
}

func (s *GetTokenVaultResponseBodyTokenVaultEncryptionConfig) SetKeyType(v string) *GetTokenVaultResponseBodyTokenVaultEncryptionConfig {
	s.KeyType = &v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVaultEncryptionConfig) SetKmsKeyArn(v string) *GetTokenVaultResponseBodyTokenVaultEncryptionConfig {
	s.KmsKeyArn = &v
	return s
}

func (s *GetTokenVaultResponseBodyTokenVaultEncryptionConfig) Validate() error {
	return dara.Validate(s)
}
