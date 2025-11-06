// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEmailVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *SubmitEmailVerificationRequest
	GetEmail() *string
	SetLang(v string) *SubmitEmailVerificationRequest
	GetLang() *string
	SetSendIfExist(v bool) *SubmitEmailVerificationRequest
	GetSendIfExist() *bool
	SetUserClientIp(v string) *SubmitEmailVerificationRequest
	GetUserClientIp() *string
}

type SubmitEmailVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// false
	SendIfExist *bool `json:"SendIfExist,omitempty" xml:"SendIfExist,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SubmitEmailVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationRequest) GetEmail() *string {
	return s.Email
}

func (s *SubmitEmailVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *SubmitEmailVerificationRequest) GetSendIfExist() *bool {
	return s.SendIfExist
}

func (s *SubmitEmailVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SubmitEmailVerificationRequest) SetEmail(v string) *SubmitEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *SubmitEmailVerificationRequest) SetLang(v string) *SubmitEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *SubmitEmailVerificationRequest) SetSendIfExist(v bool) *SubmitEmailVerificationRequest {
	s.SendIfExist = &v
	return s
}

func (s *SubmitEmailVerificationRequest) SetUserClientIp(v string) *SubmitEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *SubmitEmailVerificationRequest) Validate() error {
	return dara.Validate(s)
}
