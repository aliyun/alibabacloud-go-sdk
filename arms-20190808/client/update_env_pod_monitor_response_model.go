// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvPodMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnvPodMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnvPodMonitorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnvPodMonitorResponseBody) *UpdateEnvPodMonitorResponse
	GetBody() *UpdateEnvPodMonitorResponseBody
}

type UpdateEnvPodMonitorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnvPodMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvPodMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvPodMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvPodMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnvPodMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnvPodMonitorResponse) GetBody() *UpdateEnvPodMonitorResponseBody {
	return s.Body
}

func (s *UpdateEnvPodMonitorResponse) SetHeaders(v map[string]*string) *UpdateEnvPodMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvPodMonitorResponse) SetStatusCode(v int32) *UpdateEnvPodMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvPodMonitorResponse) SetBody(v *UpdateEnvPodMonitorResponseBody) *UpdateEnvPodMonitorResponse {
	s.Body = v
	return s
}

func (s *UpdateEnvPodMonitorResponse) Validate() error {
	return dara.Validate(s)
}
