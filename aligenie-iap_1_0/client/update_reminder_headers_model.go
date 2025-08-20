// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReminderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateReminderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateReminderHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateReminderHeaders
	GetAuthorization() *string
}

type UpdateReminderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateReminderHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderHeaders) GoString() string {
	return s.String()
}

func (s *UpdateReminderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateReminderHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateReminderHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateReminderHeaders) SetCommonHeaders(v map[string]*string) *UpdateReminderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateReminderHeaders) SetXAcsAligenieAccessToken(v string) *UpdateReminderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateReminderHeaders) SetAuthorization(v string) *UpdateReminderHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateReminderHeaders) Validate() error {
	return dara.Validate(s)
}
