// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallTimeTopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatFirewallTimeTopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatFirewallTimeTopResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatFirewallTimeTopResponseBody) *DescribeNatFirewallTimeTopResponse
	GetBody() *DescribeNatFirewallTimeTopResponseBody
}

type DescribeNatFirewallTimeTopResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallTimeTopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallTimeTopResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallTimeTopResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallTimeTopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatFirewallTimeTopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatFirewallTimeTopResponse) GetBody() *DescribeNatFirewallTimeTopResponseBody {
	return s.Body
}

func (s *DescribeNatFirewallTimeTopResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallTimeTopResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallTimeTopResponse) SetStatusCode(v int32) *DescribeNatFirewallTimeTopResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallTimeTopResponse) SetBody(v *DescribeNatFirewallTimeTopResponseBody) *DescribeNatFirewallTimeTopResponse {
	s.Body = v
	return s
}

func (s *DescribeNatFirewallTimeTopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
