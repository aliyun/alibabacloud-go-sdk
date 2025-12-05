// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendedConfig(v *UpdateSecretRequestExtendedConfig) *UpdateSecretRequest
	GetExtendedConfig() *UpdateSecretRequestExtendedConfig
	SetDescription(v string) *UpdateSecretRequest
	GetDescription() *string
	SetSecretName(v string) *UpdateSecretRequest
	GetSecretName() *string
}

type UpdateSecretRequest struct {
	ExtendedConfig *UpdateSecretRequestExtendedConfig `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty" type:"Struct"`
	// The description of the secret.
	//
	// example:
	//
	// datainfo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretRequest) GetExtendedConfig() *UpdateSecretRequestExtendedConfig {
	return s.ExtendedConfig
}

func (s *UpdateSecretRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *UpdateSecretRequest) SetExtendedConfig(v *UpdateSecretRequestExtendedConfig) *UpdateSecretRequest {
	s.ExtendedConfig = v
	return s
}

func (s *UpdateSecretRequest) SetDescription(v string) *UpdateSecretRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretRequest) SetSecretName(v string) *UpdateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretRequest) Validate() error {
	if s.ExtendedConfig != nil {
		if err := s.ExtendedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSecretRequestExtendedConfig struct {
	// The custom data in the extended configuration of the secret.
	//
	// > 	- If this parameter is specified, the existing extended configuration of the secret is updated.
	//
	// > 	- This parameter is unavailable for generic secrets.
	//
	// example:
	//
	// {"DBName":"app1","Port":"3306"}
	CustomData map[string]interface{} `json:"CustomData,omitempty" xml:"CustomData,omitempty"`
}

func (s UpdateSecretRequestExtendedConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretRequestExtendedConfig) GoString() string {
	return s.String()
}

func (s *UpdateSecretRequestExtendedConfig) GetCustomData() map[string]interface{} {
	return s.CustomData
}

func (s *UpdateSecretRequestExtendedConfig) SetCustomData(v map[string]interface{}) *UpdateSecretRequestExtendedConfig {
	s.CustomData = v
	return s
}

func (s *UpdateSecretRequestExtendedConfig) Validate() error {
	return dara.Validate(s)
}
