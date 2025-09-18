// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2TokenEndpointHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *OAuth2TokenEndpointHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *OAuth2TokenEndpointHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *OAuth2TokenEndpointHeaders
	GetAuthorization() *string
}

type OAuth2TokenEndpointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s OAuth2TokenEndpointHeaders) String() string {
	return dara.Prettify(s)
}

func (s OAuth2TokenEndpointHeaders) GoString() string {
	return s.String()
}

func (s *OAuth2TokenEndpointHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *OAuth2TokenEndpointHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *OAuth2TokenEndpointHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *OAuth2TokenEndpointHeaders) SetCommonHeaders(v map[string]*string) *OAuth2TokenEndpointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OAuth2TokenEndpointHeaders) SetXAcsAligenieAccessToken(v string) *OAuth2TokenEndpointHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *OAuth2TokenEndpointHeaders) SetAuthorization(v string) *OAuth2TokenEndpointHeaders {
	s.Authorization = &v
	return s
}

func (s *OAuth2TokenEndpointHeaders) Validate() error {
	return dara.Validate(s)
}
