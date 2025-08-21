// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlbumIsAddedHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListAlbumIsAddedHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListAlbumIsAddedHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListAlbumIsAddedHeaders
	GetAuthorization() *string
}

type ListAlbumIsAddedHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAlbumIsAddedHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListAlbumIsAddedHeaders) GoString() string {
	return s.String()
}

func (s *ListAlbumIsAddedHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListAlbumIsAddedHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListAlbumIsAddedHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListAlbumIsAddedHeaders) SetCommonHeaders(v map[string]*string) *ListAlbumIsAddedHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAlbumIsAddedHeaders) SetXAcsAligenieAccessToken(v string) *ListAlbumIsAddedHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListAlbumIsAddedHeaders) SetAuthorization(v string) *ListAlbumIsAddedHeaders {
	s.Authorization = &v
	return s
}

func (s *ListAlbumIsAddedHeaders) Validate() error {
	return dara.Validate(s)
}
