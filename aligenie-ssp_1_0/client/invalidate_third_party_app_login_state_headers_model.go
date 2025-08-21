// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidateThirdPartyAppLoginStateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvalidateThirdPartyAppLoginStateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *InvalidateThirdPartyAppLoginStateHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *InvalidateThirdPartyAppLoginStateHeaders
	GetAuthorization() *string
}

type InvalidateThirdPartyAppLoginStateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s InvalidateThirdPartyAppLoginStateHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvalidateThirdPartyAppLoginStateHeaders) GoString() string {
	return s.String()
}

func (s *InvalidateThirdPartyAppLoginStateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvalidateThirdPartyAppLoginStateHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *InvalidateThirdPartyAppLoginStateHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *InvalidateThirdPartyAppLoginStateHeaders) SetCommonHeaders(v map[string]*string) *InvalidateThirdPartyAppLoginStateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateHeaders) SetXAcsAligenieAccessToken(v string) *InvalidateThirdPartyAppLoginStateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateHeaders) SetAuthorization(v string) *InvalidateThirdPartyAppLoginStateHeaders {
	s.Authorization = &v
	return s
}

func (s *InvalidateThirdPartyAppLoginStateHeaders) Validate() error {
	return dara.Validate(s)
}
