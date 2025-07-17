// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusMonitoringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrometheusMonitoringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrometheusMonitoringResponse
	GetStatusCode() *int32
	SetBody(v *GetPrometheusMonitoringResponseBody) *GetPrometheusMonitoringResponse
	GetBody() *GetPrometheusMonitoringResponseBody
}

type GetPrometheusMonitoringResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrometheusMonitoringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrometheusMonitoringResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusMonitoringResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusMonitoringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrometheusMonitoringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrometheusMonitoringResponse) GetBody() *GetPrometheusMonitoringResponseBody {
	return s.Body
}

func (s *GetPrometheusMonitoringResponse) SetHeaders(v map[string]*string) *GetPrometheusMonitoringResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusMonitoringResponse) SetStatusCode(v int32) *GetPrometheusMonitoringResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusMonitoringResponse) SetBody(v *GetPrometheusMonitoringResponseBody) *GetPrometheusMonitoringResponse {
	s.Body = v
	return s
}

func (s *GetPrometheusMonitoringResponse) Validate() error {
	return dara.Validate(s)
}
