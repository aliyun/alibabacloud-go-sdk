// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScheduleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScheduleResponseBody) *UpdateScheduleResponse
	GetBody() *UpdateScheduleResponseBody
}

type UpdateScheduleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScheduleResponse) GetBody() *UpdateScheduleResponseBody {
	return s.Body
}

func (s *UpdateScheduleResponse) SetHeaders(v map[string]*string) *UpdateScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleResponse) SetStatusCode(v int32) *UpdateScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduleResponse) SetBody(v *UpdateScheduleResponseBody) *UpdateScheduleResponse {
	s.Body = v
	return s
}

func (s *UpdateScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
