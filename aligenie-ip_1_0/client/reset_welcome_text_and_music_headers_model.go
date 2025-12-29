// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetWelcomeTextAndMusicHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ResetWelcomeTextAndMusicHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ResetWelcomeTextAndMusicHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ResetWelcomeTextAndMusicHeaders
	GetAuthorization() *string
}

type ResetWelcomeTextAndMusicHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ResetWelcomeTextAndMusicHeaders) String() string {
	return dara.Prettify(s)
}

func (s ResetWelcomeTextAndMusicHeaders) GoString() string {
	return s.String()
}

func (s *ResetWelcomeTextAndMusicHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ResetWelcomeTextAndMusicHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ResetWelcomeTextAndMusicHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ResetWelcomeTextAndMusicHeaders) SetCommonHeaders(v map[string]*string) *ResetWelcomeTextAndMusicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ResetWelcomeTextAndMusicHeaders) SetXAcsAligenieAccessToken(v string) *ResetWelcomeTextAndMusicHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ResetWelcomeTextAndMusicHeaders) SetAuthorization(v string) *ResetWelcomeTextAndMusicHeaders {
	s.Authorization = &v
	return s
}

func (s *ResetWelcomeTextAndMusicHeaders) Validate() error {
	return dara.Validate(s)
}
