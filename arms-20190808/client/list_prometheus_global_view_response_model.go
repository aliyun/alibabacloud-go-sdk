// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusGlobalViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrometheusGlobalViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrometheusGlobalViewResponse
	GetStatusCode() *int32
	SetBody(v *ListPrometheusGlobalViewResponseBody) *ListPrometheusGlobalViewResponse
	GetBody() *ListPrometheusGlobalViewResponseBody
}

type ListPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrometheusGlobalViewResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusGlobalViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrometheusGlobalViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrometheusGlobalViewResponse) GetBody() *ListPrometheusGlobalViewResponseBody {
	return s.Body
}

func (s *ListPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *ListPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusGlobalViewResponse) SetStatusCode(v int32) *ListPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusGlobalViewResponse) SetBody(v *ListPrometheusGlobalViewResponseBody) *ListPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

func (s *ListPrometheusGlobalViewResponse) Validate() error {
	return dara.Validate(s)
}
