// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelRoomDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelRoomDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelRoomDeviceResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelRoomDeviceResponseBody) *GetHotelRoomDeviceResponse
	GetBody() *GetHotelRoomDeviceResponseBody
}

type GetHotelRoomDeviceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelRoomDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelRoomDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelRoomDeviceResponse) GoString() string {
	return s.String()
}

func (s *GetHotelRoomDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelRoomDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelRoomDeviceResponse) GetBody() *GetHotelRoomDeviceResponseBody {
	return s.Body
}

func (s *GetHotelRoomDeviceResponse) SetHeaders(v map[string]*string) *GetHotelRoomDeviceResponse {
	s.Headers = v
	return s
}

func (s *GetHotelRoomDeviceResponse) SetStatusCode(v int32) *GetHotelRoomDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelRoomDeviceResponse) SetBody(v *GetHotelRoomDeviceResponseBody) *GetHotelRoomDeviceResponse {
	s.Body = v
	return s
}

func (s *GetHotelRoomDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
