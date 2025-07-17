// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDispatchRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDispatchRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDispatchRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDispatchRuleResponseBody) *DescribeDispatchRuleResponse
	GetBody() *DescribeDispatchRuleResponseBody
}

type DescribeDispatchRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDispatchRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDispatchRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDispatchRuleResponse) GetBody() *DescribeDispatchRuleResponseBody {
	return s.Body
}

func (s *DescribeDispatchRuleResponse) SetHeaders(v map[string]*string) *DescribeDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDispatchRuleResponse) SetStatusCode(v int32) *DescribeDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDispatchRuleResponse) SetBody(v *DescribeDispatchRuleResponseBody) *DescribeDispatchRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeDispatchRuleResponse) Validate() error {
	return dara.Validate(s)
}
