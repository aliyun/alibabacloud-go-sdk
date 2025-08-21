// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCurrentPlayingListHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetCurrentPlayingListHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetCurrentPlayingListHeaders
	GetAuthorization() *string
}

type GetCurrentPlayingListHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetCurrentPlayingListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListHeaders) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCurrentPlayingListHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetCurrentPlayingListHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetCurrentPlayingListHeaders) SetCommonHeaders(v map[string]*string) *GetCurrentPlayingListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCurrentPlayingListHeaders) SetXAcsAligenieAccessToken(v string) *GetCurrentPlayingListHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetCurrentPlayingListHeaders) SetAuthorization(v string) *GetCurrentPlayingListHeaders {
	s.Authorization = &v
	return s
}

func (s *GetCurrentPlayingListHeaders) Validate() error {
	return dara.Validate(s)
}
