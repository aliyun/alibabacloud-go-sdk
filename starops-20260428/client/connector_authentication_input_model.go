// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConnectorAuthenticationInput interface {
	dara.Model
	String() string
	GoString() string
	SetBasic(v *ConnectorAuthenticationInputBasic) *ConnectorAuthenticationInput
	GetBasic() *ConnectorAuthenticationInputBasic
	SetBotToken(v *ConnectorAuthenticationInputBotToken) *ConnectorAuthenticationInput
	GetBotToken() *ConnectorAuthenticationInputBotToken
	SetOauth(v *ConnectorAuthenticationInputOauth) *ConnectorAuthenticationInput
	GetOauth() *ConnectorAuthenticationInputOauth
	SetPatToken(v *ConnectorAuthenticationInputPatToken) *ConnectorAuthenticationInput
	GetPatToken() *ConnectorAuthenticationInputPatToken
	SetRole(v *ConnectorAuthenticationInputRole) *ConnectorAuthenticationInput
	GetRole() *ConnectorAuthenticationInputRole
	SetSatellite(v *ConnectorAuthenticationInputSatellite) *ConnectorAuthenticationInput
	GetSatellite() *ConnectorAuthenticationInputSatellite
	SetType(v string) *ConnectorAuthenticationInput
	GetType() *string
}

type ConnectorAuthenticationInput struct {
	// Authenticates by using a username and password.
	Basic *ConnectorAuthenticationInputBasic `json:"basic,omitempty" xml:"basic,omitempty" type:"Struct"`
	// Authenticates by using a bot token and a signing key.
	BotToken *ConnectorAuthenticationInputBotToken `json:"botToken,omitempty" xml:"botToken,omitempty" type:"Struct"`
	// Authenticates by using an OAuth client identity.
	Oauth *ConnectorAuthenticationInputOauth `json:"oauth,omitempty" xml:"oauth,omitempty" type:"Struct"`
	// Authenticates by using a personal access token.
	PatToken *ConnectorAuthenticationInputPatToken `json:"patToken,omitempty" xml:"patToken,omitempty" type:"Struct"`
	// Authenticates by using a RAM role ARN.
	Role *ConnectorAuthenticationInputRole `json:"role,omitempty" xml:"role,omitempty" type:"Struct"`
	// Authenticates by using local credential binding.
	Satellite *ConnectorAuthenticationInputSatellite `json:"satellite,omitempty" xml:"satellite,omitempty" type:"Struct"`
	// Authentication type
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN_RESOURCE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConnectorAuthenticationInput) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationInput) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationInput) GetBasic() *ConnectorAuthenticationInputBasic {
	return s.Basic
}

func (s *ConnectorAuthenticationInput) GetBotToken() *ConnectorAuthenticationInputBotToken {
	return s.BotToken
}

func (s *ConnectorAuthenticationInput) GetOauth() *ConnectorAuthenticationInputOauth {
	return s.Oauth
}

func (s *ConnectorAuthenticationInput) GetPatToken() *ConnectorAuthenticationInputPatToken {
	return s.PatToken
}

func (s *ConnectorAuthenticationInput) GetRole() *ConnectorAuthenticationInputRole {
	return s.Role
}

func (s *ConnectorAuthenticationInput) GetSatellite() *ConnectorAuthenticationInputSatellite {
	return s.Satellite
}

func (s *ConnectorAuthenticationInput) GetType() *string {
	return s.Type
}

func (s *ConnectorAuthenticationInput) SetBasic(v *ConnectorAuthenticationInputBasic) *ConnectorAuthenticationInput {
	s.Basic = v
	return s
}

func (s *ConnectorAuthenticationInput) SetBotToken(v *ConnectorAuthenticationInputBotToken) *ConnectorAuthenticationInput {
	s.BotToken = v
	return s
}

func (s *ConnectorAuthenticationInput) SetOauth(v *ConnectorAuthenticationInputOauth) *ConnectorAuthenticationInput {
	s.Oauth = v
	return s
}

func (s *ConnectorAuthenticationInput) SetPatToken(v *ConnectorAuthenticationInputPatToken) *ConnectorAuthenticationInput {
	s.PatToken = v
	return s
}

func (s *ConnectorAuthenticationInput) SetRole(v *ConnectorAuthenticationInputRole) *ConnectorAuthenticationInput {
	s.Role = v
	return s
}

func (s *ConnectorAuthenticationInput) SetSatellite(v *ConnectorAuthenticationInputSatellite) *ConnectorAuthenticationInput {
	s.Satellite = v
	return s
}

func (s *ConnectorAuthenticationInput) SetType(v string) *ConnectorAuthenticationInput {
	s.Type = &v
	return s
}

