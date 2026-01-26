// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrometheusMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrometheusMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrometheusMonitoringResponseBody) *DeletePrometheusMonitoringResponse
	GetBody() *DeletePrometheusMonitoringResponseBody
}

type DeletePrometheusMonitoringResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrometheusMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrometheusMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusMonitoringResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrometheusMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrometheusMonitoringResponse) GetBody() *DeletePrometheusMonitoringResponseBody {
	return s.Body
}

func (s *DeletePrometheusMonitoringResponse) SetHeaders(v map[string]*string) *DeletePrometheusMonitoringResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusMonitoringResponse) SetStatusCode(v int32) *DeletePrometheusMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusMonitoringResponse) SetBody(v *DeletePrometheusMonitoringResponseBody) *DeletePrometheusMonitoringResponse {
	s.Body = v
	return s
}

func (s *DeletePrometheusMonitoringResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
