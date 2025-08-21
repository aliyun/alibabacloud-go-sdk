// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ReadMessageHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ReadMessageHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ReadMessageHeaders
	GetAuthorization() *string
}

type ReadMessageHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ReadMessageHeaders) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageHeaders) GoString() string {
	return s.String()
}

func (s *ReadMessageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ReadMessageHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ReadMessageHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ReadMessageHeaders) SetCommonHeaders(v map[string]*string) *ReadMessageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ReadMessageHeaders) SetXAcsAligenieAccessToken(v string) *ReadMessageHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ReadMessageHeaders) SetAuthorization(v string) *ReadMessageHeaders {
	s.Authorization = &v
	return s
}

func (s *ReadMessageHeaders) Validate() error {
	return dara.Validate(s)
}
