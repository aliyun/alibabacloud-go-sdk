// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRuleEvaluationResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigRuleEvaluationResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigRuleEvaluationResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigRuleEvaluationResultsResponseBody) *ListConfigRuleEvaluationResultsResponse
	GetBody() *ListConfigRuleEvaluationResultsResponseBody
}

type ListConfigRuleEvaluationResultsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigRuleEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigRuleEvaluationResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigRuleEvaluationResultsResponse) GetBody() *ListConfigRuleEvaluationResultsResponseBody {
	return s.Body
}

func (s *ListConfigRuleEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListConfigRuleEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponse) SetStatusCode(v int32) *ListConfigRuleEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponse) SetBody(v *ListConfigRuleEvaluationResultsResponseBody) *ListConfigRuleEvaluationResultsResponse {
	s.Body = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
