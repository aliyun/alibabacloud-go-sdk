// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelsHeaders
	GetAuthorization() *string
}

type ListHotelsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelsHeaders) SetCommonHeaders(v map[string]*string) *ListHotelsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelsHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelsHeaders) SetAuthorization(v string) *ListHotelsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelsHeaders) Validate() error {
	return dara.Validate(s)
}
