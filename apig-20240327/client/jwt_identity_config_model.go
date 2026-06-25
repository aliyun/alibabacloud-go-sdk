// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJwtIdentityConfig interface {
	dara.Model
	String() string
	GoString() string
	SetClaimsToHeadersConfigs(v []*JwtIdentityConfigClaimsToHeadersConfigs) *JwtIdentityConfig
	GetClaimsToHeadersConfigs() []*JwtIdentityConfigClaimsToHeadersConfigs
	SetJwks(v string) *JwtIdentityConfig
	GetJwks() *string
	SetJwtPayloadConfig(v *JwtIdentityConfigJwtPayloadConfig) *JwtIdentityConfig
	GetJwtPayloadConfig() *JwtIdentityConfigJwtPayloadConfig
	SetJwtTokenConfig(v *JwtIdentityConfigJwtTokenConfig) *JwtIdentityConfig
	GetJwtTokenConfig() *JwtIdentityConfigJwtTokenConfig
	SetRemoteJwks(v string) *JwtIdentityConfig
	GetRemoteJwks() *string
	SetSecretType(v string) *JwtIdentityConfig
	GetSecretType() *string
	SetType(v string) *JwtIdentityConfig
	GetType() *string
}

type JwtIdentityConfig struct {
	// The claims-to-headers configurations.
	ClaimsToHeadersConfigs []*JwtIdentityConfigClaimsToHeadersConfigs `json:"claimsToHeadersConfigs,omitempty" xml:"claimsToHeadersConfigs,omitempty" type:"Repeated"`
	// The JWKS configuration.
	//
	// example:
	//
	// xxxx
	Jwks *string `json:"jwks,omitempty" xml:"jwks,omitempty"`
	// The JWT payload configuration.
	JwtPayloadConfig *JwtIdentityConfigJwtPayloadConfig `json:"jwtPayloadConfig,omitempty" xml:"jwtPayloadConfig,omitempty" type:"Struct"`
	// The JWT token configuration.
	JwtTokenConfig *JwtIdentityConfigJwtTokenConfig `json:"jwtTokenConfig,omitempty" xml:"jwtTokenConfig,omitempty" type:"Struct"`
	// The remote JWKS.
	RemoteJwks *string `json:"remoteJwks,omitempty" xml:"remoteJwks,omitempty"`
	// The secret type.
	//
	// example:
	//
	// Symmetry
	SecretType *string `json:"secretType,omitempty" xml:"secretType,omitempty"`
	// The type of authentication configuration.
	//
	// example:
	//
	// Jwt
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s JwtIdentityConfig) String() string {
	return dara.Prettify(s)
}

func (s JwtIdentityConfig) GoString() string {
	return s.String()
}

func (s *JwtIdentityConfig) GetClaimsToHeadersConfigs() []*JwtIdentityConfigClaimsToHeadersConfigs {
	return s.ClaimsToHeadersConfigs
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

func (s *JwtIdentityConfig) GetRemoteJwks() *string {
	return s.RemoteJwks
}

func (s *JwtIdentityConfig) GetSecretType() *string {
	return s.SecretType
}

func (s *JwtIdentityConfig) GetType() *string {
	return s.Type
}

func (s *JwtIdentityConfig) SetClaimsToHeadersConfigs(v []*JwtIdentityConfigClaimsToHeadersConfigs) *JwtIdentityConfig {
	s.ClaimsToHeadersConfigs = v
	return s
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

func (s *JwtIdentityConfig) SetRemoteJwks(v string) *JwtIdentityConfig {
	s.RemoteJwks = &v
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
	if s.ClaimsToHeadersConfigs != nil {
		for _, item := range s.ClaimsToHeadersConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type JwtIdentityConfigClaimsToHeadersConfigs struct {
	// The claim.
	Claim *string `json:"claim,omitempty" xml:"claim,omitempty"`
	// The header.
	Header *string `json:"header,omitempty" xml:"header,omitempty"`
	// The override.
	Override *bool `json:"override,omitempty" xml:"override,omitempty"`
}

func (s JwtIdentityConfigClaimsToHeadersConfigs) String() string {
	return dara.Prettify(s)
}

func (s JwtIdentityConfigClaimsToHeadersConfigs) GoString() string {
	return s.String()
}

func (s *JwtIdentityConfigClaimsToHeadersConfigs) GetClaim() *string {
	return s.Claim
}

func (s *JwtIdentityConfigClaimsToHeadersConfigs) GetHeader() *string {
	return s.Header
}

func (s *JwtIdentityConfigClaimsToHeadersConfigs) GetOverride() *bool {
	return s.Override
}

func (s *JwtIdentityConfigClaimsToHeadersConfigs) SetClaim(v string) *JwtIdentityConfigClaimsToHeadersConfigs {
	s.Claim = &v
	return s
}

func (s *JwtIdentityConfigClaimsToHeadersConfigs) SetHeader(v string) *JwtIdentityConfigClaimsToHeadersConfigs {
	s.Header = &v
	return s
}

func (s *JwtIdentityConfigClaimsToHeadersConfigs) SetOverride(v bool) *JwtIdentityConfigClaimsToHeadersConfigs {
	s.Override = &v
	return s
}

func (s *JwtIdentityConfigClaimsToHeadersConfigs) Validate() error {
	return dara.Validate(s)
}

type JwtIdentityConfigJwtPayloadConfig struct {
	// The JWT payload key configuration.
	//
	// example:
	//
	// uid
	PayloadKeyName *string `json:"payloadKeyName,omitempty" xml:"payloadKeyName,omitempty"`
	// The JWT payload value configuration.
	//
	// example:
	//
	// 2222
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
	// The JWT key configuration.
	//
	// example:
	//
	// Authorization
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Specifies whether to pass through.
	//
	// example:
	//
	// true
	Pass *bool `json:"pass,omitempty" xml:"pass,omitempty"`
	// The storage location of the JWT.
	//
	// example:
	//
	// HEADER
	Position *string `json:"position,omitempty" xml:"position,omitempty"`
	// The prefix configuration.
	//
	// example:
	//
	// test
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
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
