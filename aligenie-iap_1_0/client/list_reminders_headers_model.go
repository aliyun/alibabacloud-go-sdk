// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemindersHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListRemindersHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListRemindersHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListRemindersHeaders
	GetAuthorization() *string
}

type ListRemindersHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListRemindersHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListRemindersHeaders) GoString() string {
	return s.String()
}

func (s *ListRemindersHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListRemindersHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListRemindersHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListRemindersHeaders) SetCommonHeaders(v map[string]*string) *ListRemindersHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListRemindersHeaders) SetXAcsAligenieAccessToken(v string) *ListRemindersHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListRemindersHeaders) SetAuthorization(v string) *ListRemindersHeaders {
	s.Authorization = &v
	return s
}

func (s *ListRemindersHeaders) Validate() error {
	return dara.Validate(s)
}
