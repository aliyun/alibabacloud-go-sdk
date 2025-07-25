// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobInstanceMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrainingJobInstanceMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrainingJobInstanceMetricsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrainingJobInstanceMetricsResponseBody) *ListTrainingJobInstanceMetricsResponse
	GetBody() *ListTrainingJobInstanceMetricsResponseBody
}

type ListTrainingJobInstanceMetricsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobInstanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobInstanceMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobInstanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobInstanceMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrainingJobInstanceMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrainingJobInstanceMetricsResponse) GetBody() *ListTrainingJobInstanceMetricsResponseBody {
	return s.Body
}

func (s *ListTrainingJobInstanceMetricsResponse) SetHeaders(v map[string]*string) *ListTrainingJobInstanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponse) SetStatusCode(v int32) *ListTrainingJobInstanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponse) SetBody(v *ListTrainingJobInstanceMetricsResponseBody) *ListTrainingJobInstanceMetricsResponse {
	s.Body = v
	return s
}

func (s *ListTrainingJobInstanceMetricsResponse) Validate() error {
	return dara.Validate(s)
}
