// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpControlPolicyItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpControlPolicyItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpControlPolicyItemsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpControlPolicyItemsResponseBody) *DescribeIpControlPolicyItemsResponse
	GetBody() *DescribeIpControlPolicyItemsResponseBody
}

type DescribeIpControlPolicyItemsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpControlPolicyItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpControlPolicyItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpControlPolicyItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpControlPolicyItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpControlPolicyItemsResponse) GetBody() *DescribeIpControlPolicyItemsResponseBody {
	return s.Body
}

func (s *DescribeIpControlPolicyItemsResponse) SetHeaders(v map[string]*string) *DescribeIpControlPolicyItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpControlPolicyItemsResponse) SetStatusCode(v int32) *DescribeIpControlPolicyItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponse) SetBody(v *DescribeIpControlPolicyItemsResponseBody) *DescribeIpControlPolicyItemsResponse {
	s.Body = v
	return s
}

func (s *DescribeIpControlPolicyItemsResponse) Validate() error {
	return dara.Validate(s)
}
