// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestThirdAppAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetLatestThirdAppAccessTokenResponseBody
	GetAccessToken() *string
	SetExpiresAt(v int64) *GetLatestThirdAppAccessTokenResponseBody
	GetExpiresAt() *int64
	SetRequestId(v string) *GetLatestThirdAppAccessTokenResponseBody
	GetRequestId() *string
}

type GetLatestThirdAppAccessTokenResponseBody struct {
	// The current valid static access token for the specified user and single cloud desktop. The original saved token is returned, and the query does not renew the token. This is a sensitive credential. Do not log or expose the full value.
	//
	// example:
	//
	// taatv1.v2****
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The absolute expiration time of the token, represented as a Unix timestamp in milliseconds. This parameter is not empty when the call is successful. GetLatestThirdAppAccessToken returns the original expiration time and does not renew the token.
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

func (s GetLatestThirdAppAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLatestThirdAppAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestThirdAppAccessTokenResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetLatestThirdAppAccessTokenResponseBody) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *GetLatestThirdAppAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLatestThirdAppAccessTokenResponseBody) SetAccessToken(v string) *GetLatestThirdAppAccessTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenResponseBody) SetExpiresAt(v int64) *GetLatestThirdAppAccessTokenResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenResponseBody) SetRequestId(v string) *GetLatestThirdAppAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLatestThirdAppAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
