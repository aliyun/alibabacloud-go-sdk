// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelControlDeviceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelControlDeviceHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelControlDeviceHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelControlDeviceHeaders
	GetAuthorization() *string
}

type ListHotelControlDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelControlDeviceHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelControlDeviceHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelControlDeviceHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelControlDeviceHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelControlDeviceHeaders) SetCommonHeaders(v map[string]*string) *ListHotelControlDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelControlDeviceHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelControlDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelControlDeviceHeaders) SetAuthorization(v string) *ListHotelControlDeviceHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelControlDeviceHeaders) Validate() error {
	return dara.Validate(s)
}
