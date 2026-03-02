// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceMetricsResponse
	GetStatusCode() *int32
	SetBody(v *MonitorMetricResult) *ListServiceMetricsResponse
	GetBody() *MonitorMetricResult
}

type ListServiceMetricsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MonitorMetricResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceMetricsResponse) GetBody() *MonitorMetricResult {
	return s.Body
}

func (s *ListServiceMetricsResponse) SetHeaders(v map[string]*string) *ListServiceMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceMetricsResponse) SetStatusCode(v int32) *ListServiceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceMetricsResponse) SetBody(v *MonitorMetricResult) *ListServiceMetricsResponse {
	s.Body = v
	return s
}

func (s *ListServiceMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
