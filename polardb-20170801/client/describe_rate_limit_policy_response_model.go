// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRateLimitPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRateLimitPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRateLimitPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRateLimitPolicyResponseBody) *DescribeRateLimitPolicyResponse
	GetBody() *DescribeRateLimitPolicyResponseBody
}

type DescribeRateLimitPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRateLimitPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRateLimitPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRateLimitPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeRateLimitPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRateLimitPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRateLimitPolicyResponse) GetBody() *DescribeRateLimitPolicyResponseBody {
	return s.Body
}

func (s *DescribeRateLimitPolicyResponse) SetHeaders(v map[string]*string) *DescribeRateLimitPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeRateLimitPolicyResponse) SetStatusCode(v int32) *DescribeRateLimitPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRateLimitPolicyResponse) SetBody(v *DescribeRateLimitPolicyResponseBody) *DescribeRateLimitPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeRateLimitPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
