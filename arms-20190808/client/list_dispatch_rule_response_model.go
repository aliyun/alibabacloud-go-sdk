// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDispatchRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDispatchRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDispatchRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListDispatchRuleResponseBody) *ListDispatchRuleResponse
	GetBody() *ListDispatchRuleResponseBody
}

type ListDispatchRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDispatchRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDispatchRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDispatchRuleResponse) GetBody() *ListDispatchRuleResponseBody {
	return s.Body
}

func (s *ListDispatchRuleResponse) SetHeaders(v map[string]*string) *ListDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *ListDispatchRuleResponse) SetStatusCode(v int32) *ListDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDispatchRuleResponse) SetBody(v *ListDispatchRuleResponseBody) *ListDispatchRuleResponse {
	s.Body = v
	return s
}

func (s *ListDispatchRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
