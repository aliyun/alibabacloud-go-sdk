// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateResourceEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateResourceEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateResourceEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateResourceEvaluationResultsResponseBody) *ListAggregateResourceEvaluationResultsResponse
	GetBody() *ListAggregateResourceEvaluationResultsResponseBody
}

type ListAggregateResourceEvaluationResultsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateResourceEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateResourceEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateResourceEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateResourceEvaluationResultsResponse) GetBody() *ListAggregateResourceEvaluationResultsResponseBody {
	return s.Body
}

func (s *ListAggregateResourceEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListAggregateResourceEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponse) SetStatusCode(v int32) *ListAggregateResourceEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponse) SetBody(v *ListAggregateResourceEvaluationResultsResponseBody) *ListAggregateResourceEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
