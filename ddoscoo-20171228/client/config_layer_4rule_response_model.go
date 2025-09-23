// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer4RuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer4RuleResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer4RuleResponseBody) *ConfigLayer4RuleResponse
	GetBody() *ConfigLayer4RuleResponseBody
}

type ConfigLayer4RuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer4RuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer4RuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer4RuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer4RuleResponse) GetBody() *ConfigLayer4RuleResponseBody {
	return s.Body
}

func (s *ConfigLayer4RuleResponse) SetHeaders(v map[string]*string) *ConfigLayer4RuleResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer4RuleResponse) SetStatusCode(v int32) *ConfigLayer4RuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer4RuleResponse) SetBody(v *ConfigLayer4RuleResponseBody) *ConfigLayer4RuleResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer4RuleResponse) Validate() error {
	return dara.Validate(s)
}
