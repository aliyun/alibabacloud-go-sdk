// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlbumDetailByIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAlbumDetailByIdHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetAlbumDetailByIdHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetAlbumDetailByIdHeaders
	GetAuthorization() *string
}

type GetAlbumDetailByIdHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAlbumDetailByIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAlbumDetailByIdHeaders) GoString() string {
	return s.String()
}

func (s *GetAlbumDetailByIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAlbumDetailByIdHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetAlbumDetailByIdHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetAlbumDetailByIdHeaders) SetCommonHeaders(v map[string]*string) *GetAlbumDetailByIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAlbumDetailByIdHeaders) SetXAcsAligenieAccessToken(v string) *GetAlbumDetailByIdHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAlbumDetailByIdHeaders) SetAuthorization(v string) *GetAlbumDetailByIdHeaders {
	s.Authorization = &v
	return s
}

func (s *GetAlbumDetailByIdHeaders) Validate() error {
	return dara.Validate(s)
}
