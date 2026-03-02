// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorArmsAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMonitorArmsAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMonitorArmsAlertResponse
	GetStatusCode() *int32
	SetBody(v *MonitorArmsAlert) *UpdateMonitorArmsAlertResponse
	GetBody() *MonitorArmsAlert
}

type UpdateMonitorArmsAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorArmsAlert  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorArmsAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorArmsAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitorArmsAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMonitorArmsAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMonitorArmsAlertResponse) GetBody() *MonitorArmsAlert {
	return s.Body
}

func (s *UpdateMonitorArmsAlertResponse) SetHeaders(v map[string]*string) *UpdateMonitorArmsAlertResponse {
	s.Headers = v
	return s
}

func (s *UpdateMonitorArmsAlertResponse) SetStatusCode(v int32) *UpdateMonitorArmsAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMonitorArmsAlertResponse) SetBody(v *MonitorArmsAlert) *UpdateMonitorArmsAlertResponse {
	s.Body = v
	return s
}

func (s *UpdateMonitorArmsAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
