// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOauthTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateOauthTokenResponseBody
	GetRequestId() *string
	SetTokenResponse(v *GenerateOauthTokenResponseBodyTokenResponse) *GenerateOauthTokenResponseBody
	GetTokenResponse() *GenerateOauthTokenResponseBodyTokenResponse
}

type GenerateOauthTokenResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4XXXXXXX
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TokenResponse *GenerateOauthTokenResponseBodyTokenResponse `json:"TokenResponse,omitempty" xml:"TokenResponse,omitempty" type:"Struct"`
}

func (s GenerateOauthTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateOauthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOauthTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateOauthTokenResponseBody) GetTokenResponse() *GenerateOauthTokenResponseBodyTokenResponse {
	return s.TokenResponse
}

func (s *GenerateOauthTokenResponseBody) SetRequestId(v string) *GenerateOauthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateOauthTokenResponseBody) SetTokenResponse(v *GenerateOauthTokenResponseBodyTokenResponse) *GenerateOauthTokenResponseBody {
	s.TokenResponse = v
	return s
}

func (s *GenerateOauthTokenResponseBody) Validate() error {
	if s.TokenResponse != nil {
		if err := s.TokenResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateOauthTokenResponseBodyTokenResponse struct {
	// Access Token。
	//
	// example:
	//
	// ***
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// 1770186372
	ExpiresAt *int64 `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// example:
	//
	// 1200
	ExpiresIn *int64 `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	// example:
	//
	// Bearer
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s GenerateOauthTokenResponseBodyTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateOauthTokenResponseBodyTokenResponse) GoString() string {
	return s.String()
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) GetExpiresAt() *int64 {
	return s.ExpiresAt
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) GetTokenType() *string {
	return s.TokenType
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) SetAccessToken(v string) *GenerateOauthTokenResponseBodyTokenResponse {
	s.AccessToken = &v
	return s
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) SetExpiresAt(v int64) *GenerateOauthTokenResponseBodyTokenResponse {
	s.ExpiresAt = &v
	return s
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) SetExpiresIn(v int64) *GenerateOauthTokenResponseBodyTokenResponse {
	s.ExpiresIn = &v
	return s
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) SetTokenType(v string) *GenerateOauthTokenResponseBodyTokenResponse {
	s.TokenType = &v
	return s
}

func (s *GenerateOauthTokenResponseBodyTokenResponse) Validate() error {
	return dara.Validate(s)
}
