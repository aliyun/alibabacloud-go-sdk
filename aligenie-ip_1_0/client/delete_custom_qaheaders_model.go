// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomQAHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteCustomQAHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteCustomQAHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteCustomQAHeaders
	GetAuthorization() *string
}

type DeleteCustomQAHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteCustomQAHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomQAHeaders) GoString() string {
	return s.String()
}

func (s *DeleteCustomQAHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteCustomQAHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteCustomQAHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteCustomQAHeaders) SetCommonHeaders(v map[string]*string) *DeleteCustomQAHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteCustomQAHeaders) SetXAcsAligenieAccessToken(v string) *DeleteCustomQAHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteCustomQAHeaders) SetAuthorization(v string) *DeleteCustomQAHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteCustomQAHeaders) Validate() error {
	return dara.Validate(s)
}
