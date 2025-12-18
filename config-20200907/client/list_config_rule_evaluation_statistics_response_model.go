// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRuleEvaluationStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigRuleEvaluationStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigRuleEvaluationStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigRuleEvaluationStatisticsResponseBody) *ListConfigRuleEvaluationStatisticsResponse
	GetBody() *ListConfigRuleEvaluationStatisticsResponseBody
}

type ListConfigRuleEvaluationStatisticsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigRuleEvaluationStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigRuleEvaluationStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigRuleEvaluationStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigRuleEvaluationStatisticsResponse) GetBody() *ListConfigRuleEvaluationStatisticsResponseBody {
	return s.Body
}

func (s *ListConfigRuleEvaluationStatisticsResponse) SetHeaders(v map[string]*string) *ListConfigRuleEvaluationStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponse) SetStatusCode(v int32) *ListConfigRuleEvaluationStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponse) SetBody(v *ListConfigRuleEvaluationStatisticsResponseBody) *ListConfigRuleEvaluationStatisticsResponse {
	s.Body = v
	return s
}

func (s *ListConfigRuleEvaluationStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
