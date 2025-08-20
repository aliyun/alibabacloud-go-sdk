// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetPhoneNumberHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetPhoneNumberHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetPhoneNumberHeaders
	GetAuthorization() *string
}

type GetPhoneNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetPhoneNumberHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberHeaders) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetPhoneNumberHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetPhoneNumberHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetPhoneNumberHeaders) SetCommonHeaders(v map[string]*string) *GetPhoneNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetPhoneNumberHeaders) SetXAcsAligenieAccessToken(v string) *GetPhoneNumberHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetPhoneNumberHeaders) SetAuthorization(v string) *GetPhoneNumberHeaders {
	s.Authorization = &v
	return s
}

func (s *GetPhoneNumberHeaders) Validate() error {
	return dara.Validate(s)
}
