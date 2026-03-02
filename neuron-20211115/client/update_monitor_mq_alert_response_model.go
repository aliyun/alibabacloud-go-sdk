// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorMqAlertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMonitorMqAlertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMonitorMqAlertResponse
	GetStatusCode() *int32
	SetBody(v *MonitorMqAlert) *UpdateMonitorMqAlertResponse
	GetBody() *MonitorMqAlert
}

type UpdateMonitorMqAlertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorMqAlert    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorMqAlertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorMqAlertResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitorMqAlertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMonitorMqAlertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMonitorMqAlertResponse) GetBody() *MonitorMqAlert {
	return s.Body
}

func (s *UpdateMonitorMqAlertResponse) SetHeaders(v map[string]*string) *UpdateMonitorMqAlertResponse {
	s.Headers = v
	return s
}

func (s *UpdateMonitorMqAlertResponse) SetStatusCode(v int32) *UpdateMonitorMqAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMonitorMqAlertResponse) SetBody(v *MonitorMqAlert) *UpdateMonitorMqAlertResponse {
	s.Body = v
	return s
}

func (s *UpdateMonitorMqAlertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
