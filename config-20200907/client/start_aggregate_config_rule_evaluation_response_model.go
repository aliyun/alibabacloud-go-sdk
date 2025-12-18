// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAggregateConfigRuleEvaluationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAggregateConfigRuleEvaluationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAggregateConfigRuleEvaluationResponse
	GetStatusCode() *int32
	SetBody(v *StartAggregateConfigRuleEvaluationResponseBody) *StartAggregateConfigRuleEvaluationResponse
	GetBody() *StartAggregateConfigRuleEvaluationResponseBody
}

type StartAggregateConfigRuleEvaluationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAggregateConfigRuleEvaluationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAggregateConfigRuleEvaluationResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAggregateConfigRuleEvaluationResponse) GoString() string {
	return s.String()
}

func (s *StartAggregateConfigRuleEvaluationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAggregateConfigRuleEvaluationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAggregateConfigRuleEvaluationResponse) GetBody() *StartAggregateConfigRuleEvaluationResponseBody {
	return s.Body
}

func (s *StartAggregateConfigRuleEvaluationResponse) SetHeaders(v map[string]*string) *StartAggregateConfigRuleEvaluationResponse {
	s.Headers = v
	return s
}

func (s *StartAggregateConfigRuleEvaluationResponse) SetStatusCode(v int32) *StartAggregateConfigRuleEvaluationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationResponse) SetBody(v *StartAggregateConfigRuleEvaluationResponseBody) *StartAggregateConfigRuleEvaluationResponse {
	s.Body = v
	return s
}

func (s *StartAggregateConfigRuleEvaluationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
