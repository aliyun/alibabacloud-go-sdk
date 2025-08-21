// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserMessageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListUserMessageHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListUserMessageHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListUserMessageHeaders
	GetAuthorization() *string
}

type ListUserMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListUserMessageHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListUserMessageHeaders) GoString() string {
	return s.String()
}

func (s *ListUserMessageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListUserMessageHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListUserMessageHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListUserMessageHeaders) SetCommonHeaders(v map[string]*string) *ListUserMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListUserMessageHeaders) SetXAcsAligenieAccessToken(v string) *ListUserMessageHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListUserMessageHeaders) SetAuthorization(v string) *ListUserMessageHeaders {
	s.Authorization = &v
	return s
}

func (s *ListUserMessageHeaders) Validate() error {
	return dara.Validate(s)
}
