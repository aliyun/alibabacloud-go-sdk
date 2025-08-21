// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMusicTypeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMusicTypeHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *QueryMusicTypeHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *QueryMusicTypeHeaders
	GetAuthorization() *string
}

type QueryMusicTypeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryMusicTypeHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMusicTypeHeaders) GoString() string {
	return s.String()
}

func (s *QueryMusicTypeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMusicTypeHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *QueryMusicTypeHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *QueryMusicTypeHeaders) SetCommonHeaders(v map[string]*string) *QueryMusicTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMusicTypeHeaders) SetXAcsAligenieAccessToken(v string) *QueryMusicTypeHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryMusicTypeHeaders) SetAuthorization(v string) *QueryMusicTypeHeaders {
	s.Authorization = &v
	return s
}

func (s *QueryMusicTypeHeaders) Validate() error {
	return dara.Validate(s)
}
