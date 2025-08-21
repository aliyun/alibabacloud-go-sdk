// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSubHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddSubHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddSubHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddSubHeaders
	GetAuthorization() *string
}

type AddSubHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddSubHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddSubHeaders) GoString() string {
	return s.String()
}

func (s *AddSubHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddSubHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddSubHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddSubHeaders) SetCommonHeaders(v map[string]*string) *AddSubHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddSubHeaders) SetXAcsAligenieAccessToken(v string) *AddSubHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddSubHeaders) SetAuthorization(v string) *AddSubHeaders {
	s.Authorization = &v
	return s
}

func (s *AddSubHeaders) Validate() error {
	return dara.Validate(s)
}
