// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotelRoomDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHotelRoomDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHotelRoomDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryHotelRoomDetailResponseBody) *QueryHotelRoomDetailResponse
	GetBody() *QueryHotelRoomDetailResponseBody
}

type QueryHotelRoomDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHotelRoomDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHotelRoomDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHotelRoomDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryHotelRoomDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHotelRoomDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHotelRoomDetailResponse) GetBody() *QueryHotelRoomDetailResponseBody {
	return s.Body
}

func (s *QueryHotelRoomDetailResponse) SetHeaders(v map[string]*string) *QueryHotelRoomDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryHotelRoomDetailResponse) SetStatusCode(v int32) *QueryHotelRoomDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotelRoomDetailResponse) SetBody(v *QueryHotelRoomDetailResponseBody) *QueryHotelRoomDetailResponse {
	s.Body = v
	return s
}

func (s *QueryHotelRoomDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
