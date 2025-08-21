// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayModeControlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PlayModeControlHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PlayModeControlHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PlayModeControlHeaders
	GetAuthorization() *string
}

type PlayModeControlHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PlayModeControlHeaders) String() string {
	return dara.Prettify(s)
}

func (s PlayModeControlHeaders) GoString() string {
	return s.String()
}

func (s *PlayModeControlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PlayModeControlHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PlayModeControlHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PlayModeControlHeaders) SetCommonHeaders(v map[string]*string) *PlayModeControlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PlayModeControlHeaders) SetXAcsAligenieAccessToken(v string) *PlayModeControlHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PlayModeControlHeaders) SetAuthorization(v string) *PlayModeControlHeaders {
	s.Authorization = &v
	return s
}

func (s *PlayModeControlHeaders) Validate() error {
	return dara.Validate(s)
}
