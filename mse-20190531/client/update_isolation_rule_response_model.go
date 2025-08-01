// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIsolationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIsolationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIsolationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIsolationRuleResponseBody) *UpdateIsolationRuleResponse
	GetBody() *UpdateIsolationRuleResponseBody
}

type UpdateIsolationRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIsolationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIsolationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIsolationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateIsolationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIsolationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIsolationRuleResponse) GetBody() *UpdateIsolationRuleResponseBody {
	return s.Body
}

func (s *UpdateIsolationRuleResponse) SetHeaders(v map[string]*string) *UpdateIsolationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateIsolationRuleResponse) SetStatusCode(v int32) *UpdateIsolationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIsolationRuleResponse) SetBody(v *UpdateIsolationRuleResponseBody) *UpdateIsolationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateIsolationRuleResponse) Validate() error {
	return dara.Validate(s)
}
