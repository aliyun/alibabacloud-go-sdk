// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRuleResponseBody) *UpdateRuleResponse
	GetBody() *UpdateRuleResponseBody
}

type UpdateRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRuleResponse) GetBody() *UpdateRuleResponseBody {
	return s.Body
}

func (s *UpdateRuleResponse) SetHeaders(v map[string]*string) *UpdateRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleResponse) SetStatusCode(v int32) *UpdateRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRuleResponse) SetBody(v *UpdateRuleResponseBody) *UpdateRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateRuleResponse) Validate() error {
	return dara.Validate(s)
}
