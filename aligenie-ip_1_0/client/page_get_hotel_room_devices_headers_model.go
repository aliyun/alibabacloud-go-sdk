// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageGetHotelRoomDevicesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PageGetHotelRoomDevicesHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *PageGetHotelRoomDevicesHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *PageGetHotelRoomDevicesHeaders
	GetAuthorization() *string
}

type PageGetHotelRoomDevicesHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PageGetHotelRoomDevicesHeaders) String() string {
	return dara.Prettify(s)
}

func (s PageGetHotelRoomDevicesHeaders) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PageGetHotelRoomDevicesHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *PageGetHotelRoomDevicesHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PageGetHotelRoomDevicesHeaders) SetCommonHeaders(v map[string]*string) *PageGetHotelRoomDevicesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PageGetHotelRoomDevicesHeaders) SetXAcsAligenieAccessToken(v string) *PageGetHotelRoomDevicesHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *PageGetHotelRoomDevicesHeaders) SetAuthorization(v string) *PageGetHotelRoomDevicesHeaders {
	s.Authorization = &v
	return s
}

func (s *PageGetHotelRoomDevicesHeaders) Validate() error {
	return dara.Validate(s)
}
