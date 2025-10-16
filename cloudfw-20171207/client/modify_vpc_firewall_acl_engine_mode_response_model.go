// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallAclEngineModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallAclEngineModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallAclEngineModeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallAclEngineModeResponseBody) *ModifyVpcFirewallAclEngineModeResponse
	GetBody() *ModifyVpcFirewallAclEngineModeResponseBody
}

type ModifyVpcFirewallAclEngineModeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallAclEngineModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallAclEngineModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallAclEngineModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallAclEngineModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallAclEngineModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallAclEngineModeResponse) GetBody() *ModifyVpcFirewallAclEngineModeResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallAclEngineModeResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallAclEngineModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallAclEngineModeResponse) SetStatusCode(v int32) *ModifyVpcFirewallAclEngineModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallAclEngineModeResponse) SetBody(v *ModifyVpcFirewallAclEngineModeResponseBody) *ModifyVpcFirewallAclEngineModeResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallAclEngineModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
