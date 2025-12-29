// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetHotelContactsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *GetHotelContactsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *GetHotelContactsHeaders
	GetAuthorization() *string
}

type GetHotelContactsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetHotelContactsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactsHeaders) GoString() string {
	return s.String()
}

func (s *GetHotelContactsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetHotelContactsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *GetHotelContactsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetHotelContactsHeaders) SetCommonHeaders(v map[string]*string) *GetHotelContactsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetHotelContactsHeaders) SetXAcsAligenieAccessToken(v string) *GetHotelContactsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *GetHotelContactsHeaders) SetAuthorization(v string) *GetHotelContactsHeaders {
	s.Authorization = &v
	return s
}

func (s *GetHotelContactsHeaders) Validate() error {
	return dara.Validate(s)
}
