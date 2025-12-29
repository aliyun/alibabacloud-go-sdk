// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelQrBindHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *HotelQrBindHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *HotelQrBindHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *HotelQrBindHeaders
	GetAuthorization() *string
}

type HotelQrBindHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s HotelQrBindHeaders) String() string {
	return dara.Prettify(s)
}

func (s HotelQrBindHeaders) GoString() string {
	return s.String()
}

func (s *HotelQrBindHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *HotelQrBindHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *HotelQrBindHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *HotelQrBindHeaders) SetCommonHeaders(v map[string]*string) *HotelQrBindHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HotelQrBindHeaders) SetXAcsAligenieAccessToken(v string) *HotelQrBindHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *HotelQrBindHeaders) SetAuthorization(v string) *HotelQrBindHeaders {
	s.Authorization = &v
	return s
}

func (s *HotelQrBindHeaders) Validate() error {
	return dara.Validate(s)
}
