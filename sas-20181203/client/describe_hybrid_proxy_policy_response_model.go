// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHybridProxyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHybridProxyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHybridProxyPolicyResponseBody) *DescribeHybridProxyPolicyResponse
	GetBody() *DescribeHybridProxyPolicyResponseBody
}

type DescribeHybridProxyPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHybridProxyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHybridProxyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHybridProxyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHybridProxyPolicyResponse) GetBody() *DescribeHybridProxyPolicyResponseBody {
	return s.Body
}

func (s *DescribeHybridProxyPolicyResponse) SetHeaders(v map[string]*string) *DescribeHybridProxyPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeHybridProxyPolicyResponse) SetStatusCode(v int32) *DescribeHybridProxyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHybridProxyPolicyResponse) SetBody(v *DescribeHybridProxyPolicyResponseBody) *DescribeHybridProxyPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeHybridProxyPolicyResponse) Validate() error {
	return dara.Validate(s)
}
