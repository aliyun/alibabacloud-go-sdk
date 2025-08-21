// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSubHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListSubHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListSubHeaders
	GetAuthorization() *string
}

type ListSubHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSubHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSubHeaders) GoString() string {
	return s.String()
}

func (s *ListSubHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSubHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListSubHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListSubHeaders) SetCommonHeaders(v map[string]*string) *ListSubHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubHeaders) SetXAcsAligenieAccessToken(v string) *ListSubHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSubHeaders) SetAuthorization(v string) *ListSubHeaders {
	s.Authorization = &v
	return s
}

func (s *ListSubHeaders) Validate() error {
	return dara.Validate(s)
}
