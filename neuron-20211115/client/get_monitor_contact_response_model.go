// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonitorContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMonitorContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMonitorContactResponse
	GetStatusCode() *int32
	SetBody(v *MonitorContact) *GetMonitorContactResponse
	GetBody() *MonitorContact
}

type GetMonitorContactResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorContact    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMonitorContactResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMonitorContactResponse) GoString() string {
	return s.String()
}

func (s *GetMonitorContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMonitorContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMonitorContactResponse) GetBody() *MonitorContact {
	return s.Body
}

func (s *GetMonitorContactResponse) SetHeaders(v map[string]*string) *GetMonitorContactResponse {
	s.Headers = v
	return s
}

func (s *GetMonitorContactResponse) SetStatusCode(v int32) *GetMonitorContactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMonitorContactResponse) SetBody(v *MonitorContact) *GetMonitorContactResponse {
	s.Body = v
	return s
}

func (s *GetMonitorContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
