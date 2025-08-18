// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompressionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCompressionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCompressionRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetCompressionRuleResponseBody) *GetCompressionRuleResponse
	GetBody() *GetCompressionRuleResponseBody
}

type GetCompressionRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCompressionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCompressionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCompressionRuleResponse) GoString() string {
	return s.String()
}

func (s *GetCompressionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCompressionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCompressionRuleResponse) GetBody() *GetCompressionRuleResponseBody {
	return s.Body
}

func (s *GetCompressionRuleResponse) SetHeaders(v map[string]*string) *GetCompressionRuleResponse {
	s.Headers = v
	return s
}

func (s *GetCompressionRuleResponse) SetStatusCode(v int32) *GetCompressionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompressionRuleResponse) SetBody(v *GetCompressionRuleResponseBody) *GetCompressionRuleResponse {
	s.Body = v
	return s
}

func (s *GetCompressionRuleResponse) Validate() error {
	return dara.Validate(s)
}
