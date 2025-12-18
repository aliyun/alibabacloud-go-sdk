// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigRuleEvaluationStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAggregateConfigRuleEvaluationStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAggregateConfigRuleEvaluationStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *ListAggregateConfigRuleEvaluationStatisticsResponseBody) *ListAggregateConfigRuleEvaluationStatisticsResponse
	GetBody() *ListAggregateConfigRuleEvaluationStatisticsResponseBody
}

type ListAggregateConfigRuleEvaluationStatisticsResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAggregateConfigRuleEvaluationStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponse) GetBody() *ListAggregateConfigRuleEvaluationStatisticsResponseBody {
	return s.Body
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponse) SetHeaders(v map[string]*string) *ListAggregateConfigRuleEvaluationStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponse) SetStatusCode(v int32) *ListAggregateConfigRuleEvaluationStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponse) SetBody(v *ListAggregateConfigRuleEvaluationStatisticsResponseBody) *ListAggregateConfigRuleEvaluationStatisticsResponse {
	s.Body = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
