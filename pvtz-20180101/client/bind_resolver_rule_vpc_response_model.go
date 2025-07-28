// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindResolverRuleVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindResolverRuleVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindResolverRuleVpcResponse
	GetStatusCode() *int32
	SetBody(v *BindResolverRuleVpcResponseBody) *BindResolverRuleVpcResponse
	GetBody() *BindResolverRuleVpcResponseBody
}

type BindResolverRuleVpcResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindResolverRuleVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindResolverRuleVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s BindResolverRuleVpcResponse) GoString() string {
	return s.String()
}

func (s *BindResolverRuleVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindResolverRuleVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindResolverRuleVpcResponse) GetBody() *BindResolverRuleVpcResponseBody {
	return s.Body
}

func (s *BindResolverRuleVpcResponse) SetHeaders(v map[string]*string) *BindResolverRuleVpcResponse {
	s.Headers = v
	return s
}

func (s *BindResolverRuleVpcResponse) SetStatusCode(v int32) *BindResolverRuleVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *BindResolverRuleVpcResponse) SetBody(v *BindResolverRuleVpcResponseBody) *BindResolverRuleVpcResponse {
	s.Body = v
	return s
}

func (s *BindResolverRuleVpcResponse) Validate() error {
	return dara.Validate(s)
}
