// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResolverRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateResolverRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateResolverRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateResolverRuleResponseBody) *UpdateResolverRuleResponse
	GetBody() *UpdateResolverRuleResponseBody
}

type UpdateResolverRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResolverRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResolverRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateResolverRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateResolverRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateResolverRuleResponse) GetBody() *UpdateResolverRuleResponseBody {
	return s.Body
}

func (s *UpdateResolverRuleResponse) SetHeaders(v map[string]*string) *UpdateResolverRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateResolverRuleResponse) SetStatusCode(v int32) *UpdateResolverRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResolverRuleResponse) SetBody(v *UpdateResolverRuleResponseBody) *UpdateResolverRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateResolverRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
