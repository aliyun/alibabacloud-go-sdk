// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubAlbumHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSubAlbumHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListSubAlbumHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListSubAlbumHeaders
	GetAuthorization() *string
}

type ListSubAlbumHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSubAlbumHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSubAlbumHeaders) GoString() string {
	return s.String()
}

func (s *ListSubAlbumHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSubAlbumHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListSubAlbumHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListSubAlbumHeaders) SetCommonHeaders(v map[string]*string) *ListSubAlbumHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubAlbumHeaders) SetXAcsAligenieAccessToken(v string) *ListSubAlbumHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSubAlbumHeaders) SetAuthorization(v string) *ListSubAlbumHeaders {
	s.Authorization = &v
	return s
}

func (s *ListSubAlbumHeaders) Validate() error {
	return dara.Validate(s)
}
