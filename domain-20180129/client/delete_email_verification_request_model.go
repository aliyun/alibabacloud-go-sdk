// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmailVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *DeleteEmailVerificationRequest
	GetEmail() *string
	SetLang(v string) *DeleteEmailVerificationRequest
	GetLang() *string
	SetUserClientIp(v string) *DeleteEmailVerificationRequest
	GetUserClientIp() *string
}

type DeleteEmailVerificationRequest struct {
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

func (s DeleteEmailVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationRequest) GetEmail() *string {
	return s.Email
}

func (s *DeleteEmailVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteEmailVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteEmailVerificationRequest) SetEmail(v string) *DeleteEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *DeleteEmailVerificationRequest) SetLang(v string) *DeleteEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *DeleteEmailVerificationRequest) SetUserClientIp(v string) *DeleteEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteEmailVerificationRequest) Validate() error {
	return dara.Validate(s)
}
