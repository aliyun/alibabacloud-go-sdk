// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallVswitchResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallVswitchResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallVswitchResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallVswitchResourcesResponseBody) *DescribeFirewallVswitchResourcesResponse
	GetBody() *DescribeFirewallVswitchResourcesResponseBody
}

type DescribeFirewallVswitchResourcesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallVswitchResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallVswitchResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVswitchResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVswitchResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallVswitchResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallVswitchResourcesResponse) GetBody() *DescribeFirewallVswitchResourcesResponseBody {
	return s.Body
}

func (s *DescribeFirewallVswitchResourcesResponse) SetHeaders(v map[string]*string) *DescribeFirewallVswitchResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponse) SetStatusCode(v int32) *DescribeFirewallVswitchResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponse) SetBody(v *DescribeFirewallVswitchResourcesResponseBody) *DescribeFirewallVswitchResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallVswitchResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
