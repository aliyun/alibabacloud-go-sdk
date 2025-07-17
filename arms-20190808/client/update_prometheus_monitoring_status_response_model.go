// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusMonitoringStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrometheusMonitoringStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrometheusMonitoringStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrometheusMonitoringStatusResponseBody) *UpdatePrometheusMonitoringStatusResponse
	GetBody() *UpdatePrometheusMonitoringStatusResponseBody
}

type UpdatePrometheusMonitoringStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrometheusMonitoringStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrometheusMonitoringStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusMonitoringStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusMonitoringStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrometheusMonitoringStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrometheusMonitoringStatusResponse) GetBody() *UpdatePrometheusMonitoringStatusResponseBody {
	return s.Body
}

func (s *UpdatePrometheusMonitoringStatusResponse) SetHeaders(v map[string]*string) *UpdatePrometheusMonitoringStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusMonitoringStatusResponse) SetStatusCode(v int32) *UpdatePrometheusMonitoringStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusResponse) SetBody(v *UpdatePrometheusMonitoringStatusResponseBody) *UpdatePrometheusMonitoringStatusResponse {
	s.Body = v
	return s
}

func (s *UpdatePrometheusMonitoringStatusResponse) Validate() error {
	return dara.Validate(s)
}
