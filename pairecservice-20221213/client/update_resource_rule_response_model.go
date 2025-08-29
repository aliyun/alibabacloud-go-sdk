// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResourceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResourceRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResourceRuleResponseBody) *UpdateResourceRuleResponse
	GetBody() *UpdateResourceRuleResponseBody
}

type UpdateResourceRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResourceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResourceRuleResponse) GetBody() *UpdateResourceRuleResponseBody {
	return s.Body
}

func (s *UpdateResourceRuleResponse) SetHeaders(v map[string]*string) *UpdateResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceRuleResponse) SetStatusCode(v int32) *UpdateResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceRuleResponse) SetBody(v *UpdateResourceRuleResponseBody) *UpdateResourceRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateResourceRuleResponse) Validate() error {
	return dara.Validate(s)
}
