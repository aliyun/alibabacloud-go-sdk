// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigLayer4RuleAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigLayer4RuleAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ConfigLayer4RuleAttributeResponseBody) *ConfigLayer4RuleAttributeResponse
	GetBody() *ConfigLayer4RuleAttributeResponseBody
}

type ConfigLayer4RuleAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigLayer4RuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigLayer4RuleAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigLayer4RuleAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigLayer4RuleAttributeResponse) GetBody() *ConfigLayer4RuleAttributeResponseBody {
	return s.Body
}

func (s *ConfigLayer4RuleAttributeResponse) SetHeaders(v map[string]*string) *ConfigLayer4RuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *ConfigLayer4RuleAttributeResponse) SetStatusCode(v int32) *ConfigLayer4RuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigLayer4RuleAttributeResponse) SetBody(v *ConfigLayer4RuleAttributeResponseBody) *ConfigLayer4RuleAttributeResponse {
	s.Body = v
	return s
}

func (s *ConfigLayer4RuleAttributeResponse) Validate() error {
	return dara.Validate(s)
}
