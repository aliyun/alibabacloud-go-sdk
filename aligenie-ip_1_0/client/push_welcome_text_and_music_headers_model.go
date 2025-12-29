// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeTextAndMusicHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PushWelcomeTextAndMusicHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PushWelcomeTextAndMusicHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PushWelcomeTextAndMusicHeaders
	GetAuthorization() *string
}

type PushWelcomeTextAndMusicHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushWelcomeTextAndMusicHeaders) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeTextAndMusicHeaders) GoString() string {
	return s.String()
}

func (s *PushWelcomeTextAndMusicHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PushWelcomeTextAndMusicHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PushWelcomeTextAndMusicHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PushWelcomeTextAndMusicHeaders) SetCommonHeaders(v map[string]*string) *PushWelcomeTextAndMusicHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushWelcomeTextAndMusicHeaders) SetXAcsAligenieAccessToken(v string) *PushWelcomeTextAndMusicHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushWelcomeTextAndMusicHeaders) SetAuthorization(v string) *PushWelcomeTextAndMusicHeaders {
	s.Authorization = &v
	return s
}

func (s *PushWelcomeTextAndMusicHeaders) Validate() error {
	return dara.Validate(s)
}
