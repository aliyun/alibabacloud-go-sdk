// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationCodeForBindSecureMobilePhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpirationDate(v string) *SendVerificationCodeForBindSecureMobilePhoneResponseBody
	GetExpirationDate() *string
	SetRequestId(v string) *SendVerificationCodeForBindSecureMobilePhoneResponseBody
	GetRequestId() *string
}

type SendVerificationCodeForBindSecureMobilePhoneResponseBody struct {
	// The time when the verification code expires.
	//
	// example:
	//
	// 2021-12-17T07:38:41.747Z
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DCD43660-75DD-5D15-B342-1B83FCA5B913
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationCodeForBindSecureMobilePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) SetExpirationDate(v string) *SendVerificationCodeForBindSecureMobilePhoneResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) SetRequestId(v string) *SendVerificationCodeForBindSecureMobilePhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendVerificationCodeForBindSecureMobilePhoneResponseBody) Validate() error {
	return dara.Validate(s)
}
