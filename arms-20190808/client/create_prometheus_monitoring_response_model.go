// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrometheusMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrometheusMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrometheusMonitoringResponseBody) *CreatePrometheusMonitoringResponse
	GetBody() *CreatePrometheusMonitoringResponseBody
}

type CreatePrometheusMonitoringResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrometheusMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrometheusMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusMonitoringResponse) GoString() string {
	return s.String()
}

func (s *CreatePrometheusMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrometheusMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrometheusMonitoringResponse) GetBody() *CreatePrometheusMonitoringResponseBody {
	return s.Body
}

func (s *CreatePrometheusMonitoringResponse) SetHeaders(v map[string]*string) *CreatePrometheusMonitoringResponse {
	s.Headers = v
	return s
}

func (s *CreatePrometheusMonitoringResponse) SetStatusCode(v int32) *CreatePrometheusMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrometheusMonitoringResponse) SetBody(v *CreatePrometheusMonitoringResponseBody) *CreatePrometheusMonitoringResponse {
	s.Body = v
	return s
}

func (s *CreatePrometheusMonitoringResponse) Validate() error {
	return dara.Validate(s)
}
