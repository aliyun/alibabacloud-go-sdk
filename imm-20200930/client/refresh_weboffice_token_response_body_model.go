// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshWebofficeTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *RefreshWebofficeTokenResponseBody
	GetAccessToken() *string
	SetAccessTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody
	GetAccessTokenExpiredTime() *string
	SetRefreshToken(v string) *RefreshWebofficeTokenResponseBody
	GetRefreshToken() *string
	SetRefreshTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody
	GetRefreshTokenExpiredTime() *string
	SetRequestId(v string) *RefreshWebofficeTokenResponseBody
	GetRequestId() *string
}

type RefreshWebofficeTokenResponseBody struct {
	// Weboffice access token.
	//
	// example:
	//
	// 4996466c690a4902846ce00f96********
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// Expiration time of the access token. The expiration period is 30 minutes.
	//
	// example:
	//
	// 2021-08-31T13:07:28.950065359Z
	AccessTokenExpiredTime *string `json:"AccessTokenExpiredTime,omitempty" xml:"AccessTokenExpiredTime,omitempty"`
	// Weboffice refresh token.
	//
	// example:
	//
	// 72a52ab3702a4123ab5594671a********
	RefreshToken *string `json:"RefreshToken,omitempty" xml:"RefreshToken,omitempty"`
	// Expiration time of the refresh token. The expiration period is 1 day.
	//
	// example:
	//
	// 2021-09-01T12:37:28.950065359Z
	RefreshTokenExpiredTime *string `json:"RefreshTokenExpiredTime,omitempty" xml:"RefreshTokenExpiredTime,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 501339F9-4B70-0CE2-AB8C-866C********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshWebofficeTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshWebofficeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshWebofficeTokenResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *RefreshWebofficeTokenResponseBody) GetAccessTokenExpiredTime() *string {
	return s.AccessTokenExpiredTime
}

func (s *RefreshWebofficeTokenResponseBody) GetRefreshToken() *string {
	return s.RefreshToken
}

func (s *RefreshWebofficeTokenResponseBody) GetRefreshTokenExpiredTime() *string {
	return s.RefreshTokenExpiredTime
}

func (s *RefreshWebofficeTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessToken(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetAccessTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.AccessTokenExpiredTime = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshToken(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshToken = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRefreshTokenExpiredTime(v string) *RefreshWebofficeTokenResponseBody {
	s.RefreshTokenExpiredTime = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) SetRequestId(v string) *RefreshWebofficeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshWebofficeTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
