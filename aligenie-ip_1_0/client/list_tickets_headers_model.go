// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTicketsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListTicketsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListTicketsHeaders
	GetAuthorization() *string
}

type ListTicketsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListTicketsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsHeaders) GoString() string {
	return s.String()
}

func (s *ListTicketsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTicketsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListTicketsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListTicketsHeaders) SetCommonHeaders(v map[string]*string) *ListTicketsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTicketsHeaders) SetXAcsAligenieAccessToken(v string) *ListTicketsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListTicketsHeaders) SetAuthorization(v string) *ListTicketsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListTicketsHeaders) Validate() error {
	return dara.Validate(s)
}
