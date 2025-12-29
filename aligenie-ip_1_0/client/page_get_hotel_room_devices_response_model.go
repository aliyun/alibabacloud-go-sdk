// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageGetHotelRoomDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageGetHotelRoomDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageGetHotelRoomDevicesResponse
	GetStatusCode() *int32
	SetBody(v *PageGetHotelRoomDevicesResponseBody) *PageGetHotelRoomDevicesResponse
	GetBody() *PageGetHotelRoomDevicesResponseBody
}

type PageGetHotelRoomDevicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageGetHotelRoomDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageGetHotelRoomDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s PageGetHotelRoomDevicesResponse) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageGetHotelRoomDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageGetHotelRoomDevicesResponse) GetBody() *PageGetHotelRoomDevicesResponseBody {
	return s.Body
}

func (s *PageGetHotelRoomDevicesResponse) SetHeaders(v map[string]*string) *PageGetHotelRoomDevicesResponse {
	s.Headers = v
	return s
}

func (s *PageGetHotelRoomDevicesResponse) SetStatusCode(v int32) *PageGetHotelRoomDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponse) SetBody(v *PageGetHotelRoomDevicesResponseBody) *PageGetHotelRoomDevicesResponse {
	s.Body = v
	return s
}

func (s *PageGetHotelRoomDevicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
