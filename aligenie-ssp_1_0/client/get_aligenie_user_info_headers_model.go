// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAligenieUserInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAligenieUserInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetAligenieUserInfoHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetAligenieUserInfoHeaders
	GetAuthorization() *string
}

type GetAligenieUserInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAligenieUserInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAligenieUserInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAligenieUserInfoHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetAligenieUserInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetAligenieUserInfoHeaders) SetCommonHeaders(v map[string]*string) *GetAligenieUserInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAligenieUserInfoHeaders) SetXAcsAligenieAccessToken(v string) *GetAligenieUserInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAligenieUserInfoHeaders) SetAuthorization(v string) *GetAligenieUserInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *GetAligenieUserInfoHeaders) Validate() error {
	return dara.Validate(s)
}
