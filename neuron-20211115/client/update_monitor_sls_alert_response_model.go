// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorSlsAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMonitorSlsAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMonitorSlsAlertResponse
	GetStatusCode() *int32
	SetBody(v *MonitorSlsAlert) *UpdateMonitorSlsAlertResponse
	GetBody() *MonitorSlsAlert
}

type UpdateMonitorSlsAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorSlsAlert   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorSlsAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorSlsAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitorSlsAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMonitorSlsAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMonitorSlsAlertResponse) GetBody() *MonitorSlsAlert {
	return s.Body
}

func (s *UpdateMonitorSlsAlertResponse) SetHeaders(v map[string]*string) *UpdateMonitorSlsAlertResponse {
	s.Headers = v
	return s
}

func (s *UpdateMonitorSlsAlertResponse) SetStatusCode(v int32) *UpdateMonitorSlsAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMonitorSlsAlertResponse) SetBody(v *MonitorSlsAlert) *UpdateMonitorSlsAlertResponse {
	s.Body = v
	return s
}

func (s *UpdateMonitorSlsAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
