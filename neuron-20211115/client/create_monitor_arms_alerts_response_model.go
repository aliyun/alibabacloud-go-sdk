// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorArmsAlertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorArmsAlertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorArmsAlertsResponse
	GetStatusCode() *int32
	SetBody(v *MonitorArmsAlerts) *CreateMonitorArmsAlertsResponse
	GetBody() *MonitorArmsAlerts
}

type CreateMonitorArmsAlertsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorArmsAlerts `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorArmsAlertsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorArmsAlertsResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorArmsAlertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorArmsAlertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorArmsAlertsResponse) GetBody() *MonitorArmsAlerts {
	return s.Body
}

func (s *CreateMonitorArmsAlertsResponse) SetHeaders(v map[string]*string) *CreateMonitorArmsAlertsResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorArmsAlertsResponse) SetStatusCode(v int32) *CreateMonitorArmsAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorArmsAlertsResponse) SetBody(v *MonitorArmsAlerts) *CreateMonitorArmsAlertsResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorArmsAlertsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
