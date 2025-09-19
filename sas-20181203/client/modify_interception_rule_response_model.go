// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInterceptionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInterceptionRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInterceptionRuleResponseBody) *ModifyInterceptionRuleResponse
	GetBody() *ModifyInterceptionRuleResponseBody
}

type ModifyInterceptionRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInterceptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInterceptionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInterceptionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInterceptionRuleResponse) GetBody() *ModifyInterceptionRuleResponseBody {
	return s.Body
}

func (s *ModifyInterceptionRuleResponse) SetHeaders(v map[string]*string) *ModifyInterceptionRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyInterceptionRuleResponse) SetStatusCode(v int32) *ModifyInterceptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInterceptionRuleResponse) SetBody(v *ModifyInterceptionRuleResponseBody) *ModifyInterceptionRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyInterceptionRuleResponse) Validate() error {
	return dara.Validate(s)
}
