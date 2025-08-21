// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateContentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListCateContentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListCateContentHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListCateContentHeaders
	GetAuthorization() *string
}

type ListCateContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCateContentHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCateContentHeaders) GoString() string {
	return s.String()
}

func (s *ListCateContentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListCateContentHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListCateContentHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListCateContentHeaders) SetCommonHeaders(v map[string]*string) *ListCateContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCateContentHeaders) SetXAcsAligenieAccessToken(v string) *ListCateContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCateContentHeaders) SetAuthorization(v string) *ListCateContentHeaders {
	s.Authorization = &v
	return s
}

func (s *ListCateContentHeaders) Validate() error {
	return dara.Validate(s)
}
