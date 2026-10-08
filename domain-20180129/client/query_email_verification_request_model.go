// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmailVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *QueryEmailVerificationRequest
	GetEmail() *string
	SetLang(v string) *QueryEmailVerificationRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryEmailVerificationRequest
	GetUserClientIp() *string
}

type QueryEmailVerificationRequest struct {
	// The Email to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Language of the error message returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default Value is **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// User IP address, which can be set to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryEmailVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *QueryEmailVerificationRequest) GetEmail() *string {
	return s.Email
}

func (s *QueryEmailVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryEmailVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryEmailVerificationRequest) SetEmail(v string) *QueryEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *QueryEmailVerificationRequest) SetLang(v string) *QueryEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *QueryEmailVerificationRequest) SetUserClientIp(v string) *QueryEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryEmailVerificationRequest) Validate() error {
	return dara.Validate(s)
}
