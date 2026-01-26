// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusInstancesResponseBody) *ListPrometheusInstancesResponse
	GetBody() *ListPrometheusInstancesResponseBody
}

type ListPrometheusInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusInstancesResponse) GetBody() *ListPrometheusInstancesResponseBody {
	return s.Body
}

func (s *ListPrometheusInstancesResponse) SetHeaders(v map[string]*string) *ListPrometheusInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusInstancesResponse) SetStatusCode(v int32) *ListPrometheusInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusInstancesResponse) SetBody(v *ListPrometheusInstancesResponseBody) *ListPrometheusInstancesResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
