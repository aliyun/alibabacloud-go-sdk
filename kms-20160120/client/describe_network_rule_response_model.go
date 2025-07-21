// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkRuleResponseBody) *DescribeNetworkRuleResponse
	GetBody() *DescribeNetworkRuleResponseBody
}

type DescribeNetworkRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkRuleResponse) GetBody() *DescribeNetworkRuleResponseBody {
	return s.Body
}

func (s *DescribeNetworkRuleResponse) SetHeaders(v map[string]*string) *DescribeNetworkRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkRuleResponse) SetStatusCode(v int32) *DescribeNetworkRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkRuleResponse) SetBody(v *DescribeNetworkRuleResponseBody) *DescribeNetworkRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkRuleResponse) Validate() error {
	return dara.Validate(s)
}
