// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHotelAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHotelAlarmResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHotelAlarmResponseBody) *UpdateHotelAlarmResponse
	GetBody() *UpdateHotelAlarmResponseBody
}

type UpdateHotelAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotelAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotelAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelAlarmResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotelAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHotelAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHotelAlarmResponse) GetBody() *UpdateHotelAlarmResponseBody {
	return s.Body
}

func (s *UpdateHotelAlarmResponse) SetHeaders(v map[string]*string) *UpdateHotelAlarmResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotelAlarmResponse) SetStatusCode(v int32) *UpdateHotelAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelAlarmResponse) SetBody(v *UpdateHotelAlarmResponseBody) *UpdateHotelAlarmResponse {
	s.Body = v
	return s
}

func (s *UpdateHotelAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
