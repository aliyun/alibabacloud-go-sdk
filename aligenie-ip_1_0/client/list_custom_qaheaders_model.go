// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomQAHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListCustomQAHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListCustomQAHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListCustomQAHeaders
	GetAuthorization() *string
}

type ListCustomQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListCustomQAHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListCustomQAHeaders) GoString() string {
	return s.String()
}

func (s *ListCustomQAHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListCustomQAHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListCustomQAHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListCustomQAHeaders) SetCommonHeaders(v map[string]*string) *ListCustomQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListCustomQAHeaders) SetXAcsAligenieAccessToken(v string) *ListCustomQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListCustomQAHeaders) SetAuthorization(v string) *ListCustomQAHeaders {
	s.Authorization = &v
	return s
}

func (s *ListCustomQAHeaders) Validate() error {
	return dara.Validate(s)
}
