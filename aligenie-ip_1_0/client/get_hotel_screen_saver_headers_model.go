// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelScreenSaverHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelScreenSaverHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelScreenSaverHeaders
	GetAuthorization() *string
}

type GetHotelScreenSaverHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelScreenSaverHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelScreenSaverHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelScreenSaverHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelScreenSaverHeaders) SetCommonHeaders(v map[string]*string) *GetHotelScreenSaverHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelScreenSaverHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelScreenSaverHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelScreenSaverHeaders) SetAuthorization(v string) *GetHotelScreenSaverHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelScreenSaverHeaders) Validate() error {
	return dara.Validate(s)
}
