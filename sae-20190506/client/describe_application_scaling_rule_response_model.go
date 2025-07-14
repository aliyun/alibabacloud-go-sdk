// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationScalingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationScalingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationScalingRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationScalingRuleResponseBody) *DescribeApplicationScalingRuleResponse
	GetBody() *DescribeApplicationScalingRuleResponseBody
}

type DescribeApplicationScalingRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationScalingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationScalingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationScalingRuleResponse) GetBody() *DescribeApplicationScalingRuleResponseBody {
	return s.Body
}

func (s *DescribeApplicationScalingRuleResponse) SetHeaders(v map[string]*string) *DescribeApplicationScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationScalingRuleResponse) SetStatusCode(v int32) *DescribeApplicationScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationScalingRuleResponse) SetBody(v *DescribeApplicationScalingRuleResponseBody) *DescribeApplicationScalingRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationScalingRuleResponse) Validate() error {
	return dara.Validate(s)
}
