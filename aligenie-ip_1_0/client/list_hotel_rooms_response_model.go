// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelRoomsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelRoomsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelRoomsResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelRoomsResponseBody) *ListHotelRoomsResponse
	GetBody() *ListHotelRoomsResponseBody
}

type ListHotelRoomsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelRoomsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListHotelRoomsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelRoomsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelRoomsResponse) GetBody() *ListHotelRoomsResponseBody {
	return s.Body
}

func (s *ListHotelRoomsResponse) SetHeaders(v map[string]*string) *ListHotelRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListHotelRoomsResponse) SetStatusCode(v int32) *ListHotelRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelRoomsResponse) SetBody(v *ListHotelRoomsResponseBody) *ListHotelRoomsResponse {
	s.Body = v
	return s
}

func (s *ListHotelRoomsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
