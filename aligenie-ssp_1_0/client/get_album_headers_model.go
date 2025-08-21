// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlbumHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAlbumHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetAlbumHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetAlbumHeaders
	GetAuthorization() *string
}

type GetAlbumHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAlbumHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumHeaders) GoString() string {
	return s.String()
}

func (s *GetAlbumHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAlbumHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetAlbumHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetAlbumHeaders) SetCommonHeaders(v map[string]*string) *GetAlbumHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAlbumHeaders) SetXAcsAligenieAccessToken(v string) *GetAlbumHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAlbumHeaders) SetAuthorization(v string) *GetAlbumHeaders {
	s.Authorization = &v
	return s
}

func (s *GetAlbumHeaders) Validate() error {
	return dara.Validate(s)
}
