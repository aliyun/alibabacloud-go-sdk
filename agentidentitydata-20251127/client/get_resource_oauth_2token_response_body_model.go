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
	SetRequestId(v string) *GetResourceOAuth2TokenResponseBody
	GetRequestId() *string
	SetSessionStatus(v string) *GetResourceOAuth2TokenResponseBody
	GetSessionStatus() *string
	SetSessionURI(v string) *GetResourceOAuth2TokenResponseBody
	GetSessionURI() *string
}

type GetResourceOAuth2TokenResponseBody struct {
	// example:
	//
	// 66a1a2****0b93v3
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// example:
	//
	// https://agentidentity.cn-beijing.aliyuncs.com/*****
	AuthorizationURL *string `json:"AuthorizationURL,omitempty" xml:"AuthorizationURL,omitempty"`
	// example:
	//
	// 173C69C9-9C07-5B25-9477-603A33E5DA32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// IN_PROGRESS
	SessionStatus *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// example:
	//
	// urn:ietf:params:oauth:request_uri:43b7df1a-****-****-****-709375a44d8a
	SessionURI *string `json:"SessionURI,omitempty" xml:"SessionURI,omitempty"`
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
