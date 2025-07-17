// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDIJobMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDIJobMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ListDIJobMetricsResponseBody) *ListDIJobMetricsResponse
	GetBody() *ListDIJobMetricsResponseBody
}

type ListDIJobMetricsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDIJobMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDIJobMetricsResponse) GetBody() *ListDIJobMetricsResponseBody {
	return s.Body
}

func (s *ListDIJobMetricsResponse) SetHeaders(v map[string]*string) *ListDIJobMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobMetricsResponse) SetStatusCode(v int32) *ListDIJobMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobMetricsResponse) SetBody(v *ListDIJobMetricsResponseBody) *ListDIJobMetricsResponse {
	s.Body = v
	return s
}

func (s *ListDIJobMetricsResponse) Validate() error {
	return dara.Validate(s)
}
