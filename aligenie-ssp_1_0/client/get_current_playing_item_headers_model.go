// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingItemHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCurrentPlayingItemHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetCurrentPlayingItemHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetCurrentPlayingItemHeaders
	GetAuthorization() *string
}

type GetCurrentPlayingItemHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetCurrentPlayingItemHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemHeaders) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCurrentPlayingItemHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetCurrentPlayingItemHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetCurrentPlayingItemHeaders) SetCommonHeaders(v map[string]*string) *GetCurrentPlayingItemHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCurrentPlayingItemHeaders) SetXAcsAligenieAccessToken(v string) *GetCurrentPlayingItemHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetCurrentPlayingItemHeaders) SetAuthorization(v string) *GetCurrentPlayingItemHeaders {
	s.Authorization = &v
	return s
}

func (s *GetCurrentPlayingItemHeaders) Validate() error {
	return dara.Validate(s)
}
