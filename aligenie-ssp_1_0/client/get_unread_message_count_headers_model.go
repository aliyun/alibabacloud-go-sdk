// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnreadMessageCountHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUnreadMessageCountHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetUnreadMessageCountHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetUnreadMessageCountHeaders
	GetAuthorization() *string
}

type GetUnreadMessageCountHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetUnreadMessageCountHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUnreadMessageCountHeaders) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUnreadMessageCountHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetUnreadMessageCountHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetUnreadMessageCountHeaders) SetCommonHeaders(v map[string]*string) *GetUnreadMessageCountHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUnreadMessageCountHeaders) SetXAcsAligenieAccessToken(v string) *GetUnreadMessageCountHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetUnreadMessageCountHeaders) SetAuthorization(v string) *GetUnreadMessageCountHeaders {
	s.Authorization = &v
	return s
}

func (s *GetUnreadMessageCountHeaders) Validate() error {
	return dara.Validate(s)
}
