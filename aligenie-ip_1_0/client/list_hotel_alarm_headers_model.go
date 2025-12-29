// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelAlarmHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelAlarmHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelAlarmHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelAlarmHeaders
	GetAuthorization() *string
}

type ListHotelAlarmHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelAlarmHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelAlarmHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelAlarmHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelAlarmHeaders) SetCommonHeaders(v map[string]*string) *ListHotelAlarmHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelAlarmHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelAlarmHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelAlarmHeaders) SetAuthorization(v string) *ListHotelAlarmHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelAlarmHeaders) Validate() error {
	return dara.Validate(s)
}
