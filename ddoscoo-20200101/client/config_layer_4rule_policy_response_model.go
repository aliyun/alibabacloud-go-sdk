// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RulePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer4RulePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer4RulePolicyResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer4RulePolicyResponseBody) *ConfigLayer4RulePolicyResponse
	GetBody() *ConfigLayer4RulePolicyResponseBody
}

type ConfigLayer4RulePolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer4RulePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer4RulePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RulePolicyResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RulePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer4RulePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer4RulePolicyResponse) GetBody() *ConfigLayer4RulePolicyResponseBody {
	return s.Body
}

func (s *ConfigLayer4RulePolicyResponse) SetHeaders(v map[string]*string) *ConfigLayer4RulePolicyResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer4RulePolicyResponse) SetStatusCode(v int32) *ConfigLayer4RulePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer4RulePolicyResponse) SetBody(v *ConfigLayer4RulePolicyResponseBody) *ConfigLayer4RulePolicyResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer4RulePolicyResponse) Validate() error {
	return dara.Validate(s)
}
