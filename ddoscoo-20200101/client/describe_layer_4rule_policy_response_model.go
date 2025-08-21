// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RulePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLayer4RulePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLayer4RulePolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLayer4RulePolicyResponseBody) *DescribeLayer4RulePolicyResponse
	GetBody() *DescribeLayer4RulePolicyResponseBody
}

type DescribeLayer4RulePolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLayer4RulePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLayer4RulePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulePolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLayer4RulePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLayer4RulePolicyResponse) GetBody() *DescribeLayer4RulePolicyResponseBody {
	return s.Body
}

func (s *DescribeLayer4RulePolicyResponse) SetHeaders(v map[string]*string) *DescribeLayer4RulePolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeLayer4RulePolicyResponse) SetStatusCode(v int32) *DescribeLayer4RulePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLayer4RulePolicyResponse) SetBody(v *DescribeLayer4RulePolicyResponseBody) *DescribeLayer4RulePolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeLayer4RulePolicyResponse) Validate() error {
	return dara.Validate(s)
}
