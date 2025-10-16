// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallPrecheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcFirewallPrecheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcFirewallPrecheckResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcFirewallPrecheckResponseBody) *CreateVpcFirewallPrecheckResponse
	GetBody() *CreateVpcFirewallPrecheckResponseBody
}

type CreateVpcFirewallPrecheckResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallPrecheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallPrecheckResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallPrecheckResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallPrecheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcFirewallPrecheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcFirewallPrecheckResponse) GetBody() *CreateVpcFirewallPrecheckResponseBody {
	return s.Body
}

func (s *CreateVpcFirewallPrecheckResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallPrecheckResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallPrecheckResponse) SetStatusCode(v int32) *CreateVpcFirewallPrecheckResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallPrecheckResponse) SetBody(v *CreateVpcFirewallPrecheckResponseBody) *CreateVpcFirewallPrecheckResponse {
	s.Body = v
	return s
}

func (s *CreateVpcFirewallPrecheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
