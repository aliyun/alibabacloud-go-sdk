// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelAlarmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHotelAlarmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHotelAlarmResponse
	GetStatusCode() *int32
	SetBody(v *CreateHotelAlarmResponseBody) *CreateHotelAlarmResponse
	GetBody() *CreateHotelAlarmResponseBody
}

type CreateHotelAlarmResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHotelAlarmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHotelAlarmResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelAlarmResponse) GoString() string {
	return s.String()
}

func (s *CreateHotelAlarmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHotelAlarmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHotelAlarmResponse) GetBody() *CreateHotelAlarmResponseBody {
	return s.Body
}

func (s *CreateHotelAlarmResponse) SetHeaders(v map[string]*string) *CreateHotelAlarmResponse {
	s.Headers = v
	return s
}

func (s *CreateHotelAlarmResponse) SetStatusCode(v int32) *CreateHotelAlarmResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHotelAlarmResponse) SetBody(v *CreateHotelAlarmResponseBody) *CreateHotelAlarmResponse {
	s.Body = v
	return s
}

func (s *CreateHotelAlarmResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
