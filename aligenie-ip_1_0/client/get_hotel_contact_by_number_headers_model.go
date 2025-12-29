// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByNumberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelContactByNumberHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelContactByNumberHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelContactByNumberHeaders
	GetAuthorization() *string
}

type GetHotelContactByNumberHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelContactByNumberHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByNumberHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelContactByNumberHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelContactByNumberHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelContactByNumberHeaders) SetCommonHeaders(v map[string]*string) *GetHotelContactByNumberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelContactByNumberHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelContactByNumberHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelContactByNumberHeaders) SetAuthorization(v string) *GetHotelContactByNumberHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelContactByNumberHeaders) Validate() error {
	return dara.Validate(s)
}
