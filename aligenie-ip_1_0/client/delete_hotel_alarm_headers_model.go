// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelAlarmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteHotelAlarmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *DeleteHotelAlarmHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *DeleteHotelAlarmHeaders
	GetAuthorization() *string
}

type DeleteHotelAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s DeleteHotelAlarmHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelAlarmHeaders) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteHotelAlarmHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *DeleteHotelAlarmHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *DeleteHotelAlarmHeaders) SetCommonHeaders(v map[string]*string) *DeleteHotelAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteHotelAlarmHeaders) SetXAcsAligenieAccessToken(v string) *DeleteHotelAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *DeleteHotelAlarmHeaders) SetAuthorization(v string) *DeleteHotelAlarmHeaders {
	s.Authorization = &v
	return s
}

func (s *DeleteHotelAlarmHeaders) Validate() error {
	return dara.Validate(s)
}
