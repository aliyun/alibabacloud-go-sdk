// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateWelcomeTextHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddOrUpdateWelcomeTextHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddOrUpdateWelcomeTextHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddOrUpdateWelcomeTextHeaders
	GetAuthorization() *string
}

type AddOrUpdateWelcomeTextHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddOrUpdateWelcomeTextHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateWelcomeTextHeaders) GoString() string {
	return s.String()
}

func (s *AddOrUpdateWelcomeTextHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddOrUpdateWelcomeTextHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddOrUpdateWelcomeTextHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddOrUpdateWelcomeTextHeaders) SetCommonHeaders(v map[string]*string) *AddOrUpdateWelcomeTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrUpdateWelcomeTextHeaders) SetXAcsAligenieAccessToken(v string) *AddOrUpdateWelcomeTextHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddOrUpdateWelcomeTextHeaders) SetAuthorization(v string) *AddOrUpdateWelcomeTextHeaders {
	s.Authorization = &v
	return s
}

func (s *AddOrUpdateWelcomeTextHeaders) Validate() error {
	return dara.Validate(s)
}
