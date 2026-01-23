// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationToken(v string) *GetAuthorizationTokenResponseBody
	GetAuthorizationToken() *string
	SetCode(v string) *GetAuthorizationTokenResponseBody
	GetCode() *string
	SetExpireTime(v int64) *GetAuthorizationTokenResponseBody
	GetExpireTime() *int64
	SetIsSuccess(v bool) *GetAuthorizationTokenResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *GetAuthorizationTokenResponseBody
	GetRequestId() *string
	SetTempUsername(v string) *GetAuthorizationTokenResponseBody
	GetTempUsername() *string
}

type GetAuthorizationTokenResponseBody struct {
	// The password that you use to log on to the registry.
	//
	// example:
	//
	// shaunadadakks:uuczxnjcyeyhdjadkkajsjdjadhyucb
	AuthorizationToken *string `json:"AuthorizationToken,omitempty" xml:"AuthorizationToken,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The timestamp when the temporary token expired. Unit: milliseconds.
	//
	// example:
	//
	// 1571242083000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID
	//
	// example:
	//
	// E069EB86-E6AD-4A98-ADDE-0E993390239A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The username that you use to log on to the registry.
	//
	// example:
	//
	// temp_user_cr
	TempUsername *string `json:"TempUsername,omitempty" xml:"TempUsername,omitempty"`
}

func (s GetAuthorizationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationTokenResponseBody) GetAuthorizationToken() *string {
	return s.AuthorizationToken
}

func (s *GetAuthorizationTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAuthorizationTokenResponseBody) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetAuthorizationTokenResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetAuthorizationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthorizationTokenResponseBody) GetTempUsername() *string {
	return s.TempUsername
}

func (s *GetAuthorizationTokenResponseBody) SetAuthorizationToken(v string) *GetAuthorizationTokenResponseBody {
	s.AuthorizationToken = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetCode(v string) *GetAuthorizationTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetExpireTime(v int64) *GetAuthorizationTokenResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetIsSuccess(v bool) *GetAuthorizationTokenResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetRequestId(v string) *GetAuthorizationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetTempUsername(v string) *GetAuthorizationTokenResponseBody {
	s.TempUsername = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
