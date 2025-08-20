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
	// The temporary password returned after you call this API operation is a Security Token Service (STS) token whose validity period is 1 hour. Take note of the following items when you log on to Container Registry instances by using an STS token:
	//
	// 	- If the STS token belongs to an Alibaba Cloud account, you can use the STS token to log on to all Container Registry instances that belong to the Alibaba Cloud account.
	//
	// 	- If the STS token belongs to a Resource Access Management (RAM) user, you can use the STS token to log on to all Container Registry instances that belong to the RAM user.
	//
	// 	- You can use an STS token to access only Container Registry instances to which the STS token is scoped.
	//
	// example:
	//
	// shaunadadakks:uuczxnjcyeyhdjadkkajsjdjadhyucb
	AuthorizationToken *string `json:"AuthorizationToken,omitempty" xml:"AuthorizationToken,omitempty"`
	// Indicates whether the API call is successful.
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value.
	//
	// example:
	//
	// 1571242083000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The username that is used to log on to the Container Registry instance.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The timestamp when the temporary password expires. Unit: milliseconds.
	//
	// example:
	//
	// E069EB86-E6AD-4A98-ADDE-0E993390239A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The password that is used to log on to the Container Registry instance.
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
