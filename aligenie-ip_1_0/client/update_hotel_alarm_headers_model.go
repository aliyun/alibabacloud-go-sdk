// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelAlarmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateHotelAlarmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *UpdateHotelAlarmHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *UpdateHotelAlarmHeaders
	GetAuthorization() *string
}

type UpdateHotelAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s UpdateHotelAlarmHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmHeaders) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateHotelAlarmHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *UpdateHotelAlarmHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *UpdateHotelAlarmHeaders) SetCommonHeaders(v map[string]*string) *UpdateHotelAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateHotelAlarmHeaders) SetXAcsAligenieAccessToken(v string) *UpdateHotelAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *UpdateHotelAlarmHeaders) SetAuthorization(v string) *UpdateHotelAlarmHeaders {
	s.Authorization = &v
	return s
}

func (s *UpdateHotelAlarmHeaders) Validate() error {
	return dara.Validate(s)
}
