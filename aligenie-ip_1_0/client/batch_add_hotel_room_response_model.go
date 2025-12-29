// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddHotelRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAddHotelRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAddHotelRoomResponse
	GetStatusCode() *int32
	SetBody(v *BatchAddHotelRoomResponseBody) *BatchAddHotelRoomResponse
	GetBody() *BatchAddHotelRoomResponseBody
}

type BatchAddHotelRoomResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddHotelRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddHotelRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAddHotelRoomResponse) GoString() string {
	return s.String()
}

func (s *BatchAddHotelRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAddHotelRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAddHotelRoomResponse) GetBody() *BatchAddHotelRoomResponseBody {
	return s.Body
}

func (s *BatchAddHotelRoomResponse) SetHeaders(v map[string]*string) *BatchAddHotelRoomResponse {
	s.Headers = v
	return s
}

func (s *BatchAddHotelRoomResponse) SetStatusCode(v int32) *BatchAddHotelRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddHotelRoomResponse) SetBody(v *BatchAddHotelRoomResponseBody) *BatchAddHotelRoomResponse {
	s.Body = v
	return s
}

func (s *BatchAddHotelRoomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
