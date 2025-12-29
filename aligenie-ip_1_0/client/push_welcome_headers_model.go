// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushWelcomeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PushWelcomeHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PushWelcomeHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PushWelcomeHeaders
	GetAuthorization() *string
}

type PushWelcomeHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PushWelcomeHeaders) String() string {
	return dara.Prettify(s)
}

func (s PushWelcomeHeaders) GoString() string {
	return s.String()
}

func (s *PushWelcomeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PushWelcomeHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PushWelcomeHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PushWelcomeHeaders) SetCommonHeaders(v map[string]*string) *PushWelcomeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PushWelcomeHeaders) SetXAcsAligenieAccessToken(v string) *PushWelcomeHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PushWelcomeHeaders) SetAuthorization(v string) *PushWelcomeHeaders {
	s.Authorization = &v
	return s
}

func (s *PushWelcomeHeaders) Validate() error {
	return dara.Validate(s)
}
