// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrometheusIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrometheusIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrometheusIntegrationResponseBody) *DeletePrometheusIntegrationResponse
	GetBody() *DeletePrometheusIntegrationResponseBody
}

type DeletePrometheusIntegrationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrometheusIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrometheusIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusIntegrationResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrometheusIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrometheusIntegrationResponse) GetBody() *DeletePrometheusIntegrationResponseBody {
	return s.Body
}

func (s *DeletePrometheusIntegrationResponse) SetHeaders(v map[string]*string) *DeletePrometheusIntegrationResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusIntegrationResponse) SetStatusCode(v int32) *DeletePrometheusIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusIntegrationResponse) SetBody(v *DeletePrometheusIntegrationResponseBody) *DeletePrometheusIntegrationResponse {
	s.Body = v
	return s
}

func (s *DeletePrometheusIntegrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
