// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusIntegrationResponseBody) *ListPrometheusIntegrationResponse
	GetBody() *ListPrometheusIntegrationResponseBody
}

type ListPrometheusIntegrationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusIntegrationResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusIntegrationResponse) GetBody() *ListPrometheusIntegrationResponseBody {
	return s.Body
}

func (s *ListPrometheusIntegrationResponse) SetHeaders(v map[string]*string) *ListPrometheusIntegrationResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusIntegrationResponse) SetStatusCode(v int32) *ListPrometheusIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusIntegrationResponse) SetBody(v *ListPrometheusIntegrationResponseBody) *ListPrometheusIntegrationResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusIntegrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
