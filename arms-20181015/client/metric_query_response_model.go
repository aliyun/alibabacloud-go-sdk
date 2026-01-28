// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MetricQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MetricQueryResponse
	GetStatusCode() *int32
	SetBody(v *MetricQueryResponseBody) *MetricQueryResponse
	GetBody() *MetricQueryResponseBody
}

type MetricQueryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MetricQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MetricQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s MetricQueryResponse) GoString() string {
	return s.String()
}

func (s *MetricQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MetricQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MetricQueryResponse) GetBody() *MetricQueryResponseBody {
	return s.Body
}

func (s *MetricQueryResponse) SetHeaders(v map[string]*string) *MetricQueryResponse {
	s.Headers = v
	return s
}

func (s *MetricQueryResponse) SetStatusCode(v int32) *MetricQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *MetricQueryResponse) SetBody(v *MetricQueryResponseBody) *MetricQueryResponse {
	s.Body = v
	return s
}

func (s *MetricQueryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
