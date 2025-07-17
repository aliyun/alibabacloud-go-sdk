// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrometheusIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrometheusIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *GetPrometheusIntegrationResponseBody) *GetPrometheusIntegrationResponse
	GetBody() *GetPrometheusIntegrationResponseBody
}

type GetPrometheusIntegrationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrometheusIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrometheusIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusIntegrationResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrometheusIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrometheusIntegrationResponse) GetBody() *GetPrometheusIntegrationResponseBody {
	return s.Body
}

func (s *GetPrometheusIntegrationResponse) SetHeaders(v map[string]*string) *GetPrometheusIntegrationResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusIntegrationResponse) SetStatusCode(v int32) *GetPrometheusIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusIntegrationResponse) SetBody(v *GetPrometheusIntegrationResponseBody) *GetPrometheusIntegrationResponse {
	s.Body = v
	return s
}

func (s *GetPrometheusIntegrationResponse) Validate() error {
	return dara.Validate(s)
}
