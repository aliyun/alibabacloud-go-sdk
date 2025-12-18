// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertAggregateEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevertAggregateEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevertAggregateEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *RevertAggregateEvaluationResultsResponseBody) *RevertAggregateEvaluationResultsResponse
	GetBody() *RevertAggregateEvaluationResultsResponseBody
}

type RevertAggregateEvaluationResultsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevertAggregateEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertAggregateEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s RevertAggregateEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevertAggregateEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevertAggregateEvaluationResultsResponse) GetBody() *RevertAggregateEvaluationResultsResponseBody {
	return s.Body
}

func (s *RevertAggregateEvaluationResultsResponse) SetHeaders(v map[string]*string) *RevertAggregateEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *RevertAggregateEvaluationResultsResponse) SetStatusCode(v int32) *RevertAggregateEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertAggregateEvaluationResultsResponse) SetBody(v *RevertAggregateEvaluationResultsResponseBody) *RevertAggregateEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *RevertAggregateEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
