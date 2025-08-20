// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReminderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateReminderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CreateReminderHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CreateReminderHeaders
	GetAuthorization() *string
}

type CreateReminderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateReminderHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderHeaders) GoString() string {
	return s.String()
}

func (s *CreateReminderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateReminderHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CreateReminderHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateReminderHeaders) SetCommonHeaders(v map[string]*string) *CreateReminderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateReminderHeaders) SetXAcsAligenieAccessToken(v string) *CreateReminderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateReminderHeaders) SetAuthorization(v string) *CreateReminderHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateReminderHeaders) Validate() error {
	return dara.Validate(s)
}
