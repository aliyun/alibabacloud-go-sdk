// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompressionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCompressionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCompressionRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCompressionRuleResponseBody) *UpdateCompressionRuleResponse
	GetBody() *UpdateCompressionRuleResponseBody
}

type UpdateCompressionRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCompressionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCompressionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompressionRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateCompressionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCompressionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCompressionRuleResponse) GetBody() *UpdateCompressionRuleResponseBody {
	return s.Body
}

func (s *UpdateCompressionRuleResponse) SetHeaders(v map[string]*string) *UpdateCompressionRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateCompressionRuleResponse) SetStatusCode(v int32) *UpdateCompressionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCompressionRuleResponse) SetBody(v *UpdateCompressionRuleResponseBody) *UpdateCompressionRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateCompressionRuleResponse) Validate() error {
	return dara.Validate(s)
}
