// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelAlarmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateHotelAlarmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *CreateHotelAlarmHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *CreateHotelAlarmHeaders
	GetAuthorization() *string
}

type CreateHotelAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s CreateHotelAlarmHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmHeaders) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateHotelAlarmHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *CreateHotelAlarmHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *CreateHotelAlarmHeaders) SetCommonHeaders(v map[string]*string) *CreateHotelAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateHotelAlarmHeaders) SetXAcsAligenieAccessToken(v string) *CreateHotelAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *CreateHotelAlarmHeaders) SetAuthorization(v string) *CreateHotelAlarmHeaders {
	s.Authorization = &v
	return s
}

func (s *CreateHotelAlarmHeaders) Validate() error {
	return dara.Validate(s)
}
