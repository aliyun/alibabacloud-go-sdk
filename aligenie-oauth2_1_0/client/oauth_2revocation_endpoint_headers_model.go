// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2RevocationEndpointHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *OAuth2RevocationEndpointHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *OAuth2RevocationEndpointHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *OAuth2RevocationEndpointHeaders
	GetAuthorization() *string
}

type OAuth2RevocationEndpointHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s OAuth2RevocationEndpointHeaders) String() string {
	return dara.Prettify(s)
}

func (s OAuth2RevocationEndpointHeaders) GoString() string {
	return s.String()
}

func (s *OAuth2RevocationEndpointHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *OAuth2RevocationEndpointHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *OAuth2RevocationEndpointHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *OAuth2RevocationEndpointHeaders) SetCommonHeaders(v map[string]*string) *OAuth2RevocationEndpointHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OAuth2RevocationEndpointHeaders) SetXAcsAligenieAccessToken(v string) *OAuth2RevocationEndpointHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *OAuth2RevocationEndpointHeaders) SetAuthorization(v string) *OAuth2RevocationEndpointHeaders {
	s.Authorization = &v
	return s
}

func (s *OAuth2RevocationEndpointHeaders) Validate() error {
	return dara.Validate(s)
}