func (s *ConnectorAuthenticationInput) Validate() error {
	if s.Basic != nil {
		if err := s.Basic.Validate(); err != nil {
			return err
		}
	}
	if s.BotToken != nil {
		if err := s.BotToken.Validate(); err != nil {
			return err
		}
	}
	if s.Oauth != nil {
		if err := s.Oauth.Validate(); err != nil {
			return err
		}
	}
	if s.PatToken != nil {
		if err := s.PatToken.Validate(); err != nil {
			return err
		}
	}
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	if s.Satellite != nil {
		if err := s.Satellite.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConnectorAuthenticationInputBasic struct {
	// Password
	//
	// This parameter is required.
	//
	// example:
	//
	// example-password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// Username
	//
	// This parameter is required.
	//
	// example:
	//
	// starops
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ConnectorAuthenticationInputBasic) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationInputBasic) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationInputBasic) GetPassword() *string {
	return s.Password
}

func (s *ConnectorAuthenticationInputBasic) GetUsername() *string {
	return s.Username
}

func (s *ConnectorAuthenticationInputBasic) SetPassword(v string) *ConnectorAuthenticationInputBasic {
	s.Password = &v
	return s
}

func (s *ConnectorAuthenticationInputBasic) SetUsername(v string) *ConnectorAuthenticationInputBasic {
	s.Username = &v
	return s
}

func (s *ConnectorAuthenticationInputBasic) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationInputBotToken struct {
	// Bot token
	//
	// This parameter is required.
	//
	// example:
	//
	// example-bot-token
	BotToken *string `json:"botToken,omitempty" xml:"botToken,omitempty"`
	// Signing secret
	//
	// example:
	//
	// example-signing-secret
	SigningSecret *string `json:"signingSecret,omitempty" xml:"signingSecret,omitempty"`
}

func (s ConnectorAuthenticationInputBotToken) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationInputBotToken) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationInputBotToken) GetBotToken() *string {
	return s.BotToken
}

func (s *ConnectorAuthenticationInputBotToken) GetSigningSecret() *string {
	return s.SigningSecret
}

func (s *ConnectorAuthenticationInputBotToken) SetBotToken(v string) *ConnectorAuthenticationInputBotToken {
	s.BotToken = &v
	return s
}

func (s *ConnectorAuthenticationInputBotToken) SetSigningSecret(v string) *ConnectorAuthenticationInputBotToken {
	s.SigningSecret = &v
	return s
}

func (s *ConnectorAuthenticationInputBotToken) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationInputOauth struct {
	// OAuth client ID
	//
	// This parameter is required.
	//
	// example:
	//
	// client-id
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// OAuth client secret
	//
	// This parameter is required.
	//
	// example:
	//
	// example-client-secret
	ClientSecret *string `json:"clientSecret,omitempty" xml:"clientSecret,omitempty"`
}

func (s ConnectorAuthenticationInputOauth) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationInputOauth) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationInputOauth) GetClientId() *string {
	return s.ClientId
}

func (s *ConnectorAuthenticationInputOauth) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *ConnectorAuthenticationInputOauth) SetClientId(v string) *ConnectorAuthenticationInputOauth {
	s.ClientId = &v
	return s
}

func (s *ConnectorAuthenticationInputOauth) SetClientSecret(v string) *ConnectorAuthenticationInputOauth {
	s.ClientSecret = &v
	return s
}

func (s *ConnectorAuthenticationInputOauth) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationInputPatToken struct {
	// The personal access token used to access the target service.
	//
	// This parameter is required.
	//
	// example:
	//
	// example-personal-access-token
	PatToken *string `json:"patToken,omitempty" xml:"patToken,omitempty"`
}

func (s ConnectorAuthenticationInputPatToken) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationInputPatToken) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationInputPatToken) GetPatToken() *string {
	return s.PatToken
}

func (s *ConnectorAuthenticationInputPatToken) SetPatToken(v string) *ConnectorAuthenticationInputPatToken {
	s.PatToken = &v
	return s
}

func (s *ConnectorAuthenticationInputPatToken) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationInputRole struct {
	// Role ARN
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/starops-reader
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
}

func (s ConnectorAuthenticationInputRole) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationInputRole) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationInputRole) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ConnectorAuthenticationInputRole) SetRoleArn(v string) *ConnectorAuthenticationInputRole {
	s.RoleArn = &v
	return s
}

func (s *ConnectorAuthenticationInputRole) Validate() error {
	return dara.Validate(s)
}

type ConnectorAuthenticationInputSatellite struct {
	// Local credential binding name
	//
	// This parameter is required.
	//
	// example:
	//
	// private-gitlab
	BindingName *string `json:"bindingName,omitempty" xml:"bindingName,omitempty"`
}

func (s ConnectorAuthenticationInputSatellite) String() string {
	return dara.Prettify(s)
}

func (s ConnectorAuthenticationInputSatellite) GoString() string {
	return s.String()
}

func (s *ConnectorAuthenticationInputSatellite) GetBindingName() *string {
	return s.BindingName
}

func (s *ConnectorAuthenticationInputSatellite) SetBindingName(v string) *ConnectorAuthenticationInputSatellite {
	s.BindingName = &v
	return s
}

func (s *ConnectorAuthenticationInputSatellite) Validate() error {
	return dara.Validate(s)
}
