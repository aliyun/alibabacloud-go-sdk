// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7CCRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer7CCRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer7CCRuleResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer7CCRuleResponseBody) *ConfigLayer7CCRuleResponse
	GetBody() *ConfigLayer7CCRuleResponseBody
}

type ConfigLayer7CCRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer7CCRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CCRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer7CCRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer7CCRuleResponse) GetBody() *ConfigLayer7CCRuleResponseBody {
	return s.Body
}

func (s *ConfigLayer7CCRuleResponse) SetHeaders(v map[string]*string) *ConfigLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7CCRuleResponse) SetStatusCode(v int32) *ConfigLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7CCRuleResponse) SetBody(v *ConfigLayer7CCRuleResponseBody) *ConfigLayer7CCRuleResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer7CCRuleResponse) Validate() error {
	return dara.Validate(s)
}
