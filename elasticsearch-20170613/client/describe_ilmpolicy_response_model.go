// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeILMPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeILMPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeILMPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeILMPolicyResponseBody) *DescribeILMPolicyResponse
	GetBody() *DescribeILMPolicyResponseBody
}

type DescribeILMPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeILMPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeILMPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeILMPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeILMPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeILMPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeILMPolicyResponse) GetBody() *DescribeILMPolicyResponseBody {
	return s.Body
}

func (s *DescribeILMPolicyResponse) SetHeaders(v map[string]*string) *DescribeILMPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeILMPolicyResponse) SetStatusCode(v int32) *DescribeILMPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeILMPolicyResponse) SetBody(v *DescribeILMPolicyResponseBody) *DescribeILMPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeILMPolicyResponse) Validate() error {
	return dara.Validate(s)
}
