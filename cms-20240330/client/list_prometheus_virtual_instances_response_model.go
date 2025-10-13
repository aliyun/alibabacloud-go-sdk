// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusVirtualInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusVirtualInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusVirtualInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusVirtualInstancesResponseBody) *ListPrometheusVirtualInstancesResponse
	GetBody() *ListPrometheusVirtualInstancesResponseBody
}

type ListPrometheusVirtualInstancesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusVirtualInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusVirtualInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusVirtualInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusVirtualInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusVirtualInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusVirtualInstancesResponse) GetBody() *ListPrometheusVirtualInstancesResponseBody {
	return s.Body
}

func (s *ListPrometheusVirtualInstancesResponse) SetHeaders(v map[string]*string) *ListPrometheusVirtualInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusVirtualInstancesResponse) SetStatusCode(v int32) *ListPrometheusVirtualInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusVirtualInstancesResponse) SetBody(v *ListPrometheusVirtualInstancesResponseBody) *ListPrometheusVirtualInstancesResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusVirtualInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
