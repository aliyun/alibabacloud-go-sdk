// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotelRoomDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryHotelRoomDetailHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *QueryHotelRoomDetailHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *QueryHotelRoomDetailHeaders
	GetAuthorization() *string
}

type QueryHotelRoomDetailHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s QueryHotelRoomDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryHotelRoomDetailHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *QueryHotelRoomDetailHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *QueryHotelRoomDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryHotelRoomDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryHotelRoomDetailHeaders) SetXAcsAligenieAccessToken(v string) *QueryHotelRoomDetailHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *QueryHotelRoomDetailHeaders) SetAuthorization(v string) *QueryHotelRoomDetailHeaders {
	s.Authorization = &v
	return s
}

func (s *QueryHotelRoomDetailHeaders) Validate() error {
	return dara.Validate(s)
}
