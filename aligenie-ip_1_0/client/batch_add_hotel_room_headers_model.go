// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddHotelRoomHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchAddHotelRoomHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *BatchAddHotelRoomHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *BatchAddHotelRoomHeaders
	GetAuthorization() *string
}

type BatchAddHotelRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s BatchAddHotelRoomHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchAddHotelRoomHeaders) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchAddHotelRoomHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *BatchAddHotelRoomHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *BatchAddHotelRoomHeaders) SetCommonHeaders(v map[string]*string) *BatchAddHotelRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchAddHotelRoomHeaders) SetXAcsAligenieAccessToken(v string) *BatchAddHotelRoomHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *BatchAddHotelRoomHeaders) SetAuthorization(v string) *BatchAddHotelRoomHeaders {
	s.Authorization = &v
	return s
}

func (s *BatchAddHotelRoomHeaders) Validate() error {
	return dara.Validate(s)
}
