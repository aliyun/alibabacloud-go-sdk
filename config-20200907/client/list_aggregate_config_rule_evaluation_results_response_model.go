// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRuleEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateConfigRuleEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateConfigRuleEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateConfigRuleEvaluationResultsResponseBody) *ListAggregateConfigRuleEvaluationResultsResponse
	GetBody() *ListAggregateConfigRuleEvaluationResultsResponseBody
}

type ListAggregateConfigRuleEvaluationResultsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateConfigRuleEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) GetBody() *ListAggregateConfigRuleEvaluationResultsResponseBody {
	return s.Body
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListAggregateConfigRuleEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) SetStatusCode(v int32) *ListAggregateConfigRuleEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) SetBody(v *ListAggregateConfigRuleEvaluationResultsResponseBody) *ListAggregateConfigRuleEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
