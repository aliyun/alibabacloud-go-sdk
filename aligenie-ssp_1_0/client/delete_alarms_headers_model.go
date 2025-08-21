// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlarmsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteAlarmsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteAlarmsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteAlarmsHeaders
	GetAuthorization() *string
}

type DeleteAlarmsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteAlarmsHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmsHeaders) GoString() string {
	return s.String()
}

func (s *DeleteAlarmsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteAlarmsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteAlarmsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteAlarmsHeaders) SetCommonHeaders(v map[string]*string) *DeleteAlarmsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteAlarmsHeaders) SetXAcsAligenieAccessToken(v string) *DeleteAlarmsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteAlarmsHeaders) SetAuthorization(v string) *DeleteAlarmsHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteAlarmsHeaders) Validate() error {
	return dara.Validate(s)
}
