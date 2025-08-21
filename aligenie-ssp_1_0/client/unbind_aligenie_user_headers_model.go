// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAligenieUserHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UnbindAligenieUserHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UnbindAligenieUserHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UnbindAligenieUserHeaders
	GetAuthorization() *string
}

type UnbindAligenieUserHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UnbindAligenieUserHeaders) String() string {
	return dara.Prettify(s)
}

func (s UnbindAligenieUserHeaders) GoString() string {
	return s.String()
}

func (s *UnbindAligenieUserHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UnbindAligenieUserHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UnbindAligenieUserHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UnbindAligenieUserHeaders) SetCommonHeaders(v map[string]*string) *UnbindAligenieUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnbindAligenieUserHeaders) SetXAcsAligenieAccessToken(v string) *UnbindAligenieUserHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UnbindAligenieUserHeaders) SetAuthorization(v string) *UnbindAligenieUserHeaders {
	s.Authorization = &v
	return s
}

func (s *UnbindAligenieUserHeaders) Validate() error {
	return dara.Validate(s)
}
