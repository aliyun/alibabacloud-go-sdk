// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsFirewallPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDnsFirewallPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDnsFirewallPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDnsFirewallPolicyResponseBody) *DeleteDnsFirewallPolicyResponse
	GetBody() *DeleteDnsFirewallPolicyResponseBody
}

type DeleteDnsFirewallPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDnsFirewallPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDnsFirewallPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsFirewallPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteDnsFirewallPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDnsFirewallPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDnsFirewallPolicyResponse) GetBody() *DeleteDnsFirewallPolicyResponseBody {
	return s.Body
}

func (s *DeleteDnsFirewallPolicyResponse) SetHeaders(v map[string]*string) *DeleteDnsFirewallPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteDnsFirewallPolicyResponse) SetStatusCode(v int32) *DeleteDnsFirewallPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDnsFirewallPolicyResponse) SetBody(v *DeleteDnsFirewallPolicyResponseBody) *DeleteDnsFirewallPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteDnsFirewallPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
