// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorMqAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorMqAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorMqAlertResponse
	GetStatusCode() *int32
	SetBody(v *MonitorMqAlert) *CreateMonitorMqAlertResponse
	GetBody() *MonitorMqAlert
}

type CreateMonitorMqAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorMqAlert    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorMqAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorMqAlertResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorMqAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorMqAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorMqAlertResponse) GetBody() *MonitorMqAlert {
	return s.Body
}

func (s *CreateMonitorMqAlertResponse) SetHeaders(v map[string]*string) *CreateMonitorMqAlertResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorMqAlertResponse) SetStatusCode(v int32) *CreateMonitorMqAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorMqAlertResponse) SetBody(v *MonitorMqAlert) *CreateMonitorMqAlertResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorMqAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
