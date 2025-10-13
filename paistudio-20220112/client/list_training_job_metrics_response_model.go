// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrainingJobMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrainingJobMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrainingJobMetricsResponseBody) *ListTrainingJobMetricsResponse
	GetBody() *ListTrainingJobMetricsResponseBody
}

type ListTrainingJobMetricsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrainingJobMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrainingJobMetricsResponse) GetBody() *ListTrainingJobMetricsResponseBody {
	return s.Body
}

func (s *ListTrainingJobMetricsResponse) SetHeaders(v map[string]*string) *ListTrainingJobMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobMetricsResponse) SetStatusCode(v int32) *ListTrainingJobMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobMetricsResponse) SetBody(v *ListTrainingJobMetricsResponseBody) *ListTrainingJobMetricsResponse {
	s.Body = v
	return s
}

func (s *ListTrainingJobMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
