// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationMetricDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEvaluationMetricDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEvaluationMetricDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListEvaluationMetricDetailsResponseBody) *ListEvaluationMetricDetailsResponse
	GetBody() *ListEvaluationMetricDetailsResponseBody
}

type ListEvaluationMetricDetailsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationMetricDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationMetricDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationMetricDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEvaluationMetricDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEvaluationMetricDetailsResponse) GetBody() *ListEvaluationMetricDetailsResponseBody {
	return s.Body
}

func (s *ListEvaluationMetricDetailsResponse) SetHeaders(v map[string]*string) *ListEvaluationMetricDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationMetricDetailsResponse) SetStatusCode(v int32) *ListEvaluationMetricDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponse) SetBody(v *ListEvaluationMetricDetailsResponseBody) *ListEvaluationMetricDetailsResponse {
	s.Body = v
	return s
}

func (s *ListEvaluationMetricDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
