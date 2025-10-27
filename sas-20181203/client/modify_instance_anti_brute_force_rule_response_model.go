// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAntiBruteForceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceAntiBruteForceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceAntiBruteForceRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceAntiBruteForceRuleResponseBody) *ModifyInstanceAntiBruteForceRuleResponse
	GetBody() *ModifyInstanceAntiBruteForceRuleResponseBody
}

type ModifyInstanceAntiBruteForceRuleResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceAntiBruteForceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceAntiBruteForceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAntiBruteForceRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) GetBody() *ModifyInstanceAntiBruteForceRuleResponseBody {
	return s.Body
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) SetHeaders(v map[string]*string) *ModifyInstanceAntiBruteForceRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) SetStatusCode(v int32) *ModifyInstanceAntiBruteForceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) SetBody(v *ModifyInstanceAntiBruteForceRuleResponseBody) *ModifyInstanceAntiBruteForceRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceAntiBruteForceRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
