// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAntiBruteForceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAntiBruteForceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAntiBruteForceRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAntiBruteForceRuleResponseBody) *ModifyAntiBruteForceRuleResponse
	GetBody() *ModifyAntiBruteForceRuleResponseBody
}

type ModifyAntiBruteForceRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAntiBruteForceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAntiBruteForceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAntiBruteForceRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAntiBruteForceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAntiBruteForceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAntiBruteForceRuleResponse) GetBody() *ModifyAntiBruteForceRuleResponseBody {
	return s.Body
}

func (s *ModifyAntiBruteForceRuleResponse) SetHeaders(v map[string]*string) *ModifyAntiBruteForceRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAntiBruteForceRuleResponse) SetStatusCode(v int32) *ModifyAntiBruteForceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAntiBruteForceRuleResponse) SetBody(v *ModifyAntiBruteForceRuleResponseBody) *ModifyAntiBruteForceRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyAntiBruteForceRuleResponse) Validate() error {
	return dara.Validate(s)
}
