// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchContentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *SearchContentHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *SearchContentHeaders
	GetAuthorization() *string
}

type SearchContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SearchContentHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchContentHeaders) GoString() string {
	return s.String()
}

func (s *SearchContentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchContentHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *SearchContentHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *SearchContentHeaders) SetCommonHeaders(v map[string]*string) *SearchContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchContentHeaders) SetXAcsAligenieAccessToken(v string) *SearchContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SearchContentHeaders) SetAuthorization(v string) *SearchContentHeaders {
	s.Authorization = &v
	return s
}

func (s *SearchContentHeaders) Validate() error {
	return dara.Validate(s)
}
