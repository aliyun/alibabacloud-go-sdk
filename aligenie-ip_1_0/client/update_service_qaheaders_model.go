// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceQAHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateServiceQAHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateServiceQAHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateServiceQAHeaders
	GetAuthorization() *string
}

type UpdateServiceQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateServiceQAHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceQAHeaders) GoString() string {
	return s.String()
}

func (s *UpdateServiceQAHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateServiceQAHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateServiceQAHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateServiceQAHeaders) SetCommonHeaders(v map[string]*string) *UpdateServiceQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateServiceQAHeaders) SetXAcsAligenieAccessToken(v string) *UpdateServiceQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateServiceQAHeaders) SetAuthorization(v string) *UpdateServiceQAHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateServiceQAHeaders) Validate() error {
	return dara.Validate(s)
}
