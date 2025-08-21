// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAuthCodeBindForExtHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CheckAuthCodeBindForExtHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CheckAuthCodeBindForExtHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CheckAuthCodeBindForExtHeaders
	GetAuthorization() *string
}

type CheckAuthCodeBindForExtHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CheckAuthCodeBindForExtHeaders) String() string {
	return dara.Prettify(s)
}

func (s CheckAuthCodeBindForExtHeaders) GoString() string {
	return s.String()
}

func (s *CheckAuthCodeBindForExtHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CheckAuthCodeBindForExtHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CheckAuthCodeBindForExtHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CheckAuthCodeBindForExtHeaders) SetCommonHeaders(v map[string]*string) *CheckAuthCodeBindForExtHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CheckAuthCodeBindForExtHeaders) SetXAcsAligenieAccessToken(v string) *CheckAuthCodeBindForExtHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CheckAuthCodeBindForExtHeaders) SetAuthorization(v string) *CheckAuthCodeBindForExtHeaders {
	s.Authorization = &v
	return s
}

func (s *CheckAuthCodeBindForExtHeaders) Validate() error {
	return dara.Validate(s)
}
