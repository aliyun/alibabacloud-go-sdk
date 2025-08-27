// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelRoomInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelRoomInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelRoomInfoResponse
	GetStatusCode() *int32
	SetBody(v *HotelRoomInfoResponseBody) *HotelRoomInfoResponse
	GetBody() *HotelRoomInfoResponseBody
}

type HotelRoomInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelRoomInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelRoomInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelRoomInfoResponse) GoString() string {
	return s.String()
}

func (s *HotelRoomInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelRoomInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelRoomInfoResponse) GetBody() *HotelRoomInfoResponseBody {
	return s.Body
}

func (s *HotelRoomInfoResponse) SetHeaders(v map[string]*string) *HotelRoomInfoResponse {
	s.Headers = v
	return s
}

func (s *HotelRoomInfoResponse) SetStatusCode(v int32) *HotelRoomInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelRoomInfoResponse) SetBody(v *HotelRoomInfoResponseBody) *HotelRoomInfoResponse {
	s.Body = v
	return s
}

func (s *HotelRoomInfoResponse) Validate() error {
	return dara.Validate(s)
}
