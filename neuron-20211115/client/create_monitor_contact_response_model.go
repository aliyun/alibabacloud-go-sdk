// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorContactResponse
	GetStatusCode() *int32
	SetBody(v *MonitorContact) *CreateMonitorContactResponse
	GetBody() *MonitorContact
}

type CreateMonitorContactResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorContact    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorContactResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorContactResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorContactResponse) GetBody() *MonitorContact {
	return s.Body
}

func (s *CreateMonitorContactResponse) SetHeaders(v map[string]*string) *CreateMonitorContactResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorContactResponse) SetStatusCode(v int32) *CreateMonitorContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorContactResponse) SetBody(v *MonitorContact) *CreateMonitorContactResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
