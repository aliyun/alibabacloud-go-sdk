// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6EgressOnlyRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpv6EgressOnlyRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpv6EgressOnlyRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpv6EgressOnlyRulesResponseBody) *DescribeIpv6EgressOnlyRulesResponse
	GetBody() *DescribeIpv6EgressOnlyRulesResponseBody
}

type DescribeIpv6EgressOnlyRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpv6EgressOnlyRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpv6EgressOnlyRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6EgressOnlyRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpv6EgressOnlyRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpv6EgressOnlyRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpv6EgressOnlyRulesResponse) GetBody() *DescribeIpv6EgressOnlyRulesResponseBody {
	return s.Body
}

func (s *DescribeIpv6EgressOnlyRulesResponse) SetHeaders(v map[string]*string) *DescribeIpv6EgressOnlyRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponse) SetStatusCode(v int32) *DescribeIpv6EgressOnlyRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponse) SetBody(v *DescribeIpv6EgressOnlyRulesResponseBody) *DescribeIpv6EgressOnlyRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponse) Validate() error {
	return dara.Validate(s)
}
