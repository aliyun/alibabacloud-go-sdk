// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigL7GlobalRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigL7GlobalRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigL7GlobalRuleResponse
	GetStatusCode() *int32
	SetBody(v *ConfigL7GlobalRuleResponseBody) *ConfigL7GlobalRuleResponse
	GetBody() *ConfigL7GlobalRuleResponseBody
}

type ConfigL7GlobalRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigL7GlobalRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigL7GlobalRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigL7GlobalRuleResponse) GoString() string {
	return s.String()
}

func (s *ConfigL7GlobalRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigL7GlobalRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigL7GlobalRuleResponse) GetBody() *ConfigL7GlobalRuleResponseBody {
	return s.Body
}

func (s *ConfigL7GlobalRuleResponse) SetHeaders(v map[string]*string) *ConfigL7GlobalRuleResponse {
	s.Headers = v
	return s
}

func (s *ConfigL7GlobalRuleResponse) SetStatusCode(v int32) *ConfigL7GlobalRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigL7GlobalRuleResponse) SetBody(v *ConfigL7GlobalRuleResponseBody) *ConfigL7GlobalRuleResponse {
	s.Body = v
	return s
}

func (s *ConfigL7GlobalRuleResponse) Validate() error {
	return dara.Validate(s)
}
