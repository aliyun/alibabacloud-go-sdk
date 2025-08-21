// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindUserlistToAuthLoginWithPhoneNumberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FindUserlistToAuthLoginWithPhoneNumberHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *FindUserlistToAuthLoginWithPhoneNumberHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *FindUserlistToAuthLoginWithPhoneNumberHeaders
	GetAuthorization() *string
}

type FindUserlistToAuthLoginWithPhoneNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberHeaders) String() string {
	return dara.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *FindUserlistToAuthLoginWithPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) SetXAcsAligenieAccessToken(v string) *FindUserlistToAuthLoginWithPhoneNumberHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) SetAuthorization(v string) *FindUserlistToAuthLoginWithPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberHeaders) Validate() error {
	return dara.Validate(s)
}
