// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPrometheusIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPrometheusIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *AddPrometheusIntegrationResponseBody) *AddPrometheusIntegrationResponse
	GetBody() *AddPrometheusIntegrationResponseBody
}

type AddPrometheusIntegrationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPrometheusIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPrometheusIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusIntegrationResponse) GoString() string {
	return s.String()
}

func (s *AddPrometheusIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPrometheusIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPrometheusIntegrationResponse) GetBody() *AddPrometheusIntegrationResponseBody {
	return s.Body
}

func (s *AddPrometheusIntegrationResponse) SetHeaders(v map[string]*string) *AddPrometheusIntegrationResponse {
	s.Headers = v
	return s
}

func (s *AddPrometheusIntegrationResponse) SetStatusCode(v int32) *AddPrometheusIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrometheusIntegrationResponse) SetBody(v *AddPrometheusIntegrationResponseBody) *AddPrometheusIntegrationResponse {
	s.Body = v
	return s
}

func (s *AddPrometheusIntegrationResponse) Validate() error {
	return dara.Validate(s)
}
