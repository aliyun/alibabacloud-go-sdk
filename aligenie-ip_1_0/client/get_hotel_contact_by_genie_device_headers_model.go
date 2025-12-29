// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByGenieDeviceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelContactByGenieDeviceHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelContactByGenieDeviceHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelContactByGenieDeviceHeaders
	GetAuthorization() *string
}

type GetHotelContactByGenieDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelContactByGenieDeviceHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByGenieDeviceHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelContactByGenieDeviceHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelContactByGenieDeviceHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelContactByGenieDeviceHeaders) SetCommonHeaders(v map[string]*string) *GetHotelContactByGenieDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelContactByGenieDeviceHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelContactByGenieDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelContactByGenieDeviceHeaders) SetAuthorization(v string) *GetHotelContactByGenieDeviceHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelContactByGenieDeviceHeaders) Validate() error {
	return dara.Validate(s)
}
