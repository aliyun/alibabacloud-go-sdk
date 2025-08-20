// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountForAppHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAccountForAppHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetAccountForAppHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetAccountForAppHeaders
	GetAuthorization() *string
}

type GetAccountForAppHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAccountForAppHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAccountForAppHeaders) GoString() string {
	return s.String()
}

func (s *GetAccountForAppHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAccountForAppHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetAccountForAppHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetAccountForAppHeaders) SetCommonHeaders(v map[string]*string) *GetAccountForAppHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAccountForAppHeaders) SetXAcsAligenieAccessToken(v string) *GetAccountForAppHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAccountForAppHeaders) SetAuthorization(v string) *GetAccountForAppHeaders {
	s.Authorization = &v
	return s
}

func (s *GetAccountForAppHeaders) Validate() error {
	return dara.Validate(s)
}
