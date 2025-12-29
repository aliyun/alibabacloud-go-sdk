// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverStyleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelScreenSaverStyleHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelScreenSaverStyleHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelScreenSaverStyleHeaders
	GetAuthorization() *string
}

type GetHotelScreenSaverStyleHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelScreenSaverStyleHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverStyleHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelScreenSaverStyleHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelScreenSaverStyleHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelScreenSaverStyleHeaders) SetCommonHeaders(v map[string]*string) *GetHotelScreenSaverStyleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelScreenSaverStyleHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelScreenSaverStyleHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelScreenSaverStyleHeaders) SetAuthorization(v string) *GetHotelScreenSaverStyleHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelScreenSaverStyleHeaders) Validate() error {
	return dara.Validate(s)
}
