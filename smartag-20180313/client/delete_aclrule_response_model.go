// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteACLRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteACLRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteACLRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteACLRuleResponseBody) *DeleteACLRuleResponse
	GetBody() *DeleteACLRuleResponseBody
}

type DeleteACLRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteACLRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteACLRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteACLRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteACLRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteACLRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteACLRuleResponse) GetBody() *DeleteACLRuleResponseBody {
	return s.Body
}

func (s *DeleteACLRuleResponse) SetHeaders(v map[string]*string) *DeleteACLRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteACLRuleResponse) SetStatusCode(v int32) *DeleteACLRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteACLRuleResponse) SetBody(v *DeleteACLRuleResponseBody) *DeleteACLRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteACLRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
