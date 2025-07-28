// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddResolverRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddResolverRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddResolverRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddResolverRuleResponseBody) *AddResolverRuleResponse
	GetBody() *AddResolverRuleResponseBody
}

type AddResolverRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddResolverRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddResolverRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddResolverRuleResponse) GoString() string {
	return s.String()
}

func (s *AddResolverRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddResolverRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddResolverRuleResponse) GetBody() *AddResolverRuleResponseBody {
	return s.Body
}

func (s *AddResolverRuleResponse) SetHeaders(v map[string]*string) *AddResolverRuleResponse {
	s.Headers = v
	return s
}

func (s *AddResolverRuleResponse) SetStatusCode(v int32) *AddResolverRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddResolverRuleResponse) SetBody(v *AddResolverRuleResponseBody) *AddResolverRuleResponse {
	s.Body = v
	return s
}

func (s *AddResolverRuleResponse) Validate() error {
	return dara.Validate(s)
}
