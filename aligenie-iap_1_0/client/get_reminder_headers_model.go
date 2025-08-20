// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReminderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetReminderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetReminderHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetReminderHeaders
	GetAuthorization() *string
}

type GetReminderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetReminderHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetReminderHeaders) GoString() string {
	return s.String()
}

func (s *GetReminderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetReminderHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetReminderHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetReminderHeaders) SetCommonHeaders(v map[string]*string) *GetReminderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetReminderHeaders) SetXAcsAligenieAccessToken(v string) *GetReminderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetReminderHeaders) SetAuthorization(v string) *GetReminderHeaders {
	s.Authorization = &v
	return s
}

func (s *GetReminderHeaders) Validate() error {
	return dara.Validate(s)
}
