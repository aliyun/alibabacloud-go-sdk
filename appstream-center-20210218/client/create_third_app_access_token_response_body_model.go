// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThirdAppAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateThirdAppAccessTokenResponseBody
	GetAccessToken() *string
	SetExpiresAt(v int64) *CreateThirdAppAccessTokenResponseBody
	GetExpiresAt() *int64
	SetRequestId(v string) *CreateThirdAppAccessTokenResponseBody
	GetRequestId() *string
}

type CreateThirdAppAccessTokenResponseBody struct {
	// The time-limited static access token created for the specified user and single cloud desktop. This is a sensitive credential. Do not log or expose the full value.
	//
	// example:
	//
	// taatv1.v2****
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The absolute expiration time of the token, in UNIX milliseconds. This parameter is not empty when the call is successful. GetLatestThirdAppAccessToken returns the original expiration time without extending it.
	//
	// example:
	//
	// 1795450500000
	ExpiresAt *int64 `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The unique identifier of the request, which can be used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6DAB2D15-7345-4C33-9D77-48A1A756B9E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateThirdAppAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateThirdAppAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateThirdAppAccessTokenResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateThirdAppAccessTokenResponseBody) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *CreateThirdAppAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateThirdAppAccessTokenResponseBody) SetAccessToken(v string) *CreateThirdAppAccessTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *CreateThirdAppAccessTokenResponseBody) SetExpiresAt(v int64) *CreateThirdAppAccessTokenResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *CreateThirdAppAccessTokenResponseBody) SetRequestId(v string) *CreateThirdAppAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateThirdAppAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
