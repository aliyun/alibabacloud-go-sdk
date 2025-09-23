// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7RuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer7RuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer7RuleResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer7RuleResponseBody) *ConfigLayer7RuleResponse
	GetBody() *ConfigLayer7RuleResponseBody
}

type ConfigLayer7RuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer7RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer7RuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7RuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer7RuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer7RuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer7RuleResponse) GetBody() *ConfigLayer7RuleResponseBody {
	return s.Body
}

func (s *ConfigLayer7RuleResponse) SetHeaders(v map[string]*string) *ConfigLayer7RuleResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer7RuleResponse) SetStatusCode(v int32) *ConfigLayer7RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer7RuleResponse) SetBody(v *ConfigLayer7RuleResponseBody) *ConfigLayer7RuleResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer7RuleResponse) Validate() error {
	return dara.Validate(s)
}
