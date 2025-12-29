// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateScreenSaverHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddOrUpdateScreenSaverHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddOrUpdateScreenSaverHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddOrUpdateScreenSaverHeaders
	GetAuthorization() *string
}

type AddOrUpdateScreenSaverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddOrUpdateScreenSaverHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateScreenSaverHeaders) GoString() string {
	return s.String()
}

func (s *AddOrUpdateScreenSaverHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddOrUpdateScreenSaverHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddOrUpdateScreenSaverHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddOrUpdateScreenSaverHeaders) SetCommonHeaders(v map[string]*string) *AddOrUpdateScreenSaverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddOrUpdateScreenSaverHeaders) SetXAcsAligenieAccessToken(v string) *AddOrUpdateScreenSaverHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddOrUpdateScreenSaverHeaders) SetAuthorization(v string) *AddOrUpdateScreenSaverHeaders {
	s.Authorization = &v
	return s
}

func (s *AddOrUpdateScreenSaverHeaders) Validate() error {
	return dara.Validate(s)
}
