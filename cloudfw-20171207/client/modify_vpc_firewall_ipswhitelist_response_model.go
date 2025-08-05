// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallIPSWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallIPSWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallIPSWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallIPSWhitelistResponseBody) *ModifyVpcFirewallIPSWhitelistResponse
	GetBody() *ModifyVpcFirewallIPSWhitelistResponseBody
}

type ModifyVpcFirewallIPSWhitelistResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallIPSWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallIPSWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallIPSWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) GetBody() *ModifyVpcFirewallIPSWhitelistResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallIPSWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) SetStatusCode(v int32) *ModifyVpcFirewallIPSWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) SetBody(v *ModifyVpcFirewallIPSWhitelistResponseBody) *ModifyVpcFirewallIPSWhitelistResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
