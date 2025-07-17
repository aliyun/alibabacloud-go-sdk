// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvServiceMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnvServiceMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnvServiceMonitorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnvServiceMonitorResponseBody) *UpdateEnvServiceMonitorResponse
	GetBody() *UpdateEnvServiceMonitorResponseBody
}

type UpdateEnvServiceMonitorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnvServiceMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvServiceMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvServiceMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvServiceMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnvServiceMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnvServiceMonitorResponse) GetBody() *UpdateEnvServiceMonitorResponseBody {
	return s.Body
}

func (s *UpdateEnvServiceMonitorResponse) SetHeaders(v map[string]*string) *UpdateEnvServiceMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvServiceMonitorResponse) SetStatusCode(v int32) *UpdateEnvServiceMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvServiceMonitorResponse) SetBody(v *UpdateEnvServiceMonitorResponseBody) *UpdateEnvServiceMonitorResponse {
	s.Body = v
	return s
}

func (s *UpdateEnvServiceMonitorResponse) Validate() error {
	return dara.Validate(s)
}
