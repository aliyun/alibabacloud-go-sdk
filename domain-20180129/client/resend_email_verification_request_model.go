// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendEmailVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *ResendEmailVerificationRequest
	GetEmail() *string
	SetLang(v string) *ResendEmailVerificationRequest
	GetLang() *string
	SetUserClientIp(v string) *ResendEmailVerificationRequest
	GetUserClientIp() *string
}

type ResendEmailVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test1@aliyun.com,test2@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ResendEmailVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s ResendEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationRequest) GetEmail() *string {
	return s.Email
}

func (s *ResendEmailVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *ResendEmailVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ResendEmailVerificationRequest) SetEmail(v string) *ResendEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *ResendEmailVerificationRequest) SetLang(v string) *ResendEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *ResendEmailVerificationRequest) SetUserClientIp(v string) *ResendEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *ResendEmailVerificationRequest) Validate() error {
	return dara.Validate(s)
}
