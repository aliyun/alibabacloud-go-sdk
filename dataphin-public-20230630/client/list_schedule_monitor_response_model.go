// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduleMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduleMonitorResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduleMonitorResponseBody) *ListScheduleMonitorResponse
	GetBody() *ListScheduleMonitorResponseBody
}

type ListScheduleMonitorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduleMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduleMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleMonitorResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduleMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduleMonitorResponse) GetBody() *ListScheduleMonitorResponseBody {
	return s.Body
}

func (s *ListScheduleMonitorResponse) SetHeaders(v map[string]*string) *ListScheduleMonitorResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleMonitorResponse) SetStatusCode(v int32) *ListScheduleMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduleMonitorResponse) SetBody(v *ListScheduleMonitorResponseBody) *ListScheduleMonitorResponse {
	s.Body = v
	return s
}

func (s *ListScheduleMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
