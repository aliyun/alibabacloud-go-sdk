// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWelcomeTextAndMusicHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetWelcomeTextAndMusicHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetWelcomeTextAndMusicHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetWelcomeTextAndMusicHeaders
	GetAuthorization() *string
}

type GetWelcomeTextAndMusicHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetWelcomeTextAndMusicHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetWelcomeTextAndMusicHeaders) GoString() string {
	return s.String()
}

func (s *GetWelcomeTextAndMusicHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetWelcomeTextAndMusicHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetWelcomeTextAndMusicHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetWelcomeTextAndMusicHeaders) SetCommonHeaders(v map[string]*string) *GetWelcomeTextAndMusicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetWelcomeTextAndMusicHeaders) SetXAcsAligenieAccessToken(v string) *GetWelcomeTextAndMusicHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetWelcomeTextAndMusicHeaders) SetAuthorization(v string) *GetWelcomeTextAndMusicHeaders {
	s.Authorization = &v
	return s
}

func (s *GetWelcomeTextAndMusicHeaders) Validate() error {
	return dara.Validate(s)
}
