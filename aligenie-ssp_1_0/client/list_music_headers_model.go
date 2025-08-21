// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMusicHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListMusicHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListMusicHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListMusicHeaders
	GetAuthorization() *string
}

type ListMusicHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListMusicHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListMusicHeaders) GoString() string {
	return s.String()
}

func (s *ListMusicHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListMusicHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListMusicHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListMusicHeaders) SetCommonHeaders(v map[string]*string) *ListMusicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListMusicHeaders) SetXAcsAligenieAccessToken(v string) *ListMusicHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListMusicHeaders) SetAuthorization(v string) *ListMusicHeaders {
	s.Authorization = &v
	return s
}

func (s *ListMusicHeaders) Validate() error {
	return dara.Validate(s)
}
