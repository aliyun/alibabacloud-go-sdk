// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConfigRuleEvaluationByResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartConfigRuleEvaluationByResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartConfigRuleEvaluationByResourceResponse
	GetStatusCode() *int32
	SetBody(v *StartConfigRuleEvaluationByResourceResponseBody) *StartConfigRuleEvaluationByResourceResponse
	GetBody() *StartConfigRuleEvaluationByResourceResponseBody
}

type StartConfigRuleEvaluationByResourceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartConfigRuleEvaluationByResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartConfigRuleEvaluationByResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartConfigRuleEvaluationByResourceResponse) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationByResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartConfigRuleEvaluationByResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartConfigRuleEvaluationByResourceResponse) GetBody() *StartConfigRuleEvaluationByResourceResponseBody {
	return s.Body
}

func (s *StartConfigRuleEvaluationByResourceResponse) SetHeaders(v map[string]*string) *StartConfigRuleEvaluationByResourceResponse {
	s.Headers = v
	return s
}

func (s *StartConfigRuleEvaluationByResourceResponse) SetStatusCode(v int32) *StartConfigRuleEvaluationByResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartConfigRuleEvaluationByResourceResponse) SetBody(v *StartConfigRuleEvaluationByResourceResponseBody) *StartConfigRuleEvaluationByResourceResponse {
	s.Body = v
	return s
}

func (s *StartConfigRuleEvaluationByResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
