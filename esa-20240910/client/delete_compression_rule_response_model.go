// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompressionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCompressionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCompressionRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCompressionRuleResponseBody) *DeleteCompressionRuleResponse
	GetBody() *DeleteCompressionRuleResponseBody
}

type DeleteCompressionRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCompressionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCompressionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompressionRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCompressionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCompressionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCompressionRuleResponse) GetBody() *DeleteCompressionRuleResponseBody {
	return s.Body
}

func (s *DeleteCompressionRuleResponse) SetHeaders(v map[string]*string) *DeleteCompressionRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCompressionRuleResponse) SetStatusCode(v int32) *DeleteCompressionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCompressionRuleResponse) SetBody(v *DeleteCompressionRuleResponseBody) *DeleteCompressionRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteCompressionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
