// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHotelAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHotelAlarmResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHotelAlarmResponseBody) *DeleteHotelAlarmResponse
	GetBody() *DeleteHotelAlarmResponseBody
}

type DeleteHotelAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHotelAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHotelAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelAlarmResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotelAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHotelAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHotelAlarmResponse) GetBody() *DeleteHotelAlarmResponseBody {
	return s.Body
}

func (s *DeleteHotelAlarmResponse) SetHeaders(v map[string]*string) *DeleteHotelAlarmResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotelAlarmResponse) SetStatusCode(v int32) *DeleteHotelAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotelAlarmResponse) SetBody(v *DeleteHotelAlarmResponseBody) *DeleteHotelAlarmResponse {
	s.Body = v
	return s
}

func (s *DeleteHotelAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
