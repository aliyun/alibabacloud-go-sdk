// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorSlsAlertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorSlsAlertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorSlsAlertsResponse
	GetStatusCode() *int32
	SetBody(v *MonitorSlsAlerts) *CreateMonitorSlsAlertsResponse
	GetBody() *MonitorSlsAlerts
}

type CreateMonitorSlsAlertsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorSlsAlerts  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorSlsAlertsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorSlsAlertsResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorSlsAlertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorSlsAlertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorSlsAlertsResponse) GetBody() *MonitorSlsAlerts {
	return s.Body
}

func (s *CreateMonitorSlsAlertsResponse) SetHeaders(v map[string]*string) *CreateMonitorSlsAlertsResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorSlsAlertsResponse) SetStatusCode(v int32) *CreateMonitorSlsAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorSlsAlertsResponse) SetBody(v *MonitorSlsAlerts) *CreateMonitorSlsAlertsResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorSlsAlertsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
