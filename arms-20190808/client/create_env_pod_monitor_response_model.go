// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvPodMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnvPodMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnvPodMonitorResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnvPodMonitorResponseBody) *CreateEnvPodMonitorResponse
	GetBody() *CreateEnvPodMonitorResponseBody
}

type CreateEnvPodMonitorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnvPodMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnvPodMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvPodMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvPodMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnvPodMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnvPodMonitorResponse) GetBody() *CreateEnvPodMonitorResponseBody {
	return s.Body
}

func (s *CreateEnvPodMonitorResponse) SetHeaders(v map[string]*string) *CreateEnvPodMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvPodMonitorResponse) SetStatusCode(v int32) *CreateEnvPodMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnvPodMonitorResponse) SetBody(v *CreateEnvPodMonitorResponseBody) *CreateEnvPodMonitorResponse {
	s.Body = v
	return s
}

func (s *CreateEnvPodMonitorResponse) Validate() error {
	return dara.Validate(s)
}
