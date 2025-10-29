// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJwtIdentityConfig interface {
	dara.Model
	String() string
	GoString() string
	SetJwks(v string) *JwtIdentityConfig
	GetJwks() *string
	SetJwtPayloadConfig(v *JwtIdentityConfigJwtPayloadConfig) *JwtIdentityConfig
	GetJwtPayloadConfig() *JwtIdentityConfigJwtPayloadConfig
	SetJwtTokenConfig(v *JwtIdentityConfigJwtTokenConfig) *JwtIdentityConfig
	GetJwtTokenConfig() *JwtIdentityConfigJwtTokenConfig
	SetSecretType(v string) *JwtIdentityConfig
	GetSecretType() *string
	SetType(v string) *JwtIdentityConfig
	GetType() *string
}

type JwtIdentityConfig struct {
	Jwks             *string                            `json:"jwks,omitempty" xml:"jwks,omitempty"`
	JwtPayloadConfig *JwtIdentityConfigJwtPayloadConfig `json:"jwtPayloadConfig,omitempty" xml:"jwtPayloadConfig,omitempty" type:"Struct"`
	JwtTokenConfig   *JwtIdentityConfigJwtTokenConfig   `json:"jwtTokenConfig,omitempty" xml:"jwtTokenConfig,omitempty" type:"Struct"`
	SecretType       *string                            `json:"secretType,omitempty" xml:"secretType,omitempty"`
	Type             *string                            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s JwtIdentityConfig) String() string {
	return dara.Prettify(s)
}

func (s JwtIdentityConfig) GoString() string {
	return s.String()
}

func (s *JwtIdentityConfig) GetJwks() *string {
	return s.Jwks
}

func (s *JwtIdentityConfig) GetJwtPayloadConfig() *JwtIdentityConfigJwtPayloadConfig {
	return s.JwtPayloadConfig
}

func (s *JwtIdentityConfig) GetJwtTokenConfig() *JwtIdentityConfigJwtTokenConfig {
	return s.JwtTokenConfig
}

func (s *JwtIdentityConfig) GetSecretType() *string {
	return s.SecretType
}

func (s *JwtIdentityConfig) GetType() *string {
	return s.Type
}

func (s *JwtIdentityConfig) SetJwks(v string) *JwtIdentityConfig {
	s.Jwks = &v
	return s
}

func (s *JwtIdentityConfig) SetJwtPayloadConfig(v *JwtIdentityConfigJwtPayloadConfig) *JwtIdentityConfig {
	s.JwtPayloadConfig = v
	return s
}

func (s *JwtIdentityConfig) SetJwtTokenConfig(v *JwtIdentityConfigJwtTokenConfig) *JwtIdentityConfig {
	s.JwtTokenConfig = v
	return s
}

func (s *JwtIdentityConfig) SetSecretType(v string) *JwtIdentityConfig {
	s.SecretType = &v
	return s
}

func (s *JwtIdentityConfig) SetType(v string) *JwtIdentityConfig {
	s.Type = &v
	return s
}

func (s *JwtIdentityConfig) Validate() error {
	if s.JwtPayloadConfig != nil {
		if err := s.JwtPayloadConfig.Validate(); err != nil {
			return err
		}
	}
	if s.JwtTokenConfig != nil {
		if err := s.JwtTokenConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type JwtIdentityConfigJwtPayloadConfig struct {
	PayloadKeyName  *string `json:"payloadKeyName,omitempty" xml:"payloadKeyName,omitempty"`
	PayloadKeyValue *string `json:"payloadKeyValue,omitempty" xml:"payloadKeyValue,omitempty"`
}

func (s JwtIdentityConfigJwtPayloadConfig) String() string {
	return dara.Prettify(s)
}

func (s JwtIdentityConfigJwtPayloadConfig) GoString() string {
	return s.String()
}

func (s *JwtIdentityConfigJwtPayloadConfig) GetPayloadKeyName() *string {
	return s.PayloadKeyName
}

func (s *JwtIdentityConfigJwtPayloadConfig) GetPayloadKeyValue() *string {
	return s.PayloadKeyValue
}

func (s *JwtIdentityConfigJwtPayloadConfig) SetPayloadKeyName(v string) *JwtIdentityConfigJwtPayloadConfig {
	s.PayloadKeyName = &v
	return s
}

func (s *JwtIdentityConfigJwtPayloadConfig) SetPayloadKeyValue(v string) *JwtIdentityConfigJwtPayloadConfig {
	s.PayloadKeyValue = &v
	return s
}

func (s *JwtIdentityConfigJwtPayloadConfig) Validate() error {
	return dara.Validate(s)
}

type JwtIdentityConfigJwtTokenConfig struct {
	Key      *string `json:"key,omitempty" xml:"key,omitempty"`
	Pass     *bool   `json:"pass,omitempty" xml:"pass,omitempty"`
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	Prefix   *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s JwtIdentityConfigJwtTokenConfig) String() string {
	return dara.Prettify(s)
}

func (s JwtIdentityConfigJwtTokenConfig) GoString() string {
	return s.String()
}

func (s *JwtIdentityConfigJwtTokenConfig) GetKey() *string {
	return s.Key
}

func (s *JwtIdentityConfigJwtTokenConfig) GetPass() *bool {
	return s.Pass
}

func (s *JwtIdentityConfigJwtTokenConfig) GetPosition() *string {
	return s.Position
}

func (s *JwtIdentityConfigJwtTokenConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *JwtIdentityConfigJwtTokenConfig) SetKey(v string) *JwtIdentityConfigJwtTokenConfig {
	s.Key = &v
	return s
}

func (s *JwtIdentityConfigJwtTokenConfig) SetPass(v bool) *JwtIdentityConfigJwtTokenConfig {
	s.Pass = &v
	return s
}

func (s *JwtIdentityConfigJwtTokenConfig) SetPosition(v string) *JwtIdentityConfigJwtTokenConfig {
	s.Position = &v
	return s
}

func (s *JwtIdentityConfigJwtTokenConfig) SetPrefix(v string) *JwtIdentityConfigJwtTokenConfig {
	s.Prefix = &v
	return s
}

func (s *JwtIdentityConfigJwtTokenConfig) Validate() error {
	return dara.Validate(s)
}
