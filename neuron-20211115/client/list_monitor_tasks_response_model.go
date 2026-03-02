// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMonitorTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMonitorTasksResponse
	GetStatusCode() *int32
	SetBody(v *MonitorNotifyAlertPageResult) *ListMonitorTasksResponse
	GetBody() *MonitorNotifyAlertPageResult
}

type ListMonitorTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorNotifyAlertPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMonitorTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMonitorTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMonitorTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMonitorTasksResponse) GetBody() *MonitorNotifyAlertPageResult {
	return s.Body
}

func (s *ListMonitorTasksResponse) SetHeaders(v map[string]*string) *ListMonitorTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMonitorTasksResponse) SetStatusCode(v int32) *ListMonitorTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMonitorTasksResponse) SetBody(v *MonitorNotifyAlertPageResult) *ListMonitorTasksResponse {
	s.Body = v
	return s
}

func (s *ListMonitorTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
