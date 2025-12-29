// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteHotelRoomHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchDeleteHotelRoomHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *BatchDeleteHotelRoomHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *BatchDeleteHotelRoomHeaders
	GetAuthorization() *string
}

type BatchDeleteHotelRoomHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s BatchDeleteHotelRoomHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteHotelRoomHeaders) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchDeleteHotelRoomHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *BatchDeleteHotelRoomHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *BatchDeleteHotelRoomHeaders) SetCommonHeaders(v map[string]*string) *BatchDeleteHotelRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchDeleteHotelRoomHeaders) SetXAcsAligenieAccessToken(v string) *BatchDeleteHotelRoomHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *BatchDeleteHotelRoomHeaders) SetAuthorization(v string) *BatchDeleteHotelRoomHeaders {
	s.Authorization = &v
	return s
}

func (s *BatchDeleteHotelRoomHeaders) Validate() error {
	return dara.Validate(s)
}
