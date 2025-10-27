// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduleResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduleResponseBody) *CreateScheduleResponse
	GetBody() *CreateScheduleResponseBody
}

type CreateScheduleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduleResponse) GetBody() *CreateScheduleResponseBody {
	return s.Body
}

func (s *CreateScheduleResponse) SetHeaders(v map[string]*string) *CreateScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleResponse) SetStatusCode(v int32) *CreateScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduleResponse) SetBody(v *CreateScheduleResponseBody) *CreateScheduleResponse {
	s.Body = v
	return s
}

func (s *CreateScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
