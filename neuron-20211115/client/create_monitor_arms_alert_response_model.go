// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorArmsAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorArmsAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorArmsAlertResponse
	GetStatusCode() *int32
	SetBody(v *MonitorArmsAlert) *CreateMonitorArmsAlertResponse
	GetBody() *MonitorArmsAlert
}

type CreateMonitorArmsAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorArmsAlert  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorArmsAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorArmsAlertResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorArmsAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorArmsAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorArmsAlertResponse) GetBody() *MonitorArmsAlert {
	return s.Body
}

func (s *CreateMonitorArmsAlertResponse) SetHeaders(v map[string]*string) *CreateMonitorArmsAlertResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorArmsAlertResponse) SetStatusCode(v int32) *CreateMonitorArmsAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorArmsAlertResponse) SetBody(v *MonitorArmsAlert) *CreateMonitorArmsAlertResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorArmsAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
