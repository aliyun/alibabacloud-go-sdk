// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlarmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetAlarmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetAlarmHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetAlarmHeaders
	GetAuthorization() *string
}

type GetAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetAlarmHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetAlarmHeaders) GoString() string {
	return s.String()
}

func (s *GetAlarmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetAlarmHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetAlarmHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetAlarmHeaders) SetCommonHeaders(v map[string]*string) *GetAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetAlarmHeaders) SetXAcsAligenieAccessToken(v string) *GetAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetAlarmHeaders) SetAuthorization(v string) *GetAlarmHeaders {
	s.Authorization = &v
	return s
}

func (s *GetAlarmHeaders) Validate() error {
	return dara.Validate(s)
}
