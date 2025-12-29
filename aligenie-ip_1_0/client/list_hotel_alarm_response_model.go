// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelAlarmResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelAlarmResponseBody) *ListHotelAlarmResponse
	GetBody() *ListHotelAlarmResponseBody
}

type ListHotelAlarmResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelAlarmResponse) GoString() string {
	return s.String()
}

func (s *ListHotelAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelAlarmResponse) GetBody() *ListHotelAlarmResponseBody {
	return s.Body
}

func (s *ListHotelAlarmResponse) SetHeaders(v map[string]*string) *ListHotelAlarmResponse {
	s.Headers = v
	return s
}

func (s *ListHotelAlarmResponse) SetStatusCode(v int32) *ListHotelAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelAlarmResponse) SetBody(v *ListHotelAlarmResponseBody) *ListHotelAlarmResponse {
	s.Body = v
	return s
}

func (s *ListHotelAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
