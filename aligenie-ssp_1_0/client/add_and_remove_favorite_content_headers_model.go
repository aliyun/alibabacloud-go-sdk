// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAndRemoveFavoriteContentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddAndRemoveFavoriteContentHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddAndRemoveFavoriteContentHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddAndRemoveFavoriteContentHeaders
	GetAuthorization() *string
}

type AddAndRemoveFavoriteContentHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddAndRemoveFavoriteContentHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentHeaders) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddAndRemoveFavoriteContentHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddAndRemoveFavoriteContentHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddAndRemoveFavoriteContentHeaders) SetCommonHeaders(v map[string]*string) *AddAndRemoveFavoriteContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddAndRemoveFavoriteContentHeaders) SetXAcsAligenieAccessToken(v string) *AddAndRemoveFavoriteContentHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddAndRemoveFavoriteContentHeaders) SetAuthorization(v string) *AddAndRemoveFavoriteContentHeaders {
	s.Authorization = &v
	return s
}

func (s *AddAndRemoveFavoriteContentHeaders) Validate() error {
	return dara.Validate(s)
}
