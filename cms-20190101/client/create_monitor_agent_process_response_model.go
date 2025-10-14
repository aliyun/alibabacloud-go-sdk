// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorAgentProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorAgentProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorAgentProcessResponse
	GetStatusCode() *int32
	SetBody(v *CreateMonitorAgentProcessResponseBody) *CreateMonitorAgentProcessResponse
	GetBody() *CreateMonitorAgentProcessResponseBody
}

type CreateMonitorAgentProcessResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitorAgentProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorAgentProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorAgentProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorAgentProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorAgentProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorAgentProcessResponse) GetBody() *CreateMonitorAgentProcessResponseBody {
	return s.Body
}

func (s *CreateMonitorAgentProcessResponse) SetHeaders(v map[string]*string) *CreateMonitorAgentProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorAgentProcessResponse) SetStatusCode(v int32) *CreateMonitorAgentProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorAgentProcessResponse) SetBody(v *CreateMonitorAgentProcessResponseBody) *CreateMonitorAgentProcessResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorAgentProcessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
