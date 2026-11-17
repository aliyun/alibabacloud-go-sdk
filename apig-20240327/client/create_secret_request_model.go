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
	// The description of the secret.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The gateway type.
	//
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The KMS key configuration. This parameter is required when secretSource is set to KMS. If this parameter is omitted, the API returns InvalidParameter.IsEmpty(400) with the message "Parameter kmsConfig should not be empty.". This parameter contains two sub-fields: kmsInstanceId (the KMS instance ID) and kmsKeyId (the KMS key ID).
	KmsConfig *KMSConfig `json:"kmsConfig,omitempty" xml:"kmsConfig,omitempty"`
	// The name of the secret. The name can contain letters, digits, and underscores (_) and can be up to 64 characters in length.
	//
	// example:
	//
	// my_secret
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The KMS secret value. This parameter is required. If this parameter is omitted, the API returns InvalidParameter.IsEmpty(400). Because secretSource only supports KMS, you must always provide this parameter in actual calls.
	//
	// example:
	//
	// apikey-123456xxxxxxxx
	SecretData *string `json:"secretData,omitempty" xml:"secretData,omitempty"`
	// The source of the secret. Currently, only KMS is supported. If you specify other values such as MANUAL, the API returns InvalidParameter.WithValue(400). This parameter is required. If this parameter is not provided, the API also returns InvalidParameter.WithValue(400).
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
