// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResolverRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResolverRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResolverRuleResponseBody) *DescribeResolverRuleResponse
	GetBody() *DescribeResolverRuleResponseBody
}

type DescribeResolverRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResolverRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResolverRuleResponse) GetBody() *DescribeResolverRuleResponseBody {
	return s.Body
}

func (s *DescribeResolverRuleResponse) SetHeaders(v map[string]*string) *DescribeResolverRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverRuleResponse) SetStatusCode(v int32) *DescribeResolverRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverRuleResponse) SetBody(v *DescribeResolverRuleResponseBody) *DescribeResolverRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeResolverRuleResponse) Validate() error {
	return dara.Validate(s)
}
