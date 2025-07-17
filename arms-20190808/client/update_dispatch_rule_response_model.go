// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDispatchRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDispatchRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDispatchRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDispatchRuleResponseBody) *UpdateDispatchRuleResponse
	GetBody() *UpdateDispatchRuleResponseBody
}

type UpdateDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDispatchRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDispatchRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDispatchRuleResponse) GetBody() *UpdateDispatchRuleResponseBody {
	return s.Body
}

func (s *UpdateDispatchRuleResponse) SetHeaders(v map[string]*string) *UpdateDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDispatchRuleResponse) SetStatusCode(v int32) *UpdateDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDispatchRuleResponse) SetBody(v *UpdateDispatchRuleResponseBody) *UpdateDispatchRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateDispatchRuleResponse) Validate() error {
	return dara.Validate(s)
}
