// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMessageTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddMessageTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *AddMessageTemplateHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *AddMessageTemplateHeaders
	GetAuthorization() *string
}

type AddMessageTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s AddMessageTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddMessageTemplateHeaders) GoString() string {
	return s.String()
}

func (s *AddMessageTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddMessageTemplateHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *AddMessageTemplateHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *AddMessageTemplateHeaders) SetCommonHeaders(v map[string]*string) *AddMessageTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMessageTemplateHeaders) SetXAcsAligenieAccessToken(v string) *AddMessageTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *AddMessageTemplateHeaders) SetAuthorization(v string) *AddMessageTemplateHeaders {
	s.Authorization = &v
	return s
}

func (s *AddMessageTemplateHeaders) Validate() error {
	return dara.Validate(s)
}
