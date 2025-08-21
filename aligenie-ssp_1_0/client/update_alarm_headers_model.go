// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlarmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateAlarmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateAlarmHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateAlarmHeaders
	GetAuthorization() *string
}

type UpdateAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateAlarmHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlarmHeaders) GoString() string {
	return s.String()
}

func (s *UpdateAlarmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateAlarmHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateAlarmHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateAlarmHeaders) SetCommonHeaders(v map[string]*string) *UpdateAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateAlarmHeaders) SetXAcsAligenieAccessToken(v string) *UpdateAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateAlarmHeaders) SetAuthorization(v string) *UpdateAlarmHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateAlarmHeaders) Validate() error {
	return dara.Validate(s)
}
