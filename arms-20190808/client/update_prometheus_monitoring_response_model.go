// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrometheusMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrometheusMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrometheusMonitoringResponseBody) *UpdatePrometheusMonitoringResponse
	GetBody() *UpdatePrometheusMonitoringResponseBody
}

type UpdatePrometheusMonitoringResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrometheusMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrometheusMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusMonitoringResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrometheusMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrometheusMonitoringResponse) GetBody() *UpdatePrometheusMonitoringResponseBody {
	return s.Body
}

func (s *UpdatePrometheusMonitoringResponse) SetHeaders(v map[string]*string) *UpdatePrometheusMonitoringResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusMonitoringResponse) SetStatusCode(v int32) *UpdatePrometheusMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusMonitoringResponse) SetBody(v *UpdatePrometheusMonitoringResponseBody) *UpdatePrometheusMonitoringResponse {
	s.Body = v
	return s
}

func (s *UpdatePrometheusMonitoringResponse) Validate() error {
	return dara.Validate(s)
}
