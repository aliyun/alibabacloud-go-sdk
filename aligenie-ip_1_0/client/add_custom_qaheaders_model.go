// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddCustomQAHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddCustomQAHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddCustomQAHeaders
	GetAuthorization() *string
}

type AddCustomQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddCustomQAHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAHeaders) GoString() string {
	return s.String()
}

func (s *AddCustomQAHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddCustomQAHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddCustomQAHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddCustomQAHeaders) SetCommonHeaders(v map[string]*string) *AddCustomQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCustomQAHeaders) SetXAcsAligenieAccessToken(v string) *AddCustomQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddCustomQAHeaders) SetAuthorization(v string) *AddCustomQAHeaders {
	s.Authorization = &v
	return s
}

func (s *AddCustomQAHeaders) Validate() error {
	return dara.Validate(s)
}
