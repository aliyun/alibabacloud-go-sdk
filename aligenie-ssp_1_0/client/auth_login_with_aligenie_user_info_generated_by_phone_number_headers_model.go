// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders
	GetAuthorization() *string
}

type AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) String() string {
	return dara.Prettify(s)
}

func (s AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) SetXAcsAligenieAccessToken(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) SetAuthorization(v string) *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

func (s *AuthLoginWithAligenieUserInfoGeneratedByPhoneNumberHeaders) Validate() error {
	return dara.Validate(s)
}
