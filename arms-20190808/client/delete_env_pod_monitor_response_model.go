// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvPodMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnvPodMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnvPodMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEnvPodMonitorResponseBody) *DeleteEnvPodMonitorResponse
	GetBody() *DeleteEnvPodMonitorResponseBody
}

type DeleteEnvPodMonitorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEnvPodMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEnvPodMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvPodMonitorResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvPodMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnvPodMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnvPodMonitorResponse) GetBody() *DeleteEnvPodMonitorResponseBody {
	return s.Body
}

func (s *DeleteEnvPodMonitorResponse) SetHeaders(v map[string]*string) *DeleteEnvPodMonitorResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvPodMonitorResponse) SetStatusCode(v int32) *DeleteEnvPodMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnvPodMonitorResponse) SetBody(v *DeleteEnvPodMonitorResponseBody) *DeleteEnvPodMonitorResponse {
	s.Body = v
	return s
}

func (s *DeleteEnvPodMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
