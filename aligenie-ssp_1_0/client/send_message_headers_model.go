// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SendMessageHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *SendMessageHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *SendMessageHeaders
	GetAuthorization() *string
}

type SendMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s SendMessageHeaders) String() string {
	return dara.Prettify(s)
}

func (s SendMessageHeaders) GoString() string {
	return s.String()
}

func (s *SendMessageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SendMessageHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *SendMessageHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *SendMessageHeaders) SetCommonHeaders(v map[string]*string) *SendMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SendMessageHeaders) SetXAcsAligenieAccessToken(v string) *SendMessageHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *SendMessageHeaders) SetAuthorization(v string) *SendMessageHeaders {
	s.Authorization = &v
	return s
}

func (s *SendMessageHeaders) Validate() error {
	return dara.Validate(s)
}
