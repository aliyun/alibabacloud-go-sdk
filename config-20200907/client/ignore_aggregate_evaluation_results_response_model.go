// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreAggregateEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IgnoreAggregateEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IgnoreAggregateEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *IgnoreAggregateEvaluationResultsResponseBody) *IgnoreAggregateEvaluationResultsResponse
	GetBody() *IgnoreAggregateEvaluationResultsResponseBody
}

type IgnoreAggregateEvaluationResultsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IgnoreAggregateEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IgnoreAggregateEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IgnoreAggregateEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IgnoreAggregateEvaluationResultsResponse) GetBody() *IgnoreAggregateEvaluationResultsResponseBody {
	return s.Body
}

func (s *IgnoreAggregateEvaluationResultsResponse) SetHeaders(v map[string]*string) *IgnoreAggregateEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *IgnoreAggregateEvaluationResultsResponse) SetStatusCode(v int32) *IgnoreAggregateEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsResponse) SetBody(v *IgnoreAggregateEvaluationResultsResponseBody) *IgnoreAggregateEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *IgnoreAggregateEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
