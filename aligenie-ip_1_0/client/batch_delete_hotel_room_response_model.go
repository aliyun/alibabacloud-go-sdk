// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteHotelRoomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteHotelRoomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteHotelRoomResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteHotelRoomResponseBody) *BatchDeleteHotelRoomResponse
	GetBody() *BatchDeleteHotelRoomResponseBody
}

type BatchDeleteHotelRoomResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteHotelRoomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteHotelRoomResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteHotelRoomResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteHotelRoomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteHotelRoomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteHotelRoomResponse) GetBody() *BatchDeleteHotelRoomResponseBody {
	return s.Body
}

func (s *BatchDeleteHotelRoomResponse) SetHeaders(v map[string]*string) *BatchDeleteHotelRoomResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteHotelRoomResponse) SetStatusCode(v int32) *BatchDeleteHotelRoomResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteHotelRoomResponse) SetBody(v *BatchDeleteHotelRoomResponseBody) *BatchDeleteHotelRoomResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteHotelRoomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
