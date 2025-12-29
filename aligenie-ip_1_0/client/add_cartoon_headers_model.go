// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCartoonHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddCartoonHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddCartoonHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddCartoonHeaders
	GetAuthorization() *string
}

type AddCartoonHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddCartoonHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddCartoonHeaders) GoString() string {
	return s.String()
}

func (s *AddCartoonHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddCartoonHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddCartoonHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddCartoonHeaders) SetCommonHeaders(v map[string]*string) *AddCartoonHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddCartoonHeaders) SetXAcsAligenieAccessToken(v string) *AddCartoonHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddCartoonHeaders) SetAuthorization(v string) *AddCartoonHeaders {
	s.Authorization = &v
	return s
}

func (s *AddCartoonHeaders) Validate() error {
	return dara.Validate(s)
}
