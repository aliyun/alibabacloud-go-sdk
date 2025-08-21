// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListAlbumDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListAlbumDetailHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListAlbumDetailHeaders
	GetAuthorization() *string
}

type ListAlbumDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAlbumDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumDetailHeaders) GoString() string {
	return s.String()
}

func (s *ListAlbumDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListAlbumDetailHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListAlbumDetailHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListAlbumDetailHeaders) SetCommonHeaders(v map[string]*string) *ListAlbumDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAlbumDetailHeaders) SetXAcsAligenieAccessToken(v string) *ListAlbumDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListAlbumDetailHeaders) SetAuthorization(v string) *ListAlbumDetailHeaders {
	s.Authorization = &v
	return s
}

func (s *ListAlbumDetailHeaders) Validate() error {
	return dara.Validate(s)
}
