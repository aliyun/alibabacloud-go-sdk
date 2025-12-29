// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSTBServiceProvidersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSTBServiceProvidersHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListSTBServiceProvidersHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListSTBServiceProvidersHeaders
	GetAuthorization() *string
}

type ListSTBServiceProvidersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSTBServiceProvidersHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSTBServiceProvidersHeaders) GoString() string {
	return s.String()
}

func (s *ListSTBServiceProvidersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSTBServiceProvidersHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListSTBServiceProvidersHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListSTBServiceProvidersHeaders) SetCommonHeaders(v map[string]*string) *ListSTBServiceProvidersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSTBServiceProvidersHeaders) SetXAcsAligenieAccessToken(v string) *ListSTBServiceProvidersHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSTBServiceProvidersHeaders) SetAuthorization(v string) *ListSTBServiceProvidersHeaders {
	s.Authorization = &v
	return s
}

func (s *ListSTBServiceProvidersHeaders) Validate() error {
	return dara.Validate(s)
}
