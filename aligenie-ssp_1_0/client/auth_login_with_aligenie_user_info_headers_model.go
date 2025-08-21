// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithAligenieUserInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AuthLoginWithAligenieUserInfoHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AuthLoginWithAligenieUserInfoHeaders
	GetAuthorization() *string
}

type AuthLoginWithAligenieUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AuthLoginWithAligenieUserInfoHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AuthLoginWithAligenieUserInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AuthLoginWithAligenieUserInfoHeaders) SetCommonHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoHeaders) SetXAcsAligenieAccessToken(v string) *AuthLoginWithAligenieUserInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoHeaders) SetAuthorization(v string) *AuthLoginWithAligenieUserInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoHeaders) Validate() error {
	return dara.Validate(s)
}
