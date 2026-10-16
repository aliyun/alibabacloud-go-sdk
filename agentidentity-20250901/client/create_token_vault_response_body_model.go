// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenVaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTokenVaultResponseBody
	GetRequestId() *string
	SetTokenVault(v *CreateTokenVaultResponseBodyTokenVault) *CreateTokenVaultResponseBody
	GetTokenVault() *CreateTokenVaultResponseBodyTokenVault
}

type CreateTokenVaultResponseBody struct {
	// The unique identifier of the request.
	//
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the token vault.
	TokenVault *CreateTokenVaultResponseBodyTokenVault `json:"TokenVault,omitempty" xml:"TokenVault,omitempty" type:"Struct"`
}

func (s CreateTokenVaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenVaultResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenVaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTokenVaultResponseBody) GetTokenVault() *CreateTokenVaultResponseBodyTokenVault {
	return s.TokenVault
}

func (s *CreateTokenVaultResponseBody) SetRequestId(v string) *CreateTokenVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTokenVaultResponseBody) SetTokenVault(v *CreateTokenVaultResponseBodyTokenVault) *CreateTokenVaultResponseBody {
	s.TokenVault = v
	return s
}

func (s *CreateTokenVaultResponseBody) Validate() error {
	if s.TokenVault != nil {
		if err := s.TokenVault.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTokenVaultResponseBodyTokenVault struct {
	// The time when the token vault was created, in UTC ISO 8601 format.
	//
	// example:
	//
	// 2026-05-08T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the token vault.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The encryption configuration of the token vault.
	EncryptionConfig *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig `json:"EncryptionConfig,omitempty" xml:"EncryptionConfig,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the Resource Access Management (RAM) role used to authorize the AgentIdentity service to access other cloud resources such as KMS.
	//
	// example:
	//
	// acs:ram::123456:role/AliyunAgentIdentityVaultRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The ARN of the token vault.
	//
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:tokenvault/default
	TokenVaultArn *string `json:"TokenVaultArn,omitempty" xml:"TokenVaultArn,omitempty"`
	// The name of the token vault. The name must be unique within the current Alibaba Cloud account.
	//
	// example:
	//
	// default
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s CreateTokenVaultResponseBodyTokenVault) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenVaultResponseBodyTokenVault) GoString() string {
	return s.String()
}

func (s *CreateTokenVaultResponseBodyTokenVault) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateTokenVaultResponseBodyTokenVault) GetDescription() *string {
	return s.Description
}

func (s *CreateTokenVaultResponseBodyTokenVault) GetEncryptionConfig() *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig {
	return s.EncryptionConfig
}

func (s *CreateTokenVaultResponseBodyTokenVault) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateTokenVaultResponseBodyTokenVault) GetTokenVaultArn() *string {
	return s.TokenVaultArn
}

func (s *CreateTokenVaultResponseBodyTokenVault) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *CreateTokenVaultResponseBodyTokenVault) SetCreateTime(v string) *CreateTokenVaultResponseBodyTokenVault {
	s.CreateTime = &v
	return s
}

func (s *CreateTokenVaultResponseBodyTokenVault) SetDescription(v string) *CreateTokenVaultResponseBodyTokenVault {
	s.Description = &v
	return s
}

func (s *CreateTokenVaultResponseBodyTokenVault) SetEncryptionConfig(v *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig) *CreateTokenVaultResponseBodyTokenVault {
	s.EncryptionConfig = v
	return s
}

func (s *CreateTokenVaultResponseBodyTokenVault) SetRoleArn(v string) *CreateTokenVaultResponseBodyTokenVault {
	s.RoleArn = &v
	return s
}

func (s *CreateTokenVaultResponseBodyTokenVault) SetTokenVaultArn(v string) *CreateTokenVaultResponseBodyTokenVault {
	s.TokenVaultArn = &v
	return s
}

func (s *CreateTokenVaultResponseBodyTokenVault) SetTokenVaultName(v string) *CreateTokenVaultResponseBodyTokenVault {
	s.TokenVaultName = &v
	return s
}

func (s *CreateTokenVaultResponseBodyTokenVault) Validate() error {
	if s.EncryptionConfig != nil {
		if err := s.EncryptionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTokenVaultResponseBodyTokenVaultEncryptionConfig struct {
	// The type of the encryption key. Valid values:
	//
	// - SERVICE_MANAGED_KEY: service-managed key
	//
	// - CUSTOMER_MANAGED_KEY: customer-managed KMS key
	//
	// example:
	//
	// SERVICE_MANAGED_KEY
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The ARN of the KMS key used when KeyType is set to CUSTOMER_MANAGED_KEY.
	//
	// example:
	//
	// acs:kms:cn-beijing:123456:key/key-bjxxxxxxxx
	KmsKeyArn *string `json:"KmsKeyArn,omitempty" xml:"KmsKeyArn,omitempty"`
}

func (s CreateTokenVaultResponseBodyTokenVaultEncryptionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenVaultResponseBodyTokenVaultEncryptionConfig) GoString() string {
	return s.String()
}

func (s *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig) GetKeyType() *string {
	return s.KeyType
}

func (s *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig) GetKmsKeyArn() *string {
	return s.KmsKeyArn
}

func (s *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig) SetKeyType(v string) *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig {
	s.KeyType = &v
	return s
}

func (s *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig) SetKmsKeyArn(v string) *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig {
	s.KmsKeyArn = &v
	return s
}

func (s *CreateTokenVaultResponseBodyTokenVaultEncryptionConfig) Validate() error {
	return dara.Validate(s)
}
