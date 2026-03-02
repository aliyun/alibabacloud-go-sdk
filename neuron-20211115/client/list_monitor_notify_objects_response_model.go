// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMonitorNotifyObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMonitorNotifyObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMonitorNotifyObjectsResponse
	GetStatusCode() *int32
	SetBody(v *MonitorNotifyObjectResult) *ListMonitorNotifyObjectsResponse
	GetBody() *MonitorNotifyObjectResult
}

type ListMonitorNotifyObjectsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorNotifyObjectResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMonitorNotifyObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMonitorNotifyObjectsResponse) GoString() string {
	return s.String()
}

func (s *ListMonitorNotifyObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMonitorNotifyObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMonitorNotifyObjectsResponse) GetBody() *MonitorNotifyObjectResult {
	return s.Body
}

func (s *ListMonitorNotifyObjectsResponse) SetHeaders(v map[string]*string) *ListMonitorNotifyObjectsResponse {
	s.Headers = v
	return s
}

func (s *ListMonitorNotifyObjectsResponse) SetStatusCode(v int32) *ListMonitorNotifyObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMonitorNotifyObjectsResponse) SetBody(v *MonitorNotifyObjectResult) *ListMonitorNotifyObjectsResponse {
	s.Body = v
	return s
}

func (s *ListMonitorNotifyObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
