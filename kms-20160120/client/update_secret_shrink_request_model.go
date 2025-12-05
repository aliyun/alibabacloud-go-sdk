// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtendedConfig(v *UpdateSecretShrinkRequestExtendedConfig) *UpdateSecretShrinkRequest
	GetExtendedConfig() *UpdateSecretShrinkRequestExtendedConfig
	SetDescription(v string) *UpdateSecretShrinkRequest
	GetDescription() *string
	SetSecretName(v string) *UpdateSecretShrinkRequest
	GetSecretName() *string
}

type UpdateSecretShrinkRequest struct {
	ExtendedConfig *UpdateSecretShrinkRequestExtendedConfig `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty" type:"Struct"`
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

func (s UpdateSecretShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretShrinkRequest) GetExtendedConfig() *UpdateSecretShrinkRequestExtendedConfig {
	return s.ExtendedConfig
}

func (s *UpdateSecretShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSecretShrinkRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *UpdateSecretShrinkRequest) SetExtendedConfig(v *UpdateSecretShrinkRequestExtendedConfig) *UpdateSecretShrinkRequest {
	s.ExtendedConfig = v
	return s
}

func (s *UpdateSecretShrinkRequest) SetDescription(v string) *UpdateSecretShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretShrinkRequest) SetSecretName(v string) *UpdateSecretShrinkRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretShrinkRequest) Validate() error {
	if s.ExtendedConfig != nil {
		if err := s.ExtendedConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSecretShrinkRequestExtendedConfig struct {
	// The custom data in the extended configuration of the secret.
	//
	// > 	- If this parameter is specified, the existing extended configuration of the secret is updated.
	//
	// > 	- This parameter is unavailable for generic secrets.
	//
	// example:
	//
	// {"DBName":"app1","Port":"3306"}
	CustomData *string `json:"CustomData,omitempty" xml:"CustomData,omitempty"`
}

func (s UpdateSecretShrinkRequestExtendedConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretShrinkRequestExtendedConfig) GoString() string {
	return s.String()
}

func (s *UpdateSecretShrinkRequestExtendedConfig) GetCustomData() *string {
	return s.CustomData
}

func (s *UpdateSecretShrinkRequestExtendedConfig) SetCustomData(v string) *UpdateSecretShrinkRequestExtendedConfig {
	s.CustomData = &v
	return s
}

func (s *UpdateSecretShrinkRequestExtendedConfig) Validate() error {
	return dara.Validate(s)
}
