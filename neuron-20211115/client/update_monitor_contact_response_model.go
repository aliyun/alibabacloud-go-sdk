// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMonitorContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMonitorContactResponse
	GetStatusCode() *int32
	SetBody(v *MonitorContact) *UpdateMonitorContactResponse
	GetBody() *MonitorContact
}

type UpdateMonitorContactResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorContact    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorContactResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateMonitorContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMonitorContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMonitorContactResponse) GetBody() *MonitorContact {
	return s.Body
}

func (s *UpdateMonitorContactResponse) SetHeaders(v map[string]*string) *UpdateMonitorContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateMonitorContactResponse) SetStatusCode(v int32) *UpdateMonitorContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMonitorContactResponse) SetBody(v *MonitorContact) *UpdateMonitorContactResponse {
	s.Body = v
	return s
}

func (s *UpdateMonitorContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
