// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMessageTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteMessageTemplateHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteMessageTemplateHeaders
	GetAuthorization() *string
}

type DeleteMessageTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteMessageTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageTemplateHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMessageTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMessageTemplateHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteMessageTemplateHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteMessageTemplateHeaders) SetCommonHeaders(v map[string]*string) *DeleteMessageTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMessageTemplateHeaders) SetXAcsAligenieAccessToken(v string) *DeleteMessageTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteMessageTemplateHeaders) SetAuthorization(v string) *DeleteMessageTemplateHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteMessageTemplateHeaders) Validate() error {
	return dara.Validate(s)
}
