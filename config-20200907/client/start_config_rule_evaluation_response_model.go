// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConfigRuleEvaluationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartConfigRuleEvaluationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartConfigRuleEvaluationResponse
	GetStatusCode() *int32
	SetBody(v *StartConfigRuleEvaluationResponseBody) *StartConfigRuleEvaluationResponse
	GetBody() *StartConfigRuleEvaluationResponseBody
}

type StartConfigRuleEvaluationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartConfigRuleEvaluationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartConfigRuleEvaluationResponse) String() string {
	return dara.Prettify(s)
}

func (s StartConfigRuleEvaluationResponse) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartConfigRuleEvaluationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartConfigRuleEvaluationResponse) GetBody() *StartConfigRuleEvaluationResponseBody {
	return s.Body
}

func (s *StartConfigRuleEvaluationResponse) SetHeaders(v map[string]*string) *StartConfigRuleEvaluationResponse {
	s.Headers = v
	return s
}

func (s *StartConfigRuleEvaluationResponse) SetStatusCode(v int32) *StartConfigRuleEvaluationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartConfigRuleEvaluationResponse) SetBody(v *StartConfigRuleEvaluationResponseBody) *StartConfigRuleEvaluationResponse {
	s.Body = v
	return s
}

func (s *StartConfigRuleEvaluationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
