// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterceptionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInterceptionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInterceptionRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateInterceptionRuleResponseBody) *CreateInterceptionRuleResponse
	GetBody() *CreateInterceptionRuleResponseBody
}

type CreateInterceptionRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInterceptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInterceptionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInterceptionRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateInterceptionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInterceptionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInterceptionRuleResponse) GetBody() *CreateInterceptionRuleResponseBody {
	return s.Body
}

func (s *CreateInterceptionRuleResponse) SetHeaders(v map[string]*string) *CreateInterceptionRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateInterceptionRuleResponse) SetStatusCode(v int32) *CreateInterceptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInterceptionRuleResponse) SetBody(v *CreateInterceptionRuleResponseBody) *CreateInterceptionRuleResponse {
	s.Body = v
	return s
}

func (s *CreateInterceptionRuleResponse) Validate() error {
	return dara.Validate(s)
}
