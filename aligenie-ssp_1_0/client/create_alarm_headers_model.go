// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlarmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateAlarmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CreateAlarmHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CreateAlarmHeaders
	GetAuthorization() *string
}

type CreateAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateAlarmHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateAlarmHeaders) GoString() string {
	return s.String()
}

func (s *CreateAlarmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateAlarmHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CreateAlarmHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateAlarmHeaders) SetCommonHeaders(v map[string]*string) *CreateAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateAlarmHeaders) SetXAcsAligenieAccessToken(v string) *CreateAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateAlarmHeaders) SetAuthorization(v string) *CreateAlarmHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateAlarmHeaders) Validate() error {
	return dara.Validate(s)
}
