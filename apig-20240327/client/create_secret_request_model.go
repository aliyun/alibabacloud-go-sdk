// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSecretRequest
	GetDescription() *string
	SetGatewayType(v string) *CreateSecretRequest
	GetGatewayType() *string
	SetKmsConfig(v *KMSConfig) *CreateSecretRequest
	GetKmsConfig() *KMSConfig
	SetName(v string) *CreateSecretRequest
	GetName() *string
	SetSecretData(v string) *CreateSecretRequest
	GetSecretData() *string
	SetSecretSource(v string) *CreateSecretRequest
	GetSecretSource() *string
}

type CreateSecretRequest struct {
	// The description of the key.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the gateway.
	//
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The key configuration information of KMS.
	KmsConfig *KMSConfig `json:"kmsConfig,omitempty" xml:"kmsConfig,omitempty"`
	// The key name. It can be up to 64 characters in length and can contain letters, digits, and underscores (_).
	//
	// example:
	//
	// my_secret
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the KMS credential.
	//
	// example:
	//
	// apikey-123456xxxxxxxx
	SecretData *string `json:"secretData,omitempty" xml:"secretData,omitempty"`
	// The source of the key.
	//
	// example:
	//
	// KMS
	SecretSource *string `json:"secretSource,omitempty" xml:"secretSource,omitempty"`
}

func (s CreateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecretRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateSecretRequest) GetKmsConfig() *KMSConfig {
	return s.KmsConfig
}

func (s *CreateSecretRequest) GetName() *string {
	return s.Name
}

func (s *CreateSecretRequest) GetSecretData() *string {
	return s.SecretData
}

func (s *CreateSecretRequest) GetSecretSource() *string {
	return s.SecretSource
}

func (s *CreateSecretRequest) SetDescription(v string) *CreateSecretRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretRequest) SetGatewayType(v string) *CreateSecretRequest {
	s.GatewayType = &v
	return s
}

func (s *CreateSecretRequest) SetKmsConfig(v *KMSConfig) *CreateSecretRequest {
	s.KmsConfig = v
	return s
}

func (s *CreateSecretRequest) SetName(v string) *CreateSecretRequest {
	s.Name = &v
	return s
}

func (s *CreateSecretRequest) SetSecretData(v string) *CreateSecretRequest {
	s.SecretData = &v
	return s
}

func (s *CreateSecretRequest) SetSecretSource(v string) *CreateSecretRequest {
	s.SecretSource = &v
	return s
}

func (s *CreateSecretRequest) Validate() error {
	if s.KmsConfig != nil {
		if err := s.KmsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
