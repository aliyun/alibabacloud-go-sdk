// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListABMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListABMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ListABMetricsResponseBody) *ListABMetricsResponse
	GetBody() *ListABMetricsResponseBody
}

type ListABMetricsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListABMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListABMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListABMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListABMetricsResponse) GetBody() *ListABMetricsResponseBody {
	return s.Body
}

func (s *ListABMetricsResponse) SetHeaders(v map[string]*string) *ListABMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListABMetricsResponse) SetStatusCode(v int32) *ListABMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABMetricsResponse) SetBody(v *ListABMetricsResponseBody) *ListABMetricsResponse {
	s.Body = v
	return s
}

func (s *ListABMetricsResponse) Validate() error {
	return dara.Validate(s)
}
