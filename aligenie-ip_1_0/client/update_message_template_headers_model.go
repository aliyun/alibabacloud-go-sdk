// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMessageTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateMessageTemplateHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateMessageTemplateHeaders
	GetAuthorization() *string
}

type UpdateMessageTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateMessageTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageTemplateHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMessageTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMessageTemplateHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateMessageTemplateHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateMessageTemplateHeaders) SetCommonHeaders(v map[string]*string) *UpdateMessageTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMessageTemplateHeaders) SetXAcsAligenieAccessToken(v string) *UpdateMessageTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateMessageTemplateHeaders) SetAuthorization(v string) *UpdateMessageTemplateHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateMessageTemplateHeaders) Validate() error {
	return dara.Validate(s)
}
