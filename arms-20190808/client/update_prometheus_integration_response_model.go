// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrometheusIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrometheusIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrometheusIntegrationResponseBody) *UpdatePrometheusIntegrationResponse
	GetBody() *UpdatePrometheusIntegrationResponseBody
}

type UpdatePrometheusIntegrationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrometheusIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrometheusIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusIntegrationResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrometheusIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrometheusIntegrationResponse) GetBody() *UpdatePrometheusIntegrationResponseBody {
	return s.Body
}

func (s *UpdatePrometheusIntegrationResponse) SetHeaders(v map[string]*string) *UpdatePrometheusIntegrationResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusIntegrationResponse) SetStatusCode(v int32) *UpdatePrometheusIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusIntegrationResponse) SetBody(v *UpdatePrometheusIntegrationResponseBody) *UpdatePrometheusIntegrationResponse {
	s.Body = v
	return s
}

func (s *UpdatePrometheusIntegrationResponse) Validate() error {
	return dara.Validate(s)
}
