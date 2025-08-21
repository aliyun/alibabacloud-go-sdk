// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionAlbumCategoryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSubscriptionAlbumCategoryHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListSubscriptionAlbumCategoryHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListSubscriptionAlbumCategoryHeaders
	GetAuthorization() *string
}

type ListSubscriptionAlbumCategoryHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListSubscriptionAlbumCategoryHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionAlbumCategoryHeaders) GoString() string {
	return s.String()
}

func (s *ListSubscriptionAlbumCategoryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSubscriptionAlbumCategoryHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListSubscriptionAlbumCategoryHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListSubscriptionAlbumCategoryHeaders) SetCommonHeaders(v map[string]*string) *ListSubscriptionAlbumCategoryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSubscriptionAlbumCategoryHeaders) SetXAcsAligenieAccessToken(v string) *ListSubscriptionAlbumCategoryHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryHeaders) SetAuthorization(v string) *ListSubscriptionAlbumCategoryHeaders {
	s.Authorization = &v
	return s
}

func (s *ListSubscriptionAlbumCategoryHeaders) Validate() error {
	return dara.Validate(s)
}
