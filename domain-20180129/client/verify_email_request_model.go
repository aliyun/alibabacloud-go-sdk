// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *VerifyEmailRequest
	GetLang() *string
	SetToken(v string) *VerifyEmailRequest
	GetToken() *string
	SetUserClientIp(v string) *VerifyEmailRequest
	GetUserClientIp() *string
}

type VerifyEmailRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0b32247496409441e9e179ea7c2e0****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s VerifyEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyEmailRequest) GoString() string {
	return s.String()
}

func (s *VerifyEmailRequest) GetLang() *string {
	return s.Lang
}

func (s *VerifyEmailRequest) GetToken() *string {
	return s.Token
}

func (s *VerifyEmailRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *VerifyEmailRequest) SetLang(v string) *VerifyEmailRequest {
	s.Lang = &v
	return s
}

func (s *VerifyEmailRequest) SetToken(v string) *VerifyEmailRequest {
	s.Token = &v
	return s
}

func (s *VerifyEmailRequest) SetUserClientIp(v string) *VerifyEmailRequest {
	s.UserClientIp = &v
	return s
}

func (s *VerifyEmailRequest) Validate() error {
	return dara.Validate(s)
}
