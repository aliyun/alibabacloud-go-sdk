// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorization interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v *AuthorizationParameters) *Authorization
	GetParameters() *AuthorizationParameters
	SetType(v string) *Authorization
	GetType() *string
}

type Authorization struct {
	// This parameter is required.
	Parameters *AuthorizationParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// APIKey
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Authorization) String() string {
	return dara.Prettify(s)
}

func (s Authorization) GoString() string {
	return s.String()
}

func (s *Authorization) GetParameters() *AuthorizationParameters {
	return s.Parameters
}

func (s *Authorization) GetType() *string {
	return s.Type
}

func (s *Authorization) SetParameters(v *AuthorizationParameters) *Authorization {
	s.Parameters = v
	return s
}

func (s *Authorization) SetType(v string) *Authorization {
	s.Type = &v
	return s
}

func (s *Authorization) Validate() error {
	return dara.Validate(s)
}

type AuthorizationParameters struct {
	ApiKeyParameter *APIKeyAuthParameter `json:"apiKeyParameter,omitempty" xml:"apiKeyParameter,omitempty"`
}

func (s AuthorizationParameters) String() string {
	return dara.Prettify(s)
}

func (s AuthorizationParameters) GoString() string {
	return s.String()
}

func (s *AuthorizationParameters) GetApiKeyParameter() *APIKeyAuthParameter {
	return s.ApiKeyParameter
}

func (s *AuthorizationParameters) SetApiKeyParameter(v *APIKeyAuthParameter) *AuthorizationParameters {
	s.ApiKeyParameter = v
	return s
}

func (s *AuthorizationParameters) Validate() error {
	return dara.Validate(s)
}
