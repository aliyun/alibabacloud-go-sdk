// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListAlarmsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListAlarmsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListAlarmsHeaders
	GetAuthorization() *string
}

type ListAlarmsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListAlarmsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmsHeaders) GoString() string {
	return s.String()
}

func (s *ListAlarmsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListAlarmsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListAlarmsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListAlarmsHeaders) SetCommonHeaders(v map[string]*string) *ListAlarmsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListAlarmsHeaders) SetXAcsAligenieAccessToken(v string) *ListAlarmsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListAlarmsHeaders) SetAuthorization(v string) *ListAlarmsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListAlarmsHeaders) Validate() error {
	return dara.Validate(s)
}
