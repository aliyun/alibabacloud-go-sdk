// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateAuthenticationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *ValidateAuthenticationTokenRequest
	GetToken() *string
	SetTokenTypeHint(v string) *ValidateAuthenticationTokenRequest
	GetTokenTypeHint() *string
}

type ValidateAuthenticationTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eyJhbGciOixxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// -
	TokenTypeHint *string `json:"token_type_hint,omitempty" xml:"token_type_hint,omitempty"`
}

func (s ValidateAuthenticationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateAuthenticationTokenRequest) GoString() string {
	return s.String()
}

func (s *ValidateAuthenticationTokenRequest) GetToken() *string {
	return s.Token
}

func (s *ValidateAuthenticationTokenRequest) GetTokenTypeHint() *string {
	return s.TokenTypeHint
}

func (s *ValidateAuthenticationTokenRequest) SetToken(v string) *ValidateAuthenticationTokenRequest {
	s.Token = &v
	return s
}

func (s *ValidateAuthenticationTokenRequest) SetTokenTypeHint(v string) *ValidateAuthenticationTokenRequest {
	s.TokenTypeHint = &v
	return s
}

func (s *ValidateAuthenticationTokenRequest) Validate() error {
	return dara.Validate(s)
}
