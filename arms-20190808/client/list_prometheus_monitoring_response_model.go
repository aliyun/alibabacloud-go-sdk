// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusMonitoringResponseBody) *ListPrometheusMonitoringResponse
	GetBody() *ListPrometheusMonitoringResponseBody
}

type ListPrometheusMonitoringResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusMonitoringResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusMonitoringResponse) GetBody() *ListPrometheusMonitoringResponseBody {
	return s.Body
}

func (s *ListPrometheusMonitoringResponse) SetHeaders(v map[string]*string) *ListPrometheusMonitoringResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusMonitoringResponse) SetStatusCode(v int32) *ListPrometheusMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusMonitoringResponse) SetBody(v *ListPrometheusMonitoringResponseBody) *ListPrometheusMonitoringResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusMonitoringResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
