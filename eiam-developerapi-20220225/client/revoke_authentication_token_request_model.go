// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAuthenticationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *RevokeAuthenticationTokenRequest
	GetToken() *string
	SetTokenTypeHint(v string) *RevokeAuthenticationTokenRequest
	GetTokenTypeHint() *string
}

type RevokeAuthenticationTokenRequest struct {
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

func (s RevokeAuthenticationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeAuthenticationTokenRequest) GoString() string {
	return s.String()
}

func (s *RevokeAuthenticationTokenRequest) GetToken() *string {
	return s.Token
}

func (s *RevokeAuthenticationTokenRequest) GetTokenTypeHint() *string {
	return s.TokenTypeHint
}

func (s *RevokeAuthenticationTokenRequest) SetToken(v string) *RevokeAuthenticationTokenRequest {
	s.Token = &v
	return s
}

func (s *RevokeAuthenticationTokenRequest) SetTokenTypeHint(v string) *RevokeAuthenticationTokenRequest {
	s.TokenTypeHint = &v
	return s
}

func (s *RevokeAuthenticationTokenRequest) Validate() error {
	return dara.Validate(s)
}
