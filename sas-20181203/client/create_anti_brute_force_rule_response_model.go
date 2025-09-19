// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAntiBruteForceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAntiBruteForceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAntiBruteForceRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateAntiBruteForceRuleResponseBody) *CreateAntiBruteForceRuleResponse
	GetBody() *CreateAntiBruteForceRuleResponseBody
}

type CreateAntiBruteForceRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAntiBruteForceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAntiBruteForceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAntiBruteForceRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAntiBruteForceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAntiBruteForceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAntiBruteForceRuleResponse) GetBody() *CreateAntiBruteForceRuleResponseBody {
	return s.Body
}

func (s *CreateAntiBruteForceRuleResponse) SetHeaders(v map[string]*string) *CreateAntiBruteForceRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAntiBruteForceRuleResponse) SetStatusCode(v int32) *CreateAntiBruteForceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAntiBruteForceRuleResponse) SetBody(v *CreateAntiBruteForceRuleResponseBody) *CreateAntiBruteForceRuleResponse {
	s.Body = v
	return s
}

func (s *CreateAntiBruteForceRuleResponse) Validate() error {
	return dara.Validate(s)
}
