// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScheduleMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScheduleMonitorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScheduleMonitorResponseBody) *UpdateScheduleMonitorResponse
	GetBody() *UpdateScheduleMonitorResponseBody
}

type UpdateScheduleMonitorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduleMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduleMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScheduleMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScheduleMonitorResponse) GetBody() *UpdateScheduleMonitorResponseBody {
	return s.Body
}

func (s *UpdateScheduleMonitorResponse) SetHeaders(v map[string]*string) *UpdateScheduleMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleMonitorResponse) SetStatusCode(v int32) *UpdateScheduleMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduleMonitorResponse) SetBody(v *UpdateScheduleMonitorResponseBody) *UpdateScheduleMonitorResponse {
	s.Body = v
	return s
}

func (s *UpdateScheduleMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
