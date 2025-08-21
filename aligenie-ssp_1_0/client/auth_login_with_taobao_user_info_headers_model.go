// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithTaobaoUserInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AuthLoginWithTaobaoUserInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AuthLoginWithTaobaoUserInfoHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AuthLoginWithTaobaoUserInfoHeaders
	GetAuthorization() *string
}

type AuthLoginWithTaobaoUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuthLoginWithTaobaoUserInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithTaobaoUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) SetCommonHeaders(v map[string]*string) *AuthLoginWithTaobaoUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) SetXAcsAligenieAccessToken(v string) *AuthLoginWithTaobaoUserInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) SetAuthorization(v string) *AuthLoginWithTaobaoUserInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *AuthLoginWithTaobaoUserInfoHeaders) Validate() error {
	return dara.Validate(s)
}
