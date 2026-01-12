// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJwtToken(v string) *AuthUserRequest
	GetJwtToken() *string
}

type AuthUserRequest struct {
	// This parameter is required.
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s AuthUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthUserRequest) GoString() string {
	return s.String()
}

func (s *AuthUserRequest) GetJwtToken() *string {
	return s.JwtToken
}

func (s *AuthUserRequest) SetJwtToken(v string) *AuthUserRequest {
	s.JwtToken = &v
	return s
}

func (s *AuthUserRequest) Validate() error {
	return dara.Validate(s)
}
