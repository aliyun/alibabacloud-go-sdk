// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusViewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusViewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusViewsResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusViewsResponseBody) *ListPrometheusViewsResponse
	GetBody() *ListPrometheusViewsResponseBody
}

type ListPrometheusViewsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusViewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusViewsResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusViewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusViewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusViewsResponse) GetBody() *ListPrometheusViewsResponseBody {
	return s.Body
}

func (s *ListPrometheusViewsResponse) SetHeaders(v map[string]*string) *ListPrometheusViewsResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusViewsResponse) SetStatusCode(v int32) *ListPrometheusViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusViewsResponse) SetBody(v *ListPrometheusViewsResponseBody) *ListPrometheusViewsResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusViewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
