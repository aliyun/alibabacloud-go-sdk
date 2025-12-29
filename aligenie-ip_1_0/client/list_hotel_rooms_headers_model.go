// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelRoomsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListHotelRoomsHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListHotelRoomsHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListHotelRoomsHeaders
	GetAuthorization() *string
}

type ListHotelRoomsHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListHotelRoomsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListHotelRoomsHeaders) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListHotelRoomsHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListHotelRoomsHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListHotelRoomsHeaders) SetCommonHeaders(v map[string]*string) *ListHotelRoomsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListHotelRoomsHeaders) SetXAcsAligenieAccessToken(v string) *ListHotelRoomsHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListHotelRoomsHeaders) SetAuthorization(v string) *ListHotelRoomsHeaders {
	s.Authorization = &v
	return s
}

func (s *ListHotelRoomsHeaders) Validate() error {
	return dara.Validate(s)
}
