// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduleMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduleMonitorResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduleMonitorResponseBody) *CreateScheduleMonitorResponse
	GetBody() *CreateScheduleMonitorResponseBody
}

type CreateScheduleMonitorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduleMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduleMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduleMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduleMonitorResponse) GetBody() *CreateScheduleMonitorResponseBody {
	return s.Body
}

func (s *CreateScheduleMonitorResponse) SetHeaders(v map[string]*string) *CreateScheduleMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleMonitorResponse) SetStatusCode(v int32) *CreateScheduleMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduleMonitorResponse) SetBody(v *CreateScheduleMonitorResponseBody) *CreateScheduleMonitorResponse {
	s.Body = v
	return s
}

func (s *CreateScheduleMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
