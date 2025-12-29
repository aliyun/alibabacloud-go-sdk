// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomQAHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateCustomQAHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateCustomQAHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateCustomQAHeaders
	GetAuthorization() *string
}

type UpdateCustomQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateCustomQAHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomQAHeaders) GoString() string {
	return s.String()
}

func (s *UpdateCustomQAHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateCustomQAHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateCustomQAHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateCustomQAHeaders) SetCommonHeaders(v map[string]*string) *UpdateCustomQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateCustomQAHeaders) SetXAcsAligenieAccessToken(v string) *UpdateCustomQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateCustomQAHeaders) SetAuthorization(v string) *UpdateCustomQAHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateCustomQAHeaders) Validate() error {
	return dara.Validate(s)
}
