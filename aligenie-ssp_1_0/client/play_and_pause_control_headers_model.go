// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayAndPauseControlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PlayAndPauseControlHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PlayAndPauseControlHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PlayAndPauseControlHeaders
	GetAuthorization() *string
}

type PlayAndPauseControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PlayAndPauseControlHeaders) String() string {
	return dara.Prettify(s)
}

func (s PlayAndPauseControlHeaders) GoString() string {
	return s.String()
}

func (s *PlayAndPauseControlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PlayAndPauseControlHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PlayAndPauseControlHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PlayAndPauseControlHeaders) SetCommonHeaders(v map[string]*string) *PlayAndPauseControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PlayAndPauseControlHeaders) SetXAcsAligenieAccessToken(v string) *PlayAndPauseControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PlayAndPauseControlHeaders) SetAuthorization(v string) *PlayAndPauseControlHeaders {
	s.Authorization = &v
	return s
}

func (s *PlayAndPauseControlHeaders) Validate() error {
	return dara.Validate(s)
}
