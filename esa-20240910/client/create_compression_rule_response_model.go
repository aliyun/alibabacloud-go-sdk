// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompressionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCompressionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCompressionRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCompressionRuleResponseBody) *CreateCompressionRuleResponse
	GetBody() *CreateCompressionRuleResponseBody
}

type CreateCompressionRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCompressionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCompressionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCompressionRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateCompressionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCompressionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCompressionRuleResponse) GetBody() *CreateCompressionRuleResponseBody {
	return s.Body
}

func (s *CreateCompressionRuleResponse) SetHeaders(v map[string]*string) *CreateCompressionRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateCompressionRuleResponse) SetStatusCode(v int32) *CreateCompressionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCompressionRuleResponse) SetBody(v *CreateCompressionRuleResponseBody) *CreateCompressionRuleResponse {
	s.Body = v
	return s
}

func (s *CreateCompressionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
