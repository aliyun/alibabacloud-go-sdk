// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallVSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFirewallVSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFirewallVSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFirewallVSwitchResponseBody) *DescribeFirewallVSwitchResponse
	GetBody() *DescribeFirewallVSwitchResponseBody
}

type DescribeFirewallVSwitchResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallVSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallVSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFirewallVSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFirewallVSwitchResponse) GetBody() *DescribeFirewallVSwitchResponseBody {
	return s.Body
}

func (s *DescribeFirewallVSwitchResponse) SetHeaders(v map[string]*string) *DescribeFirewallVSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallVSwitchResponse) SetStatusCode(v int32) *DescribeFirewallVSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallVSwitchResponse) SetBody(v *DescribeFirewallVSwitchResponseBody) *DescribeFirewallVSwitchResponse {
	s.Body = v
	return s
}

func (s *DescribeFirewallVSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
