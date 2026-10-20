// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceOAuth2TokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetResourceOAuth2TokenResponseBody
	GetAccessToken() *string
	SetAuthorizationURL(v string) *GetResourceOAuth2TokenResponseBody
	GetAuthorizationURL() *string
	SetExpiresIn(v int64) *GetResourceOAuth2TokenResponseBody
	GetExpiresIn() *int64
	SetRequestId(v string) *GetResourceOAuth2TokenResponseBody
	GetRequestId() *string
	SetSessionStatus(v string) *GetResourceOAuth2TokenResponseBody
	GetSessionStatus() *string
	SetSessionURI(v string) *GetResourceOAuth2TokenResponseBody
	GetSessionURI() *string
}

type GetResourceOAuth2TokenResponseBody struct {
	AccessToken      *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	AuthorizationURL *string `json:"AuthorizationURL,omitempty" xml:"AuthorizationURL,omitempty"`
	ExpiresIn        *int64  `json:"ExpiresIn,omitempty" xml:"ExpiresIn,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionStatus    *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	SessionURI       *string `json:"SessionURI,omitempty" xml:"SessionURI,omitempty"`
}

func (s GetResourceOAuth2TokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceOAuth2TokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceOAuth2TokenResponseBody) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetResourceOAuth2TokenResponseBody) GetAuthorizationURL() *string {
	return s.AuthorizationURL
}

func (s *GetResourceOAuth2TokenResponseBody) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *GetResourceOAuth2TokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceOAuth2TokenResponseBody) GetSessionStatus() *string {
	return s.SessionStatus
}

func (s *GetResourceOAuth2TokenResponseBody) GetSessionURI() *string {
	return s.SessionURI
}

func (s *GetResourceOAuth2TokenResponseBody) SetAccessToken(v string) *GetResourceOAuth2TokenResponseBody {
	s.AccessToken = &v
	return s
}

func (s *GetResourceOAuth2TokenResponseBody) SetAuthorizationURL(v string) *GetResourceOAuth2TokenResponseBody {
	s.AuthorizationURL = &v
	return s
}

func (s *GetResourceOAuth2TokenResponseBody) SetExpiresIn(v int64) *GetResourceOAuth2TokenResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GetResourceOAuth2TokenResponseBody) SetRequestId(v string) *GetResourceOAuth2TokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceOAuth2TokenResponseBody) SetSessionStatus(v string) *GetResourceOAuth2TokenResponseBody {
	s.SessionStatus = &v
	return s
}

func (s *GetResourceOAuth2TokenResponseBody) SetSessionURI(v string) *GetResourceOAuth2TokenResponseBody {
	s.SessionURI = &v
	return s
}

func (s *GetResourceOAuth2TokenResponseBody) Validate() error {
	return dara.Validate(s)
}
