// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelRoomDeviceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelRoomDeviceHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelRoomDeviceHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelRoomDeviceHeaders
	GetAuthorization() *string
}

type GetHotelRoomDeviceHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelRoomDeviceHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelRoomDeviceHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelRoomDeviceHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelRoomDeviceHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelRoomDeviceHeaders) SetCommonHeaders(v map[string]*string) *GetHotelRoomDeviceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelRoomDeviceHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelRoomDeviceHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelRoomDeviceHeaders) SetAuthorization(v string) *GetHotelRoomDeviceHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelRoomDeviceHeaders) Validate() error {
	return dara.Validate(s)
}
