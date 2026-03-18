// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReinstateAuthenticationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetToken(v string) *ReinstateAuthenticationTokenRequest
	GetToken() *string
	SetTokenTypeHint(v string) *ReinstateAuthenticationTokenRequest
	GetTokenTypeHint() *string
}

type ReinstateAuthenticationTokenRequest struct {
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

func (s ReinstateAuthenticationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ReinstateAuthenticationTokenRequest) GoString() string {
	return s.String()
}

func (s *ReinstateAuthenticationTokenRequest) GetToken() *string {
	return s.Token
}

func (s *ReinstateAuthenticationTokenRequest) GetTokenTypeHint() *string {
	return s.TokenTypeHint
}

func (s *ReinstateAuthenticationTokenRequest) SetToken(v string) *ReinstateAuthenticationTokenRequest {
	s.Token = &v
	return s
}

func (s *ReinstateAuthenticationTokenRequest) SetTokenTypeHint(v string) *ReinstateAuthenticationTokenRequest {
	s.TokenTypeHint = &v
	return s
}

func (s *ReinstateAuthenticationTokenRequest) Validate() error {
	return dara.Validate(s)
}
