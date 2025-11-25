// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreatedNatFirewallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCreatedNatFirewallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCreatedNatFirewallResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCreatedNatFirewallResponseBody) *DescribeCreatedNatFirewallResponse
	GetBody() *DescribeCreatedNatFirewallResponseBody
}

type DescribeCreatedNatFirewallResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCreatedNatFirewallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCreatedNatFirewallResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreatedNatFirewallResponse) GoString() string {
	return s.String()
}

func (s *DescribeCreatedNatFirewallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCreatedNatFirewallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCreatedNatFirewallResponse) GetBody() *DescribeCreatedNatFirewallResponseBody {
	return s.Body
}

func (s *DescribeCreatedNatFirewallResponse) SetHeaders(v map[string]*string) *DescribeCreatedNatFirewallResponse {
	s.Headers = v
	return s
}

func (s *DescribeCreatedNatFirewallResponse) SetStatusCode(v int32) *DescribeCreatedNatFirewallResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCreatedNatFirewallResponse) SetBody(v *DescribeCreatedNatFirewallResponseBody) *DescribeCreatedNatFirewallResponse {
	s.Body = v
	return s
}

func (s *DescribeCreatedNatFirewallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
