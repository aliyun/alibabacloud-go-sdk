// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorMqAlertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorMqAlertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorMqAlertsResponse
	GetStatusCode() *int32
	SetBody(v *MonitorMqAlerts) *CreateMonitorMqAlertsResponse
	GetBody() *MonitorMqAlerts
}

type CreateMonitorMqAlertsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorMqAlerts   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorMqAlertsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorMqAlertsResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorMqAlertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorMqAlertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorMqAlertsResponse) GetBody() *MonitorMqAlerts {
	return s.Body
}

func (s *CreateMonitorMqAlertsResponse) SetHeaders(v map[string]*string) *CreateMonitorMqAlertsResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorMqAlertsResponse) SetStatusCode(v int32) *CreateMonitorMqAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorMqAlertsResponse) SetBody(v *MonitorMqAlerts) *CreateMonitorMqAlertsResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorMqAlertsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
