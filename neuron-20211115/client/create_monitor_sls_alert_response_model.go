// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorSlsAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorSlsAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorSlsAlertResponse
	GetStatusCode() *int32
	SetBody(v *MonitorSlsAlert) *CreateMonitorSlsAlertResponse
	GetBody() *MonitorSlsAlert
}

type CreateMonitorSlsAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorSlsAlert   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorSlsAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorSlsAlertResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorSlsAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorSlsAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorSlsAlertResponse) GetBody() *MonitorSlsAlert {
	return s.Body
}

func (s *CreateMonitorSlsAlertResponse) SetHeaders(v map[string]*string) *CreateMonitorSlsAlertResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorSlsAlertResponse) SetStatusCode(v int32) *CreateMonitorSlsAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorSlsAlertResponse) SetBody(v *MonitorSlsAlert) *CreateMonitorSlsAlertResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorSlsAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
