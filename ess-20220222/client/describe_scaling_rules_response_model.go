// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScalingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScalingRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScalingRulesResponseBody) *DescribeScalingRulesResponse
	GetBody() *DescribeScalingRulesResponseBody
}

type DescribeScalingRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScalingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScalingRulesResponse) GetBody() *DescribeScalingRulesResponseBody {
	return s.Body
}

func (s *DescribeScalingRulesResponse) SetHeaders(v map[string]*string) *DescribeScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingRulesResponse) SetStatusCode(v int32) *DescribeScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingRulesResponse) SetBody(v *DescribeScalingRulesResponseBody) *DescribeScalingRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeScalingRulesResponse) Validate() error {
	return dara.Validate(s)
}
