// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListCateInfoHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListCateInfoHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListCateInfoHeaders
	GetAuthorization() *string
}

type ListCateInfoHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCateInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCateInfoHeaders) GoString() string {
	return s.String()
}

func (s *ListCateInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListCateInfoHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListCateInfoHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListCateInfoHeaders) SetCommonHeaders(v map[string]*string) *ListCateInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCateInfoHeaders) SetXAcsAligenieAccessToken(v string) *ListCateInfoHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCateInfoHeaders) SetAuthorization(v string) *ListCateInfoHeaders {
	s.Authorization = &v
	return s
}

func (s *ListCateInfoHeaders) Validate() error {
	return dara.Validate(s)
}
