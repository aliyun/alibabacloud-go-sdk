// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReminderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteReminderHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteReminderHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteReminderHeaders
	GetAuthorization() *string
}

type DeleteReminderHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteReminderHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteReminderHeaders) GoString() string {
	return s.String()
}

func (s *DeleteReminderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteReminderHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteReminderHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteReminderHeaders) SetCommonHeaders(v map[string]*string) *DeleteReminderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteReminderHeaders) SetXAcsAligenieAccessToken(v string) *DeleteReminderHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteReminderHeaders) SetAuthorization(v string) *DeleteReminderHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteReminderHeaders) Validate() error {
	return dara.Validate(s)
}
