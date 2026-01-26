// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvServiceMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnvServiceMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnvServiceMonitorResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnvServiceMonitorResponseBody) *CreateEnvServiceMonitorResponse
	GetBody() *CreateEnvServiceMonitorResponseBody
}

type CreateEnvServiceMonitorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnvServiceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnvServiceMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvServiceMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvServiceMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnvServiceMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnvServiceMonitorResponse) GetBody() *CreateEnvServiceMonitorResponseBody {
	return s.Body
}

func (s *CreateEnvServiceMonitorResponse) SetHeaders(v map[string]*string) *CreateEnvServiceMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvServiceMonitorResponse) SetStatusCode(v int32) *CreateEnvServiceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnvServiceMonitorResponse) SetBody(v *CreateEnvServiceMonitorResponseBody) *CreateEnvServiceMonitorResponse {
	s.Body = v
	return s
}

func (s *CreateEnvServiceMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
